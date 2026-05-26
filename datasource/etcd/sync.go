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
	"errors"
)

const (
	SyncAllKey     = "/cse-sr/sync-all"
	SyncAllLockKey = "/cse-sr/sync-all-lock"

	// one minutes
	defaultLockTime = 60
)

var (
	ErrWithoutDomainProject = errors.New("key without domain and project")
)

type SyncManager struct {
}

// SyncAll will list all services,accounts,roles,schemas,tags,deps and use tasks to store
func (s *SyncManager) SyncAll(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func syncAllAccounts(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func syncAllRoles(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// syncAllTags func use kv resource task to store tags
func syncAllTags(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func syncAllServices(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// syncAllSchemas func use kv resource task to store schemas
func syncAllSchemas(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func syncAllServiceSchemas(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func syncAllServiceSchemaRefs(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func syncAllServiceSchemaContents(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func syncAllServiceSchemaSummaries(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func syncAllDependencies(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func getDomainProject(key string, prefixKey string) (domain string, project string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
