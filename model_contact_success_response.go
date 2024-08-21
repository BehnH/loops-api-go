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

// checks if the ContactSuccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactSuccessResponse{}

// ContactSuccessResponse struct for ContactSuccessResponse
type ContactSuccessResponse struct {
	Success bool `json:"success"`
	Id string `json:"id"`
}

type _ContactSuccessResponse ContactSuccessResponse

// NewContactSuccessResponse instantiates a new ContactSuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactSuccessResponse(success bool, id string) *ContactSuccessResponse {
	this := ContactSuccessResponse{}
	this.Success = success
	this.Id = id
	return &this
}

// NewContactSuccessResponseWithDefaults instantiates a new ContactSuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactSuccessResponseWithDefaults() *ContactSuccessResponse {
	this := ContactSuccessResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *ContactSuccessResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ContactSuccessResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ContactSuccessResponse) SetSuccess(v bool) {
	o.Success = v
}

// GetId returns the Id field value
func (o *ContactSuccessResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContactSuccessResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContactSuccessResponse) SetId(v string) {
	o.Id = v
}

func (o ContactSuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactSuccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ContactSuccessResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"id",
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

	varContactSuccessResponse := _ContactSuccessResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContactSuccessResponse)

	if err != nil {
		return err
	}

	*o = ContactSuccessResponse(varContactSuccessResponse)

	return err
}

type NullableContactSuccessResponse struct {
	value *ContactSuccessResponse
	isSet bool
}

func (v NullableContactSuccessResponse) Get() *ContactSuccessResponse {
	return v.value
}

func (v *NullableContactSuccessResponse) Set(val *ContactSuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContactSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContactSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactSuccessResponse(val *ContactSuccessResponse) *NullableContactSuccessResponse {
	return &NullableContactSuccessResponse{value: val, isSet: true}
}

func (v NullableContactSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


