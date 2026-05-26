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
	"github.com/little-cui/etcdadpt"

	"github.com/apache/servicecomb-service-center/datasource/rbac"
)

func init() {
	rbac.Install("etcd", NewRbacDAO)
	rbac.Install("embeded_etcd", NewRbacDAO)
	rbac.Install("embedded_etcd", NewRbacDAO)
}

func NewRbacDAO(_ rbac.Options) (rbac.DAO, error) {
	_ = "STUB: not implemented"
	return *new(rbac.DAO), nil
}

type RbacDAO struct {
}

func (ds *RbacDAO) CreateAccount(ctx context.Context, a *crbac.Account) error {
	_ = "STUB: not implemented"
	return nil
}

func GenAccountOpts(a *crbac.Account, action etcdadpt.Action) ([]etcdadpt.OpOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *RbacDAO) AccountExist(ctx context.Context, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ds *RbacDAO) GetAccount(ctx context.Context, name string) (*crbac.Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds *RbacDAO) compatibleOldVersionAccount(a *crbac.Account) {
	_ = "STUB: not implemented"
	// old version use Role, now use Roles
	// Role/Roles will not exist at the same time
	return
}

func (ds *RbacDAO) ListAccount(ctx context.Context) ([]*crbac.Account, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

//do not fail if some account is invalid

func (ds *RbacDAO) DeleteAccount(ctx context.Context, names []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

//do not fail if some account is invalid

//do not fail if some account is invalid

func (ds *RbacDAO) UpdateAccount(ctx context.Context, _ string, account *crbac.Account) error {
	_ = "STUB: not implemented"
	return nil
}

func hasRole(account *crbac.Account, r string) bool { _ = "STUB: not implemented"; return false }
