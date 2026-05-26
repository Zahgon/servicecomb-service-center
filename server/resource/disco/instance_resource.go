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

package disco

import (
	"net/http"

	"github.com/apache/servicecomb-service-center/datasource"
	"github.com/apache/servicecomb-service-center/pkg/rest"
)

type InstanceResource struct {
	//
}

func (s *InstanceResource) URLPatterns() []rest.Route { _ = "STUB: not implemented"; return nil }

func (s *InstanceResource) LegacyRegisterInstance(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *InstanceResource) SendHeartbeat(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *InstanceResource) SendManyHeartbeat(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *InstanceResource) UnregisterInstance(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *InstanceResource) FindInstances(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *InstanceResource) InstancesAction(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func findManyInstances(w http.ResponseWriter, r *http.Request, body []byte) {
	_ = "STUB: not implemented"
	return
}

func (s *InstanceResource) GetInstance(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *InstanceResource) ListInstance(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *InstanceResource) PutInstanceStatus(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *InstanceResource) PutInstanceProperties(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *InstanceResource) UpdateManyInstanceStatus(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

type UpdateManyInstanceStatusRequest struct {
	Matches datasource.MatchPolicy `json:"matches,omitempty"`
	Status  string                 `json:"status,omitempty"`
}
