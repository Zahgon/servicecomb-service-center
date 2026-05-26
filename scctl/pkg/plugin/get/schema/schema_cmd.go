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

package schema

import (
	"github.com/apache/servicecomb-service-center/pkg/dump"
	"github.com/apache/servicecomb-service-center/scctl/pkg/plugin/get"
	"github.com/spf13/cobra"
)

var (
	AppID       string
	ServiceName string
	Version     string
	SaveDir     string
)

func init() {
	NewSchemaCommand(get.RootCmd)
}

func NewSchemaCommand(parent *cobra.Command) *cobra.Command { _ = "STUB: not implemented"; return nil }

// schemas/[${domain}/][${project}/][${env}/]${app}/${microservice}.${version}/${schemaId}.yaml
func saveDirectory(root string, ms *dump.Microservice) string { _ = "STUB: not implemented"; return "" }

func CommandFunc(_ *cobra.Command, _ []string) { _ = "STUB: not implemented"; return }
