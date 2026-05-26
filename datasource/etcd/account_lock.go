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

package etcd

import (
	"context"

	"github.com/apache/servicecomb-service-center/datasource/rbac"
)

func (al *RbacDAO) UpsertLock(ctx context.Context, lock *rbac.Lock) error {
	_ = "STUB: not implemented"
	return nil
}

func (al *RbacDAO) GetLock(ctx context.Context, key string) (*rbac.Lock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (al *RbacDAO) ListLock(ctx context.Context) ([]*rbac.Lock, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

//do not fail if some account is invalid

func (al *RbacDAO) DeleteLock(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (al *RbacDAO) DeleteLockList(ctx context.Context, keys []string) error {
	_ = "STUB: not implemented"
	return nil
}
