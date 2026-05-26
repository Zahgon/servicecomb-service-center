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

package account

import (
	"context"
	"time"
)

const CleanupInterval = 1 * time.Minute

func init() {
	startReleasedLockHistoryCleanupJob()
}

// startReleasedLockHistoryCleanupJob cause of accountsvc.IsBanned may be never called
// after locked account, then run a job to cleanup the released lock
func startReleasedLockHistoryCleanupJob() { _ = "STUB: not implemented"; return }

func CleanupReleasedLockHistory(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
