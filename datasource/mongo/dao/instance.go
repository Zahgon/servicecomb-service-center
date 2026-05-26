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
	"github.com/go-chassis/cari/pkg/errsvc"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/apache/servicecomb-service-center/datasource/mongo/model"
)

func GetInstances(ctx context.Context, filter interface{}) ([]*model.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetMicroServiceInstancesByID(ctx context.Context, serviceID string) ([]*discovery.MicroServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetMicroServiceInstances(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]*discovery.MicroServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CountInstance(ctx context.Context, filter interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func UpdateInstance(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) *errsvc.Error {
	_ = "STUB: not implemented"
	return nil
}

func ExistInstance(ctx context.Context, serviceID string, instanceID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
