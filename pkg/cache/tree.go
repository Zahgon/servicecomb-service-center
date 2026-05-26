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

package cache

import (
	"context"
	"errors"
	"sync"

	"github.com/karlseguin/ccache"
)

var errNilNode = errors.New("nil node")

type Tree struct {
	Config  *Config
	roots   *ccache.Cache
	filters []Filter
	lock    sync.RWMutex
}

func (t *Tree) AddFilter(fs ...Filter) *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) Get(ctx context.Context, ops ...Option) (node *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parent may be a temp root in concurrent scene

func (t *Tree) Remove(ctx context.Context) { _ = "STUB: not implemented"; return }

func (t *Tree) getOrCreateRoot(ctx context.Context) (node *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Tree) getOrCreateNode(ctx context.Context, idx int, parent *Node) (node *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// new a temp node

func (t *Tree) nodeFullName(name string, parent *Node) string { _ = "STUB: not implemented"; return "" }

func (t *Tree) createNode(ctx context.Context, idx int, name string, parent *Node) (node *Node, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTree(cfg *Config) *Tree { _ = "STUB: not implemented"; return nil }
