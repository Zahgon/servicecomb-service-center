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

package state

var (
	repoPlugins   = make(map[string]newRepoFunc)
	manager       *Manager
	configuration Config
)

type newRepoFunc func(opts Config) Repository

// Install load plugins configuration into plugins
func Install(pluginImplName string, newFunc newRepoFunc) { _ = "STUB: not implemented"; return }

func Init(opts Config) error { _ = "STUB: not implemented"; return nil }

func NewRepository(opts Config) (Repository, error) {
	_ = "STUB: not implemented"
	return *new(Repository), nil
}

func Instance() *Manager { _ = "STUB: not implemented"; return nil }

func Revision() int64 { _ = "STUB: not implemented"; return 0 }

func Configuration() Config { _ = "STUB: not implemented"; return *new(Config) }
