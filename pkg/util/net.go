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
	"context"
	"net"
	"net/http"
)

const CtxRemoteIP CtxKey = "x-remote-ip"

type IPPort struct {
	IP   string
	Port uint16
}

func GetIPFromContext(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func ParseEndpoint(ep string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ParseIPPort(addr string) IPPort { _ = "STUB: not implemented"; return *new(IPPort) }

func GetRealIP(r *http.Request) string { _ = "STUB: not implemented"; return "" }

func InetNtoIP(ipnr uint32) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func InetNtoa(ipnr uint32) string { _ = "STUB: not implemented"; return "" }

func InetAton(ip string) (ipnr uint32) { _ = "STUB: not implemented"; return 0 }

func ParseRequestURL(r *http.Request) string { _ = "STUB: not implemented"; return "" }
