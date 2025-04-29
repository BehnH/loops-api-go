# IdempotencyKeyFailureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Message** | **string** |  | 

## Methods

### NewIdempotencyKeyFailureResponse

`func NewIdempotencyKeyFailureResponse(success bool, message string, ) *IdempotencyKeyFailureResponse`

NewIdempotencyKeyFailureResponse instantiates a new IdempotencyKeyFailureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdempotencyKeyFailureResponseWithDefaults

`func NewIdempotencyKeyFailureResponseWithDefaults() *IdempotencyKeyFailureResponse`

NewIdempotencyKeyFailureResponseWithDefaults instantiates a new IdempotencyKeyFailureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *IdempotencyKeyFailureResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *IdempotencyKeyFailureResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *IdempotencyKeyFailureResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *IdempotencyKeyFailureResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IdempotencyKeyFailureResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IdempotencyKeyFailureResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


