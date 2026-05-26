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

package admin

import (
	"errors"

	"google.golang.org/grpc"
)

const (
	scheme      = "grpc"
	serviceName = "syncer"
)

var (
	peerInfos        []*PeerInfo
	ErrConfigIsEmpty = errors.New("sync config is empty")
)

type Resp struct {
	Peers []*Peer `json:"peers"`
}

type PeerInfo struct {
	Peer       *Peer
	ClientConn *grpc.ClientConn
}

type Peer struct {
	Name      string   `json:"name"`
	Kind      string   `json:"kind"`
	Mode      []string `json:"mode"`
	Endpoints []string `json:"endpoints"`
	Status    string   `json:"status"`
	Token     string   `json:"-"`
}

func Init() { _ = "STUB: not implemented"; return }

func Health() (*Resp, error) { _ = "STUB: not implemented"; return nil, nil }

func getPeerStatus(peerInfo *PeerInfo) string { _ = "STUB: not implemented"; return "" }

func reportClockDiff(peerName string, local int64, resp int64) { _ = "STUB: not implemented"; return }

func reportMetrics(peers []*Peer) { _ = "STUB: not implemented"; return }

func newRPCConn(endpoints []string) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Peers() []*PeerInfo { _ = "STUB: not implemented"; return nil }
