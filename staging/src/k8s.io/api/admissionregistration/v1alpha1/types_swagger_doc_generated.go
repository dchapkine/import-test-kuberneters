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

package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_MatchResources = map[string]string{
	"":                     "MatchResources decides whether to run the admission control policy on an object based on whether it meets the match criteria. The exclude rules take precedence over include rules (if a resource matches both, it is excluded)",
	"namespaceSelector":    "NamespaceSelector decides whether to run the admission control policy on an object based on whether the namespace for that object matches the selector. If the object itself is a namespace, the matching is performed on object.metadata.labels. If the object is another cluster scoped resource, it never skips the policy.\n\nFor example, to run the webhook on any objects whose namespace is not associated with \"runlevel\" of \"0\" or \"1\";  you will set the selector as follows: \"namespaceSelector\": {\n  \"matchExpressions\": [\n    {\n      \"key\": \"runlevel\",\n      \"operator\": \"NotIn\",\n      \"values\": [\n        \"0\",\n        \"1\"\n      ]\n    }\n  ]\n}\n\nIf instead you want to only run the policy on any objects whose namespace is associated with the \"environment\" of \"prod\" or \"staging\"; you will set the selector as follows: \"namespaceSelector\": {\n  \"matchExpressions\": [\n    {\n      \"key\": \"environment\",\n      \"operator\": \"In\",\n      \"values\": [\n        \"prod\",\n        \"staging\"\n      ]\n    }\n  ]\n}\n\nSee https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more examples of label selectors.\n\nDefault to the empty LabelSelector, which matches everything.",
	"objectSelector":       "ObjectSelector decides whether to run the validation based on if the object has matching labels. objectSelector is evaluated against both the oldObject and newObject that would be sent to the cel validation, and is considered to match if either object matches the selector. A null object (oldObject in the case of create, or newObject in the case of delete) or an object that cannot have labels (like a DeploymentRollback or a PodProxyOptions object) is not considered to match. Use the object selector only if the webhook is opt-in, because end users may skip the admission webhook by setting the labels. Default to the empty LabelSelector, which matches everything.",
	"resourceRules":        "ResourceRules describes what operations on what resources/subresources the ValidatingAdmissionPolicy matches. The policy cares about an operation if it matches _any_ Rule.",
	"excludeResourceRules": "ExcludeResourceRules describes what operations on what resources/subresources the ValidatingAdmissionPolicy should not care about. The exclude rules take precedence over include rules (if a resource matches both, it is excluded)",
	"matchPolicy":          "matchPolicy defines how the \"MatchResources\" list is used to match incoming requests. Allowed values are \"Exact\" or \"Equivalent\".\n\n- Exact: match a request only if it exactly matches a specified rule. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but \"rules\" only included `apiGroups:[\"apps\"], apiVersions:[\"v1\"], resources: [\"deployments\"]`, a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the ValidatingAdmissionPolicy.\n\n- Equivalent: match a request if modifies a resource listed in rules, even via another API group or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, and \"rules\" only included `apiGroups:[\"apps\"], apiVersions:[\"v1\"], resources: [\"deployments\"]`, a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the ValidatingAdmissionPolicy.\n\nDefaults to \"Equivalent\"",
}

func (MatchResources) SwaggerDoc() map[string]string {
	return map_MatchResources
}

var map_NamedRuleWithOperations = map[string]string{
	"":              "NamedRuleWithOperations is a tuple of Operations and Resources with ResourceNames.",
	"resourceNames": "ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.",
}

func (NamedRuleWithOperations) SwaggerDoc() map[string]string {
	return map_NamedRuleWithOperations
}

var map_ParamKind = map[string]string{
	"":           "ParamKind is a tuple of Group Kind and Version.",
	"apiVersion": "APIVersion is the API group version the resources belong to. In format of \"group/version\". Required.",
	"kind":       "Kind is the API kind the resources belong to. Required.",
}

func (ParamKind) SwaggerDoc() map[string]string {
	return map_ParamKind
}

var map_ParamRef = map[string]string{
	"":          "ParamRef references a parameter resource",
	"name":      "Name of the resource being referenced.",
	"namespace": "Namespace of the referenced resource. Should be empty for the cluster-scoped resources",
}

