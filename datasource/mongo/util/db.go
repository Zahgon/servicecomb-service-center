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

package util

import (
	"context"

	"github.com/go-chassis/cari/discovery"
	"github.com/go-chassis/cari/rbac"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Option func(filter bson.M)

func Domain(domain string) Option { _ = "STUB: not implemented"; return *new(Option) }

func Project(project string) Option { _ = "STUB: not implemented"; return *new(Option) }

func AccountName(name interface{}) Option { _ = "STUB: not implemented"; return *new(Option) }

func Password(password string) Option { _ = "STUB: not implemented"; return *new(Option) }

func Roles(roles []string) Option { _ = "STUB: not implemented"; return *new(Option) }

func TokenExpirationTime(tokenExpirationTime string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func CurrentPassword(password string) Option { _ = "STUB: not implemented"; return *new(Option) }

func Status(status string) Option { _ = "STUB: not implemented"; return *new(Option) }

func ID(id string) Option { _ = "STUB: not implemented"; return *new(Option) }

func RoleName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

func Perms(perms []*rbac.Permission) Option { _ = "STUB: not implemented"; return *new(Option) }

func AccountUpdateTime(dt interface{}) Option { _ = "STUB: not implemented"; return *new(Option) }

func RoleUpdateTime(dt interface{}) Option { _ = "STUB: not implemented"; return *new(Option) }

func AccountLockKey(key interface{}) Option { _ = "STUB: not implemented"; return *new(Option) }

func AccountLockStatus(status interface{}) Option { _ = "STUB: not implemented"; return *new(Option) }

func AccountLockReleaseAt(releaseAt interface{}) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func In(data interface{}) Option { _ = "STUB: not implemented"; return *new(Option) }

func NotIn(data interface{}) Option { _ = "STUB: not implemented"; return *new(Option) }

func Set(data interface{}) Option { _ = "STUB: not implemented"; return *new(Option) }

func Nor(options ...Option) Option { _ = "STUB: not implemented"; return *new(Option) }

func Or(options ...Option) Option { _ = "STUB: not implemented"; return *new(Option) }

func NewFilter(options ...Option) bson.M { _ = "STUB: not implemented"; return *new(bson.M) }

func NewDomainProjectFilter(domain string, project string, options ...func(filter bson.M)) bson.M {
	_ = "STUB: not implemented"
	return *new(bson.M)
}

func NewBasicFilter(ctx context.Context, options ...func(filter bson.M)) bson.M {
	_ = "STUB: not implemented"
	return *new(bson.M)
}

func InstanceServiceID(serviceID interface{}) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func InstanceInstanceID(instanceID string) Option { _ = "STUB: not implemented"; return *new(Option) }

// ServiceServiceID serviceID can be string or bson.M
func ServiceServiceID(serviceID interface{}) Option { _ = "STUB: not implemented"; return *new(Option) }

func ServiceEnv(env string) Option { _ = "STUB: not implemented"; return *new(Option) }

func ServiceAppID(appID string) Option { _ = "STUB: not implemented"; return *new(Option) }

func ServiceModTime(modTime string) Option { _ = "STUB: not implemented"; return *new(Option) }

func ServiceProperty(property map[string]string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func ServiceServiceName(serviceName interface{}) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func ServiceID(serviceID string) Option { _ = "STUB: not implemented"; return *new(Option) }

func ServiceAlias(alias string) Option { _ = "STUB: not implemented"; return *new(Option) }

func ServiceSchemas(schemas []string) Option { _ = "STUB: not implemented"; return *new(Option) }

func ServiceVersion(version interface{}) Option { _ = "STUB: not implemented"; return *new(Option) }

func ServiceType(serviceType string) Option { _ = "STUB: not implemented"; return *new(Option) }

func ServiceKeyTenant(tenant string) Option { _ = "STUB: not implemented"; return *new(Option) }

func ServiceKeyAppID(appID string) Option { _ = "STUB: not implemented"; return *new(Option) }

func ServiceKeyServiceName(serviceName string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func ServiceKeyServiceEnv(env string) Option { _ = "STUB: not implemented"; return *new(Option) }

func ServiceKeyServiceVersion(version string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func Schema(schema string) Option { _ = "STUB: not implemented"; return *new(Option) }

func SchemaID(schemaID string) Option { _ = "STUB: not implemented"; return *new(Option) }

func SchemaSummary(schemaSummary string) Option { _ = "STUB: not implemented"; return *new(Option) }

func Tags(tags map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

func Instance(instance *discovery.MicroServiceInstance) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func InstanceModTime(modTime string) Option { _ = "STUB: not implemented"; return *new(Option) }

func InstanceStatus(status string) Option { _ = "STUB: not implemented"; return *new(Option) }

func InstanceProperties(properties map[string]string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func BuildIndexDoc(keys ...string) mongo.IndexModel {
	_ = "STUB: not implemented"
	return *new(mongo.IndexModel)
}

func NotGlobal() Option { _ = "STUB: not implemented"; return *new(Option) }

func Global() Option { _ = "STUB: not implemented"; return *new(Option) }
