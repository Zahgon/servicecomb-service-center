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

package kvstore

import (
	"sync"
)

// KvCache implements Cache.
// KvCache is dedicated to stores service discovery data,
// e.g. service, instance, lease.
type KvCache struct {
	Cfg   *Options
	name  string
	store map[string]map[string]*KeyValue
	rwMux sync.RWMutex
	dirty bool
}

func (c *KvCache) Name() string { _ = "STUB: not implemented"; return "" }

func (c *KvCache) Size() (l int) { _ = "STUB: not implemented"; return 0 }

func (c *KvCache) Get(key string) (v *KeyValue) { _ = "STUB: not implemented"; return nil }

func (c *KvCache) GetAll(arr *[]*KeyValue) (count int) { _ = "STUB: not implemented"; return 0 }

func (c *KvCache) GetPrefix(prefix string, arr *[]*KeyValue) (count int) {
	_ = "STUB: not implemented"
	return 0
}

func (c *KvCache) Put(key string, v *KeyValue) { _ = "STUB: not implemented"; return }

func (c *KvCache) Remove(key string) { _ = "STUB: not implemented"; return }

func (c *KvCache) MarkDirty() { _ = "STUB: not implemented"; return }

func (c *KvCache) Dirty() bool { _ = "STUB: not implemented"; return false }

func (c *KvCache) Clear() { _ = "STUB: not implemented"; return }

func (c *KvCache) ForEach(iter func(k string, v *KeyValue) (next bool)) {
	_ = "STUB: not implemented"
	return
}

func (c *KvCache) prefix(key string) string { _ = "STUB: not implemented"; return "" }

func (c *KvCache) getPrefixKey(arr *[]*KeyValue, prefix string) (count int) {
	_ = "STUB: not implemented"
	return 0
}

// TODO support sort option

func (c *KvCache) addPrefixKey(key string, val *KeyValue) { _ = "STUB: not implemented"; return }

// build parent index key and new child nodes

// override the value

func (c *KvCache) deletePrefixKey(key string) { _ = "STUB: not implemented"; return }

// remove parent which has no child

func NewKvCache(name string, cfg *Options) *KvCache { _ = "STUB: not implemented"; return nil }
