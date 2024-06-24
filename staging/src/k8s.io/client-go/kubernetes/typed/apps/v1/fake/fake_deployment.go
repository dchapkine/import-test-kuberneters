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
	json "encoding/json"
	"fmt"

	v1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	appsv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	applyconfigurationsautoscalingv1 "k8s.io/client-go/applyconfigurations/autoscaling/v1"
	testing "k8s.io/client-go/testing"
)

// FakeDeployments implements DeploymentInterface
type FakeDeployments struct {
	Fake *FakeAppsV1
	ns   string
}

var deploymentsResource = v1.SchemeGroupVersion.WithResource("deployments")

var deploymentsKind = v1.SchemeGroupVersion.WithKind("Deployment")

// Get takes name of the deployment, and returns the corresponding deployment object, and an error if there is any.
func (c *FakeDeployments) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Deployment, err error) {
	emptyResult := &v1.Deployment{}
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(deploymentsResource, c.ns, name), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Deployment), err
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (c *FakeDeployments) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DeploymentList, err error) {
	emptyResult := &v1.DeploymentList{}
	obj, err := c.Fake.
		Invokes(testing.NewListAction(deploymentsResource, deploymentsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	list := &v1.DeploymentList{ListMeta: obj.(*v1.DeploymentList).ListMeta}
	for _, item := range obj.(*v1.DeploymentList).Items {
		if testing.ExtractFromListOptions(opts).Labels.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deployments.
func (c *FakeDeployments) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(deploymentsResource, c.ns, opts))

}

// Create takes the representation of a deployment and creates it.  Returns the server's representation of the deployment, and an error, if there is any.
func (c *FakeDeployments) Create(ctx context.Context, deployment *v1.Deployment, opts metav1.CreateOptions) (result *v1.Deployment, err error) {
	emptyResult := &v1.Deployment{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(deploymentsResource, c.ns, deployment), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Deployment), err
}

// Update takes the representation of a deployment and updates it. Returns the server's representation of the deployment, and an error, if there is any.
func (c *FakeDeployments) Update(ctx context.Context, deployment *v1.Deployment, opts metav1.UpdateOptions) (result *v1.Deployment, err error) {
	emptyResult := &v1.Deployment{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(deploymentsResource, c.ns, deployment), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Deployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeployments) UpdateStatus(ctx context.Context, deployment *v1.Deployment, opts metav1.UpdateOptions) (result *v1.Deployment, err error) {
	emptyResult := &v1.Deployment{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(deploymentsResource, "status", c.ns, deployment), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Deployment), err
}

// Delete takes name of the deployment and deletes it. Returns an error if one occurs.
func (c *FakeDeployments) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(deploymentsResource, c.ns, name, opts), &v1.Deployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeployments) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(deploymentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.DeploymentList{})
	return err
}

// Patch applies the patch and returns the patched deployment.
func (c *FakeDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Deployment, err error) {
	emptyResult := &v1.Deployment{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deploymentsResource, c.ns, name, pt, data, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Deployment), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied deployment.
func (c *FakeDeployments) Apply(ctx context.Context, deployment *appsv1.DeploymentApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Deployment, err error) {
	if deployment == nil {
		return nil, fmt.Errorf("deployment provided to Apply must not be nil")
	}
	data, err := json.Marshal(deployment)
	if err != nil {
		return nil, err
	}
	name := deployment.Name
	if name == nil {
		return nil, fmt.Errorf("deployment.Name must be provided to Apply")
	}
	emptyResult := &v1.Deployment{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deploymentsResource, c.ns, *name, types.ApplyPatchType, data), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Deployment), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeDeployments) ApplyStatus(ctx context.Context, deployment *appsv1.DeploymentApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Deployment, err error) {
	if deployment == nil {
		return nil, fmt.Errorf("deployment provided to Apply must not be nil")
	}
	data, err := json.Marshal(deployment)
	if err != nil {
		return nil, err
	}
	name := deployment.Name
	if name == nil {
		return nil, fmt.Errorf("deployment.Name must be provided to Apply")
	}
	emptyResult := &v1.Deployment{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deploymentsResource, c.ns, *name, types.ApplyPatchType, data, "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Deployment), err
}

// GetScale takes name of the deployment, and returns the corresponding scale object, and an error if there is any.
func (c *FakeDeployments) GetScale(ctx context.Context, deploymentName string, options metav1.GetOptions) (result *autoscalingv1.Scale, err error) {
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceAction(deploymentsResource, c.ns, "scale", deploymentName), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakeDeployments) UpdateScale(ctx context.Context, deploymentName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(deploymentsResource, "scale", c.ns, scale), &autoscalingv1.Scale{})

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}

// ApplyScale takes top resource name and the apply declarative configuration for scale,
// applies it and returns the applied scale, and an error, if there is any.
func (c *FakeDeployments) ApplyScale(ctx context.Context, deploymentName string, scale *applyconfigurationsautoscalingv1.ScaleApplyConfiguration, opts metav1.ApplyOptions) (result *autoscalingv1.Scale, err error) {
	if scale == nil {
		return nil, fmt.Errorf("scale provided to ApplyScale must not be nil")
	}
	data, err := json.Marshal(scale)
	if err != nil {
		return nil, err
	}
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deploymentsResource, c.ns, deploymentName, types.ApplyPatchType, data, "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}
