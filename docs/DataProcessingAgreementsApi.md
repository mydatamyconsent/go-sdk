# MyDataMyConsent\DataProcessingAgreementsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DataAgreementsGet**](DataProcessingAgreementsApi.md#V1DataAgreementsGet) | **Get** /v1/data-agreements | Get all data processing agreements.
[**V1DataAgreementsIdDelete**](DataProcessingAgreementsApi.md#V1DataAgreementsIdDelete) | **Delete** /v1/data-agreements/{id} | Delete a data processing agreement. This will not delete a published or a agreement in use with consents.
[**V1DataAgreementsIdGet**](DataProcessingAgreementsApi.md#V1DataAgreementsIdGet) | **Get** /v1/data-agreements/{id} | Get data processing agreement by Id.
[**V1DataAgreementsIdPut**](DataProcessingAgreementsApi.md#V1DataAgreementsIdPut) | **Put** /v1/data-agreements/{id} | Update a data processing agreement.
[**V1DataAgreementsIdTerminatePut**](DataProcessingAgreementsApi.md#V1DataAgreementsIdTerminatePut) | **Put** /v1/data-agreements/{id}/terminate | Terminate a data processing agreement.
[**V1DataAgreementsPost**](DataProcessingAgreementsApi.md#V1DataAgreementsPost) | **Post** /v1/data-agreements | Create a data processing agreement.



## V1DataAgreementsGet

> DataProcessingAgreementPaginatedList V1DataAgreementsGet(ctx).PageNo(pageNo).PageSize(pageSize).Execute()

Get all data processing agreements.

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
    pageNo := int32(56) // int32 | Page number. (optional) (default to 1)
    pageSize := int32(56) // int32 | Number of items to return. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataProcessingAgreementsApi.V1DataAgreementsGet(context.Background()).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.V1DataAgreementsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DataAgreementsGet`: DataProcessingAgreementPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingAgreementsApi.V1DataAgreementsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DataAgreementsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNo** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Number of items to return. | [default to 25]

### Return type

[**DataProcessingAgreementPaginatedList**](DataProcessingAgreementPaginatedList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataAgreementsIdDelete

> V1DataAgreementsIdDelete(ctx, id).Execute()

Delete a data processing agreement. This will not delete a published or a agreement in use with consents.

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
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataProcessingAgreementsApi.V1DataAgreementsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.V1DataAgreementsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DataAgreementsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataAgreementsIdGet

> DataProcessingAgreement V1DataAgreementsIdGet(ctx, id).Execute()

Get data processing agreement by Id.

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
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataProcessingAgreementsApi.V1DataAgreementsIdGet(context.Background(), id).Execute()
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
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DataAgreementsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataProcessingAgreement**](DataProcessingAgreement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataAgreementsIdPut

> DataProcessingAgreement V1DataAgreementsIdPut(ctx, id).DataProcessingAgreement(dataProcessingAgreement).Execute()

Update a data processing agreement.

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
    id := TODO // string | 
    dataProcessingAgreement := *openapiclient.NewDataProcessingAgreement() // DataProcessingAgreement |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataProcessingAgreementsApi.V1DataAgreementsIdPut(context.Background(), id).DataProcessingAgreement(dataProcessingAgreement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.V1DataAgreementsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DataAgreementsIdPut`: DataProcessingAgreement
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingAgreementsApi.V1DataAgreementsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DataAgreementsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataProcessingAgreement** | [**DataProcessingAgreement**](DataProcessingAgreement.md) |  | 

### Return type

[**DataProcessingAgreement**](DataProcessingAgreement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataAgreementsIdTerminatePut

> V1DataAgreementsIdTerminatePut(ctx, id).Execute()

Terminate a data processing agreement.

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
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataProcessingAgreementsApi.V1DataAgreementsIdTerminatePut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.V1DataAgreementsIdTerminatePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DataAgreementsIdTerminatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataAgreementsPost

> DataProcessingAgreement V1DataAgreementsPost(ctx).DataProcessingAgreement(dataProcessingAgreement).Execute()

Create a data processing agreement.

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
    dataProcessingAgreement := *openapiclient.NewDataProcessingAgreement() // DataProcessingAgreement |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataProcessingAgreementsApi.V1DataAgreementsPost(context.Background()).DataProcessingAgreement(dataProcessingAgreement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.V1DataAgreementsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DataAgreementsPost`: DataProcessingAgreement
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingAgreementsApi.V1DataAgreementsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DataAgreementsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataProcessingAgreement** | [**DataProcessingAgreement**](DataProcessingAgreement.md) |  | 

### Return type

[**DataProcessingAgreement**](DataProcessingAgreement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

