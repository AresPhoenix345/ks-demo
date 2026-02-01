/*
Copyright 2025 The KubeStellar Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ArgoCDBindingSpec defines the desired state of ArgoCDBinding.
// It references a KubeStellar BindingPolicy and provides a template for Argo CD Applications.
type ArgoCDBindingSpec struct {
	// BindingPolicyRef references a KubeStellar BindingPolicy (cluster-scoped).
	BindingPolicyRef LocalObjectReference `json:"bindingPolicyRef"`

	// ApplicationTemplate defines the Argo CD Application template (project, source, syncPolicy).
	// Destination is derived per cluster from the BindingPolicy/Binding resolution.
	ApplicationTemplate ApplicationTemplate `json:"applicationTemplate"`

	// SyncPolicy defines sync behavior (e.g. automated prune/selfHeal).
	// +optional
	SyncPolicy *SyncPolicy `json:"syncPolicy,omitempty"`
}

// LocalObjectReference references an object by name (cluster-scoped or namespaced).
type LocalObjectReference struct {
	// Name of the referenced object.
	Name string `json:"name"`
	// Namespace of the referenced object (optional; omit for cluster-scoped).
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// ApplicationTemplate is the template for creating Argo CD Application resources.
type ApplicationTemplate struct {
	// Project is the Argo CD project name.
	// +optional
	Project string `json:"project,omitempty"`

	// Source defines the Git repo and path (and optionally chart for Helm).
	Source ApplicationSource `json:"source"`

	// Destination is the base destination; server/name may be overridden per cluster.
	// +optional
	Destination *ApplicationDestination `json:"destination,omitempty"`
}

// ApplicationSource mirrors Argo CD ApplicationSource (simplified).
type ApplicationSource struct {
	RepoURL        string `json:"repoURL"`
	Path           string `json:"path,omitempty"`
	TargetRevision string `json:"targetRevision,omitempty"`
}

// ApplicationDestination mirrors Argo CD ApplicationDestination.
type ApplicationDestination struct {
	Server    string `json:"server,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
}

// SyncPolicy defines automated sync behavior.
type SyncPolicy struct {
	Automated *SyncPolicyAutomated `json:"automated,omitempty"`
}

// SyncPolicyAutomated enables automated sync.
type SyncPolicyAutomated struct {
	Prune    bool `json:"prune,omitempty"`
	SelfHeal bool `json:"selfHeal,omitempty"`
}

// ArgoCDBindingStatus defines the observed state of ArgoCDBinding.
type ArgoCDBindingStatus struct {
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	ObservedGeneration int64 `json:"observedGeneration"`

	// ApplicationCount is the number of Argo CD Applications created/updated.
	// +optional
	ApplicationCount int32 `json:"applicationCount,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,shortName=acdb

// ArgoCDBinding binds a KubeStellar BindingPolicy to Argo CD Applications (one per destination).
type ArgoCDBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArgoCDBindingSpec   `json:"spec,omitempty"`
	Status ArgoCDBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ArgoCDBindingList contains a list of ArgoCDBinding.
type ArgoCDBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items          []ArgoCDBinding `json:"items"`
}
