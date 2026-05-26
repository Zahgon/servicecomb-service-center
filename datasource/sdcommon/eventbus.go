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

package sdcommon

import (
	"context"
	"sync"
)

type EventBus struct {
	Cfg    ListWatchConfig
	Lw     ListWatch
	Bus    chan *ListWatchResp
	stopCh chan struct{}
	stop   bool
	mux    sync.Mutex
}

func (w *EventBus) ResourceEventBus() <-chan *ListWatchResp { _ = "STUB: not implemented"; return nil }

func (w *EventBus) process(_ context.Context) { _ = "STUB: not implemented"; return }

// timed out or exception

func (w *EventBus) sendEvent(resp *ListWatchResp) { _ = "STUB: not implemented"; return }

func (w *EventBus) Stop() { _ = "STUB: not implemented"; return }

func NewEventBus(lw ListWatch, cfg ListWatchConfig) *EventBus {
	_ = "STUB: not implemented"
	return nil
}
