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

package task

import (
	"context"
	"sync"
	"time"

	"github.com/apache/servicecomb-service-center/syncer/service/event"

	carisync "github.com/go-chassis/cari/sync"
)

const (
	defaultInternal = 2 * time.Second

	heartbeatInternal = 15 * time.Second
	taskTTL           = 30
	taskName          = "load--handle-task"
)

func Work() { _ = "STUB: not implemented"; return }

func work() { _ = "STUB: not implemented"; return }

// Manager defines task manager, transfer task to event, and send event to event manager
type Manager interface {
	LoadAndHandleTask(ctx context.Context)
	UpdateResultTask(ctx context.Context)
}

type ManagerOption func(*managerOptions)

type managerOptions struct {
	internal    time.Duration
	operator    Operator
	eventSender event.Sender
}

func toManagerOptions(os ...ManagerOption) *managerOptions { _ = "STUB: not implemented"; return nil }

func ManagerInternal(i time.Duration) ManagerOption {
	_ = "STUB: not implemented"
	return *new(ManagerOption)
}

func EventSender(e event.Sender) ManagerOption {
	_ = "STUB: not implemented"
	return *new(ManagerOption)
}

func ManagerOperator(l Operator) ManagerOption {
	_ = "STUB: not implemented"
	return *new(ManagerOption)
}

func NewManager(os ...ManagerOption) Manager { _ = "STUB: not implemented"; return *new(Manager) }

type manager struct {
	internal time.Duration
	ticker   *time.Ticker

	toHandleTasks []*carisync.Task

	isClosing bool
	result    chan *event.Result
	cache     sync.Map

	operator    Operator
	eventSender event.Sender
}

// Operator define task operator, to list tasks and delete task
type Operator interface {
	ListTasks(ctx context.Context) ([]*carisync.Task, error)
	DeleteTask(ctx context.Context, t *carisync.Task) error
}

func (m *manager) LoadAndHandleTask(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *manager) Close() { _ = "STUB: not implemented"; return }

func (m *manager) ListTasks(ctx context.Context) ([]*carisync.Task, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *manager) DeleteTask(ctx context.Context, t *carisync.Task) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) UpdateResultTask(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *manager) closeUpdateTasks() { _ = "STUB: not implemented"; return }

func (m *manager) handleResult(res *event.Result) { _ = "STUB: not implemented"; return }

func (m *manager) handleTasks(sts syncTasks) { _ = "STUB: not implemented"; return }

func toEvent(task *carisync.Task, result chan<- *event.Result) *event.Event {
	_ = "STUB: not implemented"
	return nil
}
