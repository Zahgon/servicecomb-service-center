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

package bootstrap

import (
	"context"
	"time"

	"github.com/apache/servicecomb-service-center/istio/pkg/controllers/istioconnector"
	"github.com/apache/servicecomb-service-center/istio/pkg/controllers/servicecenter"
	"github.com/apache/servicecomb-service-center/istio/pkg/event"
)

// cli args
type Args struct {
	// servicecomb-service-center address
	ServiceCenterAddr string
	// kubeconfig file path
	Kubeconfig string
	// enable leader election or not for high abalibility
	HA bool
}

const (
	// leader election check locked resource namespace
	lockNameSpace = "istio-system"
	// leader election check locked resource name
	resourceName = "servicecenter2mesh"
	// LeaseDuration is the duration that non-leader candidates will
	// wait to force acquire leadership. This is measured against time of
	// last observed ack.
	//
	// A client needs to wait a full LeaseDuration without observing a change to
	// the record before it can attempt to take over. When all clients are
	// shutdown and a new set of clients are started with different names against
	// the same leader record, they must wait the full LeaseDuration before
	// attempting to acquire the lease. Thus LeaseDuration should be as short as
	// possible (within your tolerance for clock skew rate) to avoid a possible
	// long waits in the scenario.
	//
	// Core clients default this value to 15 seconds.
	defaultLeaseDuration = 15 * time.Second
	// RenewDeadline is the duration that the acting master will retry
	// refreshing leadership before giving up.
	//
	// Core clients default this value to 10 seconds.
	defaultRenewDeadline = 10 * time.Second
	// RetryPeriod is the duration the LeaderElector clients should wait
	// between tries of actions.
	//
	// Core clients default this value to 2 seconds.
	defaultRetryPeriod = 2 * time.Second
)

type Server struct {
	// service center controller watches service center update, and push to istio controller
	serviceCenterController *servicecenter.Controller
	// istio controller receives updates from service center controller and push to k8s api server
	istioController *istioconnector.Controller
	// channel for passing service center event from service center controller to istio controller
	serviceCenterEvent chan []event.ChangeEvent
}

func NewServer(args *Args) (*Server, error) {
	_ = "STUB: not implemented"
	// only allow 1 eventlist at a time
	return nil, nil
}

// Create a new istio controller, the controller is ready to push configs to istio

// start the server need to start both service center and istio controller
func (s *Server) Start(ctx context.Context, args *Args) error {
	_ = "STUB: not implemented"
	// by default the leader election is disabled, just do regular start
	return nil
}

// This function is used to enable leader election using k8s client-go api. leaderElectAndRun runs the leader election,
// and runs the callbacks once the leader lease is acquired.
//
// For k8s clint-go API:
// DISCLAIMER: this is an alpha API. This library will likely change significantly or even be removed entirely in subsequent releases.
// Depend on this API at your own risk.
//
// Note: this API is also used by K8S controller and Cluster auto scaler.
func (s *Server) doLeaderElectionRun(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// creates the in-cluster config

func (s *Server) doRun(ctx context.Context) { _ = "STUB: not implemented"; return }

// on server stop
func (s *Server) waitForShutdown(ctx context.Context) { _ = "STUB: not implemented"; return }
