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

	pb "github.com/go-chassis/cari/discovery"
	"github.com/little-cui/etcdadpt"
)

func GetConsumerIds(ctx context.Context, domainProject string, provider *pb.MicroService) ([]string, error) {
	_ = "STUB: not implemented"
	// 查询所有consumer
	return nil, nil
}

func GetConsumers(ctx context.Context, domainProject string, provider *pb.MicroService,
	opts ...DependencyRelationFilterOption) ([]*pb.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetProviderIds(ctx context.Context, domainProject string, consumer *pb.MicroService) ([]string, error) {
	_ = "STUB: not implemented"
	// 查询所有provider
	return nil, nil
}

func GetProviders(ctx context.Context, domainProject string, consumer *pb.MicroService,
	opts ...DependencyRelationFilterOption) ([]*pb.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DependencyRuleExist(ctx context.Context, provider *pb.MicroServiceKey, consumer *pb.MicroServiceKey) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func DependencyRuleExistWithKey(ctx context.Context, key string, target *pb.MicroServiceKey) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

//删除之前的依赖

func AddServiceVersionRule(ctx context.Context, domainProject string, consumer *pb.MicroService, provider *pb.MicroServiceKey) error {
	_ = "STUB: not implemented"
	//创建依赖一致
	return nil
}

func TransferToMicroServiceDependency(ctx context.Context, key string) (*pb.MicroServiceDependency, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EqualServiceDependency(serviceA *pb.MicroServiceKey, serviceB *pb.MicroServiceKey) bool {
	_ = "STUB: not implemented"
	return false
}

func DiffServiceVersion(serviceA *pb.MicroServiceKey, serviceB *pb.MicroServiceKey) bool {
	_ = "STUB: not implemented"
	return false
}

func toString(in *pb.MicroServiceKey) string { _ = "STUB: not implemented"; return "" }

func parseAddOrUpdateRules(ctx context.Context, dep *Dependency) (createDependencyRuleList, existDependencyRuleList, deleteDependencyRuleList []*pb.MicroServiceKey) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func parseOverrideRules(ctx context.Context, dep *Dependency) (createDependencyRuleList, existDependencyRuleList, deleteDependencyRuleList []*pb.MicroServiceKey) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func syncDependencyRule(ctx context.Context, dep *Dependency, filter func(context.Context, *Dependency) (_, _, _ []*pb.MicroServiceKey)) error {
	_ = "STUB: not implemented"
	//更新consumer的providers的值,consumer的版本是确定的
	return nil
}

func AddDependencyRule(ctx context.Context, dep *Dependency) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateDependencyRule(ctx context.Context, dep *Dependency) error {
	_ = "STUB: not implemented"
	return nil
}

func IsNeedUpdate(services []*pb.MicroServiceKey, service *pb.MicroServiceKey) *pb.MicroServiceKey {
	_ = "STUB: not implemented"
	return nil
}

func ContainServiceDependency(services []*pb.MicroServiceKey, service *pb.MicroServiceKey) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func DeleteDependencyForDeleteService(domainProject string, serviceID string, service *pb.MicroServiceKey) (etcdadpt.OpOptions, error) {
	_ = "STUB: not implemented"
	return *new(etcdadpt.OpOptions), nil
}

func removeProviderRuleOfConsumer(ctx context.Context, domainProject string, cache map[string]bool) ([]etcdadpt.OpOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RemoveProviderRuleKeys(ctx context.Context, domainProject string, cache map[string]bool) ([]etcdadpt.OpOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CleanUpDependencyRules(ctx context.Context, domainProject string) error {
	_ = "STUB: not implemented"
	return nil
}
