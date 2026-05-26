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

package dao

import (
	"context"

	"github.com/go-chassis/cari/discovery"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/apache/servicecomb-service-center/datasource/mongo/model"
)

func GetServiceByID(ctx context.Context, serviceID string) (*model.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetService(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (*model.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetServiceID(ctx context.Context, key *discovery.MicroServiceKey) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getServiceID(ctx context.Context, filter bson.M) (serviceID string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetServices(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]*model.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetMicroServices(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]*discovery.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UpdateService(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// means no doc find, if the operation is update,should return err

func CountService(ctx context.Context, filter interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
