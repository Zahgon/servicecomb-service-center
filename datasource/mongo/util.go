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

	"github.com/apache/servicecomb-service-center/datasource"
	pb "github.com/go-chassis/cari/discovery"
)

type InstanceSlice []*pb.MicroServiceInstance

func (s InstanceSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (s InstanceSlice) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s InstanceSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func statistics(ctx context.Context, withShared bool) (*pb.Statistics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getInstanceCountByDomain(ctx context.Context, svcIDToNonVerKey map[string]string, resp chan datasource.GetInstanceCountByDomainResponse) {
	_ = "STUB: not implemented"
	return
}
