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

package client

import (
	"context"

	"github.com/go-chassis/cari/discovery"
	"github.com/go-chassis/cari/pkg/errsvc"
)

const (
	apiDiscoveryInstancesURL = "/v4/%s/registry/instances"
	apiHeartbeatSetURL       = "/v4/%s/registry/heartbeats"
	apiInstancesURL          = "/v4/%s/registry/microservices/%s/instances"
	apiInstanceURL           = "/v4/%s/registry/microservices/%s/instances/%s"
	apiInstanceHeartbeatURL  = "/v4/%s/registry/microservices/%s/instances/%s/heartbeat"
)

func (c *Client) RegisterInstance(ctx context.Context, domain, project, serviceID string, instance *discovery.MicroServiceInstance) (string, *errsvc.Error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Client) UnregisterInstance(ctx context.Context, domain, project, serviceID, instanceID string) *errsvc.Error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DiscoveryInstances(ctx context.Context, domain, project, consumerID, providerAppID, providerServiceName, providerVersionRule string) ([]*discovery.MicroServiceInstance, *errsvc.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Heartbeat(ctx context.Context, domain, project, serviceID, instanceID string) *errsvc.Error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) HeartbeatSet(ctx context.Context, domain, project string, instances ...*discovery.HeartbeatSetElement) ([]*discovery.InstanceHbRst, *errsvc.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) GetInstancesByServiceID(ctx context.Context, domain, project, providerID, consumerID string) ([]*discovery.MicroServiceInstance, *errsvc.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) GetInstanceByInstanceID(ctx context.Context, domain, project, providerID, instanceID, consumerID string) (*discovery.MicroServiceInstance, *errsvc.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}
