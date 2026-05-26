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
	"sync"

	"github.com/apache/servicecomb-service-center/pkg/rest"
	"github.com/apache/servicecomb-service-center/version"
)

var (
	versionJSONCache []byte
	parseVersionOnce sync.Once
)

const APIVersion = "4.0.0"

type Result struct {
	*version.Set
	APIVersion string `json:"apiVersion"`
}

type MainService struct {
	//
}

func (s *MainService) URLPatterns() []rest.Route { _ = "STUB: not implemented"; return nil }

func (s *MainService) ClusterHealth(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *MainService) GetVersion(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
