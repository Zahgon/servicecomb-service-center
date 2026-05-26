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

import (
	"github.com/apache/servicecomb-service-center/pkg/util"
)

type Group struct {
	name    string
	members *util.ConcurrentMap
}

func (g *Group) Name() string { _ = "STUB: not implemented"; return "" }

func (g *Group) Member(name string) Subscriber { _ = "STUB: not implemented"; return *new(Subscriber) }

func (g *Group) ForEach(iter func(m Subscriber)) { _ = "STUB: not implemented"; return }

func (g *Group) AddMember(subscriber Subscriber) Subscriber {
	_ = "STUB: not implemented"
	return *new(Subscriber)
}

func (g *Group) RemoveMember(name string) { _ = "STUB: not implemented"; return }

func (g *Group) Size() int { _ = "STUB: not implemented"; return 0 }

func NewGroup(name string) *Group { _ = "STUB: not implemented"; return nil }
