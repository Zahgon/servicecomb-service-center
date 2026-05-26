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

package replicator

import (
	"context"

	"google.golang.org/grpc"

	v1sync "github.com/apache/servicecomb-service-center/syncer/api/v1"
	"github.com/apache/servicecomb-service-center/syncer/service/replicator/resource"
)

const (
	schema      = "grpc"
	serviceName = "syncer"
)

const (
	reservedSize = 512 * 1024
	maxSize      = 10*1024*1024 - reservedSize
)

var (
	manager = NewManager(make(map[string]struct{}, 1000))
)

var (
	conn      *grpc.ClientConn
	peerToken = ""
)

func Work() error { _ = "STUB: not implemented"; return nil }

func InitSyncClient() error { _ = "STUB: not implemented"; return nil }

func Close() { _ = "STUB: not implemented"; return }

func Manager() Replicator {
	_ = "STUB: not implemented"

	// Replicator define replicator manager, receive events from event manager
	// and send events to remote syncer
	return *new(Replicator)
}

type Replicator interface {
	Replicate(ctx context.Context, el *v1sync.EventList) (*v1sync.Results, error)
	Persist(ctx context.Context, el *v1sync.EventList) []*resource.Result
}

func NewManager(cache map[string]struct{}) Replicator {
	_ = "STUB: not implemented"
	return *new(Replicator)
}

type replicatorManager struct {
	cache map[string]struct{}
}

func (r *replicatorManager) Replicate(ctx context.Context, el *v1sync.EventList) (*v1sync.Results, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pageEvents(source *v1sync.EventList, max int) []*v1sync.EventList {
	_ = "STUB: not implemented"
	return nil
}

func (r *replicatorManager) replicate(ctx context.Context, el *v1sync.EventList) (*v1sync.Results, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *replicatorManager) Persist(ctx context.Context, el *v1sync.EventList) []*resource.Result {
	_ = "STUB: not implemented"
	return nil
}
