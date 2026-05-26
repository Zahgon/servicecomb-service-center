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

func GetServiceByID(ctx context.Context, serviceID string) (*model.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetServiceByIDAcrossDomain(ctx context.Context, serviceID string) (*model.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ServiceExistID(ctx context.Context, serviceID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func GetAllMicroServicesByDomainProject(ctx context.Context) ([]*discovery.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetServiceID(ctx context.Context, key *discovery.MicroServiceKey) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
