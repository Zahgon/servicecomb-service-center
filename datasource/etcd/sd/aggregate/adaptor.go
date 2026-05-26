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

package aggregate

import (
	"github.com/apache/servicecomb-service-center/datasource/etcd/state"
	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
)

// Aggregator implements state.State.
// Aggregator is an aggregator of multi Adaptors, and it aggregates all the
// Adaptors' data as it's result.
type Aggregator struct {
	// Indexer searches data from all the adapters
	kvstore.Indexer
	Type     kvstore.Type
	Adaptors []state.State
}

// Cache gets all the adapters' cache
func (as *Aggregator) Cache() kvstore.CacheReader {
	_ = "STUB: not implemented"
	return *new(kvstore.CacheReader)
}

func (as *Aggregator) Run() { _ = "STUB: not implemented"; return }

func (as *Aggregator) Stop() { _ = "STUB: not implemented"; return }

func (as *Aggregator) Ready() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func getLogConflictFunc(t kvstore.Type) func(origin, conflict *kvstore.KeyValue) {
	_ = "STUB: not implemented"
	return nil
}

func NewAggregator(t kvstore.Type, cfg *kvstore.Options) *Aggregator {
	_ = "STUB: not implemented"
	return nil
}

// create and get all plugin instances
