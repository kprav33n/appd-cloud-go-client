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

// ServiceClientsContainer Generic response object which contains an array of ServiceClientResponse.
type ServiceClientsContainer struct {
	Items []ServiceClientResponse `json:"items,omitempty"`
}

// NewServiceClientsContainer instantiates a new ServiceClientsContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceClientsContainer() *ServiceClientsContainer {
	this := ServiceClientsContainer{}
	return &this
}

// NewServiceClientsContainerWithDefaults instantiates a new ServiceClientsContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceClientsContainerWithDefaults() *ServiceClientsContainer {
	this := ServiceClientsContainer{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ServiceClientsContainer) GetItems() []ServiceClientResponse {
	if o == nil || isNil(o.Items) {
		var ret []ServiceClientResponse
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceClientsContainer) GetItemsOk() ([]ServiceClientResponse, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ServiceClientsContainer) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ServiceClientResponse and assigns it to the Items field.
func (o *ServiceClientsContainer) SetItems(v []ServiceClientResponse) {
	o.Items = v
}

func (o ServiceClientsContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableServiceClientsContainer struct {
	value *ServiceClientsContainer
	isSet bool
}

func (v NullableServiceClientsContainer) Get() *ServiceClientsContainer {
	return v.value
}

func (v *NullableServiceClientsContainer) Set(val *ServiceClientsContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceClientsContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceClientsContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceClientsContainer(val *ServiceClientsContainer) *NullableServiceClientsContainer {
	return &NullableServiceClientsContainer{value: val, isSet: true}
}

func (v NullableServiceClientsContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceClientsContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


