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

package disco

import (
	"context"

	pb "github.com/go-chassis/cari/discovery"
)

func RegisterService(ctx context.Context, request *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error) {
	_ = "STUB: not implemented"
	//create service
	return nil, nil
}

//create tag,rule,instances

func registerService(ctx context.Context, request *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func assignDefaultValue(service *pb.MicroService) { _ = "STUB: not implemented"; return }

func registerServiceDetails(ctx context.Context, in *pb.CreateServiceRequest, serviceID string) (*pb.CreateServiceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//create tags

// create instance

// handle result

func hasServiceDetails(in *pb.CreateServiceRequest) bool { _ = "STUB: not implemented"; return false }

func checkServiceQuota(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func UnregisterService(ctx context.Context, request *pb.DeleteServiceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func GetService(ctx context.Context, in *pb.GetServiceRequest) (*pb.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListService(ctx context.Context, in *pb.GetServicesRequest) (*pb.GetServicesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindService(ctx context.Context, in *pb.MicroServiceKey) (*pb.GetServicesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UnregisterManyService(ctx context.Context, request *pb.DelServicesRequest) (*pb.DelServicesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 合法性检查

// 批量删除服务

//ServiceId重复性检查

//执行删除服务操作

//获取批量删除服务的结果

//结果收集over，关闭通道

func getDeleteServiceFunc(ctx context.Context, serviceID string, force bool, serviceRespChan chan<- *pb.DelServicesRspInfo) func(context.Context) {
	_ = "STUB: not implemented"
	return nil
}

func ExistService(ctx context.Context, in *pb.GetExistenceRequest) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func PutServiceProperties(ctx context.Context, request *pb.UpdateServicePropsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ServiceUsage(ctx context.Context, request *pb.GetServiceCountRequest) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
