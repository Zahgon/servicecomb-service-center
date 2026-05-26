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

package v4

import (
	"net/http"

	"github.com/apache/servicecomb-service-center/pkg/rest"
	"github.com/apache/servicecomb-service-center/server/handler/exception"
	"github.com/gorilla/websocket"
)

const (
	APIWatch     = "/v4/:project/registry/microservices/:serviceId/watcher"
	APIHeartbeat = "/v4/:project/registry/microservices/:serviceId/instances/:instanceId/heartbeat"
)

func init() {
	exception.RegisterWhitelist(http.MethodGet, APIWatch)
	exception.RegisterWhitelist(http.MethodGet, APIHeartbeat)
}

type WatchService struct {
	//
}

func (s *WatchService) URLPatterns() []rest.Route { _ = "STUB: not implemented"; return nil }

func upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *WatchService) Watch(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *WatchService) Heartbeat(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
