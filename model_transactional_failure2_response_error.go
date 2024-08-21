/*
Loops OpenAPI Spec

This is the OpenAPI Spec for the [Loops API](https://loops.so/docs/api).

API version: 1.3.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loops

import (
	"encoding/json"
)

// checks if the TransactionalFailure2ResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionalFailure2ResponseError{}

// TransactionalFailure2ResponseError struct for TransactionalFailure2ResponseError
type TransactionalFailure2ResponseError struct {
	Path *string `json:"path,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewTransactionalFailure2ResponseError instantiates a new TransactionalFailure2ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionalFailure2ResponseError() *TransactionalFailure2ResponseError {
	this := TransactionalFailure2ResponseError{}
	return &this
}

// NewTransactionalFailure2ResponseErrorWithDefaults instantiates a new TransactionalFailure2ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionalFailure2ResponseErrorWithDefaults() *TransactionalFailure2ResponseError {
	this := TransactionalFailure2ResponseError{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *TransactionalFailure2ResponseError) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionalFailure2ResponseError) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *TransactionalFailure2ResponseError) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *TransactionalFailure2ResponseError) SetPath(v string) {
	o.Path = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TransactionalFailure2ResponseError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionalFailure2ResponseError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TransactionalFailure2ResponseError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *TransactionalFailure2ResponseError) SetMessage(v string) {
	o.Message = &v
}

func (o TransactionalFailure2ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionalFailure2ResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableTransactionalFailure2ResponseError struct {
	value *TransactionalFailure2ResponseError
	isSet bool
}

func (v NullableTransactionalFailure2ResponseError) Get() *TransactionalFailure2ResponseError {
	return v.value
}

func (v *NullableTransactionalFailure2ResponseError) Set(val *TransactionalFailure2ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionalFailure2ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionalFailure2ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionalFailure2ResponseError(val *TransactionalFailure2ResponseError) *NullableTransactionalFailure2ResponseError {
	return &NullableTransactionalFailure2ResponseError{value: val, isSet: true}
}

func (v NullableTransactionalFailure2ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionalFailure2ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


