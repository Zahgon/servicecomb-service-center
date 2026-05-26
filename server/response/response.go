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

package response

import (
	"github.com/go-chassis/cari/discovery"
)

func init() {
	RegisterFilter("/v4/:project/registry/microservices", MicroserviceListFilter)
	RegisterFilter("/v4/:project/registry/microservices/:providerId/consumers", ProvidersListFilter)
	RegisterFilter("/v4/:project/registry/microservices/:consumerId/providers", ConsumersListFilter)
	// control panel apis
	RegisterFilter("/v4/:project/govern/microservices", MicroServiceInfoListFilter)
	RegisterFilter("/v4/:project/govern/apps", AppIDListFilter)
}

func MicroserviceListFilter(obj interface{}, labels []map[string]string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func matchOne(service *discovery.MicroService, labels map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

func filterMicroservices(sources []*discovery.MicroService, labelsList []map[string]string) []*discovery.MicroService {
	_ = "STUB: not implemented"
	return nil
}

func ProvidersListFilter(obj interface{}, labels []map[string]string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func ConsumersListFilter(obj interface{}, labels []map[string]string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func MicroServiceInfoListFilter(obj interface{}, labelsList []map[string]string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func AppIDListFilter(obj interface{}, labelsList []map[string]string) interface{} {
	_ = "STUB: not implemented"
	return nil
}
