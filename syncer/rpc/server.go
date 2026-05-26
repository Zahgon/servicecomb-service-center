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

package rpc

import (
	"context"

	v1sync "github.com/apache/servicecomb-service-center/syncer/api/v1"
	"github.com/apache/servicecomb-service-center/syncer/service/replicator"
	"github.com/apache/servicecomb-service-center/syncer/service/replicator/resource"
)

const (
	HealthStatusConnected = "CONNECTED"
	HealthStatusAbnormal  = "ABNORMAL"
	HealthStatusClose     = "CLOSE"
	HealthStatusAuthFail  = "AuthFail"

	RbacAllowedAccountName = "sync-user"
	RbacAllowedRoleName    = "sync-admin"
)

func NewServer() *Server { _ = "STUB: not implemented"; return nil }

type Server struct {
	v1sync.UnimplementedEventServiceServer

	replicator replicator.Replicator
}

func (s *Server) Sync(ctx context.Context, events *v1sync.EventList) (*v1sync.Results, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateFailedResults(events *v1sync.EventList, err error) (*v1sync.Results, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) toResults(results []*resource.Result) *v1sync.Results {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) Health(ctx context.Context, _ *v1sync.HealthRequest) (*v1sync.HealthReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO enable to close syncer
