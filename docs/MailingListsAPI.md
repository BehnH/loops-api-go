# \MailingListsAPI

All URIs are relative to *https://app.loops.so/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListsGet**](MailingListsAPI.md#ListsGet) | **Get** /lists | Get a list of mailing lists



## ListsGet

> []MailingList ListsGet(ctx).Execute()

Get a list of mailing lists



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MailingListsAPI.ListsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailingListsAPI.ListsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsGet`: []MailingList
	fmt.Fprintf(os.Stdout, "Response from `MailingListsAPI.ListsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListsGetRequest struct via the builder pattern


### Return type

[**[]MailingList**](MailingList.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

