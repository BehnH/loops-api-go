# TransactionalRequestAttachmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** | The name of the file, shown in email clients. | 
**ContentType** | **string** | The MIME type of the file. | 
**Data** | **string** | The base64-encoded content of the file. | 

## Methods

### NewTransactionalRequestAttachmentsInner

`func NewTransactionalRequestAttachmentsInner(filename string, contentType string, data string, ) *TransactionalRequestAttachmentsInner`

NewTransactionalRequestAttachmentsInner instantiates a new TransactionalRequestAttachmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionalRequestAttachmentsInnerWithDefaults

`func NewTransactionalRequestAttachmentsInnerWithDefaults() *TransactionalRequestAttachmentsInner`

NewTransactionalRequestAttachmentsInnerWithDefaults instantiates a new TransactionalRequestAttachmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *TransactionalRequestAttachmentsInner) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *TransactionalRequestAttachmentsInner) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *TransactionalRequestAttachmentsInner) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetContentType

`func (o *TransactionalRequestAttachmentsInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *TransactionalRequestAttachmentsInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *TransactionalRequestAttachmentsInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetData

`func (o *TransactionalRequestAttachmentsInner) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionalRequestAttachmentsInner) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionalRequestAttachmentsInner) SetData(v string)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


