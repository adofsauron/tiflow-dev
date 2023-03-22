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

package integration

import (
	"context"

	pb "sdbflow/engine/enginepb"
	"sdbflow/engine/framework/model"
	"sdbflow/engine/pkg/client"
	"sdbflow/engine/pkg/externalresource/broker"
	"sdbflow/engine/pkg/externalresource/manager"
	resModel "sdbflow/engine/pkg/externalresource/model"
)

type resourceClientStub struct {
	service *manager.Service
}

func (c *resourceClientStub) CreateResource(
	ctx context.Context,
	req *pb.CreateResourceRequest,
) error {
	_, err := c.service.CreateResource(ctx, req)
	return err
}

func (c *resourceClientStub) QueryResource(
	ctx context.Context,
	req *pb.QueryResourceRequest,
) (*pb.QueryResourceResponse, error) {
	return c.service.QueryResource(ctx, req)
}

func (c *resourceClientStub) RemoveResource(
	ctx context.Context,
	req *pb.RemoveResourceRequest,
) error {
	_, err := c.service.RemoveResource(ctx, req)
	return err
}

type executorClientStub struct {
	client.ExecutorClient
	brk broker.Broker
}

func (c *executorClientStub) RemoveResource(
	ctx context.Context,
	creatorWorkerID model.WorkerID,
	resourceID resModel.ResourceID,
) error {
	_, err := c.brk.RemoveResource(ctx, &pb.RemoveLocalResourceRequest{
		ResourceId: resourceID,
		WorkerId:   creatorWorkerID,
	})
	return err
}
