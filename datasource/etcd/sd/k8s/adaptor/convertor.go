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

package adaptor

import (
	pb "github.com/go-chassis/cari/discovery"
	v1 "k8s.io/api/core/v1"

	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
)

func getLabel(labels map[string]string, key, def string) string {
	_ = "STUB: not implemented"
	return ""
}

func getRegionAZ(node *v1.Node) (string, string) { _ = "STUB: not implemented"; return "", "" }

func getFullName(namespace, name string) string { _ = "STUB: not implemented"; return "" }

func getProtocol(port v1.EndpointPort) (string, bool) { _ = "STUB: not implemented"; return "", false }

func generateEndpoint(ip string, port v1.EndpointPort) string { _ = "STUB: not implemented"; return "" }

func generateServiceKey(domainProject string, svc *v1.Service) *pb.MicroServiceKey {
	_ = "STUB: not implemented"
	return nil
}

func FromK8sService(domainProject string, svc *v1.Service) (ms *pb.MicroService) {
	_ = "STUB: not implemented"
	return nil
}

func AsKeyValue(key string, v interface{}, resourceVersion string) *kvstore.KeyValue {
	_ = "STUB: not implemented"
	return nil
}
