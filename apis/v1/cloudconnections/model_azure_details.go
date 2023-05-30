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
AppDynamics Cloud Connections API

Enables you to manage cloud connections

API version: 1.0.0
Contact: support@appdynamics.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudconnections

import (
	"encoding/json"
)

// AzureDetails struct for AzureDetails
type AzureDetails struct {
	// Client IDs, also known as Application IDs, are long-term credentials for an Azure user, or account root user. The Client ID is one of three properties needed to authenticate to Azure, the other two being Client Secret and Tenant (Directory) ID
	ClientId string `json:"clientId"`
	// A Client Secret allows an Azure application to provide its identity when requesting an access token. The Client Secret is one of three properties needed to authenticate to Azure, the other two being Client ID (Application ID) and Tenant (Directory) ID
	ClientSecret string `json:"clientSecret"`
	// The Azure AD Tenant (Directory) IDis one of three properties needed to authenticate to Azure. The other two are Client Secret and Client ID (Application ID).
	TenantId string `json:"tenantId"`
	// Specify a GUID Subscription ID to monitor. If monitoring all subscriptions, do not specify a Subscription ID.
	SubscriptionId string `json:"subscriptionId"`
}

// NewAzureDetails instantiates a new AzureDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureDetails(clientId string, clientSecret string, tenantId string, subscriptionId string) *AzureDetails {
	this := AzureDetails{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.TenantId = tenantId
	this.SubscriptionId = subscriptionId
	return &this
}

// NewAzureDetailsWithDefaults instantiates a new AzureDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureDetailsWithDefaults() *AzureDetails {
	this := AzureDetails{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *AzureDetails) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *AzureDetails) GetClientIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *AzureDetails) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *AzureDetails) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *AzureDetails) GetClientSecretOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *AzureDetails) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetTenantId returns the TenantId field value
func (o *AzureDetails) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *AzureDetails) GetTenantIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *AzureDetails) SetTenantId(v string) {
	o.TenantId = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *AzureDetails) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *AzureDetails) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *AzureDetails) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o AzureDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clientId"] = o.ClientId
	}
	if true {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	return json.Marshal(toSerialize)
}

type NullableAzureDetails struct {
	value *AzureDetails
	isSet bool
}

func (v NullableAzureDetails) Get() *AzureDetails {
	return v.value
}

func (v *NullableAzureDetails) Set(val *AzureDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureDetails(val *AzureDetails) *NullableAzureDetails {
	return &NullableAzureDetails{value: val, isSet: true}
}

func (v NullableAzureDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


