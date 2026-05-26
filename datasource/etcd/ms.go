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

package etcd

import (
	"context"

	pb "github.com/go-chassis/cari/discovery"
	"github.com/go-chassis/cari/pkg/errsvc"

	"github.com/apache/servicecomb-service-center/datasource"
	"github.com/apache/servicecomb-service-center/datasource/etcd/cache"
)

type MetadataManager struct {
	// InstanceTTL options
	InstanceTTL int64
}

// RegisterService implement:
// 1. capsule request to etcd kv format
// 2. invoke etcd client to store data
// 3. check etcd-client response && construct createServiceResponse
func (ds *MetadataManager) RegisterService(ctx context.Context, request *pb.CreateServiceRequest) (
	*pb.CreateServiceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 产生全局service id

// internal error?

func (ds *MetadataManager) ListService(ctx context.Context, _ *pb.GetServicesRequest) (
	*pb.GetServicesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) GetService(ctx context.Context, request *pb.GetServiceRequest) (
	*pb.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) GetOverview(ctx context.Context, _ *pb.GetServicesRequest) (
	*pb.Statistics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) ListApp(ctx context.Context, request *pb.GetAppsRequest) (*pb.GetAppsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) ExistServiceByID(ctx context.Context, request *pb.GetExistenceByIDRequest) (*pb.GetExistenceByIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) ExistService(ctx context.Context, request *pb.GetExistenceRequest) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// 约定多个时，取较新版本

func (ds *MetadataManager) FindService(ctx context.Context, request *pb.MicroServiceKey) (*pb.GetServicesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) PutServiceProperties(ctx context.Context, request *pb.UpdateServicePropsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Set key file

func (ds *MetadataManager) RegisterInstance(ctx context.Context, request *pb.RegisterInstanceRequest) (
	*pb.RegisterInstanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) registerInstance(ctx context.Context, request *pb.RegisterInstanceRequest) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

//允许自定义id

//先以domain/project的方式组装

// build the request options

func sendEvent(ctx context.Context, action string, resourceType string, resource interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ds *MetadataManager) calcInstanceTTL(instance *pb.MicroServiceInstance) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (ds *MetadataManager) sendHeartbeatInstead(ctx context.Context, instance *pb.MicroServiceInstance) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// keep alive the lease ttl
// there are two reasons for sending a heartbeat here:
// 1. request the scenario the instance has been removed,
//    the cast of registration operation can be reduced.
// 2. request the self-protection scenario, the instance is unhealthy
//    and needs to be re-registered.

// register a new one

func (ds *MetadataManager) ExistInstance(ctx context.Context, request *pb.MicroServiceInstanceKey) (*pb.GetExistenceByIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) GetInstance(ctx context.Context, request *pb.GetOneInstanceRequest) (
	*pb.GetOneInstanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for gRPC

func (ds *MetadataManager) ListInstance(ctx context.Context, request *pb.GetInstancesRequest) (*pb.GetInstancesResponse,
	error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for gRPC

func (ds *MetadataManager) FindInstances(ctx context.Context, request *pb.FindInstancesRequest) (*pb.FindInstancesResponse,
	error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) findInstance(ctx context.Context, request *pb.FindInstancesRequest,
	provider *pb.MicroServiceKey, rev string) (*pb.FindInstancesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// provider is not a shared micro-service,
// only allow shared micro-service instances found request different domains.

// cache

// add dependency queue

func (ds *MetadataManager) findSharedServiceInstance(ctx context.Context, request *pb.FindInstancesRequest,
	provider *pb.MicroServiceKey, rev string) (*pb.FindInstancesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// it means the shared micro-services must be the same env with SC.

// cache

func (ds *MetadataManager) genFindResult(ctx context.Context, oldRev string, item *cache.VersionRuleCacheItem) (
	*pb.FindInstancesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for gRPC

// TODO support gRPC output context

func (ds *MetadataManager) reshapeProviderKey(ctx context.Context, provider *pb.MicroServiceKey, providerID string) (
	*pb.MicroServiceKey, error) {
	_ = "STUB: not implemented"
	// service name 可能是别名，所以重新获取
	return nil, nil
}

// just compatible to old version

func (ds *MetadataManager) PutInstance(ctx context.Context, request *pb.RegisterInstanceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) PutInstanceStatus(ctx context.Context, request *pb.UpdateInstanceStatusRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) PutInstanceProperties(ctx context.Context, request *pb.UpdateInstancePropsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) SendManyHeartbeat(ctx context.Context, request *pb.HeartbeatSetRequest) (*pb.HeartbeatSetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) UnregisterInstance(ctx context.Context, request *pb.UnregisterInstanceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) SendHeartbeat(ctx context.Context, request *pb.HeartbeatRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) ListManyInstances(ctx context.Context, _ *pb.GetAllInstancesRequest) (*pb.GetAllInstancesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) ModifySchemas(ctx context.Context, request *pb.ModifySchemasRequest) (
	*pb.ModifySchemasResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) ModifySchema(ctx context.Context, request *pb.ModifySchemaRequest) (
	*pb.ModifySchemaResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) ExistSchema(ctx context.Context, request *pb.GetExistenceRequest) (
	*pb.GetExistenceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) GetSchema(ctx context.Context, request *pb.GetSchemaRequest) (*pb.GetSchemaResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) GetAllSchemas(ctx context.Context, request *pb.GetAllSchemaRequest) (
	*pb.GetAllSchemaResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) DeleteSchema(ctx context.Context, request *pb.DeleteSchemaRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) PutManyTags(ctx context.Context, request *pb.AddServiceTagsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// service id存在性校验

func (ds *MetadataManager) ListTag(ctx context.Context, request *pb.GetServiceTagsRequest) (*pb.GetServiceTagsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) PutTag(ctx context.Context, request *pb.UpdateServiceTagRequest) error {
	_ = "STUB: not implemented"
	return nil
}

//check if the tag exists

func (ds *MetadataManager) DeleteManyTags(ctx context.Context, request *pb.DeleteServiceTagsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// the capacity of tags may be 0

func (ds *MetadataManager) modifySchemas(ctx context.Context, domainProject string, service *pb.MicroService,
	schemas []*pb.Schema) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) modifySchema(ctx context.Context, serviceID string, schema *pb.Schema) *errsvc.Error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) UnregisterService(ctx context.Context, request *pb.DeleteServiceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// 强制删除，则与该服务相关的信息删除，非强制删除： 如果作为该被依赖（作为provider，提供服务,且不是只存在自依赖）或者存在实例，则不能删除

//删除依赖规则

//删除schemas

//删除tags

//删除instances

//删除实例

func (ds *MetadataManager) Statistics(ctx context.Context, withShared bool) (*pb.Statistics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) UpdateManyInstanceStatus(ctx context.Context, match *datasource.MatchPolicy, status string) error {
	_ = "STUB: not implemented"
	return nil
}

//更新状态
