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

package rbac

import (
	"net/http"

	"github.com/go-chassis/cari/rbac"

	"github.com/apache/servicecomb-service-center/pkg/rest"
)

var ErrConflictRole int32 = 409002

type RoleResource struct {
}

// URLPatterns define http pattern
func (rr *RoleResource) URLPatterns() []rest.Route { _ = "STUB: not implemented"; return nil }

// ListRoles list all roles and there's permissions
func (rr *RoleResource) ListRoles(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// roleParse parse the role info from the request body
func (rr *RoleResource) roleParse(body []byte) (*rbac.Role, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: validate role

// CreateRole create new role and assign permissions
func (rr *RoleResource) CreateRole(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// UpdateRole update role permissions
func (rr *RoleResource) UpdateRole(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// GetRole get the role info according to role name
func (rr *RoleResource) GetRole(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// DeleteRole delete the role info by role name
func (rr *RoleResource) DeleteRole(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}
