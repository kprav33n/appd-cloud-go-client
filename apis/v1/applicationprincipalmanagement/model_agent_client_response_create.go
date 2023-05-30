// Copyright 2023 Cisco Systems, Inc.
// 
// Licensed under the MPL License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     https://www.mozilla.org/en-US/MPL/2.0/
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/*
Application Principal Management Service

Handles all administrative requests to manage application identities, including both Agents and Service principals.

API version: 1.0
Contact: info@appdynamics.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package applicationprincipalmanagement

import (
	"encoding/json"
)

// AgentClientResponseCreate Response object for requests to manage Agent client identities. Includes a generated client secret.
type AgentClientResponseCreate struct {
	// The display name for the client.
	DisplayName *string `json:"displayName,omitempty"`
	// A user provided description of the client.
	Description *string `json:"description,omitempty"`
	// Supported authentication methods used to request oAuth tokens: * `client_secret_basic` - The client credentials will be sent in the authorization header. * `client_secret_post` - The client credentials will be sent in the request body.
	AuthType *string `json:"authType,omitempty"`
	// The unique client identifier.
	Id *string `json:"id,omitempty"`
	// Indicates if the client has rotated secrets. Rotated client secrets can be revoked.
	HasRotatedSecrets *bool `json:"hasRotatedSecrets,omitempty"`
	// The RFC3339 timestamp when the client was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The RFC3339 timestamp when the client was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The client's secret, used to authenticate during an oAuth token request.
	ClientSecret *string `json:"clientSecret,omitempty"`
}

// NewAgentClientResponseCreate instantiates a new AgentClientResponseCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentClientResponseCreate() *AgentClientResponseCreate {
	this := AgentClientResponseCreate{}
	return &this
}

// NewAgentClientResponseCreateWithDefaults instantiates a new AgentClientResponseCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentClientResponseCreateWithDefaults() *AgentClientResponseCreate {
	this := AgentClientResponseCreate{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AgentClientResponseCreate) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientResponseCreate) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AgentClientResponseCreate) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AgentClientResponseCreate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AgentClientResponseCreate) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientResponseCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AgentClientResponseCreate) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AgentClientResponseCreate) SetDescription(v string) {
	o.Description = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *AgentClientResponseCreate) GetAuthType() string {
	if o == nil || isNil(o.AuthType) {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientResponseCreate) GetAuthTypeOk() (*string, bool) {
	if o == nil || isNil(o.AuthType) {
    return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *AgentClientResponseCreate) HasAuthType() bool {
	if o != nil && !isNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *AgentClientResponseCreate) SetAuthType(v string) {
	o.AuthType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgentClientResponseCreate) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientResponseCreate) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgentClientResponseCreate) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgentClientResponseCreate) SetId(v string) {
	o.Id = &v
}

// GetHasRotatedSecrets returns the HasRotatedSecrets field value if set, zero value otherwise.
func (o *AgentClientResponseCreate) GetHasRotatedSecrets() bool {
	if o == nil || isNil(o.HasRotatedSecrets) {
		var ret bool
		return ret
	}
	return *o.HasRotatedSecrets
}

// GetHasRotatedSecretsOk returns a tuple with the HasRotatedSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientResponseCreate) GetHasRotatedSecretsOk() (*bool, bool) {
	if o == nil || isNil(o.HasRotatedSecrets) {
    return nil, false
	}
	return o.HasRotatedSecrets, true
}

// HasHasRotatedSecrets returns a boolean if a field has been set.
func (o *AgentClientResponseCreate) HasHasRotatedSecrets() bool {
	if o != nil && !isNil(o.HasRotatedSecrets) {
		return true
	}

	return false
}

// SetHasRotatedSecrets gets a reference to the given bool and assigns it to the HasRotatedSecrets field.
func (o *AgentClientResponseCreate) SetHasRotatedSecrets(v bool) {
	o.HasRotatedSecrets = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AgentClientResponseCreate) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientResponseCreate) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AgentClientResponseCreate) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AgentClientResponseCreate) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AgentClientResponseCreate) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientResponseCreate) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AgentClientResponseCreate) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AgentClientResponseCreate) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *AgentClientResponseCreate) GetClientSecret() string {
	if o == nil || isNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentClientResponseCreate) GetClientSecretOk() (*string, bool) {
	if o == nil || isNil(o.ClientSecret) {
    return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *AgentClientResponseCreate) HasClientSecret() bool {
	if o != nil && !isNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *AgentClientResponseCreate) SetClientSecret(v string) {
	o.ClientSecret = &v
}

func (o AgentClientResponseCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.AuthType) {
		toSerialize["authType"] = o.AuthType
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.HasRotatedSecrets) {
		toSerialize["hasRotatedSecrets"] = o.HasRotatedSecrets
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	return json.Marshal(toSerialize)
}

type NullableAgentClientResponseCreate struct {
	value *AgentClientResponseCreate
	isSet bool
}

func (v NullableAgentClientResponseCreate) Get() *AgentClientResponseCreate {
	return v.value
}

func (v *NullableAgentClientResponseCreate) Set(val *AgentClientResponseCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentClientResponseCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentClientResponseCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentClientResponseCreate(val *AgentClientResponseCreate) *NullableAgentClientResponseCreate {
	return &NullableAgentClientResponseCreate{value: val, isSet: true}
}

func (v NullableAgentClientResponseCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentClientResponseCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


