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

import (
	"github.com/go-chassis/openlog"
)

var (
	dataSourceInst DataSource
	plugins        = make(map[string]dataSourceEngine)

	logger openlog.Logger
)

func Logger() openlog.Logger { _ = "STUB: not implemented"; return *new(openlog.Logger) }

type dataSourceEngine func() DataSource

func GetDataSource() DataSource { _ = "STUB: not implemented"; return *new(DataSource) }

func RegisterPlugin(name string, engineFunc dataSourceEngine) { _ = "STUB: not implemented"; return }

type Config struct {
	Kind   string
	Logger openlog.Logger
}

func Init(c *Config) error { _ = "STUB: not implemented"; return nil }

func GetTaskDao() TaskDao { _ = "STUB: not implemented"; return *new(TaskDao) }

func GetTombstoneDao() TombstoneDao { _ = "STUB: not implemented"; return *new(TombstoneDao) }
