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
	"errors"
	"sync"

	v1sync "github.com/apache/servicecomb-service-center/syncer/api/v1"
)

const (
	KV = "kv"

	ComparableKey = "comparable"
)

const (
	KVKey         = "key"
	KVKeyNonExist = "key not exist in opts"
)

var (
	manager KeyManager

	ErrRecordNonExist = errors.New("record non exist")
)

func NewKV(e *v1sync.Event) Resource { _ = "STUB: not implemented"; return *new(Resource) }

type kv struct {
	event *v1sync.Event
	key   string

	manager         KeyManager
	tombstoneLoader tombstoneLoader

	cur []byte

	defaultFailHandler
}

func (k *kv) LoadCurrentResource(ctx context.Context) *Result {
	_ = "STUB: not implemented"
	return nil
}

type Value struct {
	Timestamp int64 `json:"$timestamp"`
}

func (k *kv) getUpdateTime() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (k *kv) NeedOperate(ctx context.Context) *Result { _ = "STUB: not implemented"; return nil }

func (k *kv) CreateHandle(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (k *kv) UpdateHandle(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (k *kv) DeleteHandle(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

var once sync.Once

func keyManage() KeyManager { _ = "STUB: not implemented"; return *new(KeyManager) }

func (k *kv) Operate(ctx context.Context) *Result { _ = "STUB: not implemented"; return nil }

type KeyManager interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Put(ctx context.Context, key string, value []byte) error
	Post(ctx context.Context, key string, value []byte) error
	Delete(ctx context.Context, key string) error
}

type etcdManager struct {
}

func InitManager() { _ = "STUB: not implemented"; return }

func (e *etcdManager) Get(ctx context.Context, key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e etcdManager) Put(ctx context.Context, key string, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e etcdManager) Post(ctx context.Context, key string, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e etcdManager) Delete(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}
