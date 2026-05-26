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

package resource

import (
	"context"

	v1sync "github.com/apache/servicecomb-service-center/syncer/api/v1"

	crbac "github.com/go-chassis/cari/rbac"
)

const (
	Role = "role"
)

func NewRole(e *v1sync.Event) Resource { _ = "STUB: not implemented"; return *new(Resource) }

type role struct {
	event *v1sync.Event

	input       *crbac.Role
	deleteInput *string

	roleName string

	cur *crbac.Role

	manager roleManager

	defaultFailHandler
}

type roleManager interface {
	GetRole(ctx context.Context, name string) (*crbac.Role, error)
	EditRole(ctx context.Context, name string, r *crbac.Role) error
	CreateRole(ctx context.Context, r *crbac.Role) error
	DeleteRole(ctx context.Context, name string) error
}

func (r *role) loadInput() error { _ = "STUB: not implemented"; return nil }

func (r *role) LoadCurrentResource(ctx context.Context) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (r *role) NeedOperate(ctx context.Context) *Result { _ = "STUB: not implemented"; return nil }

func (r *role) CreateHandle(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *role) UpdateHandle(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *role) DeleteHandle(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *role) Operate(ctx context.Context) *Result { _ = "STUB: not implemented"; return nil }

func (r *role) GetRole(ctx context.Context, name string) (*crbac.Role, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *role) EditRole(ctx context.Context, name string, role *crbac.Role) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *role) CreateRole(ctx context.Context, role *crbac.Role) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *role) DeleteRole(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}
