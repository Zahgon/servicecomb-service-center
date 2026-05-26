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
	"github.com/apache/servicecomb-service-center/pkg/queue"
	"github.com/go-chassis/foundation/gopool"
)

type Executor struct {
	pool       *gopool.Pool
	tasks      *queue.UniQueue
	latestTask Task
}

func (s *Executor) AddTask(task Task) (err error) { _ = "STUB: not implemented"; return nil }

func (s *Executor) Execute() { _ = "STUB: not implemented"; return }

func (s *Executor) Close() { _ = "STUB: not implemented"; return }

func NewExecutor(pool *gopool.Pool, task Task) *Executor { _ = "STUB: not implemented"; return nil }
