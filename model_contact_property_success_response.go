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

// checks if the ContactPropertySuccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactPropertySuccessResponse{}

// ContactPropertySuccessResponse struct for ContactPropertySuccessResponse
type ContactPropertySuccessResponse struct {
	Success bool `json:"success"`
}

type _ContactPropertySuccessResponse ContactPropertySuccessResponse

// NewContactPropertySuccessResponse instantiates a new ContactPropertySuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactPropertySuccessResponse(success bool) *ContactPropertySuccessResponse {
	this := ContactPropertySuccessResponse{}
	this.Success = success
	return &this
}

// NewContactPropertySuccessResponseWithDefaults instantiates a new ContactPropertySuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactPropertySuccessResponseWithDefaults() *ContactPropertySuccessResponse {
	this := ContactPropertySuccessResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *ContactPropertySuccessResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ContactPropertySuccessResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ContactPropertySuccessResponse) SetSuccess(v bool) {
	o.Success = v
}

func (o ContactPropertySuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactPropertySuccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

func (o *ContactPropertySuccessResponse) UnmarshalJSON(data []byte) (err error) {
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

	varContactPropertySuccessResponse := _ContactPropertySuccessResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContactPropertySuccessResponse)

	if err != nil {
		return err
	}

	*o = ContactPropertySuccessResponse(varContactPropertySuccessResponse)

	return err
}

type NullableContactPropertySuccessResponse struct {
	value *ContactPropertySuccessResponse
	isSet bool
}

func (v NullableContactPropertySuccessResponse) Get() *ContactPropertySuccessResponse {
	return v.value
}

func (v *NullableContactPropertySuccessResponse) Set(val *ContactPropertySuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContactPropertySuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContactPropertySuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactPropertySuccessResponse(val *ContactPropertySuccessResponse) *NullableContactPropertySuccessResponse {
	return &NullableContactPropertySuccessResponse{value: val, isSet: true}
}

func (v NullableContactPropertySuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactPropertySuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


