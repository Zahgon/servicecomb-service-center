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
	"context"
	"sync"

	"github.com/apache/servicecomb-service-center/istio/pkg/event"
	"github.com/go-chassis/cari/discovery"
)

type Controller struct {
	// servicecomb service center go-chassis API client
	conn *Connector
	// Channel used to send and receive servicecomb service center change events from the service center controller
	events chan []event.ChangeEvent
	// Cache of retrieved servicecomb service center microservices, mapped to their service ids
	serviceCache sync.Map
}

func NewController(addr string, e chan []event.ChangeEvent) *Controller {
	_ = "STUB: not implemented"
	return nil
}

// Run until a stop signal is received
func (c *Controller) Run(ctx context.Context) {
	_ = "STUB: not implemented"
	// start a new go routine to watch service center update
	return
}

// Stop the controller.
func (c *Controller) Stop() {
	_ = "STUB: not implemented"
	// Unregister app instance watcher services
	return
}

// Watch the service center registry for MicroService changes.
func (c *Controller) watchServiceCenter(ctx context.Context) { _ = "STUB: not implemented"; return }

// Full sync all services from service center

// Process received services

// Number of seconds to wait between syncs

// Send all new services to Istio controller.
func (c *Controller) onServiceCenterUpdate(services []*discovery.MicroService) {
	_ = "STUB: not implemented"
	return
}

// Send all instance events to Istio controller.
func (c *Controller) onInstanceUpdate(e event.ChangeEvent) { _ = "STUB: not implemented"; return }

// Get service center service changes, register watcher for newly created services.
func (c *Controller) getChangedServices(services []*discovery.MicroService) []event.ChangeEvent {
	_ = "STUB: not implemented"
	// All non-watcher service center services mapped to their ids
	return nil
}

// All new non-watcher service center services

// IDs of current service center watcher services

// Service events that must be pushed

// Register new app instance watcher service

// Record the id of the watcher service for this app

// Collect newly created service

// Collect updated service

// No change, keep cache entry

// Watcher still exists as expected, record its current id

// Collect deleted services

// Initial sync-up for newly created services; retrieve and start watching their instances

// Update service ID cache with current services

// Check for app instance watcher changes

// Save MicroService(s) retrieved from service center registry
func (c *Controller) refreshServiceCache(services sync.Map) { _ = "STUB: not implemented"; return }

// Detect missing watcher services in registry. If a watcher service was expected but is missing, flag it to be re-registered.
func (c *Controller) checkAppInstanceWatchers(currAppInstanceWatcherIds sync.Map) {
	_ = "STUB: not implemented"
	return
}

// Watcher is missing for this app, remove all app's services from cache

// Cache current watcher ids (if any are missing, will be re-registered on next sync)

// Watch services, has side effect of adding instances to MicroserviceEntry(s)
func (c *Controller) initNewServices(newServices []*event.MicroserviceEntry) {
	_ = "STUB: not implemented"
	return
}
