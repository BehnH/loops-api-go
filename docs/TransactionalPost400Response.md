# TransactionalPost400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Path** | **string** |  | 
**Message** | **string** |  | 
**Error** | [**TransactionalFailure2ResponseError**](TransactionalFailure2ResponseError.md) |  | 

## Methods

### NewTransactionalPost400Response

`func NewTransactionalPost400Response(success bool, path string, message string, error_ TransactionalFailure2ResponseError, ) *TransactionalPost400Response`

NewTransactionalPost400Response instantiates a new TransactionalPost400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionalPost400ResponseWithDefaults

`func NewTransactionalPost400ResponseWithDefaults() *TransactionalPost400Response`

NewTransactionalPost400ResponseWithDefaults instantiates a new TransactionalPost400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *TransactionalPost400Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TransactionalPost400Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TransactionalPost400Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetPath

`func (o *TransactionalPost400Response) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TransactionalPost400Response) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TransactionalPost400Response) SetPath(v string)`

SetPath sets Path field to given value.


### GetMessage

`func (o *TransactionalPost400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransactionalPost400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransactionalPost400Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *TransactionalPost400Response) GetError() TransactionalFailure2ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TransactionalPost400Response) GetErrorOk() (*TransactionalFailure2ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TransactionalPost400Response) SetError(v TransactionalFailure2ResponseError)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


