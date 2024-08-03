/*
Loops OpenAPI Spec

This is the OpenAPI Spec for the [Loops API](https://loops.so/docs/api).

API version: 1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loops

import (
	"encoding/json"
)

// checks if the ContactFailureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactFailureResponse{}

// ContactFailureResponse struct for ContactFailureResponse
type ContactFailureResponse struct {
	Success *bool `json:"success,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewContactFailureResponse instantiates a new ContactFailureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactFailureResponse() *ContactFailureResponse {
	this := ContactFailureResponse{}
	return &this
}

// NewContactFailureResponseWithDefaults instantiates a new ContactFailureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactFailureResponseWithDefaults() *ContactFailureResponse {
	this := ContactFailureResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ContactFailureResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactFailureResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ContactFailureResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ContactFailureResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ContactFailureResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactFailureResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ContactFailureResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ContactFailureResponse) SetMessage(v string) {
	o.Message = &v
}

func (o ContactFailureResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactFailureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableContactFailureResponse struct {
	value *ContactFailureResponse
	isSet bool
}

func (v NullableContactFailureResponse) Get() *ContactFailureResponse {
	return v.value
}

func (v *NullableContactFailureResponse) Set(val *ContactFailureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContactFailureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContactFailureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactFailureResponse(val *ContactFailureResponse) *NullableContactFailureResponse {
	return &NullableContactFailureResponse{value: val, isSet: true}
}

func (v NullableContactFailureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactFailureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


