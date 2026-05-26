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

package chain

import (
	"context"

	"github.com/apache/servicecomb-service-center/pkg/util"
)

type InvocationOption func(op InvocationOp) InvocationOp

type InvocationOp struct {
	Func  CallbackFunc
	Async bool
}

// WithFunc append a func to the begin of invocation callback list
func WithFunc(f func(r Result)) InvocationOption {
	_ = "STUB: not implemented"
	return *new(InvocationOption)
}

// WithAsyncFunc called concurrently after all WithFunc finish
func WithAsyncFunc(f func(r Result)) InvocationOption {
	_ = "STUB: not implemented"
	return *new(InvocationOption)
}

type Invocation struct {
	Callback
	context *util.StringContext
	chain   Chain
}

func (i *Invocation) Init(ctx context.Context, ch Chain) { _ = "STUB: not implemented"; return }

func (i *Invocation) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (i *Invocation) WithContext(key util.CtxKey, val interface{}) *Invocation {
	_ = "STUB: not implemented"
	return nil
}

// Next is the method to go next step in handler chain
// WithFunc and WithAsyncFunc options can add customize callbacks in chain
// and the callbacks seq like below:
// when i.Next(WithFunc(CB1)).Next(WithAsyncFunc(CB2)).Next(WithFunc(CB3)).Invoke(CB0)
// then i.Success/Fail() -> CB3 -> CB1 -> CB0(invoke)             goroutine 0
//
//	\-> CB2(async)    goroutine 1
func (i *Invocation) Next(opts ...InvocationOption) { _ = "STUB: not implemented"; return }

func (i *Invocation) setCallback(f CallbackFunc, async bool) { _ = "STUB: not implemented"; return }

func callback(prev, next CallbackFunc, async bool, r Result) {
	_ = "STUB: not implemented"
	// we make sure the all sync funcs called before the async funcs
	return
}

func (i *Invocation) Invoke(last CallbackFunc) { _ = "STUB: not implemented"; return }

// this recover only catch the exceptions raised in sync invocations.
// The async invocations will be catch by gopool pkg then it never
// change the callback results.
// i.Fail(discovery.NewError(discovery.ErrInternal, fmt.Sprintf("%v", itf)))

func NewInvocation(ctx context.Context, ch Chain) (inv Invocation) {
	_ = "STUB: not implemented"
	return *new(Invocation)
}
