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

// checks if the ContactDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactDeleteRequest{}

// ContactDeleteRequest struct for ContactDeleteRequest
type ContactDeleteRequest struct {
	Email string `json:"email"`
	UserId string `json:"userId"`
}

type _ContactDeleteRequest ContactDeleteRequest

// NewContactDeleteRequest instantiates a new ContactDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactDeleteRequest(email string, userId string) *ContactDeleteRequest {
	this := ContactDeleteRequest{}
	this.Email = email
	this.UserId = userId
	return &this
}

// NewContactDeleteRequestWithDefaults instantiates a new ContactDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactDeleteRequestWithDefaults() *ContactDeleteRequest {
	this := ContactDeleteRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *ContactDeleteRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ContactDeleteRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ContactDeleteRequest) SetEmail(v string) {
	o.Email = v
}

// GetUserId returns the UserId field value
func (o *ContactDeleteRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ContactDeleteRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ContactDeleteRequest) SetUserId(v string) {
	o.UserId = v
}

func (o ContactDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *ContactDeleteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"userId",
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

	varContactDeleteRequest := _ContactDeleteRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContactDeleteRequest)

	if err != nil {
		return err
	}

	*o = ContactDeleteRequest(varContactDeleteRequest)

	return err
}

type NullableContactDeleteRequest struct {
	value *ContactDeleteRequest
	isSet bool
}

func (v NullableContactDeleteRequest) Get() *ContactDeleteRequest {
	return v.value
}

func (v *NullableContactDeleteRequest) Set(val *ContactDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContactDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContactDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactDeleteRequest(val *ContactDeleteRequest) *NullableContactDeleteRequest {
	return &NullableContactDeleteRequest{value: val, isSet: true}
}

func (v NullableContactDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


