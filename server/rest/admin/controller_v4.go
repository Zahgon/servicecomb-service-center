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

package admin

import (
	"net/http"

	"github.com/apache/servicecomb-service-center/pkg/rest"
)

// ControllerV4 治理相关接口服务
type ControllerV4 struct {
}

// URLPatterns 路由.
func (ctrl *ControllerV4) URLPatterns() []rest.Route { _ = "STUB: not implemented"; return nil }

func (ctrl *ControllerV4) Dump(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (ctrl *ControllerV4) Clusters(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (ctrl *ControllerV4) AlarmList(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (ctrl *ControllerV4) ClearAlarm(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
