// Copyright (C) 2015 The Gravitee team (http://gravitee.io)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +kubebuilder:object:generate=true
package base

type ApiBase struct {
	// The API ID. If empty, this field will take the value of the `metadata.uid`
	// field of the resource.
	ID string `json:"id,omitempty"`
	// When promoting an API from one environment to the other,
	// this ID identifies the API across those different environments.
	// Setting this ID also allows to take control over an existing API on an APIM instance
	// (by setting the same value as defined in APIM).
	// If empty, a UUID will be generated based on the namespace and name of the resource.
	CrossID string `json:"crossId,omitempty"`
	// API name
	Name string `json:"name,omitempty"`
	// +kubebuilder:validation:Required
	Version string `json:"version,omitempty"`
	// API description
	Description string `json:"description,omitempty"`
	// +kubebuilder:default:=`STARTED`
	// +kubebuilder:validation:Enum=STARTED;STOPPED;
	// The state of API (setting the value to `STOPPED` will make the API un-reachable from the gateway)
	State string `json:"state,omitempty"`
	// +kubebuilder:default:=`CREATED`
	// API life cycle state can be one of the values CREATED, PUBLISHED, UNPUBLISHED, DEPRECATED, ARCHIVED
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// List of Tags of the API
	Tags []string `json:"tags,omitempty"`
	// List of labels of the API
	Labels []string `json:"labels,omitempty"`
	// +kubebuilder:default:=PRIVATE
	// Should the API be publicly available from the portal or not ?
	Visibility ApiVisibility `json:"visibility,omitempty"`
	// Specify the primary member that owns the API
	PrimaryOwner *Member `json:"primaryOwner,omitempty"`
	// +kubebuilder:default:={}
	// List of Properties for the API
	Properties []*Property `json:"properties,omitempty"`
	// +kubebuilder:default:={}
	// List of API metadata entries
	Metadata []*MetadataEntry `json:"metadata,omitempty"`
	// A list of Response Templates for the API
	ResponseTemplates map[string]map[string]*ResponseTemplate `json:"response_templates,omitempty"`
	// +kubebuilder:default:={}
	// Resources can be either inlined or reference the namespace and name
	// of an <a href="#apiresource">existing API resource definition</a>.
	Resources []*ResourceOrRef `json:"resources,omitempty"`
}

// +kubebuilder:validation:Enum=PUBLIC;PRIVATE;
type ApiVisibility string

type DefinitionVersion string

const (
	DefinitionVersionV1 DefinitionVersion = "1.0.0"
	DefinitionVersionV2 DefinitionVersion = "2.0.0"
	DefinitionVersionV4 DefinitionVersion = "V4"
	GatewayDefinitionV4 DefinitionVersion = "4.0.0"
)

// +kubebuilder:validation:Enum=CREATED;PUBLISHED;UNPUBLISHED;DEPRECATED;ARCHIVED;
type LifecycleState string

const (
	StateStarted string = "STARTED"
	StateStopped string = "STOPPED"
)

type ResponseTemplate struct {
	// Response Template status code
	StatusCode int `json:"status,omitempty"`

	// Response Template headers, arbitrary map of string key-value headers
	Headers map[string]string `json:"headers,omitempty"`

	// Response Template body
	Body string `json:"body,omitempty"`
}
