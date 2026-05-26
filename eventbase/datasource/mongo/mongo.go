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

package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/apache/servicecomb-service-center/eventbase/datasource"
)

func init() {
	datasource.RegisterPlugin("mongo", NewDatasource)
}

type Datasource struct {
	taskDao      datasource.TaskDao
	tombstoneDao datasource.TombstoneDao
}

func (d *Datasource) TaskDao() datasource.TaskDao {
	_ = "STUB: not implemented"
	return *new(datasource.TaskDao)
}

func (d *Datasource) TombstoneDao() datasource.TombstoneDao {
	_ = "STUB: not implemented"
	return *new(datasource.TombstoneDao)
}

func NewDatasource() datasource.DataSource {
	_ = "STUB: not implemented"
	return *new(datasource.DataSource)
}

func ensureDB() { _ = "STUB: not implemented"; return }

func ensureTask() { _ = "STUB: not implemented"; return }

func ensureTombstone() { _ = "STUB: not implemented"; return }

func buildIndexDoc(keys ...string) mongo.IndexModel {
	_ = "STUB: not implemented"
	return *new(mongo.IndexModel)
}
