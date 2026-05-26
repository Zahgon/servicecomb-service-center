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

package schema

import (
	"errors"
	"net"

	"github.com/labstack/echo/v4"
)

var (
	ErrInvalidInstanceIP = errors.New("invalid HTTP Header X-InstanceIP")
)

type Mux struct {
	// Disable represents frontend proxy service api or not
	Disable        bool
	SchemaTestCIDR *net.IPNet
}

func (m *Mux) SchemaHandleFunc(c echo.Context) (err error) { _ = "STUB: not implemented"; return nil }

func (m *Mux) checkInstanceHost(instanceIP string) error { _ = "STUB: not implemented"; return nil }
