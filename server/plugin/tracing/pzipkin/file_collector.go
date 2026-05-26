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

package pzipkin

import (
	"time"

	"github.com/openzipkin/zipkin-go-opentracing/thrift/gen-go/zipkincore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type FileCollector struct {
	// Timeout is the timeout of sending span to chan
	Timeout time.Duration
	// Interval is interval to log
	Interval time.Duration
	// BatchSize is the log batch size
	BatchSize int
	logger    *lumberjack.Logger
	c         chan *zipkincore.Span
}

func (f *FileCollector) Collect(span *zipkincore.Span) error { _ = "STUB: not implemented"; return nil }

func (f *FileCollector) Close() error { _ = "STUB: not implemented"; return nil }

func (f *FileCollector) write(batch []*zipkincore.Span) (c int) {
	_ = "STUB: not implemented"
	return 0
}

func (f *FileCollector) Run() { _ = "STUB: not implemented"; return }

// allocate more

// new one

func NewFileCollector(path string) (*FileCollector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// megabytes
