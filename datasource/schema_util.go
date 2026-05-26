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

package datasource

import (
	pb "github.com/go-chassis/cari/discovery"
)

// @titile SchemasAnalysis
// @description schemasanlysis decide the schema to be deleted,updated or added
// @param schemas []*pb.Schema "the schemas from request"
// @param schemasFromDb []*pb.Schema "the schemas from database by the serviceID"
// @param schemaIDsInService []string "the schema field in service struct"
func SchemasAnalysis(schemas []*pb.Schema, schemasFromDb []*pb.Schema, schemaIDsInService []string) (
	[]*pb.Schema, []*pb.Schema, []*pb.Schema, []string) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}
