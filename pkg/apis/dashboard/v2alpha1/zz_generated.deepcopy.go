//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by deepcopy-gen. DO NOT EDIT.

package v2alpha1

import (
	v0alpha1 "github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnnotationActions) DeepCopyInto(out *AnnotationActions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnnotationActions.
func (in *AnnotationActions) DeepCopy() *AnnotationActions {
	if in == nil {
		return nil
	}
	out := new(AnnotationActions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnnotationPermission) DeepCopyInto(out *AnnotationPermission) {
	*out = *in
	out.Dashboard = in.Dashboard
	out.Organization = in.Organization
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnnotationPermission.
func (in *AnnotationPermission) DeepCopy() *AnnotationPermission {
	if in == nil {
		return nil
	}
	out := new(AnnotationPermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConversionStatus) DeepCopyInto(out *ConversionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConversionStatus.
func (in *ConversionStatus) DeepCopy() *ConversionStatus {
	if in == nil {
		return nil
	}
	out := new(ConversionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dashboard) DeepCopyInto(out *Dashboard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(DashboardStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dashboard.
func (in *Dashboard) DeepCopy() *Dashboard {
	if in == nil {
		return nil
	}
	out := new(Dashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Dashboard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardAccess) DeepCopyInto(out *DashboardAccess) {
	*out = *in
	if in.AnnotationsPermissions != nil {
		in, out := &in.AnnotationsPermissions, &out.AnnotationsPermissions
		*out = new(AnnotationPermission)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardAccess.
func (in *DashboardAccess) DeepCopy() *DashboardAccess {
	if in == nil {
		return nil
	}
	out := new(DashboardAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardList) DeepCopyInto(out *DashboardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Dashboard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardList.
func (in *DashboardList) DeepCopy() *DashboardList {
	if in == nil {
		return nil
	}
	out := new(DashboardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DashboardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardStatus) DeepCopyInto(out *DashboardStatus) {
	*out = *in
	if in.ConversionStatus != nil {
		in, out := &in.ConversionStatus, &out.ConversionStatus
		*out = new(ConversionStatus)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardStatus.
func (in *DashboardStatus) DeepCopy() *DashboardStatus {
	if in == nil {
		return nil
	}
	out := new(DashboardStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardVersionInfo) DeepCopyInto(out *DashboardVersionInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardVersionInfo.
func (in *DashboardVersionInfo) DeepCopy() *DashboardVersionInfo {
	if in == nil {
		return nil
	}
	out := new(DashboardVersionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardVersionList) DeepCopyInto(out *DashboardVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DashboardVersionInfo, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardVersionList.
func (in *DashboardVersionList) DeepCopy() *DashboardVersionList {
	if in == nil {
		return nil
	}
	out := new(DashboardVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DashboardVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardWithAccessInfo) DeepCopyInto(out *DashboardWithAccessInfo) {
	*out = *in
	in.Dashboard.DeepCopyInto(&out.Dashboard)
	in.Access.DeepCopyInto(&out.Access)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardWithAccessInfo.
func (in *DashboardWithAccessInfo) DeepCopy() *DashboardWithAccessInfo {
	if in == nil {
		return nil
	}
	out := new(DashboardWithAccessInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DashboardWithAccessInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibraryPanel) DeepCopyInto(out *LibraryPanel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(LibraryPanelStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibraryPanel.
func (in *LibraryPanel) DeepCopy() *LibraryPanel {
	if in == nil {
		return nil
	}
	out := new(LibraryPanel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LibraryPanel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibraryPanelList) DeepCopyInto(out *LibraryPanelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LibraryPanel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibraryPanelList.
func (in *LibraryPanelList) DeepCopy() *LibraryPanelList {
	if in == nil {
		return nil
	}
	out := new(LibraryPanelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LibraryPanelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibraryPanelSpec) DeepCopyInto(out *LibraryPanelSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	in.FieldConfig.DeepCopyInto(&out.FieldConfig)
	if in.Datasource != nil {
		in, out := &in.Datasource, &out.Datasource
		*out = new(v0alpha1.DataSourceRef)
		**out = **in
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]v0alpha1.DataQuery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibraryPanelSpec.
func (in *LibraryPanelSpec) DeepCopy() *LibraryPanelSpec {
	if in == nil {
		return nil
	}
	out := new(LibraryPanelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibraryPanelStatus) DeepCopyInto(out *LibraryPanelStatus) {
	*out = *in
	if in.Warnings != nil {
		in, out := &in.Warnings, &out.Warnings
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Missing.DeepCopyInto(&out.Missing)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibraryPanelStatus.
func (in *LibraryPanelStatus) DeepCopy() *LibraryPanelStatus {
	if in == nil {
		return nil
	}
	out := new(LibraryPanelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionsQueryOptions) DeepCopyInto(out *VersionsQueryOptions) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionsQueryOptions.
func (in *VersionsQueryOptions) DeepCopy() *VersionsQueryOptions {
	if in == nil {
		return nil
	}
	out := new(VersionsQueryOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VersionsQueryOptions) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
