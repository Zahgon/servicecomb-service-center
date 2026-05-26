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
	getTagsReqValidator   validate.Validator
	addTagsReqValidator   validate.Validator
	updateTagReqValidator validate.Validator
	deleteTagReqValidator validate.Validator
)

var (
	tagRegex, _ = regexp.Compile(`^[a-zA-Z][a-zA-Z0-9_\-.]{0,63}$`)
)

func GetTagsReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func AddTagsReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func UpdateTagReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func DeleteTagReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func ValidateAddServiceTagsRequest(v *discovery.AddServiceTagsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateUpdateServiceTagRequest(v *discovery.UpdateServiceTagRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateDeleteServiceTagsRequest(v *discovery.DeleteServiceTagsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateGetServiceTagsRequest(v *discovery.GetServiceTagsRequest) error {
	_ = "STUB: not implemented"
	return nil
}
