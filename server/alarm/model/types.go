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

package model

import (
	nf "github.com/apache/servicecomb-service-center/pkg/event"
	"github.com/apache/servicecomb-service-center/pkg/util"
)

type ID string

type Status string

type Field struct {
	Key   string
	Value interface{}
}

type AlarmEvent struct {
	nf.Event `json:"-"`
	Status   Status          `json:"status"`
	ID       ID              `json:"id"`
	Fields   util.JSONObject `json:"fields,omitempty"`
}

func (ae *AlarmEvent) FieldBool(key string) bool { _ = "STUB: not implemented"; return false }

func (ae *AlarmEvent) FieldString(key string) string { _ = "STUB: not implemented"; return "" }

func (ae *AlarmEvent) FieldInt64(key string) int64 { _ = "STUB: not implemented"; return 0 }

func (ae *AlarmEvent) FieldInt(key string) int { _ = "STUB: not implemented"; return 0 }

func (ae *AlarmEvent) FieldFloat64(key string) float64 { _ = "STUB: not implemented"; return 0 }
