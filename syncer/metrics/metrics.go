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

const (
	FamilyName        = "syncer"
	KeyPendingEvent   = FamilyName + "_pending_event"
	KeyAbandonEvent   = FamilyName + "_abandon_event"
	KeyPendingTask    = FamilyName + "_pending_task"
	KeyConnectedPeers = FamilyName + "_connected_peers"
	KeyPeersTotal     = FamilyName + "_peers_total"
	KeyPeersClockDiff = FamilyName + "_peers_clock_diff"
)

var Instance string

func Init() error { _ = "STUB: not implemented"; return nil }

//TODO should call metrics.Init()

func PendingEventSet(n int64) { _ = "STUB: not implemented"; return }

func AbandonEventAdd() { _ = "STUB: not implemented"; return }

func PendingTaskSet(n int64) { _ = "STUB: not implemented"; return }

func ConnectedPeersSet(n int64) { _ = "STUB: not implemented"; return }

func PeersTotalSet(n int64) { _ = "STUB: not implemented"; return }

func PeersClockDiffSet(peerName string, n int64) { _ = "STUB: not implemented"; return }
