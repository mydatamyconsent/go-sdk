# \DataConsentRequestsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelConsentRequest**](DataConsentRequestsApi.md#CancelConsentRequest) | **Delete** /v1/consent-requests/{requestId}/cancel | Revoke / Cancel the ConsentRequest based on Id.
[**CreateRequest**](DataConsentRequestsApi.md#CreateRequest) | **Post** /v1/consent-requests | Create a consent request.
[**GetAllConsentRequestsToIndividuals**](DataConsentRequestsApi.md#GetAllConsentRequestsToIndividuals) | **Get** /v1/consent-requests/individuals | Get all Consent Requests sent to Individuals.
[**GetAllConsentRequestsToOrganizations**](DataConsentRequestsApi.md#GetAllConsentRequestsToOrganizations) | **Get** /v1/consent-requests/organizations | Get All Consent Requests sent to Organizations
[**GetIndividualConsentRequestById**](DataConsentRequestsApi.md#GetIndividualConsentRequestById) | **Get** /v1/consent-requests/individuals/{requestId} | Get a Consent Request by ID.
[**GetOrganizationConsentRequestById**](DataConsentRequestsApi.md#GetOrganizationConsentRequestById) | **Get** /v1/consent-requests/organizations/{requestId} | Get a OrganizationConsent Request by Id



## CancelConsentRequest

> bool CancelConsentRequest(ctx, requestId).Execute()

Revoke / Cancel the ConsentRequest based on Id.

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
    requestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.CancelConsentRequest(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.CancelConsentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelConsentRequest`: bool
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.CancelConsentRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelConsentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRequest

> bool CreateRequest(ctx).DataConsentRequestModel(dataConsentRequestModel).Execute()

Create a consent request.

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
    dataConsentRequestModel := *openapiclient.NewDataConsentRequestModel(*openapiclient.NewReceiver()) // DataConsentRequestModel | MyDataMyConsent.Models.Consents.DataConsentRequestModel. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.CreateRequest(context.Background()).DataConsentRequestModel(dataConsentRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.CreateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRequest`: bool
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.CreateRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataConsentRequestModel** | [**DataConsentRequestModel**](DataConsentRequestModel.md) | MyDataMyConsent.Models.Consents.DataConsentRequestModel. | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllConsentRequestsToIndividuals

> map[string]interface{} GetAllConsentRequestsToIndividuals(ctx).PageNo(pageNo).PageSize(pageSize).Status(status).Execute()

Get all Consent Requests sent to Individuals.

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
    pageNo := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    status := openapiclient.DataConsentStatus("Pending") // DataConsentStatus |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.GetAllConsentRequestsToIndividuals(context.Background()).PageNo(pageNo).PageSize(pageSize).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.GetAllConsentRequestsToIndividuals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllConsentRequestsToIndividuals`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.GetAllConsentRequestsToIndividuals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConsentRequestsToIndividualsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNo** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllConsentRequestsToOrganizations

> map[string]interface{} GetAllConsentRequestsToOrganizations(ctx).PageNo(pageNo).PageSize(pageSize).Status(status).Execute()

Get All Consent Requests sent to Organizations

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
    pageNo := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    status := openapiclient.DataConsentStatus("Pending") // DataConsentStatus |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentRequestsApi.GetAllConsentRequestsToOrganizations(context.Background()).PageNo(pageNo).PageSize(pageSize).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.GetAllConsentRequestsToOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllConsentRequestsToOrganizations`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.GetAllConsentRequestsToOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConsentRequestsToOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNo** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 

### Return type

**map[string]interface{}**

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
    requestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

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
**requestId** | **string** |  | 

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

Get a OrganizationConsent Request by Id

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
    requestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

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
**requestId** | **string** |  | 

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

