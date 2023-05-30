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
AppDynamics Cloud Query Service API

An API providing access to observation data using an AppDynamics domain-specific language.  The language is declarative, read-only (it does not allow for data modification) and specific to the AppD MELT model. It presents MELT data (metrics, events, logs, traces) and State in the scope of monitored topology.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudquery

import (
	"encoding/json"
	"time"
)

// ErrorResultChunkError struct for ErrorResultChunkError
type ErrorResultChunkError struct {
	// URI to the documentation of the error.
	Type *string `json:"type,omitempty"`
	// The code of the error.
	ErrorCode *string `json:"errorCode,omitempty"`
	// Time and date the `ErrorResultChunk` was created.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// It correlates events and logs messages across dependent services for single business transaction.
	TraceId *string `json:"traceId,omitempty"`
	// Brief description of the error.
	Title *string `json:"title,omitempty"`
	// Usually more detailed description of the error.
	Detail *string `json:"detail,omitempty"`
	// The input query.
	Query *string `json:"query,omitempty"`
}

// NewErrorResultChunkError instantiates a new ErrorResultChunkError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResultChunkError() *ErrorResultChunkError {
	this := ErrorResultChunkError{}
	return &this
}

// NewErrorResultChunkErrorWithDefaults instantiates a new ErrorResultChunkError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResultChunkErrorWithDefaults() *ErrorResultChunkError {
	this := ErrorResultChunkError{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ErrorResultChunkError) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResultChunkError) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ErrorResultChunkError) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ErrorResultChunkError) SetType(v string) {
	o.Type = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ErrorResultChunkError) GetErrorCode() string {
	if o == nil || isNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResultChunkError) GetErrorCodeOk() (*string, bool) {
	if o == nil || isNil(o.ErrorCode) {
    return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ErrorResultChunkError) HasErrorCode() bool {
	if o != nil && !isNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *ErrorResultChunkError) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ErrorResultChunkError) GetTimestamp() time.Time {
	if o == nil || isNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResultChunkError) GetTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.Timestamp) {
    return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ErrorResultChunkError) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *ErrorResultChunkError) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *ErrorResultChunkError) GetTraceId() string {
	if o == nil || isNil(o.TraceId) {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResultChunkError) GetTraceIdOk() (*string, bool) {
	if o == nil || isNil(o.TraceId) {
    return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *ErrorResultChunkError) HasTraceId() bool {
	if o != nil && !isNil(o.TraceId) {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *ErrorResultChunkError) SetTraceId(v string) {
	o.TraceId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ErrorResultChunkError) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResultChunkError) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ErrorResultChunkError) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ErrorResultChunkError) SetTitle(v string) {
	o.Title = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ErrorResultChunkError) GetDetail() string {
	if o == nil || isNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResultChunkError) GetDetailOk() (*string, bool) {
	if o == nil || isNil(o.Detail) {
    return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ErrorResultChunkError) HasDetail() bool {
	if o != nil && !isNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ErrorResultChunkError) SetDetail(v string) {
	o.Detail = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ErrorResultChunkError) GetQuery() string {
	if o == nil || isNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResultChunkError) GetQueryOk() (*string, bool) {
	if o == nil || isNil(o.Query) {
    return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ErrorResultChunkError) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ErrorResultChunkError) SetQuery(v string) {
	o.Query = &v
}

func (o ErrorResultChunkError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !isNil(o.TraceId) {
		toSerialize["traceId"] = o.TraceId
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !isNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableErrorResultChunkError struct {
	value *ErrorResultChunkError
	isSet bool
}

func (v NullableErrorResultChunkError) Get() *ErrorResultChunkError {
	return v.value
}

func (v *NullableErrorResultChunkError) Set(val *ErrorResultChunkError) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResultChunkError) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResultChunkError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResultChunkError(val *ErrorResultChunkError) *NullableErrorResultChunkError {
	return &NullableErrorResultChunkError{value: val, isSet: true}
}

func (v NullableErrorResultChunkError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResultChunkError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


