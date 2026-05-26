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

package mongo

var fastRegisterInstanceService *RegisterFastInstanceService

type RegisterFastInstanceService struct {
	InstEventCh  chan *InstanceRegisterEvent
	FailedInstCh chan *InstanceRegisterEvent
}

func NewFastRegisterInstanceService() *RegisterFastInstanceService {
	_ = "STUB: not implemented"
	return nil
}

func GetFastRegisterInstanceService() *RegisterFastInstanceService {
	_ = "STUB: not implemented"
	return nil
}

func SetFastRegisterInstanceService(service *RegisterFastInstanceService) {
	_ = "STUB: not implemented"
	return
}

func (s *RegisterFastInstanceService) AddEvent(event *InstanceRegisterEvent) {
	_ = "STUB: not implemented"
	return
}

func (s *RegisterFastInstanceService) AddFailedEvent(event *InstanceRegisterEvent) {
	_ = "STUB: not implemented"
	return
}

func (s *RegisterFastInstanceService) AddFailedEvents(events []*InstanceRegisterEvent) {
	_ = "STUB: not implemented"
	return
}
