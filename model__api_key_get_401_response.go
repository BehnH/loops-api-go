/*
Loops OpenAPI Spec

This is the OpenAPI Spec for the [Loops API](https://loops.so/docs/api).

API version: 1.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loops

import (
	"encoding/json"
)

// checks if the ApiKeyGet401Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyGet401Response{}

// ApiKeyGet401Response struct for ApiKeyGet401Response
type ApiKeyGet401Response struct {
	Error *string `json:"error,omitempty"`
}

// NewApiKeyGet401Response instantiates a new ApiKeyGet401Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyGet401Response() *ApiKeyGet401Response {
	this := ApiKeyGet401Response{}
	return &this
}

// NewApiKeyGet401ResponseWithDefaults instantiates a new ApiKeyGet401Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyGet401ResponseWithDefaults() *ApiKeyGet401Response {
	this := ApiKeyGet401Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ApiKeyGet401Response) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyGet401Response) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ApiKeyGet401Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ApiKeyGet401Response) SetError(v string) {
	o.Error = &v
}

func (o ApiKeyGet401Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyGet401Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableApiKeyGet401Response struct {
	value *ApiKeyGet401Response
	isSet bool
}

func (v NullableApiKeyGet401Response) Get() *ApiKeyGet401Response {
	return v.value
}

func (v *NullableApiKeyGet401Response) Set(val *ApiKeyGet401Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyGet401Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyGet401Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyGet401Response(val *ApiKeyGet401Response) *NullableApiKeyGet401Response {
	return &NullableApiKeyGet401Response{value: val, isSet: true}
}

func (v NullableApiKeyGet401Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyGet401Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


