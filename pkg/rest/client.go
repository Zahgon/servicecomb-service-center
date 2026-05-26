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
	"compress/gzip"
	"crypto/tls"
	"io"
	"net/http"
	"time"

	"context"
)

var defaultURLClientOption = URLClientOption{
	Compressed:            true,
	VerifyPeer:            true,
	SSLVersion:            tls.VersionTLS12,
	HandshakeTimeout:      10 * time.Second,
	ResponseHeaderTimeout: 30 * time.Second,
	RequestTimeout:        60 * time.Second,
	ConnsPerHost:          DefaultConnPoolPerHostSize,
}

type URLClientOption struct {
	SSLEnabled            bool
	Compressed            bool
	VerifyPeer            bool
	CAFile                string
	CertFile              string
	CertKeyFile           string
	CertKeyPWD            string
	SSLVersion            uint16
	HandshakeTimeout      time.Duration
	ResponseHeaderTimeout time.Duration
	RequestTimeout        time.Duration
	ConnsPerHost          int
}

type gzipBodyReader struct {
	*gzip.Reader
	Body io.ReadCloser
}

func (w *gzipBodyReader) Close() error { _ = "STUB: not implemented"; return nil }

func NewGZipBodyReader(body io.ReadCloser) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

type URLClient struct {
	*http.Client

	TLS *tls.Config

	Cfg URLClientOption
}

func (client *URLClient) HTTPDoWithContext(ctx context.Context, method string, rawURL string, headers http.Header, body []byte) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DumpRequestOut(req *http.Request) { _ = "STUB: not implemented"; return }

func DumpResponse(resp *http.Response) { _ = "STUB: not implemented"; return }

func (client *URLClient) HTTPDo(method string, rawURL string, headers http.Header, body []byte) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DefaultURLClientOption() URLClientOption {
	_ = "STUB: not implemented"
	return *new(URLClientOption)
}

func setOptionDefaultValue(o *URLClientOption) URLClientOption {
	_ = "STUB: not implemented"
	return *new(URLClientOption)
}

func GetURLClient(o URLClientOption) (client *URLClient, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
