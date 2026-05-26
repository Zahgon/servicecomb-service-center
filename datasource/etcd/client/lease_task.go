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

package client

import (
	"context"
	"time"

	simple "github.com/apache/servicecomb-service-center/pkg/time"
	"github.com/little-cui/etcdadpt"
)

const leaseProfTimeFmt = "15:04:05.000"

type LeaseTask struct {
	Client etcdadpt.Client

	key     string
	LeaseID int64
	TTL     int64

	recvTime simple.Time
	err      error
}

func (lat *LeaseTask) Key() string { _ = "STUB: not implemented"; return "" }

func (lat *LeaseTask) Do(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }

// it means instance is deleted

// DON'T care about other errors, so client should send heartbeat in next interval

func (lat *LeaseTask) Err() error { _ = "STUB: not implemented"; return nil }

func (lat *LeaseTask) ReceiveTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func NewLeaseAsyncTask(op etcdadpt.OpOptions) *LeaseTask { _ = "STUB: not implemented"; return nil }

func ToLeaseAsyncTaskKey(key string) string { _ = "STUB: not implemented"; return "" }
