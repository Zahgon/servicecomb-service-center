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

package etcd

import (
	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
)

// State implements kvstore.State.
// State does service pkg with etcd as it's cache.
type State struct {
	kvstore.Cacher
	kvstore.Indexer
}

func (se *State) Run() { _ = "STUB: not implemented"; return }

func (se *State) Stop() { _ = "STUB: not implemented"; return }

func (se *State) Ready() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func NewEtcdState(name string, cfg *kvstore.Options) *State { _ = "STUB: not implemented"; return nil }
