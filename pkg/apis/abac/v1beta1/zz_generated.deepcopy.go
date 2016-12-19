// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1beta1

import (
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_Policy, InType: reflect.TypeOf(&Policy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_PolicySpec, InType: reflect.TypeOf(&PolicySpec{})},
	)
}

func DeepCopy_v1beta1_Policy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Policy)
		out := out.(*Policy)
		out.TypeMeta = in.TypeMeta
		out.Spec = in.Spec
		return nil
	}
}

func DeepCopy_v1beta1_PolicySpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicySpec)
		out := out.(*PolicySpec)
		out.User = in.User
		out.Group = in.Group
		out.Readonly = in.Readonly
		out.APIGroup = in.APIGroup
		out.Resource = in.Resource
		out.Namespace = in.Namespace
		out.NonResourcePath = in.NonResourcePath
		return nil
	}
}
