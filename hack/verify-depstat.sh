#!/usr/bin/env bash

# Copyright 2021 The Kubernetes Authors.
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

# This script checks if the dependency-stats.json file needs to be
# updated or not. If yes then please run `hack/update-depstat.sh`
# to update.
# Usage: `hack/verify-depstat.sh`

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${KUBE_ROOT}/hack/lib/init.sh"
source "${KUBE_ROOT}/hack/lib/util.sh"

kube::golang::verify_go_version
kube::golang::setup_env

# Explicitly opt into go modules, even though we're inside a GOPATH directory
export GO111MODULE=on

echo 'Installing depstat'
pushd "${KUBE_ROOT}/hack/tools" >/dev/null
  go install github.com/kubernetes-sigs/depstat
popd >/dev/null

echo 'Verifying depstat'
diff -s --ignore-all-space "${KUBE_ROOT}/hack/dependency-stats.json" <(depstat stats --json)
