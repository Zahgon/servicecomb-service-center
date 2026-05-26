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
	"crypto/tls"
	"sync"

	"github.com/apache/servicecomb-service-center/client"
	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
	"github.com/apache/servicecomb-service-center/pkg/dump"
	"github.com/go-chassis/cari/pkg/errsvc"
)

var (
	scClient   *SCClientAggregate
	clientOnce sync.Once
	clientTLS  *tls.Config
)

type SCClientAggregate []*client.Client

func getClientTLS() (*tls.Config, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *SCClientAggregate) GetScCache(ctx context.Context) (*dump.Cache, map[string]error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *SCClientAggregate) cacheAppend(name string, setter dump.Setter, getter dump.Getter) {
	_ = "STUB: not implemented"
	return
}

func (c *SCClientAggregate) GetSchemasByServiceID(ctx context.Context, domainProject, serviceID string) (*kvstore.Response, *errsvc.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *SCClientAggregate) GetSchemaBySchemaID(ctx context.Context, domainProject, serviceID, schemaID string) (*kvstore.Response, *errsvc.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *SCClientAggregate) GetInstancesByServiceID(ctx context.Context, domain, project, providerID, consumerID string) (*kvstore.Response, *errsvc.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *SCClientAggregate) GetInstanceByInstanceID(ctx context.Context, domain, project, providerID, instanceID, consumerID string) (*kvstore.Response, *errsvc.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetOrCreateSCClient() *SCClientAggregate { _ = "STUB: not implemented"; return nil }

// TODO should not use the etcd config

// TLS
