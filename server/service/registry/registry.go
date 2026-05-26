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

package registry

import (
	"context"
)

func addDefaultContextValue(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// set default domain/project

// register without quota check

// is a sync operation

func SelfRegister(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// start send heart beat job

func selfRegister(pCtx context.Context) error { _ = "STUB: not implemented"; return nil }

// 实例信息

func registerService(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func registerNewService(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func registerInstance(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func selfHeartBeat(pCtx context.Context) error { _ = "STUB: not implemented"; return nil }

func autoSelfHeartBeat() { _ = "STUB: not implemented"; return }

//服务不存在，创建服务

func SelfUnregister(pCtx context.Context) error { _ = "STUB: not implemented"; return nil }
