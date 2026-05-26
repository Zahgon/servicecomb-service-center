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

package validator

import (
	"regexp"

	"github.com/apache/servicecomb-service-center/pkg/validate"
	pb "github.com/go-chassis/cari/discovery"
)

var (
	getSchemaReqValidator     validate.Validator
	modifySchemasReqValidator validate.Validator
	modifySchemaReqValidator  validate.Validator
)

var (
	schemaIDUnlimitedRegex, _ = regexp.Compile(`^[a-zA-Z0-9]+$|^[a-zA-Z0-9][a-zA-Z0-9_\-.]*[a-zA-Z0-9]$`)
	schemaSummaryRegex, _     = regexp.Compile(`^[a-zA-Z0-9]*$`)
)

func GetSchemaReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func ModifySchemasReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func ModifySchemaReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

// forward compatibility: allow empty

func ValidateGetSchema(request *pb.GetSchemaRequest) error { _ = "STUB: not implemented"; return nil }

func ValidateListSchema(request *pb.GetAllSchemaRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidatePutSchema(request *pb.ModifySchemaRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidatePutSchemas(request *pb.ModifySchemasRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateDeleteSchema(request *pb.DeleteSchemaRequest) error {
	_ = "STUB: not implemented"
	return nil
}
