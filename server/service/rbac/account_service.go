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

// Package rbac is dao layer API to help service center manage account, policy and role info
package rbac

import (
	"context"

	rbacmodel "github.com/go-chassis/cari/rbac"
)

// CreateAccount save account info
func CreateAccount(ctx context.Context, a *rbacmodel.Account) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateAccount updates an account's info, except the password
func UpdateAccount(ctx context.Context, name string, a *rbacmodel.Account) error {
	_ = "STUB: not implemented"
	// todo params validation
	return nil
}

func GetAccount(ctx context.Context, name string) (*rbacmodel.Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListAccount(ctx context.Context) ([]*rbacmodel.Account, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func AccountExist(ctx context.Context, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func DeleteAccount(ctx context.Context, name string) error { _ = "STUB: not implemented"; return nil }

// EditAccount save account info
func EditAccount(ctx context.Context, a *rbacmodel.Account) error {
	_ = "STUB: not implemented"
	return nil
}

func checkRoleNames(ctx context.Context, roles []string) error {
	_ = "STUB: not implemented"
	return nil
}

func illegalAccountCheck(ctx context.Context, target string) error {
	_ = "STUB: not implemented"
	return nil
}

func AccountUsage(ctx context.Context) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func BatchCreateAccounts(ctx context.Context, req *rbacmodel.BatchCreateAccountsRequest) (*rbacmodel.BatchCreateAccountsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func populateAccounts(accounts []*rbacmodel.Account) error { _ = "STUB: not implemented"; return nil }

func populateAccount(account *rbacmodel.Account) error { _ = "STUB: not implemented"; return nil }
