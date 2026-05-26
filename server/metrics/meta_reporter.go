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

package metrics

import (
	"github.com/apache/servicecomb-service-center/datasource"
)

var metaReporter = &MetaReporter{}

type MetaReporter struct {
}

func (m *MetaReporter) DomainAdd(delta float64) { _ = "STUB: not implemented"; return }

func (m *MetaReporter) ServiceAdd(delta float64, ml datasource.MetricsLabels) {
	_ = "STUB: not implemented"
	return
}

func (m *MetaReporter) ServiceUsageSet() { _ = "STUB: not implemented"; return }

func (m *MetaReporter) InstanceAdd(delta float64, ml datasource.MetricsLabels) {
	_ = "STUB: not implemented"
	return
}

func (m *MetaReporter) InstanceUsageSet() { _ = "STUB: not implemented"; return }

func (m *MetaReporter) SchemaAdd(delta float64, ml datasource.MetricsLabels) {
	_ = "STUB: not implemented"
	return
}

func (m *MetaReporter) FrameworkSet(ml datasource.MetricsLabels) { _ = "STUB: not implemented"; return }

func GetMetaReporter() *MetaReporter { _ = "STUB: not implemented"; return nil }

func ResetMetaMetrics() { _ = "STUB: not implemented"; return }

func ReportMetaMetrics() { _ = "STUB: not implemented"; return }
