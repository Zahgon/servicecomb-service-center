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

package cache

import (
	"context"
	"time"

	pb "github.com/go-chassis/cari/discovery"

	"github.com/apache/servicecomb-service-center/pkg/cache"
)

var FindInstances = &FindInstancesCache{
	Tree: cache.NewTree(cache.Configure().WithTTL(2 * time.Minute).WithMaxSize(DefaultCacheMaxSize))}

func init() {
	FindInstances.AddFilter(
		&ServiceFilter{},
		&VersionFilter{},
		&TagsFilter{},
		&AccessibleFilter{},
		&InstancesFilter{},
		&ConsistencyFilter{},
	)
}

type VersionRuleCacheItem struct {
	ServiceIds []string
	Instances  []*pb.MicroServiceInstance
	Rev        string

	broken bool
	queue  chan struct{}
}

func (vi *VersionRuleCacheItem) InitBrokenQueue() { _ = "STUB: not implemented"; return }

func (vi *VersionRuleCacheItem) BrokenWait() bool { _ = "STUB: not implemented"; return false }

func (vi *VersionRuleCacheItem) Broken() { _ = "STUB: not implemented"; return }

type FindInstancesCache struct {
	*cache.Tree
}

func (f *FindInstancesCache) Get(ctx context.Context, consumer *pb.MicroService, provider *pb.MicroServiceKey,
	tags []string, rev string) (*VersionRuleCacheItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FindInstancesCache) GetWithProviderID(ctx context.Context, consumer *pb.MicroService, provider *pb.MicroServiceKey,
	instanceKey *pb.HeartbeatSetElement, tags []string, rev string) (*VersionRuleCacheItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FindInstancesCache) Remove(provider *pb.MicroServiceKey) {
	_ = "STUB: not implemented"
	return
}
