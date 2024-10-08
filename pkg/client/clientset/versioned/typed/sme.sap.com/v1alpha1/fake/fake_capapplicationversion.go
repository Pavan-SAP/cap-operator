/*
SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and cap-operator contributors
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/sap/cap-operator/pkg/apis/sme.sap.com/v1alpha1"
	smesapcomv1alpha1 "github.com/sap/cap-operator/pkg/client/applyconfiguration/sme.sap.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCAPApplicationVersions implements CAPApplicationVersionInterface
type FakeCAPApplicationVersions struct {
	Fake *FakeSmeV1alpha1
	ns   string
}

var capapplicationversionsResource = v1alpha1.SchemeGroupVersion.WithResource("capapplicationversions")

var capapplicationversionsKind = v1alpha1.SchemeGroupVersion.WithKind("CAPApplicationVersion")

// Get takes name of the cAPApplicationVersion, and returns the corresponding cAPApplicationVersion object, and an error if there is any.
func (c *FakeCAPApplicationVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CAPApplicationVersion, err error) {
	emptyResult := &v1alpha1.CAPApplicationVersion{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(capapplicationversionsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CAPApplicationVersion), err
}

// List takes label and field selectors, and returns the list of CAPApplicationVersions that match those selectors.
func (c *FakeCAPApplicationVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CAPApplicationVersionList, err error) {
	emptyResult := &v1alpha1.CAPApplicationVersionList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(capapplicationversionsResource, capapplicationversionsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CAPApplicationVersionList{ListMeta: obj.(*v1alpha1.CAPApplicationVersionList).ListMeta}
	for _, item := range obj.(*v1alpha1.CAPApplicationVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cAPApplicationVersions.
func (c *FakeCAPApplicationVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(capapplicationversionsResource, c.ns, opts))

}

// Create takes the representation of a cAPApplicationVersion and creates it.  Returns the server's representation of the cAPApplicationVersion, and an error, if there is any.
func (c *FakeCAPApplicationVersions) Create(ctx context.Context, cAPApplicationVersion *v1alpha1.CAPApplicationVersion, opts v1.CreateOptions) (result *v1alpha1.CAPApplicationVersion, err error) {
	emptyResult := &v1alpha1.CAPApplicationVersion{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(capapplicationversionsResource, c.ns, cAPApplicationVersion, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CAPApplicationVersion), err
}

// Update takes the representation of a cAPApplicationVersion and updates it. Returns the server's representation of the cAPApplicationVersion, and an error, if there is any.
func (c *FakeCAPApplicationVersions) Update(ctx context.Context, cAPApplicationVersion *v1alpha1.CAPApplicationVersion, opts v1.UpdateOptions) (result *v1alpha1.CAPApplicationVersion, err error) {
	emptyResult := &v1alpha1.CAPApplicationVersion{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(capapplicationversionsResource, c.ns, cAPApplicationVersion, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CAPApplicationVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCAPApplicationVersions) UpdateStatus(ctx context.Context, cAPApplicationVersion *v1alpha1.CAPApplicationVersion, opts v1.UpdateOptions) (result *v1alpha1.CAPApplicationVersion, err error) {
	emptyResult := &v1alpha1.CAPApplicationVersion{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(capapplicationversionsResource, "status", c.ns, cAPApplicationVersion, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CAPApplicationVersion), err
}

// Delete takes name of the cAPApplicationVersion and deletes it. Returns an error if one occurs.
func (c *FakeCAPApplicationVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(capapplicationversionsResource, c.ns, name, opts), &v1alpha1.CAPApplicationVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCAPApplicationVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(capapplicationversionsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CAPApplicationVersionList{})
	return err
}

// Patch applies the patch and returns the patched cAPApplicationVersion.
func (c *FakeCAPApplicationVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CAPApplicationVersion, err error) {
	emptyResult := &v1alpha1.CAPApplicationVersion{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(capapplicationversionsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CAPApplicationVersion), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied cAPApplicationVersion.
func (c *FakeCAPApplicationVersions) Apply(ctx context.Context, cAPApplicationVersion *smesapcomv1alpha1.CAPApplicationVersionApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.CAPApplicationVersion, err error) {
	if cAPApplicationVersion == nil {
		return nil, fmt.Errorf("cAPApplicationVersion provided to Apply must not be nil")
	}
	data, err := json.Marshal(cAPApplicationVersion)
	if err != nil {
		return nil, err
	}
	name := cAPApplicationVersion.Name
	if name == nil {
		return nil, fmt.Errorf("cAPApplicationVersion.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.CAPApplicationVersion{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(capapplicationversionsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CAPApplicationVersion), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeCAPApplicationVersions) ApplyStatus(ctx context.Context, cAPApplicationVersion *smesapcomv1alpha1.CAPApplicationVersionApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.CAPApplicationVersion, err error) {
	if cAPApplicationVersion == nil {
		return nil, fmt.Errorf("cAPApplicationVersion provided to Apply must not be nil")
	}
	data, err := json.Marshal(cAPApplicationVersion)
	if err != nil {
		return nil, err
	}
	name := cAPApplicationVersion.Name
	if name == nil {
		return nil, fmt.Errorf("cAPApplicationVersion.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.CAPApplicationVersion{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(capapplicationversionsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CAPApplicationVersion), err
}
