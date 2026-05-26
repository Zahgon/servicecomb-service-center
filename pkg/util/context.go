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

package util

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	CtxDomain        CtxKey = "domain"
	CtxProject       CtxKey = "project"
	CtxTargetDomain  CtxKey = "target-domain"
	CtxTargetProject CtxKey = "target-project"
	SPLIT                   = "/"
)

type StringContext struct {
	parentCtx context.Context
	kv        *ConcurrentMap
}

func (c *StringContext) Deadline() (deadline time.Time, ok bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func (c *StringContext) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (c *StringContext) Err() error { _ = "STUB: not implemented"; return nil }

func (c *StringContext) Value(key interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func (c *StringContext) SetKV(key CtxKey, val interface{}) { _ = "STUB: not implemented"; return }

func NewStringContext(ctx context.Context) *StringContext { _ = "STUB: not implemented"; return nil }

func SetContext(ctx context.Context, key CtxKey, val interface{}) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func CloneContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// FromContext return the value from ctx, return empty STRING if not found
func FromContext(ctx context.Context, key CtxKey) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func SetRequestContext(r *http.Request, key CtxKey, val interface{}) *http.Request {
	_ = "STUB: not implemented"
	return nil
}

func SetFiberContext(c *fiber.Ctx, key CtxKey, val interface{}) { _ = "STUB: not implemented"; return }

func ParseDomainProject(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func ParseTargetDomainProject(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func ParseDomain(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func ParseTargetDomain(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func ParseProject(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func ParseTargetProject(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func SetDomain(ctx context.Context, domain string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func SetProject(ctx context.Context, project string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func SetTargetDomain(ctx context.Context, domain string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func SetTargetProject(ctx context.Context, project string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func SetDomainProject(ctx context.Context, domain string, project string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func SetDomainProjectString(ctx context.Context, domainProject string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func SetTargetDomainProject(ctx context.Context, domain string, project string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func WithNoCache(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func NoCache(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func WithCacheOnly(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func CacheOnly(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func WithGlobal(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func Global(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func EnableSync(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func WithRequestRev(ctx context.Context, rev string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func WithResponseRev(ctx context.Context, rev string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
