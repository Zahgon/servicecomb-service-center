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

package util

import (
	"unsafe"
)

const intSize = int(unsafe.Sizeof(0))

var bs *[intSize]byte

func init() {
	i := 0x1
	bs = (*[intSize]byte)(unsafe.Pointer(&i))
}

func IsBigEndian() bool { _ = "STUB: not implemented"; return false }

func IsLittleEndian() bool { _ = "STUB: not implemented"; return false }

func PathExist(path string) bool { _ = "STUB: not implemented"; return false }

func HostName() (hostname string) { _ = "STUB: not implemented"; return "" }

func GetEnvInt(name string, def int) int { _ = "STUB: not implemented"; return 0 }

func GetEnvString(name string, def string) string { _ = "STUB: not implemented"; return "" }

func GetProcCPUUsage() (pt float64, ct float64) { _ = "STUB: not implemented"; return 0, 0 }
