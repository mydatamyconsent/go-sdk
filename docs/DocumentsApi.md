# MyDataMyConsent\DocumentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssueDocument**](DocumentsApi.md#IssueDocument) | **Post** /v1/documents/issue | Issue a new document.
[**V1DocumentsIssuedDocumentIdGet**](DocumentsApi.md#V1DocumentsIssuedDocumentIdGet) | **Get** /v1/documents/issued/{documentId} | Get issued document.
[**V1DocumentsIssuedGet**](DocumentsApi.md#V1DocumentsIssuedGet) | **Get** /v1/documents/issued | Get issued documents.
[**V1DocumentsTypesGet**](DocumentsApi.md#V1DocumentsTypesGet) | **Get** /v1/documents/types | Get registered document types.



## IssueDocument

> bool IssueDocument(ctx).DocumentIssueRequest(documentIssueRequest).Execute()

Issue a new document.

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
    documentIssueRequest := *openapiclient.NewDocumentIssueRequest("DocumentTypeId_example", "Identifier_example", "Name_example", "Description_example", *openapiclient.NewReceiver(), "Base64PDFDocument_example") // DocumentIssueRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.IssueDocument(context.Background()).DocumentIssueRequest(documentIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.IssueDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssueDocument`: bool
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.IssueDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentIssueRequest** | [**DocumentIssueRequest**](DocumentIssueRequest.md) |  | 

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


## V1DocumentsIssuedDocumentIdGet

> V1DocumentsIssuedDocumentIdGet(ctx, documentId).Execute()

Get issued document.

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
    documentId := TODO // string | Document id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.V1DocumentsIssuedDocumentIdGet(context.Background(), documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.V1DocumentsIssuedDocumentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | [**string**](.md) | Document id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DocumentsIssuedDocumentIdGetRequest struct via the builder pattern


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


## V1DocumentsIssuedGet

> V1DocumentsIssuedGet(ctx).DocumentTypeId(documentTypeId).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageSize(pageSize).PageNo(pageNo).Execute()

Get issued documents.

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
    documentTypeId := TODO // string |  (optional)
    fromDateTime := time.Now() // time.Time |  (optional)
    toDateTime := time.Now() // time.Time |  (optional)
    pageSize := int32(56) // int32 |  (optional) (default to 25)
    pageNo := int32(56) // int32 |  (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.V1DocumentsIssuedGet(context.Background()).DocumentTypeId(documentTypeId).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageSize(pageSize).PageNo(pageNo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.V1DocumentsIssuedGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DocumentsIssuedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentTypeId** | [**string**](string.md) |  | 
 **fromDateTime** | **time.Time** |  | 
 **toDateTime** | **time.Time** |  | 
 **pageSize** | **int32** |  | [default to 25]
 **pageNo** | **int32** |  | [default to 1]

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


## V1DocumentsTypesGet

> V1DocumentsTypesGet(ctx).PageSize(pageSize).PageNo(pageNo).Execute()

Get registered document types.

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
    pageSize := int32(56) // int32 |  (optional)
    pageNo := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.V1DocumentsTypesGet(context.Background()).PageSize(pageSize).PageNo(pageNo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.V1DocumentsTypesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DocumentsTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** |  | 
 **pageNo** | **int32** |  | 

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

