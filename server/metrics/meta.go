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
	"time"

	metricsvc "github.com/apache/servicecomb-service-center/pkg/metrics"
)

const (
	SubSystem            = "db"
	KeyServiceTotal      = metricsvc.FamilyName + "_" + SubSystem + "_" + "service_total"
	KeyInstanceTotal     = metricsvc.FamilyName + "_" + SubSystem + "_" + "instance_total"
	KeyServiceUsage      = metricsvc.FamilyName + "_" + SubSystem + "_" + "service_usage"
	KeyInstanceUsage     = metricsvc.FamilyName + "_" + SubSystem + "_" + "instance_usage"
	KeyDomainTotal       = metricsvc.FamilyName + "_" + SubSystem + "_" + "domain_total"
	KeySchemaTotal       = metricsvc.FamilyName + "_" + SubSystem + "_" + "schema_total"
	KeyFrameworkTotal    = metricsvc.FamilyName + "_" + SubSystem + "_" + "framework_total"
	KeyHeartbeatTotal    = metricsvc.FamilyName + "_" + SubSystem + "_" + "heartbeat_total"
	KeyHeartbeatDuration = metricsvc.FamilyName + "_" + SubSystem + "_" + "heartbeat_durations_microseconds"
	KeySCTotal           = metricsvc.FamilyName + "_" + SubSystem + "_" + "sc_total"
)

var metaEnabled = false

func InitMetaMetrics() (err error) { _ = "STUB: not implemented"; return nil }

func GetTotalService(domain, project string) int64 { _ = "STUB: not implemented"; return 0 }

func GetTotalInstance(domain, project string) int64 { _ = "STUB: not implemented"; return 0 }

func ReportScInstance() { _ = "STUB: not implemented"; return }

func ReportHeartbeatCompleted(err error, start time.Time) { _ = "STUB: not implemented"; return }
