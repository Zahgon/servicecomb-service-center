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

// Package grc include API of governance(grc is the abbreviation of governance)
package grc

import (
	"k8s.io/kube-openapi/pkg/validation/spec"
)

type ValueType string

var PolicyNames []string

// policySchemas saves policy kind and schema
var policySchemas = make(map[string]*spec.Schema)

// RegisterPolicySchema register a contract of one kind of policy
// this API is not thread safe, only use it during sc init
func RegisterPolicySchema(kind string, schema *spec.Schema) { _ = "STUB: not implemented"; return }

// ValidatePolicySpec validates spec attributes
func ValidatePolicySpec(kind string, spec interface{}) error { _ = "STUB: not implemented"; return nil }
