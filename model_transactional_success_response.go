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

// checks if the TransactionalSuccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionalSuccessResponse{}

// TransactionalSuccessResponse struct for TransactionalSuccessResponse
type TransactionalSuccessResponse struct {
	Success bool `json:"success"`
}

type _TransactionalSuccessResponse TransactionalSuccessResponse

// NewTransactionalSuccessResponse instantiates a new TransactionalSuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionalSuccessResponse(success bool) *TransactionalSuccessResponse {
	this := TransactionalSuccessResponse{}
	this.Success = success
	return &this
}

// NewTransactionalSuccessResponseWithDefaults instantiates a new TransactionalSuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionalSuccessResponseWithDefaults() *TransactionalSuccessResponse {
	this := TransactionalSuccessResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *TransactionalSuccessResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *TransactionalSuccessResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *TransactionalSuccessResponse) SetSuccess(v bool) {
	o.Success = v
}

func (o TransactionalSuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionalSuccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

func (o *TransactionalSuccessResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
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

	varTransactionalSuccessResponse := _TransactionalSuccessResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionalSuccessResponse)

	if err != nil {
		return err
	}

	*o = TransactionalSuccessResponse(varTransactionalSuccessResponse)

	return err
}

type NullableTransactionalSuccessResponse struct {
	value *TransactionalSuccessResponse
	isSet bool
}

func (v NullableTransactionalSuccessResponse) Get() *TransactionalSuccessResponse {
	return v.value
}

func (v *NullableTransactionalSuccessResponse) Set(val *TransactionalSuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionalSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionalSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionalSuccessResponse(val *TransactionalSuccessResponse) *NullableTransactionalSuccessResponse {
	return &NullableTransactionalSuccessResponse{value: val, isSet: true}
}

func (v NullableTransactionalSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionalSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


