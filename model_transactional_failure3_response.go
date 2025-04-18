/*
Loops OpenAPI Spec

This is the OpenAPI Spec for the [Loops API](https://loops.so/docs/api).

API version: 1.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loops

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TransactionalFailure3Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionalFailure3Response{}

// TransactionalFailure3Response struct for TransactionalFailure3Response
type TransactionalFailure3Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

type _TransactionalFailure3Response TransactionalFailure3Response

// NewTransactionalFailure3Response instantiates a new TransactionalFailure3Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionalFailure3Response(success bool, message string) *TransactionalFailure3Response {
	this := TransactionalFailure3Response{}
	this.Success = success
	this.Message = message
	return &this
}

// NewTransactionalFailure3ResponseWithDefaults instantiates a new TransactionalFailure3Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionalFailure3ResponseWithDefaults() *TransactionalFailure3Response {
	this := TransactionalFailure3Response{}
	return &this
}

// GetSuccess returns the Success field value
func (o *TransactionalFailure3Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *TransactionalFailure3Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *TransactionalFailure3Response) SetSuccess(v bool) {
	o.Success = v
}

// GetMessage returns the Message field value
func (o *TransactionalFailure3Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *TransactionalFailure3Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *TransactionalFailure3Response) SetMessage(v string) {
	o.Message = v
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
	toSerialize["success"] = o.Success
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *TransactionalFailure3Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"message",
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

	varTransactionalFailure3Response := _TransactionalFailure3Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionalFailure3Response)

	if err != nil {
		return err
	}

	*o = TransactionalFailure3Response(varTransactionalFailure3Response)

	return err
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


