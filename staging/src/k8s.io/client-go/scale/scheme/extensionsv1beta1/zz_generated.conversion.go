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

package extensionsv1beta1

import (
	apiextensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	apimachinerypkgconversion "k8s.io/apimachinery/pkg/conversion"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	clientgoscalescheme "k8s.io/client-go/scale/scheme"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *apimachinerypkgruntime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*apiextensionsv1beta1.Scale)(nil), (*clientgoscalescheme.Scale)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1beta1_Scale_To_scheme_Scale(a.(*apiextensionsv1beta1.Scale), b.(*clientgoscalescheme.Scale), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*clientgoscalescheme.Scale)(nil), (*apiextensionsv1beta1.Scale)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_scheme_Scale_To_v1beta1_Scale(a.(*clientgoscalescheme.Scale), b.(*apiextensionsv1beta1.Scale), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiextensionsv1beta1.ScaleSpec)(nil), (*clientgoscalescheme.ScaleSpec)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1beta1_ScaleSpec_To_scheme_ScaleSpec(a.(*apiextensionsv1beta1.ScaleSpec), b.(*clientgoscalescheme.ScaleSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*clientgoscalescheme.ScaleSpec)(nil), (*apiextensionsv1beta1.ScaleSpec)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_scheme_ScaleSpec_To_v1beta1_ScaleSpec(a.(*clientgoscalescheme.ScaleSpec), b.(*apiextensionsv1beta1.ScaleSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*clientgoscalescheme.ScaleStatus)(nil), (*apiextensionsv1beta1.ScaleStatus)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_scheme_ScaleStatus_To_v1beta1_ScaleStatus(a.(*clientgoscalescheme.ScaleStatus), b.(*apiextensionsv1beta1.ScaleStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*apiextensionsv1beta1.ScaleStatus)(nil), (*clientgoscalescheme.ScaleStatus)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1beta1_ScaleStatus_To_scheme_ScaleStatus(a.(*apiextensionsv1beta1.ScaleStatus), b.(*clientgoscalescheme.ScaleStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_Scale_To_scheme_Scale(in *apiextensionsv1beta1.Scale, out *clientgoscalescheme.Scale, s apimachinerypkgconversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ScaleSpec_To_scheme_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ScaleStatus_To_scheme_ScaleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Scale_To_scheme_Scale is an autogenerated conversion function.
func Convert_v1beta1_Scale_To_scheme_Scale(in *apiextensionsv1beta1.Scale, out *clientgoscalescheme.Scale, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1beta1_Scale_To_scheme_Scale(in, out, s)
}

func autoConvert_scheme_Scale_To_v1beta1_Scale(in *clientgoscalescheme.Scale, out *apiextensionsv1beta1.Scale, s apimachinerypkgconversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_scheme_ScaleSpec_To_v1beta1_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_scheme_ScaleStatus_To_v1beta1_ScaleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_scheme_Scale_To_v1beta1_Scale is an autogenerated conversion function.
func Convert_scheme_Scale_To_v1beta1_Scale(in *clientgoscalescheme.Scale, out *apiextensionsv1beta1.Scale, s apimachinerypkgconversion.Scope) error {
	return autoConvert_scheme_Scale_To_v1beta1_Scale(in, out, s)
}

func autoConvert_v1beta1_ScaleSpec_To_scheme_ScaleSpec(in *apiextensionsv1beta1.ScaleSpec, out *clientgoscalescheme.ScaleSpec, s apimachinerypkgconversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

// Convert_v1beta1_ScaleSpec_To_scheme_ScaleSpec is an autogenerated conversion function.
func Convert_v1beta1_ScaleSpec_To_scheme_ScaleSpec(in *apiextensionsv1beta1.ScaleSpec, out *clientgoscalescheme.ScaleSpec, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1beta1_ScaleSpec_To_scheme_ScaleSpec(in, out, s)
}

func autoConvert_scheme_ScaleSpec_To_v1beta1_ScaleSpec(in *clientgoscalescheme.ScaleSpec, out *apiextensionsv1beta1.ScaleSpec, s apimachinerypkgconversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

// Convert_scheme_ScaleSpec_To_v1beta1_ScaleSpec is an autogenerated conversion function.
func Convert_scheme_ScaleSpec_To_v1beta1_ScaleSpec(in *clientgoscalescheme.ScaleSpec, out *apiextensionsv1beta1.ScaleSpec, s apimachinerypkgconversion.Scope) error {
	return autoConvert_scheme_ScaleSpec_To_v1beta1_ScaleSpec(in, out, s)
}

func autoConvert_v1beta1_ScaleStatus_To_scheme_ScaleStatus(in *apiextensionsv1beta1.ScaleStatus, out *clientgoscalescheme.ScaleStatus, s apimachinerypkgconversion.Scope) error {
	out.Replicas = in.Replicas
	// WARNING: in.Selector requires manual conversion: inconvertible types (map[string]string vs *k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector)
	// WARNING: in.TargetSelector requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_scheme_ScaleStatus_To_v1beta1_ScaleStatus(in *clientgoscalescheme.ScaleStatus, out *apiextensionsv1beta1.ScaleStatus, s apimachinerypkgconversion.Scope) error {
	out.Replicas = in.Replicas
	// WARNING: in.Selector requires manual conversion: inconvertible types (*k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector vs map[string]string)
	return nil
}
