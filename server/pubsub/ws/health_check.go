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

package ws

import (
	"context"
	"sync"
)

var checker *HealthCheck

func init() {
	checker = NewHealthCheck()
	checker.Run()
}

type HealthCheck struct {
	wss  []*WebSocket
	lock sync.Mutex
}

func (wh *HealthCheck) Run() { _ = "STUB: not implemented"; return }

func (wh *HealthCheck) loop(ctx context.Context) { _ = "STUB: not implemented"; return }

// server shutdown

func (wh *HealthCheck) check(ws *WebSocket) { _ = "STUB: not implemented"; return }

func (wh *HealthCheck) Accept(ws *WebSocket) int { _ = "STUB: not implemented"; return 0 }

func (wh *HealthCheck) Remove(ws *WebSocket) int { _ = "STUB: not implemented"; return 0 }

func NewHealthCheck() *HealthCheck { _ = "STUB: not implemented"; return nil }

func HealthChecker() *HealthCheck { _ = "STUB: not implemented"; return nil }
