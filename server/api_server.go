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

package server

import (
	"github.com/apache/servicecomb-service-center/pkg/rest"
	"github.com/go-chassis/foundation/gopool"
)

var apiServer *APIServer

func init() {
	apiServer = &APIServer{
		isClose:   true,
		err:       make(chan error, 1),
		goroutine: gopool.New(gopool.Configure().Workers(5)),
	}
}

type APIServer struct {
	HostPort   string
	HTTPServer *rest.Server

	isClose   bool
	forked    bool
	err       chan error
	goroutine *gopool.Pool
}

func (s *APIServer) Err() <-chan error { _ = "STUB: not implemented"; return nil }

func (s *APIServer) graceDone() { _ = "STUB: not implemented"; return }

func (s *APIServer) MarkForked() { _ = "STUB: not implemented"; return }

func (s *APIServer) SetHostPort(ip, port string) { _ = "STUB: not implemented"; return }

func (s *APIServer) serve() (err error) { _ = "STUB: not implemented"; return nil }

func (s *APIServer) Start() { _ = "STUB: not implemented"; return }

// 自注册

func (s *APIServer) Stop() { _ = "STUB: not implemented"; return }

func (s *APIServer) selfRegister() { _ = "STUB: not implemented"; return }

// report the metrics

func (s *APIServer) selfUnregister() { _ = "STUB: not implemented"; return }

func GetAPIServer() *APIServer { _ = "STUB: not implemented"; return nil }
