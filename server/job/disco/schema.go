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

package disco

import (
	"fmt"

	"github.com/apache/servicecomb-service-center/pkg/log"
	"github.com/apache/servicecomb-service-center/server/config"
	"github.com/robfig/cron/v3"
)

const (
	defaultRetireSchemaCron = "0 2 * * *"
	retireSchemaLockTTL     = 60
	retireSchemaLockKey     = "retire-schema-job"
)

func init() {
	cronExpr := config.GetString("registry.schema.retire.cron", defaultRetireSchemaCron)
	log.Info(fmt.Sprintf("start retire schema job, plan is %v", cronExpr))
	c := cron.New()
	_, err := c.AddFunc(cronExpr, func() {
		retireSchema()
	})
	if err != nil {
		log.Error("cron add func failed", err)
		return
	}
	c.Start()
}

func retireSchema() { _ = "STUB: not implemented"; return }
