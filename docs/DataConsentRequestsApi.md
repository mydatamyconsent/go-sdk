# MyDataMyConsent\DataConsentRequestsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelConsentRequest**](DataConsentRequestsApi.md#CancelConsentRequest) | **Delete** /v1/consent-requests/{requestId}/cancel | Cancel a Consent Request by ID.
[**CreateRequest**](DataConsentRequestsApi.md#CreateRequest) | **Post** /v1/consent-requests | Create a consent request.
[**GetAllConsentRequests**](DataConsentRequestsApi.md#GetAllConsentRequests) | **Get** /v1/consent-requests | Get all Consent Requests.
[**GetConsentRequestById**](DataConsentRequestsApi.md#GetConsentRequestById) | **Get** /v1/consent-requests/{requestId} | Get a Consent Request by ID.



## CancelConsentRequest

> CancelConsentRequest(ctx, requestId).Execute()

Cancel a Consent Request by ID.



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
    requestId := TODO // string | consent request id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataConsentRequestsApi.CancelConsentRequest(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.CancelConsentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | [**string**](.md) | consent request id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelConsentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRequest

> DataConsent CreateRequest(ctx).DataConsentRequestModel(dataConsentRequestModel).Execute()

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
    dataConsentRequestModel := *openapiclient.NewDataConsentRequestModel() // DataConsentRequestModel | MyDataMyConsent.Models.Consents.DataConsentRequestModel. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataConsentRequestsApi.CreateRequest(context.Background()).DataConsentRequestModel(dataConsentRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.CreateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRequest`: DataConsent
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

[**DataConsent**](DataConsent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllConsentRequests

> map[string]interface{} GetAllConsentRequests(ctx).Status(status).Execute()

Get all Consent Requests.

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
    status := openapiclient.DataConsentStatus("Pending") // DataConsentStatus |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataConsentRequestsApi.GetAllConsentRequests(context.Background()).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.GetAllConsentRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllConsentRequests`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.GetAllConsentRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConsentRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsentRequestById

> DataConsentDetailsDto GetConsentRequestById(ctx, requestId).Execute()

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
    requestId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataConsentRequestsApi.GetConsentRequestById(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentRequestsApi.GetConsentRequestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsentRequestById`: DataConsentDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentRequestsApi.GetConsentRequestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentRequestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataConsentDetailsDto**](DataConsentDetailsDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

