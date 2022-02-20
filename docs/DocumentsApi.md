# \DocumentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIssuedDocumentById**](DocumentsApi.md#GetIssuedDocumentById) | **Get** /v1/documents/issued/{documentId} | Get issued document.
[**GetIssuedDocuments**](DocumentsApi.md#GetIssuedDocuments) | **Get** /v1/documents/issued | Get issued documents.
[**GetRegisteredDocumentTypes**](DocumentsApi.md#GetRegisteredDocumentTypes) | **Get** /v1/documents/types | Get registered document types.
[**IssueDocument**](DocumentsApi.md#IssueDocument) | **Post** /v1/documents/issue | Issue a new document.



## GetIssuedDocumentById

> GetIssuedDocumentById(ctx, documentId).Execute()

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
    documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Document id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.GetIssuedDocumentById(context.Background(), documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetIssuedDocumentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** | Document id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuedDocumentByIdRequest struct via the builder pattern


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


## GetIssuedDocuments

> GetIssuedDocuments(ctx).DocumentTypeId(documentTypeId).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageSize(pageSize).PageNo(pageNo).Execute()

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
    documentTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    fromDateTime := time.Now() // time.Time |  (optional)
    toDateTime := time.Now() // time.Time |  (optional)
    pageSize := int32(56) // int32 |  (optional) (default to 25)
    pageNo := int32(56) // int32 |  (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.GetIssuedDocuments(context.Background()).DocumentTypeId(documentTypeId).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageSize(pageSize).PageNo(pageNo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetIssuedDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuedDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentTypeId** | **string** |  | 
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


## GetRegisteredDocumentTypes

> GetRegisteredDocumentTypes(ctx).PageSize(pageSize).PageNo(pageNo).Execute()

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
    pageSize := int32(56) // int32 |  (optional) (default to 25)
    pageNo := int32(56) // int32 |  (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.GetRegisteredDocumentTypes(context.Background()).PageSize(pageSize).PageNo(pageNo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetRegisteredDocumentTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredDocumentTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
    documentIssueRequest := *openapiclient.NewDocumentIssueRequest("DocumentTypeId_example", "DocumentIdentifier_example", "Name_example", "Description_example", *openapiclient.NewReceiver(), "Base64PdfDocument_example") // DocumentIssueRequest | Document issue request MyDataMyConsent.Models.Documents.DocumentIssueRequest.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.IssueDocument(context.Background()).DocumentIssueRequest(documentIssueRequest).Execute()
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
 **documentIssueRequest** | [**DocumentIssueRequest**](DocumentIssueRequest.md) | Document issue request MyDataMyConsent.Models.Documents.DocumentIssueRequest. | 

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

