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

package resource

import (
	"context"

	kiemodel "github.com/apache/servicecomb-kie/pkg/model"
	v1sync "github.com/apache/servicecomb-service-center/syncer/api/v1"
)

const Config = "config"

func NewConfig(e *v1sync.Event) Resource { _ = "STUB: not implemented"; return *new(Resource) }

type kvConfig struct {
	defaultFailHandler

	event      *v1sync.Event
	input      *kiemodel.KVDoc
	cur        *kiemodel.KVDoc
	resource   docResource
	resourceID string
}

func (c *kvConfig) WithDomainProjectContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *kvConfig) loadInput() error { _ = "STUB: not implemented"; return nil }

func (c *kvConfig) LoadCurrentResource(ctx context.Context) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (c *kvConfig) NeedOperate(ctx context.Context) *Result { _ = "STUB: not implemented"; return nil }

func (c *kvConfig) Operate(ctx context.Context) *Result { _ = "STUB: not implemented"; return nil }

type docResource interface {
	Create(ctx context.Context, doc *kiemodel.KVDoc) error
	Get(ctx context.Context, ID string) (*kiemodel.KVDoc, error)
	Update(ctx context.Context, doc *kiemodel.KVDoc) error
	Delete(ctx context.Context, ID string) error
}

func (c *kvConfig) Create(ctx context.Context, doc *kiemodel.KVDoc) error {
	_ = "STUB: not implemented"
	return nil
}

func completeKV(kv *kiemodel.KVDoc, revision int64) { _ = "STUB: not implemented"; return }

func (c *kvConfig) Get(ctx context.Context, ID string) (*kiemodel.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *kvConfig) Update(ctx context.Context, doc *kiemodel.KVDoc) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *kvConfig) Delete(ctx context.Context, ID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *kvConfig) CreateHandle(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *kvConfig) UpdateHandle(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *kvConfig) DeleteHandle(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
