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
	"github.com/apache/servicecomb-service-center/pkg/metrics"
	dto "github.com/prometheus/client_model/go"
)

const (
	httpRequestTotal = "http_request_total"
)

var qpsLabelMap = map[string]int{
	"method":   0,
	"instance": 1,
	"api":      2,
	"domain":   3,
}

// APIReporter is used to calc the http TPS
type APIReporter struct {
	cache *metrics.Details
}

func (r *APIReporter) Report() { _ = "STUB: not implemented"; return }

func (r *APIReporter) toLabels(pairs []*dto.LabelPair) (labels []string) {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	metrics.RegisterReporter("rest", NewAPIReporter())
}

func NewAPIReporter() *APIReporter { _ = "STUB: not implemented"; return nil }
