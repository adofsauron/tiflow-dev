// Copyright 2020 PingCAP, Inc.
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

package server

import (
	"sdbflow/cdc/entry"
	"sdbflow/cdc/kv"
	"sdbflow/cdc/owner"
	"sdbflow/cdc/processor"
	"sdbflow/cdc/puller"
	redo "sdbflow/cdc/redo/common"
	"sdbflow/cdc/scheduler"
	sink "sdbflow/cdc/sink/metrics"
	"sdbflow/cdc/sink/mq/producer/kafka"
	sinkv2 "sdbflow/cdc/sinkv2/metrics"
	"sdbflow/cdc/sorter"
	dbsroter "sdbflow/cdc/sorter/db"
	"sdbflow/cdc/sorter/memory"
	"sdbflow/cdc/sorter/unified"
	"sdbflow/pkg/actor"
	"sdbflow/pkg/db"
	"sdbflow/pkg/etcd"
	"sdbflow/pkg/orchestrator"
	"sdbflow/pkg/p2p"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	tikvmetrics "github.com/tikv/client-go/v2/metrics"
)

var registry = prometheus.NewRegistry()

func init() {
	registry.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	registry.MustRegister(prometheus.NewGoCollector(
		collectors.WithGoCollections(collectors.GoRuntimeMemStatsCollection | collectors.GoRuntimeMetricsCollection)))

	initServerMetrics(registry)
	kv.InitMetrics(registry)
	puller.InitMetrics(registry)
	sink.InitMetrics(registry)
	sinkv2.InitMetrics(registry)
	entry.InitMetrics(registry)
	processor.InitMetrics(registry)
	owner.InitMetrics(registry)
	etcd.InitMetrics(registry)
	actor.InitMetrics(registry)
	orchestrator.InitMetrics(registry)
	p2p.InitMetrics(registry)
	sorter.InitMetrics(registry)
	memory.InitMetrics(registry)
	unified.InitMetrics(registry)
	dbsroter.InitMetrics(registry)
	redo.InitMetrics(registry)
	db.InitMetrics(registry)
	kafka.InitMetrics(registry)
	scheduler.InitMetrics(registry)
	// TiKV client metrics, including metrics about resolved and region cache.
	originalRegistry := prometheus.DefaultRegisterer
	prometheus.DefaultRegisterer = registry
	tikvmetrics.InitMetrics("ticdc", "tikvclient")
	tikvmetrics.RegisterMetrics()
	prometheus.DefaultRegisterer = originalRegistry
}
