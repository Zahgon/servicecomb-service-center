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

package mongo

import (
	"context"

	"github.com/go-chassis/cari/discovery"

	"github.com/apache/servicecomb-service-center/datasource/mongo/model"
)

func ExistInstance(ctx context.Context, serviceID string, instanceID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func GetInstances(ctx context.Context) ([]*model.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CountInstance(ctx context.Context, serviceID string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetAllInstancesOfOneService(ctx context.Context, serviceID string) ([]*discovery.MicroServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
