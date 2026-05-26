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

package dump

import (
	"github.com/apache/servicecomb-service-center/version"
	"github.com/go-chassis/cari/discovery"
)

type Getter interface {
	ForEach(i func(i int, v *KV) bool)
}

type Setter interface {
	SetValue(v *KV)
}

type MicroserviceSlice []*Microservice
type MicroserviceIndexSlice []*MicroserviceIndex
type MicroserviceAliasSlice []*MicroserviceAlias
type TagSlice []*Tag
type MicroServiceDependencyRuleSlice []*MicroServiceDependencyRule
type SummarySlice []*Summary
type InstanceSlice []*Instance

func (s *MicroserviceSlice) ForEach(f func(i int, v *KV) bool) { _ = "STUB: not implemented"; return }

func (s *MicroserviceIndexSlice) ForEach(f func(i int, v *KV) bool) {
	_ = "STUB: not implemented"
	return
}

func (s *MicroserviceAliasSlice) ForEach(f func(i int, v *KV) bool) {
	_ = "STUB: not implemented"
	return
}

func (s *TagSlice) ForEach(f func(i int, v *KV) bool) { _ = "STUB: not implemented"; return }

func (s *MicroServiceDependencyRuleSlice) ForEach(f func(i int, v *KV) bool) {
	_ = "STUB: not implemented"
	return
}

func (s *SummarySlice) ForEach(f func(i int, v *KV) bool) { _ = "STUB: not implemented"; return }

func (s *InstanceSlice) ForEach(f func(i int, v *KV) bool) { _ = "STUB: not implemented"; return }

func (s *MicroserviceSlice) SetValue(v *KV)               { _ = "STUB: not implemented"; return }
func (s *MicroserviceIndexSlice) SetValue(v *KV)          { _ = "STUB: not implemented"; return }
func (s *MicroserviceAliasSlice) SetValue(v *KV)          { _ = "STUB: not implemented"; return }
func (s *TagSlice) SetValue(v *KV)                        { _ = "STUB: not implemented"; return }
func (s *MicroServiceDependencyRuleSlice) SetValue(v *KV) { _ = "STUB: not implemented"; return }

func (s *SummarySlice) SetValue(v *KV)  { _ = "STUB: not implemented"; return }
func (s *InstanceSlice) SetValue(v *KV) { _ = "STUB: not implemented"; return }

func NewMicroservice(kv *KV) *Microservice { _ = "STUB: not implemented"; return nil }

func NewMicroserviceIndex(kv *KV) *MicroserviceIndex { _ = "STUB: not implemented"; return nil }

func NewMicroserviceAlias(kv *KV) *MicroserviceAlias { _ = "STUB: not implemented"; return nil }

func NewTag(kv *KV) *Tag { _ = "STUB: not implemented"; return nil }
func NewMicroServiceDependencyRule(kv *KV) *MicroServiceDependencyRule {
	_ = "STUB: not implemented"
	return nil
}

func NewSummary(kv *KV) *Summary   { _ = "STUB: not implemented"; return nil }
func NewInstance(kv *KV) *Instance { _ = "STUB: not implemented"; return nil }

type Cache struct {
	Microservices   MicroserviceSlice               `json:"services,omitempty"`
	Indexes         MicroserviceIndexSlice          `json:"serviceIndexes,omitempty"`
	Aliases         MicroserviceAliasSlice          `json:"serviceAliases,omitempty"`
	Tags            TagSlice                        `json:"serviceTags,omitempty"`
	DependencyRules MicroServiceDependencyRuleSlice `json:"dependencyRules,omitempty"`
	Summaries       SummarySlice                    `json:"summaries,omitempty"`
	Instances       InstanceSlice                   `json:"instances,omitempty"`
}

type KV struct {
	Key         string      `json:"key"`
	Rev         int64       `json:"rev"`
	Value       interface{} `json:"-"`
	ClusterName string      `json:"cluster"`
}

type Microservice struct {
	*KV
	Value *discovery.MicroService `json:"value,omitempty"`
}

type MicroserviceIndex struct {
	*KV
	Value string `json:"value,omitempty"`
}

type MicroserviceAlias struct {
	*KV
	Value string `json:"value,omitempty"`
}

type MicroServiceDependencyRule struct {
	*KV
	Value *discovery.MicroServiceDependency `json:"value,omitempty"`
}

type Summary struct {
	*KV
	Value string `json:"value,omitempty"`
}

type Tag struct {
	*KV
	Value map[string]string `json:"value,omitempty"`
}

type Instance struct {
	*KV
	Value *discovery.MicroServiceInstance `json:"value,omitempty"`
}

type Request struct {
	Options []string
}

type Response struct {
	Response *discovery.Response `json:"-"`
	Info     *version.Set        `json:"info,omitempty"`
	Cache    *Cache              `json:"cache,omitempty"`
}

type WatchInstanceChangedEvent struct {
	Action   string        `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Service  *Microservice `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
	Instance *Instance     `protobuf:"bytes,3,opt,name=instance" json:"instance,omitempty"`
}
