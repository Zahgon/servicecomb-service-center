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
	"crypto/rsa"
	"errors"
)

const (
	RootName     = "root"
	InitPassword = "SC_INIT_ROOT_PASSWORD"
)

var (
	ErrEmptyCurrentPassword = errors.New("current password should not be empty")
	ErrNoPermChangeAccount  = errors.New("can not change other account password")
	ErrWrongPassword        = errors.New("current pwd is wrong")
	ErrSamePassword         = errors.New("the password can not be same as old one")
	ErrNoPrivateKey         = errors.New("read private key failed")

	privateKey *rsa.PrivateKey
)

// Init decide whether enable rbac function and save the build-in roles to db
func Init() { _ = "STUB: not implemented"; return }

// build-in role init

func add2WhiteAPIList() { _ = "STUB: not implemented"; return }

// user can list self permission without account get permission

// user can change self password without account modify permission

func initBuildInAccount() { _ = "STUB: not implemented"; return }

// read key to memory
func readPrivateKey() { _ = "STUB: not implemented"; return }

// 打开文件

// read key to memory
func readPublicKey() { _ = "STUB: not implemented"; return }

// 打开文件

func initFirstTime() {
	_ = "STUB: not implemented"
	// handle root account
	return
}

func getPassword() string { _ = "STUB: not implemented"; return "" }

func Enabled() bool { _ = "STUB: not implemented"; return false }

// PublicKey get public key to verify a token
func PublicKey() string { _ = "STUB: not implemented"; return "" }

// GetPrivateKey return rsa key instance
func GetPrivateKey() (*rsa.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

// MakeBanKey return ban key
func MakeBanKey(name, ip string) string { _ = "STUB: not implemented"; return "" }
