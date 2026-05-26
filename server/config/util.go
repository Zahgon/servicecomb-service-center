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

func newOptions(key string, opts []Option) *Options { _ = "STUB: not implemented"; return nil }

// GetString return the string type value by specified key
func GetString(key, def string, opts ...Option) string { _ = "STUB: not implemented"; return "" }

// GetBool return the boolean type value by specified key
func GetBool(key string, def bool, opts ...Option) bool { _ = "STUB: not implemented"; return false }

// GetInt return the int type value by specified key
func GetInt(key string, def int, opts ...Option) int { _ = "STUB: not implemented"; return 0 }

// GetInt64 return the int64 type value by specified key
func GetInt64(key string, def int64, opts ...Option) int64 { _ = "STUB: not implemented"; return 0 }

func GetStringMap(key string) map[string]string { _ = "STUB: not implemented"; return nil }

func getMapFunc(key string, dataFunc func(k string, v interface{})) {
	_ = "STUB: not implemented"
	return
}

// GetDuration return the time.Duration type value by specified key
func GetDuration(key string, def time.Duration, opts ...Option) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
