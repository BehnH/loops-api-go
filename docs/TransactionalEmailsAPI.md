# \TransactionalEmailsAPI

All URIs are relative to *https://app.loops.so/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionalGet**](TransactionalEmailsAPI.md#TransactionalGet) | **Get** /transactional | List transactional emails
[**TransactionalPost**](TransactionalEmailsAPI.md#TransactionalPost) | **Post** /transactional | Send a transactional email



## TransactionalGet

> []TransactionalEmail TransactionalGet(ctx).PerPage(perPage).Cursor(cursor).Execute()

List transactional emails



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
	perPage := "perPage_example" // string | How many results to return in each request. Must be between 10 and 50. Default is 20. (optional)
	cursor := "cursor_example" // string | A cursor, to return a specific page of results. Cursors can be found from the `pagination.nextCursor` value in each response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionalEmailsAPI.TransactionalGet(context.Background()).PerPage(perPage).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionalEmailsAPI.TransactionalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionalGet`: []TransactionalEmail
	fmt.Fprintf(os.Stdout, "Response from `TransactionalEmailsAPI.TransactionalGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **string** | How many results to return in each request. Must be between 10 and 50. Default is 20. | 
 **cursor** | **string** | A cursor, to return a specific page of results. Cursors can be found from the &#x60;pagination.nextCursor&#x60; value in each response. | 

### Return type

[**[]TransactionalEmail**](TransactionalEmail.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionalPost

> TransactionalSuccessResponse TransactionalPost(ctx).TransactionalRequest(transactionalRequest).IdempotencyKey(idempotencyKey).Execute()

Send a transactional email



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
	transactionalRequest := *openapiclient.NewTransactionalRequest("Email_example", "TransactionalId_example") // TransactionalRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Include a unique ID for this request (maximum 100 characters) to avoid duplicate emails. [More info](https://loops.so/docs/api-reference/send-transactional-email#param-idempotency-key) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionalEmailsAPI.TransactionalPost(context.Background()).TransactionalRequest(transactionalRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionalEmailsAPI.TransactionalPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionalPost`: TransactionalSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionalEmailsAPI.TransactionalPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionalPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionalRequest** | [**TransactionalRequest**](TransactionalRequest.md) |  | 
 **idempotencyKey** | **string** | Include a unique ID for this request (maximum 100 characters) to avoid duplicate emails. [More info](https://loops.so/docs/api-reference/send-transactional-email#param-idempotency-key) | 

### Return type

[**TransactionalSuccessResponse**](TransactionalSuccessResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

