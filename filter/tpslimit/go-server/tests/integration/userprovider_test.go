// +build integration

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

package integration

import (
	"context"
	"strings"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	user := &User{}
	var err error

	// loop 50 times to make sure rejection happens
	for i := 0; i < 50; i++ {
		if user, err = userProvider.GetUser(context.TODO(), "A001"); err != nil {
			break
		}

		assert.Equal(t, "A001", user.ID)
		assert.Equal(t, "Alex Stocks", user.Name)
		assert.Equal(t, int32(18), user.Age)
		assert.NotNil(t, user.Time)
	}

	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "rejected"))
}
