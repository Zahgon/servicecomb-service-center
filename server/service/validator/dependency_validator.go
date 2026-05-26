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
	"github.com/go-chassis/cari/discovery"
)

var (
	addDependenciesReqValidator       validate.Validator
	overwriteDependenciesReqValidator validate.Validator
)

var versionAllowEmptyRegex, _ = regexp.Compile(`^(^\d+(\.\d+){0,2}\+?$|^\d+(\.\d+){0,2}-\d+(\.\d+){0,2}$|^latest$)?$`)

func defaultDependencyValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func AddDependenciesReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func CreateDependenciesReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func ValidateGetDependenciesRequest(v *discovery.GetDependenciesRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateCreateDependenciesRequest(v *discovery.CreateDependenciesRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateAddDependenciesRequest(v *discovery.AddDependenciesRequest) error {
	_ = "STUB: not implemented"
	return nil
}
