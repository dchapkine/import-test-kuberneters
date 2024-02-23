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

package openapi

import (
	"context"
	"net/url"

	"k8s.io/kube-openapi/pkg/handler3"
)

const ContentTypeOpenAPIV3PB = "application/com.github.proto-openapi.spec.v3@v1.0+protobuf"

// HashParamName is the name of the query parameter in the schema URL.
// the schema URL is expected to be like /openapi/v3/apis/apps/v1?hash=014fbff9a07c
const HashParamName = "hash"

type GroupVersion interface {
	Schema(contentType string) ([]byte, error)

	// Hash returns the hash that uniquely identifies the version of the requested schema.
	// It returns an empty string if the hash does not present in the URL to the schema,
	// or an error if the URL fails to parse.
	Hash() (string, error)
}

type groupversion struct {
	client          *client
	item            handler3.OpenAPIV3DiscoveryGroupVersion
	useClientPrefix bool
}

func newGroupVersion(client *client, item handler3.OpenAPIV3DiscoveryGroupVersion, useClientPrefix bool) *groupversion {
	return &groupversion{client: client, item: item, useClientPrefix: useClientPrefix}
}

func (g *groupversion) Schema(contentType string) ([]byte, error) {
	if !g.useClientPrefix {
		return g.client.restClient.Get().
			RequestURI(g.item.ServerRelativeURL).
			SetHeader("Accept", contentType).
			Do(context.TODO()).
			Raw()
	}

	locator, err := url.Parse(g.item.ServerRelativeURL)
	if err != nil {
		return nil, err
	}

	path := g.client.restClient.Get().
		AbsPath(locator.Path).
		SetHeader("Accept", contentType)

	// Other than root endpoints(openapiv3/apis), resources have hash query parameter to support etags.
	// However, absPath does not support handling query parameters internally,
	// so that hash query parameter is added manually
	for k, value := range locator.Query() {
		for _, v := range value {
			path.Param(k, v)
		}
	}

	return path.Do(context.TODO()).Raw()
}

func (g *groupversion) Hash() (string, error) {
	locator, err := url.Parse(g.item.ServerRelativeURL)
	if err != nil {
		return "", err
	}
	if query := locator.Query(); query.Has(HashParamName) {
		return query.Get(HashParamName), nil
	}
	return "", nil
}
