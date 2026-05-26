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
	"sync"

	"github.com/apache/servicecomb-service-center/pkg/plugin"
	"github.com/apache/servicecomb-service-center/server/plugin/tracing"
	"github.com/opentracing/opentracing-go"
)

var once sync.Once

func init() {
	plugin.RegisterPlugin(plugin.Plugin{Kind: tracing.TRACING, Name: "buildin", New: New})
}

func New() plugin.Instance { _ = "STUB: not implemented"; return *new(plugin.Instance) }

type Zipkin struct {
}

func (zp *Zipkin) ServerBegin(operationName string, itf interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// grpc?

func (zp *Zipkin) ServerEnd(itf interface{}, code int, message string) {
	_ = "STUB: not implemented"
	return
}

func (zp *Zipkin) ClientBegin(operationName string, itf interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// inject context

func (zp *Zipkin) ClientEnd(itf interface{}, code int, message string) {
	_ = "STUB: not implemented"
	return
}

func setResultTags(span opentracing.Span, code int, message string) {
	_ = "STUB: not implemented"
	return
}
