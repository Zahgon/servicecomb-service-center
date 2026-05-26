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

package disco

import (
	"context"
	"net/http"

	"github.com/apache/servicecomb-service-center/pkg/rest"
	pb "github.com/go-chassis/cari/discovery"
)

var trueOrFalse = map[string]bool{"true": true, "false": false, "1": true, "0": false}

type ServiceResource struct {
	//
}

func (s *ServiceResource) URLPatterns() []rest.Route { _ = "STUB: not implemented"; return nil }

// tags

func (s *ServiceResource) RegisterService(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *ServiceResource) PutServiceProperties(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *ServiceResource) UnregisterService(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *ServiceResource) ListService(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *ServiceResource) ResourceExist(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func resourceExist(ctx context.Context, w http.ResponseWriter, request *pb.GetExistenceRequest) (*pb.GetExistenceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ServiceResource) GetService(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *ServiceResource) UnregisterManyService(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *ServiceResource) PutManyTags(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *ServiceResource) PutTag(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *ServiceResource) ListTag(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *ServiceResource) DeleteManyTags(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
