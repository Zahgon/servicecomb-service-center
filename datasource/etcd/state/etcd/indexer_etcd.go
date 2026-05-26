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

	"github.com/apache/servicecomb-service-center/datasource/etcd/state/kvstore"
	"github.com/apache/servicecomb-service-center/datasource/etcd/state/parser"
	"github.com/little-cui/etcdadpt"
)

// Indexer implements kvstore.Indexer.
// Indexer searches data from etcd server.
type Indexer struct {
	Client etcdadpt.Client
	Parser parser.Parser
	Root   string
}

func (i *Indexer) CheckPrefix(key string) error { _ = "STUB: not implemented"; return nil }

func (i *Indexer) Search(ctx context.Context, opts ...etcdadpt.OpOption) (*kvstore.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Creditable implements kvstore.Indexer#Creditable.
func (i *Indexer) Creditable() bool { _ = "STUB: not implemented"; return false }

func NewEtcdIndexer(root string, p parser.Parser) (indexer *Indexer) {
	_ = "STUB: not implemented"
	return nil
}
