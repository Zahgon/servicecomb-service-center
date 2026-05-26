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

package util

import (
	"context"

	pb "github.com/go-chassis/cari/discovery"
	"github.com/little-cui/etcdadpt"

	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
)

/*
get Service by service id
*/
func GetServiceWithRev(ctx context.Context, domain string, id string, rev int64) (*pb.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetService(ctx context.Context, domainProject string, serviceID string) (*pb.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getServicesRawData(ctx context.Context, domainProject string) ([]*kvstore.KeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAllServicesAcrossDomainProject get services of all domains, projects
// the map's key is domainProject
func GetAllServicesAcrossDomainProject(ctx context.Context) (map[string][]*pb.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetServicesByDomainProject(ctx context.Context, domainProject string) ([]*pb.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetServiceID(ctx context.Context, key *pb.MicroServiceKey) (serviceID string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// 别名查询

func searchServiceID(ctx context.Context, key *pb.MicroServiceKey) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func searchServiceIDFromAlias(ctx context.Context, key *pb.MicroServiceKey) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetServiceAllVersions(ctx context.Context, key *pb.MicroServiceKey, alias bool) (*kvstore.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindServiceIds return serviceIDs match the key, the existence of the micro-service without consider of version
func FindServiceIds(ctx context.Context, key *pb.MicroServiceKey, matchVersion bool) ([]string, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func ServiceExist(ctx context.Context, domainProject string, serviceID string) bool {
	_ = "STUB: not implemented"
	return false
}

func GetAllServiceUtil(ctx context.Context) ([]*pb.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UpdateService(ctx context.Context, domainProject string, serviceID string, service *pb.MicroService) ([]etcdadpt.OpOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetOneDomainProjectServiceCount(ctx context.Context, domainProject string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetOneDomainProjectInstanceCount(ctx context.Context, domainProject string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetGlobalInstanceCount(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetGlobalServiceIDs(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetGlobalServiceCount(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getGlobalEnvironment() string { _ = "STUB: not implemented"; return "" }
