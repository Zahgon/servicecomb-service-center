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
	"time"

	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
)

const (
	eventBlockSize             = 1000
	deferCheckWindow           = 2 * time.Second // instance DELETE event will be delay.
	selfPreservationPercentage = 0.8
	selfPreservationMaxTTL     = 10 * 60 // 10min
	selfPreservationInitCount  = 5
)

var (
	TypeDomain          kvstore.Type
	TypeProject         kvstore.Type
	TypeService         kvstore.Type
	TypeServiceIndex    kvstore.Type
	TypeServiceAlias    kvstore.Type
	TypeServiceTag      kvstore.Type
	TypeDependencyRule  kvstore.Type
	TypeDependencyQueue kvstore.Type
	TypeInstance        kvstore.Type
	TypeLease           kvstore.Type
)

func RegisterInnerTypes() { _ = "STUB: not implemented"; return }
