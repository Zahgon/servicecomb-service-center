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

package etcd

import (
	"context"

	crbac "github.com/go-chassis/cari/rbac"

	rbacsvc "github.com/apache/servicecomb-service-center/server/service/rbac"
)

const isMigrated = "/cse-sr/role-migrated"

var (
	resources   = crbac.BuildResourceList(rbacsvc.ResourceConfig)
	configPerms = &crbac.Permission{
		Resources: resources,
		Verbs:     []string{"*"},
	}
)

func (rm *RbacDAO) CreateRole(ctx context.Context, r *crbac.Role) error {
	_ = "STUB: not implemented"
	return nil
}

func (rm *RbacDAO) RoleExist(ctx context.Context, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rm *RbacDAO) GetRole(ctx context.Context, name string) (*crbac.Role, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rm *RbacDAO) ListRole(ctx context.Context) ([]*crbac.Role, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

//do not fail if some role is invalid

func (rm *RbacDAO) DeleteRole(ctx context.Context, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func RoleBindingExists(ctx context.Context, role string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rm *RbacDAO) UpdateRole(ctx context.Context, name string, role *crbac.Role) error {
	_ = "STUB: not implemented"
	return nil
}

func (rm *RbacDAO) MigrateOldRoles(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
