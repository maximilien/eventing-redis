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
	v1alpha1 "knative.dev/eventing-redis/sink/pkg/apis/sinks/v1alpha1"
)

// FakeRedisStreamSinks implements RedisStreamSinkInterface
type FakeRedisStreamSinks struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var redisstreamsinksResource = schema.GroupVersionResource{Group: "sources.knative.dev", Version: "v1alpha1", Resource: "redisstreamsinks"}

var redisstreamsinksKind = schema.GroupVersionKind{Group: "sources.knative.dev", Version: "v1alpha1", Kind: "RedisStreamSink"}

// Get takes name of the redisStreamSink, and returns the corresponding redisStreamSink object, and an error if there is any.
func (c *FakeRedisStreamSinks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RedisStreamSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(redisstreamsinksResource, c.ns, name), &v1alpha1.RedisStreamSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisStreamSink), err
}

// List takes label and field selectors, and returns the list of RedisStreamSinks that match those selectors.
func (c *FakeRedisStreamSinks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RedisStreamSinkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(redisstreamsinksResource, redisstreamsinksKind, c.ns, opts), &v1alpha1.RedisStreamSinkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RedisStreamSinkList{ListMeta: obj.(*v1alpha1.RedisStreamSinkList).ListMeta}
	for _, item := range obj.(*v1alpha1.RedisStreamSinkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested redisStreamSinks.
func (c *FakeRedisStreamSinks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(redisstreamsinksResource, c.ns, opts))

}

// Create takes the representation of a redisStreamSink and creates it.  Returns the server's representation of the redisStreamSink, and an error, if there is any.
func (c *FakeRedisStreamSinks) Create(ctx context.Context, redisStreamSink *v1alpha1.RedisStreamSink) (result *v1alpha1.RedisStreamSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(redisstreamsinksResource, c.ns, redisStreamSink), &v1alpha1.RedisStreamSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisStreamSink), err
}

// Update takes the representation of a redisStreamSink and updates it. Returns the server's representation of the redisStreamSink, and an error, if there is any.
func (c *FakeRedisStreamSinks) Update(ctx context.Context, redisStreamSink *v1alpha1.RedisStreamSink) (result *v1alpha1.RedisStreamSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(redisstreamsinksResource, c.ns, redisStreamSink), &v1alpha1.RedisStreamSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisStreamSink), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRedisStreamSinks) UpdateStatus(ctx context.Context, redisStreamSink *v1alpha1.RedisStreamSink) (*v1alpha1.RedisStreamSink, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(redisstreamsinksResource, "status", c.ns, redisStreamSink), &v1alpha1.RedisStreamSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisStreamSink), err
}

// Delete takes name of the redisStreamSink and deletes it. Returns an error if one occurs.
func (c *FakeRedisStreamSinks) Delete(ctx context.Context, name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(redisstreamsinksResource, c.ns, name), &v1alpha1.RedisStreamSink{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRedisStreamSinks) DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(redisstreamsinksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RedisStreamSinkList{})
	return err
}

// Patch applies the patch and returns the patched redisStreamSink.
func (c *FakeRedisStreamSinks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedisStreamSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(redisstreamsinksResource, c.ns, name, pt, data, subresources...), &v1alpha1.RedisStreamSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisStreamSink), err
}
