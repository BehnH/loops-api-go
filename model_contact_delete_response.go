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

// checks if the ContactDeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactDeleteResponse{}

// ContactDeleteResponse struct for ContactDeleteResponse
type ContactDeleteResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

type _ContactDeleteResponse ContactDeleteResponse

// NewContactDeleteResponse instantiates a new ContactDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactDeleteResponse(success bool, message string) *ContactDeleteResponse {
	this := ContactDeleteResponse{}
	this.Success = success
	this.Message = message
	return &this
}

// NewContactDeleteResponseWithDefaults instantiates a new ContactDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactDeleteResponseWithDefaults() *ContactDeleteResponse {
	this := ContactDeleteResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *ContactDeleteResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ContactDeleteResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ContactDeleteResponse) SetSuccess(v bool) {
	o.Success = v
}

// GetMessage returns the Message field value
func (o *ContactDeleteResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ContactDeleteResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ContactDeleteResponse) SetMessage(v string) {
	o.Message = v
}

func (o ContactDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactDeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *ContactDeleteResponse) UnmarshalJSON(data []byte) (err error) {
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

	varContactDeleteResponse := _ContactDeleteResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContactDeleteResponse)

	if err != nil {
		return err
	}

	*o = ContactDeleteResponse(varContactDeleteResponse)

	return err
}

type NullableContactDeleteResponse struct {
	value *ContactDeleteResponse
	isSet bool
}

func (v NullableContactDeleteResponse) Get() *ContactDeleteResponse {
	return v.value
}

func (v *NullableContactDeleteResponse) Set(val *ContactDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContactDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContactDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactDeleteResponse(val *ContactDeleteResponse) *NullableContactDeleteResponse {
	return &NullableContactDeleteResponse{value: val, isSet: true}
}

func (v NullableContactDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


