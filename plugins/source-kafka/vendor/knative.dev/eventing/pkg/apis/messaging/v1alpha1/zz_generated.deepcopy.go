// +build !ignore_autogenerated

/*
Copyright 2020 The Knative Authors

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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	duckv1alpha1 "knative.dev/eventing/pkg/apis/duck/v1alpha1"
	duckv1beta1 "knative.dev/eventing/pkg/apis/duck/v1beta1"
	v1beta1 "knative.dev/eventing/pkg/apis/messaging/v1beta1"
	apis "knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Channel) DeepCopyInto(out *Channel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Channel.
func (in *Channel) DeepCopy() *Channel {
	if in == nil {
		return nil
	}
	out := new(Channel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Channel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelList) DeepCopyInto(out *ChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Channel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelList.
func (in *ChannelList) DeepCopy() *ChannelList {
	if in == nil {
		return nil
	}
	out := new(ChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelSpec) DeepCopyInto(out *ChannelSpec) {
	*out = *in
	if in.ChannelTemplate != nil {
		in, out := &in.ChannelTemplate, &out.ChannelTemplate
		*out = new(v1beta1.ChannelTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Subscribable != nil {
		in, out := &in.Subscribable, &out.Subscribable
		*out = new(duckv1alpha1.Subscribable)
		(*in).DeepCopyInto(*out)
	}
	if in.Delivery != nil {
		in, out := &in.Delivery, &out.Delivery
		*out = new(duckv1beta1.DeliverySpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelSpec.
func (in *ChannelSpec) DeepCopy() *ChannelSpec {
	if in == nil {
		return nil
	}
	out := new(ChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelStatus) DeepCopyInto(out *ChannelStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	in.SubscribableTypeStatus.DeepCopyInto(&out.SubscribableTypeStatus)
	if in.Channel != nil {
		in, out := &in.Channel, &out.Channel
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelStatus.
func (in *ChannelStatus) DeepCopy() *ChannelStatus {
	if in == nil {
		return nil
	}
	out := new(ChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryChannel) DeepCopyInto(out *InMemoryChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryChannel.
func (in *InMemoryChannel) DeepCopy() *InMemoryChannel {
	if in == nil {
		return nil
	}
	out := new(InMemoryChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InMemoryChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryChannelList) DeepCopyInto(out *InMemoryChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InMemoryChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryChannelList.
func (in *InMemoryChannelList) DeepCopy() *InMemoryChannelList {
	if in == nil {
		return nil
	}
	out := new(InMemoryChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InMemoryChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryChannelSpec) DeepCopyInto(out *InMemoryChannelSpec) {
	*out = *in
	if in.Subscribable != nil {
		in, out := &in.Subscribable, &out.Subscribable
		*out = new(duckv1alpha1.Subscribable)
		(*in).DeepCopyInto(*out)
	}
	if in.Delivery != nil {
		in, out := &in.Delivery, &out.Delivery
		*out = new(duckv1beta1.DeliverySpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryChannelSpec.
func (in *InMemoryChannelSpec) DeepCopy() *InMemoryChannelSpec {
	if in == nil {
		return nil
	}
	out := new(InMemoryChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryChannelStatus) DeepCopyInto(out *InMemoryChannelStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	in.SubscribableTypeStatus.DeepCopyInto(&out.SubscribableTypeStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryChannelStatus.
func (in *InMemoryChannelStatus) DeepCopy() *InMemoryChannelStatus {
	if in == nil {
		return nil
	}
	out := new(InMemoryChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subscription) DeepCopyInto(out *Subscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subscription.
func (in *Subscription) DeepCopy() *Subscription {
	if in == nil {
		return nil
	}
	out := new(Subscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Subscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionList) DeepCopyInto(out *SubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Subscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionList.
func (in *SubscriptionList) DeepCopy() *SubscriptionList {
	if in == nil {
		return nil
	}
	out := new(SubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionSpec) DeepCopyInto(out *SubscriptionSpec) {
	*out = *in
	out.Channel = in.Channel
	if in.Subscriber != nil {
		in, out := &in.Subscriber, &out.Subscriber
		*out = new(duckv1.Destination)
		(*in).DeepCopyInto(*out)
	}
	if in.Reply != nil {
		in, out := &in.Reply, &out.Reply
		*out = new(duckv1.Destination)
		(*in).DeepCopyInto(*out)
	}
	if in.Delivery != nil {
		in, out := &in.Delivery, &out.Delivery
		*out = new(duckv1beta1.DeliverySpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionSpec.
func (in *SubscriptionSpec) DeepCopy() *SubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(SubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionStatus) DeepCopyInto(out *SubscriptionStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.PhysicalSubscription.DeepCopyInto(&out.PhysicalSubscription)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionStatus.
func (in *SubscriptionStatus) DeepCopy() *SubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(SubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionStatusPhysicalSubscription) DeepCopyInto(out *SubscriptionStatusPhysicalSubscription) {
	*out = *in
	if in.SubscriberURI != nil {
		in, out := &in.SubscriberURI, &out.SubscriberURI
		*out = new(apis.URL)
		(*in).DeepCopyInto(*out)
	}
	if in.ReplyURI != nil {
		in, out := &in.ReplyURI, &out.ReplyURI
		*out = new(apis.URL)
		(*in).DeepCopyInto(*out)
	}
	if in.DeadLetterSinkURI != nil {
		in, out := &in.DeadLetterSinkURI, &out.DeadLetterSinkURI
		*out = new(apis.URL)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionStatusPhysicalSubscription.
func (in *SubscriptionStatusPhysicalSubscription) DeepCopy() *SubscriptionStatusPhysicalSubscription {
	if in == nil {
		return nil
	}
	out := new(SubscriptionStatusPhysicalSubscription)
	in.DeepCopyInto(out)
	return out
}
