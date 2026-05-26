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
	"os"

	"github.com/go-chassis/openlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	StdoutSyncer = zapcore.Lock(os.Stdout)
	StderrSyncer = zapcore.Lock(os.Stderr)

	zapLevelMap = map[string]zapcore.Level{
		"DEBUG": zap.DebugLevel,
		"INFO":  zap.InfoLevel,
		"WARN":  zap.WarnLevel,
		"ERROR": zap.ErrorLevel,
		"FATAL": zap.FatalLevel,
	}
)

func toZapConfig(c Config) zapcore.Core {
	_ = "STUB: not implemented"
	// level config
	return *new(zapcore.Core)
}

// log format

// log rotate

//zap.NewDevelopment()

type ZapLogger struct {
	Config Config

	zapLogger *zap.Logger
	zapSugar  *zap.SugaredLogger
}

func (l *ZapLogger) Debug(msg string, _ ...openlog.Option) { _ = "STUB: not implemented"; return }

func (l *ZapLogger) Info(msg string, _ ...openlog.Option) { _ = "STUB: not implemented"; return }

func (l *ZapLogger) Warn(msg string, _ ...openlog.Option) { _ = "STUB: not implemented"; return }

func (l *ZapLogger) Error(msg string, opts ...openlog.Option) { _ = "STUB: not implemented"; return }

func (l *ZapLogger) Fatal(msg string, opts ...openlog.Option) { _ = "STUB: not implemented"; return }

// Recover callSkip equals to 0 identify the caller of Recover()
func (l *ZapLogger) Recover(r interface{}, callerSkip int) { _ = "STUB: not implemented"; return }

// zapcore sync automatically when larger than ErrorLevel

// recover logs also output to stderr

// sync immediately, for server may exit abnormally

func (l *ZapLogger) Sync() { _ = "STUB: not implemented"; return }

func NewZapLogger(cfg Config) *ZapLogger { _ = "STUB: not implemented"; return nil }
