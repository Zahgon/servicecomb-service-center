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
	"sync"
)

// BusService is the daemon service to manage multiple type Bus
// and wrap handle methods of Bus
type BusService struct {
	// buses is the map of event handler, key is event source type
	buses   map[Type]*Bus
	mux     sync.RWMutex
	isClose bool
}

func (s *BusService) newBus(t Type) *Bus { _ = "STUB: not implemented"; return nil }

func (s *BusService) Start() { _ = "STUB: not implemented"; return }

// 错误subscriber清理

func (s *BusService) AddSubscriber(n Subscriber) error { _ = "STUB: not implemented"; return nil }

func (s *BusService) RemoveSubscriber(n Subscriber) { _ = "STUB: not implemented"; return }

func (s *BusService) closeBuses() { _ = "STUB: not implemented"; return }

// 通知内容塞到队列里
func (s *BusService) Fire(evt Event) error { _ = "STUB: not implemented"; return nil }

func (s *BusService) Closed() (b bool) { _ = "STUB: not implemented"; return false }

func (s *BusService) Stop() { _ = "STUB: not implemented"; return }

func NewBusService() *BusService { _ = "STUB: not implemented"; return nil }
