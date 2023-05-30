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

// ServiceClientRequest Request object for requests to manage service clients.
type ServiceClientRequest struct {
	// The display name for the client.
	DisplayName *string `json:"displayName,omitempty"`
	// A user provided description of the client.
	Description *string `json:"description,omitempty"`
	// Supported authentication methods used to request oAuth tokens: * `client_secret_basic` - The client credentials will be sent in the authorization header. * `client_secret_post` - The client credentials will be sent in the request body.
	AuthType *string `json:"authType,omitempty"`
}

// NewServiceClientRequest instantiates a new ServiceClientRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceClientRequest() *ServiceClientRequest {
	this := ServiceClientRequest{}
	return &this
}

// NewServiceClientRequestWithDefaults instantiates a new ServiceClientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceClientRequestWithDefaults() *ServiceClientRequest {
	this := ServiceClientRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ServiceClientRequest) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceClientRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ServiceClientRequest) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ServiceClientRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceClientRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceClientRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceClientRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceClientRequest) SetDescription(v string) {
	o.Description = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *ServiceClientRequest) GetAuthType() string {
	if o == nil || isNil(o.AuthType) {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceClientRequest) GetAuthTypeOk() (*string, bool) {
	if o == nil || isNil(o.AuthType) {
    return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *ServiceClientRequest) HasAuthType() bool {
	if o != nil && !isNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *ServiceClientRequest) SetAuthType(v string) {
	o.AuthType = &v
}

func (o ServiceClientRequest) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableServiceClientRequest struct {
	value *ServiceClientRequest
	isSet bool
}

func (v NullableServiceClientRequest) Get() *ServiceClientRequest {
	return v.value
}

func (v *NullableServiceClientRequest) Set(val *ServiceClientRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceClientRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceClientRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceClientRequest(val *ServiceClientRequest) *NullableServiceClientRequest {
	return &NullableServiceClientRequest{value: val, isSet: true}
}

func (v NullableServiceClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceClientRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


