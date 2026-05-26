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

package queue

import (
	"context"

	"github.com/go-chassis/foundation/gopool"
)

const (
	eventQueueSize = 1000
)

type Worker interface {
	Handle(ctx context.Context, obj interface{})
}

type Task struct {
	Payload interface{}
	// Async can let workers handle this task concurrently, but
	// it will make this task unordered
	Async bool
}

type TaskQueue struct {
	Workers []Worker

	taskCh    chan Task
	goroutine *gopool.Pool
}

// AddWorker is the method to add Worker
func (q *TaskQueue) AddWorker(w Worker) { _ = "STUB: not implemented"; return }

// Add is the method to add task in queue, one task will be handled by all workers
func (q *TaskQueue) Add(t Task) { _ = "STUB: not implemented"; return }

func (q *TaskQueue) dispatch(ctx context.Context, w Worker, obj interface{}) {
	_ = "STUB: not implemented"
	return

	// Do is the method to trigger workers handle the task immediately
}

func (q *TaskQueue) Do(ctx context.Context, task Task) { _ = "STUB: not implemented"; return }

// Run is the method to start a goroutine to pull and handle tasks from queue
func (q *TaskQueue) Run() { _ = "STUB: not implemented"; return }

// Stop is the method to stop the workers gracefully
func (q *TaskQueue) Stop() { _ = "STUB: not implemented"; return }

func NewTaskQueue(size int) *TaskQueue { _ = "STUB: not implemented"; return nil }
