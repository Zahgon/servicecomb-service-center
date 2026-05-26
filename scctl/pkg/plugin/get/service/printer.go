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

package service

import (
	"github.com/apache/servicecomb-service-center/scctl/pkg/model"
	"github.com/apache/servicecomb-service-center/scctl/pkg/writer"
)

const maxWidth = 35

var (
	longServiceTableHeader   = []string{"DOMAIN", "NAME", "APPID", "VERSIONS", "ENV", "FRAMEWORK", "ENDPOINTS", "AGE"}
	domainServiceTableHeader = []string{"DOMAIN", "NAME", "APPID", "VERSIONS", "ENV", "FRAMEWORK", "AGE"}
	shortServiceTableHeader  = []string{"NAME", "APPID", "VERSIONS", "ENV", "FRAMEWORK", "AGE"}
)

type Record struct {
	model.Service
}

func (s *Record) VersionsString() string { _ = "STUB: not implemented"; return "" }

func (s *Record) FrameworksString() string { _ = "STUB: not implemented"; return "" }

func (s *Record) EndpointsString() string { _ = "STUB: not implemented"; return "" }

func (s *Record) AgeString() string { _ = "STUB: not implemented"; return "" }

func (s *Record) Domain() string { _ = "STUB: not implemented"; return "" }

func (s *Record) PrintBody(fmt string, all bool) []string { _ = "STUB: not implemented"; return nil }

type Printer struct {
	Records map[string]*Record
	flags   []interface{}
}

func (sp *Printer) SetOutputFormat(f string, all bool) { _ = "STUB: not implemented"; return }

func (sp *Printer) Flags(flags ...interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func (sp *Printer) PrintBody() (slice [][]string) { _ = "STUB: not implemented"; return nil }

func (sp *Printer) PrintTitle() []string { _ = "STUB: not implemented"; return nil }

func (sp *Printer) Sorter() *writer.RecordsSorter { _ = "STUB: not implemented"; return nil }
