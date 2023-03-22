// Copyright 2022 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package mq

import (
	"context"

	"github.com/pingcap/errors"
	"github.com/pingcap/log"
	"sdbflow/cdc/contextutil"
	"sdbflow/cdc/model"
	"sdbflow/cdc/sink/codec/builder"
	"sdbflow/cdc/sink/codec/common"
	mqv1 "sdbflow/cdc/sink/mq"
	"sdbflow/cdc/sink/mq/dispatcher"
	"sdbflow/cdc/sink/mq/manager"
	"sdbflow/cdc/sinkv2/eventsink"
	"sdbflow/cdc/sinkv2/eventsink/mq/dmlproducer"
	"sdbflow/cdc/sinkv2/metrics"
	"sdbflow/cdc/sinkv2/tablesink/state"
	"sdbflow/pkg/config"
	cerror "sdbflow/pkg/errors"
	"sdbflow/pkg/sink"
	"go.uber.org/zap"
)

// Assert EventSink[E event.TableEvent] implementation
var _ eventsink.EventSink[*model.RowChangedEvent] = (*dmlSink)(nil)

// dmlSink is the mq sink.
// It will send the events to the MQ system.
type dmlSink struct {
	// id indicates this sink belongs to which processor(changefeed).
	id model.ChangeFeedID
	// protocol indicates the protocol used by this sink.
	protocol config.Protocol

	worker *worker
	// eventRouter used to route events to the right topic and partition.
	eventRouter *dispatcher.EventRouter
	// topicManager used to manage topics.
	// It is also responsible for creating topics.
	topicManager manager.TopicManager
}

func newSink(ctx context.Context,
	producer dmlproducer.DMLProducer,
	topicManager manager.TopicManager,
	eventRouter *dispatcher.EventRouter,
	encoderConfig *common.Config,
	encoderConcurrency int,
	errCh chan error,
) (*dmlSink, error) {
	changefeedID := contextutil.ChangefeedIDFromCtx(ctx)

	encoderBuilder, err := builder.NewEventBatchEncoderBuilder(ctx, encoderConfig)
	if err != nil {
		return nil, cerror.WrapError(cerror.ErrKafkaInvalidConfig, err)
	}

	statistics := metrics.NewStatistics(ctx, sink.RowSink)
	worker := newWorker(changefeedID, encoderConfig.Protocol,
		encoderBuilder, encoderConcurrency, producer, statistics)
	s := &dmlSink{
		id:           changefeedID,
		protocol:     encoderConfig.Protocol,
		worker:       worker,
		eventRouter:  eventRouter,
		topicManager: topicManager,
	}

	// Spawn a goroutine to send messages by the worker.
	go func() {
		if err := s.worker.run(ctx); err != nil && errors.Cause(err) != context.Canceled {
			select {
			case <-ctx.Done():
				return
			case errCh <- err:
			default:
				log.Error("Error channel is full in DML sink",
					zap.String("namespace", changefeedID.Namespace),
					zap.String("changefeed", changefeedID.ID),
					zap.Error(err))
			}
		}
	}()

	return s, nil
}

// WriteEvents writes events to the sink.
// This is an asynchronously and thread-safe method.
func (s *dmlSink) WriteEvents(rows ...*eventsink.RowChangeCallbackableEvent) error {
	for _, row := range rows {
		if row.GetTableSinkState() != state.TableSinkSinking {
			// The table where the event comes from is in stopping, so it's safe
			// to drop the event directly.
			row.Callback()
			continue
		}
		topic := s.eventRouter.GetTopicForRowChange(row.Event)
		partitionNum, err := s.topicManager.GetPartitionNum(topic)
		if err != nil {
			return errors.Trace(err)
		}
		partition := s.eventRouter.GetPartitionForRowChange(row.Event, partitionNum)
		// This never be blocked because this is an unbounded channel.
		s.worker.msgChan.In() <- mqEvent{
			key: mqv1.TopicPartitionKey{
				Topic: topic, Partition: partition,
			},
			rowEvent: row,
		}
	}

	return nil
}

// Close closes the sink.
func (s *dmlSink) Close() error {
	s.worker.close()
	return nil
}
