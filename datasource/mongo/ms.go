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

package mongo

import (
	"context"

	"github.com/go-chassis/cari/discovery"
	"github.com/go-chassis/cari/pkg/errsvc"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/apache/servicecomb-service-center/datasource"
	"github.com/apache/servicecomb-service-center/datasource/mongo/model"
)

const baseTen = 10

type MetadataManager struct {
	// InstanceTTL options
	InstanceTTL int64
}

func (ds *MetadataManager) RegisterService(ctx context.Context, request *discovery.CreateServiceRequest) (*discovery.CreateServiceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the service unique index in table is (serviceId/serviceEnv,serviceAppid,servicename,serviceVersion)

// serviceid conflict with the service in the database

func createServiceTxn(ctx context.Context, request *discovery.CreateServiceRequest, domain string, project string,
	service *discovery.MicroService) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) ListService(ctx context.Context, _ *discovery.GetServicesRequest) (*discovery.GetServicesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) ListApp(ctx context.Context, request *discovery.GetAppsRequest) (*discovery.GetAppsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) GetService(ctx context.Context, request *discovery.GetServiceRequest) (*discovery.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) ExistServiceByID(ctx context.Context, request *discovery.GetExistenceByIDRequest) (*discovery.GetExistenceByIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) ExistService(ctx context.Context, request *discovery.GetExistenceRequest) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// 约定多个时，取较新版本

