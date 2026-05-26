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

package datasource

type TaskFindOptions struct {
	Domain       string
	Project      string
	Action       string
	Status       string
	ResourceType string
}

type TombstoneFindOptions struct {
	Domain          string
	Project         string
	ResourceType    string
	BeforeTimestamp int64
}

type TaskFindOption func(options *TaskFindOptions)

type TombstoneFindOption func(options *TombstoneFindOptions)

func NewTaskFindOptions() TaskFindOptions { _ = "STUB: not implemented"; return *new(TaskFindOptions) }

func NewTombstoneFindOptions() TombstoneFindOptions {
	_ = "STUB: not implemented"
	return *new(TombstoneFindOptions)
}

// WithDomain find task with domain
func WithDomain(domain string) TaskFindOption {
	_ = "STUB: not implemented"
	return *new(TaskFindOption)
}

// WithProject find task with project
func WithProject(project string) TaskFindOption {
	_ = "STUB: not implemented"
	return *new(TaskFindOption)
}

// WithAction find task with action
func WithAction(action string) TaskFindOption {
	_ = "STUB: not implemented"
	return *new(TaskFindOption)
}

// WithStatus find task with status
func WithStatus(status string) TaskFindOption {
	_ = "STUB: not implemented"
	return *new(TaskFindOption)
}

// WithDataType find task with dataType
func WithDataType(dataType string) TaskFindOption {
	_ = "STUB: not implemented"
	return *new(TaskFindOption)
}

// WithTombstoneDomain find tombstone with domain
func WithTombstoneDomain(domain string) TombstoneFindOption {
	_ = "STUB: not implemented"
	return *new(TombstoneFindOption)
}

// WithTombstoneProject find tombstone with project
func WithTombstoneProject(project string) TombstoneFindOption {
	_ = "STUB: not implemented"
	return *new(TombstoneFindOption)
}

// WithResourceType find tombstone with resource type
func WithResourceType(resourceType string) TombstoneFindOption {
	_ = "STUB: not implemented"
	return *new(TombstoneFindOption)
}

// WithBeforeTimestamp find tombstone with beforeTimestamp
func WithBeforeTimestamp(timestamp int64) TombstoneFindOption {
	_ = "STUB: not implemented"
	return *new(TombstoneFindOption)
}
