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

package util

import (
	"github.com/cloudflare/gokey"
)

const TypePass = "pass"

var passwordSpec = &gokey.PasswordSpec{
	Length:         8,
	Upper:          1,
	Lower:          1,
	Digits:         1,
	Special:        1,
	AllowedSpecial: "-~!@#$%^&*()_=+|<>{}[]",
}

func SafeCloseChan(c chan struct{}) { _ = "STUB: not implemented"; return }

func BytesToStringWithNoCopy(bytes []byte) string { _ = "STUB: not implemented"; return "" }

func StringToBytesWithNoCopy(s string) []byte { _ = "STUB: not implemented"; return nil }

func ListToMap(list []string) map[string]struct{} { _ = "STUB: not implemented"; return nil }

func MapToList(dict map[string]struct{}) []string { _ = "STUB: not implemented"; return nil }

func StringJoin(args []string, sep string) string { _ = "STUB: not implemented"; return "" }

func GetCaller(skip int) (string, string, int, bool) {
	_ = "STUB: not implemented"
	return "", "", 0, false
}

func Int16ToInt64(bs []int16) (in int64) { _ = "STUB: not implemented"; return 0 }

func SliceHave(arr []string, str string) bool { _ = "STUB: not implemented"; return false }

func StringTRUE(s string) bool { _ = "STUB: not implemented"; return false }

func FromDomainProject(domainProject string) (domain, project string) {
	_ = "STUB: not implemented"
	return "", ""
}

func ToDomainProject(domain, project string) (domainProject string) {
	_ = "STUB: not implemented"
	return ""
}

func IsVersionOrHealthPattern(pattern string) bool { _ = "STUB: not implemented"; return false }

func ToSnake(name string) string { _ = "STUB: not implemented"; return "" }

//首字母大写

func GeneratePassword() (string, error) { _ = "STUB: not implemented"; return "", nil }
