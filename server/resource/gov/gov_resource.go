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

package gov

import (
	"net/http"

	"github.com/apache/servicecomb-service-center/pkg/rest"
)

type Governance struct {
}

const (
	AppKey         = "app"
	EnvironmentKey = "environment"
	KindKey        = ":kind"
	ProjectKey     = ":project"
	IDKey          = ":id"
	DisplayKey     = "display"
)

// Create gov config
func (t *Governance) Create(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Put gov config
func (t *Governance) Put(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// ListOrDisPlay return all gov config
func (t *Governance) ListOrDisPlay(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Get gov config
func (t *Governance) Get(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Delete delete gov config
func (t *Governance) Delete(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func processError(w http.ResponseWriter, err error, msg string) { _ = "STUB: not implemented"; return }

func (t *Governance) URLPatterns() []rest.Route { _ = "STUB: not implemented"; return nil }

//servicecomb.marker.{name}
//servicecomb.rateLimiter.{name}
//....
