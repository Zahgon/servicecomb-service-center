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

package rbac

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

const (
	MaxAttempts   = 2
	BlockInterval = 15 * time.Minute
)

var clients sync.Map

type LoginFailureLimiter struct {
	limiter *rate.Limiter
	Key     string
}

// TryLockAccount try to lock the account login attempt
// it use time/rate to allow certainty failure,
// it will ban client if rate limiter can not accept failures
func TryLockAccount(key string) { _ = "STUB: not implemented"; return }

// IsBanned check if a client is banned, and if client ban time expire,
// it will release the client from banned status
// use account name plus ip as key will maximum reduce the client conflicts
func IsBanned(key string) bool { _ = "STUB: not implemented"; return false }
