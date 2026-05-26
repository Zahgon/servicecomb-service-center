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

package diagnose

import (
	"bytes"
	"context"

	"github.com/apache/servicecomb-service-center/pkg/dump"
	"github.com/spf13/cobra"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	service  = "service"
	instance = "instance"
)

const (
	greater = iota
	mismatch
	less
)

var typeMap = map[string]string{
	service:  "/cse-sr/ms/files/",
	instance: "/cse-sr/inst/files/",
}

type etcdResponse map[string][]*mvccpb.KeyValue

func CommandFunc(_ *cobra.Command, _ []string) {
	_ = "STUB: not implemented"
	// initialize sc/etcd clients
	return
}

// query etcd

// query sc

// diagnose go...

// stdout
// stderr

func getEtcdResponse(ctx context.Context, etcdClient *clientv3.Client) (etcdResponse, error) {
	_ = "STUB: not implemented"
	return *new(etcdResponse), nil
}

func setResponse(ctx context.Context, etcdClient *clientv3.Client, key, prefix string, etcdResp etcdResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func diagnose(cache *dump.Cache, etcdResp etcdResponse) (details string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func writeResult(b *bytes.Buffer, full *bytes.Buffer, rss ...*CompareResult) {
	_ = "STUB: not implemented"
	return
}

func writeBody(b *bytes.Buffer, r map[string][]string) { _ = "STUB: not implemented"; return }

func writeSection(b *bytes.Buffer, t string) { _ = "STUB: not implemented"; return }
