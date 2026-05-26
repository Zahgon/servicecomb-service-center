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

package exception

import (
	"net/http"

	"github.com/apache/servicecomb-service-center/pkg/chain"
)

var whitelists = make(map[string]struct{})

// Handler provide a common response writer to handle exceptions
type Handler struct {
}

func (h *Handler) Handle(i *chain.Invocation) { _ = "STUB: not implemented"; return }

func (h *Handler) responseError(w http.ResponseWriter, e error) (statusCode int) {
	_ = "STUB: not implemented"
	return 0
}

func (h *Handler) alarmIfInternalError(statusCode int, errMsg string) {
	_ = "STUB: not implemented"
	return
}

func RegisterHandlers() { _ = "STUB: not implemented"; return }

func RegisterWhitelist(method, apiPath string) { _ = "STUB: not implemented"; return }

func InWhitelist(method, apiPath string) bool { _ = "STUB: not implemented"; return false }
