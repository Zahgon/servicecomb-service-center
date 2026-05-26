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

package sync

import (
	"context"
)

const (
	CollectionTask      = "task"
	CollectionTombstone = "tombstone"
)

type Options struct {
	ResourceID string
	Opts       map[string]string
}

type Option func(options *Options)

func NewSyncOption() Options { _ = "STUB: not implemented"; return *new(Options) }

func WithResourceID(resourceID string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithOpts(opts map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

func DoCreateOpts(ctx context.Context, resourceType string, resource interface{}, options ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

func DoUpdateOpts(ctx context.Context, resourceType string, resource interface{}, options ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

func DoDeleteOpts(ctx context.Context, resourceType, resourceID string, resource interface{}, options ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

func doOpts(ctx context.Context, action string, resourceType string, resource interface{}, options ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

func doTaskOpt(ctx context.Context, action string, resourceType string, resource interface{}, syncOpts *Options) error {
	_ = "STUB: not implemented"
	return nil
}

func doTombstoneOpt(ctx context.Context, resourceType, resourceID string) error {
	_ = "STUB: not implemented"
	return nil
}
