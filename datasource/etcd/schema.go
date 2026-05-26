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

	mapset "github.com/deckarep/golang-set"
	"github.com/little-cui/etcdadpt"
	"go.etcd.io/etcd/api/v3/mvccpb"

	"github.com/apache/servicecomb-service-center/datasource/schema"
)

func init() {
	schema.Install("etcd", NewSchemaDAO)
	schema.Install("embeded_etcd", NewSchemaDAO)
	schema.Install("embedded_etcd", NewSchemaDAO)
}

func NewSchemaDAO(_ schema.Options) (schema.DAO, error) {
	_ = "STUB: not implemented"
	return *new(schema.DAO), nil
}

type SchemaDAO struct{}

func (dao *SchemaDAO) GetRef(ctx context.Context, refRequest *schema.RefRequest) (*schema.Ref, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getSummary(ctx context.Context, serviceID string, schemaID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (dao *SchemaDAO) ListRef(ctx context.Context, refRequest *schema.RefRequest) ([]*schema.Ref, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// may be empty

func getSummaryMap(ctx context.Context, serviceID string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dao *SchemaDAO) DeleteRef(ctx context.Context, refRequest *schema.RefRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (dao *SchemaDAO) GetContent(ctx context.Context, contentRequest *schema.ContentRequest) (*schema.Content, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dao *SchemaDAO) PutContent(ctx context.Context, contentRequest *schema.PutContentRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// append the schemaID into service.Schemas if schemaID is new

func (dao *SchemaDAO) PutManyContent(ctx context.Context, contentRequest *schema.PutManyContentRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// unsafe!

// should update service.Schemas

// update service task

func transformSchemaIDsAndOptions(ctx context.Context, domainProject string, serviceID string,
	oldSchemaIDs []string, contentRequest *schema.PutManyContentRequest) ([]string, []etcdadpt.OpOptions) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dao *SchemaDAO) DeleteContent(ctx context.Context, contentRequest *schema.ContentRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO bad performance

func getContentHashMap(ctx context.Context) (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dao *SchemaDAO) DeleteNoRefContents(ctx context.Context) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func filterNoRefContentHashes(ctx context.Context, kvs []*mvccpb.KeyValue) (mapset.Set, error) {
	_ = "STUB: not implemented"
	return *new(mapset.Set), nil
}
