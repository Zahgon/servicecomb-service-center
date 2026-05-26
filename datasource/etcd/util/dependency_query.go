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

import (
	"context"

	pb "github.com/go-chassis/cari/discovery"
)

// DependencyRelationFilterOpt contains SameDomainProject and NonSelf flag
type DependencyRelationFilterOpt struct {
	SameDomainProject bool
	NonSelf           bool
}

type DependencyRelationFilterOption func(opt DependencyRelationFilterOpt) DependencyRelationFilterOpt

func WithSameDomainProject() DependencyRelationFilterOption {
	_ = "STUB: not implemented"
	return *new(DependencyRelationFilterOption)
}

func WithoutSelfDependency() DependencyRelationFilterOption {
	_ = "STUB: not implemented"
	return *new(DependencyRelationFilterOption)
}

func ToDependencyRelationFilterOpt(opts ...DependencyRelationFilterOption) (op DependencyRelationFilterOpt) {
	_ = "STUB: not implemented"
	return *new(DependencyRelationFilterOpt)
}

type DependencyRelation struct {
	ctx           context.Context
	domainProject string
	consumer      *pb.MicroService
	provider      *pb.MicroService
}

func (dr *DependencyRelation) GetDependencyProviders(opts ...DependencyRelationFilterOption) ([]*pb.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) getDependencyProviderIds() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) getProviderKeys() ([]*pb.MicroServiceKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) GetProviderIdsByRules(providerRules []*pb.MicroServiceKey) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) parseDependencyRule(dependencyRule *pb.MicroServiceKey) (serviceIDs []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) GetDependencyConsumers(opts ...DependencyRelationFilterOption) ([]*pb.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) GetServiceByMicroServiceKey(service *pb.MicroServiceKey) (*pb.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) getDependencyConsumerIds() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) GetDependencyConsumersOfProvider() ([]*pb.MicroServiceKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DependencyRelation) GetConsumerOfSameServiceNameAndAppID(provider *pb.MicroServiceKey) ([]*pb.MicroServiceKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewProviderDependencyRelation(ctx context.Context, domainProject string, provider *pb.MicroService) *DependencyRelation {
	_ = "STUB: not implemented"
	return nil
}

func NewConsumerDependencyRelation(ctx context.Context, domainProject string, consumer *pb.MicroService) *DependencyRelation {
	_ = "STUB: not implemented"
	return nil
}

func NewDependencyRelation(ctx context.Context, domainProject string, consumer *pb.MicroService, provider *pb.MicroService) *DependencyRelation {
	_ = "STUB: not implemented"
	return nil
}
