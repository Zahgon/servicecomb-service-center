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
	"net/http"

	"github.com/apache/servicecomb-service-center/pkg/lb"
	"github.com/apache/servicecomb-service-center/pkg/rest"
)

func NewLBClient(endpoints []string, options rest.URLClientOption) (*LBClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type LBClient struct {
	*rest.URLClient
	Retries int
	LB      lb.LoadBalancer
}

func (c *LBClient) Next() string { _ = "STUB: not implemented"; return "" }

func (c *LBClient) RestDoWithContext(ctx context.Context, method string, api string, headers http.Header, body []byte) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
