# ApiKeyGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**TeamName** | **string** | The name of the team the API key belongs to. | 

## Methods

### NewApiKeyGet200Response

`func NewApiKeyGet200Response(success bool, teamName string, ) *ApiKeyGet200Response`

NewApiKeyGet200Response instantiates a new ApiKeyGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyGet200ResponseWithDefaults

`func NewApiKeyGet200ResponseWithDefaults() *ApiKeyGet200Response`

NewApiKeyGet200ResponseWithDefaults instantiates a new ApiKeyGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ApiKeyGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ApiKeyGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ApiKeyGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetTeamName

`func (o *ApiKeyGet200Response) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *ApiKeyGet200Response) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *ApiKeyGet200Response) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


