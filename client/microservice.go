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
	"net/url"

	pb "github.com/go-chassis/cari/discovery"
	"github.com/go-chassis/cari/pkg/errsvc"
)

const (
	apiExistenceURL     = "/v4/%s/registry/existence"
	apiMicroServicesURL = "/v4/%s/registry/microservices"
	apiMicroServiceURL  = "/v4/%s/registry/microservices/%s"
)

func (c *Client) CreateService(ctx context.Context, domain, project string, service *pb.MicroService) (string, *errsvc.Error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Client) DeleteService(ctx context.Context, domain, project, serviceID string) *errsvc.Error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) ServiceExistence(ctx context.Context, domain, project string, appID, serviceName, versionRule, env string) (string, *errsvc.Error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Client) existence(ctx context.Context, domain, project string, query url.Values) (*pb.GetExistenceResponse, *errsvc.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}
