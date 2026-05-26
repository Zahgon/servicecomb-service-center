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

package state

import (
	"time"

	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
	"github.com/apache/servicecomb-service-center/datasource/etcd/state/parser"
)

type Option func(options *kvstore.Options)

func WithPrefix(key string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithInitSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTimeout(ot time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPeriod(ot time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDeferHandler(h kvstore.DeferHandler) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithParser(parser parser.Parser) Option { _ = "STUB: not implemented"; return *new(Option) }

func ToOptions(opts ...Option) kvstore.Options {
	_ = "STUB: not implemented"
	return *new(kvstore.Options)
}