func (ParamRef) SwaggerDoc() map[string]string {
	return map_ParamRef
}

var map_ValidatingAdmissionPolicy = map[string]string{
	"":         "ValidatingAdmissionPolicy describes the definition of an admission validation policy that accepts or rejects an object without changing it.",
	"metadata": "Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.",
	"spec":     "Specification of the desired behavior of the ValidatingAdmissionPolicy.",
}

func (ValidatingAdmissionPolicy) SwaggerDoc() map[string]string {
	return map_ValidatingAdmissionPolicy
}

var map_ValidatingAdmissionPolicyBinding = map[string]string{
	"":         "ValidatingAdmissionPolicyBinding binds the ValidatingAdmissionPolicy with paramerized resources. ValidatingAdmissionPolicyBinding and parameter CRDs together define how cluster administrators configure policies for clusters.",
	"metadata": "Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.",
	"spec":     "Specification of the desired behavior of the ValidatingAdmissionPolicyBinding.",
}

func (ValidatingAdmissionPolicyBinding) SwaggerDoc() map[string]string {
	return map_ValidatingAdmissionPolicyBinding
}

var map_ValidatingAdmissionPolicyBindingList = map[string]string{
	"":         "ValidatingAdmissionPolicyBindingList is a list of ValidatingAdmissionPolicyBinding.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
	"items":    "List of PolicyBinding.",
}

func (ValidatingAdmissionPolicyBindingList) SwaggerDoc() map[string]string {
	return map_ValidatingAdmissionPolicyBindingList
}

var map_ValidatingAdmissionPolicyBindingSpec = map[string]string{
	"":               "ValidatingAdmissionPolicyBindingSpec is the specification of the ValidatingAdmissionPolicyBinding.",
	"policyName":     "PolicyName references a ValidatingAdmissionPolicy name which the ValidatingAdmissionPolicyBinding binds to. If the referenced resource does not exist, this binding is considered invalid and will be ignored Required.",
	"paramRef":       "ParamRef specifies the parameter resource used to configure the admission control policy. It should point to a resource of the type specified in ParamKind of the bound ValidatingAdmissionPolicy. If the resource referred to by ParamRef does not exist, this binding is considered mis-configured and the FailurePolicy of the policy definition applied.",
	"matchResources": "MatchResources declares what resources match this binding and will be validated by it. Note that this is intersected with the policy's matchConstraints, so only requests that are matched by the policy can be selected by this. If this is unset, all resources matched by the policy are validated by this binding When resourceRules is unset, it does not constrain resource matching. If a resource is matched by the other fields of this object, it will be validated. Note that this is differs from ValidatingAdmissionPolicy matchConstraints, where resourceRules are required.",
}

func (ValidatingAdmissionPolicyBindingSpec) SwaggerDoc() map[string]string {
	return map_ValidatingAdmissionPolicyBindingSpec
}

var map_ValidatingAdmissionPolicyList = map[string]string{
	"":         "ValidatingAdmissionPolicyList is a list of ValidatingAdmissionPolicy.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
	"items":    "List of ValidatingAdmissionPolicy.",
}

func (ValidatingAdmissionPolicyList) SwaggerDoc() map[string]string {
	return map_ValidatingAdmissionPolicyList
}

var map_ValidatingAdmissionPolicySpec = map[string]string{
	"":                 "ValidatingAdmissionPolicySpec is the specification of the desired behavior of the AdmissionPolicy.",
	"paramKind":        "ParamKind specifies the kind of resources used to parameterize this policy. If absent, there are no parameters for this policy and the param CEL variable will not be provided to validation expressions. If ParamKind refers to a non-existent kind, this policy definition is mis-configured and the FailurePolicy is applied.",
	"matchConstraints": "MatchConstraints specifies what resources this policy is designed to validate. The AdmissionPolicy cares about a request if it matches _all_ Constraints. However, in order to prevent clusters from being put into an unstable state that cannot be recovered from via the API ValidatingAdmissionPolicy cannot match ValidatingAdmissionPolicy and ValidatingAdmissionPolicyBinding. Required.",
	"validations":      "Validations contain CEL expressions which is used to apply the validation. A minimum of one validation is required for a policy definition. Required.",
	"failurePolicy":    "FailurePolicy defines how to handle failures for the admission policy. Failures can occur from invalid or mis-configured policy definitions or bindings. A policy is invalid if spec.paramKind refers to a non-existent Kind. A binding is invalid if spec.paramRef.name refers to a non-existent resource. Allowed values are Ignore or Fail. Defaults to Fail.",
}

