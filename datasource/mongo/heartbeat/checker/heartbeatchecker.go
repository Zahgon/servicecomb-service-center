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

package checker

import (
	"context"

	pb "github.com/go-chassis/cari/discovery"

	"github.com/apache/servicecomb-service-center/datasource/mongo/heartbeat"
)

func init() {
	heartbeat.Install("checker", NewHeartBeatChecker)
}

type HeartBeatChecker struct {
}

func NewHeartBeatChecker() (heartbeat.HealthCheck, error) {
	_ = "STUB: not implemented"
	return *new(heartbeat.HealthCheck), nil
}

func (h *HeartBeatChecker) Heartbeat(ctx context.Context, request *pb.HeartbeatRequest) (*pb.HeartbeatResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HeartBeatChecker) CheckInstance(_ context.Context, _ *pb.MicroServiceInstance) error {
	_ = "STUB: not implemented"
	// do nothing
	return nil
}
