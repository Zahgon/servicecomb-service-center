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

package buildin

import (
	"context"
	"net/http"

	"github.com/apache/servicecomb-service-center/server/plugin/auth"
)

const (
	LabelEnvironment = "environment"
	LabelAppID       = "appId"
	LabelServiceName = "serviceName"
	QueryEnv         = "env"
	HeaderConsumerID = "X-ConsumerId"
)

var (
	// APIServiceExistence Apply by service key or serviceId
	// - /v4/:project/registry/existence?env=xxx&appId=xxx&serviceName=xxx
	// - /v4/:project/registry/existence?serviceId=xxx&schemaId=xxx
	APIServiceExistence = "/v4/:project/registry/existence"
	// APIServicesList Method GET: apply all by optional service key
	// Method POST or DELETE: apply by request body
	APIServicesList     = "/v4/:project/registry/microservices"
	APIServiceInfo      = "/v4/:project/registry/microservices/:serviceId"
	APIProConDependency = "/v4/:project/registry/microservices/:providerId/consumers"
	APIConProDependency = "/v4/:project/registry/microservices/:consumerId/providers"
	// APIDiscovery Apply by service key
	APIDiscovery = "/v4/:project/registry/instances"
	// APIBatchDiscovery Apply by request body
	APIBatchDiscovery = "/v4/:project/registry/instances/action"
	// APIHeartbeats Apply by request body
	APIHeartbeats = "/v4/:project/registry/heartbeats"
	// APIGovServicesList Apply by optional service key
	// - /v4/:project/govern/microservices?appId=xxx
	// Apply all:
	// - /v4/:project/govern/microservices?options=statistics
	// - /v4/:project/govern/microservices/statistics
	APIGovServicesList = "/v4/:project/govern/microservices"
	APIGovServiceInfo  = "/v4/:project/govern/microservices/:serviceId"
)

func init() {
	RegisterParseFunc(APIServiceInfo, ByServiceID)
	RegisterParseFunc(APIGovServiceInfo, ByServiceID)
	RegisterParseFunc(APIProConDependency, func(r *http.Request) (*auth.ResourceScope, error) {
		return fromQueryKey(r, ":providerId")
	})
	RegisterParseFunc(APIConProDependency, func(r *http.Request) (*auth.ResourceScope, error) {
		return fromQueryKey(r, ":consumerId")
	})
	RegisterParseFunc(APIDiscovery, ByServiceKey)
	RegisterParseFunc(APIServiceExistence, ByServiceKey)
	RegisterParseFunc(APIGovServicesList, ApplyAll)
	RegisterParseFunc(APIServicesList, ByRequestBody)
	RegisterParseFunc(APIBatchDiscovery, ByDiscoveryRequestBody)
	RegisterParseFunc(APIHeartbeats, ByHeartbeatRequestBody)
}

func ByServiceID(r *http.Request) (*auth.ResourceScope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fromQueryKey(r *http.Request, queryKey string) (*auth.ResourceScope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func serviceIDToLabels(ctx context.Context, serviceID string) ([]map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ByServiceKey(r *http.Request) (*auth.ResourceScope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fromServiceKeyEnv(r *http.Request, queryEnv string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ByRequestBody(r *http.Request) (*auth.ResourceScope, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// get or list by query string
}

func fromRequestBody(r *http.Request) (*auth.ResourceScope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// batch delete

// create service

func createServiceToLabels(r *http.Request) ([]map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deleteServicesToLabels(r *http.Request) ([]map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ByDiscoveryRequestBody(r *http.Request) (*auth.ResourceScope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ByHeartbeatRequestBody(r *http.Request) (*auth.ResourceScope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
