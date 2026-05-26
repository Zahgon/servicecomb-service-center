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

package sd

type IndexFunc func(interface{}) string

type IndexCols struct {
	indexFuncs []IndexFunc
}

var DepIndexCols *IndexCols
var InstIndexCols *IndexCols
var ServiceIndexCols *IndexCols
var RuleIndexCols *IndexCols

func NewIndexCols() *IndexCols { _ = "STUB: not implemented"; return nil }

func (i *IndexCols) AddIndexFunc(f IndexFunc) { _ = "STUB: not implemented"; return }

func (i *IndexCols) GetIndexes(data interface{}) (res []string) {
	_ = "STUB: not implemented"
	return nil
}
