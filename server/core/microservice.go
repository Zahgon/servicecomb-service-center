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

package core

import (
	"context"

	"github.com/apache/servicecomb-service-center/pkg/util"
	"github.com/go-chassis/cari/discovery"
)

var (
	Service  = &discovery.MicroService{}
	Instance = &discovery.MicroServiceInstance{}
)

const (
	RegistryServiceName  = "SERVICECENTER"
	RegistryServiceAlias = "SERVICECENTER"

	RegistryDefaultLeaseRenewalInterval int32 = 30
	RegistryDefaultLeaseRetryTimes      int32 = 3

	CtxScSelf util.CtxKey = "_sc_self"
)

func InitRegistration() { _ = "STUB: not implemented"; return }

func getEndpoints() []string { _ = "STUB: not implemented"; return nil }

func RegisterGlobalServices() { _ = "STUB: not implemented"; return }

func IsSCInstance(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func GetExistenceRequest() *discovery.GetExistenceRequest { _ = "STUB: not implemented"; return nil }

func GetServiceRequest(serviceID string) *discovery.GetServiceRequest {
	_ = "STUB: not implemented"
	return nil
}

func CreateServiceRequest() *discovery.CreateServiceRequest { _ = "STUB: not implemented"; return nil }

func RegisterInstanceRequest() *discovery.RegisterInstanceRequest {
	_ = "STUB: not implemented"
	return nil
}

func UnregisterInstanceRequest() *discovery.UnregisterInstanceRequest {
	_ = "STUB: not implemented"
	return nil
}

func HeartbeatRequest() *discovery.HeartbeatRequest { _ = "STUB: not implemented"; return nil }
