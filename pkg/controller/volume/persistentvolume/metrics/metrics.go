/*
Copyright 2017 The Kubernetes Authors.

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

package metrics

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/klog"
	metricutil "k8s.io/kubernetes/pkg/volume/util"
)

const (
	// Subsystem names.
	pvControllerSubsystem = "pv_collector"

	// Metric names.
	boundPVKey    = "bound_pv_count"
	unboundPVKey  = "unbound_pv_count"
	boundPVCKey   = "bound_pvc_count"
	unboundPVCKey = "unbound_pvc_count"

	// Label names.
	namespaceLabel    = "namespace"
	storageClassLabel = "storage_class"
)

var registerMetrics sync.Once

// PVLister used to list persistent volumes.
type PVLister interface {
	List() []interface{}
}

// PVCLister used to list persistent volume claims.
type PVCLister interface {
	List() []interface{}
}

// Register all metrics for pv controller.
func Register(pvLister PVLister, pvcLister PVCLister) {
	registerMetrics.Do(func() {
		prometheus.MustRegister(newPVAndPVCCountCollector(pvLister, pvcLister))
		prometheus.MustRegister(volumeOperationErrorsMetric)
	})
}

func newPVAndPVCCountCollector(pvLister PVLister, pvcLister PVCLister) *pvAndPVCCountCollector {
	return &pvAndPVCCountCollector{pvLister, pvcLister}
}

// Custom collector for current pod and container counts.
type pvAndPVCCountCollector struct {
	// Cache for accessing information about PersistentVolumes.
	pvLister PVLister
	// Cache for accessing information about PersistentVolumeClaims.
	pvcLister PVCLister
}

var (
	boundPVCountDesc = prometheus.NewDesc(
		prometheus.BuildFQName("", pvControllerSubsystem, boundPVKey),
		"Gauge measuring number of persistent volume currently bound",
		[]string{storageClassLabel}, nil)
	unboundPVCountDesc = prometheus.NewDesc(
		prometheus.BuildFQName("", pvControllerSubsystem, unboundPVKey),
		"Gauge measuring number of persistent volume currently unbound",
		[]string{storageClassLabel}, nil)

	boundPVCCountDesc = prometheus.NewDesc(
		prometheus.BuildFQName("", pvControllerSubsystem, boundPVCKey),
		"Gauge measuring number of persistent volume claim currently bound",
		[]string{namespaceLabel}, nil)
	unboundPVCCountDesc = prometheus.NewDesc(
		prometheus.BuildFQName("", pvControllerSubsystem, unboundPVCKey),
		"Gauge measuring number of persistent volume claim currently unbound",
		[]string{namespaceLabel}, nil)

	volumeOperationErrorsMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "volume_operation_total_errors",
			Help: "Total volume operation erros",
		},
		[]string{"plugin_name", "operation_name"})
)

func (collector *pvAndPVCCountCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- boundPVCountDesc
	ch <- unboundPVCountDesc
	ch <- boundPVCCountDesc
	ch <- unboundPVCCountDesc
}

func (collector *pvAndPVCCountCollector) Collect(ch chan<- prometheus.Metric) {
	collector.pvCollect(ch)
	collector.pvcCollect(ch)
}

func (collector *pvAndPVCCountCollector) pvCollect(ch chan<- prometheus.Metric) {
	boundNumberByStorageClass := make(map[string]int)
	unboundNumberByStorageClass := make(map[string]int)
	for _, pvObj := range collector.pvLister.List() {
		pv, ok := pvObj.(*v1.PersistentVolume)
		if !ok {
			continue
		}
		if pv.Status.Phase == v1.VolumeBound {
			boundNumberByStorageClass[pv.Spec.StorageClassName]++
		} else {
			unboundNumberByStorageClass[pv.Spec.StorageClassName]++
		}
	}
	for storageClassName, number := range boundNumberByStorageClass {
		metric, err := prometheus.NewConstMetric(
			boundPVCountDesc,
			prometheus.GaugeValue,
			float64(number),
			storageClassName)
		if err != nil {
			klog.Warningf("Create bound pv number metric failed: %v", err)
			continue
		}
		ch <- metric
	}
	for storageClassName, number := range unboundNumberByStorageClass {
		metric, err := prometheus.NewConstMetric(
			unboundPVCountDesc,
			prometheus.GaugeValue,
			float64(number),
			storageClassName)
		if err != nil {
			klog.Warningf("Create unbound pv number metric failed: %v", err)
			continue
		}
		ch <- metric
	}
}

func (collector *pvAndPVCCountCollector) pvcCollect(ch chan<- prometheus.Metric) {
	boundNumberByNamespace := make(map[string]int)
	unboundNumberByNamespace := make(map[string]int)
	for _, pvcObj := range collector.pvcLister.List() {
		pvc, ok := pvcObj.(*v1.PersistentVolumeClaim)
		if !ok {
			continue
		}
		if pvc.Status.Phase == v1.ClaimBound {
			boundNumberByNamespace[pvc.Namespace]++
		} else {
			unboundNumberByNamespace[pvc.Namespace]++
		}
	}
	for namespace, number := range boundNumberByNamespace {
		metric, err := prometheus.NewConstMetric(
			boundPVCCountDesc,
			prometheus.GaugeValue,
			float64(number),
			namespace)
		if err != nil {
			klog.Warningf("Create bound pvc number metric failed: %v", err)
			continue
		}
		ch <- metric
	}
	for namespace, number := range unboundNumberByNamespace {
		metric, err := prometheus.NewConstMetric(
			unboundPVCCountDesc,
			prometheus.GaugeValue,
			float64(number),
			namespace)
		if err != nil {
			klog.Warningf("Create unbound pvc number metric failed: %v", err)
			continue
		}
		ch <- metric
	}
}

// RecordVolumeOperationErrorMetric records error count into metric
// volume_operation_total_errors for provisioning/deletion operations
func RecordVolumeOperationErrorMetric(pluginName, opName string) {
	if pluginName == "" {
		pluginName = "N/A"
	}
	volumeOperationErrorsMetric.WithLabelValues(pluginName, opName).Inc()
}

// operationTimestamp stores the start time of an operation by an operator
type operationTimestamp struct {
	operator  string
	operation string
	startTs   time.Time
}

func newOperationTimestamp(operatorName, operationName string) *operationTimestamp {
	return &operationTimestamp{
		operator:  operatorName,
		operation: operationName,
		startTs:   time.Now(),
	}
}

// OperationStartTimeCache concurrent safe cache for operation start timestamps
type OperationStartTimeCache struct {
	cache sync.Map // [string]operationTimestamp
}

func NewOperationStartTimeCache() OperationStartTimeCache {
	return OperationStartTimeCache{
		cache: sync.Map{}, //[string]operationTimestamp {}
	}
}

// LoadOrStore returns directly if there exists an entry with the key. Otherwise, it
// creates a new operation timestamp using operationName, operatorName, and current timestamp
// and stores the operation timestamp with the key
func (c *OperationStartTimeCache) AddIfNotExist(key, operatorName, operationName string) {
	ts := newOperationTimestamp(operatorName, operationName)
	c.cache.LoadOrStore(key, ts)
}

// Delete deletes a value for a key.
func (c *OperationStartTimeCache) Delete(key string) {
	c.cache.Delete(key)
}

// Store creates an operationTimestamp with current time, and stores it for a key
func (c *OperationStartTimeCache) Store(key, operatorName, operationName string) {
	ts := newOperationTimestamp(operatorName, operationName)
	c.cache.Store(key, ts)
}

// Load returns the value stored in the cache for a key, or nil no value is found
// the ok result will be TRUE only if both the following conditions have been met
// 1. there exists a value in the cache with the key
// 2. the value can be cast into *operatioTimestamp
func (c *OperationStartTimeCache) Load(key string) (*operationTimestamp, bool) {
	obj, exists := c.cache.Load(key)
	if !exists {
		return nil, exists
	}
	ts, ok := obj.(*operationTimestamp)
	return ts, ok
}

// RecordMetric records either an error count metric or a latency metric if there
// exists a start timestamp entry in the cache
func RecordMetric(key string, c *OperationStartTimeCache, err error) {
	ts, exists := c.Load(key)
	if !exists {
		return
	}
	if err != nil {
		RecordVolumeOperationErrorMetric(ts.operator, ts.operation)
	} else {
		timeTaken := time.Since(ts.startTs).Seconds()
		metricutil.RecordOperationLatencyMetric(ts.operator, ts.operation, timeTaken)
	}
}
