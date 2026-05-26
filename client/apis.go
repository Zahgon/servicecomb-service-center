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

	"github.com/apache/servicecomb-service-center/pkg/dump"
	"github.com/apache/servicecomb-service-center/pkg/util"
	"github.com/apache/servicecomb-service-center/version"
	"github.com/go-chassis/cari/pkg/errsvc"
	"github.com/little-cui/etcdadpt"
)

const (
	apiVersionURL  = "/version"
	apiDumpURL     = "/v4/default/admin/dump"
	apiClustersURL = "/v4/default/admin/clusters"
	apiHealthURL   = "/v4/default/registry/health"

	QueryGlobal util.CtxKey = "global"
)

func (c *Client) toError(body []byte) *errsvc.Error { _ = "STUB: not implemented"; return nil }

func (c *Client) parseQuery(ctx context.Context) (q string) { _ = "STUB: not implemented"; return "" }

func (c *Client) GetScVersion(ctx context.Context) (*version.Set, *errsvc.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) GetScCache(ctx context.Context) (*dump.Cache, *errsvc.Error) {
	_ = "STUB: not implemented"
	return nil, nil

	// only default domain has admin permission
}

func (c *Client) GetClusters(ctx context.Context) (etcdadpt.Clusters, *errsvc.Error) {
	_ = "STUB: not implemented"
	return *new(etcdadpt.Clusters), nil
}

// only default domain has admin permission

func (c *Client) HealthCheck(ctx context.Context) *errsvc.Error {
	_ = "STUB: not implemented"
	return nil
}

// only default domain has admin permission
