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

package datasource

import (
	"github.com/go-chassis/cari/discovery"
)

var GlobalServiceNames = make(map[string]struct{})

type GetInstanceCountByDomainResponse struct {
	Err           error
	CountByDomain int64
}

func SetServiceDefaultValue(service *discovery.MicroService) { _ = "STUB: not implemented"; return }

// SetStaticServices calculate the service/application num under a domainProject
func SetStaticServices(statistics *discovery.Statistics, svcKeys []*discovery.MicroServiceKey, svcIDs []string, withShared bool) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// SetStaticInstances calculate the instance/onlineService num under a domainProject
func SetStaticInstances(statistics *discovery.Statistics, svcIDToNonVerKey map[string]string, instServiceIDs []string) {
	_ = "STUB: not implemented"
	return
}

func generateServiceKey(key *discovery.MicroServiceKey) string {
	_ = "STUB: not implemented"
	return ""
}

func TransServiceToKey(domainProject string, service *discovery.MicroService) *discovery.MicroServiceKey {
	_ = "STUB: not implemented"
	return nil
}

func RegisterGlobalService(serviceName string) { _ = "STUB: not implemented"; return }

func IsGlobal(key *discovery.MicroServiceKey) bool { _ = "STUB: not implemented"; return false }

func RemoveGlobalServices(withShared bool, domainProject string,
	services []*discovery.MicroService) []*discovery.MicroService {
	_ = "STUB: not implemented"
	return nil
}

func IsDefaultDomainProject(domainProject string) bool { _ = "STUB: not implemented"; return false }
