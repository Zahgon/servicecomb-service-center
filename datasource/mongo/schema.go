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

package mongo

import (
	"context"

	"github.com/apache/servicecomb-service-center/datasource/schema"
)

func init() {
	schema.Install("mongo", NewSchemaDAO)
}

func NewSchemaDAO(_ schema.Options) (schema.DAO, error) {
	_ = "STUB: not implemented"
	return *new(schema.DAO), nil
}

type SchemaDAO struct{}

func (s *SchemaDAO) GetRef(_ context.Context, _ *schema.RefRequest) (*schema.Ref, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SchemaDAO) ListRef(_ context.Context, _ *schema.RefRequest) ([]*schema.Ref, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SchemaDAO) DeleteRef(_ context.Context, _ *schema.RefRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SchemaDAO) GetContent(_ context.Context, _ *schema.ContentRequest) (*schema.Content, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SchemaDAO) PutContent(ctx context.Context, contentRequest *schema.PutContentRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SchemaDAO) PutManyContent(ctx context.Context, contentRequest *schema.PutManyContentRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SchemaDAO) DeleteContent(_ context.Context, _ *schema.ContentRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SchemaDAO) DeleteNoRefContents(_ context.Context) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
