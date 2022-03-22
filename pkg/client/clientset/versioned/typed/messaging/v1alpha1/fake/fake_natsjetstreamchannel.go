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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/eventing-natss/pkg/apis/messaging/v1alpha1"
)

// FakeNatsJetStreamChannels implements NatsJetStreamChannelInterface
type FakeNatsJetStreamChannels struct {
	Fake *FakeMessagingV1alpha1
	ns   string
}

var natsjetstreamchannelsResource = schema.GroupVersionResource{Group: "messaging.knative.dev", Version: "v1alpha1", Resource: "natsjetstreamchannels"}

var natsjetstreamchannelsKind = schema.GroupVersionKind{Group: "messaging.knative.dev", Version: "v1alpha1", Kind: "NatsJetStreamChannel"}

// Get takes name of the natsJetStreamChannel, and returns the corresponding natsJetStreamChannel object, and an error if there is any.
func (c *FakeNatsJetStreamChannels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NatsJetStreamChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(natsjetstreamchannelsResource, c.ns, name), &v1alpha1.NatsJetStreamChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatsJetStreamChannel), err
}

// List takes label and field selectors, and returns the list of NatsJetStreamChannels that match those selectors.
func (c *FakeNatsJetStreamChannels) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NatsJetStreamChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(natsjetstreamchannelsResource, natsjetstreamchannelsKind, c.ns, opts), &v1alpha1.NatsJetStreamChannelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NatsJetStreamChannelList{ListMeta: obj.(*v1alpha1.NatsJetStreamChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.NatsJetStreamChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested natsJetStreamChannels.
func (c *FakeNatsJetStreamChannels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(natsjetstreamchannelsResource, c.ns, opts))

}

// Create takes the representation of a natsJetStreamChannel and creates it.  Returns the server's representation of the natsJetStreamChannel, and an error, if there is any.
func (c *FakeNatsJetStreamChannels) Create(ctx context.Context, natsJetStreamChannel *v1alpha1.NatsJetStreamChannel, opts v1.CreateOptions) (result *v1alpha1.NatsJetStreamChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(natsjetstreamchannelsResource, c.ns, natsJetStreamChannel), &v1alpha1.NatsJetStreamChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatsJetStreamChannel), err
}

// Update takes the representation of a natsJetStreamChannel and updates it. Returns the server's representation of the natsJetStreamChannel, and an error, if there is any.
func (c *FakeNatsJetStreamChannels) Update(ctx context.Context, natsJetStreamChannel *v1alpha1.NatsJetStreamChannel, opts v1.UpdateOptions) (result *v1alpha1.NatsJetStreamChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(natsjetstreamchannelsResource, c.ns, natsJetStreamChannel), &v1alpha1.NatsJetStreamChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatsJetStreamChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNatsJetStreamChannels) UpdateStatus(ctx context.Context, natsJetStreamChannel *v1alpha1.NatsJetStreamChannel, opts v1.UpdateOptions) (*v1alpha1.NatsJetStreamChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(natsjetstreamchannelsResource, "status", c.ns, natsJetStreamChannel), &v1alpha1.NatsJetStreamChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatsJetStreamChannel), err
}

// Delete takes name of the natsJetStreamChannel and deletes it. Returns an error if one occurs.
func (c *FakeNatsJetStreamChannels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(natsjetstreamchannelsResource, c.ns, name, opts), &v1alpha1.NatsJetStreamChannel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNatsJetStreamChannels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(natsjetstreamchannelsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NatsJetStreamChannelList{})
	return err
}

// Patch applies the patch and returns the patched natsJetStreamChannel.
func (c *FakeNatsJetStreamChannels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NatsJetStreamChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(natsjetstreamchannelsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NatsJetStreamChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatsJetStreamChannel), err
}
