# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

kind: ConfigMap
apiVersion: v1
metadata:
  name: dns-autoscaler-nanny-params
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    version: v19
    kubernetes.io/cluster-service: "true"
data:
  params: '{ "cores_to_replicas_map": { "1":1, "4":2, "128":3, "512":5, "1024":7, "2048":10, "4096":15, "8192":20, "12288":30, "16384":40, "20480":50, "24576":60, "28672":70, "32768":80, "65535":100} }'


