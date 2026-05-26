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

package rest

import (
	"errors"
	"net/http"

	"github.com/go-chassis/cari/discovery"
)

var errNilRequestBody = errors.New("request body is nil")

func WriteError(w http.ResponseWriter, code int32, detail string) {
	_ = "STUB: not implemented"
	return
}

func WriteServiceError(w http.ResponseWriter, err error) { _ = "STUB: not implemented"; return }

// WriteResponse writes http response
// If the resp is nil or represents success, response status is http.StatusOK,
// response content is obj.
// If the resp represents fail, response status is from the code in the
// resp, response content is from the message in the resp.
func WriteResponse(w http.ResponseWriter, r *http.Request, resp *discovery.Response, obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func WriteSuccess(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// ReadBody can re-read the request body
func ReadBody(r *http.Request) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
