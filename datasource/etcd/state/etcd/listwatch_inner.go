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

package etcd

import (
	"context"

	"github.com/apache/servicecomb-service-center/datasource/sdcommon"
	"github.com/little-cui/etcdadpt"
)

type innerListWatch struct {
	Client etcdadpt.Client
	Prefix string

	rev int64
}

func (lw *innerListWatch) List(op sdcommon.ListWatchConfig) (*sdcommon.ListWatchResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lw *innerListWatch) Revision() int64 { _ = "STUB: not implemented"; return 0 }

func (lw *innerListWatch) setRevision(rev int64) { _ = "STUB: not implemented"; return }

func (lw *innerListWatch) EventBus(op sdcommon.ListWatchConfig) *sdcommon.EventBus {
	_ = "STUB: not implemented"
	return nil
}

func (lw *innerListWatch) DoWatch(ctx context.Context, f func(*sdcommon.ListWatchResp)) error {
	_ = "STUB: not implemented"
	return nil
}

// compact可能会导致watch失败 or message body size lager than 4MB

func (lw *innerListWatch) doParsePluginRspToLwRsp(pluginRsp *etcdadpt.Response) *sdcommon.ListWatchResp {
	_ = "STUB: not implemented"
	return nil
}
