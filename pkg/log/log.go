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

package log

import (
	"time"

	"github.com/go-chassis/openlog"
)

const (
	globalCallerSkip        = 2
	globalRecoverCallerSkip = 4
	defaultLogLevel         = "DEBUG"
)

var (
	flushFunc   = func() {}
	recoverFunc = func(r interface{}) {}
	Logger      = NewLogger(DefaultConfig())
)

func Init(cfg Config) { _ = "STUB: not implemented"; return }

func NewLogger(cfg Config) openlog.Logger { _ = "STUB: not implemented"; return *new(openlog.Logger) }

func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

func Debug(msg string) { _ = "STUB: not implemented"; return }

func Info(msg string) { _ = "STUB: not implemented"; return }

func Warn(msg string) { _ = "STUB: not implemented"; return }

func Error(msg string, err error) { _ = "STUB: not implemented"; return }

func Fatal(msg string, err error) { _ = "STUB: not implemented"; return }

func Flush() { _ = "STUB: not implemented"; return }

func NilOrWarn(start time.Time, message string) { _ = "STUB: not implemented"; return }

func DebugOrWarn(start time.Time, message string) { _ = "STUB: not implemented"; return }

func InfoOrWarn(start time.Time, message string) { _ = "STUB: not implemented"; return }

// Panic is a function can only be called in defer function.
func Panic(r interface{}) {
	_ = "STUB: not implemented"

	// Recover is a function call recover() and print the stack in log
	// Please call this function like 'defer log.Recover()' in your code
	return
}

func Recover() { _ = "STUB: not implemented"; return }
