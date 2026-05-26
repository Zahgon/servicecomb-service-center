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

package event

import (
	"context"
	"sync"

	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
	"github.com/apache/servicecomb-service-center/pkg/queue"
	pb "github.com/go-chassis/cari/discovery"
)

const depQueueLockKey = "/dep-queue"

// just for unit test
var testMux sync.Mutex

// DependencyEventHandler add or remove the service dependencies
// when user call find instance api or dependence operation api
type DependencyEventHandler struct {
	signals *queue.UniQueue
}

func (h *DependencyEventHandler) Type() kvstore.Type {
	_ = "STUB: not implemented"
	return *new(kvstore.Type)
}

func (h *DependencyEventHandler) OnEvent(evt kvstore.Event) { _ = "STUB: not implemented"; return }

func (h *DependencyEventHandler) notify() { _ = "STUB: not implemented"; return }

func (h *DependencyEventHandler) backoff(f func(), retries int) int {
	_ = "STUB: not implemented"
	return 0
}

func (h *DependencyEventHandler) tryWithBackoff(success func() error, backoff func(), retries int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h *DependencyEventHandler) eventLoop() { _ = "STUB: not implemented"; return }

// the events will lose, need to handle dependence records periodically

type DependencyEventHandlerResource struct {
	dep           *pb.ConsumerDependency
	kv            *kvstore.KeyValue
	domainProject string
}

func NewDependencyEventHandlerResource(dep *pb.ConsumerDependency, kv *kvstore.KeyValue, domainProject string) *DependencyEventHandlerResource {
	_ = "STUB: not implemented"
	return nil
}

func (h *DependencyEventHandler) Handle() error { _ = "STUB: not implemented"; return nil }

// maintain dependency rules.

func (h *DependencyEventHandler) dependencyRuleHandle(res interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *DependencyEventHandler) removeKV(ctx context.Context, kv *kvstore.KeyValue) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *DependencyEventHandler) CleanUp(domainProjects map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

func NewDependencyEventHandler() *DependencyEventHandler { _ = "STUB: not implemented"; return nil }
