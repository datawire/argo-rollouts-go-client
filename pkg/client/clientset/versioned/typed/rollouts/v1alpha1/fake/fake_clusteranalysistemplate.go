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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/datawire/argo-rollouts-go-client/pkg/apis/rollouts/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterAnalysisTemplates implements ClusterAnalysisTemplateInterface
type FakeClusterAnalysisTemplates struct {
	Fake *FakeArgoprojV1alpha1
}

var clusteranalysistemplatesResource = v1alpha1.SchemeGroupVersion.WithResource("clusteranalysistemplates")

var clusteranalysistemplatesKind = v1alpha1.SchemeGroupVersion.WithKind("ClusterAnalysisTemplate")

// Get takes name of the clusterAnalysisTemplate, and returns the corresponding clusterAnalysisTemplate object, and an error if there is any.
func (c *FakeClusterAnalysisTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterAnalysisTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusteranalysistemplatesResource, name), &v1alpha1.ClusterAnalysisTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterAnalysisTemplate), err
}

// List takes label and field selectors, and returns the list of ClusterAnalysisTemplates that match those selectors.
func (c *FakeClusterAnalysisTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterAnalysisTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusteranalysistemplatesResource, clusteranalysistemplatesKind, opts), &v1alpha1.ClusterAnalysisTemplateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterAnalysisTemplateList{ListMeta: obj.(*v1alpha1.ClusterAnalysisTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterAnalysisTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterAnalysisTemplates.
func (c *FakeClusterAnalysisTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusteranalysistemplatesResource, opts))
}

// Create takes the representation of a clusterAnalysisTemplate and creates it.  Returns the server's representation of the clusterAnalysisTemplate, and an error, if there is any.
func (c *FakeClusterAnalysisTemplates) Create(ctx context.Context, clusterAnalysisTemplate *v1alpha1.ClusterAnalysisTemplate, opts v1.CreateOptions) (result *v1alpha1.ClusterAnalysisTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusteranalysistemplatesResource, clusterAnalysisTemplate), &v1alpha1.ClusterAnalysisTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterAnalysisTemplate), err
}

// Update takes the representation of a clusterAnalysisTemplate and updates it. Returns the server's representation of the clusterAnalysisTemplate, and an error, if there is any.
func (c *FakeClusterAnalysisTemplates) Update(ctx context.Context, clusterAnalysisTemplate *v1alpha1.ClusterAnalysisTemplate, opts v1.UpdateOptions) (result *v1alpha1.ClusterAnalysisTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusteranalysistemplatesResource, clusterAnalysisTemplate), &v1alpha1.ClusterAnalysisTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterAnalysisTemplate), err
}

// Delete takes name of the clusterAnalysisTemplate and deletes it. Returns an error if one occurs.
func (c *FakeClusterAnalysisTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusteranalysistemplatesResource, name, opts), &v1alpha1.ClusterAnalysisTemplate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterAnalysisTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusteranalysistemplatesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterAnalysisTemplateList{})
	return err
}

// Patch applies the patch and returns the patched clusterAnalysisTemplate.
func (c *FakeClusterAnalysisTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterAnalysisTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusteranalysistemplatesResource, name, pt, data, subresources...), &v1alpha1.ClusterAnalysisTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterAnalysisTemplate), err
}
