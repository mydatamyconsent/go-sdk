# \DocumentsApi

All URIs are relative to *https://api.mydatamyconsent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIssuedDocumentById**](DocumentsApi.md#GetIssuedDocumentById) | **Get** /v1/documents/issued/{documentId} | Get issued document.
[**GetIssuedDocuments**](DocumentsApi.md#GetIssuedDocuments) | **Get** /v1/documents/issued/{documentTypeId} | Get paginated list of issued documents of given document type.
[**GetRegisteredDocumentTypes**](DocumentsApi.md#GetRegisteredDocumentTypes) | **Get** /v1/documents/types | Get registered document types.
[**IssueDocumentToIndividual**](DocumentsApi.md#IssueDocumentToIndividual) | **Post** /v1/documents/issue/individual | Issue a new document to an individual user.
[**IssueDocumentToOrganization**](DocumentsApi.md#IssueDocumentToOrganization) | **Post** /v1/documents/issue/organization | Issue a new document to an organization.
[**UploadDocumentForIndividual**](DocumentsApi.md#UploadDocumentForIndividual) | **Post** /v1/documents/issue/individual/upload/{issueRequestId} | Upload a document for issuance request of individual.
[**UploadDocumentForOrganization**](DocumentsApi.md#UploadDocumentForOrganization) | **Post** /v1/documents/issue/organization/upload/{issueRequestId} | Upload a document for issuance request of organization.



## GetIssuedDocumentById

> IssuedDocument GetIssuedDocumentById(ctx, documentId).Execute()

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
    // response from `GetIssuedDocumentById`: IssuedDocument
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetIssuedDocumentById`: %v\n", resp)
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

[**IssuedDocument**](IssuedDocument.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssuedDocuments

> IssuedDocumentPaginatedList GetIssuedDocuments(ctx, documentTypeId).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()

Get paginated list of issued documents of given document type.

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
    documentTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Document type id.
    fromDateTime := time.Now() // time.Time | From DateTime. (optional)
    toDateTime := time.Now() // time.Time | To DateTime. (optional)
    pageNo := int32(56) // int32 | Page number. (optional) (default to 1)
    pageSize := int32(56) // int32 | Number of items to return. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.GetIssuedDocuments(context.Background(), documentTypeId).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetIssuedDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIssuedDocuments`: IssuedDocumentPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetIssuedDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentTypeId** | **string** | Document type id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuedDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDateTime** | **time.Time** | From DateTime. | 
 **toDateTime** | **time.Time** | To DateTime. | 
 **pageNo** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Number of items to return. | [default to 25]

### Return type

[**IssuedDocumentPaginatedList**](IssuedDocumentPaginatedList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegisteredDocumentTypes

> DocumentTypePaginatedList GetRegisteredDocumentTypes(ctx).PageNo(pageNo).PageSize(pageSize).Execute()

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
    pageNo := int32(56) // int32 | Page number. (optional) (default to 1)
    pageSize := int32(56) // int32 | Number of items to return. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.GetRegisteredDocumentTypes(context.Background()).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetRegisteredDocumentTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegisteredDocumentTypes`: DocumentTypePaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetRegisteredDocumentTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredDocumentTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNo** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Number of items to return. | [default to 25]

### Return type

[**DocumentTypePaginatedList**](DocumentTypePaginatedList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueDocumentToIndividual

> DocumentIssueRequestDetails IssueDocumentToIndividual(ctx).DocumentIssueRequest(documentIssueRequest).Execute()

Issue a new document to an individual user.

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
    documentIssueRequest := *openapiclient.NewDocumentIssueRequest("DocumentTypeId_example", "Identifier_example", "Description_example", *openapiclient.NewDocumentReceiver("CountryIso2Code_example", []openapiclient.StringStringKeyValuePair{*openapiclient.NewStringStringKeyValuePair()}, openapiclient.IdentificationStrategy("MatchAtLeastOneIdentifier")), time.Now(), time.Now()) // DocumentIssueRequest | Document issue request MyDataMyConsent.DeveloperApi.Models.DocumentIssueRequest.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.IssueDocumentToIndividual(context.Background()).DocumentIssueRequest(documentIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.IssueDocumentToIndividual``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssueDocumentToIndividual`: DocumentIssueRequestDetails
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.IssueDocumentToIndividual`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueDocumentToIndividualRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentIssueRequest** | [**DocumentIssueRequest**](DocumentIssueRequest.md) | Document issue request MyDataMyConsent.DeveloperApi.Models.DocumentIssueRequest. | 

### Return type

[**DocumentIssueRequestDetails**](DocumentIssueRequestDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueDocumentToOrganization

> DocumentIssueRequestDetails IssueDocumentToOrganization(ctx).DocumentIssueRequest(documentIssueRequest).Execute()

Issue a new document to an organization.

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
    documentIssueRequest := *openapiclient.NewDocumentIssueRequest("DocumentTypeId_example", "Identifier_example", "Description_example", *openapiclient.NewDocumentReceiver("CountryIso2Code_example", []openapiclient.StringStringKeyValuePair{*openapiclient.NewStringStringKeyValuePair()}, openapiclient.IdentificationStrategy("MatchAtLeastOneIdentifier")), time.Now(), time.Now()) // DocumentIssueRequest | Document issue request MyDataMyConsent.DeveloperApi.Models.DocumentIssueRequest.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.IssueDocumentToOrganization(context.Background()).DocumentIssueRequest(documentIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.IssueDocumentToOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssueDocumentToOrganization`: DocumentIssueRequestDetails
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.IssueDocumentToOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueDocumentToOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentIssueRequest** | [**DocumentIssueRequest**](DocumentIssueRequest.md) | Document issue request MyDataMyConsent.DeveloperApi.Models.DocumentIssueRequest. | 

### Return type

[**DocumentIssueRequestDetails**](DocumentIssueRequestDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadDocumentForIndividual

> string UploadDocumentForIndividual(ctx, issueRequestId).FormFile(formFile).Execute()

Upload a document for issuance request of individual.

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
    issueRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Issue Request Id System.Guid.
    formFile := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.UploadDocumentForIndividual(context.Background(), issueRequestId).FormFile(formFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.UploadDocumentForIndividual``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadDocumentForIndividual`: string
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.UploadDocumentForIndividual`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueRequestId** | **string** | Issue Request Id System.Guid. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadDocumentForIndividualRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **formFile** | ***os.File** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadDocumentForOrganization

> string UploadDocumentForOrganization(ctx, issueRequestId).FormFile(formFile).Execute()

Upload a document for issuance request of organization.

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
    issueRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Issue Request Id System.Guid.
    formFile := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.UploadDocumentForOrganization(context.Background(), issueRequestId).FormFile(formFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.UploadDocumentForOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadDocumentForOrganization`: string
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.UploadDocumentForOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueRequestId** | **string** | Issue Request Id System.Guid. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadDocumentForOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **formFile** | ***os.File** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

