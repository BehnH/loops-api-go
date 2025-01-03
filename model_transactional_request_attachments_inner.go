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

// checks if the TransactionalRequestAttachmentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionalRequestAttachmentsInner{}

// TransactionalRequestAttachmentsInner struct for TransactionalRequestAttachmentsInner
type TransactionalRequestAttachmentsInner struct {
	// The name of the file, shown in email clients.
	Filename string `json:"filename"`
	// The MIME type of the file.
	ContentType string `json:"contentType"`
	// The base64-encoded content of the file.
	Data string `json:"data"`
}

type _TransactionalRequestAttachmentsInner TransactionalRequestAttachmentsInner

// NewTransactionalRequestAttachmentsInner instantiates a new TransactionalRequestAttachmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionalRequestAttachmentsInner(filename string, contentType string, data string) *TransactionalRequestAttachmentsInner {
	this := TransactionalRequestAttachmentsInner{}
	this.Filename = filename
	this.ContentType = contentType
	this.Data = data
	return &this
}

// NewTransactionalRequestAttachmentsInnerWithDefaults instantiates a new TransactionalRequestAttachmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionalRequestAttachmentsInnerWithDefaults() *TransactionalRequestAttachmentsInner {
	this := TransactionalRequestAttachmentsInner{}
	return &this
}

// GetFilename returns the Filename field value
func (o *TransactionalRequestAttachmentsInner) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *TransactionalRequestAttachmentsInner) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *TransactionalRequestAttachmentsInner) SetFilename(v string) {
	o.Filename = v
}

// GetContentType returns the ContentType field value
func (o *TransactionalRequestAttachmentsInner) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *TransactionalRequestAttachmentsInner) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *TransactionalRequestAttachmentsInner) SetContentType(v string) {
	o.ContentType = v
}

// GetData returns the Data field value
func (o *TransactionalRequestAttachmentsInner) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransactionalRequestAttachmentsInner) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TransactionalRequestAttachmentsInner) SetData(v string) {
	o.Data = v
}

func (o TransactionalRequestAttachmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionalRequestAttachmentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filename"] = o.Filename
	toSerialize["contentType"] = o.ContentType
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *TransactionalRequestAttachmentsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filename",
		"contentType",
		"data",
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

	varTransactionalRequestAttachmentsInner := _TransactionalRequestAttachmentsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionalRequestAttachmentsInner)

	if err != nil {
		return err
	}

	*o = TransactionalRequestAttachmentsInner(varTransactionalRequestAttachmentsInner)

	return err
}

type NullableTransactionalRequestAttachmentsInner struct {
	value *TransactionalRequestAttachmentsInner
	isSet bool
}

func (v NullableTransactionalRequestAttachmentsInner) Get() *TransactionalRequestAttachmentsInner {
	return v.value
}

func (v *NullableTransactionalRequestAttachmentsInner) Set(val *TransactionalRequestAttachmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionalRequestAttachmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionalRequestAttachmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionalRequestAttachmentsInner(val *TransactionalRequestAttachmentsInner) *NullableTransactionalRequestAttachmentsInner {
	return &NullableTransactionalRequestAttachmentsInner{value: val, isSet: true}
}

func (v NullableTransactionalRequestAttachmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionalRequestAttachmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


