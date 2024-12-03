//go:build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSICluster) DeepCopyInto(out *CSICluster) {
	*out = *in
	if in.Primary != nil {
		in, out := &in.Primary, &out.Primary
		*out = new(CSIFilesystem)
		**out = **in
	}
	if in.RestApi != nil {
		in, out := &in.RestApi, &out.RestApi
		*out = make([]RestApi, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSICluster.
func (in *CSICluster) DeepCopy() *CSICluster {
	if in == nil {
		return nil
	}
	out := new(CSICluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIFilesystem) DeepCopyInto(out *CSIFilesystem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIFilesystem.
func (in *CSIFilesystem) DeepCopy() *CSIFilesystem {
	if in == nil {
		return nil
	}
	out := new(CSIFilesystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSINodeSelector) DeepCopyInto(out *CSINodeSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSINodeSelector.
func (in *CSINodeSelector) DeepCopy() *CSINodeSelector {
	if in == nil {
		return nil
	}
	out := new(CSINodeSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIScaleOperator) DeepCopyInto(out *CSIScaleOperator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIScaleOperator.
func (in *CSIScaleOperator) DeepCopy() *CSIScaleOperator {
	if in == nil {
		return nil
	}
	out := new(CSIScaleOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CSIScaleOperator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIScaleOperatorList) DeepCopyInto(out *CSIScaleOperatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CSIScaleOperator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIScaleOperatorList.
func (in *CSIScaleOperatorList) DeepCopy() *CSIScaleOperatorList {
	if in == nil {
		return nil
	}
	out := new(CSIScaleOperatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CSIScaleOperatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIScaleOperatorSpec) DeepCopyInto(out *CSIScaleOperatorSpec) {
	*out = *in
	if in.AttacherNodeSelector != nil {
		in, out := &in.AttacherNodeSelector, &out.AttacherNodeSelector
		*out = make([]CSINodeSelector, len(*in))
		copy(*out, *in)
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]CSICluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeMapping != nil {
		in, out := &in.NodeMapping, &out.NodeMapping
		*out = make([]NodeMapping, len(*in))
		copy(*out, *in)
	}
	if in.PluginNodeSelector != nil {
		in, out := &in.PluginNodeSelector, &out.PluginNodeSelector
		*out = make([]CSINodeSelector, len(*in))
		copy(*out, *in)
	}
	if in.ProvisionerNodeSelector != nil {
		in, out := &in.ProvisionerNodeSelector, &out.ProvisionerNodeSelector
		*out = make([]CSINodeSelector, len(*in))
		copy(*out, *in)
	}
	if in.SnapshotterNodeSelector != nil {
		in, out := &in.SnapshotterNodeSelector, &out.SnapshotterNodeSelector
		*out = make([]CSINodeSelector, len(*in))
		copy(*out, *in)
	}
	if in.ResizerNodeSelector != nil {
		in, out := &in.ResizerNodeSelector, &out.ResizerNodeSelector
		*out = make([]CSINodeSelector, len(*in))
		copy(*out, *in)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIScaleOperatorSpec.
func (in *CSIScaleOperatorSpec) DeepCopy() *CSIScaleOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(CSIScaleOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIScaleOperatorStatus) DeepCopyInto(out *CSIScaleOperatorStatus) {
	*out = *in
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]Version, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIScaleOperatorStatus.
func (in *CSIScaleOperatorStatus) DeepCopy() *CSIScaleOperatorStatus {
	if in == nil {
		return nil
	}
	out := new(CSIScaleOperatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMapping) DeepCopyInto(out *NodeMapping) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMapping.
func (in *NodeMapping) DeepCopy() *NodeMapping {
	if in == nil {
		return nil
	}
	out := new(NodeMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestApi) DeepCopyInto(out *RestApi) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestApi.
func (in *RestApi) DeepCopy() *RestApi {
	if in == nil {
		return nil
	}
	out := new(RestApi)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Version) DeepCopyInto(out *Version) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Version.
func (in *Version) DeepCopy() *Version {
	if in == nil {
		return nil
	}
	out := new(Version)
	in.DeepCopyInto(out)
	return out
}
