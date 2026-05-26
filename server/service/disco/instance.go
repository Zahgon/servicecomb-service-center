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

package disco

import (
	"context"
	"sync"
	"time"

	"github.com/apache/servicecomb-service-center/datasource"
	pb "github.com/go-chassis/cari/discovery"
)

const (
	defaultMinInterval = 5 * time.Second
	defaultMinTimes    = 3
)

var (
	once          sync.Once
	propertiesMap map[string]string
)

func getInnerProperties() map[string]string { _ = "STUB: not implemented"; return nil }

func RegisterInstance(ctx context.Context, in *pb.RegisterInstanceRequest) (*pb.RegisterInstanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// instance util
func populateInstanceDefaultValue(ctx context.Context, instance *pb.MicroServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

// 这里应该根据租约计时
// Health check对象仅用于呈现服务健康检查逻辑，如果CHECK_BY_PLATFORM类型，表明由sidecar代发心跳，实例120s超时

func appendInnerPropertiesToInstance(instance *pb.MicroServiceInstance) {
	_ = "STUB: not implemented"
	return
}

func UnregisterInstance(ctx context.Context, in *pb.UnregisterInstanceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func SendHeartbeat(ctx context.Context, in *pb.HeartbeatRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// append the inner properties

func appendInnerProperties(ctx context.Context, serviceID string, instanceID string) error {
	_ = "STUB: not implemented"
	return nil
}

func shouldAppendInnerProperties(instance *pb.MicroServiceInstance) bool {
	_ = "STUB: not implemented"
	return false
}

func SendManyHeartbeat(ctx context.Context, in *pb.HeartbeatSetRequest) (*pb.HeartbeatSetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetInstance(ctx context.Context, in *pb.GetOneInstanceRequest) (*pb.GetOneInstanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListInstance(ctx context.Context, in *pb.GetInstancesRequest) (*pb.GetInstancesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindInstances(ctx context.Context, in *pb.FindInstancesRequest) (*pb.FindInstancesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindManyInstances(ctx context.Context, request *pb.BatchFindInstancesRequest) (*pb.BatchFindInstancesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// find services

// find instance

func batchFindServices(ctx context.Context, request *pb.BatchFindInstancesRequest) (*pb.BatchFindResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func batchFindInstances(ctx context.Context, request *pb.BatchFindInstancesRequest) (*pb.BatchFindResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// can not find the shared provider instances

func parseError(err error) (int32, string, error) { _ = "STUB: not implemented"; return 0, "", nil }

func AppendFindResponse(ctx context.Context, index int64, resp *pb.Response, instances []*pb.MicroServiceInstance,
	updatedResult *[]*pb.FindResult, notModifiedResult *[]int64, failedResult **pb.FindFailedResult) {
	_ = "STUB: not implemented"
	return
}

func PutInstance(ctx context.Context, in *pb.RegisterInstanceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func PutInstanceStatus(ctx context.Context, in *pb.UpdateInstanceStatusRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func PutInstanceProperties(ctx context.Context, in *pb.UpdateInstancePropsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ClusterHealth(ctx context.Context) (*pb.GetInstancesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkInstanceQuota(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func InstanceUsage(ctx context.Context, request *pb.GetServiceCountRequest) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func UpdateManyInstanceStatus(ctx context.Context, match *datasource.MatchPolicy, status string) error {
	_ = "STUB: not implemented"
	return nil
}
