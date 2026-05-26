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
	"context"

	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
)

type TagsChangedTask struct {
	kvstore.Event

	key string
	err error

	DomainProject string
	ConsumerID    string
}

func (apt *TagsChangedTask) Key() string { _ = "STUB: not implemented"; return "" }

func (apt *TagsChangedTask) Do(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (apt *TagsChangedTask) Err() error { _ = "STUB: not implemented"; return nil }

func (apt *TagsChangedTask) publish(ctx context.Context, domainProject, consumerID string) error {
	_ = "STUB: not implemented"
	return nil
}

// TagEventHandler is the handler to handle:
// 1. publish the EVT_EXPIRE event to subscribers when tag is changed
// 2. reset the find instance cache
type TagEventHandler struct {
}

func (h *TagEventHandler) Type() kvstore.Type { _ = "STUB: not implemented"; return *new(kvstore.Type) }

func (h *TagEventHandler) OnEvent(evt kvstore.Event) { _ = "STUB: not implemented"; return }

func NewTagEventHandler() *TagEventHandler { _ = "STUB: not implemented"; return nil }

func NewTagsChangedAsyncTask(domainProject, consumerID string, evt kvstore.Event) *TagsChangedTask {
	_ = "STUB: not implemented"
	return nil
}
