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

package rest

import (
	"crypto/tls"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	serverStateInit = iota
	serverStateRunning
	serverStateTerminating
	serverStateClosed
)

type ServerConfig struct {
	Addr              string
	Handler           http.Handler
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	IdleTimeout       time.Duration
	WriteTimeout      time.Duration
	KeepAliveTimeout  time.Duration
	GraceTimeout      time.Duration
	MaxHeaderBytes    int
	TLSConfig         *tls.Config
	Compressed        bool
	CompressMinBytes  int
}

func DefaultServerConfig() *ServerConfig { _ = "STUB: not implemented"; return nil }

// 1.4KB

func NewServer(srvCfg *ServerConfig) *Server { _ = "STUB: not implemented"; return nil }

type Server struct {
	*http.Server

	Network          string
	KeepaliveTimeout time.Duration
	GraceTimeout     time.Duration

	Listener    net.Listener
	netListener net.Listener
	tcpListener *TCPListener

	conns int64
	wg    sync.WaitGroup
	state uint8
}

func (srv *Server) Serve() (err error) { _ = "STUB: not implemented"; return nil }

func (srv *Server) AcceptOne() { _ = "STUB: not implemented"; return }

func (srv *Server) CloseOne() bool { _ = "STUB: not implemented"; return false }

func (srv *Server) Listen() error { _ = "STUB: not implemented"; return nil }

func (srv *Server) ListenTLS() error { _ = "STUB: not implemented"; return nil }

func (srv *Server) ListenAndServe() (err error) { _ = "STUB: not implemented"; return nil }

func (srv *Server) ListenAndServeTLS(certFile, keyFile string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// RegisterListener register the instance created outside by net.Listen() in server
func (srv *Server) RegisterListener(l net.Listener) { _ = "STUB: not implemented"; return }

func (srv *Server) getOrCreateListener(addr string) (l net.Listener, err error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func (srv *Server) newListener(addr string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func (srv *Server) Shutdown() { _ = "STUB: not implemented"; return }

func (srv *Server) gracefulStop(d time.Duration) { _ = "STUB: not implemented"; return }

func (srv *Server) File() *os.File { _ = "STUB: not implemented"; return nil }
