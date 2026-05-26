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

	"github.com/go-chassis/foundation/gopool"
)

const (
	initExecutorCount      = 1000
	removeExecutorInterval = 30 * time.Second
	initExecutorTTL        = 4
	executeInterval        = 1 * time.Second
	compactTimes           = 2
)

type executorWithTTL struct {
	*Executor
	TTL int64
}

type AsyncTaskService struct {
	executors map[string]*executorWithTTL
	goroutine *gopool.Pool
	lock      sync.RWMutex
	ready     chan struct{}
	isClose   bool
}

func (lat *AsyncTaskService) getOrNewExecutor(task Task) (s *Executor, isNew bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (lat *AsyncTaskService) Add(ctx context.Context, task Task) error {
	_ = "STUB: not implemented"
	return nil
}

// do immediately at first time

func (lat *AsyncTaskService) removeExecutor(key string) { _ = "STUB: not implemented"; return }

func (lat *AsyncTaskService) LatestHandled(key string) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

func (lat *AsyncTaskService) daemon(ctx context.Context) { _ = "STUB: not implemented"; return }

// non-blocked

func (lat *AsyncTaskService) Run() { _ = "STUB: not implemented"; return }

func (lat *AsyncTaskService) Stop() { _ = "STUB: not implemented"; return }

func (lat *AsyncTaskService) Ready() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (lat *AsyncTaskService) renew() { _ = "STUB: not implemented"; return }

func NewTaskService() Service { _ = "STUB: not implemented"; return *new(Service) }