func (ds *MetadataManager) FindService(ctx context.Context, request *discovery.MicroServiceKey) (*discovery.GetServicesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) UnregisterService(ctx context.Context, request *discovery.DeleteServiceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// 强制删除，则与该服务相关的信息删除，非强制删除： 如果作为该被依赖（作为provider，提供服务,且不是只存在自依赖）或者存在实例，则不能删除

//todo wait for dep interface

func deleteServiceTxn(ctx context.Context, err error, serviceID string, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) PutServiceProperties(ctx context.Context, request *discovery.UpdateServicePropsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func updateServiceTxn(ctx context.Context, request *discovery.UpdateServicePropsRequest, filter bson.M, updateFilter bson.M) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) GetOverview(ctx context.Context, _ *discovery.GetServicesRequest) (
	*discovery.Statistics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) PutManyTags(ctx context.Context, request *discovery.AddServiceTagsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) ListTag(ctx context.Context, request *discovery.GetServiceTagsRequest) (*discovery.GetServiceTagsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) PutTag(ctx context.Context, request *discovery.UpdateServiceTagRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) DeleteManyTags(ctx context.Context, request *discovery.DeleteServiceTagsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) GetSchema(ctx context.Context, request *discovery.GetSchemaRequest) (*discovery.GetSchemaResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) GetAllSchemas(ctx context.Context, request *discovery.GetAllSchemaRequest) (*discovery.GetAllSchemaResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) ExistSchema(ctx context.Context, request *discovery.GetExistenceRequest) (*discovery.GetExistenceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) DeleteSchema(ctx context.Context, request *discovery.DeleteSchemaRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) ModifySchema(ctx context.Context, request *discovery.ModifySchemaRequest) (*discovery.ModifySchemaResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) ModifySchemas(ctx context.Context, request *discovery.ModifySchemasRequest) (*discovery.ModifySchemasResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) modifySchemas(ctx context.Context, service *discovery.MicroService, schemas []*discovery.Schema) error {
	_ = "STUB: not implemented"
	return nil
}

// modifySchema will be modified in the following cases
// 1.service have no relation --> update the schema && update the service
// 2.service is editable && service have relation with the schema --> update the schema
// 3.service is editable && service have no relation with the schema --> update the schema && update the service
// 4.service can't edit && service have relation with the schema && schema summary not exist --> update the schema
func (ds *MetadataManager) modifySchema(ctx context.Context, serviceID string, schema *discovery.Schema) *errsvc.Error {
	_ = "STUB: not implemented"
	return nil
}

// Instance management
func (ds *MetadataManager) RegisterInstance(ctx context.Context,
	request *discovery.RegisterInstanceRequest) (*discovery.RegisterInstanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if queueSize is more than 0 and channel is not full, then do fast register instance

// fast register, just add instance to channel and batch register them later

func sendEvent(ctx context.Context, action string, resourceType string, resource interface{}) {
	_ = "STUB: not implemented"
	return
}

func RegisterInstanceSingle(ctx context.Context, request *discovery.RegisterInstanceRequest,
	isUserDefinedID bool) (*discovery.RegisterInstanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RegisterInstanceBatch(ctx context.Context, events []*InstanceRegisterEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func preProcessRegister(ctx context.Context, instance *discovery.MicroServiceInstance,
	isUserDefinedID bool) (*discovery.RegisterInstanceResponse, bool, error) {
	_ = "STUB: not implemented"
	// 允许自定义 id
	return nil, false, nil
}

func sendHeartbeatInstead(ctx context.Context, instance *discovery.MicroServiceInstance) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// register a new one

func (ds *MetadataManager) ExistInstance(ctx context.Context, request *discovery.MicroServiceInstanceKey) (*discovery.GetExistenceByIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetInstance returns instance under the current domain
func (ds *MetadataManager) GetInstance(ctx context.Context, request *discovery.GetOneInstanceRequest) (*discovery.GetOneInstanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use explicit instanceId to query

// for gRPC

// TODO support gRPC output context

func (ds *MetadataManager) ListInstance(ctx context.Context, request *discovery.GetInstancesRequest) (*discovery.GetInstancesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for gRPC

func (ds *MetadataManager) ListManyInstances(ctx context.Context, _ *discovery.GetAllInstancesRequest) (*discovery.GetAllInstancesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindInstances returns instances under the specified domain
func (ds *MetadataManager) FindInstances(ctx context.Context, request *discovery.FindInstancesRequest) (*discovery.FindInstancesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) PutInstance(ctx context.Context, request *discovery.RegisterInstanceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) PutInstanceStatus(ctx context.Context, request *discovery.UpdateInstanceStatusRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// todo finish get instance

func (ds *MetadataManager) PutInstanceProperties(ctx context.Context, request *discovery.UpdateInstancePropsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) UnregisterInstance(ctx context.Context, request *discovery.UnregisterInstanceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) SendHeartbeat(ctx context.Context, request *discovery.HeartbeatRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) SendManyHeartbeat(ctx context.Context, request *discovery.HeartbeatSetRequest) (*discovery.HeartbeatSetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func registryInstance(ctx context.Context, request *discovery.RegisterInstanceRequest) (*discovery.RegisterInstanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// need to complete the instance offline function in time, so you need to check the heartbeat after registering the instance

func registryInstances(ctx context.Context, instances []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *MetadataManager) findSharedServiceInstance(ctx context.Context, request *discovery.FindInstancesRequest, provider *discovery.MicroServiceKey, rev string) (*discovery.FindInstancesResponse, error) {
	_ = "STUB: not implemented"

	// it means the shared micro-services must be the same env with SC.
	return nil, nil
}

// for gRPC

// TODO support gRPC output context

func (ds *MetadataManager) findInstance(ctx context.Context, request *discovery.FindInstancesRequest, provider *discovery.MicroServiceKey, rev string) (*discovery.FindInstancesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// provider is not a shared micro-service,
// only allow shared micro-service instances found request different domains.

// add dependency queue

// for gRPC

func (ds *MetadataManager) reshapeProviderKey(ctx context.Context, provider *discovery.MicroServiceKey, providerID string) (*discovery.MicroServiceKey, error) {
	_ = "STUB: not implemented"
	//维护version的规则,service name 可能是别名，所以重新获取
	return nil, nil
}

func AddServiceVersionRule(ctx context.Context, domainProject string, consumer *discovery.MicroService, provider *discovery.MicroServiceKey) error {
	_ = "STUB: not implemented"
	return nil
}

func DependencyRuleExist(ctx context.Context, provider *discovery.MicroServiceKey, consumer *discovery.MicroServiceKey) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func DependencyRuleExistUtil(ctx context.Context, key bson.M, target *discovery.MicroServiceKey) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func KeepAliveLease(ctx context.Context, request *discovery.HeartbeatRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func getHeartbeatFunc(ctx context.Context, _ string, instancesHbRst chan<- *discovery.InstanceHbRst, element *discovery.HeartbeatSetElement) func(context.Context) {
	_ = "STUB: not implemented"
	return nil
}

func filterServices(ctx context.Context, key *discovery.MicroServiceKey) ([]*model.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterServiceIDs(ctx context.Context, consumerID string, tags []string, services []*model.Service) []string {
	_ = "STUB: not implemented"
	return nil
}

func filterTags(services []*model.Service, tags []string) []*model.Service {
	_ = "STUB: not implemented"
	return nil
}

func filterAccess(ctx context.Context, consumerID string, services []*model.Service) []*model.Service {
	_ = "STUB: not implemented"
	return nil
}

func accessible(ctx context.Context, consumerID string, providerID string) *errsvc.Error {
	_ = "STUB: not implemented"
	return nil
}

// 跨应用权限

func allowAcrossDimension(ctx context.Context, providerService *model.Service, consumerService *model.Service) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteDependencyForDeleteService(domainProject string, _ string, service *discovery.MicroServiceKey) error {
	_ = "STUB: not implemented"
	return nil
}

func formatRevision(consumerServiceID string, instances []*discovery.MicroServiceInstance) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ds *MetadataManager) Statistics(ctx context.Context, withShared bool) (*discovery.Statistics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *MetadataManager) UpdateManyInstanceStatus(_ context.Context, _ *datasource.MatchPolicy, _ string) error {
	_ = "STUB: not implemented"
	return nil
}
