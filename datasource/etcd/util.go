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

package etcd

import (
	"context"

	pb "github.com/go-chassis/cari/discovery"
	"github.com/go-chassis/cari/pkg/errsvc"
	"github.com/little-cui/etcdadpt"

	"github.com/apache/servicecomb-service-center/datasource"
	serviceUtil "github.com/apache/servicecomb-service-center/datasource/etcd/util"
)

// schema
func getSchemaSummary(ctx context.Context, domainProject string, serviceID string, schemaID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getSchemasFromDatabase(ctx context.Context, domainProject string, serviceID string) ([]*pb.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkSchemaInfoExist(ctx context.Context, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func putSchema(ctx context.Context, domainProject string, serviceID string, schema *pb.Schema) ([]etcdadpt.OpOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deleteSchema(ctx context.Context, domainProject string, serviceID string, schema *pb.Schema) ([]etcdadpt.OpOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isExistSchemaID(service *pb.MicroService, schemas []*pb.Schema) bool {
	_ = "STUB: not implemented"
	return false
}

func commitSchemaInfo(ctx context.Context, domainProject string, serviceID string, schema *pb.Schema) ([]etcdadpt.OpOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getHeartbeatFunc(ctx context.Context, domainProject string, instancesHbRst chan<- *pb.InstanceHbRst, element *pb.HeartbeatSetElement) func(context.Context) {
	_ = "STUB: not implemented"
	return nil
}

func revokeInstance(ctx context.Context, domainProject string, serviceID string, instanceID string) *errsvc.Error {
	_ = "STUB: not implemented"
	return nil
}

func statistics(ctx context.Context, withShared bool) (*pb.Statistics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// services

// instance

func getInstanceCountByDomain(ctx context.Context, svcIDToNonVerKey map[string]string, resp chan datasource.GetInstanceCountByDomainResponse) {
	_ = "STUB: not implemented"
	return
}

// dep util
func toDependencyFilterOptions(in *pb.GetDependenciesRequest) (opts []serviceUtil.DependencyRelationFilterOption) {
	_ = "STUB: not implemented"
	return nil
}
