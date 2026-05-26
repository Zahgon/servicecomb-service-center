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

package cache

import (
	"context"

	"github.com/go-chassis/cari/discovery"

	"github.com/apache/servicecomb-service-center/datasource/mongo/model"
)

func GetProviderServiceOfDeps(provider *discovery.MicroService) (*discovery.MicroServiceDependency, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func transCacheToDep(cache []interface{}) ([]*discovery.MicroServiceDependency, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func genDepServiceKey(ruleType string, service *discovery.MicroService) string {
	_ = "STUB: not implemented"
	return ""
}

func GetMicroServiceInstancesByID(ctx context.Context, serviceID string) ([]*discovery.MicroServiceInstance, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func GetServiceByID(ctx context.Context, serviceID string) (*model.Service, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func GetServiceByName(ctx context.Context, key *discovery.MicroServiceKey) ([]*model.Service, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func GetServiceID(ctx context.Context, key *discovery.MicroServiceKey) (serviceID string, exist bool) {
	_ = "STUB: not implemented"
	return "", false
}

func GetServiceByIDAcrossDomain(ctx context.Context, serviceID string) (*model.Service, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func GetServicesByDomainProject(domainProject string) (service []*model.Service, exist bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func GetMicroServicesByDomainProject(domainProject string) (service []*discovery.MicroService, exist bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func transCacheToService(services []interface{}) ([]*model.Service, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func genServiceIDIndexAcrossDomain(ctx context.Context, serviceID string) string {
	_ = "STUB: not implemented"
	return ""
}

func genServiceIDIndex(ctx context.Context, serviceID string) string {
	_ = "STUB: not implemented"
	return ""
}

func genServiceKeyIndex(ctx context.Context, key *discovery.MicroServiceKey) string {
	_ = "STUB: not implemented"
	return ""
}

func genServiceNameIndex(ctx context.Context, key *discovery.MicroServiceKey) string {
	_ = "STUB: not implemented"
	return ""
}

func CountInstances(ctx context.Context, serviceID string) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func GetInstance(ctx context.Context, serviceID string, instanceID string) (*model.Instance, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func GetInstances(ctx context.Context) ([]*model.Instance, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func transCacheToMicroInsts(cache []interface{}) ([]*discovery.MicroServiceInstance, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func transCacheToInsts(cache []interface{}) ([]*model.Instance, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func generateInstanceIDIndex(domainProject string, serviceID string, instanceID string) string {
	_ = "STUB: not implemented"
	return ""
}
