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
	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
)

type VersionRule func(sorted []string, kvs []*kvstore.KeyValue, start, end string) []string

func Sort(kvs []*kvstore.KeyValue, cmp func(start, end string) bool) {
	_ = "STUB: not implemented"
	return
}

func newSorter(kvs []*kvstore.KeyValue, cmp func(start string, end string) bool, ref bool) *serviceKeySorter {
	_ = "STUB: not implemented"
	return nil
}

func (vr VersionRule) Match(kvs []*kvstore.KeyValue, ops ...string) []string {
	_ = "STUB: not implemented"
	return nil
}

type serviceKeySorter struct {
	sortArr []string
	kvs     []*kvstore.KeyValue
	cmp     func(i, j string) bool
}

func (sks *serviceKeySorter) Len() int { _ = "STUB: not implemented"; return 0 }

func (sks *serviceKeySorter) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (sks *serviceKeySorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func Larger(start, end string) bool { _ = "STUB: not implemented"; return false }

func LessEqual(start, end string) bool { _ = "STUB: not implemented"; return false }

// Latest return latest version kv
func Latest(sorted []string, kvs []*kvstore.KeyValue, _, _ string) []string {
	_ = "STUB: not implemented"
	return nil
}

// Range return start <= version < end
func Range(sorted []string, kvs []*kvstore.KeyValue, start, end string) []string {
	_ = "STUB: not implemented"
	return nil
}

// end >= k >= start

// AtLess return version >= start
func AtLess(sorted []string, kvs []*kvstore.KeyValue, start, _ string) []string {
	_ = "STUB: not implemented"
	return nil
}

func ParseVersionRule(versionRule string) func(kvs []*kvstore.KeyValue) []string {
	_ = "STUB: not implemented"
	return nil
}

// 取最低版本及高版本集合

// 取版本范围集合

// 精确匹配

func VersionMatchRule(version string, versionRule string) bool {
	_ = "STUB: not implemented"
	return false
}
