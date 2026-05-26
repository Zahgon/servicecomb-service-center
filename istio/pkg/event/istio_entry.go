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
	istioAPI "istio.io/api/networking/v1alpha3"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
)

// An Istio WorkloadEntry and the known ports of its parent ServiceEntry.
type WorkloadEntry struct {
	*istioAPI.WorkloadEntry
	// A subset of ports used by the WorkloadEntry's parent ServiceEntry.
	// Used in conjunction with other WorkloadEntry's ServicePorts to construct a complete set of the ServiceEntry's ports.
	ServicePorts []*istioAPI.Port
}

// Convert an Istio WorkloadEntry to a service center Microservice Instance event.
func (c *WorkloadEntry) Convert() *InstanceEntry { _ = "STUB: not implemented"; return nil }

// WorkloadEntry was previously converted from service center Microservice instance; restore its id

// Convert an Istio host/endpoint's address and ports to a service center Microservice Instance entry.
func convertIstioAddressToMicroserviceInstance(address string, ports []*istioAPI.Port) *InstanceEntry {
	_ = "STUB: not implemented"
	return nil

	// Create equivalent Microservice endpoints for Istio address+port combinations
}

// we.Address could be a hostname if using DNS

// An Istio ServiceEntry and its associated WorkloadEntry(s).
type ServiceEntry struct {
	// An Istio ServiceEntry struct supported by the Istio client library.
	ServiceEntry *v1alpha3.ServiceEntry
	// WorkloadEntry(s) representing instances of the Istio ServiceEntry.
	WorkloadEntries []*WorkloadEntry
}

func NewServiceEntry(curr *v1alpha3.ServiceEntry) *ServiceEntry {
	_ = "STUB: not implemented"
	return nil
}

// Record ServiceEntry's ports in each WorkloadEntry

// Convert an Istio ServiceEntry to a MicroService event
func (s *ServiceEntry) Convert() *MicroserviceEntry { _ = "STUB: not implemented"; return nil }

// Convert a ServiceEntry with WorkloadEntry(s)

// Convert a ServiceEntry with Host(s) and Port(s) only

// Construct a Service center instance using only a host address and its ports

// Reports whether a protocol string is supported by Service center
func isMicroserviceRestProtocol(protocol string) bool { _ = "STUB: not implemented"; return false }

// Construct an endpoint for a Service center MicroServiceInstance from an Istio address and port
func getMicroserviceEndpointFromIstio(address string, port *istioAPI.Port) string {
	_ = "STUB: not implemented"
	return ""
}

// Enable SSL for the Service center REST endpoint

// ServiceEntry's "internal" port used by its corresponding WorkloadEntry

// Use "external" ServiceEntry port instead

// Service center endpoint string uses format <protocol>://<host>:<port>

// Construct a Service center service name from an Istio service name
func getMicroserviceNameFromIstio(svcName string) string { _ = "STUB: not implemented"; return "" }

// No need to return titlecased service names
// (example: https://github.com/apache/servicecomb-service-center/blob/6f26aaa7698691d40e17c6644ac71d51b6770772/integration/microservices_test.go#L592)
