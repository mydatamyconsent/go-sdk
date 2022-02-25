# \DataConsentRequestsApi

All URIs are relative to *https://api.mydatamyconsent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelIndividualDataConsentRequest**](DataConsentRequestsApi.md#CancelIndividualDataConsentRequest) | **Put** /v1/consent-requests/individual/{requestId}/cancel | Cancel the individual data consent request based on Id.
[**CancelOrganizationDataConsentRequest**](DataConsentRequestsApi.md#CancelOrganizationDataConsentRequest) | **Put** /v1/consent-requests/organization/{requestId}/cancel | Cancel the Organization data consent request based on Id.
[**CreateIndividualDataConsentRequest**](DataConsentRequestsApi.md#CreateIndividualDataConsentRequest) | **Post** /v1/consent-requests/individual | Create a individual data consent request.
[**CreateOrganizationDataConsentRequest**](DataConsentRequestsApi.md#CreateOrganizationDataConsentRequest) | **Post** /v1/consent-requests/organization | Create a organization data consent request.
[**GetAllConsentRequestsToIndividuals**](DataConsentRequestsApi.md#GetAllConsentRequestsToIndividuals) | **Get** /v1/consent-requests/individuals | Get all Consent Requests sent to Individuals.
[**GetAllConsentRequestsToOrganizations**](DataConsentRequestsApi.md#GetAllConsentRequestsToOrganizations) | **Get** /v1/consent-requests/organizations | Get All Consent Requests sent to Organizations.
[**GetIndividualConsentRequestById**](DataConsentRequestsApi.md#GetIndividualConsentRequestById) | **Get** /v1/consent-requests/individuals/{requestId} | Get a Consent Request by ID.
[**GetOrganizationConsentRequestById**](DataConsentRequestsApi.md#GetOrganizationConsentRequestById) | **Get** /v1/consent-requests/organizations/{requestId} | Get a OrganizationConsent Request by Id.



## CancelIndividualDataConsentRequest

> IndividualDataConsentRequestResponse CancelIndividualDataConsentRequest(ctx, requestId).Execute()

Cancel the individual data consent request based on Id.

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
    // response from `CancelIndividualDataConsentRequest`: IndividualDataConsentRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.CancelIndividualDataConsentRequest`: %v\n", resp)
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

[**IndividualDataConsentRequestResponse**](IndividualDataConsentRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelOrganizationDataConsentRequest

> OrganizationDataConsentRequestResponse CancelOrganizationDataConsentRequest(ctx, requestId).Execute()

Cancel the Organization data consent request based on Id.

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
    // response from `CancelOrganizationDataConsentRequest`: OrganizationDataConsentRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.CancelOrganizationDataConsentRequest`: %v\n", resp)
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

[**OrganizationDataConsentRequestResponse**](OrganizationDataConsentRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIndividualDataConsentRequest

> IndividualDataConsentRequestResponse CreateIndividualDataConsentRequest(ctx).CreateIndividualDataConsentRequest(createIndividualDataConsentRequest).Execute()

Create a individual data consent request.



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
    createIndividualDataConsentRequest := *openapiclient.NewCreateIndividualDataConsentRequest(*openapiclient.NewConsentRequestReceiver()) // CreateIndividualDataConsentRequest | The Individual data consent request payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.CreateIndividualDataConsentRequest(context.Background()).CreateIndividualDataConsentRequest(createIndividualDataConsentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.CreateIndividualDataConsentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndividualDataConsentRequest`: IndividualDataConsentRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.CreateIndividualDataConsentRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualDataConsentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIndividualDataConsentRequest** | [**CreateIndividualDataConsentRequest**](CreateIndividualDataConsentRequest.md) | The Individual data consent request payload | 

### Return type

[**IndividualDataConsentRequestResponse**](IndividualDataConsentRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationDataConsentRequest

> OrganizationDataConsentRequestResponse CreateOrganizationDataConsentRequest(ctx).CreateOrganizationDataConsentRequest(createOrganizationDataConsentRequest).Execute()

Create a organization data consent request.



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
    createOrganizationDataConsentRequest := *openapiclient.NewCreateOrganizationDataConsentRequest(*openapiclient.NewConsentRequestReceiver()) // CreateOrganizationDataConsentRequest | M:MyDataMyConsent.DeveloperApi.Controllers.DataConsentRequestsController.CreateOrganizationDataConsentRequest(MyDataMyConsent.DeveloperApi.Models.CreateOrganizationDataConsentRequest).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.CreateOrganizationDataConsentRequest(context.Background()).CreateOrganizationDataConsentRequest(createOrganizationDataConsentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.CreateOrganizationDataConsentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationDataConsentRequest`: OrganizationDataConsentRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.CreateOrganizationDataConsentRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationDataConsentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrganizationDataConsentRequest** | [**CreateOrganizationDataConsentRequest**](CreateOrganizationDataConsentRequest.md) | M:MyDataMyConsent.DeveloperApi.Controllers.DataConsentRequestsController.CreateOrganizationDataConsentRequest(MyDataMyConsent.DeveloperApi.Models.CreateOrganizationDataConsentRequest). | 

### Return type

[**OrganizationDataConsentRequestResponse**](OrganizationDataConsentRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllConsentRequestsToIndividuals

> UserDataConsentInfoDtoPaginatedList GetAllConsentRequestsToIndividuals(ctx).Status(status).StartDateTime(startDateTime).EndDateTime(endDateTime).PageNo(pageNo).PageSize(pageSize).Execute()

Get all Consent Requests sent to Individuals.

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
    startDateTime := time.Now() // time.Time | Start date time. (optional)
    endDateTime := time.Now() // time.Time | End date time. (optional)
    pageNo := int32(56) // int32 | Page number. (optional) (default to 1)
    pageSize := int32(56) // int32 | Number of items to return. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.GetAllConsentRequestsToIndividuals(context.Background()).Status(status).StartDateTime(startDateTime).EndDateTime(endDateTime).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.GetAllConsentRequestsToIndividuals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllConsentRequestsToIndividuals`: UserDataConsentInfoDtoPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.GetAllConsentRequestsToIndividuals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConsentRequestsToIndividualsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**DataConsentStatus**](DataConsentStatus.md) | Data consent status. | 
 **startDateTime** | **time.Time** | Start date time. | 
 **endDateTime** | **time.Time** | End date time. | 
 **pageNo** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Number of items to return. | [default to 25]

### Return type

[**UserDataConsentInfoDtoPaginatedList**](UserDataConsentInfoDtoPaginatedList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllConsentRequestsToOrganizations

> OrganizationDataConsentInfoDtoPaginatedList GetAllConsentRequestsToOrganizations(ctx).Status(status).StartDateTime(startDateTime).EndDateTime(endDateTime).PageNo(pageNo).PageSize(pageSize).Execute()

Get All Consent Requests sent to Organizations.

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
    startDateTime := time.Now() // time.Time | Start date time. (optional)
    endDateTime := time.Now() // time.Time | End date time. (optional)
    pageNo := int32(56) // int32 | Page number. (optional) (default to 1)
    pageSize := int32(56) // int32 | Number of items to return. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.GetAllConsentRequestsToOrganizations(context.Background()).Status(status).StartDateTime(startDateTime).EndDateTime(endDateTime).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.GetAllConsentRequestsToOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllConsentRequestsToOrganizations`: OrganizationDataConsentInfoDtoPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.GetAllConsentRequestsToOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConsentRequestsToOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**DataConsentStatus**](DataConsentStatus.md) | Data consent status. | 
 **startDateTime** | **time.Time** | Start date time. | 
 **endDateTime** | **time.Time** | End date time. | 
 **pageNo** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Number of items to return. | [default to 25]

### Return type

[**OrganizationDataConsentInfoDtoPaginatedList**](OrganizationDataConsentInfoDtoPaginatedList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndividualConsentRequestById

> DataConsentDetailsDto GetIndividualConsentRequestById(ctx, requestId).Execute()

Get a Consent Request by ID.

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
    resp, r, err := apiClient.DataConsentRequestsApi.GetIndividualConsentRequestById(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.GetIndividualConsentRequestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndividualConsentRequestById`: DataConsentDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.GetIndividualConsentRequestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Individual consent request id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndividualConsentRequestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataConsentDetailsDto**](DataConsentDetailsDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConsentRequestById

> DataConsentDetailsDto GetOrganizationConsentRequestById(ctx, requestId).Execute()

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
    // response from `GetOrganizationConsentRequestById`: DataConsentDetailsDto
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

[**DataConsentDetailsDto**](DataConsentDetailsDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

