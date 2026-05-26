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

func splitKey(key []byte) (keys []string) { _ = "STUB: not implemented"; return nil }

func getLast2Keys(key []byte) (string, string) { _ = "STUB: not implemented"; return "", "" }

func getLast3Keys(key []byte) (string, string, string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

func GetInfoFromSvcKV(key []byte) (serviceID, domainProject string) {
	_ = "STUB: not implemented"
	return "", ""
}

func GetInfoFromInstKV(key []byte) (serviceID, instanceID, domainProject string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

func GetInfoFromDomainKV(key []byte) (domain string) { _ = "STUB: not implemented"; return "" }

func GetInfoFromProjectKV(key []byte) (domain, project string) {
	_ = "STUB: not implemented"
	return "", ""
}

func GetInfoFromTagKV(key []byte) (serviceID, domainProject string) {
	_ = "STUB: not implemented"
	return "", ""
}

func GetInfoFromSvcIndexKV(key []byte) *discovery.MicroServiceKey {
	_ = "STUB: not implemented"
	return nil
}

func GetInfoFromSvcAliasKV(key []byte) *discovery.MicroServiceKey {
	_ = "STUB: not implemented"
	return nil
}

func GetInfoFromSchemaRefKV(key []byte) (domainProject, serviceID, schemaID string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

func GetInfoFromSchemaSummaryKV(key []byte) (domainProject, serviceID, schemaID string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

func GetInfoFromSchemaKV(key []byte) (domainProject, serviceID, schemaID string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

func GetInfoFromSchemaContentKV(key []byte) (domainProject, hash string) {
	_ = "STUB: not implemented"
	return "", ""
}

func GetInfoFromDependencyQueueKV(key []byte) (consumerID, domainProject, uuid string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

func GetInfoFromDependencyRuleKV(key []byte) (t string, _ *discovery.MicroServiceKey) {
	_ = "STUB: not implemented"
	return "", nil
}

func SplitDomainProject(domainProject string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}
