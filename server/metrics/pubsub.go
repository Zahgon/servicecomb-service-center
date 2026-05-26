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
	"github.com/apache/servicecomb-service-center/pkg/event"

	"github.com/apache/servicecomb-service-center/pkg/metrics"
	helper "github.com/apache/servicecomb-service-center/pkg/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	notifyCounter = helper.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metrics.FamilyName,
			Subsystem: "notify",
			Name:      "publish_total",
			Help:      "Counter of publishing instance events",
		}, []string{"instance", "source", "status"})

	notifyLatency = helper.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  metrics.FamilyName,
			Subsystem:  "notify",
			Name:       "publish_durations_microseconds",
			Help:       "Latency of publishing instance events",
			Objectives: metrics.Pxx,
		}, []string{"instance", "source", "status"})

	pendingGauge = helper.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metrics.FamilyName,
			Subsystem: "notify",
			Name:      "pending_total",
			Help:      "Counter of pending instance events",
		}, []string{"instance", "source"})

	pendingLatency = helper.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  metrics.FamilyName,
			Subsystem:  "notify",
			Name:       "pending_durations_microseconds",
			Help:       "Latency of pending instance events",
			Objectives: metrics.Pxx,
		}, []string{"instance", "source"})

	subscriberGauge = helper.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metrics.FamilyName,
			Subsystem: "notify",
			Name:      "subscriber_total",
			Help:      "Gauge of subscribers",
		}, []string{"instance", "domain", "scheme"})
)

func ReportPublishCompleted(evt event.Event, err error) { _ = "STUB: not implemented"; return }

func ReportPendingCompleted(evt event.Event) { _ = "STUB: not implemented"; return }

func ReportSubscriber(domain, scheme string, n float64) { _ = "STUB: not implemented"; return }
