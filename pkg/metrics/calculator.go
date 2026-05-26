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
	dto "github.com/prometheus/client_model/go"
)

var (
	DefaultCalculator Calculator = &CommonCalculator{}
)

// Calculator is the interface to implement customize algorithm of MetricFamily
type Calculator interface {
	Calc(mf *dto.MetricFamily) *Details
}

type CommonCalculator struct {
}

// Get value of metricFamily
func (c *CommonCalculator) Calc(mf *dto.MetricFamily) *Details {
	_ = "STUB: not implemented"
	return nil
}

func metricGaugeOf(details *Details, m []*dto.Metric) { _ = "STUB: not implemented"; return }

func metricCounterOf(details *Details, m []*dto.Metric) { _ = "STUB: not implemented"; return }

func metricSummaryOf(details *Details, m []*dto.Metric) { _ = "STUB: not implemented"; return }

func metricHistogramOf(details *Details, m []*dto.Metric) { _ = "STUB: not implemented"; return }

func RegisterCalculator(c Calculator) { _ = "STUB: not implemented"; return }

func Calculate(mf *dto.MetricFamily) *Details { _ = "STUB: not implemented"; return nil }
