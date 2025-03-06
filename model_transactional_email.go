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

// checks if the TransactionalEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionalEmail{}

// TransactionalEmail struct for TransactionalEmail
type TransactionalEmail struct {
	Id string `json:"id"`
	Name string `json:"name"`
	LastUpdated string `json:"lastUpdated"`
	DataVariables []string `json:"dataVariables"`
}

type _TransactionalEmail TransactionalEmail

// NewTransactionalEmail instantiates a new TransactionalEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionalEmail(id string, name string, lastUpdated string, dataVariables []string) *TransactionalEmail {
	this := TransactionalEmail{}
	this.Id = id
	this.Name = name
	this.LastUpdated = lastUpdated
	this.DataVariables = dataVariables
	return &this
}

// NewTransactionalEmailWithDefaults instantiates a new TransactionalEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionalEmailWithDefaults() *TransactionalEmail {
	this := TransactionalEmail{}
	return &this
}

// GetId returns the Id field value
func (o *TransactionalEmail) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransactionalEmail) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransactionalEmail) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TransactionalEmail) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TransactionalEmail) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TransactionalEmail) SetName(v string) {
	o.Name = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *TransactionalEmail) GetLastUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *TransactionalEmail) GetLastUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *TransactionalEmail) SetLastUpdated(v string) {
	o.LastUpdated = v
}

// GetDataVariables returns the DataVariables field value
func (o *TransactionalEmail) GetDataVariables() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DataVariables
}

// GetDataVariablesOk returns a tuple with the DataVariables field value
// and a boolean to check if the value has been set.
func (o *TransactionalEmail) GetDataVariablesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataVariables, true
}

// SetDataVariables sets field value
func (o *TransactionalEmail) SetDataVariables(v []string) {
	o.DataVariables = v
}

func (o TransactionalEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionalEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["dataVariables"] = o.DataVariables
	return toSerialize, nil
}

func (o *TransactionalEmail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"lastUpdated",
		"dataVariables",
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

	varTransactionalEmail := _TransactionalEmail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionalEmail)

	if err != nil {
		return err
	}

	*o = TransactionalEmail(varTransactionalEmail)

	return err
}

type NullableTransactionalEmail struct {
	value *TransactionalEmail
	isSet bool
}

func (v NullableTransactionalEmail) Get() *TransactionalEmail {
	return v.value
}

func (v *NullableTransactionalEmail) Set(val *TransactionalEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionalEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionalEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionalEmail(val *TransactionalEmail) *NullableTransactionalEmail {
	return &NullableTransactionalEmail{value: val, isSet: true}
}

func (v NullableTransactionalEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionalEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


