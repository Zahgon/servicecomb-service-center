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

package diagnose

import (
	"github.com/apache/servicecomb-service-center/datasource/etcd/state/parser"
	"github.com/apache/servicecomb-service-center/pkg/dump"
	"go.etcd.io/etcd/api/v3/mvccpb"
)

type CompareHolder interface {
	Compare() *CompareResult
}

type DataStore struct {
	Data       []*mvccpb.KeyValue
	DataParser parser.Parser
}

func (d *DataStore) ForEach(f func(i int, v *dump.KV) bool) { _ = "STUB: not implemented"; return }

type CompareResult struct {
	Name    string
	Results map[int][]string
}

type abstractCompareHolder struct {
	Cache        dump.Getter
	DataStore    *DataStore
	MismatchFunc func(v *dump.KV) string
}

func (h *abstractCompareHolder) toMap(getter dump.Getter) map[string]*dump.KV {
	_ = "STUB: not implemented"
	return nil
}

func (h *abstractCompareHolder) Compare() *CompareResult { _ = "STUB: not implemented"; return nil }

// add or update

// delete

type ServiceCompareHolder struct {
	*abstractCompareHolder
	Cache dump.MicroserviceSlice
	Kvs   []*mvccpb.KeyValue
}

func (h *ServiceCompareHolder) Compare() *CompareResult { _ = "STUB: not implemented"; return nil }

func (h *ServiceCompareHolder) toName(kv *dump.KV) string { _ = "STUB: not implemented"; return "" }

type InstanceCompareHolder struct {
	*abstractCompareHolder
	Cache dump.InstanceSlice
	Kvs   []*mvccpb.KeyValue
}

func (h *InstanceCompareHolder) Compare() *CompareResult { _ = "STUB: not implemented"; return nil }

func (h *InstanceCompareHolder) toName(kv *dump.KV) string { _ = "STUB: not implemented"; return "" }
