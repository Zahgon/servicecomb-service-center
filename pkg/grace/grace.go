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

package grace

import (
	"flag"
	"os"
	"sync"
	"syscall"
)

const (
	PreSignal = iota
	PostSignal
)

var (
	isFork         bool
	filesOrder     string
	files          []*os.File
	filesOffsetMap map[string]int

	registerSignals []os.Signal
	SignalHooks     map[int]map[os.Signal][]func()
	graceMux        sync.Mutex
	forked          bool
)

func init() {
	flag.BoolVar(&isFork, "fork", false, "listen on open fd (after forking)")
	flag.StringVar(&filesOrder, "filesorder", "", "previous initialization FDs order")

	registerSignals = []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
	}
	filesOffsetMap = make(map[string]int)
	SignalHooks = map[int]map[os.Signal][]func(){
		PreSignal:  {},
		PostSignal: {},
	}
	for _, sig := range registerSignals {
		SignalHooks[PreSignal][sig] = []func(){}
		SignalHooks[PostSignal][sig] = []func(){}
	}

	go handleSignals()
}

func ParseCommandLine() { _ = "STUB: not implemented"; return }

func Before(f func()) { _ = "STUB: not implemented"; return }

func After(f func()) { _ = "STUB: not implemented"; return }

func RegisterSignalHook(phase int, f func(), sigs ...os.Signal) { _ = "STUB: not implemented"; return }

func RegisterFiles(name string, f *os.File) { _ = "STUB: not implemented"; return }

func fireSignalHook(ppFlag int, sig os.Signal) { _ = "STUB: not implemented"; return }

func handleSignals() { _ = "STUB: not implemented"; return }

func fork() (err error) { _ = "STUB: not implemented"; return nil }

// add fork and file descriptions order flags

func parseCommandLine() (args []string) { _ = "STUB: not implemented"; return nil }

// ignore process path

// ignore fork flags

func newCommand(args ...string) error { _ = "STUB: not implemented"; return nil }

func IsFork() bool { _ = "STUB: not implemented"; return false }

func ExtraFileOrder(name string) int { _ = "STUB: not implemented"; return 0 }

func Done() error { _ = "STUB: not implemented"; return nil }
