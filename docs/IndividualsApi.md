# \IndividualsApi

All URIs are relative to *https://api.mydatamyconsent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssuerIssuedoc1XmlPost**](IndividualsApi.md#IssuerIssuedoc1XmlPost) | **Post** /issuer/issuedoc/1/xml | Digilocker compatible endpoint to issue document.
[**V1IndividualsConsentRequestsGet**](IndividualsApi.md#V1IndividualsConsentRequestsGet) | **Get** /v1/individuals/consent-requests | Get all consent requests sent to Individuals.
[**V1IndividualsConsentRequestsPost**](IndividualsApi.md#V1IndividualsConsentRequestsPost) | **Post** /v1/individuals/consent-requests | Create individual consent request.
[**V1IndividualsConsentRequestsRequestIdCancelPut**](IndividualsApi.md#V1IndividualsConsentRequestsRequestIdCancelPut) | **Put** /v1/individuals/consent-requests/{request_id}/cancel | Cancel the Individual data request by id.
[**V1IndividualsConsentRequestsRequestIdGet**](IndividualsApi.md#V1IndividualsConsentRequestsRequestIdGet) | **Get** /v1/individuals/consent-requests/{request_id} | Get Individual data consent request by id.
[**V1IndividualsConsentTemplatesGet**](IndividualsApi.md#V1IndividualsConsentTemplatesGet) | **Get** /v1/individuals/consent-templates | Get the paginated list of individual consent templates.
[**V1IndividualsConsentTemplatesTemplateIdGet**](IndividualsApi.md#V1IndividualsConsentTemplatesTemplateIdGet) | **Get** /v1/individuals/consent-templates/{template_id} | Get Individual consent template details by consent id.
[**V1IndividualsConsentsConsentIdDocumentsDocumentIdDownloadGet**](IndividualsApi.md#V1IndividualsConsentsConsentIdDocumentsDocumentIdDownloadGet) | **Get** /v1/individuals/consents/{consent_id}/documents/{document_id}/download | Download Individual consented document by document id.
[**V1IndividualsConsentsConsentIdDocumentsDocumentIdGet**](IndividualsApi.md#V1IndividualsConsentsConsentIdDocumentsDocumentIdGet) | **Get** /v1/individuals/consents/{consent_id}/documents/{document_id} | Get Individual consented document by document id.
[**V1IndividualsConsentsConsentIdDocumentsGet**](IndividualsApi.md#V1IndividualsConsentsConsentIdDocumentsGet) | **Get** /v1/individuals/consents/{consent_id}/documents | Get Individual consented document by consent id.
[**V1IndividualsConsentsConsentIdFinancialAccountsAccountIdGet**](IndividualsApi.md#V1IndividualsConsentsConsentIdFinancialAccountsAccountIdGet) | **Get** /v1/individuals/consents/{consent_id}/financial-accounts/{account_id} | Get individual consented financial account details.
[**V1IndividualsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet**](IndividualsApi.md#V1IndividualsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet) | **Get** /v1/individuals/consents/{consent_id}/financial-accounts/{account_id}/transactions | Get individual consented financial account transactions.
[**V1IndividualsConsentsConsentIdFinancialAccountsGet**](IndividualsApi.md#V1IndividualsConsentsConsentIdFinancialAccountsGet) | **Get** /v1/individuals/consents/{consent_id}/financial-accounts | Get all individual consented financial accounts.
[**V1IndividualsConsentsConsentIdGet**](IndividualsApi.md#V1IndividualsConsentsConsentIdGet) | **Get** /v1/individuals/consents/{consent_id} | Get Individuals consent details by consent id.
[**V1IndividualsConsentsConsentIdHealthFhirBundleGet**](IndividualsApi.md#V1IndividualsConsentsConsentIdHealthFhirBundleGet) | **Get** /v1/individuals/consents/{consent_id}/health/fhir/bundle | Get Individual consented Health Records by consent id.
[**V1IndividualsConsentsGet**](IndividualsApi.md#V1IndividualsConsentsGet) | **Get** /v1/individuals/consents | Get the paginated list of Individual consents.
[**V1IndividualsDocumentsIssueIssueRequestIdUploadPost**](IndividualsApi.md#V1IndividualsDocumentsIssueIssueRequestIdUploadPost) | **Post** /v1/individuals/documents/issue/{issue_request_id}/upload | Upload a document for issuance request of individual.
[**V1IndividualsDocumentsIssuePost**](IndividualsApi.md#V1IndividualsDocumentsIssuePost) | **Post** /v1/individuals/documents/issue | Issue a new document to an individual user.
[**V1IndividualsDocumentsIssuedDocumentIdGet**](IndividualsApi.md#V1IndividualsDocumentsIssuedDocumentIdGet) | **Get** /v1/individuals/documents/issued/{document_id} | Get issued document.
[**V1IndividualsDocumentsIssuedGet**](IndividualsApi.md#V1IndividualsDocumentsIssuedGet) | **Get** /v1/individuals/documents/issued | Get paginated list of issued documents of given document type.
[**V1IndividualsDocumentsTypesGet**](IndividualsApi.md#V1IndividualsDocumentsTypesGet) | **Get** /v1/individuals/documents/types | Get paginated list of registered document types.



## IssuerIssuedoc1XmlPost

> PushUriResponse IssuerIssuedoc1XmlPost(ctx).PushUriRequest(pushUriRequest).Execute()

Digilocker compatible endpoint to issue document.

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
    pushUriRequest := *openapiclient.NewPushUriRequest(*openapiclient.NewUriDetails("Aadhaar_example", "Uri_example", "DocType_example", "DocName_example", "DocId_example", "IssuedOn_example", "ValidFrom_example")) // PushUriRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualsApi.IssuerIssuedoc1XmlPost(context.Background()).PushUriRequest(pushUriRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.IssuerIssuedoc1XmlPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuerIssuedoc1XmlPost`: PushUriResponse
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.IssuerIssuedoc1XmlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssuerIssuedoc1XmlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushUriRequest** | [**PushUriRequest**](PushUriRequest.md) |  | 

### Return type

[**PushUriResponse**](PushUriResponse.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1IndividualsConsentRequestsGet

> PaginatedListOfIndividualConsentRequestDetailss V1IndividualsConsentRequestsGet(ctx).Status(status).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()

Get all consent requests sent to Individuals.

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsConsentRequestsGet(context.Background()).Status(status).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsConsentRequestsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsConsentRequestsGet`: PaginatedListOfIndividualConsentRequestDetailss
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsConsentRequestsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsConsentRequestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 
 **fromDateTime** | **time.Time** |  | 
 **toDateTime** | **time.Time** |  | 
 **pageNo** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**PaginatedListOfIndividualConsentRequestDetailss**](PaginatedListOfIndividualConsentRequestDetailss.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1IndividualsConsentRequestsPost

> IndividualConsentRequestDetails V1IndividualsConsentRequestsPost(ctx).CreateConsentRequest(createConsentRequest).Execute()

Create individual consent request.

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsConsentRequestsPost(context.Background()).CreateConsentRequest(createConsentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsConsentRequestsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsConsentRequestsPost`: IndividualConsentRequestDetails
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsConsentRequestsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsConsentRequestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConsentRequest** | [**CreateConsentRequest**](CreateConsentRequest.md) |  | 

### Return type

[**IndividualConsentRequestDetails**](IndividualConsentRequestDetails.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1IndividualsConsentRequestsRequestIdCancelPut

> bool V1IndividualsConsentRequestsRequestIdCancelPut(ctx, requestId).Execute()

Cancel the Individual data request by id.

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsConsentRequestsRequestIdCancelPut(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsConsentRequestsRequestIdCancelPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsConsentRequestsRequestIdCancelPut`: bool
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsConsentRequestsRequestIdCancelPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsConsentRequestsRequestIdCancelPutRequest struct via the builder pattern


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


## V1IndividualsConsentRequestsRequestIdGet

> ConsentRequest V1IndividualsConsentRequestsRequestIdGet(ctx, requestId).Execute()

Get Individual data consent request by id.

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsConsentRequestsRequestIdGet(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsConsentRequestsRequestIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsConsentRequestsRequestIdGet`: ConsentRequest
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsConsentRequestsRequestIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsConsentRequestsRequestIdGetRequest struct via the builder pattern


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


## V1IndividualsConsentTemplatesGet

> PaginatedListOfConsentRequestTemplates V1IndividualsConsentTemplatesGet(ctx).PageNo(pageNo).PageSize(pageSize).Execute()

Get the paginated list of individual consent templates.

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsConsentTemplatesGet(context.Background()).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsConsentTemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsConsentTemplatesGet`: PaginatedListOfConsentRequestTemplates
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsConsentTemplatesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsConsentTemplatesGetRequest struct via the builder pattern


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


## V1IndividualsConsentTemplatesTemplateIdGet

> IndividualConsentRequestTemplateDetails V1IndividualsConsentTemplatesTemplateIdGet(ctx, templateId).Execute()

Get Individual consent template details by consent id.

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsConsentTemplatesTemplateIdGet(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsConsentTemplatesTemplateIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsConsentTemplatesTemplateIdGet`: IndividualConsentRequestTemplateDetails
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsConsentTemplatesTemplateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsConsentTemplatesTemplateIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndividualConsentRequestTemplateDetails**](IndividualConsentRequestTemplateDetails.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1IndividualsConsentsConsentIdDocumentsDocumentIdDownloadGet

> *os.File V1IndividualsConsentsConsentIdDocumentsDocumentIdDownloadGet(ctx, consentId, documentId).Execute()

Download Individual consented document by document id.

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsConsentsConsentIdDocumentsDocumentIdDownloadGet(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsConsentsConsentIdDocumentsDocumentIdDownloadGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsConsentsConsentIdDocumentsDocumentIdDownloadGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsConsentsConsentIdDocumentsDocumentIdDownloadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsConsentsConsentIdDocumentsDocumentIdDownloadGetRequest struct via the builder pattern


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


## V1IndividualsConsentsConsentIdDocumentsDocumentIdGet

> ConsentedDocument V1IndividualsConsentsConsentIdDocumentsDocumentIdGet(ctx, consentId, documentId).Execute()

Get Individual consented document by document id.

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsConsentsConsentIdDocumentsDocumentIdGet(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsConsentsConsentIdDocumentsDocumentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsConsentsConsentIdDocumentsDocumentIdGet`: ConsentedDocument
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsConsentsConsentIdDocumentsDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsConsentsConsentIdDocumentsDocumentIdGetRequest struct via the builder pattern


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


## V1IndividualsConsentsConsentIdDocumentsGet

> []ConsentedDocument V1IndividualsConsentsConsentIdDocumentsGet(ctx, consentId).Execute()

Get Individual consented document by consent id.

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsConsentsConsentIdDocumentsGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsConsentsConsentIdDocumentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsConsentsConsentIdDocumentsGet`: []ConsentedDocument
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsConsentsConsentIdDocumentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsConsentsConsentIdDocumentsGetRequest struct via the builder pattern


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


## V1IndividualsConsentsConsentIdFinancialAccountsAccountIdGet

> FinancialAccount V1IndividualsConsentsConsentIdFinancialAccountsAccountIdGet(ctx, consentId, accountId).Execute()

Get individual consented financial account details.

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsConsentsConsentIdFinancialAccountsAccountIdGet(context.Background(), consentId, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsConsentsConsentIdFinancialAccountsAccountIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsConsentsConsentIdFinancialAccountsAccountIdGet`: FinancialAccount
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsConsentsConsentIdFinancialAccountsAccountIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsConsentsConsentIdFinancialAccountsAccountIdGetRequest struct via the builder pattern


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


## V1IndividualsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet

> PaginatedListOfFinancialAccountTransactions V1IndividualsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet(ctx, consentId, accountId).Filters(filters).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()

Get individual consented financial account transactions.

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
    filters := "filters_example" // string |  (optional)
    fromDateTime := "fromDateTime_example" // string |  (optional)
    toDateTime := "toDateTime_example" // string |  (optional)
    pageNo := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualsApi.V1IndividualsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet(context.Background(), consentId, accountId).Filters(filters).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet`: PaginatedListOfFinancialAccountTransactions
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsConsentsConsentIdFinancialAccountsAccountIdTransactionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsConsentsConsentIdFinancialAccountsAccountIdTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filters** | **string** |  | 
 **fromDateTime** | **string** |  | 
 **toDateTime** | **string** |  | 
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


## V1IndividualsConsentsConsentIdFinancialAccountsGet

> []FinancialAccount V1IndividualsConsentsConsentIdFinancialAccountsGet(ctx, consentId).Execute()

Get all individual consented financial accounts.

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsConsentsConsentIdFinancialAccountsGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsConsentsConsentIdFinancialAccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsConsentsConsentIdFinancialAccountsGet`: []FinancialAccount
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsConsentsConsentIdFinancialAccountsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsConsentsConsentIdFinancialAccountsGetRequest struct via the builder pattern


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


## V1IndividualsConsentsConsentIdGet

> ConsentDetails V1IndividualsConsentsConsentIdGet(ctx, consentId).Execute()

Get Individuals consent details by consent id.

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsConsentsConsentIdGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsConsentsConsentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsConsentsConsentIdGet`: ConsentDetails
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsConsentsConsentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsConsentsConsentIdGetRequest struct via the builder pattern


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


## V1IndividualsConsentsConsentIdHealthFhirBundleGet

> []HealthRecord V1IndividualsConsentsConsentIdHealthFhirBundleGet(ctx, consentId).Execute()

Get Individual consented Health Records by consent id.

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsConsentsConsentIdHealthFhirBundleGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsConsentsConsentIdHealthFhirBundleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsConsentsConsentIdHealthFhirBundleGet`: []HealthRecord
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsConsentsConsentIdHealthFhirBundleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsConsentsConsentIdHealthFhirBundleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HealthRecord**](HealthRecord.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1IndividualsConsentsGet

> PaginatedListOfConsents V1IndividualsConsentsGet(ctx).Status(status).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()

Get the paginated list of Individual consents.

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsConsentsGet(context.Background()).Status(status).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsConsentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsConsentsGet`: PaginatedListOfConsents
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsConsentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsConsentsGetRequest struct via the builder pattern


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


## V1IndividualsDocumentsIssueIssueRequestIdUploadPost

> UploadDocumentResponse V1IndividualsDocumentsIssueIssueRequestIdUploadPost(ctx, issueRequestId).File(file).Execute()

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
    issueRequestId := "issueRequestId_example" // string | 
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualsApi.V1IndividualsDocumentsIssueIssueRequestIdUploadPost(context.Background(), issueRequestId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsDocumentsIssueIssueRequestIdUploadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsDocumentsIssueIssueRequestIdUploadPost`: UploadDocumentResponse
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsDocumentsIssueIssueRequestIdUploadPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsDocumentsIssueIssueRequestIdUploadPostRequest struct via the builder pattern


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


## V1IndividualsDocumentsIssuePost

> DocumentIssueRequestDetails V1IndividualsDocumentsIssuePost(ctx).DocumentIssueRequest(documentIssueRequest).Execute()

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
    documentIssueRequest := *openapiclient.NewDocumentIssueRequest("DocumentTypeId_example", "Identifier_example", "Description_example", *openapiclient.NewDocumentIssueRequestReceiver("CountryIso2Code_example", []openapiclient.KeyValuePair{*openapiclient.NewKeyValuePair("Key_example", "Value_example")}, openapiclient.IdentificationStrategy("MatchAtLeastOneIdentifier")), time.Now(), time.Now()) // DocumentIssueRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualsApi.V1IndividualsDocumentsIssuePost(context.Background()).DocumentIssueRequest(documentIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsDocumentsIssuePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsDocumentsIssuePost`: DocumentIssueRequestDetails
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsDocumentsIssuePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsDocumentsIssuePostRequest struct via the builder pattern


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


## V1IndividualsDocumentsIssuedDocumentIdGet

> IssuedDocument V1IndividualsDocumentsIssuedDocumentIdGet(ctx, documentId).Execute()

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsDocumentsIssuedDocumentIdGet(context.Background(), documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsDocumentsIssuedDocumentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsDocumentsIssuedDocumentIdGet`: IssuedDocument
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsDocumentsIssuedDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsDocumentsIssuedDocumentIdGetRequest struct via the builder pattern


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


## V1IndividualsDocumentsIssuedGet

> PaginatedListOfIssuedDocuments V1IndividualsDocumentsIssuedGet(ctx).DocumentTypeId(documentTypeId).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsDocumentsIssuedGet(context.Background()).DocumentTypeId(documentTypeId).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsDocumentsIssuedGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsDocumentsIssuedGet`: PaginatedListOfIssuedDocuments
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsDocumentsIssuedGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsDocumentsIssuedGetRequest struct via the builder pattern


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


## V1IndividualsDocumentsTypesGet

> PaginatedListOfDocumentTypes V1IndividualsDocumentsTypesGet(ctx).PageNo(pageNo).PageSize(pageSize).Execute()

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
    resp, r, err := apiClient.IndividualsApi.V1IndividualsDocumentsTypesGet(context.Background()).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualsApi.V1IndividualsDocumentsTypesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1IndividualsDocumentsTypesGet`: PaginatedListOfDocumentTypes
    fmt.Fprintf(os.Stdout, "Response from `IndividualsApi.V1IndividualsDocumentsTypesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1IndividualsDocumentsTypesGetRequest struct via the builder pattern


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

