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

package util

// The Tree is binary sort Tree
type Tree struct {
	root        *Node
	isAddToLeft func(node *Node, addRes interface{}) bool
}

func NewTree(isAddToLeft func(node *Node, addRes interface{}) bool) *Tree {
	_ = "STUB: not implemented"
	return nil
}

type Node struct {
	Res         interface{}
	left, right *Node
}

func (t *Tree) GetRoot() *Node {
	_ = "STUB: not implemented"

	// add res into Tree
	return nil
}

func (t *Tree) AddNode(res interface{}) *Node { _ = "STUB: not implemented"; return nil }

func (t *Tree) addNode(n *Node, res interface{}) *Node { _ = "STUB: not implemented"; return nil }

// middle oder traversal, handle is the func that deals with the res, n is the start node to traversal
func (t *Tree) InOrderTraversal(n *Node, handle func(res interface{}) error) error {
	_ = "STUB: not implemented"
	return nil
}

//todo add asynchronous handle handle func: go handle
