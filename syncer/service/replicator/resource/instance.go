/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package resource

import (
	"context"

	v1sync "github.com/apache/servicecomb-service-center/syncer/api/v1"

	pb "github.com/go-chassis/cari/discovery"
)

const (
	Instance = "instance"
)

func NewInstance(e *v1sync.Event) Resource { _ = "STUB: not implemented"; return *new(Resource) }

type instance struct {
	event *v1sync.Event

	createInput *pb.RegisterInstanceRequest
	updateInput *pb.MicroServiceInstance
	deleteInput *pb.UnregisterInstanceRequest

	serviceID  string
	instanceID string

	service *pb.MicroService
	cur     *pb.MicroServiceInstance

	manager metadataManager
}

func (i *instance) loadInput() error { _ = "STUB: not implemented"; return nil }

type metadataManage struct {
}

func (m *metadataManage) RegisterService(ctx context.Context, request *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *metadataManage) GetService(ctx context.Context, in *pb.GetServiceRequest) (*pb.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *metadataManage) PutServiceProperties(ctx context.Context, request *pb.UpdateServicePropsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *metadataManage) UnregisterService(ctx context.Context, request *pb.DeleteServiceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *metadataManage) RegisterInstance(ctx context.Context, in *pb.RegisterInstanceRequest) (*pb.RegisterInstanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *metadataManage) SendHeartbeat(ctx context.Context, in *pb.HeartbeatRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *metadataManage) GetInstance(ctx context.Context, in *pb.GetOneInstanceRequest) (*pb.GetOneInstanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *metadataManage) PutInstance(ctx context.Context, in *pb.RegisterInstanceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *metadataManage) UnregisterInstance(ctx context.Context, in *pb.UnregisterInstanceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

type metadataManager interface {
	serviceManager
	instanceManager
}

type instanceManager interface {
	RegisterInstance(ctx context.Context, in *pb.RegisterInstanceRequest) (*pb.RegisterInstanceResponse, error)
	SendHeartbeat(ctx context.Context, in *pb.HeartbeatRequest) error
	GetInstance(ctx context.Context, in *pb.GetOneInstanceRequest) (*pb.GetOneInstanceResponse, error)
	PutInstance(ctx context.Context, in *pb.RegisterInstanceRequest) error
	UnregisterInstance(ctx context.Context, in *pb.UnregisterInstanceRequest) error
}

func (i *instance) LoadCurrentResource(ctx context.Context) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (i *instance) NeedOperate(ctx context.Context) *Result { _ = "STUB: not implemented"; return nil }

func (i *instance) FailHandle(ctx context.Context, code int32) (*v1sync.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *instance) CanDrop() bool { _ = "STUB: not implemented"; return false }

func (i *instance) Operate(ctx context.Context) *Result { _ = "STUB: not implemented"; return nil }

func (i *instance) CreateHandle(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (i *instance) UpdateHandle(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (i *instance) DeleteHandle(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
