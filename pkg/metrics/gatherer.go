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
	"context"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set"
)

const prefixName = FamilyName + "_"

var families = mapset.NewSet(FamilyName)

// EmptyGather just active when metrics disabled
var EmptyGather = &Gather{
	Records: NewMetrics(),
	closed:  false,
}

func NewGatherer(opts Options) *Gather { _ = "STUB: not implemented"; return nil }

type Gather struct {
	Records  *Metrics
	Interval time.Duration

	lock   sync.Mutex
	closed bool
}

func (mm *Gather) Start() { _ = "STUB: not implemented"; return }

func (mm *Gather) loop(ctx context.Context) { _ = "STUB: not implemented"; return }

func (mm *Gather) Collect() error { _ = "STUB: not implemented"; return nil }

// clean the old cache here

func RecordName(metricName string) string { _ = "STUB: not implemented"; return "" }

// just compatible with sc old metric name without familyName

func CollectFamily(familyName string) { _ = "STUB: not implemented"; return }

func ParseFamily(metricName string) string { _ = "STUB: not implemented"; return "" }
