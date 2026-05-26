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

var (
	NullCache  = &nullCache{}
	NullCacher = &nullCacher{}
)

type nullCache struct {
}

func (n *nullCache) Name() string                           { _ = "STUB: not implemented"; return "" }
func (n *nullCache) Size() int                              { _ = "STUB: not implemented"; return 0 }
func (n *nullCache) Get(_ string) *KeyValue                 { _ = "STUB: not implemented"; return nil }
func (n *nullCache) GetAll(_ *[]*KeyValue) int              { _ = "STUB: not implemented"; return 0 }
func (n *nullCache) GetPrefix(_ string, _ *[]*KeyValue) int { _ = "STUB: not implemented"; return 0 }
func (n *nullCache) ForEach(_ func(k string, v *KeyValue) (next bool)) {
	_ = "STUB: not implemented"
	return
}
func (n *nullCache) Put(_ string, _ *KeyValue) { _ = "STUB: not implemented"; return }
func (n *nullCache) Remove(_ string)           { _ = "STUB: not implemented"; return }
func (n *nullCache) MarkDirty()                { _ = "STUB: not implemented"; return }
func (n *nullCache) Dirty() bool               { _ = "STUB: not implemented"; return false }
func (n *nullCache) Clear()                    { _ = "STUB: not implemented"; return }

type nullCacher struct {
}

func (n *nullCacher) Cache() CacheReader { _ = "STUB: not implemented"; return *new(CacheReader) }
