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

// Package state provides a State to manage the implementations of sd package, see types.go
package state

import (
	"context"
	"sync"

	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
	"github.com/go-chassis/foundation/gopool"
)

type Manager struct {
	Repository Repository
	Rev        int64

	states     map[kvstore.Type]State
	statesLock sync.RWMutex

	goroutine *gopool.Pool

	ready   chan struct{}
	isClose bool
}

func (s *Manager) Initialize() { _ = "STUB: not implemented"; return }

func (s *Manager) OnCacheEvent(evt kvstore.Event) { _ = "STUB: not implemented"; return }

func (s *Manager) InjectConfig(cfg *kvstore.Options) *kvstore.Options {
	_ = "STUB: not implemented"
	return nil
}

func (s *Manager) repo() Repository { _ = "STUB: not implemented"; return *new(Repository) }

func (s *Manager) getOrCreateState(t kvstore.Type) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (s *Manager) stopStates() { _ = "STUB: not implemented"; return }

func (s *Manager) Run() { _ = "STUB: not implemented"; return }

func (s *Manager) store(ctx context.Context) {
	_ = "STUB: not implemented"
	// new all types
	return
}

func (s *Manager) autoClearCache(ctx context.Context) { _ = "STUB: not implemented"; return }

func (s *Manager) Stop() { _ = "STUB: not implemented"; return }

func (s *Manager) Ready() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (s *Manager) States(id kvstore.Type) State { _ = "STUB: not implemented"; return *new(State) }
