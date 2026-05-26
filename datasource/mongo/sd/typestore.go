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

// Package sd provides a TypeStore to manage the implementations of sd package, see types.go
package sd

import (
	"context"

	"github.com/apache/servicecomb-service-center/pkg/util"
	"github.com/go-chassis/foundation/gopool"
)

var store = &TypeStore{}

func init() {
	store.Initialize()
}

type TypeStore struct {
	caches    util.ConcurrentMap
	ready     chan struct{}
	goroutine *gopool.Pool
	isClose   bool
}

func (s *TypeStore) Initialize() { _ = "STUB: not implemented"; return }

func (s *TypeStore) Run() { _ = "STUB: not implemented"; return }

func (s *TypeStore) store(ctx context.Context) {
	_ = "STUB: not implemented"
	// new all types
	return
}

func (s *TypeStore) autoClearCache(ctx context.Context) { _ = "STUB: not implemented"; return }

func (s *TypeStore) getOrCreateCache(t string) *MongoCacher { _ = "STUB: not implemented"; return nil }

func (s *TypeStore) Stop() { _ = "STUB: not implemented"; return }

func (s *TypeStore) Ready() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (s *TypeStore) TypeCacher(id string) *MongoCacher { _ = "STUB: not implemented"; return nil }
func (s *TypeStore) Service() *MongoCacher             { _ = "STUB: not implemented"; return nil }
func (s *TypeStore) Instance() *MongoCacher            { _ = "STUB: not implemented"; return nil }
func (s *TypeStore) Dep() *MongoCacher                 { _ = "STUB: not implemented"; return nil }

func Store() *TypeStore { _ = "STUB: not implemented"; return nil }
