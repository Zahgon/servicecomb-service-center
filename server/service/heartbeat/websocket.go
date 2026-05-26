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

package heartbeat

import (
	"context"
	"sync"
	"time"

	pb "github.com/go-chassis/cari/discovery"
	"github.com/gorilla/websocket"
)

const (
	Websocket         = "Websocket"
	defaultPingPeriod = 30 * time.Second
	minPeriod         = 1 * time.Second
	maxPeriod         = 1 * time.Hour
)

var (
	once       sync.Once
	pingPeriod time.Duration
)

type client struct {
	cxt        context.Context
	conn       *websocket.Conn
	serviceID  string
	instanceID string
}

func configuration() { _ = "STUB: not implemented"; return }

func newClient(ctx context.Context, conn *websocket.Conn, serviceID string, instanceID string) *client {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) sendClose(code int, text string) error { _ = "STUB: not implemented"; return nil }

func (c *client) heartbeat() { _ = "STUB: not implemented"; return }

func (c *client) handleMessage() { _ = "STUB: not implemented"; return }

func SendEstablishError(conn *websocket.Conn, err error) { _ = "STUB: not implemented"; return }

func Heartbeat(ctx context.Context, conn *websocket.Conn, serviceID string, instanceID string) {
	_ = "STUB: not implemented"
	return
}

func process(client *client) { _ = "STUB: not implemented"; return }

func WatchHeartbeat(ctx context.Context, in *pb.HeartbeatRequest, conn *websocket.Conn) {
	_ = "STUB: not implemented"
	return
}

func preOp(ctx context.Context, in *pb.HeartbeatRequest) error {
	_ = "STUB: not implemented"
	return nil
}
