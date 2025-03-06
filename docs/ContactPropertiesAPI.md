# \ContactPropertiesAPI

All URIs are relative to *https://app.loops.so/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactsPropertiesGet**](ContactPropertiesAPI.md#ContactsPropertiesGet) | **Get** /contacts/properties | Get a list of contact properties
[**ContactsPropertiesPost**](ContactPropertiesAPI.md#ContactsPropertiesPost) | **Post** /contacts/properties | Create a contact property



## ContactsPropertiesGet

> []ContactProperty ContactsPropertiesGet(ctx).List(list).Execute()

Get a list of contact properties



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/behnh/loops-api-go"
)

func main() {
	list := "list_example" // string | \\\"all\\\" (default) or \\\"custom\\\" (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactPropertiesAPI.ContactsPropertiesGet(context.Background()).List(list).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactPropertiesAPI.ContactsPropertiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContactsPropertiesGet`: []ContactProperty
	fmt.Fprintf(os.Stdout, "Response from `ContactPropertiesAPI.ContactsPropertiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsPropertiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | \\\&quot;all\\\&quot; (default) or \\\&quot;custom\\\&quot; | 

### Return type

[**[]ContactProperty**](ContactProperty.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsPropertiesPost

> ContactPropertySuccessResponse ContactsPropertiesPost(ctx).ContactPropertyCreateRequest(contactPropertyCreateRequest).Execute()

Create a contact property



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/behnh/loops-api-go"
)

func main() {
	contactPropertyCreateRequest := *openapiclient.NewContactPropertyCreateRequest("Name_example", "Type_example") // ContactPropertyCreateRequest | The name value must be in camelCase, like `planName`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactPropertiesAPI.ContactsPropertiesPost(context.Background()).ContactPropertyCreateRequest(contactPropertyCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactPropertiesAPI.ContactsPropertiesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContactsPropertiesPost`: ContactPropertySuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `ContactPropertiesAPI.ContactsPropertiesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsPropertiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactPropertyCreateRequest** | [**ContactPropertyCreateRequest**](ContactPropertyCreateRequest.md) | The name value must be in camelCase, like &#x60;planName&#x60;. | 

### Return type

[**ContactPropertySuccessResponse**](ContactPropertySuccessResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

