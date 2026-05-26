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

package exception

import (
	"net/http"
)

// Writer is the async response writer, it is not thread safe!
type Writer struct {
	StatusCode int
	Body       []byte
	w          http.ResponseWriter
	flushed    bool
}

func (aw *Writer) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (aw *Writer) Write(body []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (aw *Writer) WriteHeader(statusCode int) { _ = "STUB: not implemented"; return }

func (aw *Writer) Flush() error { _ = "STUB: not implemented"; return nil }

func NewWriter(w http.ResponseWriter) *Writer { _ = "STUB: not implemented"; return nil }
