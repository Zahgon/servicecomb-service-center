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

	pb "github.com/go-chassis/cari/discovery"
)

// ExistSchema only return the summary without content if schema exist
func ExistSchema(ctx context.Context, request *pb.GetSchemaRequest) (*pb.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func existOldSchema(ctx context.Context, request *pb.GetSchemaRequest) (*pb.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetSchema(ctx context.Context, request *pb.GetSchemaRequest) (*pb.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getOldSchema(ctx context.Context, request *pb.GetSchemaRequest) (*pb.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListSchema(ctx context.Context, request *pb.GetAllSchemaRequest) ([]*pb.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getSchema(ctx context.Context, req *pb.GetSchemaRequest, withSchema bool) (*pb.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getOldSchemaIDs(ctx context.Context, serviceID string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeRequests(ctx context.Context, serviceID string, oldSchemaIDs []string) ([]*pb.GetSchemaRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeleteSchema(ctx context.Context, request *pb.DeleteSchemaRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func deleteOldSchema(ctx context.Context, request *pb.DeleteSchemaRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// PutSchemas covers all the schemas of a service.
// To cover the old schemas, ModifySchemas adds new schemas into, delete and
// modify the old schemas.
func PutSchemas(ctx context.Context, request *pb.ModifySchemasRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// no need to check quota usage because overwrite existing.

// PutSchema modifies a specific schema.
func PutSchema(ctx context.Context, request *pb.ModifySchemaRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func checkSchemaQuota(ctx context.Context, serviceID string, schemaID string) error {
	_ = "STUB: not implemented"
	return nil
}

func Usage(ctx context.Context, serviceID string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
