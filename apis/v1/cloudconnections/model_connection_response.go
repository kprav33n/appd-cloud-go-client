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
	"time"
)

// ConnectionResponse struct for ConnectionResponse
type ConnectionResponse struct {
	ConnectionDetail
	Details ConnectionResponseDetails `json:"details"`
}

// NewConnectionResponse instantiates a new ConnectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionResponse(details ConnectionResponseDetails, id string, createdAt time.Time, updatedAt time.Time, state string, stateMessage string, displayName string, type_ ProviderType) *ConnectionResponse {
	this := ConnectionResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.State = state
	this.StateMessage = stateMessage
	this.DisplayName = displayName
	this.Type = type_
	this.Details = details
	return &this
}

// NewConnectionResponseWithDefaults instantiates a new ConnectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionResponseWithDefaults() *ConnectionResponse {
	this := ConnectionResponse{}
	return &this
}

// GetDetails returns the Details field value
func (o *ConnectionResponse) GetDetails() ConnectionResponseDetails {
	if o == nil {
		var ret ConnectionResponseDetails
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *ConnectionResponse) GetDetailsOk() (*ConnectionResponseDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *ConnectionResponse) SetDetails(v ConnectionResponseDetails) {
	o.Details = v
}

func (o ConnectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectionDetail, errConnectionDetail := json.Marshal(o.ConnectionDetail)
	if errConnectionDetail != nil {
		return []byte{}, errConnectionDetail
	}
	errConnectionDetail = json.Unmarshal([]byte(serializedConnectionDetail), &toSerialize)
	if errConnectionDetail != nil {
		return []byte{}, errConnectionDetail
	}
	if true {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionResponse struct {
	value *ConnectionResponse
	isSet bool
}

func (v NullableConnectionResponse) Get() *ConnectionResponse {
	return v.value
}

func (v *NullableConnectionResponse) Set(val *ConnectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionResponse(val *ConnectionResponse) *NullableConnectionResponse {
	return &NullableConnectionResponse{value: val, isSet: true}
}

func (v NullableConnectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}