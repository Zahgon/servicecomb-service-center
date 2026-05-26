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
	"time"

	v1sync "github.com/apache/servicecomb-service-center/syncer/api/v1"
	"github.com/apache/servicecomb-service-center/syncer/service/replicator"
)

const (
	DefaultInternal    = 500 * time.Millisecond
	eventChanSize      = 1000
	batchEventChanSize = 100
	resChanSize        = 1000
	eventSliceSize     = 100
)

var m Manager

type Event struct {
	*v1sync.Event
	CanNotAbandon bool

	Result chan<- *Result
}

type Result struct {
	ID    string
	Data  *v1sync.Result
	Error error
}

func Work() { _ = "STUB: not implemented"; return }

func GetManager() Manager { _ = "STUB: not implemented"; return *new(Manager) }

type ManagerOption func(*managerOptions)

type managerOptions struct {
	internal   time.Duration
	replicator replicator.Replicator
}

func ManagerInternal(i time.Duration) ManagerOption {
	_ = "STUB: not implemented"
	return *new(ManagerOption)
}

func toManagerOptions(os ...ManagerOption) *managerOptions { _ = "STUB: not implemented"; return nil }

func Replicator(r replicator.Replicator) ManagerOption {
	_ = "STUB: not implemented"
	return *new(ManagerOption)
}

func NewManager(os ...ManagerOption) Manager { _ = "STUB: not implemented"; return *new(Manager) }

// Sender send events
type Sender interface {
	Send(et *Event)
}

// Manager manage events, including send events, handle events and handle result
type Manager interface {
	Sender

	HandleEvent()
	HandleResult()
}

type ManagerImpl struct {
	events      chan *Event
	batchEvents chan []*Event

	internal time.Duration
	ticker   *time.Ticker

	cache  sync.Map
	result chan *Result

	Replicator replicator.Replicator
}

func (e *ManagerImpl) Send(et *Event) { _ = "STUB: not implemented"; return }

func (e *ManagerImpl) checkThreshold(et *Event) bool { _ = "STUB: not implemented"; return false }

func (e *ManagerImpl) HandleResult() { _ = "STUB: not implemented"; return }

func (e *ManagerImpl) resultHandle(ctx context.Context) { _ = "STUB: not implemented"; return }

func (e *ManagerImpl) Close() { _ = "STUB: not implemented"; return }

type syncEvents []*Event

func (s syncEvents) Len() int { _ = "STUB: not implemented"; return 0 }

func (s syncEvents) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (s syncEvents) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (e *ManagerImpl) HandleEvent() { _ = "STUB: not implemented"; return }

func (e *ManagerImpl) readAndPackEvents(ctx context.Context) { _ = "STUB: not implemented"; return }

func (e *ManagerImpl) handleBatchEvents(ctx context.Context) { _ = "STUB: not implemented"; return }

func (e *ManagerImpl) handle(ctx context.Context, es syncEvents) { _ = "STUB: not implemented"; return }

// Send sends event to replicator
func Send(e *Event) { _ = "STUB: not implemented"; return }
