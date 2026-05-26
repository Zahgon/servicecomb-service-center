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

// metrics pkg declare http metrics and impl a reporter to set the
// 'query_per_seconds' value periodically
package metrics

import (
	"net/http"
	"time"

	"github.com/apache/servicecomb-service-center/pkg/metrics"
	helper "github.com/apache/servicecomb-service-center/pkg/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	incomingRequests = helper.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metrics.FamilyName,
			Subsystem: "http",
			Name:      "request_total",
			Help:      "Counter of requests received into ROA handler",
		}, []string{"method", "code", "instance", "api", "domain"})

	successfulRequests = helper.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metrics.FamilyName,
			Subsystem: "http",
			Name:      "success_total",
			Help:      "Counter of successful requests processed by ROA handler",
		}, []string{"method", "code", "instance", "api", "domain"})

	reqDurations = helper.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  metrics.FamilyName,
			Subsystem:  "http",
			Name:       "request_durations_microseconds",
			Help:       "HTTP request latency summary of ROA handler",
			Objectives: metrics.Pxx,
		}, []string{"method", "instance", "api", "domain"})

	queryPerSeconds = helper.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metrics.FamilyName,
			Subsystem: "http",
			Name:      "query_per_seconds",
			Help:      "HTTP requests per seconds of ROA handler",
		}, []string{"method", "instance", "api", "domain"})
)

func ReportRequestCompleted(_ http.ResponseWriter, r *http.Request, start time.Time) {
	_ = "STUB: not implemented"
	return
}

func parseStatus(code int) (bool, string) { _ = "STUB: not implemented"; return false, "" }
