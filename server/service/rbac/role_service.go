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
	"context"

	rbacmodel "github.com/go-chassis/cari/rbac"
)

const isMigrated = "/cse-sr/role-migrated-lock"

func CreateRole(ctx context.Context, r *rbacmodel.Role) error {
	_ = "STUB: not implemented"
	return nil
}

func GetRole(ctx context.Context, name string) (*rbacmodel.Role, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListRole(ctx context.Context) ([]*rbacmodel.Role, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func RoleExist(ctx context.Context, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func DeleteRole(ctx context.Context, name string) error { _ = "STUB: not implemented"; return nil }

func EditRole(ctx context.Context, name string, a *rbacmodel.Role) error {
	_ = "STUB: not implemented"
	return nil
}

func illegalRoleCheck(role string) error { _ = "STUB: not implemented"; return nil }

func RoleUsage(ctx context.Context) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func migrateOldRoles() { _ = "STUB: not implemented"; return }
