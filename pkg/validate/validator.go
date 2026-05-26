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

package validate

import (
	"sync"
)

type Validator struct {
	rules map[string]*Rule
	subs  map[string]*Validator
	once  sync.Once
}

func (v *Validator) Init(f func(*Validator)) *Validator { _ = "STUB: not implemented"; return nil }

func (v *Validator) GetRule(name string) *Rule { _ = "STUB: not implemented"; return nil }

func (v *Validator) AddRule(name string, rule *Rule) { _ = "STUB: not implemented"; return }

func (v *Validator) RemoveRule(name string) { _ = "STUB: not implemented"; return }

func (v *Validator) GetRules() map[string](*Rule) { _ = "STUB: not implemented"; return nil }

func (v *Validator) AddRules(in map[string](*Rule)) { _ = "STUB: not implemented"; return }

func (v *Validator) GetSub(name string) *Validator { _ = "STUB: not implemented"; return nil }

func (v *Validator) AddSub(name string, s *Validator) { _ = "STUB: not implemented"; return }

func (v *Validator) GetSubs() map[string](*Validator) { _ = "STUB: not implemented"; return nil }

func (v *Validator) AddSubs(in map[string](*Validator)) { _ = "STUB: not implemented"; return }

func (v *Validator) Validate(s interface{}) error { _ = "STUB: not implemented"; return nil }

// check current rule
// if pointer, check it's a nil pointer or not
// if array, slice and map, check the length and regex
// if sub type is not a string when do regex check, return OK

// check sub rule
// do not support sub type is not pointer or struct

// TODO how to validate non-base type key

func NewValidator() *Validator { _ = "STUB: not implemented"; return nil }
