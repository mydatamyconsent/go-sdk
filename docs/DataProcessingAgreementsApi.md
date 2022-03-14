# \DataProcessingAgreementsApi

All URIs are relative to *https://api.mydatamyconsent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataProcessingAgreement**](DataProcessingAgreementsApi.md#CreateDataProcessingAgreement) | **Post** /v1/data-agreements | Create a data processing agreement.
[**DeleteDataProcessingAgreementById**](DataProcessingAgreementsApi.md#DeleteDataProcessingAgreementById) | **Delete** /v1/data-agreements/{id} | Delete a data processing agreement. This will not delete a published or a agreement in use with consents.
[**GetDataProcessingAgreementById**](DataProcessingAgreementsApi.md#GetDataProcessingAgreementById) | **Get** /v1/data-agreements/{id} | Get data processing agreement by id.
[**GetDataProcessingAgreements**](DataProcessingAgreementsApi.md#GetDataProcessingAgreements) | **Get** /v1/data-agreements | Get paginated data processing agreements.
[**TerminateDataProcessingAgreementById**](DataProcessingAgreementsApi.md#TerminateDataProcessingAgreementById) | **Put** /v1/data-agreements/{id}/terminate | Terminate a data processing agreement.
[**UpdateDataProcessingAgreement**](DataProcessingAgreementsApi.md#UpdateDataProcessingAgreement) | **Put** /v1/data-agreements/{id} | Update a data processing agreement.



## CreateDataProcessingAgreement

> DataProcessingAgreement CreateDataProcessingAgreement(ctx).CreateDataProcessingAgreement(createDataProcessingAgreement).Execute()

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
    createDataProcessingAgreement := *openapiclient.NewCreateDataProcessingAgreement("Version_example", "Body_example", "AttachmentUrl_example") // CreateDataProcessingAgreement | Create data processing agreement payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.CreateDataProcessingAgreement(context.Background()).CreateDataProcessingAgreement(createDataProcessingAgreement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.CreateDataProcessingAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDataProcessingAgreement`: DataProcessingAgreement
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingAgreementsApi.CreateDataProcessingAgreement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataProcessingAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDataProcessingAgreement** | [**CreateDataProcessingAgreement**](CreateDataProcessingAgreement.md) | Create data processing agreement payload | 

### Return type

[**DataProcessingAgreement**](DataProcessingAgreement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataProcessingAgreementById

> DeleteDataProcessingAgreementById(ctx, id).Execute()

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Agreement id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.DeleteDataProcessingAgreementById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.DeleteDataProcessingAgreementById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Agreement id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataProcessingAgreementByIdRequest struct via the builder pattern


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


## GetDataProcessingAgreementById

> DataProcessingAgreement GetDataProcessingAgreementById(ctx, id).Execute()

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Agreement id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.GetDataProcessingAgreementById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.GetDataProcessingAgreementById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataProcessingAgreementById`: DataProcessingAgreement
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingAgreementsApi.GetDataProcessingAgreementById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Agreement id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataProcessingAgreementByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataProcessingAgreement**](DataProcessingAgreement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataProcessingAgreements

> DataProcessingAgreementPaginatedList GetDataProcessingAgreements(ctx).PageNo(pageNo).PageSize(pageSize).Execute()

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
    pageNo := int32(56) // int32 | Page number. (optional) (default to 1)
    pageSize := int32(56) // int32 | Number of items to return. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.GetDataProcessingAgreements(context.Background()).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.GetDataProcessingAgreements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataProcessingAgreements`: DataProcessingAgreementPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingAgreementsApi.GetDataProcessingAgreements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataProcessingAgreementsRequest struct via the builder pattern


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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateDataProcessingAgreementById

> TerminateDataProcessingAgreementById(ctx, id).Execute()

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Agreement id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.TerminateDataProcessingAgreementById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.TerminateDataProcessingAgreementById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Agreement id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateDataProcessingAgreementByIdRequest struct via the builder pattern


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


## UpdateDataProcessingAgreement

> DataProcessingAgreement UpdateDataProcessingAgreement(ctx, id).UpdateDataProcessingAgreement(updateDataProcessingAgreement).Execute()

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Agreement id.
    updateDataProcessingAgreement := *openapiclient.NewUpdateDataProcessingAgreement("Version_example", "Body_example", "AttachmentUrl_example") // UpdateDataProcessingAgreement | Update data processing agreement payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.UpdateDataProcessingAgreement(context.Background(), id).UpdateDataProcessingAgreement(updateDataProcessingAgreement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.UpdateDataProcessingAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDataProcessingAgreement`: DataProcessingAgreement
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingAgreementsApi.UpdateDataProcessingAgreement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Agreement id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataProcessingAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDataProcessingAgreement** | [**UpdateDataProcessingAgreement**](UpdateDataProcessingAgreement.md) | Update data processing agreement payload | 

### Return type

[**DataProcessingAgreement**](DataProcessingAgreement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

