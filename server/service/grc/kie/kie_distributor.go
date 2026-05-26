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

package kie

import (
	"context"

	"github.com/go-chassis/kie-client"

	"github.com/apache/servicecomb-service-center/pkg/gov"
	"github.com/apache/servicecomb-service-center/server/config"
	grcsvc "github.com/apache/servicecomb-service-center/server/service/grc"
)

const Priority = -1

type Distributor struct {
	name   string
	client *kie.Client
}

func (d *Distributor) Create(ctx context.Context, kind, project string, p *gov.Policy) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Distributor) Update(ctx context.Context, kind, id, project string, p *gov.Policy) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Distributor) Delete(ctx context.Context, kind, id, project string) error {
	_ = "STUB: not implemented"
	return nil
}

// should remove all policies of this group

func (d *Distributor) DeleteMatchGroup(ctx context.Context, id string, project string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Distributor) Display(ctx context.Context, project, app, env string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setAliasIfEmpty(spec map[string]interface{}, name string) { _ = "STUB: not implemented"; return }

func (d *Distributor) List(ctx context.Context, kind, project, app, env string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Distributor) Get(ctx context.Context, kind, id, project string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Distributor) getPolicy(ctx context.Context, kind string, id string, project string) (*gov.Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Distributor) Type() string { _ = "STUB: not implemented"; return "" }

func (d *Distributor) Name() string { _ = "STUB: not implemented"; return "" }

func initClient(endpoint string) *kie.Client { _ = "STUB: not implemented"; return nil }

func kieDistributorNew(opts config.DistributorOptions) (grcsvc.ConfigDistributor, error) {
	_ = "STUB: not implemented"
	return *new(grcsvc.ConfigDistributor), nil
}

func (d *Distributor) listDataByKind(ctx context.Context, kind, project, app, env string) (*kie.KVResponse, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (d *Distributor) generateID(ctx context.Context, project string, p *gov.Policy) error {
	_ = "STUB: not implemented"
	return nil
}

func getID() string { _ = "STUB: not implemented"; return "" }

func (d *Distributor) transform(kv *kie.KVDoc, kind string) (*gov.Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toGovKeyPrefix(kind string) string { _ = "STUB: not implemented"; return "" }

func init() {
	grcsvc.InstallDistributor(grcsvc.ConfigDistributorKie, kieDistributorNew)
}
