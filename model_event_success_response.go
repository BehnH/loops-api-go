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

// checks if the EventSuccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventSuccessResponse{}

// EventSuccessResponse struct for EventSuccessResponse
type EventSuccessResponse struct {
	Success bool `json:"success"`
}

type _EventSuccessResponse EventSuccessResponse

// NewEventSuccessResponse instantiates a new EventSuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSuccessResponse(success bool) *EventSuccessResponse {
	this := EventSuccessResponse{}
	this.Success = success
	return &this
}

// NewEventSuccessResponseWithDefaults instantiates a new EventSuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSuccessResponseWithDefaults() *EventSuccessResponse {
	this := EventSuccessResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *EventSuccessResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *EventSuccessResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *EventSuccessResponse) SetSuccess(v bool) {
	o.Success = v
}

func (o EventSuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventSuccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

func (o *EventSuccessResponse) UnmarshalJSON(data []byte) (err error) {
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

	varEventSuccessResponse := _EventSuccessResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventSuccessResponse)

	if err != nil {
		return err
	}

	*o = EventSuccessResponse(varEventSuccessResponse)

	return err
}

type NullableEventSuccessResponse struct {
	value *EventSuccessResponse
	isSet bool
}

func (v NullableEventSuccessResponse) Get() *EventSuccessResponse {
	return v.value
}

func (v *NullableEventSuccessResponse) Set(val *EventSuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSuccessResponse(val *EventSuccessResponse) *NullableEventSuccessResponse {
	return &NullableEventSuccessResponse{value: val, isSet: true}
}

func (v NullableEventSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


