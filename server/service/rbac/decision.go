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

package rbac

import (
	"context"

	"github.com/apache/servicecomb-service-center/server/plugin/auth"
	rbacmodel "github.com/go-chassis/cari/rbac"
)

// Allow return: matched labels(empty if no label defined), error
func Allow(ctx context.Context, _ string, roleList []string,
	targetResource *auth.ResourceScope) ([]map[string]string, error) {
	_ = "STUB: not implemented"
	//TODO check project
	return nil, nil
}

// allow, but no label found, means we can ignore the labels

// target resource needs no label, return without filter

// allow, and labels found, filter the labels

// target resource label matches no label in permission, means not allow

func FilterLabel(targetResourceLabel []map[string]string, permLabelList []map[string]string) []map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func LabelMatched(targetResourceLabel map[string]string, permLabel map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

func getPermsByRoles(ctx context.Context, roleList []string) ([]*rbacmodel.Permission, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetLabel checks if the perms have permission to operate the resource(ignore label),
// if one perm have the permission, add it's label to the result.
func GetLabel(perms []*rbacmodel.Permission, targetResource, verb string) (allow bool, labelList []map[string]string) {
	_ = "STUB: not implemented"
	return false, nil
}

// allow and has no label, return fast

// GetLabel checks if the perm have permission to operate the resource(ignore label),
// if the perm have the permission, return it's label.
func GetLabelFromSinglePerm(perm *rbacmodel.Permission, targetResource, verb string) (allow bool, labelList []map[string]string) {
	_ = "STUB: not implemented"
	return false, nil
}

func allowVerb(haystack []string, needle string) bool { _ = "STUB: not implemented"; return false }

func getResourceLabel(resources []*rbacmodel.Resource, needle string) (allow bool, labelList []map[string]string) {
	_ = "STUB: not implemented"
	return false, nil
}

// filter the same resource

// has no label, return fast
