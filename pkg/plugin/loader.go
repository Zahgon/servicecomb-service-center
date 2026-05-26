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
	"plugin"
	"regexp"
	"sync"
)

var (
	loader   Loader
	once     sync.Once
	regex, _ = regexp.Compile(`([A-Za-z0-9_.-]+)_plugin.so$`)
)

type wrapPlugin struct {
	p     *plugin.Plugin
	funcs map[string]plugin.Symbol
}

type Loader struct {
	Plugins map[string]*wrapPlugin
	mux     sync.RWMutex
}

func (pm *Loader) Init() { _ = "STUB: not implemented"; return }

func (pm *Loader) ReloadPlugins() error { _ = "STUB: not implemented"; return nil }

// golang 1.8+ feature

func (pm *Loader) Find(pluginName, funcName string) (plugin.Symbol, error) {
	_ = "STUB: not implemented"
	return *new(plugin.Symbol), nil
}

func (pm *Loader) Exist(pluginName string) bool { _ = "STUB: not implemented"; return false }

func GetLoader() *Loader { _ = "STUB: not implemented"; return nil }

func Reload() error { _ = "STUB: not implemented"; return nil }

func FindFunc(pluginName, funcName string) (plugin.Symbol, error) {
	_ = "STUB: not implemented"
	return *new(plugin.Symbol), nil
}
