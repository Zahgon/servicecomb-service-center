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

package mock

import (
	"context"

	"github.com/apache/servicecomb-service-center/pkg/gov"
	"github.com/apache/servicecomb-service-center/server/config"
	grcsvc "github.com/apache/servicecomb-service-center/server/service/grc"
)

type Distributor struct {
	lbPolicies map[string]*gov.Policy
	name       string
}

const MatchGroup = "match-group"

func (d *Distributor) Create(_ context.Context, kind, _ string, p *gov.Policy) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Distributor) Update(_ context.Context, kind, id, _ string, p *gov.Policy) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Distributor) Delete(_ context.Context, _, id, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Distributor) Display(_ context.Context, _, app, env string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Distributor) List(_ context.Context, kind, _, app, env string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkPolicy(g *gov.Policy, kind, app, env string) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *Distributor) Get(_ context.Context, _, id, _ string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Distributor) Type() string { _ = "STUB: not implemented"; return "" }

func (d *Distributor) Name() string { _ = "STUB: not implemented"; return "" }

func newMock(opts config.DistributorOptions) (grcsvc.ConfigDistributor, error) {
	_ = "STUB: not implemented"
	return *new(grcsvc.ConfigDistributor), nil
}

func init() {
	grcsvc.InstallDistributor(grcsvc.ConfigDistributorMock, newMock)
}
