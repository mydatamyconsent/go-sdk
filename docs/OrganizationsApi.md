# \OrganizationsApi

All URIs are relative to *https://api.mydatamyconsent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1OrganizationsConsentRequestsGet**](OrganizationsApi.md#V1OrganizationsConsentRequestsGet) | **Get** /v1/organizations/consent-requests | Get all consent requests sent to Organizations.
[**V1OrganizationsConsentRequestsPost**](OrganizationsApi.md#V1OrganizationsConsentRequestsPost) | **Post** /v1/organizations/consent-requests | Create consent request for an Organization.
[**V1OrganizationsConsentRequestsRequestIdCancelPut**](OrganizationsApi.md#V1OrganizationsConsentRequestsRequestIdCancelPut) | **Put** /v1/organizations/consent-requests/{request_id}/cancel | Cancel the  Organization data request by id.
[**V1OrganizationsConsentRequestsRequestIdGet**](OrganizationsApi.md#V1OrganizationsConsentRequestsRequestIdGet) | **Get** /v1/organizations/consent-requests/{request_id} | Get Organization data request by id.
[**V1OrganizationsConsentTemplatesGet**](OrganizationsApi.md#V1OrganizationsConsentTemplatesGet) | **Get** /v1/organizations/consent-templates | Get the paginated list of organization consent templates.
[**V1OrganizationsConsentTemplatesTemplateIdGet**](OrganizationsApi.md#V1OrganizationsConsentTemplatesTemplateIdGet) | **Get** /v1/organizations/consent-templates/{template_id} | Get Organization consent template details by consent id.
[**V1OrganizationsConsentsConsentIdDocumentsDocumentIdDownloadGet**](OrganizationsApi.md#V1OrganizationsConsentsConsentIdDocumentsDocumentIdDownloadGet) | **Get** /v1/organizations/consents/{consent_id}/documents/{document_id}/download | Download Organization consent document by document id.
[**V1OrganizationsConsentsConsentIdDocumentsDocumentIdGet**](OrganizationsApi.md#V1OrganizationsConsentsConsentIdDocumentsDocumentIdGet) | **Get** /v1/organizations/consents/{consent_id}/documents/{document_id} | Get Organization consent document based on document id and consent id.
[**V1OrganizationsConsentsConsentIdDocumentsGet**](OrganizationsApi.md#V1OrganizationsConsentsConsentIdDocumentsGet) | **Get** /v1/organizations/consents/{consent_id}/documents | Get Organization consent document by consent id.
[**V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdGet**](OrganizationsApi.md#V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdGet) | **Get** /v1/organizations/consents/{consent_id}/financial-accounts/{account_id} | Get organization consented financial account details.
[**V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet**](OrganizationsApi.md#V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet) | **Get** /v1/organizations/consents/{consent_id}/financial-accounts/{account_id}/transactions | Get organization consented financial account transactions.
[**V1OrganizationsConsentsConsentIdFinancialAccountsGet**](OrganizationsApi.md#V1OrganizationsConsentsConsentIdFinancialAccountsGet) | **Get** /v1/organizations/consents/{consent_id}/financial-accounts | Get all organization consented financial accounts.
[**V1OrganizationsConsentsConsentIdGet**](OrganizationsApi.md#V1OrganizationsConsentsConsentIdGet) | **Get** /v1/organizations/consents/{consent_id} | Get Organization consent request details by consent id.
[**V1OrganizationsConsentsGet**](OrganizationsApi.md#V1OrganizationsConsentsGet) | **Get** /v1/organizations/consents | Get the paginated list of Organization consents.
[**V1OrganizationsDocumentsIssuePost**](OrganizationsApi.md#V1OrganizationsDocumentsIssuePost) | **Post** /v1/organizations/documents/issue | Issue a new document to an organization.
[**V1OrganizationsDocumentsIssueUploadIssueRequestIdPost**](OrganizationsApi.md#V1OrganizationsDocumentsIssueUploadIssueRequestIdPost) | **Post** /v1/organizations/documents/issue/upload/{issue_request_id} | Upload a document for issuance request of Organization.
[**V1OrganizationsDocumentsIssuedDocumentIdGet**](OrganizationsApi.md#V1OrganizationsDocumentsIssuedDocumentIdGet) | **Get** /v1/organizations/documents/issued/{document_id} | Get issued document.
[**V1OrganizationsDocumentsIssuedGet**](OrganizationsApi.md#V1OrganizationsDocumentsIssuedGet) | **Get** /v1/organizations/documents/issued | Get paginated list of issued documents of given document type.
[**V1OrganizationsDocumentsTypesGet**](OrganizationsApi.md#V1OrganizationsDocumentsTypesGet) | **Get** /v1/organizations/documents/types | Get paginated list of registered document types.



## V1OrganizationsConsentRequestsGet

> PaginatedListOfOrganizationConsentRequestDetailss V1OrganizationsConsentRequestsGet(ctx).Status(status).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()

Get all consent requests sent to Organizations.

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
    status := openapiclient.DataConsentStatus("Pending") // DataConsentStatus |  (optional)
    fromDateTime := time.Now() // time.Time |  (optional)
    toDateTime := time.Now() // time.Time |  (optional)
    pageNo := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsConsentRequestsGet(context.Background()).Status(status).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsConsentRequestsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsConsentRequestsGet`: PaginatedListOfOrganizationConsentRequestDetailss
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsConsentRequestsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsConsentRequestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 
 **fromDateTime** | **time.Time** |  | 
 **toDateTime** | **time.Time** |  | 
 **pageNo** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**PaginatedListOfOrganizationConsentRequestDetailss**](PaginatedListOfOrganizationConsentRequestDetailss.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsConsentRequestsPost

> OrganizationConsentRequestDetails V1OrganizationsConsentRequestsPost(ctx).CreateConsentRequest(createConsentRequest).Execute()

Create consent request for an Organization.

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
    createConsentRequest := *openapiclient.NewCreateConsentRequest("ConsentTemplateId_example", *openapiclient.NewConsentRequestReceiver("CountryIso2Code_example", []openapiclient.KeyValuePair{*openapiclient.NewKeyValuePair("Key_example", "Value_example")}, openapiclient.IdentificationStrategy("MatchAtLeastOneIdentifier"))) // CreateConsentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsConsentRequestsPost(context.Background()).CreateConsentRequest(createConsentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsConsentRequestsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsConsentRequestsPost`: OrganizationConsentRequestDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsConsentRequestsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsConsentRequestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConsentRequest** | [**CreateConsentRequest**](CreateConsentRequest.md) |  | 

### Return type

[**OrganizationConsentRequestDetails**](OrganizationConsentRequestDetails.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsConsentRequestsRequestIdCancelPut

> bool V1OrganizationsConsentRequestsRequestIdCancelPut(ctx, requestId).Execute()

Cancel the  Organization data request by id.

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
    requestId := "requestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsConsentRequestsRequestIdCancelPut(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsConsentRequestsRequestIdCancelPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsConsentRequestsRequestIdCancelPut`: bool
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsConsentRequestsRequestIdCancelPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsConsentRequestsRequestIdCancelPutRequest struct via the builder pattern


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


## V1OrganizationsConsentRequestsRequestIdGet

> ConsentRequest V1OrganizationsConsentRequestsRequestIdGet(ctx, requestId).Execute()

Get Organization data request by id.

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
    requestId := "requestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsConsentRequestsRequestIdGet(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsConsentRequestsRequestIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsConsentRequestsRequestIdGet`: ConsentRequest
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsConsentRequestsRequestIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsConsentRequestsRequestIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsentRequest**](ConsentRequest.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsConsentTemplatesGet

> PaginatedListOfConsentRequestTemplates V1OrganizationsConsentTemplatesGet(ctx).PageNo(pageNo).PageSize(pageSize).Execute()

Get the paginated list of organization consent templates.

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
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsConsentTemplatesGet(context.Background()).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsConsentTemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsConsentTemplatesGet`: PaginatedListOfConsentRequestTemplates
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsConsentTemplatesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsConsentTemplatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNo** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**PaginatedListOfConsentRequestTemplates**](PaginatedListOfConsentRequestTemplates.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsConsentTemplatesTemplateIdGet

> OrganizationConsentRequestTemplateDetails V1OrganizationsConsentTemplatesTemplateIdGet(ctx, templateId).Execute()

Get Organization consent template details by consent id.

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
    templateId := "templateId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsConsentTemplatesTemplateIdGet(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsConsentTemplatesTemplateIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsConsentTemplatesTemplateIdGet`: OrganizationConsentRequestTemplateDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsConsentTemplatesTemplateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsConsentTemplatesTemplateIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationConsentRequestTemplateDetails**](OrganizationConsentRequestTemplateDetails.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsConsentsConsentIdDocumentsDocumentIdDownloadGet

> *os.File V1OrganizationsConsentsConsentIdDocumentsDocumentIdDownloadGet(ctx, consentId, documentId).Execute()

Download Organization consent document by document id.

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
    consentId := "consentId_example" // string | 
    documentId := "documentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsConsentsConsentIdDocumentsDocumentIdDownloadGet(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsConsentsConsentIdDocumentsDocumentIdDownloadGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsConsentsConsentIdDocumentsDocumentIdDownloadGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsConsentsConsentIdDocumentsDocumentIdDownloadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsConsentsConsentIdDocumentsDocumentIdDownloadGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsConsentsConsentIdDocumentsDocumentIdGet

> ConsentedDocument V1OrganizationsConsentsConsentIdDocumentsDocumentIdGet(ctx, consentId, documentId).Execute()

Get Organization consent document based on document id and consent id.

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
    consentId := "consentId_example" // string | 
    documentId := "documentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsConsentsConsentIdDocumentsDocumentIdGet(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsConsentsConsentIdDocumentsDocumentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsConsentsConsentIdDocumentsDocumentIdGet`: ConsentedDocument
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsConsentsConsentIdDocumentsDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsConsentsConsentIdDocumentsDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConsentedDocument**](ConsentedDocument.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsConsentsConsentIdDocumentsGet

> []ConsentedDocument V1OrganizationsConsentsConsentIdDocumentsGet(ctx, consentId).Execute()

Get Organization consent document by consent id.

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
    consentId := "consentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsConsentsConsentIdDocumentsGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsConsentsConsentIdDocumentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsConsentsConsentIdDocumentsGet`: []ConsentedDocument
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsConsentsConsentIdDocumentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsConsentsConsentIdDocumentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ConsentedDocument**](ConsentedDocument.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdGet

> FinancialAccount V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdGet(ctx, consentId, accountId).Execute()

Get organization consented financial account details.

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
    consentId := "consentId_example" // string | 
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdGet(context.Background(), consentId, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdGet`: FinancialAccount
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsConsentsConsentIdFinancialAccountsAccountIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FinancialAccount**](FinancialAccount.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet

> PaginatedListOfFinancialAccountTransactions V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet(ctx, consentId, accountId).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()

Get organization consented financial account transactions.

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
    consentId := "consentId_example" // string | 
    accountId := "accountId_example" // string | 
    fromDateTime := time.Now() // time.Time |  (optional)
    toDateTime := time.Now() // time.Time |  (optional)
    pageNo := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet(context.Background(), consentId, accountId).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet`: PaginatedListOfFinancialAccountTransactions
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsConsentsConsentIdFinancialAccountsAccountIdTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromDateTime** | **time.Time** |  | 
 **toDateTime** | **time.Time** |  | 
 **pageNo** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**PaginatedListOfFinancialAccountTransactions**](PaginatedListOfFinancialAccountTransactions.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsConsentsConsentIdFinancialAccountsGet

> []FinancialAccount V1OrganizationsConsentsConsentIdFinancialAccountsGet(ctx, consentId).Execute()

Get all organization consented financial accounts.

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
    consentId := "consentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsConsentsConsentIdFinancialAccountsGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsConsentsConsentIdFinancialAccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsConsentsConsentIdFinancialAccountsGet`: []FinancialAccount
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsConsentsConsentIdFinancialAccountsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsConsentsConsentIdFinancialAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FinancialAccount**](FinancialAccount.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsConsentsConsentIdGet

> ConsentDetails V1OrganizationsConsentsConsentIdGet(ctx, consentId).Execute()

Get Organization consent request details by consent id.

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
    consentId := "consentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsConsentsConsentIdGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsConsentsConsentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsConsentsConsentIdGet`: ConsentDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsConsentsConsentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsConsentsConsentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsentDetails**](ConsentDetails.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsConsentsGet

> PaginatedListOfConsents V1OrganizationsConsentsGet(ctx).Status(status).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()

Get the paginated list of Organization consents.

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
    status := openapiclient.DataConsentStatus("Pending") // DataConsentStatus |  (optional)
    fromDateTime := time.Now() // time.Time |  (optional)
    toDateTime := time.Now() // time.Time |  (optional)
    pageNo := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsConsentsGet(context.Background()).Status(status).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsConsentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsConsentsGet`: PaginatedListOfConsents
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsConsentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsConsentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 
 **fromDateTime** | **time.Time** |  | 
 **toDateTime** | **time.Time** |  | 
 **pageNo** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**PaginatedListOfConsents**](PaginatedListOfConsents.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsDocumentsIssuePost

> DocumentIssueRequestDetails V1OrganizationsDocumentsIssuePost(ctx).DocumentIssueRequest(documentIssueRequest).Execute()

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
    documentIssueRequest := *openapiclient.NewDocumentIssueRequest("DocumentTypeId_example", "Identifier_example", "Description_example", *openapiclient.NewDocumentIssueRequestReceiver("CountryIso2Code_example", []openapiclient.KeyValuePair{*openapiclient.NewKeyValuePair("Key_example", "Value_example")}, openapiclient.IdentificationStrategy("MatchAtLeastOneIdentifier")), time.Now(), time.Now()) // DocumentIssueRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsDocumentsIssuePost(context.Background()).DocumentIssueRequest(documentIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsDocumentsIssuePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsDocumentsIssuePost`: DocumentIssueRequestDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsDocumentsIssuePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsDocumentsIssuePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentIssueRequest** | [**DocumentIssueRequest**](DocumentIssueRequest.md) |  | 

### Return type

[**DocumentIssueRequestDetails**](DocumentIssueRequestDetails.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsDocumentsIssueUploadIssueRequestIdPost

> UploadDocumentResponse V1OrganizationsDocumentsIssueUploadIssueRequestIdPost(ctx, issueRequestId).File(file).Execute()

Upload a document for issuance request of Organization.

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
    issueRequestId := "issueRequestId_example" // string | 
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsDocumentsIssueUploadIssueRequestIdPost(context.Background(), issueRequestId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsDocumentsIssueUploadIssueRequestIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsDocumentsIssueUploadIssueRequestIdPost`: UploadDocumentResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsDocumentsIssueUploadIssueRequestIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsDocumentsIssueUploadIssueRequestIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 

### Return type

[**UploadDocumentResponse**](UploadDocumentResponse.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsDocumentsIssuedDocumentIdGet

> IssuedDocument V1OrganizationsDocumentsIssuedDocumentIdGet(ctx, documentId).Execute()

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
    documentId := "documentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsDocumentsIssuedDocumentIdGet(context.Background(), documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsDocumentsIssuedDocumentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsDocumentsIssuedDocumentIdGet`: IssuedDocument
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsDocumentsIssuedDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsDocumentsIssuedDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IssuedDocument**](IssuedDocument.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsDocumentsIssuedGet

> PaginatedListOfIssuedDocuments V1OrganizationsDocumentsIssuedGet(ctx).DocumentTypeId(documentTypeId).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()

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
    documentTypeId := "documentTypeId_example" // string | 
    fromDateTime := time.Now() // time.Time |  (optional)
    toDateTime := time.Now() // time.Time |  (optional)
    pageNo := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsDocumentsIssuedGet(context.Background()).DocumentTypeId(documentTypeId).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsDocumentsIssuedGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsDocumentsIssuedGet`: PaginatedListOfIssuedDocuments
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsDocumentsIssuedGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsDocumentsIssuedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentTypeId** | **string** |  | 
 **fromDateTime** | **time.Time** |  | 
 **toDateTime** | **time.Time** |  | 
 **pageNo** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**PaginatedListOfIssuedDocuments**](PaginatedListOfIssuedDocuments.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsDocumentsTypesGet

> PaginatedListOfDocumentTypes V1OrganizationsDocumentsTypesGet(ctx).PageNo(pageNo).PageSize(pageSize).Execute()

Get paginated list of registered document types.

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
    resp, r, err := apiClient.OrganizationsApi.V1OrganizationsDocumentsTypesGet(context.Background()).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.V1OrganizationsDocumentsTypesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsDocumentsTypesGet`: PaginatedListOfDocumentTypes
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.V1OrganizationsDocumentsTypesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsDocumentsTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNo** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**PaginatedListOfDocumentTypes**](PaginatedListOfDocumentTypes.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

