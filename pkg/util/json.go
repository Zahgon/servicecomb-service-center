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

package util

type JSONObject map[string]interface{}

func (c JSONObject) Set(k interface{}, v interface{}) JSONObject {
	_ = "STUB: not implemented"
	return *new(JSONObject)
}

func (c JSONObject) Bool(k interface{}, def bool) bool { _ = "STUB: not implemented"; return false }

func (c JSONObject) Int(k interface{}, def int) int { _ = "STUB: not implemented"; return 0 }

func (c JSONObject) String(k interface{}, def string) string { _ = "STUB: not implemented"; return "" }

func (c JSONObject) Object(k interface{}) JSONObject {
	_ = "STUB: not implemented"
	return *new(JSONObject)
}

func toString(v interface{}) string { _ = "STUB: not implemented"; return "" }

func NewJSONObject() JSONObject { _ = "STUB: not implemented"; return *new(JSONObject) }
