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

package alarm

import (
	"github.com/apache/servicecomb-service-center/pkg/event"
	"github.com/apache/servicecomb-service-center/server/alarm/model"
)

const (
	Activated model.Status = "ACTIVATED"
	Cleared   model.Status = "CLEARED"
)

const (
	IDBackendConnectionRefuse model.ID = "BackendConnectionRefuse"
	IDInternalError           model.ID = "InternalError"
	IDIncrementPullError      model.ID = "IncrementPullError"
	IDWebsocketOfScSyncerLost model.ID = "WebsocketOfScSyncerLost"
)

const (
	FieldAdditionalContext = "detail"
)

const (
	Subject = "__ALARM_SUBJECT__"
	Group   = "__ALARM_GROUP__"
)

var ALARM = event.RegisterType("ALARM", 0)

func FieldBool(key string, v bool) model.Field { _ = "STUB: not implemented"; return *new(model.Field) }

func FieldString(key string, v string) model.Field {
	_ = "STUB: not implemented"
	return *new(model.Field)
}

func FieldInt64(key string, v int64) model.Field {
	_ = "STUB: not implemented"
	return *new(model.Field)
}

func FieldInt(key string, v int) model.Field { _ = "STUB: not implemented"; return *new(model.Field) }

func FieldFloat64(key string, v float64) model.Field {
	_ = "STUB: not implemented"
	return *new(model.Field)
}

func AdditionalContext(format string, args ...interface{}) model.Field {
	_ = "STUB: not implemented"
	return *new(model.Field)
}

func ListAll() []*model.AlarmEvent { _ = "STUB: not implemented"; return nil }

func Raise(id model.ID, fields ...model.Field) error { _ = "STUB: not implemented"; return nil }

func Clear(id model.ID) error { _ = "STUB: not implemented"; return nil }

func ClearAll() { _ = "STUB: not implemented"; return }
