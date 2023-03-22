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

package factory

import (
	"context"
	"strings"

	"sdbflow/cdc/sink/mq/producer/kafka"
	"sdbflow/cdc/sinkv2/ddlsink"
	"sdbflow/cdc/sinkv2/ddlsink/blackhole"
	"sdbflow/cdc/sinkv2/ddlsink/cloudstorage"
	"sdbflow/cdc/sinkv2/ddlsink/mq"
	"sdbflow/cdc/sinkv2/ddlsink/mq/ddlproducer"
	"sdbflow/cdc/sinkv2/ddlsink/mysql"
	"sdbflow/pkg/config"
	cerror "sdbflow/pkg/errors"
	"sdbflow/pkg/sink"
	pmysql "sdbflow/pkg/sink/mysql"
)

// New creates a new ddlsink.DDLEventSink by schema.
func New(
	ctx context.Context,
	sinkURIStr string,
	cfg *config.ReplicaConfig,
) (ddlsink.DDLEventSink, error) {
	sinkURI, err := config.GetSinkURIAndAdjustConfigWithSinkURI(sinkURIStr, cfg)
	if err != nil {
		return nil, err
	}
	schema := strings.ToLower(sinkURI.Scheme)
	switch schema {
	case sink.KafkaScheme, sink.KafkaSSLScheme:
		return mq.NewKafkaDDLSink(ctx, sinkURI, cfg,
			kafka.NewAdminClientImpl, ddlproducer.NewKafkaDDLProducer)
	case sink.BlackHoleScheme:
		return blackhole.New(), nil
	case sink.MySQLSSLScheme, sink.MySQLScheme, sink.TiDBScheme, sink.TiDBSSLScheme:
		return mysql.NewMySQLDDLSink(ctx, sinkURI, cfg, pmysql.CreateMySQLDBConn)
	case sink.S3Scheme, sink.FileScheme, sink.GCSScheme, sink.GSScheme, sink.AzblobScheme, sink.AzureScheme, sink.CloudStorageNoopScheme:
		return cloudstorage.NewCloudStorageDDLSink(ctx, sinkURI)
	default:
		return nil,
			cerror.ErrSinkURIInvalid.GenWithStack("the sink scheme (%s) is not supported", schema)
	}
}
