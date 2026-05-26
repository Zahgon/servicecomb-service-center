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

	"github.com/apache/servicecomb-service-center/pkg/util"
	"github.com/apache/servicecomb-service-center/pkg/validate"
	"github.com/go-chassis/cari/discovery"
)

var (
	findInstanceReqValidator        validate.Validator
	batchFindInstanceReqValidator   validate.Validator
	getInstanceReqValidator         validate.Validator
	updateInstanceReqValidator      validate.Validator
	registerInstanceReqValidator    validate.Validator
	heartbeatReqValidator           validate.Validator
	updateInstancePropsReqValidator validate.Validator
)

var (
	instStatusRegex, _ = regexp.Compile("^(" + util.StringJoin([]string{
		discovery.MSI_UP, discovery.MSI_DOWN, discovery.MSI_STARTING, discovery.MSI_TESTING, discovery.MSI_OUTOFSERVICE}, "|") + ")?$")
	updateInstStatusRegex, _ = regexp.Compile("^(" + util.StringJoin([]string{
		discovery.MSI_UP, discovery.MSI_DOWN, discovery.MSI_STARTING, discovery.MSI_TESTING, discovery.MSI_OUTOFSERVICE}, "|") + ")$")
	hbModeRegex, _               = regexp.Compile(`^(push|pull)$`)
	urlRegex, _                  = regexp.Compile(`^\S*$`)
	epRegex, _                   = regexp.Compile(`\S+`)
	simpleNameAllowEmptyRegex, _ = regexp.Compile(`^[A-Za-z0-9_.-]*$`)
	simpleNameRegex, _           = regexp.Compile(`^[A-Za-z0-9_.-]+$`)
	regionRegex, _               = regexp.Compile(`^[A-Za-z0-9_.-]+$`)
)

func FindInstanceReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func FindManyInstanceReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func GetInstanceReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func HeartbeatReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func UpdateInstanceStatusReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func UpdateInstancePropsReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func RegisterInstanceReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

// allow empty endpoint register for client only

func ValidateRegisterInstanceRequest(in *discovery.RegisterInstanceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateUnregisterInstanceRequest(in *discovery.UnregisterInstanceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateHeartbeatRequest(in *discovery.HeartbeatRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateGetOneInstanceRequest(in *discovery.GetOneInstanceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateGetInstancesRequest(in *discovery.GetInstancesRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateFindInstancesRequest(in *discovery.FindInstancesRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateFindManyInstancesRequest(in *discovery.BatchFindInstancesRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateUpdateInstanceStatusRequest(in *discovery.UpdateInstanceStatusRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateUpdateInstancePropsRequest(in *discovery.UpdateInstancePropsRequest) error {
	_ = "STUB: not implemented"
	return nil
}
