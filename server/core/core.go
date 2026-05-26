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

package core

import (
	"time"

	//go-chassis plugin
	_ "github.com/go-chassis/go-chassis-extension/codec/gojson"
	_ "github.com/go-chassis/go-chassis-extension/protocol/fiber4r"
	_ "github.com/go-chassis/go-chassis-extension/protocol/grpc/server"
	_ "github.com/go-chassis/go-chassis/v2/middleware/monitoring"
	_ "github.com/go-chassis/go-chassis/v2/middleware/ratelimiter"

	// import the grace package and parse grace cmd line
	_ "github.com/apache/servicecomb-service-center/pkg/grace"
)

const (
	defaultCollectPeriod = 30 * time.Second
)

// Init init chassis and sc configs
func Init() { _ = "STUB: not implemented"; return }

// initialize configuration

// Logging

// go pool

// init the sc registration

// Register global services

// init metrics

func initLogger() { _ = "STUB: not implemented"; return }

func initMetrics() { _ = "STUB: not implemented"; return }
