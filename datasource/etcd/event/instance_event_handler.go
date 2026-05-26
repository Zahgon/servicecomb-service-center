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
	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
	pb "github.com/go-chassis/cari/discovery"
)

// InstanceEventHandler is the handler to handle:
// 1. report instance metrics
// 2. recover the instance quota
// 3. publish the instance events to the subscribers
// 4. reset the find instance cache
type InstanceEventHandler struct {
}

func (h *InstanceEventHandler) Type() kvstore.Type {
	_ = "STUB: not implemented"
	return *new(kvstore.Type)
}

func (h *InstanceEventHandler) OnEvent(evt kvstore.Event) { _ = "STUB: not implemented"; return }

// 查询服务版本信息

// 查询所有consumer

func NewInstanceEventHandler() *InstanceEventHandler { _ = "STUB: not implemented"; return nil }

func PublishInstanceEvent(evt kvstore.Event, serviceKey *pb.MicroServiceKey, subscribers []string) {
	_ = "STUB: not implemented"
	return
}
