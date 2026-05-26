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

package etcd

import (
	"github.com/apache/servicecomb-service-center/datasource"
	"github.com/apache/servicecomb-service-center/datasource/etcd/sd"
)

const compactLockKey = "/etcd-compact"

var clustersIndex = make(map[string]int)

func init() {
	datasource.Install("etcd", NewDataSource)
	datasource.Install("embeded_etcd", NewDataSource) //TODO remove misspell in future
	datasource.Install("embedded_etcd", NewDataSource)

	sd.RegisterInnerTypes()
}

type DataSource struct {
	Options *datasource.Options

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
	return *new(datasource.DataSource), nil
}

func (ds *DataSource) initialize() error {
	_ = "STUB: not implemented"
	// Wait for kv store ready
	return nil
}

// Compact

func (ds *DataSource) initClustersIndex() { _ = "STUB: not implemented"; return }

func (ds *DataSource) initPlugins() {
	_ = "STUB: not implemented"
	// registry
	return
}

// clusters

// discovery

func (ds *DataSource) initKvStore() {
	_ = "STUB: not implemented"
	// init client/sd plugins
	return
}

// Add events handlers

func (ds *DataSource) autoCompact() { _ = "STUB: not implemented"; return }
