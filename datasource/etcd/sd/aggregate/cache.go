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
	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
)

// Cache implements kvstore.CacheReader.
// Cache is a multi-CacheReader, it reads cache from all CacheReaders.
type Cache []kvstore.CacheReader

func (c Cache) Name() string { _ = "STUB: not implemented"; return "" }

func (c Cache) Size() (s int) { _ = "STUB: not implemented"; return 0 }

func (c Cache) Get(k string) (kv *kvstore.KeyValue) { _ = "STUB: not implemented"; return nil }

func (c Cache) GetAll(arr *[]*kvstore.KeyValue) (s int) { _ = "STUB: not implemented"; return 0 }

func (c Cache) GetPrefix(prefix string, arr *[]*kvstore.KeyValue) (s int) {
	_ = "STUB: not implemented"
	return 0
}

func (c Cache) append(tmp []*kvstore.KeyValue, arr *[]*kvstore.KeyValue,
	exists map[string]struct{}) (s int) {
	_ = "STUB: not implemented"
	return 0
}

func (c Cache) ForEach(iter func(k string, v *kvstore.KeyValue) (next bool)) {
	_ = "STUB: not implemented"
	return
}
