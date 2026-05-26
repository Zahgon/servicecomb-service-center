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

package servicecenter

import (
	"sync"
	"time"

	"github.com/apache/servicecomb-service-center/istio/pkg/event"

	"github.com/go-chassis/cari/discovery"
	"github.com/go-chassis/sc-client"
)

const (
	// Time in seconds to wait before re-registering a deleted Servicecomb Service Center Watcher service
	REREGISTER_INTERVAL time.Duration = time.Second * 5
)

// Servicecomb Service Center go-chassis client
type Connector struct {
	client                  *sc.Client
	AppInstanceWatcherCache sync.Map // Maps appId to id of a instance watcher service. Need app-specific watchers to avoid cross-app errors.
}

func NewConnector(addr string) *Connector { _ = "STUB: not implemented"; return nil }

// Check whether a service center MicroService exists in the registry.
func (c *Connector) GetServiceExistence(microServiceId string) bool {
	_ = "STUB: not implemented"
	return false
}

// Retrieve all service center MicroServices, without their instances, from the registry.
func (c *Connector) GetAllServices() ([]*discovery.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Register a new service center Watcher service that watches instance-level change events for all service center services sharing a specific appId.
func (c *Connector) RegisterAppInstanceWatcher(name string, appId string, callback func(event event.ChangeEvent)) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Need to reregister existing watcher for this app to reestablish the websocket connection

// Sleep allows time for service center to unregister watcher consumer/producer relationships
// Re-registering too early will cause race conditions

// Cache the id of the app instance watcher service

// Unregister a service center Watcher service.
func (c *Connector) UnregisterInstanceWatcher(serviceId string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetServiceInstances fetch newly received service instances.
func (c *Connector) GetServiceInstances(entries []*event.MicroserviceEntry) map[string][]*discovery.MicroServiceInstance {
	_ = "STUB: not implemented"
	return nil
}

// Initial instance sync of services with same appId, will use app instance watcher for future instance updates

// Save the ids of currently active service center Watcher services mapped to the appIds that they are responsible for.
func (c *Connector) RefreshAppInstanceWatcherCache(appWatcherIds sync.Map) {
	_ = "STUB: not implemented"
	return
}
