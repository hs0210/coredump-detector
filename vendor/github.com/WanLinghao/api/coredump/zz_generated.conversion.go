// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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
	unsafe "unsafe"

	coredump "github.com/WanLinghao/coredump-detector/pkg/apis/coredump"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*CoredumpEndpoint)(nil), (*coredump.CoredumpEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CoredumpEndpoint_To_coredump_CoredumpEndpoint(a.(*CoredumpEndpoint), b.(*coredump.CoredumpEndpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*coredump.CoredumpEndpoint)(nil), (*CoredumpEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_coredump_CoredumpEndpoint_To_v1alpha1_CoredumpEndpoint(a.(*coredump.CoredumpEndpoint), b.(*CoredumpEndpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CoredumpEndpointDump)(nil), (*coredump.CoredumpEndpointDump)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CoredumpEndpointDump_To_coredump_CoredumpEndpointDump(a.(*CoredumpEndpointDump), b.(*coredump.CoredumpEndpointDump), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*coredump.CoredumpEndpointDump)(nil), (*CoredumpEndpointDump)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_coredump_CoredumpEndpointDump_To_v1alpha1_CoredumpEndpointDump(a.(*coredump.CoredumpEndpointDump), b.(*CoredumpEndpointDump), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CoredumpEndpointDumpList)(nil), (*coredump.CoredumpEndpointDumpList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CoredumpEndpointDumpList_To_coredump_CoredumpEndpointDumpList(a.(*CoredumpEndpointDumpList), b.(*coredump.CoredumpEndpointDumpList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*coredump.CoredumpEndpointDumpList)(nil), (*CoredumpEndpointDumpList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_coredump_CoredumpEndpointDumpList_To_v1alpha1_CoredumpEndpointDumpList(a.(*coredump.CoredumpEndpointDumpList), b.(*CoredumpEndpointDumpList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CoredumpEndpointList)(nil), (*coredump.CoredumpEndpointList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CoredumpEndpointList_To_coredump_CoredumpEndpointList(a.(*CoredumpEndpointList), b.(*coredump.CoredumpEndpointList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*coredump.CoredumpEndpointList)(nil), (*CoredumpEndpointList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_coredump_CoredumpEndpointList_To_v1alpha1_CoredumpEndpointList(a.(*coredump.CoredumpEndpointList), b.(*CoredumpEndpointList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CoredumpEndpointSpec)(nil), (*coredump.CoredumpEndpointSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CoredumpEndpointSpec_To_coredump_CoredumpEndpointSpec(a.(*CoredumpEndpointSpec), b.(*coredump.CoredumpEndpointSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*coredump.CoredumpEndpointSpec)(nil), (*CoredumpEndpointSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_coredump_CoredumpEndpointSpec_To_v1alpha1_CoredumpEndpointSpec(a.(*coredump.CoredumpEndpointSpec), b.(*CoredumpEndpointSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CoredumpEndpointStatus)(nil), (*coredump.CoredumpEndpointStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CoredumpEndpointStatus_To_coredump_CoredumpEndpointStatus(a.(*CoredumpEndpointStatus), b.(*coredump.CoredumpEndpointStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*coredump.CoredumpEndpointStatus)(nil), (*CoredumpEndpointStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_coredump_CoredumpEndpointStatus_To_v1alpha1_CoredumpEndpointStatus(a.(*coredump.CoredumpEndpointStatus), b.(*CoredumpEndpointStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CoredumpGetOptions)(nil), (*coredump.CoredumpGetOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CoredumpGetOptions_To_coredump_CoredumpGetOptions(a.(*CoredumpGetOptions), b.(*coredump.CoredumpGetOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*coredump.CoredumpGetOptions)(nil), (*CoredumpGetOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_coredump_CoredumpGetOptions_To_v1alpha1_CoredumpGetOptions(a.(*coredump.CoredumpGetOptions), b.(*CoredumpGetOptions), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_CoredumpEndpoint_To_coredump_CoredumpEndpoint(in *CoredumpEndpoint, out *coredump.CoredumpEndpoint, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_CoredumpEndpointSpec_To_coredump_CoredumpEndpointSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_CoredumpEndpointStatus_To_coredump_CoredumpEndpointStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_CoredumpEndpoint_To_coredump_CoredumpEndpoint is an autogenerated conversion function.
func Convert_v1alpha1_CoredumpEndpoint_To_coredump_CoredumpEndpoint(in *CoredumpEndpoint, out *coredump.CoredumpEndpoint, s conversion.Scope) error {
	return autoConvert_v1alpha1_CoredumpEndpoint_To_coredump_CoredumpEndpoint(in, out, s)
}

func autoConvert_coredump_CoredumpEndpoint_To_v1alpha1_CoredumpEndpoint(in *coredump.CoredumpEndpoint, out *CoredumpEndpoint, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_coredump_CoredumpEndpointSpec_To_v1alpha1_CoredumpEndpointSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_coredump_CoredumpEndpointStatus_To_v1alpha1_CoredumpEndpointStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_coredump_CoredumpEndpoint_To_v1alpha1_CoredumpEndpoint is an autogenerated conversion function.
func Convert_coredump_CoredumpEndpoint_To_v1alpha1_CoredumpEndpoint(in *coredump.CoredumpEndpoint, out *CoredumpEndpoint, s conversion.Scope) error {
	return autoConvert_coredump_CoredumpEndpoint_To_v1alpha1_CoredumpEndpoint(in, out, s)
}

func autoConvert_v1alpha1_CoredumpEndpointDump_To_coredump_CoredumpEndpointDump(in *CoredumpEndpointDump, out *coredump.CoredumpEndpointDump, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	return nil
}

// Convert_v1alpha1_CoredumpEndpointDump_To_coredump_CoredumpEndpointDump is an autogenerated conversion function.
func Convert_v1alpha1_CoredumpEndpointDump_To_coredump_CoredumpEndpointDump(in *CoredumpEndpointDump, out *coredump.CoredumpEndpointDump, s conversion.Scope) error {
	return autoConvert_v1alpha1_CoredumpEndpointDump_To_coredump_CoredumpEndpointDump(in, out, s)
}

func autoConvert_coredump_CoredumpEndpointDump_To_v1alpha1_CoredumpEndpointDump(in *coredump.CoredumpEndpointDump, out *CoredumpEndpointDump, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	return nil
}

// Convert_coredump_CoredumpEndpointDump_To_v1alpha1_CoredumpEndpointDump is an autogenerated conversion function.
func Convert_coredump_CoredumpEndpointDump_To_v1alpha1_CoredumpEndpointDump(in *coredump.CoredumpEndpointDump, out *CoredumpEndpointDump, s conversion.Scope) error {
	return autoConvert_coredump_CoredumpEndpointDump_To_v1alpha1_CoredumpEndpointDump(in, out, s)
}

func autoConvert_v1alpha1_CoredumpEndpointDumpList_To_coredump_CoredumpEndpointDumpList(in *CoredumpEndpointDumpList, out *coredump.CoredumpEndpointDumpList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]coredump.CoredumpEndpointDump)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_CoredumpEndpointDumpList_To_coredump_CoredumpEndpointDumpList is an autogenerated conversion function.
func Convert_v1alpha1_CoredumpEndpointDumpList_To_coredump_CoredumpEndpointDumpList(in *CoredumpEndpointDumpList, out *coredump.CoredumpEndpointDumpList, s conversion.Scope) error {
	return autoConvert_v1alpha1_CoredumpEndpointDumpList_To_coredump_CoredumpEndpointDumpList(in, out, s)
}

func autoConvert_coredump_CoredumpEndpointDumpList_To_v1alpha1_CoredumpEndpointDumpList(in *coredump.CoredumpEndpointDumpList, out *CoredumpEndpointDumpList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]CoredumpEndpointDump)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_coredump_CoredumpEndpointDumpList_To_v1alpha1_CoredumpEndpointDumpList is an autogenerated conversion function.
func Convert_coredump_CoredumpEndpointDumpList_To_v1alpha1_CoredumpEndpointDumpList(in *coredump.CoredumpEndpointDumpList, out *CoredumpEndpointDumpList, s conversion.Scope) error {
	return autoConvert_coredump_CoredumpEndpointDumpList_To_v1alpha1_CoredumpEndpointDumpList(in, out, s)
}

func autoConvert_v1alpha1_CoredumpEndpointList_To_coredump_CoredumpEndpointList(in *CoredumpEndpointList, out *coredump.CoredumpEndpointList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]coredump.CoredumpEndpoint)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_CoredumpEndpointList_To_coredump_CoredumpEndpointList is an autogenerated conversion function.
func Convert_v1alpha1_CoredumpEndpointList_To_coredump_CoredumpEndpointList(in *CoredumpEndpointList, out *coredump.CoredumpEndpointList, s conversion.Scope) error {
	return autoConvert_v1alpha1_CoredumpEndpointList_To_coredump_CoredumpEndpointList(in, out, s)
}

func autoConvert_coredump_CoredumpEndpointList_To_v1alpha1_CoredumpEndpointList(in *coredump.CoredumpEndpointList, out *CoredumpEndpointList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]CoredumpEndpoint)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_coredump_CoredumpEndpointList_To_v1alpha1_CoredumpEndpointList is an autogenerated conversion function.
func Convert_coredump_CoredumpEndpointList_To_v1alpha1_CoredumpEndpointList(in *coredump.CoredumpEndpointList, out *CoredumpEndpointList, s conversion.Scope) error {
	return autoConvert_coredump_CoredumpEndpointList_To_v1alpha1_CoredumpEndpointList(in, out, s)
}

func autoConvert_v1alpha1_CoredumpEndpointSpec_To_coredump_CoredumpEndpointSpec(in *CoredumpEndpointSpec, out *coredump.CoredumpEndpointSpec, s conversion.Scope) error {
	out.PodUID = types.UID(in.PodUID)
	return nil
}

// Convert_v1alpha1_CoredumpEndpointSpec_To_coredump_CoredumpEndpointSpec is an autogenerated conversion function.
func Convert_v1alpha1_CoredumpEndpointSpec_To_coredump_CoredumpEndpointSpec(in *CoredumpEndpointSpec, out *coredump.CoredumpEndpointSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_CoredumpEndpointSpec_To_coredump_CoredumpEndpointSpec(in, out, s)
}

func autoConvert_coredump_CoredumpEndpointSpec_To_v1alpha1_CoredumpEndpointSpec(in *coredump.CoredumpEndpointSpec, out *CoredumpEndpointSpec, s conversion.Scope) error {
	out.PodUID = types.UID(in.PodUID)
	return nil
}

// Convert_coredump_CoredumpEndpointSpec_To_v1alpha1_CoredumpEndpointSpec is an autogenerated conversion function.
func Convert_coredump_CoredumpEndpointSpec_To_v1alpha1_CoredumpEndpointSpec(in *coredump.CoredumpEndpointSpec, out *CoredumpEndpointSpec, s conversion.Scope) error {
	return autoConvert_coredump_CoredumpEndpointSpec_To_v1alpha1_CoredumpEndpointSpec(in, out, s)
}

func autoConvert_v1alpha1_CoredumpEndpointStatus_To_coredump_CoredumpEndpointStatus(in *CoredumpEndpointStatus, out *coredump.CoredumpEndpointStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_CoredumpEndpointStatus_To_coredump_CoredumpEndpointStatus is an autogenerated conversion function.
func Convert_v1alpha1_CoredumpEndpointStatus_To_coredump_CoredumpEndpointStatus(in *CoredumpEndpointStatus, out *coredump.CoredumpEndpointStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_CoredumpEndpointStatus_To_coredump_CoredumpEndpointStatus(in, out, s)
}

func autoConvert_coredump_CoredumpEndpointStatus_To_v1alpha1_CoredumpEndpointStatus(in *coredump.CoredumpEndpointStatus, out *CoredumpEndpointStatus, s conversion.Scope) error {
	return nil
}

// Convert_coredump_CoredumpEndpointStatus_To_v1alpha1_CoredumpEndpointStatus is an autogenerated conversion function.
func Convert_coredump_CoredumpEndpointStatus_To_v1alpha1_CoredumpEndpointStatus(in *coredump.CoredumpEndpointStatus, out *CoredumpEndpointStatus, s conversion.Scope) error {
	return autoConvert_coredump_CoredumpEndpointStatus_To_v1alpha1_CoredumpEndpointStatus(in, out, s)
}

func autoConvert_v1alpha1_CoredumpGetOptions_To_coredump_CoredumpGetOptions(in *CoredumpGetOptions, out *coredump.CoredumpGetOptions, s conversion.Scope) error {
	out.Container = in.Container
	return nil
}

// Convert_v1alpha1_CoredumpGetOptions_To_coredump_CoredumpGetOptions is an autogenerated conversion function.
func Convert_v1alpha1_CoredumpGetOptions_To_coredump_CoredumpGetOptions(in *CoredumpGetOptions, out *coredump.CoredumpGetOptions, s conversion.Scope) error {
	return autoConvert_v1alpha1_CoredumpGetOptions_To_coredump_CoredumpGetOptions(in, out, s)
}

func autoConvert_coredump_CoredumpGetOptions_To_v1alpha1_CoredumpGetOptions(in *coredump.CoredumpGetOptions, out *CoredumpGetOptions, s conversion.Scope) error {
	out.Container = in.Container
	return nil
}

// Convert_coredump_CoredumpGetOptions_To_v1alpha1_CoredumpGetOptions is an autogenerated conversion function.
func Convert_coredump_CoredumpGetOptions_To_v1alpha1_CoredumpGetOptions(in *coredump.CoredumpGetOptions, out *CoredumpGetOptions, s conversion.Scope) error {
	return autoConvert_coredump_CoredumpGetOptions_To_v1alpha1_CoredumpGetOptions(in, out, s)
}
