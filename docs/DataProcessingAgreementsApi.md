# \DataProcessingAgreementsApi

All URIs are relative to *https://api.mydatamyconsent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DataAgreementsGet**](DataProcessingAgreementsApi.md#V1DataAgreementsGet) | **Get** /v1/data-agreements | Get paginated data processing agreements.
[**V1DataAgreementsIdDelete**](DataProcessingAgreementsApi.md#V1DataAgreementsIdDelete) | **Delete** /v1/data-agreements/{id} | Delete a data processing agreement this will not delete a published or a agreement in use with consents.
[**V1DataAgreementsIdGet**](DataProcessingAgreementsApi.md#V1DataAgreementsIdGet) | **Get** /v1/data-agreements/{id} | Get data processing agreement by id.
[**V1DataAgreementsIdTerminatePut**](DataProcessingAgreementsApi.md#V1DataAgreementsIdTerminatePut) | **Put** /v1/data-agreements/{id}/terminate | Terminate a data processing agreement by id.



## V1DataAgreementsGet

> PaginatedListOfDataProcessingAgreements V1DataAgreementsGet(ctx).PageNo(pageNo).PageSize(pageSize).Execute()

Get paginated data processing agreements.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.V1DataAgreementsGet(context.Background()).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.V1DataAgreementsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DataAgreementsGet`: PaginatedListOfDataProcessingAgreements
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingAgreementsApi.V1DataAgreementsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DataAgreementsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNo** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**PaginatedListOfDataProcessingAgreements**](PaginatedListOfDataProcessingAgreements.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataAgreementsIdDelete

> bool V1DataAgreementsIdDelete(ctx, id).Execute()

Delete a data processing agreement this will not delete a published or a agreement in use with consents.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.V1DataAgreementsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.V1DataAgreementsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DataAgreementsIdDelete`: bool
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingAgreementsApi.V1DataAgreementsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DataAgreementsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataAgreementsIdGet

> DataProcessingAgreement V1DataAgreementsIdGet(ctx, id).Execute()

Get data processing agreement by id.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.V1DataAgreementsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.V1DataAgreementsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DataAgreementsIdGet`: DataProcessingAgreement
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingAgreementsApi.V1DataAgreementsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DataAgreementsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataProcessingAgreement**](DataProcessingAgreement.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataAgreementsIdTerminatePut

> bool V1DataAgreementsIdTerminatePut(ctx, id).Execute()

Terminate a data processing agreement by id.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.V1DataAgreementsIdTerminatePut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.V1DataAgreementsIdTerminatePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DataAgreementsIdTerminatePut`: bool
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingAgreementsApi.V1DataAgreementsIdTerminatePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DataAgreementsIdTerminatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

