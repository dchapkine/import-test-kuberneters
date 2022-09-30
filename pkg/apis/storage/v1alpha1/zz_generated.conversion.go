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
	"unsafe"

	apicorev1 "k8s.io/api/core/v1"
	apistoragev1alpha1 "k8s.io/api/storage/v1alpha1"
	pkgapiresource "k8s.io/apimachinery/pkg/api/resource"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgconversion "k8s.io/apimachinery/pkg/conversion"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	pkgapiscore "k8s.io/kubernetes/pkg/apis/core"
	apiscorev1 "k8s.io/kubernetes/pkg/apis/core/v1"
	pkgapisstorage "k8s.io/kubernetes/pkg/apis/storage"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *apimachinerypkgruntime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*apistoragev1alpha1.CSIStorageCapacity)(nil), (*pkgapisstorage.CSIStorageCapacity)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1alpha1_CSIStorageCapacity_To_storage_CSIStorageCapacity(a.(*apistoragev1alpha1.CSIStorageCapacity), b.(*pkgapisstorage.CSIStorageCapacity), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisstorage.CSIStorageCapacity)(nil), (*apistoragev1alpha1.CSIStorageCapacity)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_storage_CSIStorageCapacity_To_v1alpha1_CSIStorageCapacity(a.(*pkgapisstorage.CSIStorageCapacity), b.(*apistoragev1alpha1.CSIStorageCapacity), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apistoragev1alpha1.CSIStorageCapacityList)(nil), (*pkgapisstorage.CSIStorageCapacityList)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1alpha1_CSIStorageCapacityList_To_storage_CSIStorageCapacityList(a.(*apistoragev1alpha1.CSIStorageCapacityList), b.(*pkgapisstorage.CSIStorageCapacityList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisstorage.CSIStorageCapacityList)(nil), (*apistoragev1alpha1.CSIStorageCapacityList)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_storage_CSIStorageCapacityList_To_v1alpha1_CSIStorageCapacityList(a.(*pkgapisstorage.CSIStorageCapacityList), b.(*apistoragev1alpha1.CSIStorageCapacityList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apistoragev1alpha1.VolumeAttachment)(nil), (*pkgapisstorage.VolumeAttachment)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1alpha1_VolumeAttachment_To_storage_VolumeAttachment(a.(*apistoragev1alpha1.VolumeAttachment), b.(*pkgapisstorage.VolumeAttachment), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisstorage.VolumeAttachment)(nil), (*apistoragev1alpha1.VolumeAttachment)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_storage_VolumeAttachment_To_v1alpha1_VolumeAttachment(a.(*pkgapisstorage.VolumeAttachment), b.(*apistoragev1alpha1.VolumeAttachment), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apistoragev1alpha1.VolumeAttachmentList)(nil), (*pkgapisstorage.VolumeAttachmentList)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1alpha1_VolumeAttachmentList_To_storage_VolumeAttachmentList(a.(*apistoragev1alpha1.VolumeAttachmentList), b.(*pkgapisstorage.VolumeAttachmentList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisstorage.VolumeAttachmentList)(nil), (*apistoragev1alpha1.VolumeAttachmentList)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_storage_VolumeAttachmentList_To_v1alpha1_VolumeAttachmentList(a.(*pkgapisstorage.VolumeAttachmentList), b.(*apistoragev1alpha1.VolumeAttachmentList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apistoragev1alpha1.VolumeAttachmentSource)(nil), (*pkgapisstorage.VolumeAttachmentSource)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1alpha1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(a.(*apistoragev1alpha1.VolumeAttachmentSource), b.(*pkgapisstorage.VolumeAttachmentSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisstorage.VolumeAttachmentSource)(nil), (*apistoragev1alpha1.VolumeAttachmentSource)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_storage_VolumeAttachmentSource_To_v1alpha1_VolumeAttachmentSource(a.(*pkgapisstorage.VolumeAttachmentSource), b.(*apistoragev1alpha1.VolumeAttachmentSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apistoragev1alpha1.VolumeAttachmentSpec)(nil), (*pkgapisstorage.VolumeAttachmentSpec)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1alpha1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(a.(*apistoragev1alpha1.VolumeAttachmentSpec), b.(*pkgapisstorage.VolumeAttachmentSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisstorage.VolumeAttachmentSpec)(nil), (*apistoragev1alpha1.VolumeAttachmentSpec)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_storage_VolumeAttachmentSpec_To_v1alpha1_VolumeAttachmentSpec(a.(*pkgapisstorage.VolumeAttachmentSpec), b.(*apistoragev1alpha1.VolumeAttachmentSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apistoragev1alpha1.VolumeAttachmentStatus)(nil), (*pkgapisstorage.VolumeAttachmentStatus)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1alpha1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(a.(*apistoragev1alpha1.VolumeAttachmentStatus), b.(*pkgapisstorage.VolumeAttachmentStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisstorage.VolumeAttachmentStatus)(nil), (*apistoragev1alpha1.VolumeAttachmentStatus)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_storage_VolumeAttachmentStatus_To_v1alpha1_VolumeAttachmentStatus(a.(*pkgapisstorage.VolumeAttachmentStatus), b.(*apistoragev1alpha1.VolumeAttachmentStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apistoragev1alpha1.VolumeError)(nil), (*pkgapisstorage.VolumeError)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1alpha1_VolumeError_To_storage_VolumeError(a.(*apistoragev1alpha1.VolumeError), b.(*pkgapisstorage.VolumeError), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisstorage.VolumeError)(nil), (*apistoragev1alpha1.VolumeError)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_storage_VolumeError_To_v1alpha1_VolumeError(a.(*pkgapisstorage.VolumeError), b.(*apistoragev1alpha1.VolumeError), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_CSIStorageCapacity_To_storage_CSIStorageCapacity(in *apistoragev1alpha1.CSIStorageCapacity, out *pkgapisstorage.CSIStorageCapacity, s apimachinerypkgconversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.NodeTopology = (*apismetav1.LabelSelector)(unsafe.Pointer(in.NodeTopology))
	out.StorageClassName = in.StorageClassName
	out.Capacity = (*pkgapiresource.Quantity)(unsafe.Pointer(in.Capacity))
	out.MaximumVolumeSize = (*pkgapiresource.Quantity)(unsafe.Pointer(in.MaximumVolumeSize))
	return nil
}

// Convert_v1alpha1_CSIStorageCapacity_To_storage_CSIStorageCapacity is an autogenerated conversion function.
func Convert_v1alpha1_CSIStorageCapacity_To_storage_CSIStorageCapacity(in *apistoragev1alpha1.CSIStorageCapacity, out *pkgapisstorage.CSIStorageCapacity, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1alpha1_CSIStorageCapacity_To_storage_CSIStorageCapacity(in, out, s)
}

func autoConvert_storage_CSIStorageCapacity_To_v1alpha1_CSIStorageCapacity(in *pkgapisstorage.CSIStorageCapacity, out *apistoragev1alpha1.CSIStorageCapacity, s apimachinerypkgconversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.NodeTopology = (*apismetav1.LabelSelector)(unsafe.Pointer(in.NodeTopology))
	out.StorageClassName = in.StorageClassName
	out.Capacity = (*pkgapiresource.Quantity)(unsafe.Pointer(in.Capacity))
	out.MaximumVolumeSize = (*pkgapiresource.Quantity)(unsafe.Pointer(in.MaximumVolumeSize))
	return nil
}

// Convert_storage_CSIStorageCapacity_To_v1alpha1_CSIStorageCapacity is an autogenerated conversion function.
func Convert_storage_CSIStorageCapacity_To_v1alpha1_CSIStorageCapacity(in *pkgapisstorage.CSIStorageCapacity, out *apistoragev1alpha1.CSIStorageCapacity, s apimachinerypkgconversion.Scope) error {
	return autoConvert_storage_CSIStorageCapacity_To_v1alpha1_CSIStorageCapacity(in, out, s)
}

func autoConvert_v1alpha1_CSIStorageCapacityList_To_storage_CSIStorageCapacityList(in *apistoragev1alpha1.CSIStorageCapacityList, out *pkgapisstorage.CSIStorageCapacityList, s apimachinerypkgconversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]pkgapisstorage.CSIStorageCapacity)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_CSIStorageCapacityList_To_storage_CSIStorageCapacityList is an autogenerated conversion function.
func Convert_v1alpha1_CSIStorageCapacityList_To_storage_CSIStorageCapacityList(in *apistoragev1alpha1.CSIStorageCapacityList, out *pkgapisstorage.CSIStorageCapacityList, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1alpha1_CSIStorageCapacityList_To_storage_CSIStorageCapacityList(in, out, s)
}

func autoConvert_storage_CSIStorageCapacityList_To_v1alpha1_CSIStorageCapacityList(in *pkgapisstorage.CSIStorageCapacityList, out *apistoragev1alpha1.CSIStorageCapacityList, s apimachinerypkgconversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]apistoragev1alpha1.CSIStorageCapacity)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_storage_CSIStorageCapacityList_To_v1alpha1_CSIStorageCapacityList is an autogenerated conversion function.
func Convert_storage_CSIStorageCapacityList_To_v1alpha1_CSIStorageCapacityList(in *pkgapisstorage.CSIStorageCapacityList, out *apistoragev1alpha1.CSIStorageCapacityList, s apimachinerypkgconversion.Scope) error {
	return autoConvert_storage_CSIStorageCapacityList_To_v1alpha1_CSIStorageCapacityList(in, out, s)
}

func autoConvert_v1alpha1_VolumeAttachment_To_storage_VolumeAttachment(in *apistoragev1alpha1.VolumeAttachment, out *pkgapisstorage.VolumeAttachment, s apimachinerypkgconversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_VolumeAttachment_To_storage_VolumeAttachment is an autogenerated conversion function.
func Convert_v1alpha1_VolumeAttachment_To_storage_VolumeAttachment(in *apistoragev1alpha1.VolumeAttachment, out *pkgapisstorage.VolumeAttachment, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1alpha1_VolumeAttachment_To_storage_VolumeAttachment(in, out, s)
}

func autoConvert_storage_VolumeAttachment_To_v1alpha1_VolumeAttachment(in *pkgapisstorage.VolumeAttachment, out *apistoragev1alpha1.VolumeAttachment, s apimachinerypkgconversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_storage_VolumeAttachmentSpec_To_v1alpha1_VolumeAttachmentSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_storage_VolumeAttachmentStatus_To_v1alpha1_VolumeAttachmentStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_storage_VolumeAttachment_To_v1alpha1_VolumeAttachment is an autogenerated conversion function.
func Convert_storage_VolumeAttachment_To_v1alpha1_VolumeAttachment(in *pkgapisstorage.VolumeAttachment, out *apistoragev1alpha1.VolumeAttachment, s apimachinerypkgconversion.Scope) error {
	return autoConvert_storage_VolumeAttachment_To_v1alpha1_VolumeAttachment(in, out, s)
}

func autoConvert_v1alpha1_VolumeAttachmentList_To_storage_VolumeAttachmentList(in *apistoragev1alpha1.VolumeAttachmentList, out *pkgapisstorage.VolumeAttachmentList, s apimachinerypkgconversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]pkgapisstorage.VolumeAttachment, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_VolumeAttachment_To_storage_VolumeAttachment(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_VolumeAttachmentList_To_storage_VolumeAttachmentList is an autogenerated conversion function.
func Convert_v1alpha1_VolumeAttachmentList_To_storage_VolumeAttachmentList(in *apistoragev1alpha1.VolumeAttachmentList, out *pkgapisstorage.VolumeAttachmentList, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1alpha1_VolumeAttachmentList_To_storage_VolumeAttachmentList(in, out, s)
}

func autoConvert_storage_VolumeAttachmentList_To_v1alpha1_VolumeAttachmentList(in *pkgapisstorage.VolumeAttachmentList, out *apistoragev1alpha1.VolumeAttachmentList, s apimachinerypkgconversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]apistoragev1alpha1.VolumeAttachment, len(*in))
		for i := range *in {
			if err := Convert_storage_VolumeAttachment_To_v1alpha1_VolumeAttachment(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_storage_VolumeAttachmentList_To_v1alpha1_VolumeAttachmentList is an autogenerated conversion function.
func Convert_storage_VolumeAttachmentList_To_v1alpha1_VolumeAttachmentList(in *pkgapisstorage.VolumeAttachmentList, out *apistoragev1alpha1.VolumeAttachmentList, s apimachinerypkgconversion.Scope) error {
	return autoConvert_storage_VolumeAttachmentList_To_v1alpha1_VolumeAttachmentList(in, out, s)
}

func autoConvert_v1alpha1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(in *apistoragev1alpha1.VolumeAttachmentSource, out *pkgapisstorage.VolumeAttachmentSource, s apimachinerypkgconversion.Scope) error {
	out.PersistentVolumeName = (*string)(unsafe.Pointer(in.PersistentVolumeName))
	if in.InlineVolumeSpec != nil {
		in, out := &in.InlineVolumeSpec, &out.InlineVolumeSpec
		*out = new(pkgapiscore.PersistentVolumeSpec)
		if err := apiscorev1.Convert_v1_PersistentVolumeSpec_To_core_PersistentVolumeSpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.InlineVolumeSpec = nil
	}
	return nil
}

// Convert_v1alpha1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource is an autogenerated conversion function.
func Convert_v1alpha1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(in *apistoragev1alpha1.VolumeAttachmentSource, out *pkgapisstorage.VolumeAttachmentSource, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1alpha1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(in, out, s)
}

func autoConvert_storage_VolumeAttachmentSource_To_v1alpha1_VolumeAttachmentSource(in *pkgapisstorage.VolumeAttachmentSource, out *apistoragev1alpha1.VolumeAttachmentSource, s apimachinerypkgconversion.Scope) error {
	out.PersistentVolumeName = (*string)(unsafe.Pointer(in.PersistentVolumeName))
	if in.InlineVolumeSpec != nil {
		in, out := &in.InlineVolumeSpec, &out.InlineVolumeSpec
		*out = new(apicorev1.PersistentVolumeSpec)
		if err := apiscorev1.Convert_core_PersistentVolumeSpec_To_v1_PersistentVolumeSpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.InlineVolumeSpec = nil
	}
	return nil
}

// Convert_storage_VolumeAttachmentSource_To_v1alpha1_VolumeAttachmentSource is an autogenerated conversion function.
func Convert_storage_VolumeAttachmentSource_To_v1alpha1_VolumeAttachmentSource(in *pkgapisstorage.VolumeAttachmentSource, out *apistoragev1alpha1.VolumeAttachmentSource, s apimachinerypkgconversion.Scope) error {
	return autoConvert_storage_VolumeAttachmentSource_To_v1alpha1_VolumeAttachmentSource(in, out, s)
}

func autoConvert_v1alpha1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(in *apistoragev1alpha1.VolumeAttachmentSpec, out *pkgapisstorage.VolumeAttachmentSpec, s apimachinerypkgconversion.Scope) error {
	out.Attacher = in.Attacher
	if err := Convert_v1alpha1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(&in.Source, &out.Source, s); err != nil {
		return err
	}
	out.NodeName = in.NodeName
	return nil
}

// Convert_v1alpha1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec is an autogenerated conversion function.
func Convert_v1alpha1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(in *apistoragev1alpha1.VolumeAttachmentSpec, out *pkgapisstorage.VolumeAttachmentSpec, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1alpha1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(in, out, s)
}

func autoConvert_storage_VolumeAttachmentSpec_To_v1alpha1_VolumeAttachmentSpec(in *pkgapisstorage.VolumeAttachmentSpec, out *apistoragev1alpha1.VolumeAttachmentSpec, s apimachinerypkgconversion.Scope) error {
	out.Attacher = in.Attacher
	if err := Convert_storage_VolumeAttachmentSource_To_v1alpha1_VolumeAttachmentSource(&in.Source, &out.Source, s); err != nil {
		return err
	}
	out.NodeName = in.NodeName
	return nil
}

// Convert_storage_VolumeAttachmentSpec_To_v1alpha1_VolumeAttachmentSpec is an autogenerated conversion function.
func Convert_storage_VolumeAttachmentSpec_To_v1alpha1_VolumeAttachmentSpec(in *pkgapisstorage.VolumeAttachmentSpec, out *apistoragev1alpha1.VolumeAttachmentSpec, s apimachinerypkgconversion.Scope) error {
	return autoConvert_storage_VolumeAttachmentSpec_To_v1alpha1_VolumeAttachmentSpec(in, out, s)
}

func autoConvert_v1alpha1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(in *apistoragev1alpha1.VolumeAttachmentStatus, out *pkgapisstorage.VolumeAttachmentStatus, s apimachinerypkgconversion.Scope) error {
	out.Attached = in.Attached
	out.AttachmentMetadata = *(*map[string]string)(unsafe.Pointer(&in.AttachmentMetadata))
	out.AttachError = (*pkgapisstorage.VolumeError)(unsafe.Pointer(in.AttachError))
	out.DetachError = (*pkgapisstorage.VolumeError)(unsafe.Pointer(in.DetachError))
	return nil
}

// Convert_v1alpha1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus is an autogenerated conversion function.
func Convert_v1alpha1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(in *apistoragev1alpha1.VolumeAttachmentStatus, out *pkgapisstorage.VolumeAttachmentStatus, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1alpha1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(in, out, s)
}

func autoConvert_storage_VolumeAttachmentStatus_To_v1alpha1_VolumeAttachmentStatus(in *pkgapisstorage.VolumeAttachmentStatus, out *apistoragev1alpha1.VolumeAttachmentStatus, s apimachinerypkgconversion.Scope) error {
	out.Attached = in.Attached
	out.AttachmentMetadata = *(*map[string]string)(unsafe.Pointer(&in.AttachmentMetadata))
	out.AttachError = (*apistoragev1alpha1.VolumeError)(unsafe.Pointer(in.AttachError))
	out.DetachError = (*apistoragev1alpha1.VolumeError)(unsafe.Pointer(in.DetachError))
	return nil
}

// Convert_storage_VolumeAttachmentStatus_To_v1alpha1_VolumeAttachmentStatus is an autogenerated conversion function.
func Convert_storage_VolumeAttachmentStatus_To_v1alpha1_VolumeAttachmentStatus(in *pkgapisstorage.VolumeAttachmentStatus, out *apistoragev1alpha1.VolumeAttachmentStatus, s apimachinerypkgconversion.Scope) error {
	return autoConvert_storage_VolumeAttachmentStatus_To_v1alpha1_VolumeAttachmentStatus(in, out, s)
}

func autoConvert_v1alpha1_VolumeError_To_storage_VolumeError(in *apistoragev1alpha1.VolumeError, out *pkgapisstorage.VolumeError, s apimachinerypkgconversion.Scope) error {
	out.Time = in.Time
	out.Message = in.Message
	return nil
}

// Convert_v1alpha1_VolumeError_To_storage_VolumeError is an autogenerated conversion function.
func Convert_v1alpha1_VolumeError_To_storage_VolumeError(in *apistoragev1alpha1.VolumeError, out *pkgapisstorage.VolumeError, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1alpha1_VolumeError_To_storage_VolumeError(in, out, s)
}

func autoConvert_storage_VolumeError_To_v1alpha1_VolumeError(in *pkgapisstorage.VolumeError, out *apistoragev1alpha1.VolumeError, s apimachinerypkgconversion.Scope) error {
	out.Time = in.Time
	out.Message = in.Message
	return nil
}

// Convert_storage_VolumeError_To_v1alpha1_VolumeError is an autogenerated conversion function.
func Convert_storage_VolumeError_To_v1alpha1_VolumeError(in *pkgapisstorage.VolumeError, out *apistoragev1alpha1.VolumeError, s apimachinerypkgconversion.Scope) error {
	return autoConvert_storage_VolumeError_To_v1alpha1_VolumeError(in, out, s)
}
