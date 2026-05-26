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
	"net/http"

	"github.com/apache/servicecomb-service-center/pkg/rest"
)

const DefaultTokenExpirationDuration = "12h"

type AuthResource struct {
}

// URLPatterns define htp pattern
func (ar *AuthResource) URLPatterns() []rest.Route { _ = "STUB: not implemented"; return nil }

func (ar *AuthResource) CreateAccount(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (ar *AuthResource) DeleteAccount(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (ar *AuthResource) UpdateAccount(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (ar *AuthResource) ListAccount(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (ar *AuthResource) GetAccount(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (ar *AuthResource) ChangePassword(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (ar *AuthResource) Login(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (ar *AuthResource) ListSelfPerms(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (ar *AuthResource) ListLock(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (ar *AuthResource) BatchCreateAccount(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func MakeBanKey(name, ip string) string { _ = "STUB: not implemented"; return "" }
