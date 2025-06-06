/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hierarchy

import kueue "sigs.k8s.io/kueue/apis/kueue/v1beta1"

type ClusterQueue[C nodeBase[kueue.CohortReference]] struct {
	cohort C
}

func (c *ClusterQueue[C]) Parent() C {
	return c.cohort
}

func (c *ClusterQueue[C]) HasParent() bool {
	var zero C
	return c.Parent() != zero
}

// implements clusterQueueNode interface

func (c *ClusterQueue[C]) setParent(cohort C) {
	c.cohort = cohort
}
