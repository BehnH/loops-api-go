/*
Loops OpenAPI Spec

This is the OpenAPI Spec for the [Loops API](https://loops.so/docs/api).

API version: 1.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loops

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ContactFailureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactFailureResponse{}

// ContactFailureResponse struct for ContactFailureResponse
type ContactFailureResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

type _ContactFailureResponse ContactFailureResponse

// NewContactFailureResponse instantiates a new ContactFailureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactFailureResponse(success bool, message string) *ContactFailureResponse {
	this := ContactFailureResponse{}
	this.Success = success
	this.Message = message
	return &this
}

// NewContactFailureResponseWithDefaults instantiates a new ContactFailureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactFailureResponseWithDefaults() *ContactFailureResponse {
	this := ContactFailureResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *ContactFailureResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ContactFailureResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ContactFailureResponse) SetSuccess(v bool) {
	o.Success = v
}

// GetMessage returns the Message field value
func (o *ContactFailureResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ContactFailureResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ContactFailureResponse) SetMessage(v string) {
	o.Message = v
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
	toSerialize["success"] = o.Success
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *ContactFailureResponse) UnmarshalJSON(data []byte) (err error) {
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

	varContactFailureResponse := _ContactFailureResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContactFailureResponse)

	if err != nil {
		return err
	}

	*o = ContactFailureResponse(varContactFailureResponse)

	return err
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


