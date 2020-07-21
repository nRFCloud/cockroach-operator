/*
Copyright 2020 The Cockroach Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package condition

import (
	api "github.com/cockroachdb/cockroach-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestInitializesForEmptyConditions(t *testing.T) {
	now := metav1.Now()

	status := api.CrdbClusterStatus{}

	expected := []api.ClusterCondition{
		{
			Type:               "NotInitialized",
			Status:             metav1.ConditionTrue,
			LastTransitionTime: now,
		},
	}

	InitConditionsIfNeeded(&status, now)

	assert.ElementsMatch(t, expected, status.Conditions)
}