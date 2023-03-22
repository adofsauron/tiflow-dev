#!/bin/bash
# Copyright 2022 PingCAP, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# See the License for the specific language governing permissions and
# limitations under the License.

set -eu

cd "$(dirname "${BASH_SOURCE[0]}")"/..

MOCKGEN="tools/bin/mockgen"

if [ ! -f "$MOCKGEN" ]; then
	echo "${MOCKGEN} does not exist, please run 'make tools/bin/mockgen' first"
	exit 1
fi

# CDC mock
"$MOCKGEN" -source cdc/owner/owner.go -destination cdc/owner/mock/owner_mock.go
"$MOCKGEN" -source cdc/owner/status_provider.go -destination cdc/owner/mock/status_provider_mock.go
"$MOCKGEN" -source cdc/api/v2/api_helpers.go -destination cdc/api/v2/api_helpers_mock.go -package v2
"$MOCKGEN" -source pkg/etcd/etcd.go -destination pkg/etcd/mock/etcd_client_mock.go
"$MOCKGEN" -source cdc/processor/manager.go -destination cdc/processor/mock/manager_mock.go
"$MOCKGEN" -source cdc/capture/capture.go -destination cdc/capture/mock/capture_mock.go
"$MOCKGEN" -source pkg/cmd/factory/factory.go -destination pkg/cmd/factory/mock/factory_mock.go -package mock_factory
"$MOCKGEN" -source cdc/processor/sourcemanager/engine/engine.go -destination cdc/processor/sourcemanager/engine/mock/engine_mock.go

# DM mock
"$MOCKGEN" -package pbmock -destination dm/pbmock/dmmaster.go sdbflow/dm/pb MasterClient,MasterServer
"$MOCKGEN" -package pbmock -destination dm/pbmock/dmworker.go sdbflow/dm/pb WorkerClient,WorkerServer

# Engine mock
"$MOCKGEN" -package mock -destination engine/pkg/meta/mock/client_mock.go sdbflow/engine/pkg/meta/model KVClient
"$MOCKGEN" -package mock -destination engine/executor/server/mock/metastore_mock.go sdbflow/engine/executor/server MetastoreCreator
"$MOCKGEN" -package mock -destination engine/enginepb/mock/executor_mock.go sdbflow/engine/enginepb ExecutorServiceClient
"$MOCKGEN" -package mock -destination engine/enginepb/mock/broker_mock.go sdbflow/engine/enginepb BrokerServiceClient
"$MOCKGEN" -package mock -destination engine/enginepb/mock/resource_mock.go sdbflow/engine/enginepb ResourceManagerClient
"$MOCKGEN" -package mock -destination engine/pkg/httputil/mock/jobhttpclient_mock.go sdbflow/engine/pkg/httputil JobHTTPClient
"$MOCKGEN" -package mock -destination engine/servermaster/jobop/mock/joboperator_mock.go sdbflow/engine/servermaster/jobop JobOperator
"$MOCKGEN" -package mock -destination engine/pkg/election/mock/storage_mock.go sdbflow/engine/pkg/election Storage
"$MOCKGEN" -package mock -destination engine/pkg/election/mock/elector_mock.go sdbflow/engine/pkg/election Elector
"$MOCKGEN" -package mock -destination engine/pkg/orm/mock/client_mock.go sdbflow/engine/pkg/orm Client
"$MOCKGEN" -package mock -destination engine/servermaster/jobop/mock/backoffmanager_mock.go sdbflow/engine/servermaster/jobop BackoffManager
"$MOCKGEN" -package mock -source engine/pkg/rpcutil/checker.go -destination engine/pkg/rpcutil/mock/checker_mock.go
"$MOCKGEN" -package client -self_package sdbflow/engine/pkg/client \
	-destination engine/pkg/client/client_mock.go sdbflow/engine/pkg/client ExecutorClient,ServerMasterClient
