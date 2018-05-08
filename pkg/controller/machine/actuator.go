/*
Copyright 2018 The Kubernetes Authors.

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

package machine

import (
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
)

// Actuator controls machines on a specific infrastructure. All
// methods should be idempotent unless otherwise specified.
type Actuator interface {
	// TODO mkjelland move ProvisionClusterDependencies to a cluster actuator
	// Provision infrastructure that the cluster needs before it
	// can be created
	ProvisionClusterDependencies(*clusterv1.Cluster, []*clusterv1.Machine) error
	// Create the machine.
	Create(*clusterv1.Cluster, *clusterv1.Machine) error
	// Delete the machine.
	Delete(*clusterv1.Machine) error
	// Update the machine to the provided definition.
	Update(c *clusterv1.Cluster, machine *clusterv1.Machine) error
	// Checks if the machine currently exists.
	Exists(*clusterv1.Machine) (bool, error)
}
