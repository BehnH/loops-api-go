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

// checks if the TransactionalFailure3Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionalFailure3Response{}

// TransactionalFailure3Response struct for TransactionalFailure3Response
type TransactionalFailure3Response struct {
	Success *bool `json:"success,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewTransactionalFailure3Response instantiates a new TransactionalFailure3Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionalFailure3Response() *TransactionalFailure3Response {
	this := TransactionalFailure3Response{}
	return &this
}

// NewTransactionalFailure3ResponseWithDefaults instantiates a new TransactionalFailure3Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionalFailure3ResponseWithDefaults() *TransactionalFailure3Response {
	this := TransactionalFailure3Response{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *TransactionalFailure3Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionalFailure3Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *TransactionalFailure3Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *TransactionalFailure3Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TransactionalFailure3Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionalFailure3Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TransactionalFailure3Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *TransactionalFailure3Response) SetMessage(v string) {
	o.Message = &v
}

func (o TransactionalFailure3Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionalFailure3Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableTransactionalFailure3Response struct {
	value *TransactionalFailure3Response
	isSet bool
}

func (v NullableTransactionalFailure3Response) Get() *TransactionalFailure3Response {
	return v.value
}

func (v *NullableTransactionalFailure3Response) Set(val *TransactionalFailure3Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionalFailure3Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionalFailure3Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionalFailure3Response(val *TransactionalFailure3Response) *NullableTransactionalFailure3Response {
	return &NullableTransactionalFailure3Response{value: val, isSet: true}
}

func (v NullableTransactionalFailure3Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionalFailure3Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


