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

package config

import (
	"time"
)

const (
	InitVersion = "0"
	minCacheTTL = 5 * time.Minute
)

var (
	Server = NewServerConfig()
	// App is application root config
	App = &AppConfig{Server: Server}
)

// GetProfile return active profile
func GetProfile() *ServerConfig {
	_ = "STUB: not implemented"

	// GetGov return governance configs
	return nil
}

func GetGov() *Gov {
	_ = "STUB: not implemented"

	// GetServer return the http server configs
	return nil
}

func GetServer() ServerConfigDetail {
	_ = "STUB: not implemented"
	return *

	// GetSSL return the ssl configs
	new(ServerConfigDetail)
}

func GetSSL() ServerConfigDetail {
	_ = "STUB: not implemented"
	return *

	// GetLog return the log configs
	new(ServerConfigDetail)
}

func GetLog() ServerConfigDetail {
	_ = "STUB: not implemented"
	return *

	// GetRegistry return the registry configs
	new(ServerConfigDetail)
}

func GetRegistry() ServerConfigDetail {
	_ = "STUB: not implemented"
	return *

	// GetPlugin return the plugin configs
	new(ServerConfigDetail)
}

func GetPlugin() ServerConfigDetail {
	_ = "STUB: not implemented"
	return *

	// GetRBAC return the rbac configs
	new(ServerConfigDetail)
}

func GetRBAC() ServerConfigDetail { _ = "STUB: not implemented"; return *new(ServerConfigDetail) }

func Init() { _ = "STUB: not implemented"; return }

// Reload reload the all configurations
func Reload() error { _ = "STUB: not implemented"; return nil }

func loadServerConfig() ServerConfig { _ = "STUB: not implemented"; return *new(ServerConfig) }

// compatible with beego config's runmode

func setCPUs() { _ = "STUB: not implemented"; return }

func loadGrcConfig() *Gov { _ = "STUB: not implemented"; return nil }
