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

	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
)

func GetAllDomainRawData(ctx context.Context) ([]*kvstore.KeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetAllDomain(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AddDomain(ctx context.Context, domain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func DomainExist(ctx context.Context, domain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func AddProject(ctx context.Context, domain, project string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func ProjectExist(ctx context.Context, domain, project string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func NewDomainProject(ctx context.Context, domain, project string) error {
	_ = "STUB: not implemented"
	return nil
}
