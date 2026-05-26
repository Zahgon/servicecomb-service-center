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

package plugin

import (
	"sync"
)

const (
	defaultPluginSize     = 20
	defaultPluginImplSize = 5
)

var (
	globalConfigurator Configurator = &DefaultConfigurator{}
	pluginMgr                       = &Manager{}
)

func init() {
	pluginMgr.Initialize()
}

type wrapInstance struct {
	dynamic  bool
	instance Instance
	lock     sync.RWMutex
}

// Manager manages plugin instance generation.
// Manager keeps the plugin instance currently used by server
// for every plugin interface.
type Manager struct {
	plugins   map[Kind]map[ImplName]*Plugin
	instances map[Kind]*wrapInstance
}

// Initialize initializes the struct
func (pm *Manager) Initialize() { _ = "STUB: not implemented"; return }

// ReloadAll reloads all the plugin instances
func (pm *Manager) ReloadAll() { _ = "STUB: not implemented"; return }

// Register registers a 'Plugin'
// unsafe
func (pm *Manager) Register(p Plugin) { _ = "STUB: not implemented"; return }

// Get gets a 'Plugin'
func (pm *Manager) Get(pn Kind, name ImplName) *Plugin { _ = "STUB: not implemented"; return nil }

// Instance gets an plugin instance.
// What plugin instance you get is depended on the supplied go plugin files
// (high priority) or the plugin config(low priority)
//
// The go plugin file should be {plugins_dir}/{Kind}_plugin.so.
// ('plugins_dir' must be configured as a valid path in service-center config.)
// The plugin config in service-center config should be:
// {Kind}_plugin = {ImplName}
//
// e.g. For registry plugin, you can set a config in app.conf:
// plugins_dir = /home, and supply a go plugin file: /home/registry_plugin.so;
// or if you want to use etcd as registry, you can set a config in app.conf:
// registry_plugin = etcd.
func (pm *Manager) Instance(pn Kind) Instance { _ = "STUB: not implemented"; return *new(Instance) }

// New initializes and sets the instance of a plugin interface,
// but not returns it.
// Use 'Instance' if you want to get the plugin instance.
// We suggest you to use 'Instance' instead of 'New'.
func (pm *Manager) New(pn Kind) { _ = "STUB: not implemented"; return }

// Dynamic plugin has high priority.

// Reload reloads the instance of the specified plugin interface.
func (pm *Manager) Reload(pn Kind) { _ = "STUB: not implemented"; return }

func (pm *Manager) existDynamicPlugin(pn Kind) *Plugin { _ = "STUB: not implemented"; return nil }

// 'buildin' implement of all plugins should call DynamicPluginFunc()

func (pm *Manager) IsDynamicPlugin(pn Kind) bool { _ = "STUB: not implemented"; return false }

// Plugins returns the 'Manager'.
func Plugins() *Manager {
	_ = "STUB: not implemented"

	// RegisterPlugin registers a 'Plugin'.
	return nil
}

func RegisterPlugin(p Plugin) { _ = "STUB: not implemented"; return }

// LoadPlugins loads and sets all the plugin interfaces's instance.
func LoadPlugins() { _ = "STUB: not implemented"; return }

func GetConfigurator() Configurator {
	_ = "STUB: not implemented"
	return *

	// RegisterConfigurator registers the customize Configurator impl
	new(Configurator)
}

func RegisterConfigurator(cfg Configurator) { _ = "STUB: not implemented"; return }
