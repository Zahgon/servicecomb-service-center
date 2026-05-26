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

var config Config

type Config struct {
	Sync *Sync `yaml:"sync"`
}

type Sync struct {
	EnableOnStart bool `yaml:"enableOnStart"`
	// When RbacEnabled is true, syncer's API requires the rbac token,
	// and service-center also provides the rbac token to communicate with peer.
	// At the same time, service-center rbac must be enabled.
	RbacEnabled bool    `yaml:"rbacEnabled"`
	Peers       []*Peer `yaml:"peers"`
}

type Peer struct {
	Name      string   `yaml:"name"`
	Kind      string   `yaml:"kind"`
	Endpoints []string `yaml:"endpoints"`
	Mode      []string `yaml:"mode"`
	// The token to communicate with peer, this takes effect only when RbacEnabled is true
	Token string `yaml:"token"`
}

func Init() error { _ = "STUB: not implemented"; return nil }

// Reload all configurations
func Reload() error { _ = "STUB: not implemented"; return nil }

// GetConfig return the syncer full configurations
func GetConfig() Config {
	_ = "STUB: not implemented"

	// SetConfig for UT
	return *new(Config)
}

func SetConfig(c Config) { _ = "STUB: not implemented"; return }
