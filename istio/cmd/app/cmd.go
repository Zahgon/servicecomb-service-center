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

package app

import (
	"context"

	"github.com/apache/servicecomb-service-center/istio/pkg/bootstrap"

	"github.com/spf13/cobra"
	"istio.io/pkg/log"
)

var inputArgs *bootstrap.Args
var loggingOptions = log.DefaultOptions()

// NewRootCommand creates servicecomb-service-center-istio service cli args
func NewRootCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Create the stop channel for all of the servers.
// ctx, cancelFunc := context.WithCancel(context.Background())

// defer cancelFunc()
// Create the server for the servicecomb-service-center-istio service.

// Start the server

// WaitSignal awaits for SIGINT or SIGTERM and closes the channel
func waitSignal(ctx context.Context) { _ = "STUB: not implemented"; return }

func addFlags(c *cobra.Command) { _ = "STUB: not implemented"; return }

// Process commandline args.

// sc-addr is the service center registry centre address

// enable leader-election or not

// kubectl config file path, if not set, will use in cluster kube config

// Attach the Istio logging options to the command.
