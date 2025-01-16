# Go API client for loops

This is the OpenAPI Spec for the [Loops API](https://loops.so/docs/api).

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.3.4
- Package version: 1.0.2
- Generator version: 7.7.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import loops "github.com/behnh/loops-api-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `loops.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), loops.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `loops.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), loops.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `loops.ContextOperationServerIndices` and `loops.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), loops.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), loops.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://app.loops.so/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*APIKeyAPI* | [**ApiKeyGet**](docs/APIKeyAPI.md#apikeyget) | **Get** /api-key | Test your API key
*ContactsAPI* | [**ContactsCreatePost**](docs/ContactsAPI.md#contactscreatepost) | **Post** /contacts/create | Create a contact
*ContactsAPI* | [**ContactsDeletePost**](docs/ContactsAPI.md#contactsdeletepost) | **Post** /contacts/delete | Delete a contact
*ContactsAPI* | [**ContactsFindGet**](docs/ContactsAPI.md#contactsfindget) | **Get** /contacts/find | Find a contact
*ContactsAPI* | [**ContactsUpdatePut**](docs/ContactsAPI.md#contactsupdateput) | **Put** /contacts/update | Update a contact
*CustomFieldsAPI* | [**ContactsCustomFieldsGet**](docs/CustomFieldsAPI.md#contactscustomfieldsget) | **Get** /contacts/customFields | Get a list of custom contact properties
*EventsAPI* | [**EventsSendPost**](docs/EventsAPI.md#eventssendpost) | **Post** /events/send | Send an event
*MailingListsAPI* | [**ListsGet**](docs/MailingListsAPI.md#listsget) | **Get** /lists | Get a list of mailing lists
*TransactionalEmailsAPI* | [**TransactionalPost**](docs/TransactionalEmailsAPI.md#transactionalpost) | **Post** /transactional | Send a transactional email


## Documentation For Models

 - [ApiKeyGet200Response](docs/ApiKeyGet200Response.md)
 - [ApiKeyGet401Response](docs/ApiKeyGet401Response.md)
 - [Contact](docs/Contact.md)
 - [ContactDeleteRequest](docs/ContactDeleteRequest.md)
 - [ContactDeleteResponse](docs/ContactDeleteResponse.md)
 - [ContactFailureResponse](docs/ContactFailureResponse.md)
 - [ContactRequest](docs/ContactRequest.md)
 - [ContactSuccessResponse](docs/ContactSuccessResponse.md)
 - [CustomField](docs/CustomField.md)
 - [EventFailureResponse](docs/EventFailureResponse.md)
 - [EventRequest](docs/EventRequest.md)
 - [EventSuccessResponse](docs/EventSuccessResponse.md)
 - [MailingList](docs/MailingList.md)
 - [TransactionalFailure2Response](docs/TransactionalFailure2Response.md)
 - [TransactionalFailure2ResponseError](docs/TransactionalFailure2ResponseError.md)
 - [TransactionalFailure3Response](docs/TransactionalFailure3Response.md)
 - [TransactionalFailureResponse](docs/TransactionalFailureResponse.md)
 - [TransactionalPost400Response](docs/TransactionalPost400Response.md)
 - [TransactionalRequest](docs/TransactionalRequest.md)
 - [TransactionalRequestAttachmentsInner](docs/TransactionalRequestAttachmentsInner.md)
 - [TransactionalSuccessResponse](docs/TransactionalSuccessResponse.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### apiKey

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), loops.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



