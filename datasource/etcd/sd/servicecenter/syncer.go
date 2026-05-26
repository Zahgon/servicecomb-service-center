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
	"sync"

	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
	"github.com/apache/servicecomb-service-center/pkg/dump"
)

var (
	syncer     *Syncer
	syncerOnce sync.Once
)

type Syncer struct {
	Client *SCClientAggregate

	cachers map[kvstore.Type]*Cacher
}

func (c *Syncer) Initialize() { _ = "STUB: not implemented"; return }

func (c *Syncer) Sync(ctx context.Context) { _ = "STUB: not implemented"; return }

// microservice

// microservice meta

// instance

func (c *Syncer) check(local *Cacher, remote dump.Getter, skipClusters map[string]error) {
	_ = "STUB: not implemented"
	return
}

func (c *Syncer) checkWithConflictHandleFunc(local *Cacher, remote dump.Getter, skipClusters map[string]error,
	conflictHandleFunc func(origin *dump.KV, conflict dump.Getter, index int)) {
	_ = "STUB: not implemented"
	return
}

// because the result of the remote return may contain the same data as
// the local cache of the current SC. So we need to ignore it and
// prevent the aggregation result from increasing.

// if connect to some cluster failed, then skip to notify changes
// of these clusters to prevent publish the wrong changes events of kvs.

func (c *Syncer) skipHandleFunc(_ *dump.KV, _ dump.Getter, _ int) {
	_ = "STUB: not implemented"
	return
}

func (c *Syncer) logConflictFunc(origin *dump.KV, conflict dump.Getter, index int) {
	_ = "STUB: not implemented"
	return
}

func (c *Syncer) loop(ctx context.Context) { _ = "STUB: not implemented"; return }

// TODO support watching sc

// unsafe
func (c *Syncer) AddCacher(t kvstore.Type, cacher *Cacher) { _ = "STUB: not implemented"; return }

func (c *Syncer) Run() { _ = "STUB: not implemented"; return }

func GetOrCreateSyncer() *Syncer { _ = "STUB: not implemented"; return nil }
