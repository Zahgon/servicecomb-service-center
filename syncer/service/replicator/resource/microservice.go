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
	Microservice = "service"
)

func NewMicroservice(e *v1sync.Event) Resource { _ = "STUB: not implemented"; return *new(Resource) }

type microservice struct {
	event *v1sync.Event

	createInput *pb.CreateServiceRequest
	updateInput *pb.UpdateServicePropsRequest
	deleteInput *pb.DeleteServiceRequest

	serviceID string

	cur *pb.MicroService

	manager serviceManager

	defaultFailHandler
}

type serviceManager interface {
	RegisterService(ctx context.Context, request *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error)
	GetService(ctx context.Context, in *pb.GetServiceRequest) (*pb.MicroService, error)
	PutServiceProperties(ctx context.Context, request *pb.UpdateServicePropsRequest) error
	UnregisterService(ctx context.Context, request *pb.DeleteServiceRequest) error
}

func (m *microservice) loadInput() error { _ = "STUB: not implemented"; return nil }

func (m *microservice) LoadCurrentResource(ctx context.Context) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (m *microservice) NeedOperate(ctx context.Context) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (m *microservice) CreateHandle(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *microservice) UpdateHandle(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *microservice) DeleteHandle(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *microservice) Operate(ctx context.Context) *Result { _ = "STUB: not implemented"; return nil }