func (ValidatingAdmissionPolicySpec) SwaggerDoc() map[string]string {
	return map_ValidatingAdmissionPolicySpec
}

var map_Validation = map[string]string{
	"":           "Validation specifies the CEL expression which is used to apply the validation.",
	"expression": "Expression represents the expression which will be evaluated by CEL. ref: https://github.com/google/cel-spec CEL expressions have access to the contents of the Admission request/response, organized into CEL variables as well as some other useful variables:\n\n'object' - The object from the incoming request. The value is null for DELETE requests. 'oldObject' - The existing object. The value is null for CREATE requests. 'request' - Attributes of the admission request([ref](/pkg/apis/admission/types.go#AdmissionRequest)). 'params' - Parameter resource referred to by the policy binding being evaluated. Only populated if the policy has a ParamKind.\n\nThe `apiVersion`, `kind`, `metadata.name` and `metadata.generateName` are always accessible from the root of the object. No other metadata properties are accessible.\n\nOnly property names of the form `[a-zA-Z_.-/][a-zA-Z0-9_.-/]*` are accessible. Accessible property names are escaped according to the following rules when accessed in the expression: - '__' escapes to '__underscores__' - '.' escapes to '__dot__' - '-' escapes to '__dash__' - '/' escapes to '__slash__' - Property names that exactly match a CEL RESERVED keyword escape to '__{keyword}__'. The keywords are:\n\t  \"true\", \"false\", \"null\", \"in\", \"as\", \"break\", \"const\", \"continue\", \"else\", \"for\", \"function\", \"if\",\n\t  \"import\", \"let\", \"loop\", \"package\", \"namespace\", \"return\".\nExamples:\n  - Expression accessing a property named \"namespace\": {\"Expression\": \"object.__namespace__ > 0\"}\n  - Expression accessing a property named \"x-prop\": {\"Expression\": \"object.x__dash__prop > 0\"}\n  - Expression accessing a property named \"redact__d\": {\"Expression\": \"object.redact__underscores__d > 0\"}\n\nEquality on arrays with list type of 'set' or 'map' ignores element order, i.e. [1, 2] == [2, 1]. Concatenation on arrays with x-kubernetes-list-type use the semantics of the list type:\n  - 'set': `X + Y` performs a union where the array positions of all elements in `X` are preserved and\n    non-intersecting elements in `Y` are appended, retaining their partial order.\n  - 'map': `X + Y` performs a merge where the array positions of all keys in `X` are preserved but the values\n    are overwritten by values in `Y` when the key sets of `X` and `Y` intersect. Elements in `Y` with\n    non-intersecting keys are appended, retaining their partial order.\nRequired.",
	"message":    "Message represents the message displayed when validation fails. The message is required if the Expression contains line breaks. The message must not contain line breaks. If unset, the message is \"failed rule: {Rule}\". e.g. \"must be a URL with the host matching spec.host\" If the Expression contains line breaks. Message is required. The message must not contain line breaks. If unset, the message is \"failed Expression: {Expression}\".",
	"reason":     "Reason represents a machine-readable description of why this validation failed. If this is the first validation in the list to fail, this reason, as well as the corresponding HTTP response code, are used in the HTTP response to the client. The currently supported reasons are: \"Unauthorized\", \"Forbidden\", \"Invalid\", \"RequestEntityTooLarge\". If not set, StatusReasonInvalid is used in the response to the client.",
}

func (Validation) SwaggerDoc() map[string]string {
	return map_Validation
}

// AUTO-GENERATED FUNCTIONS END HERE
