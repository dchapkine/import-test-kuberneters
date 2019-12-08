/*
Copyright 2019 The Kubernetes Authors.

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

package scheduler

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/scheduler/internal/cache"
	"k8s.io/kubernetes/pkg/scheduler/internal/queue"

	fakecache "k8s.io/kubernetes/pkg/scheduler/internal/cache/fake"
)

func TestSkipPodUpdate(t *testing.T) {
	table := []struct {
		pod              *v1.Pod
		isAssumedPodFunc func(*v1.Pod) bool
		getPodFunc       func(*v1.Pod) *v1.Pod
		expected         bool
		name             string
	}{
		{
			name: "Non-assumed pod",
			pod: &v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "pod-0",
				},
			},
			isAssumedPodFunc: func(*v1.Pod) bool { return false },
			getPodFunc: func(*v1.Pod) *v1.Pod {
				return &v1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name: "pod-0",
					},
				}
			},
			expected: false,
		},
		{
			name: "with changes on ResourceVersion, Spec.NodeName and/or Annotations",
			pod: &v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:            "pod-0",
					Annotations:     map[string]string{"a": "b"},
					ResourceVersion: "0",
				},
				Spec: v1.PodSpec{
					NodeName: "node-0",
				},
			},
			isAssumedPodFunc: func(*v1.Pod) bool {
				return true
			},
			getPodFunc: func(*v1.Pod) *v1.Pod {
				return &v1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:            "pod-0",
						Annotations:     map[string]string{"c": "d"},
						ResourceVersion: "1",
					},
					Spec: v1.PodSpec{
						NodeName: "node-1",
					},
				}
			},
			expected: true,
		},
		{
			name: "with changes on Labels",
			pod: &v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "pod-0",
					Labels: map[string]string{"a": "b"},
				},
			},
			isAssumedPodFunc: func(*v1.Pod) bool {
				return true
			},
			getPodFunc: func(*v1.Pod) *v1.Pod {
				return &v1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:   "pod-0",
						Labels: map[string]string{"c": "d"},
					},
				}
			},
			expected: false,
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			c := &Scheduler{
				SchedulerCache: &fakecache.Cache{
					IsAssumedPodFunc: test.isAssumedPodFunc,
					GetPodFunc:       test.getPodFunc,
				},
			}
			got := c.skipPodUpdate(test.pod)
			if got != test.expected {
				t.Errorf("skipPodUpdate() = %t, expected = %t", got, test.expected)
			}
		})
	}
}

func TestNodeAllocatableChanged(t *testing.T) {
	newQuantity := func(value int64) resource.Quantity {
		return *resource.NewQuantity(value, resource.BinarySI)
	}
	for _, c := range []struct {
		Name           string
		Changed        bool
		OldAllocatable v1.ResourceList
		NewAllocatable v1.ResourceList
	}{
		{
			Name:           "no allocatable resources changed",
			Changed:        false,
			OldAllocatable: v1.ResourceList{v1.ResourceMemory: newQuantity(1024)},
			NewAllocatable: v1.ResourceList{v1.ResourceMemory: newQuantity(1024)},
		},
		{
			Name:           "new node has more allocatable resources",
			Changed:        true,
			OldAllocatable: v1.ResourceList{v1.ResourceMemory: newQuantity(1024)},
			NewAllocatable: v1.ResourceList{v1.ResourceMemory: newQuantity(1024), v1.ResourceStorage: newQuantity(1024)},
		},
	} {
		oldNode := &v1.Node{Status: v1.NodeStatus{Allocatable: c.OldAllocatable}}
		newNode := &v1.Node{Status: v1.NodeStatus{Allocatable: c.NewAllocatable}}
		changed := nodeAllocatableChanged(newNode, oldNode)
		if changed != c.Changed {
			t.Errorf("nodeAllocatableChanged should be %t, got %t", c.Changed, changed)
		}
	}
}

func TestNodeLabelsChanged(t *testing.T) {
	for _, c := range []struct {
		Name      string
		Changed   bool
		OldLabels map[string]string
		NewLabels map[string]string
	}{
		{
			Name:      "no labels changed",
			Changed:   false,
			OldLabels: map[string]string{"foo": "bar"},
			NewLabels: map[string]string{"foo": "bar"},
		},
		// Labels changed.
		{
			Name:      "new node has more labels",
			Changed:   true,
			OldLabels: map[string]string{"foo": "bar"},
			NewLabels: map[string]string{"foo": "bar", "test": "value"},
		},
	} {
		oldNode := &v1.Node{ObjectMeta: metav1.ObjectMeta{Labels: c.OldLabels}}
		newNode := &v1.Node{ObjectMeta: metav1.ObjectMeta{Labels: c.NewLabels}}
		changed := nodeLabelsChanged(newNode, oldNode)
		if changed != c.Changed {
			t.Errorf("Test case %q failed: should be %t, got %t", c.Name, c.Changed, changed)
		}
	}
}

func TestNodeTaintsChanged(t *testing.T) {
	for _, c := range []struct {
		Name      string
		Changed   bool
		OldTaints []v1.Taint
		NewTaints []v1.Taint
	}{
		{
			Name:      "no taint changed",
			Changed:   false,
			OldTaints: []v1.Taint{{Key: "key", Value: "value"}},
			NewTaints: []v1.Taint{{Key: "key", Value: "value"}},
		},
		{
			Name:      "taint value changed",
			Changed:   true,
			OldTaints: []v1.Taint{{Key: "key", Value: "value1"}},
			NewTaints: []v1.Taint{{Key: "key", Value: "value2"}},
		},
	} {
		oldNode := &v1.Node{Spec: v1.NodeSpec{Taints: c.OldTaints}}
		newNode := &v1.Node{Spec: v1.NodeSpec{Taints: c.NewTaints}}
		changed := nodeTaintsChanged(newNode, oldNode)
		if changed != c.Changed {
			t.Errorf("Test case %q failed: should be %t, not %t", c.Name, c.Changed, changed)
		}
	}
}

func TestNodeConditionsChanged(t *testing.T) {
	nodeConditionType := reflect.TypeOf(v1.NodeCondition{})
	if nodeConditionType.NumField() != 6 {
		t.Errorf("NodeCondition type has changed. The nodeConditionsChanged() function must be reevaluated.")
	}

	for _, c := range []struct {
		Name          string
		Changed       bool
		OldConditions []v1.NodeCondition
		NewConditions []v1.NodeCondition
	}{
		{
			Name:          "no condition changed",
			Changed:       false,
			OldConditions: []v1.NodeCondition{{Type: v1.NodeDiskPressure, Status: v1.ConditionTrue}},
			NewConditions: []v1.NodeCondition{{Type: v1.NodeDiskPressure, Status: v1.ConditionTrue}},
		},
		{
			Name:          "only LastHeartbeatTime changed",
			Changed:       false,
			OldConditions: []v1.NodeCondition{{Type: v1.NodeDiskPressure, Status: v1.ConditionTrue, LastHeartbeatTime: metav1.Unix(1, 0)}},
			NewConditions: []v1.NodeCondition{{Type: v1.NodeDiskPressure, Status: v1.ConditionTrue, LastHeartbeatTime: metav1.Unix(2, 0)}},
		},
		{
			Name:          "new node has more healthy conditions",
			Changed:       true,
			OldConditions: []v1.NodeCondition{},
			NewConditions: []v1.NodeCondition{{Type: v1.NodeReady, Status: v1.ConditionTrue}},
		},
		{
			Name:          "new node has less unhealthy conditions",
			Changed:       true,
			OldConditions: []v1.NodeCondition{{Type: v1.NodeDiskPressure, Status: v1.ConditionTrue}},
			NewConditions: []v1.NodeCondition{},
		},
		{
			Name:          "condition status changed",
			Changed:       true,
			OldConditions: []v1.NodeCondition{{Type: v1.NodeReady, Status: v1.ConditionFalse}},
			NewConditions: []v1.NodeCondition{{Type: v1.NodeReady, Status: v1.ConditionTrue}},
		},
	} {
		oldNode := &v1.Node{Status: v1.NodeStatus{Conditions: c.OldConditions}}
		newNode := &v1.Node{Status: v1.NodeStatus{Conditions: c.NewConditions}}
		changed := nodeConditionsChanged(newNode, oldNode)
		if changed != c.Changed {
			t.Errorf("Test case %q failed: should be %t, got %t", c.Name, c.Changed, changed)
		}
	}
}

func TestUpdatePodInCache(t *testing.T) {
	ttl := 10 * time.Second
	nodeName := "node"

	tests := []struct {
		name   string
		oldObj interface{}
		newObj interface{}
	}{
		{
			name:   "test update pod with same uid",
			oldObj: makeBasePod(t, nodeName, "oldUID", "testName", "100m", "500", "", []v1.ContainerPort{{HostIP: "127.0.0.1", HostPort: 80, Protocol: "TCP"}}),
			newObj: makeBasePod(t, nodeName, "oldUID", "testName", "100m", "500", "", []v1.ContainerPort{{HostIP: "127.0.0.1", HostPort: 8080, Protocol: "TCP"}}),
		},
		{
			name:   "test update pod with different uid",
			oldObj: makeBasePod(t, nodeName, "oldUID", "testName", "100m", "500", "", []v1.ContainerPort{{HostIP: "127.0.0.1", HostPort: 80, Protocol: "TCP"}}),
			newObj: makeBasePod(t, nodeName, "newUID", "testName", "100m", "500", "", []v1.ContainerPort{{HostIP: "127.0.0.1", HostPort: 8080, Protocol: "TCP"}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stopCh := make(chan struct{})
			defer close(stopCh)
			schedulerCache := cache.New(ttl, stopCh)
			schedulerQueue := queue.NewPriorityQueue(nil)
			sched := &Scheduler{
				SchedulerCache:  schedulerCache,
				SchedulingQueue: schedulerQueue,
			}
			sched.addPodToCache(tt.oldObj)
			sched.updatePodInCache(tt.oldObj, tt.newObj)
			pod, err := sched.SchedulerCache.GetPod(tt.newObj.(*v1.Pod))
			if err != nil {
				t.Errorf("Failed to get pod from scheduler: %v", err)
			}
			if pod.UID != tt.newObj.(*v1.Pod).UID {
				t.Errorf("Expected get pod UID %v, got %v", tt.newObj.(*v1.Pod).UID, pod.UID)
			}
		})
	}
}

func makeBasePod(t *testing.T, nodeName, objUID, objName, cpu, mem, extended string, ports []v1.ContainerPort) *v1.Pod {
	req := v1.ResourceList{}
	if cpu != "" {
		req = v1.ResourceList{
			v1.ResourceCPU:    resource.MustParse(cpu),
			v1.ResourceMemory: resource.MustParse(mem),
		}
		if extended != "" {
			parts := strings.Split(extended, ":")
			if len(parts) != 2 {
				t.Fatalf("Invalid extended resource string: \"%s\"", extended)
			}
			req[v1.ResourceName(parts[0])] = resource.MustParse(parts[1])
		}
	}
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			UID:       types.UID(objUID),
			Namespace: "node_info_cache_test",
			Name:      objName,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{{
				Resources: v1.ResourceRequirements{
					Requests: req,
				},
				Ports: ports,
			}},
			NodeName: nodeName,
		},
	}
}
