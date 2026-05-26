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
	"github.com/go-chassis/cari/rbac"
)

const ErrUserOrPwdWrongInHalfOpening int32 = 401302

var (
	ErrTokenExpired     = rbac.NewError(rbac.ErrTokenExpired, "")
	ErrAccountBlocked   = rbac.NewError(rbac.ErrAccountBlocked, "")
	ErrUserOrPwdWrong   = rbac.NewError(rbac.ErrUserOrPwdWrong, "")
	ErrUserOrPwdWrongEx = rbac.NewError(ErrUserOrPwdWrongInHalfOpening,
		"User name or password is wrong, RBAC system is half opening")
	ErrOldPwdWrong = rbac.NewError(rbac.ErrOldPwdWrong, "")
)

var roleMap = map[string]*rbac.Role{}

func init() {
	rbac.MustRegisterErr(ErrUserOrPwdWrongInHalfOpening, ErrUserOrPwdWrong.Error())

	// Assign resources to admin role, admin role own all permissions
	roleMap[rbac.RoleAdmin] = &rbac.Role{
		Name:  rbac.RoleAdmin,
		Perms: AdminPerms(),
	}
	roleMap[rbac.RoleDeveloper] = &rbac.Role{
		Name:  rbac.RoleDeveloper,
		Perms: DevPerms(),
	}
}

func initBuildInRole() { _ = "STUB: not implemented"; return }

func createBuildInRole(r *rbac.Role) { _ = "STUB: not implemented"; return }
