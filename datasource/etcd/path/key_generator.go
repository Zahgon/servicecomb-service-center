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

package path

import (
	"github.com/go-chassis/cari/discovery"
)

const (
	SPLIT                    = "/"
	RegistryRootKey          = "cse-sr"
	RegistrySysKey           = "sys"
	RegistryServiceKey       = "ms"
	RegistryInstanceKey      = "inst"
	RegistryFile             = "files"
	RegistryIndex            = "indexes"
	RegistryDomainKey        = "domains"
	RegistryProjectKey       = "projects"
	RegistryAliasKey         = "alias"
	RegistryTagKey           = "tags"
	RegistrySchemaRefKey     = "schema-ref"
	RegistrySchemaContentKey = "schema-content"
	RegistrySchemaKey        = "schemas"
	RegistrySchemaSummaryKey = "schema-sum"
	RegistryLeaseKey         = "leases"
	RegistryDepsRuleKey      = "dep-rules"
	RegistryDepsQueueKey     = "dep-queue"
	RegistryMetricsKey       = "metrics"
	DepsQueueUUID            = "0"
	DepsConsumer             = "c"
	DepsProvider             = "p"
)

func GetRootKey() string { _ = "STUB: not implemented"; return "" }

func GenerateDomainKey(domain string) string { _ = "STUB: not implemented"; return "" }

func GetProjectRootKey(domain string) string { _ = "STUB: not implemented"; return "" }

func GenerateProjectKey(domain, project string) string { _ = "STUB: not implemented"; return "" }

func GenerateRBACAccountKey(name string) string { _ = "STUB: not implemented"; return "" }

func GenerateRBACRoleKey(name string) string { _ = "STUB: not implemented"; return "" }

func GenRoleAccountIdxKey(role, account string) string { _ = "STUB: not implemented"; return "" }

func GenRoleAccountPrefixIdxKey(role string) string { _ = "STUB: not implemented"; return "" }

func GetServiceRootKey(domainProject string) string { _ = "STUB: not implemented"; return "" }

func GenerateServiceKey(domainProject string, serviceID string) string {
	_ = "STUB: not implemented"
	return ""
}

func GetServiceIndexRootKey(domainProject string) string { _ = "STUB: not implemented"; return "" }

func GetServiceAliasRootKey(domainProject string) string { _ = "STUB: not implemented"; return "" }

func GetServiceAppKey(domainProject, env, appID string) string {
	_ = "STUB: not implemented"
	return ""
}

func GetServiceTagRootKey(domainProject string) string { _ = "STUB: not implemented"; return "" }

func GetServiceSchemaRootKey(domainProject string) string { _ = "STUB: not implemented"; return "" }

func GetServiceSchemaRefRootKey(domainProject string) string { _ = "STUB: not implemented"; return "" }

func GetServiceSchemaContentRootKey(domainProject string) string {
	_ = "STUB: not implemented"
	return ""
}

func GetInstanceRootKey(domainProject string) string { _ = "STUB: not implemented"; return "" }

func GetInstanceLeaseRootKey(domainProject string) string { _ = "STUB: not implemented"; return "" }

func GenerateServiceIndexKey(key *discovery.MicroServiceKey) string {
	_ = "STUB: not implemented"
	return ""
}

func GenerateServiceAliasKey(key *discovery.MicroServiceKey) string {
	_ = "STUB: not implemented"
	return ""
}

func GenerateServiceTagKey(domainProject string, serviceID string) string {
	_ = "STUB: not implemented"
	return ""
}

func GenerateServiceSchemaRefKey(domainProject string, serviceID string, schemaID string) string {
	_ = "STUB: not implemented"
	return ""
}

func GenerateServiceSchemaContentKey(domainProject string, hash string) string {
	_ = "STUB: not implemented"
	return ""
}

func GenerateServiceSchemaKey(domainProject string, serviceID string, schemaID string) string {
	_ = "STUB: not implemented"
	return ""
}

func GenerateServiceSchemaSummaryKey(domainProject string, serviceID string, schemaID string) string {
	_ = "STUB: not implemented"
	return ""
}

func GetServiceSchemaSummaryRootKey(domainProject string) string {
	_ = "STUB: not implemented"
	return ""
}

func GenerateInstanceKey(domainProject string, serviceID string, instanceID string) string {
	_ = "STUB: not implemented"
	return ""
}

func GenerateInstanceLeaseKey(domainProject string, serviceID string, instanceID string) string {
	_ = "STUB: not implemented"
	return ""
}

func GenerateServiceDependencyRuleKey(serviceType string, domainProject string, in *discovery.MicroServiceKey) string {
	_ = "STUB: not implemented"
	return ""
}

func GenerateConsumerDependencyRuleKey(domainProject string, in *discovery.MicroServiceKey) string {
	_ = "STUB: not implemented"
	return ""
}

func GenerateProviderDependencyRuleKey(domainProject string, in *discovery.MicroServiceKey) string {
	_ = "STUB: not implemented"
	return ""
}

func GetServiceDependencyRuleRootKey(domainProject string) string {
	_ = "STUB: not implemented"
	return ""
}

func GetServiceDependencyQueueRootKey(domainProject string) string {
	_ = "STUB: not implemented"
	return ""
}

func GenerateConsumerDependencyQueueKey(domainProject, consumerID, uuid string) string {
	_ = "STUB: not implemented"
	return ""
}

func GenerateAccountKey(name string) string { _ = "STUB: not implemented"; return "" }

func GenerateAccountLockKey(key string) string { _ = "STUB: not implemented"; return "" }

func GenerateRBACSecretKey() string { _ = "STUB: not implemented"; return "" }

func GetServerInfoKey() string { _ = "STUB: not implemented"; return "" }

func GetMetricsRootKey() string { _ = "STUB: not implemented"; return "" }

func GenerateMetricsKey(name, utc, domain string) string { _ = "STUB: not implemented"; return "" }
