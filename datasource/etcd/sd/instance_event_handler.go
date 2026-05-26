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

package sd

import (
	"context"
	"sync"

	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
)

type deferItem struct {
	ReplayAfter int32 // in seconds
	event       kvstore.Event
}

type InstanceEventDeferHandler struct {
	Percent float64

	cache    kvstore.CacheReader
	once     sync.Once
	enabled  bool
	items    map[string]*deferItem
	evts     chan []kvstore.Event
	replayCh chan kvstore.Event
	resetCh  chan struct{}
}

func (iedh *InstanceEventDeferHandler) OnCondition(cache kvstore.CacheReader, evts []kvstore.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func (iedh *InstanceEventDeferHandler) recoverOrDefer(evt kvstore.Event) {
	_ = "STUB: not implemented"
	return
}

// return nil // no need to publish event to subscribers?

func (iedh *InstanceEventDeferHandler) HandleChan() <-chan kvstore.Event {
	_ = "STUB: not implemented"
	return nil
}

func (iedh *InstanceEventDeferHandler) check(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (iedh *InstanceEventDeferHandler) ReplayEvents() { _ = "STUB: not implemented"; return }

func (iedh *InstanceEventDeferHandler) replayEvent(evt kvstore.Event) {
	_ = "STUB: not implemented"
	return
}

func (iedh *InstanceEventDeferHandler) Reset() bool { _ = "STUB: not implemented"; return false }

func NewInstanceEventDeferHandler() *InstanceEventDeferHandler {
	_ = "STUB: not implemented"
	return nil
}
