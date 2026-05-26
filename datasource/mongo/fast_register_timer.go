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
	"time"

	"github.com/go-chassis/foundation/gopool"
)

const (
	loopTime              = 100 * time.Millisecond
	batchLen              = 10000
	fuseMinCount          = 3
	fuseTime              = 5 * time.Second
	maxRegisterFailedTime = 500
	ctxCancelTimeOut      = 60 * time.Second
)

var fastRegisterTimeTask *FastRegisterTimeTask

type FastRegisterTimeTask struct {
	goroutine *gopool.Pool
}

func NewRegisterTimeTask() *FastRegisterTimeTask { _ = "STUB: not implemented"; return nil }

func (rt *FastRegisterTimeTask) Start() { _ = "STUB: not implemented"; return }

func (rt *FastRegisterTimeTask) Stop() { _ = "STUB: not implemented"; return }

func (rt *FastRegisterTimeTask) loopRegister(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// server shutdown

// fuse is triggered if failed registry counts more than fuseMinCount

// if instance batch register failed, register it single

func (rt *FastRegisterTimeTask) generateEvents(length int) []*InstanceRegisterEvent {
	_ = "STUB: not implemented"
	// if channel len >= batch len, use batch len, otherwise use channel len
	return nil
}

func (rt *FastRegisterTimeTask) RegisterInstancesAsync(events []*InstanceRegisterEvent, blockCh chan struct{}, failedCount chan int) {
	_ = "STUB: not implemented"
	return
}

//failed count

//add to failed instance channel, will retry register

func (rt *FastRegisterTimeTask) RegisterInstance(event *InstanceRegisterEvent, blockCh chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func endBlock(blockCh chan struct{}) { _ = "STUB: not implemented"; return }

func refreshCanceledCtx(event *InstanceRegisterEvent) context.CancelFunc {
	_ = "STUB: not implemented"
	return *new(context.CancelFunc)
}
