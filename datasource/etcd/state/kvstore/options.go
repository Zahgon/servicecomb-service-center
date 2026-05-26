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

package kvstore

import (
	"time"

	"github.com/apache/servicecomb-service-center/datasource/etcd/state/parser"
)

type Options struct {
	// Key is the prefix to unique specify resource type
	Key          string
	InitSize     int
	Timeout      time.Duration
	Period       time.Duration
	DeferHandler DeferHandler
	OnEvent      EventFunc
	Parser       parser.Parser
}

func (opts *Options) String() string { _ = "STUB: not implemented"; return "" }

func (opts *Options) WithPrefix(key string) *Options { _ = "STUB: not implemented"; return nil }

func (opts *Options) WithInitSize(size int) *Options { _ = "STUB: not implemented"; return nil }

func (opts *Options) WithTimeout(ot time.Duration) *Options { _ = "STUB: not implemented"; return nil }

func (opts *Options) WithPeriod(ot time.Duration) *Options { _ = "STUB: not implemented"; return nil }

func (opts *Options) WithDeferHandler(h DeferHandler) *Options {
	_ = "STUB: not implemented"
	return nil
}

func (opts *Options) WithEventFunc(f EventFunc) *Options { _ = "STUB: not implemented"; return nil }

func (opts *Options) AppendEventFunc(f EventFunc) *Options { _ = "STUB: not implemented"; return nil }

func (opts *Options) WithParser(parser parser.Parser) *Options {
	_ = "STUB: not implemented"
	return nil
}

func NewOptions() *Options { _ = "STUB: not implemented"; return nil }
