/*
Loops OpenAPI Spec

This is the OpenAPI Spec for the [Loops API](https://loops.so/docs/api).

API version: 1.3.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loops

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TransactionalFailure2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionalFailure2Response{}

// TransactionalFailure2Response struct for TransactionalFailure2Response
type TransactionalFailure2Response struct {
	Success bool `json:"success"`
	Error TransactionalFailure2ResponseError `json:"error"`
}

type _TransactionalFailure2Response TransactionalFailure2Response

// NewTransactionalFailure2Response instantiates a new TransactionalFailure2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionalFailure2Response(success bool, error_ TransactionalFailure2ResponseError) *TransactionalFailure2Response {
	this := TransactionalFailure2Response{}
	this.Success = success
	this.Error = error_
	return &this
}

// NewTransactionalFailure2ResponseWithDefaults instantiates a new TransactionalFailure2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionalFailure2ResponseWithDefaults() *TransactionalFailure2Response {
	this := TransactionalFailure2Response{}
	return &this
}

// GetSuccess returns the Success field value
func (o *TransactionalFailure2Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *TransactionalFailure2Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *TransactionalFailure2Response) SetSuccess(v bool) {
	o.Success = v
}

// GetError returns the Error field value
func (o *TransactionalFailure2Response) GetError() TransactionalFailure2ResponseError {
	if o == nil {
		var ret TransactionalFailure2ResponseError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *TransactionalFailure2Response) GetErrorOk() (*TransactionalFailure2ResponseError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *TransactionalFailure2Response) SetError(v TransactionalFailure2ResponseError) {
	o.Error = v
}

func (o TransactionalFailure2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionalFailure2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *TransactionalFailure2Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"error",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTransactionalFailure2Response := _TransactionalFailure2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionalFailure2Response)

	if err != nil {
		return err
	}

	*o = TransactionalFailure2Response(varTransactionalFailure2Response)

	return err
}

type NullableTransactionalFailure2Response struct {
	value *TransactionalFailure2Response
	isSet bool
}

func (v NullableTransactionalFailure2Response) Get() *TransactionalFailure2Response {
	return v.value
}

func (v *NullableTransactionalFailure2Response) Set(val *TransactionalFailure2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionalFailure2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionalFailure2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionalFailure2Response(val *TransactionalFailure2Response) *NullableTransactionalFailure2Response {
	return &NullableTransactionalFailure2Response{value: val, isSet: true}
}

func (v NullableTransactionalFailure2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionalFailure2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


