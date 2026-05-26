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

	"github.com/apache/servicecomb-service-center/datasource"
	"github.com/apache/servicecomb-service-center/datasource/mongo/model"
)

type DepManager struct {
}

func (ds *DepManager) ListConsumers(ctx context.Context, request *discovery.GetDependenciesRequest) (*discovery.GetProDependenciesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *DepManager) ListProviders(ctx context.Context, request *discovery.GetDependenciesRequest) (*discovery.GetConDependenciesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *DepManager) PutDependencies(ctx context.Context, dependencyInfos []*discovery.ConsumerDependency, override bool) error {
	_ = "STUB: not implemented"
	return nil
}

func updateDepTxn(ctx context.Context, dependencyInfos []*discovery.ConsumerDependency, override bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *DepManager) DependencyHandle(_ context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func syncDependencyRule(ctx context.Context, domainProject string, r *discovery.ConsumerDependency) error {
	_ = "STUB: not implemented"
	return nil
}

//var err error

// add mongo get dep here

func GetOldProviderRules(dep *datasource.Dependency) (*discovery.MicroServiceDependency, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func updateDeps(domainProject string, dep *datasource.Dependency) error {
	_ = "STUB: not implemented"
	return nil
}

func CleanUpDepRules(ctx context.Context, domainProject string) error {
	_ = "STUB: not implemented"
	return nil
}

func removeProviderRuleOfConsumer(ctx context.Context, domainProject string, cache map[*model.DelDepCacheKey]bool) error {
	_ = "STUB: not implemented"
	return nil
}

func removeProviderRuleKeys(ctx context.Context, domainProject string, cache map[*model.DelDepCacheKey]bool) error {
	_ = "STUB: not implemented"
	return nil
}

func GetDepRules(ctx context.Context, filter bson.M) ([]*model.DependencyRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func removeProviderDeps(ctx context.Context, depRule *model.DependencyRule, cache map[*model.DelDepCacheKey]bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func removeConsumerDeps(ctx context.Context, depRule *model.DependencyRule, cache map[*model.DelDepCacheKey]bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func TransferToMicroServiceDependency(ctx context.Context, filter bson.M) (*discovery.MicroServiceDependency, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
