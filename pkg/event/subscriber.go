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

package event

type Subscriber interface {
	ID() string
	Subject() string
	Group() string
	Type() Type
	Bus() *BusService
	SetBus(*BusService)

	// Err event bus remove subscriber automatically, if return not nil.
	// Implement of OnMessage should call SetError when run exception
	Err() error
	SetError(err error)

	Close()
	// OnAccept call when subscriber appended in event bus successfully
	OnAccept()
	// OnMessage call when event bus fire a msg, it must be non-blocked
	OnMessage(Event)
}

type baseSubscriber struct {
	nType   Type
	id      string
	subject string
	group   string
	service *BusService
	err     error
}

func (s *baseSubscriber) ID() string             { _ = "STUB: not implemented"; return "" }
func (s *baseSubscriber) Subject() string        { _ = "STUB: not implemented"; return "" }
func (s *baseSubscriber) Group() string          { _ = "STUB: not implemented"; return "" }
func (s *baseSubscriber) Type() Type             { _ = "STUB: not implemented"; return *new(Type) }
func (s *baseSubscriber) Bus() *BusService       { _ = "STUB: not implemented"; return nil }
func (s *baseSubscriber) SetBus(svc *BusService) { _ = "STUB: not implemented"; return }
func (s *baseSubscriber) Err() error             { _ = "STUB: not implemented"; return nil }
func (s *baseSubscriber) SetError(err error)     { _ = "STUB: not implemented"; return }
func (s *baseSubscriber) Close()                 { _ = "STUB: not implemented"; return }
func (s *baseSubscriber) OnAccept()              { _ = "STUB: not implemented"; return }
func (s *baseSubscriber) OnMessage(_ Event)      { _ = "STUB: not implemented"; return }

func NewSubscriber(nType Type, subject, group string) Subscriber {
	_ = "STUB: not implemented"
	return *new(Subscriber)
}
