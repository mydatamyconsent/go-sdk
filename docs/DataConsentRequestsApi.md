# \DataConsentRequestsApi

All URIs are relative to *https://api.mydatamyconsent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelIndividualDataConsentRequest**](DataConsentRequestsApi.md#CancelIndividualDataConsentRequest) | **Put** /v1/consent-requests/individual/{requestId}/cancel | Cancel the individual data consent request by Id.
[**CancelOrganizationDataConsentRequest**](DataConsentRequestsApi.md#CancelOrganizationDataConsentRequest) | **Put** /v1/consent-requests/organization/{requestId}/cancel | Cancel the organization data consent request by Id.
[**CreateIndividualDataConsentRequest**](DataConsentRequestsApi.md#CreateIndividualDataConsentRequest) | **Post** /v1/consent-requests/individual | Create data consent request for an individual.
[**CreateOrganizationDataConsentRequest**](DataConsentRequestsApi.md#CreateOrganizationDataConsentRequest) | **Post** /v1/consent-requests/organization | Create data consent request for an organization.
[**GetAllConsentRequestsToIndividuals**](DataConsentRequestsApi.md#GetAllConsentRequestsToIndividuals) | **Get** /v1/consent-requests/individuals | Get all Consent Requests sent to individuals.
[**GetAllConsentRequestsToOrganizations**](DataConsentRequestsApi.md#GetAllConsentRequestsToOrganizations) | **Get** /v1/consent-requests/organizations | Get all Consent Requests sent to organizations.
[**GetIndividualConsentRequestById**](DataConsentRequestsApi.md#GetIndividualConsentRequestById) | **Get** /v1/consent-requests/individuals/{requestId} | Get individual data consent request by id.
[**GetOrganizationConsentRequestById**](DataConsentRequestsApi.md#GetOrganizationConsentRequestById) | **Get** /v1/consent-requests/organizations/{requestId} | Get a OrganizationConsent Request by Id.



## CancelIndividualDataConsentRequest

> CancelIndividualDataConsentRequest(ctx, requestId).Execute()

Cancel the individual data consent request by Id.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    requestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Individual consent request id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.CancelIndividualDataConsentRequest(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.CancelIndividualDataConsentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Individual consent request id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelIndividualDataConsentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelOrganizationDataConsentRequest

> CancelOrganizationDataConsentRequest(ctx, requestId).Execute()

Cancel the organization data consent request by Id.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    requestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization consent request id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.CancelOrganizationDataConsentRequest(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.CancelOrganizationDataConsentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Organization consent request id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrganizationDataConsentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIndividualDataConsentRequest

> IndividualDataConsentRequestDetails CreateIndividualDataConsentRequest(ctx).CreateDataConsentRequest(createDataConsentRequest).Execute()

Create data consent request for an individual.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createDataConsentRequest := *openapiclient.NewCreateDataConsentRequest("ConsentTemplateId_example", *openapiclient.NewConsentRequestReceiver("CountryIso2Code_example", []openapiclient.StringStringKeyValuePair{*openapiclient.NewStringStringKeyValuePair()}, openapiclient.IdentificationStrategy("MatchAtLeastOneIdentifier"))) // CreateDataConsentRequest | The Individual data consent request payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.CreateIndividualDataConsentRequest(context.Background()).CreateDataConsentRequest(createDataConsentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.CreateIndividualDataConsentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndividualDataConsentRequest`: IndividualDataConsentRequestDetails
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.CreateIndividualDataConsentRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualDataConsentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDataConsentRequest** | [**CreateDataConsentRequest**](CreateDataConsentRequest.md) | The Individual data consent request payload | 

### Return type

[**IndividualDataConsentRequestDetails**](IndividualDataConsentRequestDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationDataConsentRequest

> OrganizationDataConsentRequestDetails CreateOrganizationDataConsentRequest(ctx).CreateDataConsentRequest(createDataConsentRequest).Execute()

Create data consent request for an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createDataConsentRequest := *openapiclient.NewCreateDataConsentRequest("ConsentTemplateId_example", *openapiclient.NewConsentRequestReceiver("CountryIso2Code_example", []openapiclient.StringStringKeyValuePair{*openapiclient.NewStringStringKeyValuePair()}, openapiclient.IdentificationStrategy("MatchAtLeastOneIdentifier"))) // CreateDataConsentRequest | The Organization data consent request payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.CreateOrganizationDataConsentRequest(context.Background()).CreateDataConsentRequest(createDataConsentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.CreateOrganizationDataConsentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationDataConsentRequest`: OrganizationDataConsentRequestDetails
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.CreateOrganizationDataConsentRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationDataConsentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDataConsentRequest** | [**CreateDataConsentRequest**](CreateDataConsentRequest.md) | The Organization data consent request payload | 

### Return type

[**OrganizationDataConsentRequestDetails**](OrganizationDataConsentRequestDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllConsentRequestsToIndividuals

> IndividualDataConsentRequestDetailsPaginatedList GetAllConsentRequestsToIndividuals(ctx).Status(status).StartDateTime(startDateTime).EndDateTime(endDateTime).PageNo(pageNo).PageSize(pageSize).Execute()

Get all Consent Requests sent to individuals.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    status := openapiclient.DataConsentStatus("Pending") // DataConsentStatus | Data consent status. (optional)
    startDateTime := time.Now() // time.Time | Start datetime in UTC timezone. (optional)
    endDateTime := time.Now() // time.Time | End datetime in UTC timezone. (optional)
    pageNo := int32(56) // int32 | Page number. (optional) (default to 1)
    pageSize := int32(56) // int32 | Number of items to return. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.GetAllConsentRequestsToIndividuals(context.Background()).Status(status).StartDateTime(startDateTime).EndDateTime(endDateTime).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.GetAllConsentRequestsToIndividuals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllConsentRequestsToIndividuals`: IndividualDataConsentRequestDetailsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.GetAllConsentRequestsToIndividuals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConsentRequestsToIndividualsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**DataConsentStatus**](DataConsentStatus.md) | Data consent status. | 
 **startDateTime** | **time.Time** | Start datetime in UTC timezone. | 
 **endDateTime** | **time.Time** | End datetime in UTC timezone. | 
 **pageNo** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Number of items to return. | [default to 25]

### Return type

[**IndividualDataConsentRequestDetailsPaginatedList**](IndividualDataConsentRequestDetailsPaginatedList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllConsentRequestsToOrganizations

> OrganizationDataConsentRequestDetailsPaginatedList GetAllConsentRequestsToOrganizations(ctx).Status(status).StartDateTime(startDateTime).EndDateTime(endDateTime).PageNo(pageNo).PageSize(pageSize).Execute()

Get all Consent Requests sent to organizations.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    status := openapiclient.DataConsentStatus("Pending") // DataConsentStatus | Data consent status. (optional)
    startDateTime := time.Now() // time.Time | Start datetime in UTC timezone. (optional)
    endDateTime := time.Now() // time.Time | End datetime in UTC timezone. (optional)
    pageNo := int32(56) // int32 | Page number. (optional) (default to 1)
    pageSize := int32(56) // int32 | Number of items to return. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.GetAllConsentRequestsToOrganizations(context.Background()).Status(status).StartDateTime(startDateTime).EndDateTime(endDateTime).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.GetAllConsentRequestsToOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllConsentRequestsToOrganizations`: OrganizationDataConsentRequestDetailsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.GetAllConsentRequestsToOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConsentRequestsToOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**DataConsentStatus**](DataConsentStatus.md) | Data consent status. | 
 **startDateTime** | **time.Time** | Start datetime in UTC timezone. | 
 **endDateTime** | **time.Time** | End datetime in UTC timezone. | 
 **pageNo** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Number of items to return. | [default to 25]

### Return type

[**OrganizationDataConsentRequestDetailsPaginatedList**](OrganizationDataConsentRequestDetailsPaginatedList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndividualConsentRequestById

> DataConsentRequest GetIndividualConsentRequestById(ctx, requestId).Execute()

Get individual data consent request by id.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    requestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Individual data consent request id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.GetIndividualConsentRequestById(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.GetIndividualConsentRequestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndividualConsentRequestById`: DataConsentRequest
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.GetIndividualConsentRequestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Individual data consent request id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndividualConsentRequestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataConsentRequest**](DataConsentRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConsentRequestById

> DataConsentRequest GetOrganizationConsentRequestById(ctx, requestId).Execute()

Get a OrganizationConsent Request by Id.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    requestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization consent request id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.GetOrganizationConsentRequestById(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.GetOrganizationConsentRequestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConsentRequestById`: DataConsentRequest
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.GetOrganizationConsentRequestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Organization consent request id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConsentRequestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataConsentRequest**](DataConsentRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

