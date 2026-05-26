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

package mongo

import (
	"context"

	"github.com/apache/servicecomb-service-center/datasource/rbac"
	rbacmodel "github.com/go-chassis/cari/rbac"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	rbac.Install("mongo", NewRbacDAO)
}

func NewRbacDAO(_ rbac.Options) (rbac.DAO, error) {
	_ = "STUB: not implemented"
	return *new(rbac.DAO), nil
}

type RbacDAO struct {
}

func (ds *RbacDAO) CreateAccount(ctx context.Context, a *rbacmodel.Account) error {
	_ = "STUB: not implemented"
	return nil
}

func createAccountTxn(ctx context.Context, a *rbacmodel.Account) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *RbacDAO) AccountExist(ctx context.Context, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ds *RbacDAO) GetAccount(ctx context.Context, name string) (*rbacmodel.Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *RbacDAO) ListAccount(ctx context.Context) ([]*rbacmodel.Account, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (ds *RbacDAO) DeleteAccount(ctx context.Context, names []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func deleteAccountTxn(ctx context.Context, names []string, ds *RbacDAO) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *RbacDAO) UpdateAccount(ctx context.Context, name string, account *rbacmodel.Account) error {
	_ = "STUB: not implemented"
	return nil
}

func updateAccountTxn(ctx context.Context, filter bson.M, updateFilter bson.M, account *rbacmodel.Account) error {
	_ = "STUB: not implemented"
	return nil
}
