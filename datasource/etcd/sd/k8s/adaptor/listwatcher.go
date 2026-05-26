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

package adaptor

import (
	"context"

	pb "github.com/go-chassis/cari/discovery"
	"k8s.io/client-go/tools/cache"

	"github.com/apache/servicecomb-service-center/pkg/queue"
)

type K8sEvent struct {
	EventType  pb.EventType
	Object     interface{}
	PrevObject interface{}
}

type OnEventFunc func(evt K8sEvent)

type Watcher interface {
	queue.Worker
	OnEvent(evt K8sEvent)
}

type ListWatcher interface {
	cache.SharedIndexInformer
	queue.Worker
}

type k8sListWatcher struct {
	cache.SharedIndexInformer
	cb OnEventFunc
}

func (w *k8sListWatcher) Handle(_ context.Context, obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func NewListWatcher(t K8sType, lister cache.SharedIndexInformer, f OnEventFunc) (lw ListWatcher) {
	_ = "STUB: not implemented"
	return *new(ListWatcher)
}
