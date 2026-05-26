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

package event

import (
	"net/url"

	"github.com/go-chassis/cari/discovery"
	istioAPI "istio.io/api/networking/v1alpha3"
)

const (
	// microservice protocol used for REST endpoints, equivalent to HTTP/HTTPS.
	RestProtocol = "rest"
	// Istio label corresponding to id of original microservice instance before being converted to WorkloadEntry.
	InstanceIdLabel = "instanceId"
	// microservice endpoint query string that enables SSL.
	EnableSSL = "sslEnabled=true"
)

type ChangeEvent struct {
	// Type of change event
	Action discovery.EventType
	// Event payload
	Event
}

type InstanceEntry struct {
	*discovery.MicroServiceInstance
}

// A Service center MicroService and its associated instance(s).
type MicroserviceEntry struct {
	// Microservice struct
	MicroService *discovery.MicroService
	// Instances of the MicroService
	Instances []*InstanceEntry
}

// Convert a MicroServiceInstance to an Istio WorkloadEntry event.
func (c *InstanceEntry) Convert() *WorkloadEntry {
	_ = "STUB: not implemented"
	// Istio ServiceEntry port names mapped to the "internal" port number of their target WorkloadEntry
	return nil
}

// Istio ServiceEntry port names mapped to port structs

// Only using first endpoint's hostname as the WorkloadEntry's address, Service center Microservice expects all endpoints for an instance to share a hostname

// Name for Istio ports are set to <protocol>-<portNumber>

// Convert ServiceEntry port map to array

// Convert a MicroService to an Istio ServiceEntry.
func (c *MicroserviceEntry) Convert() *ServiceEntry { _ = "STUB: not implemented"; return nil }

// Ensures that only one port struct exists for each port name

// Convert port map to array representing all ports exposed by ServiceEntry

func getIstioProtocolFromURL(u url.URL) string { _ = "STUB: not implemented"; return "" }

// Microservice uses query string to signify HTTPS endpoint

// Construct an Istio ServiceEntry port.
func newIstioPort(name string, protocol string, number uint32) *istioAPI.Port {
	_ = "STUB: not implemented"
	return nil
}
