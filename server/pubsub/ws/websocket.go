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
	"errors"
	"time"

	"github.com/gorilla/websocket"
)

const Websocket = "Websocket"

var errServiceNotExist = errors.New("service does not exist")

type WebSocket struct {
	Options
	Conn          *websocket.Conn
	RemoteAddr    string
	DomainProject string
	ConsumerID    string

	ticker   *time.Ticker
	needPing bool
	idleCh   chan struct{}
}

func (wh *WebSocket) Init() { _ = "STUB: not implemented"; return }

func (wh *WebSocket) registerMessageHandler() { _ = "STUB: not implemented"; return }

// PING

// PONG

// CLOSE

func (wh *WebSocket) ReadMessage() error { _ = "STUB: not implemented"; return nil }

func (wh *WebSocket) sendClose(code int, text string) error { _ = "STUB: not implemented"; return nil }

// NeedCheck will be called by checker
func (wh *WebSocket) NeedCheck() interface{} { _ = "STUB: not implemented"; return nil }

// reset if idleCh

// CheckHealth will be called if NeedCheck() returns not nil
func (wh *WebSocket) CheckHealth(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (wh *WebSocket) WritePingPong(messageType int) error { _ = "STUB: not implemented"; return nil }

func (wh *WebSocket) WriteTextMessage(message []byte) error { _ = "STUB: not implemented"; return nil }

func (wh *WebSocket) Idle() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (wh *WebSocket) SetIdle() { _ = "STUB: not implemented"; return }

func NewWebSocket(domainProject, serviceID string, conn *websocket.Conn) *WebSocket {
	_ = "STUB: not implemented"
	return nil
}
