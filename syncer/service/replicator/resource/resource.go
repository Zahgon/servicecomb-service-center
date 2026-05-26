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

package resource

import (
	"context"

	"github.com/apache/servicecomb-service-center/eventbase/model"
	v1sync "github.com/apache/servicecomb-service-center/syncer/api/v1"

	"github.com/go-chassis/cari/sync"
)

const (
	Success int32 = iota
	Fail
	Skip
	MicroNonExist
	InstNonExist
	NonImplement
)

const (
	ResultStatusSkip          = "skip"
	ResultStatusSuccess       = "success"
	ResultStatusFail          = "fail"
	ResultStatusMicroNonExist = "microNonExist"
	ResultStatusInstNonExist  = "instNonExist"
	ResultStatusNonImplement  = "nonImplement"
)

var codeDescriber = map[int32]string{
	Skip:          ResultStatusSkip,
	Success:       ResultStatusSuccess,
	Fail:          ResultStatusFail,
	MicroNonExist: ResultStatusMicroNonExist,
	InstNonExist:  ResultStatusInstNonExist,
	NonImplement:  ResultStatusNonImplement,
}

type NewResource func(event *v1sync.Event) Resource

var (
	resources = map[string]NewResource{
		Account:      NewAccount,
		Role:         NewRole,
		Microservice: NewMicroservice,
		Instance:     NewInstance,
		Heartbeat:    NewHeartbeat,
		Config:       NewConfig,
		KV:           NewKV,
	}
)

func RegisterResources(name string, nr NewResource) { _ = "STUB: not implemented"; return }

func New(event *v1sync.Event) (Resource, *Result) {
	_ = "STUB: not implemented"
	return *new(Resource), nil
}

type operator struct {
	a ActionHandler
}

func newOperator(a ActionHandler) *operator { _ = "STUB: not implemented"; return nil }

func (o *operator) operate(ctx context.Context, action string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func NewResult(status int32, message string) *Result { _ = "STUB: not implemented"; return nil }

func FailResult(err error) *Result { _ = "STUB: not implemented"; return nil }

func SuccessResult() *Result { _ = "STUB: not implemented"; return nil }

func SkipResult() *Result { _ = "STUB: not implemented"; return nil }

type Result struct {
	EventID string
	Status  int32
	Message string
}

func (r *Result) WithMessage(m string) *Result { _ = "STUB: not implemented"; return nil }

func (r *Result) WithEventID(id string) *Result { _ = "STUB: not implemented"; return nil }

func (r *Result) Flag() string { _ = "STUB: not implemented"; return "" }

func NonImplementResult() *Result { _ = "STUB: not implemented"; return nil }

type FailHandler interface {
	FailHandle(context.Context, int32) (*v1sync.Event, error)
	CanDrop() bool
}

type defaultFailHandler struct {
}

func (d *defaultFailHandler) FailHandle(context.Context, int32) (*v1sync.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultFailHandler) CanDrop() bool { _ = "STUB: not implemented"; return false }

type OperateHandler interface {
	LoadCurrentResource(context.Context) *Result
	NeedOperate(context.Context) *Result
	Operate(context.Context) *Result
}

type Resource interface {
	OperateHandler
	FailHandler
}

type ActionHandler interface {
	CreateHandle(context.Context) error
	UpdateHandle(context.Context) error
	DeleteHandle(context.Context) error
}

func newInputParam(input interface{}, callback func()) *inputParam {
	_ = "STUB: not implemented"
	return nil
}

type inputParam struct {
	input interface{}

	callback func()
}

func newInputLoader(
	event *v1sync.Event,
	create *inputParam,
	update *inputParam,
	delete *inputParam) *inputLoader {
	_ = "STUB: not implemented"
	return nil
}

type inputLoader struct {
	event *v1sync.Event

	createInput    interface{}
	createCallback func()

	updateInput    interface{}
	updateCallback func()

	deleteInput    interface{}
	deleteCallback func()
}

func (i *inputLoader) loadInputUtil(value interface{}, callback func()) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *inputLoader) loadCreateInput() error { _ = "STUB: not implemented"; return nil }

func (i *inputLoader) loadUpdateInput() error { _ = "STUB: not implemented"; return nil }

func (i *inputLoader) loadDeleteInput() error { _ = "STUB: not implemented"; return nil }

func (i *inputLoader) loadInput() error { _ = "STUB: not implemented"; return nil }

type tombstoneLoader interface {
	get(ctx context.Context, req *model.GetTombstoneRequest) (*sync.Tombstone, error)
}

type checker struct {
	curNotNil bool

	event      *v1sync.Event
	updateTime func() (int64, error)
	resourceID string

	tombstoneLoader tombstoneLoader
}

func formatUpdateTimeSecond(src string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func secToNanoSec(timestamp int64) int64 { _ = "STUB: not implemented"; return 0 }

func (o *checker) needOperate(ctx context.Context) *Result { _ = "STUB: not implemented"; return nil }

func (o *checker) get(ctx context.Context, req *model.GetTombstoneRequest) (*sync.Tombstone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
