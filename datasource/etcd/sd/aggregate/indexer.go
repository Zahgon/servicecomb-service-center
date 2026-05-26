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
	"context"

	"github.com/apache/servicecomb-service-center/datasource/etcd/state"
	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
	"github.com/little-cui/etcdadpt"
)

// AdaptorsIndexer implements kvstore.Indexer.
// AdaptorsIndexer is an aggregator of multi Indexers, and it aggregates all the
// Indexers' data as it's result.
type AdaptorsIndexer struct {
	Adaptors []state.State
}

// Search implements kvstore.Indexer#Search.
// AdaptorsIndexer ignores the errors during search to ensure availability, so
// it always searches successfully, no matter how many Adaptors are abnormal.
// But at the cost of that, AdaptorsIndexer doesn't guarantee the correctness
// of the search results.
func (i *AdaptorsIndexer) Search(ctx context.Context, opts ...etcdadpt.OpOption) (*kvstore.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Creditable implements kvstore.Indexer#Creditable.
// AdaptorsIndexer's search result's are not creditable as it ignores the
// errors. In other words, AdaptorsIndexer makes the best efforts to search
// data, but it does not ensure the correctness.
func (i *AdaptorsIndexer) Creditable() bool { _ = "STUB: not implemented"; return false }

func NewAdaptorsIndexer(as []state.State) *AdaptorsIndexer { _ = "STUB: not implemented"; return nil }

// AggregatorIndexer implements kvstore.Indexer.
// AggregatorIndexer consists of multi Indexers and it decides which Indexer to
// use based on it's mechanism.
type AggregatorIndexer struct {
	// CacheIndexer searches data from all the adaptors's cache.
	*kvstore.CacheIndexer
	// AdaptorsIndexer searches data from all the adaptors.
	AdaptorsIndexer kvstore.Indexer
	// LocalIndexer data from local adaptor.
	LocalIndexer kvstore.Indexer
}

// Search implements kvstore.Indexer#Search.
func (i *AggregatorIndexer) Search(ctx context.Context, opts ...etcdadpt.OpOption) (resp *kvstore.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// request with global param then do not use local indexer

// Creditable implements kvstore.Indexer#Creditable.
func (i *AggregatorIndexer) Creditable() bool { _ = "STUB: not implemented"; return false }

func NewAggregatorIndexer(as *Aggregator) *AggregatorIndexer { _ = "STUB: not implemented"; return nil }
