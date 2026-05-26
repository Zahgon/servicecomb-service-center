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
	"sync"

	nf "github.com/apache/servicecomb-service-center/pkg/event"
	"github.com/apache/servicecomb-service-center/pkg/util"
	"github.com/apache/servicecomb-service-center/server/alarm/model"
)

var (
	service *Service
	once    sync.Once
)

type Service struct {
	nf.Subscriber
	alarms util.ConcurrentMap
}

func (ac *Service) Raise(id model.ID, fields ...model.Field) error {
	_ = "STUB: not implemented"
	return nil
}

func (ac *Service) Clear(id model.ID) error { _ = "STUB: not implemented"; return nil }

func (ac *Service) ListAll() (ls []*model.AlarmEvent) { _ = "STUB: not implemented"; return nil }

func (ac *Service) ClearAll() { _ = "STUB: not implemented"; return }

func (ac *Service) OnMessage(evt nf.Event) { _ = "STUB: not implemented"; return }

func NewAlarmService() *Service { _ = "STUB: not implemented"; return nil }
