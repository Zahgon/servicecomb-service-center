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

package util

import (
	"context"

	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
	pb "github.com/go-chassis/cari/discovery"
	"github.com/go-chassis/cari/pkg/errsvc"
)

func GetLeaseID(ctx context.Context, domainProject string, serviceID string, instanceID string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetInstance(ctx context.Context, domainProject string, serviceID string, instanceID string) (*pb.MicroServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExistInstance(ctx context.Context, domainProject string, serviceID string, instanceID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func FormatRevision(revs, counts []int64) (s string) { _ = "STUB: not implemented"; return "" }

func GetAllInstancesOfOneService(ctx context.Context, domainProject string, serviceID string) ([]*pb.MicroServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetInstanceCountOfOneService(ctx context.Context, domainProject string, serviceID string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type EndpointIndexValue struct {
	ServiceID  string
	InstanceID string
}

func ParseEndpointIndexValue(value []byte) EndpointIndexValue {
	_ = "STUB: not implemented"
	return *new(EndpointIndexValue)
}

func DeleteServiceAllInstances(ctx context.Context, serviceID string) error {
	_ = "STUB: not implemented"
	return nil
}

func QueryServiceInstancesKvs(ctx context.Context, serviceID string, rev int64) ([]*kvstore.KeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UpdateInstance(ctx context.Context, domainProject string, instance *pb.MicroServiceInstance) *errsvc.Error {
	_ = "STUB: not implemented"
	return nil
}
