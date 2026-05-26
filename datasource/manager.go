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

package datasource

type dataSourceEngine func(opts Options) (DataSource, error)

var (
	plugins        = make(map[string]dataSourceEngine)
	dataSourceInst DataSource
)

// load plugins configuration into plugins
func Install(pluginImplName string, engineFunc dataSourceEngine) { _ = "STUB: not implemented"; return }

// Init construct storage plugin instance
// invoked by sc main process
func Init(opts Options) error { _ = "STUB: not implemented"; return nil }

// init eventbase

func initDatasource(opts Options) error { _ = "STUB: not implemented"; return nil }

func GetSCManager() SCManager { _ = "STUB: not implemented"; return *new(SCManager) }

func GetMetadataManager() MetadataManager { _ = "STUB: not implemented"; return *new(MetadataManager) }

func GetSystemManager() SystemManager { _ = "STUB: not implemented"; return *new(SystemManager) }

func GetDependencyManager() DependencyManager {
	_ = "STUB: not implemented"
	return *new(DependencyManager)
}

func GetMetricsManager() MetricsManager { _ = "STUB: not implemented"; return *new(MetricsManager) }

func GetSyncManager() SyncManager { _ = "STUB: not implemented"; return *new(SyncManager) }
