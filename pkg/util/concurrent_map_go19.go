//go:build go1.9
// +build go1.9

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

import "sync"

type MapItem struct {
	Key   interface{}
	Value interface{}
}

type ConcurrentMap struct {
	mapper    sync.Map
	fetchLock sync.RWMutex
}

func (cm *ConcurrentMap) Put(key, val interface{}) { _ = "STUB: not implemented"; return }

func (cm *ConcurrentMap) PutIfAbsent(key, val interface{}) (exist interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func (cm *ConcurrentMap) Fetch(key interface{}, f func() (interface{}, error)) (v interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cm *ConcurrentMap) Get(key interface{}) (val interface{}, b bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (cm *ConcurrentMap) Remove(key interface{}) { _ = "STUB: not implemented"; return }

func (cm *ConcurrentMap) Clear() { _ = "STUB: not implemented"; return }

func (cm *ConcurrentMap) Size() (s int) { _ = "STUB: not implemented"; return 0 }

func (cm *ConcurrentMap) ForEach(f func(item MapItem) (next bool)) {
	_ = "STUB: not implemented"
	return
}

func NewConcurrentMap(_ int) *ConcurrentMap { _ = "STUB: not implemented"; return nil }
