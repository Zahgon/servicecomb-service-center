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

package cache

import (
	"context"

	"github.com/apache/servicecomb-service-center/pkg/cache"
	pb "github.com/go-chassis/cari/discovery"
)

type InstancesFilter struct {
}

func (f *InstancesFilter) Name(ctx context.Context, _ *cache.Node) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *InstancesFilter) Init(ctx context.Context, parent *cache.Node) (node *cache.Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *InstancesFilter) Find(ctx context.Context, parent *cache.Node) (
	instances []*pb.MicroServiceInstance, rev string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// can not find by instanceKey.ServiceID after pre-filters init

func (f *InstancesFilter) findInstances(ctx context.Context, domainProject, serviceID, instanceID string, maxRevs []int64, counts []int64) (instances []*pb.MicroServiceInstance, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *InstancesFilter) FindInstances(ctx context.Context, domainProject string, instanceKey *pb.HeartbeatSetElement) (instances []*pb.MicroServiceInstance, rev string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (f *InstancesFilter) BatchFindInstances(ctx context.Context, domainProject string, serviceIDs []string) (instances []*pb.MicroServiceInstance, rev string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}
