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
	microServiceKeyValidator       validate.Validator
	microServiceKeySearchValidator validate.Validator
	existenceReqValidator          validate.Validator
	getServiceReqValidator         validate.Validator
	createServiceReqValidator      validate.Validator
	updateServicePropsReqValidator validate.Validator
	unregisterManyServiceValidator validate.Validator
)

var (
	// 非map/slice的validator
	nameRegex, _ = regexp.Compile(`^[a-zA-Z0-9]*$|^[a-zA-Z0-9][a-zA-Z0-9_\-.]*[a-zA-Z0-9]$`)
	// find 支持alias，多个:
	serviceNameForFindRegex, _ = regexp.Compile(`^[a-zA-Z0-9]*$|^[a-zA-Z0-9][a-zA-Z0-9_\-.:]*[a-zA-Z0-9]$`)
	// version規則: x[.y[.z]]
	versionRegex           = validate.NewVersionRegexp(false)
	pathRegex, _           = regexp.Compile(`^[A-Za-z0-9.,?'\\/+&amp;%$#=~_\-@{}]*$`)
	levelRegex, _          = regexp.Compile(`^(FRONT|MIDDLE|BACK)$`)
	statusRegex, _         = regexp.Compile("^(" + discovery.MS_UP + "|" + discovery.MS_DOWN + ")?$")
	serviceIDRegex, _      = regexp.Compile(`^\S*$`)
	serviceIDRangeRegex, _ = regexp.Compile(`^\S{1,64}$`)
	aliasRegex, _          = regexp.Compile(`^[a-zA-Z0-9_\-.:]*$`)
	registerByRegex, _     = regexp.Compile("^(" + util.StringJoin([]string{discovery.REGISTERBY_SDK, discovery.REGISTERBY_SIDECAR, discovery.REGISTERBY_PLATFORM}, "|") + ")*$")
	envRegex, _            = regexp.Compile("^(" + util.StringJoin([]string{
		discovery.ENV_DEV, discovery.ENV_TEST, discovery.ENV_ACCEPT, discovery.ENV_PROD}, "|") + ")*$")
	schemaIDRegex, _      = regexp.Compile(`^[a-zA-Z0-9]{1,160}$|^[a-zA-Z0-9][a-zA-Z0-9_\-.]{0,158}[a-zA-Z0-9]$`)
	accountStatusRegex, _ = regexp.Compile(`^(active|inactive)$|^$`)
)

func MicroServiceKeyValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func MicroServiceSearchKeyValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

// support name or alias

func GetServiceReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func CreateServiceReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func UpdateServicePropsReqValidator() *validate.Validator { _ = "STUB: not implemented"; return nil }

func ValidateCreateServiceRequest(v *discovery.CreateServiceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateUnregisterManyService(in *discovery.DelServicesRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateGetServiceExistenceRequest(in *discovery.GetExistenceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateUpdateServicePropsRequest(request *discovery.UpdateServicePropsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateDeleteServiceRequest(request *discovery.DeleteServiceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateGetServiceRequest(request *discovery.GetServiceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateGetAppsRequest(v *discovery.GetAppsRequest) error {
	_ = "STUB: not implemented"
	return nil
}
