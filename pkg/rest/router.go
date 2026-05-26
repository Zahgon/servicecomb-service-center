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

package rest

import (
	"net/http"

	"github.com/apache/servicecomb-service-center/pkg/chain"
)

var doNothingFunc = func(_ chain.Result) {}

// Router is a HTTP request multiplexer
// Attention:
//  1. not thread-safe, must be initialized completely before serve http request
//  2. redirect not supported
type Router struct {
	handlers  map[string][]*urlPatternHandler
	chainName string
}

// RegisterServant registers a RouteGroup
// servant must be an pointer to service object
func (router *Router) RegisterServant(servant RouteGroup) { _ = "STUB: not implemented"; return }

func (router *Router) setChainName(name string) { _ = "STUB: not implemented"; return }

func (router *Router) addRoute(route *Route) (err error) { _ = "STUB: not implemented"; return nil }

// ServeHTTP implements http.Handler
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (router *Router) serve(ph *urlPatternHandler, w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// NewRouter news a Router
func NewRouter() *Router { _ = "STUB: not implemented"; return nil }
