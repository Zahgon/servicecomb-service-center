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

package mongo

import (
	"github.com/apache/servicecomb-service-center/datasource"
)

const defaultExpireTime = 300
const defaultPoolSize = 1000

func init() {
	datasource.Install("mongo", NewDataSource)
}

type DataSource struct {
	metadataManager datasource.MetadataManager
	sysManager      datasource.SystemManager
	depManager      datasource.DependencyManager
	scManager       datasource.SCManager
	metricsManager  datasource.MetricsManager
	syncManager     datasource.SyncManager
}

func (ds *DataSource) SystemManager() datasource.SystemManager {
	_ = "STUB: not implemented"
	return *new(datasource.SystemManager)
}

func (ds *DataSource) DependencyManager() datasource.DependencyManager {
	_ = "STUB: not implemented"
	return *new(datasource.DependencyManager)
}

func (ds *DataSource) MetadataManager() datasource.MetadataManager {
	_ = "STUB: not implemented"
	return *new(datasource.MetadataManager)
}

func (ds *DataSource) SCManager() datasource.SCManager {
	_ = "STUB: not implemented"
	return *new(datasource.SCManager)
}

func (ds *DataSource) MetricsManager() datasource.MetricsManager {
	_ = "STUB: not implemented"
	return *new(datasource.MetricsManager)
}

func (ds *DataSource) SyncManager() datasource.SyncManager {
	_ = "STUB: not implemented"
	return *new(datasource.SyncManager)
}

func NewDataSource(opts datasource.Options) (datasource.DataSource, error) {
	_ = "STUB: not implemented"
	// TODO: construct a reasonable DataSource instance
	return *new(datasource.DataSource), nil
}

// TODO: deal with exception

func (ds *DataSource) initialize() error {
	_ = "STUB: not implemented"

	// init heartbeat plugins
	return nil
}

// init mongo client

// create db index and validator

// if fast register enabled, init fast register service

// init cache

func (ds *DataSource) initPlugins() error { _ = "STUB: not implemented"; return nil }

func (ds *DataSource) initClient() error { _ = "STUB: not implemented"; return nil }

func (ds *DataSource) initStore() { _ = "STUB: not implemented"; return }

func initFastRegister() { _ = "STUB: not implemented"; return }
