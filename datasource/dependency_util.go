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

package datasource

import (
	"context"

	"github.com/go-chassis/cari/discovery"
)

type Dependency struct {
	DomainProject string
	// store the consumer Dependency from dep-queue object
	Consumer      *discovery.MicroServiceKey
	ProvidersRule []*discovery.MicroServiceKey
	// store the parsed rules from Dependency object
	DeleteDependencyRuleList []*discovery.MicroServiceKey
	CreateDependencyRuleList []*discovery.MicroServiceKey
}

func ParamsChecker(consumerInfo *discovery.MicroServiceKey, providersInfo []*discovery.MicroServiceKey) error {
	_ = "STUB: not implemented"
	return nil
}

func toString(in *discovery.MicroServiceKey) string { _ = "STUB: not implemented"; return "" }

func ParseAddOrUpdateRules(_ context.Context, dep *Dependency, oldProviderRules *discovery.MicroServiceDependency) {
	_ = "STUB: not implemented"
	return
}

func ParseOverrideRules(_ context.Context, dep *Dependency, oldProviderRules *discovery.MicroServiceDependency) {
	_ = "STUB: not implemented"
	return
}

func setDep(dep *Dependency, createDependencyRuleList, existDependencyRuleList, deleteDependencyRuleList []*discovery.MicroServiceKey) {
	_ = "STUB: not implemented"
	return
}

func ContainServiceDependency(services []*discovery.MicroServiceKey, service *discovery.MicroServiceKey) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func EqualServiceDependency(serviceA *discovery.MicroServiceKey, serviceB *discovery.MicroServiceKey) bool {
	_ = "STUB: not implemented"
	return false
}

func IsNeedUpdate(services []*discovery.MicroServiceKey, service *discovery.MicroServiceKey) *discovery.MicroServiceKey {
	_ = "STUB: not implemented"
	return nil
}

func DiffServiceVersion(serviceA *discovery.MicroServiceKey, serviceB *discovery.MicroServiceKey) bool {
	_ = "STUB: not implemented"
	return false
}
