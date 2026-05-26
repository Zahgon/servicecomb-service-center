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
	"go.mongodb.org/mongo-driver/bson"

	"github.com/apache/servicecomb-service-center/datasource/mongo/model"
)

type DependencyRelation struct {
	ctx           context.Context
	domainProject string
	consumer      *discovery.MicroService
	provider      *discovery.MicroService
}

type DependencyRelationFilterOpt struct {
	SameDomainProject bool
	NonSelf           bool
}

type DependencyRelationFilterOption func(opt DependencyRelationFilterOpt) DependencyRelationFilterOpt

func NewConsumerDependencyRelation(ctx context.Context, domainProject string, consumer *discovery.MicroService) *DependencyRelation {
	_ = "STUB: not implemented"
	return nil
}

func NewProviderDependencyRelation(ctx context.Context, domainProject string, provider *discovery.MicroService) *DependencyRelation {
	_ = "STUB: not implemented"
	return nil
}

func NewDependencyRelation(ctx context.Context, domainProject string, consumer *discovery.MicroService, provider *discovery.MicroService) *DependencyRelation {
	_ = "STUB: not implemented"
	return nil
}

func (dr *DependencyRelation) GetDependencyProviders(opts ...DependencyRelationFilterOption) ([]*discovery.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) GetDependencyConsumers(opts ...DependencyRelationFilterOption) ([]*discovery.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) GetDependencyConsumersOfProvider() ([]*discovery.MicroServiceKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) GetConsumerOfSameServiceNameAndAppID(provider *discovery.MicroServiceKey) ([]*discovery.MicroServiceKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) GetServiceByMicroServiceKey(service *discovery.MicroServiceKey) (*discovery.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getServiceKeysInDep(ctx context.Context, filter interface{}) ([]*model.DependencyRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) getProviderKeys() ([]*discovery.MicroServiceKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) parseDependencyRule(dependencyRule *discovery.MicroServiceKey) (serviceIDs []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) GetDependencyConsumerIds() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MicroServiceKeyFilter(key *discovery.MicroServiceKey) (bson.M, error) {
	_ = "STUB: not implemented"
	return *new(bson.M), nil
}

func FindServiceIds(ctx context.Context, key *discovery.MicroServiceKey, matchVersion bool) ([]string, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// service exist but version not matched

func serviceVersionFilter(ctx context.Context, version string, filter bson.D, matchVersion bool) ([]string, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func findServiceKeysByServiceName(ctx context.Context, key *discovery.MicroServiceKey, baseFilter bson.D, matchVersion bool) ([]string, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func findServiceKeysByAlias(ctx context.Context, key *discovery.MicroServiceKey, baseFilter bson.D, matchVersion bool) ([]string, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

type ServiceVersionFilter func(ctx context.Context, filter bson.D) ([]string, error)

func findServiceKeys(_ context.Context, version string, filter bson.D) (newFilter bson.D) {
	_ = "STUB: not implemented"
	return *new(bson.D)
}

func GetVersionService(ctx context.Context, m bson.D) (serviceIds []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WithSameDomainProject() DependencyRelationFilterOption {
	_ = "STUB: not implemented"
	return *new(DependencyRelationFilterOption)
}

func WithoutSelfDependency() DependencyRelationFilterOption {
	_ = "STUB: not implemented"
	return *new(DependencyRelationFilterOption)
}

func ToDependencyFilterOptions(in *discovery.GetDependenciesRequest) (opts []DependencyRelationFilterOption) {
	_ = "STUB: not implemented"
	return nil
}

func ToDependencyRelationFilterOpt(opts ...DependencyRelationFilterOption) (op DependencyRelationFilterOpt) {
	_ = "STUB: not implemented"
	return *new(DependencyRelationFilterOpt)
}

func GenerateConsumerDependencyRuleKey(domainProject string, in *discovery.MicroServiceKey) bson.M {
	_ = "STUB: not implemented"
	return *new(bson.M)
}

func GenerateProviderDependencyRuleKey(domainProject string, in *discovery.MicroServiceKey) bson.M {
	_ = "STUB: not implemented"
	return *new(bson.M)
}

func GenerateRuleKeyWithSameServiceNameAndAppID(serviceType string, domainProject string, in *discovery.MicroServiceKey) bson.M {
	_ = "STUB: not implemented"
	return *new(bson.M)
}

func GenerateServiceDependencyRuleKey(serviceType string, domainProject string, in *discovery.MicroServiceKey) bson.M {
	_ = "STUB: not implemented"
	return *new(bson.M)
}
