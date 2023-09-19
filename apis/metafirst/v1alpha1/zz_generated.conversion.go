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
	unsafe "unsafe"

	metafirst "github.com/Dimss/first/apis/metafirst"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*First)(nil), (*metafirst.First)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_First_To_metafirst_First(a.(*First), b.(*metafirst.First), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metafirst.First)(nil), (*First)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metafirst_First_To_v1alpha1_First(a.(*metafirst.First), b.(*First), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FirstList)(nil), (*metafirst.FirstList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FirstList_To_metafirst_FirstList(a.(*FirstList), b.(*metafirst.FirstList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metafirst.FirstList)(nil), (*FirstList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metafirst_FirstList_To_v1alpha1_FirstList(a.(*metafirst.FirstList), b.(*FirstList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FirstSpec)(nil), (*metafirst.FirstSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FirstSpec_To_metafirst_FirstSpec(a.(*FirstSpec), b.(*metafirst.FirstSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metafirst.FirstSpec)(nil), (*FirstSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metafirst_FirstSpec_To_v1alpha1_FirstSpec(a.(*metafirst.FirstSpec), b.(*FirstSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FirstStatus)(nil), (*metafirst.FirstStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FirstStatus_To_metafirst_FirstStatus(a.(*FirstStatus), b.(*metafirst.FirstStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metafirst.FirstStatus)(nil), (*FirstStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metafirst_FirstStatus_To_v1alpha1_FirstStatus(a.(*metafirst.FirstStatus), b.(*FirstStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_First_To_metafirst_First(in *First, out *metafirst.First, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_FirstSpec_To_metafirst_FirstSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_FirstStatus_To_metafirst_FirstStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_First_To_metafirst_First is an autogenerated conversion function.
func Convert_v1alpha1_First_To_metafirst_First(in *First, out *metafirst.First, s conversion.Scope) error {
	return autoConvert_v1alpha1_First_To_metafirst_First(in, out, s)
}

func autoConvert_metafirst_First_To_v1alpha1_First(in *metafirst.First, out *First, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_metafirst_FirstSpec_To_v1alpha1_FirstSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_metafirst_FirstStatus_To_v1alpha1_FirstStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_metafirst_First_To_v1alpha1_First is an autogenerated conversion function.
func Convert_metafirst_First_To_v1alpha1_First(in *metafirst.First, out *First, s conversion.Scope) error {
	return autoConvert_metafirst_First_To_v1alpha1_First(in, out, s)
}

func autoConvert_v1alpha1_FirstList_To_metafirst_FirstList(in *FirstList, out *metafirst.FirstList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]metafirst.First)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_FirstList_To_metafirst_FirstList is an autogenerated conversion function.
func Convert_v1alpha1_FirstList_To_metafirst_FirstList(in *FirstList, out *metafirst.FirstList, s conversion.Scope) error {
	return autoConvert_v1alpha1_FirstList_To_metafirst_FirstList(in, out, s)
}

func autoConvert_metafirst_FirstList_To_v1alpha1_FirstList(in *metafirst.FirstList, out *FirstList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]First)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_metafirst_FirstList_To_v1alpha1_FirstList is an autogenerated conversion function.
func Convert_metafirst_FirstList_To_v1alpha1_FirstList(in *metafirst.FirstList, out *FirstList, s conversion.Scope) error {
	return autoConvert_metafirst_FirstList_To_v1alpha1_FirstList(in, out, s)
}

func autoConvert_v1alpha1_FirstSpec_To_metafirst_FirstSpec(in *FirstSpec, out *metafirst.FirstSpec, s conversion.Scope) error {
	out.Message = in.Message
	return nil
}

// Convert_v1alpha1_FirstSpec_To_metafirst_FirstSpec is an autogenerated conversion function.
func Convert_v1alpha1_FirstSpec_To_metafirst_FirstSpec(in *FirstSpec, out *metafirst.FirstSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_FirstSpec_To_metafirst_FirstSpec(in, out, s)
}

func autoConvert_metafirst_FirstSpec_To_v1alpha1_FirstSpec(in *metafirst.FirstSpec, out *FirstSpec, s conversion.Scope) error {
	out.Message = in.Message
	return nil
}

// Convert_metafirst_FirstSpec_To_v1alpha1_FirstSpec is an autogenerated conversion function.
func Convert_metafirst_FirstSpec_To_v1alpha1_FirstSpec(in *metafirst.FirstSpec, out *FirstSpec, s conversion.Scope) error {
	return autoConvert_metafirst_FirstSpec_To_v1alpha1_FirstSpec(in, out, s)
}

func autoConvert_v1alpha1_FirstStatus_To_metafirst_FirstStatus(in *FirstStatus, out *metafirst.FirstStatus, s conversion.Scope) error {
	out.Status = in.Status
	return nil
}

// Convert_v1alpha1_FirstStatus_To_metafirst_FirstStatus is an autogenerated conversion function.
func Convert_v1alpha1_FirstStatus_To_metafirst_FirstStatus(in *FirstStatus, out *metafirst.FirstStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_FirstStatus_To_metafirst_FirstStatus(in, out, s)
}

func autoConvert_metafirst_FirstStatus_To_v1alpha1_FirstStatus(in *metafirst.FirstStatus, out *FirstStatus, s conversion.Scope) error {
	out.Status = in.Status
	return nil
}

// Convert_metafirst_FirstStatus_To_v1alpha1_FirstStatus is an autogenerated conversion function.
func Convert_metafirst_FirstStatus_To_v1alpha1_FirstStatus(in *metafirst.FirstStatus, out *FirstStatus, s conversion.Scope) error {
	return autoConvert_metafirst_FirstStatus_To_v1alpha1_FirstStatus(in, out, s)
}
