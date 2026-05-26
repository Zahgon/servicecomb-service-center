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

package sd

import (
	"github.com/apache/servicecomb-service-center/datasource/etcd/state"
)

func Service() state.State         { _ = "STUB: not implemented"; return *new(state.State) }
func Instance() state.State        { _ = "STUB: not implemented"; return *new(state.State) }
func Lease() state.State           { _ = "STUB: not implemented"; return *new(state.State) }
func ServiceIndex() state.State    { _ = "STUB: not implemented"; return *new(state.State) }
func ServiceAlias() state.State    { _ = "STUB: not implemented"; return *new(state.State) }
func ServiceTag() state.State      { _ = "STUB: not implemented"; return *new(state.State) }
func DependencyRule() state.State  { _ = "STUB: not implemented"; return *new(state.State) }
func DependencyQueue() state.State { _ = "STUB: not implemented"; return *new(state.State) }
func Domain() state.State          { _ = "STUB: not implemented"; return *new(state.State) }
func Project() state.State         { _ = "STUB: not implemented"; return *new(state.State) }
