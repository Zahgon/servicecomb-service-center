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
)

type DependencyService struct {
}

func (s *DependencyService) URLPatterns() []rest.Route { _ = "STUB: not implemented"; return nil }

// Deprecated
func (s *DependencyService) AddDependencies(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Deprecated
func (s *DependencyService) PutDependencies(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *DependencyService) ListProviders(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *DependencyService) ListConsumers(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
