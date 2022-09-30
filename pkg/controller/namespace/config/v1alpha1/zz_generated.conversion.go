//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgconversion "k8s.io/apimachinery/pkg/conversion"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	kubecontrollermanagerconfigv1alpha1 "k8s.io/kube-controller-manager/config/v1alpha1"
	controllernamespaceconfig "k8s.io/kubernetes/pkg/controller/namespace/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *apimachinerypkgruntime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*kubecontrollermanagerconfigv1alpha1.GroupResource)(nil), (*apismetav1.GroupResource)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1alpha1_GroupResource_To_v1_GroupResource(a.(*kubecontrollermanagerconfigv1alpha1.GroupResource), b.(*apismetav1.GroupResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apismetav1.GroupResource)(nil), (*kubecontrollermanagerconfigv1alpha1.GroupResource)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1_GroupResource_To_v1alpha1_GroupResource(a.(*apismetav1.GroupResource), b.(*kubecontrollermanagerconfigv1alpha1.GroupResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*controllernamespaceconfig.NamespaceControllerConfiguration)(nil), (*kubecontrollermanagerconfigv1alpha1.NamespaceControllerConfiguration)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_config_NamespaceControllerConfiguration_To_v1alpha1_NamespaceControllerConfiguration(a.(*controllernamespaceconfig.NamespaceControllerConfiguration), b.(*kubecontrollermanagerconfigv1alpha1.NamespaceControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*kubecontrollermanagerconfigv1alpha1.NamespaceControllerConfiguration)(nil), (*controllernamespaceconfig.NamespaceControllerConfiguration)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1alpha1_NamespaceControllerConfiguration_To_config_NamespaceControllerConfiguration(a.(*kubecontrollermanagerconfigv1alpha1.NamespaceControllerConfiguration), b.(*controllernamespaceconfig.NamespaceControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_GroupResource_To_v1_GroupResource(in *kubecontrollermanagerconfigv1alpha1.GroupResource, out *apismetav1.GroupResource, s apimachinerypkgconversion.Scope) error {
	out.Group = in.Group
	out.Resource = in.Resource
	return nil
}

// Convert_v1alpha1_GroupResource_To_v1_GroupResource is an autogenerated conversion function.
func Convert_v1alpha1_GroupResource_To_v1_GroupResource(in *kubecontrollermanagerconfigv1alpha1.GroupResource, out *apismetav1.GroupResource, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1alpha1_GroupResource_To_v1_GroupResource(in, out, s)
}

func autoConvert_v1_GroupResource_To_v1alpha1_GroupResource(in *apismetav1.GroupResource, out *kubecontrollermanagerconfigv1alpha1.GroupResource, s apimachinerypkgconversion.Scope) error {
	out.Group = in.Group
	out.Resource = in.Resource
	return nil
}

// Convert_v1_GroupResource_To_v1alpha1_GroupResource is an autogenerated conversion function.
func Convert_v1_GroupResource_To_v1alpha1_GroupResource(in *apismetav1.GroupResource, out *kubecontrollermanagerconfigv1alpha1.GroupResource, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1_GroupResource_To_v1alpha1_GroupResource(in, out, s)
}

func autoConvert_v1alpha1_NamespaceControllerConfiguration_To_config_NamespaceControllerConfiguration(in *kubecontrollermanagerconfigv1alpha1.NamespaceControllerConfiguration, out *controllernamespaceconfig.NamespaceControllerConfiguration, s apimachinerypkgconversion.Scope) error {
	out.NamespaceSyncPeriod = in.NamespaceSyncPeriod
	out.ConcurrentNamespaceSyncs = in.ConcurrentNamespaceSyncs
	return nil
}

func autoConvert_config_NamespaceControllerConfiguration_To_v1alpha1_NamespaceControllerConfiguration(in *controllernamespaceconfig.NamespaceControllerConfiguration, out *kubecontrollermanagerconfigv1alpha1.NamespaceControllerConfiguration, s apimachinerypkgconversion.Scope) error {
	out.NamespaceSyncPeriod = in.NamespaceSyncPeriod
	out.ConcurrentNamespaceSyncs = in.ConcurrentNamespaceSyncs
	return nil
}
