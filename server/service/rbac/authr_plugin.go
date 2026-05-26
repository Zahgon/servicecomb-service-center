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
	"crypto/rsa"

	"github.com/go-chassis/go-chassis/v2/security/authr"
)

// EmbeddedAuthenticator is sc default auth plugin, RBAC data is persisted in etcd
type EmbeddedAuthenticator struct {
}

func newEmbeddedAuthenticator(_ *authr.Options) (authr.Authenticator, error) {
	_ = "STUB: not implemented"
	return *new(authr.Authenticator), nil
}

// Login check db user and password,will verify and return token for valid account
func (a *EmbeddedAuthenticator) Login(ctx context.Context, user string, password string, opts ...authr.LoginOption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

//TODO config for each user

// Authenticate parse a token to claims
func (a *EmbeddedAuthenticator) Authenticate(_ context.Context, tokenStr string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *EmbeddedAuthenticator) isTokenExpiredError(err error) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *EmbeddedAuthenticator) authToken(tokenStr string, pub *rsa.PublicKey) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UserOrPwdWrongError() error { _ = "STUB: not implemented"; return nil }

func init() {
	authr.Install("default", newEmbeddedAuthenticator)
}
