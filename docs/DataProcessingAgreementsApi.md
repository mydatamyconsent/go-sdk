# \DataProcessingAgreementsApi

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

> DataProcessingAgreementDtoPaginatedList V1DataAgreementsGet(ctx).PageNo(pageNo).PageSize(pageSize).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.V1DataAgreementsGet(context.Background()).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.V1DataAgreementsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DataAgreementsGet`: DataProcessingAgreementDtoPaginatedList
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

[**DataProcessingAgreementDtoPaginatedList**](DataProcessingAgreementDtoPaginatedList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.V1DataAgreementsIdDelete(context.Background(), id).Execute()
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
**id** | **string** |  | 

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataAgreementsIdGet

> DataProcessingAgreementDto V1DataAgreementsIdGet(ctx, id).Execute()

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.V1DataAgreementsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.V1DataAgreementsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DataAgreementsIdGet`: DataProcessingAgreementDto
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

[**DataProcessingAgreementDto**](DataProcessingAgreementDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataAgreementsIdPut

> DataProcessingAgreementDto V1DataAgreementsIdPut(ctx, id).UpdateDataProcessingAgreementRequestModel(updateDataProcessingAgreementRequestModel).Execute()

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    updateDataProcessingAgreementRequestModel := *openapiclient.NewUpdateDataProcessingAgreementRequestModel("Version_example", "Body_example", "AttachmentUrl_example") // UpdateDataProcessingAgreementRequestModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.V1DataAgreementsIdPut(context.Background(), id).UpdateDataProcessingAgreementRequestModel(updateDataProcessingAgreementRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.V1DataAgreementsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DataAgreementsIdPut`: DataProcessingAgreementDto
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingAgreementsApi.V1DataAgreementsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DataAgreementsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDataProcessingAgreementRequestModel** | [**UpdateDataProcessingAgreementRequestModel**](UpdateDataProcessingAgreementRequestModel.md) |  | 

### Return type

[**DataProcessingAgreementDto**](DataProcessingAgreementDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.V1DataAgreementsIdTerminatePut(context.Background(), id).Execute()
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
**id** | **string** |  | 

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataAgreementsPost

> DataProcessingAgreementDto V1DataAgreementsPost(ctx).CreateDataProcessingAgreementRequestModel(createDataProcessingAgreementRequestModel).Execute()

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
    createDataProcessingAgreementRequestModel := *openapiclient.NewCreateDataProcessingAgreementRequestModel("Version_example", "Body_example", "AttachmentUrl_example") // CreateDataProcessingAgreementRequestModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProcessingAgreementsApi.V1DataAgreementsPost(context.Background()).CreateDataProcessingAgreementRequestModel(createDataProcessingAgreementRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProcessingAgreementsApi.V1DataAgreementsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DataAgreementsPost`: DataProcessingAgreementDto
    fmt.Fprintf(os.Stdout, "Response from `DataProcessingAgreementsApi.V1DataAgreementsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DataAgreementsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDataProcessingAgreementRequestModel** | [**CreateDataProcessingAgreementRequestModel**](CreateDataProcessingAgreementRequestModel.md) |  | 

### Return type

[**DataProcessingAgreementDto**](DataProcessingAgreementDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

