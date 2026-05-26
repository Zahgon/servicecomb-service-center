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

package buildin

import (
	"context"
	"errors"
	"net/http"

	"github.com/apache/servicecomb-service-center/pkg/plugin"
	"github.com/apache/servicecomb-service-center/server/plugin/auth"
	rbacmodel "github.com/go-chassis/cari/rbac"
)

var ErrNoRoles = errors.New("no role found in token")

func init() {
	plugin.RegisterPlugin(plugin.Plugin{Kind: auth.AUTH, Name: "buildin", New: New})
}

func New() plugin.Instance { _ = "STUB: not implemented"; return *new(plugin.Instance) }

type TokenAuthenticator struct {
}

func (ba *TokenAuthenticator) Identify(req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func getRequestPattern(req *http.Request) string { _ = "STUB: not implemented"; return "" }

func (ba *TokenAuthenticator) mustAuth(req *http.Request, pattern string) (*rbacmodel.Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ba *TokenAuthenticator) VerifyRequest(req *http.Request) (*rbacmodel.Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accountExist(ctx context.Context, user string) error {
	_ = "STUB: not implemented"
	// if root should pass, cause of root initialization
	return nil
}

func filterRoles(roleList []string) (hasAdmin bool, normalRoles []string) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ba *TokenAuthenticator) VerifyToken(req *http.Request) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this method decouple business code and perm checks
func checkPerm(roleList []string, req *http.Request) ([]map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//todo fast check for dev role

//TODO add project
