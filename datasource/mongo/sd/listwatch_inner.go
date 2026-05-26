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

package sd

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/apache/servicecomb-service-center/datasource/sdcommon"
)

type parsefunc func(doc bson.Raw) (resource sdcommon.Resource)

type mongoListWatch struct {
	Key         string
	resumeToken bson.Raw
	parseFunc   parsefunc
}

func (lw *mongoListWatch) List(op sdcommon.ListWatchConfig) (*sdcommon.ListWatchResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert mongoListResponse to ListWatchResp

func (lw *mongoListWatch) EventBus(op sdcommon.ListWatchConfig) *sdcommon.EventBus {
	_ = "STUB: not implemented"
	return nil
}

func (lw *mongoListWatch) DoWatch(ctx context.Context, f func(*sdcommon.ListWatchResp)) error {
	_ = "STUB: not implemented"
	return nil
}

// ignore instance refresh_time change event for avoid meaningless instance push.

// convert mongoWatchResponse to ListWatchResp

func (lw *mongoListWatch) ResumeToken() bson.Raw { _ = "STUB: not implemented"; return *new(bson.Raw) }

func (lw *mongoListWatch) setResumeToken(resumeToken bson.Raw) { _ = "STUB: not implemented"; return }

func (lw *mongoListWatch) doParseWatchRspToResource(wRsp *MongoWatchResponse) (resource sdcommon.Resource) {
	_ = "STUB: not implemented"
	return *new(sdcommon.Resource)
}

//delete operation has no fullDocumentValue
