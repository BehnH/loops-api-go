/*
Loops OpenAPI Spec

This is the OpenAPI Spec for the [Loops API](https://loops.so/docs/api).

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the EventSuccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventSuccessResponse{}

// EventSuccessResponse struct for EventSuccessResponse
type EventSuccessResponse struct {
	Success *bool `json:"success,omitempty"`
}

// NewEventSuccessResponse instantiates a new EventSuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSuccessResponse() *EventSuccessResponse {
	this := EventSuccessResponse{}
	return &this
}

// NewEventSuccessResponseWithDefaults instantiates a new EventSuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSuccessResponseWithDefaults() *EventSuccessResponse {
	this := EventSuccessResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *EventSuccessResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSuccessResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *EventSuccessResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *EventSuccessResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o EventSuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventSuccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableEventSuccessResponse struct {
	value *EventSuccessResponse
	isSet bool
}

func (v NullableEventSuccessResponse) Get() *EventSuccessResponse {
	return v.value
}

func (v *NullableEventSuccessResponse) Set(val *EventSuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSuccessResponse(val *EventSuccessResponse) *NullableEventSuccessResponse {
	return &NullableEventSuccessResponse{value: val, isSet: true}
}

func (v NullableEventSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


