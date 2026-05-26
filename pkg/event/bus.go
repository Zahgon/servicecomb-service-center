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

	"github.com/apache/servicecomb-service-center/pkg/queue"
	"github.com/apache/servicecomb-service-center/pkg/util"
)

// Bus can fire the event aync and dispatch events to subscriber according to subject
type Bus struct {
	*queue.TaskQueue

	name     string
	subjects *util.ConcurrentMap
}

func (bus *Bus) Name() string { _ = "STUB: not implemented"; return "" }

func (bus *Bus) Fire(evt Event) {
	_ = "STUB: not implemented"
	// TODO add option if queue is full
	return
}

func (bus *Bus) Handle(_ context.Context, payload interface{}) { _ = "STUB: not implemented"; return }

func (bus *Bus) fireAtOnce(evt Event) { _ = "STUB: not implemented"; return }

// else the evt will be discard

func (bus *Bus) Subjects(name string) *Poster { _ = "STUB: not implemented"; return nil }

func (bus *Bus) AddSubscriber(n Subscriber) { _ = "STUB: not implemented"; return }

func (bus *Bus) RemoveSubscriber(n Subscriber) { _ = "STUB: not implemented"; return }

func (bus *Bus) Clear() { _ = "STUB: not implemented"; return }

func NewBus(name string, queueSize int) *Bus { _ = "STUB: not implemented"; return nil }
