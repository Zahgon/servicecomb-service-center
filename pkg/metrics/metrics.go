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
	"github.com/apache/servicecomb-service-center/pkg/buffer"
	dto "github.com/prometheus/client_model/go"
)

// Pxx represents p99 p90 p50
var Pxx = map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}

func NewMetrics() *Metrics { _ = "STUB: not implemented"; return nil }

func NewDetails() *Details { _ = "STUB: not implemented"; return nil }

// Details is the struct to hold the calculated result and index by metric label
type Details struct {
	// Summary is the calculation results of the details
	Summary float64

	mapper map[string]float64
	buffer *buffer.Pool
}

// to format 'N1=L1,N2=L2,N3=L3,...'
func (cm *Details) toKey(labels []*dto.LabelPair) string { _ = "STUB: not implemented"; return "" }

func (cm *Details) toLabels(key string) (p []*dto.LabelPair) { _ = "STUB: not implemented"; return nil }

func (cm *Details) Get(labels []*dto.LabelPair) (val float64) { _ = "STUB: not implemented"; return 0 }

func (cm *Details) put(labels []*dto.LabelPair, val float64) { _ = "STUB: not implemented"; return }

func (cm *Details) ForEach(f func(labels []*dto.LabelPair, v float64) (next bool)) {
	_ = "STUB: not implemented"
	return
}

// Metrics is the struct to hold the Details objects store and index by metric name
type Metrics struct {
	mapper map[string]*Details
}

func (cm *Metrics) put(key string, val *Details) { _ = "STUB: not implemented"; return }

func (cm *Metrics) Get(key string) (val *Details) { _ = "STUB: not implemented"; return nil }

func (cm *Metrics) ForEach(f func(k string, v *Details) (next bool)) {
	_ = "STUB: not implemented"
	return
}

func (cm *Metrics) Summary(key string) (sum float64) { _ = "STUB: not implemented"; return 0 }

const (
	tagJSON = "json"
)

// ToRawData parses result form labels
func ToRawData(result interface{}, labels []*dto.LabelPair) { _ = "STUB: not implemented"; return }

// ToLabelNames returns label names, count is special label of v of func ForEach
func ToLabelNames(structure interface{}) []string { _ = "STUB: not implemented"; return nil }
