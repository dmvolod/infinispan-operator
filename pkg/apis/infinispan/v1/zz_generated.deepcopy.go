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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointEncryption) DeepCopyInto(out *EndpointEncryption) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointEncryption.
func (in *EndpointEncryption) DeepCopy() *EndpointEncryption {
	if in == nil {
		return nil
	}
	out := new(EndpointEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Infinispan) DeepCopyInto(out *Infinispan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Infinispan.
func (in *Infinispan) DeepCopy() *Infinispan {
	if in == nil {
		return nil
	}
	out := new(Infinispan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Infinispan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanAuthInfo) DeepCopyInto(out *InfinispanAuthInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanAuthInfo.
func (in *InfinispanAuthInfo) DeepCopy() *InfinispanAuthInfo {
	if in == nil {
		return nil
	}
	out := new(InfinispanAuthInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanCondition) DeepCopyInto(out *InfinispanCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanCondition.
func (in *InfinispanCondition) DeepCopy() *InfinispanCondition {
	if in == nil {
		return nil
	}
	out := new(InfinispanCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanContainerSpec) DeepCopyInto(out *InfinispanContainerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanContainerSpec.
func (in *InfinispanContainerSpec) DeepCopy() *InfinispanContainerSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanList) DeepCopyInto(out *InfinispanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Infinispan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanList.
func (in *InfinispanList) DeepCopy() *InfinispanList {
	if in == nil {
		return nil
	}
	out := new(InfinispanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfinispanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanLoggingSpec) DeepCopyInto(out *InfinispanLoggingSpec) {
	*out = *in
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanLoggingSpec.
func (in *InfinispanLoggingSpec) DeepCopy() *InfinispanLoggingSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanLoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanSecurity) DeepCopyInto(out *InfinispanSecurity) {
	*out = *in
	out.EndpointEncryption = in.EndpointEncryption
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanSecurity.
func (in *InfinispanSecurity) DeepCopy() *InfinispanSecurity {
	if in == nil {
		return nil
	}
	out := new(InfinispanSecurity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanServiceContainerSpec) DeepCopyInto(out *InfinispanServiceContainerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanServiceContainerSpec.
func (in *InfinispanServiceContainerSpec) DeepCopy() *InfinispanServiceContainerSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanServiceContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanServiceSpec) DeepCopyInto(out *InfinispanServiceSpec) {
	*out = *in
	out.Container = in.Container
	in.Sites.DeepCopyInto(&out.Sites)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanServiceSpec.
func (in *InfinispanServiceSpec) DeepCopy() *InfinispanServiceSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanSitesBackupSpec) DeepCopyInto(out *InfinispanSitesBackupSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanSitesBackupSpec.
func (in *InfinispanSitesBackupSpec) DeepCopy() *InfinispanSitesBackupSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanSitesBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanSitesLocalSpec) DeepCopyInto(out *InfinispanSitesLocalSpec) {
	*out = *in
	in.Expose.DeepCopyInto(&out.Expose)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanSitesLocalSpec.
func (in *InfinispanSitesLocalSpec) DeepCopy() *InfinispanSitesLocalSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanSitesLocalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanSitesSpec) DeepCopyInto(out *InfinispanSitesSpec) {
	*out = *in
	in.Local.DeepCopyInto(&out.Local)
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = make([]InfinispanSitesBackupSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanSitesSpec.
func (in *InfinispanSitesSpec) DeepCopy() *InfinispanSitesSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanSitesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanSpec) DeepCopyInto(out *InfinispanSpec) {
	*out = *in
	out.Security = in.Security
	out.Container = in.Container
	in.Service.DeepCopyInto(&out.Service)
	in.Logging.DeepCopyInto(&out.Logging)
	in.Expose.DeepCopyInto(&out.Expose)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanSpec.
func (in *InfinispanSpec) DeepCopy() *InfinispanSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanStatus) DeepCopyInto(out *InfinispanStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]InfinispanCondition, len(*in))
		copy(*out, *in)
	}
	out.Security = in.Security
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanStatus.
func (in *InfinispanStatus) DeepCopy() *InfinispanStatus {
	if in == nil {
		return nil
	}
	out := new(InfinispanStatus)
	in.DeepCopyInto(out)
	return out
}
