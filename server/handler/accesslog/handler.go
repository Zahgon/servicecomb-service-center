//go:build go1.9
// +build go1.9

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

package accesslog

import (
	"github.com/apache/servicecomb-service-center/pkg/chain"
	"github.com/go-chassis/openlog"
)

// Handler implements chain.Handler
// Handler records access log.
// Make sure to complete the initialization before handling the request.
type Handler struct {
	logger        openlog.Logger
	whiteListAPIs map[string]struct{} // not record access log
}

// AddWhiteListAPIs adds APIs to white list, where the APIs will be ignored
// in access log.
// Not safe for concurrent use.
func (h *Handler) AddWhiteListAPIs(apis ...string) { _ = "STUB: not implemented"; return }

// ShouldIgnoreAPI judges whether the API should be ignored in access log.
func (h *Handler) ShouldIgnoreAPI(api string) bool { _ = "STUB: not implemented"; return false }

// Handle handles the request
func (h *Handler) Handle(i *chain.Invocation) { _ = "STUB: not implemented"; return }

// format:  remoteIp requestReceiveTime "method requestUri proto" statusCode requestBodySize delay(ms)
// example: 127.0.0.1 2006-01-02T15:04:05.000Z07:00 "GET /v4/default/registry/microservices HTTP/1.1" 200 0 0

// NewAccessLogHandler creates a Handler
func NewAccessLogHandler(l openlog.Logger) *Handler { _ = "STUB: not implemented"; return nil }

// RegisterHandlers registers an access log handler to the handler chain
func RegisterHandlers() { _ = "STUB: not implemented"; return }

// no access log for heartbeat
