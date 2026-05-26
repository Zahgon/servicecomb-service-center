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

	"github.com/go-chassis/cari/discovery"
	"github.com/little-cui/etcdadpt"
)

// Dependency contains dependency rules
type Dependency struct {
	DomainProject string
	// store the consumer Dependency from dep-queue object
	Consumer      *discovery.MicroServiceKey
	ProvidersRule []*discovery.MicroServiceKey
	// store the parsed rules from Dependency object
	DeleteDependencyRuleList []*discovery.MicroServiceKey
	CreateDependencyRuleList []*discovery.MicroServiceKey
}

func (dep *Dependency) removeConsumerOfProviderRule(ctx context.Context) ([]etcdadpt.OpOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//删除后，如果不存在依赖规则了，就删除该provider的依赖规则，如果有，则更新该依赖规则

func (dep *Dependency) addConsumerOfProviderRule(ctx context.Context) ([]etcdadpt.OpOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dep *Dependency) updateProvidersRuleOfConsumer(_ context.Context) ([]etcdadpt.OpOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Commit is dependent rule operations
func (dep *Dependency) Commit(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
