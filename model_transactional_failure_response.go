/*
Loops OpenAPI Spec

This is the OpenAPI Spec for the [Loops API](https://loops.so/docs/api).

API version: 1.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loops

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TransactionalFailureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionalFailureResponse{}

// TransactionalFailureResponse struct for TransactionalFailureResponse
type TransactionalFailureResponse struct {
	Success bool `json:"success"`
	Path string `json:"path"`
	Message string `json:"message"`
}

type _TransactionalFailureResponse TransactionalFailureResponse

// NewTransactionalFailureResponse instantiates a new TransactionalFailureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionalFailureResponse(success bool, path string, message string) *TransactionalFailureResponse {
	this := TransactionalFailureResponse{}
	this.Success = success
	this.Path = path
	this.Message = message
	return &this
}

// NewTransactionalFailureResponseWithDefaults instantiates a new TransactionalFailureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionalFailureResponseWithDefaults() *TransactionalFailureResponse {
	this := TransactionalFailureResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *TransactionalFailureResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *TransactionalFailureResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *TransactionalFailureResponse) SetSuccess(v bool) {
	o.Success = v
}

// GetPath returns the Path field value
func (o *TransactionalFailureResponse) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *TransactionalFailureResponse) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *TransactionalFailureResponse) SetPath(v string) {
	o.Path = v
}

// GetMessage returns the Message field value
func (o *TransactionalFailureResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *TransactionalFailureResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *TransactionalFailureResponse) SetMessage(v string) {
	o.Message = v
}

func (o TransactionalFailureResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionalFailureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["path"] = o.Path
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *TransactionalFailureResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"path",
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

	varTransactionalFailureResponse := _TransactionalFailureResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionalFailureResponse)

	if err != nil {
		return err
	}

	*o = TransactionalFailureResponse(varTransactionalFailureResponse)

	return err
}

type NullableTransactionalFailureResponse struct {
	value *TransactionalFailureResponse
	isSet bool
}

func (v NullableTransactionalFailureResponse) Get() *TransactionalFailureResponse {
	return v.value
}

func (v *NullableTransactionalFailureResponse) Set(val *TransactionalFailureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionalFailureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionalFailureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionalFailureResponse(val *TransactionalFailureResponse) *NullableTransactionalFailureResponse {
	return &NullableTransactionalFailureResponse{value: val, isSet: true}
}

func (v NullableTransactionalFailureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionalFailureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


