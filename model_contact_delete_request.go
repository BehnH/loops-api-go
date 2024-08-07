/*
Loops OpenAPI Spec

This is the OpenAPI Spec for the [Loops API](https://loops.so/docs/api).

API version: 1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loops

import (
	"encoding/json"
)

// checks if the ContactDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactDeleteRequest{}

// ContactDeleteRequest struct for ContactDeleteRequest
type ContactDeleteRequest struct {
	Email *string `json:"email,omitempty"`
	UserId *string `json:"userId,omitempty"`
}

// NewContactDeleteRequest instantiates a new ContactDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactDeleteRequest() *ContactDeleteRequest {
	this := ContactDeleteRequest{}
	return &this
}

// NewContactDeleteRequestWithDefaults instantiates a new ContactDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactDeleteRequestWithDefaults() *ContactDeleteRequest {
	this := ContactDeleteRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ContactDeleteRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactDeleteRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ContactDeleteRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ContactDeleteRequest) SetEmail(v string) {
	o.Email = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ContactDeleteRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactDeleteRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ContactDeleteRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ContactDeleteRequest) SetUserId(v string) {
	o.UserId = &v
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
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
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


