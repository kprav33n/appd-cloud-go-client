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

// ConnectionUpdateDetailsOneOf2 struct for ConnectionUpdateDetailsOneOf2
type ConnectionUpdateDetailsOneOf2 struct {
	ClientId string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// NewConnectionUpdateDetailsOneOf2 instantiates a new ConnectionUpdateDetailsOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionUpdateDetailsOneOf2(clientId string, clientSecret string) *ConnectionUpdateDetailsOneOf2 {
	this := ConnectionUpdateDetailsOneOf2{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	return &this
}

// NewConnectionUpdateDetailsOneOf2WithDefaults instantiates a new ConnectionUpdateDetailsOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionUpdateDetailsOneOf2WithDefaults() *ConnectionUpdateDetailsOneOf2 {
	this := ConnectionUpdateDetailsOneOf2{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *ConnectionUpdateDetailsOneOf2) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ConnectionUpdateDetailsOneOf2) GetClientIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ConnectionUpdateDetailsOneOf2) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *ConnectionUpdateDetailsOneOf2) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *ConnectionUpdateDetailsOneOf2) GetClientSecretOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *ConnectionUpdateDetailsOneOf2) SetClientSecret(v string) {
	o.ClientSecret = v
}

func (o ConnectionUpdateDetailsOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clientId"] = o.ClientId
	}
	if true {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionUpdateDetailsOneOf2 struct {
	value *ConnectionUpdateDetailsOneOf2
	isSet bool
}

func (v NullableConnectionUpdateDetailsOneOf2) Get() *ConnectionUpdateDetailsOneOf2 {
	return v.value
}

func (v *NullableConnectionUpdateDetailsOneOf2) Set(val *ConnectionUpdateDetailsOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionUpdateDetailsOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionUpdateDetailsOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionUpdateDetailsOneOf2(val *ConnectionUpdateDetailsOneOf2) *NullableConnectionUpdateDetailsOneOf2 {
	return &NullableConnectionUpdateDetailsOneOf2{value: val, isSet: true}
}

func (v NullableConnectionUpdateDetailsOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionUpdateDetailsOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

