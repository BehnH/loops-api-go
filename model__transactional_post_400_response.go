/*
Loops OpenAPI Spec

This is the OpenAPI Spec for the [Loops API](https://loops.so/docs/api).

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loops

import (
	"encoding/json"
	"fmt"
)

// TransactionalPost400Response - struct for TransactionalPost400Response
type TransactionalPost400Response struct {
	TransactionalFailure2Response *TransactionalFailure2Response
	TransactionalFailure3Response *TransactionalFailure3Response
	TransactionalFailureResponse *TransactionalFailureResponse
}

// TransactionalFailure2ResponseAsTransactionalPost400Response is a convenience function that returns TransactionalFailure2Response wrapped in TransactionalPost400Response
func TransactionalFailure2ResponseAsTransactionalPost400Response(v *TransactionalFailure2Response) TransactionalPost400Response {
	return TransactionalPost400Response{
		TransactionalFailure2Response: v,
	}
}

// TransactionalFailure3ResponseAsTransactionalPost400Response is a convenience function that returns TransactionalFailure3Response wrapped in TransactionalPost400Response
func TransactionalFailure3ResponseAsTransactionalPost400Response(v *TransactionalFailure3Response) TransactionalPost400Response {
	return TransactionalPost400Response{
		TransactionalFailure3Response: v,
	}
}

// TransactionalFailureResponseAsTransactionalPost400Response is a convenience function that returns TransactionalFailureResponse wrapped in TransactionalPost400Response
func TransactionalFailureResponseAsTransactionalPost400Response(v *TransactionalFailureResponse) TransactionalPost400Response {
	return TransactionalPost400Response{
		TransactionalFailureResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransactionalPost400Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TransactionalFailure2Response
	err = newStrictDecoder(data).Decode(&dst.TransactionalFailure2Response)
	if err == nil {
		jsonTransactionalFailure2Response, _ := json.Marshal(dst.TransactionalFailure2Response)
		if string(jsonTransactionalFailure2Response) == "{}" { // empty struct
			dst.TransactionalFailure2Response = nil
		} else {
			match++
		}
	} else {
		dst.TransactionalFailure2Response = nil
	}

	// try to unmarshal data into TransactionalFailure3Response
	err = newStrictDecoder(data).Decode(&dst.TransactionalFailure3Response)
	if err == nil {
		jsonTransactionalFailure3Response, _ := json.Marshal(dst.TransactionalFailure3Response)
		if string(jsonTransactionalFailure3Response) == "{}" { // empty struct
			dst.TransactionalFailure3Response = nil
		} else {
			match++
		}
	} else {
		dst.TransactionalFailure3Response = nil
	}

	// try to unmarshal data into TransactionalFailureResponse
	err = newStrictDecoder(data).Decode(&dst.TransactionalFailureResponse)
	if err == nil {
		jsonTransactionalFailureResponse, _ := json.Marshal(dst.TransactionalFailureResponse)
		if string(jsonTransactionalFailureResponse) == "{}" { // empty struct
			dst.TransactionalFailureResponse = nil
		} else {
			match++
		}
	} else {
		dst.TransactionalFailureResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TransactionalFailure2Response = nil
		dst.TransactionalFailure3Response = nil
		dst.TransactionalFailureResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TransactionalPost400Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TransactionalPost400Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransactionalPost400Response) MarshalJSON() ([]byte, error) {
	if src.TransactionalFailure2Response != nil {
		return json.Marshal(&src.TransactionalFailure2Response)
	}

	if src.TransactionalFailure3Response != nil {
		return json.Marshal(&src.TransactionalFailure3Response)
	}

	if src.TransactionalFailureResponse != nil {
		return json.Marshal(&src.TransactionalFailureResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransactionalPost400Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.TransactionalFailure2Response != nil {
		return obj.TransactionalFailure2Response
	}

	if obj.TransactionalFailure3Response != nil {
		return obj.TransactionalFailure3Response
	}

	if obj.TransactionalFailureResponse != nil {
		return obj.TransactionalFailureResponse
	}

	// all schemas are nil
	return nil
}

type NullableTransactionalPost400Response struct {
	value *TransactionalPost400Response
	isSet bool
}

func (v NullableTransactionalPost400Response) Get() *TransactionalPost400Response {
	return v.value
}

func (v *NullableTransactionalPost400Response) Set(val *TransactionalPost400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionalPost400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionalPost400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionalPost400Response(val *TransactionalPost400Response) *NullableTransactionalPost400Response {
	return &NullableTransactionalPost400Response{value: val, isSet: true}
}

func (v NullableTransactionalPost400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionalPost400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

