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
	"os"
	"sync"

	"github.com/apache/servicecomb-service-center/pkg/goutil"
	"github.com/apache/servicecomb-service-center/pkg/log"
	"github.com/apache/servicecomb-service-center/pkg/util"
	"github.com/go-chassis/foundation/gopool"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

var (
	client     *K8sClient
	clientOnce sync.Once
)

type K8sType string

type K8sClient struct {
	// eventFuncs is used for store functions called k8s event handler
	eventFuncs util.ConcurrentMap
	// ipIndex is used for index pod object; key is ip, value is pod full name
	ipIndex util.ConcurrentMap

	kubeClient kubernetes.Interface
	services   ListWatcher
	endpoints  ListWatcher
	nodes      ListWatcher
	pods       ListWatcher

	ready     chan struct{}
	stopCh    chan struct{}
	goroutine *gopool.Pool
}

func (c *K8sClient) init() (err error) {
	c.ready = make(chan struct{})
	c.stopCh = make(chan struct{})
	c.goroutine = goutil.New()

	// if KUBERNETES_CONFIG_PATH is unset, then service center must be deployed in the same k8s cluster
	c.kubeClient, err = createKubeClient(os.Getenv("KUBERNETES_CONFIG_PATH"))
	if err != nil {
		log.Error("create kube client failed", err)
		return
	}

	// if KUBERNETES_NAMESPACE is unset, then list watch all namespaces
	listerFactory := informers.NewFilteredSharedInformerFactory(
		c.kubeClient, defaultResyncInterval, os.Getenv("KUBERNETES_NAMESPACE"), nil)
	c.services = c.newListWatcher(TypeService, listerFactory.Core().V1().Services().Informer())
	c.endpoints = c.newListWatcher(TypeEndpoint, listerFactory.Core().V1().Endpoints().Informer())
	c.nodes = c.newListWatcher(TypeNode, listerFactory.Core().V1().Nodes().Informer())
	c.pods = c.newListWatcher(TypePod, listerFactory.Core().V1().Pods().Informer())

	// append ipIndex build function
	c.AppendEventFunc(TypePod, c.onPodEvent)
	return
}

func (c *K8sClient) newListWatcher(t K8sType, lister cache.SharedIndexInformer) (lw ListWatcher) {
	_ = "STUB: not implemented"
	return *new(ListWatcher)
}

func (c *K8sClient) getEvent(t K8sType) OnEventFunc {
	_ = "STUB: not implemented"
	return *new(OnEventFunc)
}

// onPodEvent is method to build ipIndex
func (c *K8sClient) onPodEvent(evt K8sEvent) { _ = "STUB: not implemented"; return }

// unsafe
func (c *K8sClient) AppendEventFunc(t K8sType, f OnEventFunc) { _ = "STUB: not implemented"; return }

func (c *K8sClient) waitForSync(lw ListWatcher) ListWatcher {
	_ = "STUB: not implemented"
	return *new(ListWatcher)
}

func (c *K8sClient) Services() ListWatcher { _ = "STUB: not implemented"; return *new(ListWatcher) }

func (c *K8sClient) Endpoints() ListWatcher { _ = "STUB: not implemented"; return *new(ListWatcher) }

func (c *K8sClient) Pods() ListWatcher { _ = "STUB: not implemented"; return *new(ListWatcher) }

func (c *K8sClient) Nodes() ListWatcher { _ = "STUB: not implemented"; return *new(ListWatcher) }

func (c *K8sClient) GetDomainProject() string { _ = "STUB: not implemented"; return "" }

func (c *K8sClient) GetService(namespace, name string) (svc *v1.Service) {
	_ = "STUB: not implemented"
	return nil
}

func (c *K8sClient) GetEndpoints(namespace, name string) (ep *v1.Endpoints) {
	_ = "STUB: not implemented"
	return nil
}

func (c *K8sClient) GetPodByIP(ip string) (pod *v1.Pod) { _ = "STUB: not implemented"; return nil }

func (c *K8sClient) GetNodeByPod(pod *v1.Pod) (node *v1.Node) {
	_ = "STUB: not implemented"
	return nil
}

func (c *K8sClient) Run() { _ = "STUB: not implemented"; return }

func (c *K8sClient) Stop() { _ = "STUB: not implemented"; return }

func (c *K8sClient) Ready() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func Kubernetes() *K8sClient { _ = "STUB: not implemented"; return nil }
