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
	"crypto/tls"

	nf "github.com/apache/servicecomb-service-center/pkg/event"
)

var sc ServiceCenterServer

func Run() { _ = "STUB: not implemented"; return }

type endpoint struct {
	Host string
	Port string
}

type ServiceCenterServer struct {
	Endpoint    endpoint
	APIServer   *APIServer
	eventCenter *nf.BusService
}

func (s *ServiceCenterServer) Run() { _ = "STUB: not implemented"; return }

func (s *ServiceCenterServer) startChassis() { _ = "STUB: not implemented"; return }

func (s *ServiceCenterServer) waitForQuit() { _ = "STUB: not implemented"; return }

func (s *ServiceCenterServer) initialize() {
	_ = "STUB: not implemented"

	// SSL
	return
}

// Datasource

func (s *ServiceCenterServer) initEndpoints() { _ = "STUB: not implemented"; return }

func (s *ServiceCenterServer) initDatasource() {
	_ = "STUB: not implemented"
	// init datasource
	return
}

func getDatasourceTLSConfig() (*tls.Config, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *ServiceCenterServer) initSSL() { _ = "STUB: not implemented"; return }

func (s *ServiceCenterServer) startServices() {
	_ = "STUB: not implemented"
	// notifications
	return
}

// load sc plugins

// check version

// api service

func (s *ServiceCenterServer) startAPIService() { _ = "STUB: not implemented"; return }

func (s *ServiceCenterServer) Stop() { _ = "STUB: not implemented"; return }
