/*
Copyright 2025 The KubeStellar Authors.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "integrations.kubestellar.io"

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
	GroupVersion  = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(GroupVersion,
		&ArgoCDBinding{},
		&ArgoCDBindingList{},
	)
	metav1.AddToGroupVersion(scheme, GroupVersion)
	return nil
}
