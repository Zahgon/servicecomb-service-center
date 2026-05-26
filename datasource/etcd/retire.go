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

	"github.com/apache/servicecomb-service-center/datasource"
	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
)

const poolSizeOfRotation = 5

type RotateServiceIDKey struct {
	DomainProject string
	ServiceID     string
}

func (ds *MetadataManager) RetireService(ctx context.Context, plan *datasource.RetirePlan) error {
	_ = "STUB: not implemented"
	return nil
}

func FilterNoInstance(ctx context.Context, serviceIDKeys []*RotateServiceIDKey) []*RotateServiceIDKey {
	_ = "STUB: not implemented"
	return nil
}

func UnregisterManyService(ctx context.Context, serviceIDKeys []*RotateServiceIDKey) (deleted int64) {
	_ = "STUB: not implemented"
	return 0
}

func GetRetireServiceIDs(indexesResp *kvstore.Response, reserveVersionCount int) []*RotateServiceIDKey {
	_ = "STUB: not implemented"
	return nil
}
