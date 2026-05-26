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

	rbacmodel "github.com/go-chassis/cari/rbac"
)

const (
	Account = "account"
)

func NewAccount(e *v1sync.Event) Resource { _ = "STUB: not implemented"; return *new(Resource) }

type accountManager interface {
	CreateAccount(ctx context.Context, a *rbacmodel.Account) error
	GetAccount(ctx context.Context, name string) (*rbacmodel.Account, error)
	UpdateAccount(ctx context.Context, account *rbacmodel.Account) error
	DeleteAccount(ctx context.Context, name string) error
}

type account struct {
	event *v1sync.Event

	input       *rbacmodel.Account
	accountName string

	cur *rbacmodel.Account

	defaultFailHandler

	manager accountManager
}

func (a *account) loadInput() error { _ = "STUB: not implemented"; return nil }

func (a *account) LoadCurrentResource(ctx context.Context) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (a *account) NeedOperate(ctx context.Context) *Result { _ = "STUB: not implemented"; return nil }

func (a *account) CreateHandle(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *account) UpdateHandle(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *account) DeleteHandle(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *account) Operate(ctx context.Context) *Result { _ = "STUB: not implemented"; return nil }

func (a *account) CreateAccount(ctx context.Context, at *rbacmodel.Account) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *account) GetAccount(ctx context.Context, name string) (*rbacmodel.Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *account) UpdateAccount(ctx context.Context, account *rbacmodel.Account) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *account) DeleteAccount(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}
