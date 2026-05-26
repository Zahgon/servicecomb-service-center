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

package govern

import (
	"context"

	pb "github.com/go-chassis/cari/discovery"
)

var defaultOptions = []string{"tags", "instances", "schemas", "dependencies"}

type ServiceDetailOpt struct {
	domainProject string
	service       *pb.MicroService
	countOnly     bool
	options       []string
}

func ListServiceDetail(ctx context.Context, in *pb.GetServicesInfoRequest) (*pb.GetServicesInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//获取所有服务

func getServiceDetailUtil(ctx context.Context, opts ServiceDetailOpt) (*pb.ServiceDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterServices(domainProject string, request *pb.GetServicesInfoRequest, service *pb.MicroService) bool {
	_ = "STUB: not implemented"
	return false
}

func matchAllProperties(properties map[string]string, service *pb.MicroService) bool {
	_ = "STUB: not implemented"
	return false
}

func NewServiceOverview(serviceDetail *pb.ServiceDetail, innerProperties map[string]string) (*pb.ServiceDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func removeCustomProperties(properties, innerProperties map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func GetServiceDetail(ctx context.Context, in *pb.GetServiceRequest) (*pb.ServiceDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getServiceAllVersions(ctx context.Context, key *pb.MicroServiceKey) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListApp(ctx context.Context, in *pb.GetAppsRequest) (*pb.GetAppsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetOverview(ctx context.Context, in *pb.GetServicesRequest) (*pb.Statistics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
