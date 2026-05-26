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

package servicecenter

import (
	"context"

	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
	"github.com/little-cui/etcdadpt"
)

// ClusterIndexer implements kvstore.Indexer.
// ClusterIndexer searches data from cache(firstly) and
// other service-centers(secondly).
type ClusterIndexer struct {
	*kvstore.CacheIndexer
	Client *SCClientAggregate
	Type   kvstore.Type
}

func (i *ClusterIndexer) Search(ctx context.Context, opts ...etcdadpt.OpOption) (resp *kvstore.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *ClusterIndexer) search(ctx context.Context, opts ...etcdadpt.OpOption) (r *kvstore.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *ClusterIndexer) searchSchemas(ctx context.Context, op etcdadpt.OpOptions) (*kvstore.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *ClusterIndexer) searchInstances(ctx context.Context, op etcdadpt.OpOptions) (r *kvstore.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Creditable implements kvstore.Indexer#Creditable.
// ClusterIndexer's search result's are not creditable as SCClientAggregate
// ignores sc clients' errors.
func (i *ClusterIndexer) Creditable() bool { _ = "STUB: not implemented"; return false }

func NewClusterIndexer(t kvstore.Type, cache kvstore.Cache) *ClusterIndexer {
	_ = "STUB: not implemented"
	return nil
}
