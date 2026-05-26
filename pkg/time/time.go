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

package time

import "time"

type Time struct {
	sec  int64
	nsec int64
}

func (t Time) String() string { _ = "STUB: not implemented"; return "" }

func (t Time) UTC() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (t Time) Local() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func FromTime(t time.Time) Time { _ = "STUB: not implemented"; return *new(Time) }
