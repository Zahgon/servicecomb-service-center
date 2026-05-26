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

package buildin

import (
	"errors"
	"net/http"

	"github.com/apache/servicecomb-service-center/server/plugin/auth"
)

var ErrCtxMatchPatternNotFound = errors.New("CtxMatchPattern not found")

var APIMapping = map[string]ParseFunc{}

type ParseFunc func(r *http.Request) (*auth.ResourceScope, error)

// ApplyAll work when no api registered by RegisterParseFunc matched
func ApplyAll(r *http.Request) (*auth.ResourceScope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FromRequest return the scope parsed from request
// return nil mean apply all resources
func FromRequest(r *http.Request) *auth.ResourceScope { _ = "STUB: not implemented"; return nil }

func GetAPIParseFunc(apiPattern string) ParseFunc {
	_ = "STUB: not implemented"
	return *new(ParseFunc)
}

func RegisterParseFunc(apiPathPrefix string, f ParseFunc) { _ = "STUB: not implemented"; return }
