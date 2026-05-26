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

package quota

import (
	"context"

	"github.com/apache/servicecomb-service-center/pkg/plugin"
)

const QUOTA plugin.Kind = "quota"

type Manager interface {
	RemandQuotas(ctx context.Context, t ResourceType)
	GetQuota(ctx context.Context, t ResourceType) int64
	Usage(ctx context.Context, req *Request) (int64, error)
}

func GetQuota(ctx context.Context, resourceType ResourceType) int64 {
	_ = "STUB: not implemented"
	return 0
}

func Remand(ctx context.Context, resourceType ResourceType) { _ = "STUB: not implemented"; return }

func Usage(ctx context.Context, req *Request) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Apply 申请配额sourceType serviceinstance servicetype
func Apply(ctx context.Context, res *Request) error { _ = "STUB: not implemented"; return nil }
