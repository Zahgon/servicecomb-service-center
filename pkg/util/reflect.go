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

import (
	"reflect"
	"sync"
)

var (
	reflector *Reflector
	unknown   = new(ReflectObject)
)

func init() {
	reflector = &Reflector{
		types: make(map[*uintptr]*ReflectObject),
	}
}

type ReflectObject struct {
	// full name
	FullName string
	Type     reflect.Type
	// if type is not struct, Fields is nil
	Fields []reflect.StructField
}

// Name returns a short name of the object type
func (o *ReflectObject) Name() string { _ = "STUB: not implemented"; return "" }

type Reflector struct {
	types map[*uintptr]*ReflectObject
	mux   sync.RWMutex
}

func (r *Reflector) Load(obj interface{}) *ReflectObject { _ = "STUB: not implemented"; return nil }

func Reflect(obj interface{}) *ReflectObject { _ = "STUB: not implemented"; return nil }

func Sizeof(obj interface{}) uint64 { _ = "STUB: not implemented"; return 0 }

func sizeof(v reflect.Value, selfRecurseMap map[uintptr]struct{}) (s uint64) {
	_ = "STUB: not implemented"
	return 0
}

func isValueType(kind reflect.Kind) bool { _ = "STUB: not implemented"; return false }

func FormatFuncName(f string) string { _ = "STUB: not implemented"; return "" }

// trim the suffix of function closure name

func FuncName(f interface{}) string { _ = "STUB: not implemented"; return "" }
