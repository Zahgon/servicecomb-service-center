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

package heartbeatcache

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/apache/servicecomb-service-center/datasource/mongo/model"
)

const (
	DefaultTTL              = 30
	defaultCacheCapacity    = 10000
	defaultWorkNum          = 10
	defaultTimeout          = 10
	instanceCheckerInternal = 1 * time.Second
	ctxTimeout              = 5 * time.Second
)

var ErrHeartbeatTimeout = errors.New("heartbeat task waiting for processing timeout. ")

var (
	once sync.Once
	cfg  CacheConfig
)

type CacheConfig struct {
	CacheChan              chan *InstanceHeartbeatInfo
	InstanceHeartbeatStore *cache.Cache
	WorkerNum              int
	HeartbeatTaskTimeout   int
}

func Configuration() *CacheConfig { _ = "STUB: not implemented"; return nil }

func (c *CacheConfig) AddHeartbeatTask(serviceID string, instanceID string, ttl int32) error {
	_ = "STUB: not implemented"
	// Unassigned setting default value is 30s
	return nil
}

func (c *CacheConfig) RemoveCacheInstance(instanceID string) { _ = "STUB: not implemented"; return }

func cleanInstance(ctx context.Context, serviceID string, instanceID string) error {
	_ = "STUB: not implemented"
	return nil
}

func removeDBInstance(ctx context.Context, serviceID string, instanceID string) error {
	_ = "STUB: not implemented"
	return nil
}

func findInstance(ctx context.Context, serviceID string, instanceID string) (*model.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func updateInstance(ctx context.Context, serviceID string, instanceID string) error {
	_ = "STUB: not implemented"
	return nil
}

func isOutDate(refreshTime time.Time, ttl int32) bool { _ = "STUB: not implemented"; return false }
