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

package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

// Vectors is unsafe, so all NewXXXVec funcs should be called during the initialization phase
var Vectors = make(map[string]prometheus.Collector)

func registerMetrics(name string, vec prometheus.Collector) { _ = "STUB: not implemented"; return }

func NewCounterVec(opts prometheus.CounterOpts, labelNames []string) *prometheus.CounterVec {
	_ = "STUB: not implemented"
	return nil
}

func NewGaugeVec(opts prometheus.GaugeOpts, labelNames []string) *prometheus.GaugeVec {
	_ = "STUB: not implemented"
	return nil
}

func NewSummaryVec(opts prometheus.SummaryOpts, labelNames []string) *prometheus.SummaryVec {
	_ = "STUB: not implemented"
	return nil
}

func Gather() ([]*dto.MetricFamily, error) { _ = "STUB: not implemented"; return nil, nil }

func HTTPHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }
