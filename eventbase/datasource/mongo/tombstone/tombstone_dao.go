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

package tombstone

import (
	"context"

	"github.com/go-chassis/cari/sync"

	"github.com/apache/servicecomb-service-center/eventbase/datasource"
	emodel "github.com/apache/servicecomb-service-center/eventbase/model"
)

type Dao struct {
}

func (d *Dao) Get(ctx context.Context, req *emodel.GetTombstoneRequest) (*sync.Tombstone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Dao) Create(ctx context.Context, tombstone *sync.Tombstone) (*sync.Tombstone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Dao) Delete(ctx context.Context, tombstones ...*sync.Tombstone) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dao) List(ctx context.Context, options ...datasource.TombstoneFindOption) ([]*sync.Tombstone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
