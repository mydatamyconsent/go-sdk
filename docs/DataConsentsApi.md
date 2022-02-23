# \DataConsentsApi

All URIs are relative to *https://api.mydatamyconsent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadConsentedDocumentById**](DataConsentsApi.md#DownloadConsentedDocumentById) | **Get** /v1/consents/individuals/{consentId}/documents/{documentId}/download | Download a individuals consented document.
[**DownloadOrgConsentedDocumentById**](DataConsentsApi.md#DownloadOrgConsentedDocumentById) | **Get** /v1/consents/organizations/{consentId}/documents/{documentId}/download | Download a organizations consented document.
[**GetAllConsentedDocuments**](DataConsentsApi.md#GetAllConsentedDocuments) | **Get** /v1/consents/individuals/{consentId}/documents | Get the individual documents based on ConsentId.
[**GetAllConsentedFinancialAccounts**](DataConsentsApi.md#GetAllConsentedFinancialAccounts) | **Get** /v1/consents/individuals/{consentId}/accounts | Get all individual consented financial accounts.
[**GetAllOrganizationConsentedDocuments**](DataConsentsApi.md#GetAllOrganizationConsentedDocuments) | **Get** /v1/consents/organizations/{consentId}/documents | Get the organization documents based on ConsentId.
[**GetConsentDetailsById**](DataConsentsApi.md#GetConsentDetailsById) | **Get** /v1/consents/individuals/{consentId} | Get all individuals consent details by consent id.
[**GetConsentFinancialAccounts**](DataConsentsApi.md#GetConsentFinancialAccounts) | **Get** /v1/consents/organizations/{consentId}/accounts | Get all organizational consented financial accounts.
[**GetConsentedAccountById**](DataConsentsApi.md#GetConsentedAccountById) | **Get** /v1/consents/individuals/{consentId}/accounts/{accountId} | Get individual consented financial account details based on account id.
[**GetConsentedDocumentById**](DataConsentsApi.md#GetConsentedDocumentById) | **Get** /v1/consents/individuals/{consentId}/documents/{documentId} | Get individuals consent document based on document id.
[**GetConsentedFinancialAccount**](DataConsentsApi.md#GetConsentedFinancialAccount) | **Get** /v1/consents/organizations/{consentId}/accounts/{accountId} | Get organization consented financial account details based on account id.
[**GetConsentedFinancialAccountTransactions**](DataConsentsApi.md#GetConsentedFinancialAccountTransactions) | **Get** /v1/consents/individuals/{consentId}/accounts/{accountId}/transactions | Get individual consented financial account transactions of an individual based on accountId.
[**GetConsentsForOrganizations**](DataConsentsApi.md#GetConsentsForOrganizations) | **Get** /v1/consents/organizations | Get the list of data consents sent for organizations.
[**GetConsentsSentToIndividuals**](DataConsentsApi.md#GetConsentsSentToIndividuals) | **Get** /v1/consents/individuals | Get the list of Consents Sent to Individuals.
[**GetOrgConsentedAccountTransactions**](DataConsentsApi.md#GetOrgConsentedAccountTransactions) | **Get** /v1/consents/organizations/{consentId}/accounts/{accountId}/transactions | Get organization consented financial account transactions of an individual based on accountId.
[**GetOrganizationConsentDetailsById**](DataConsentsApi.md#GetOrganizationConsentDetailsById) | **Get** /v1/consents/organizations/{consentId} | Get all organization consent details by consent id.
[**GetOrganizationConsentedDocumentById**](DataConsentsApi.md#GetOrganizationConsentedDocumentById) | **Get** /v1/consents/organizations/{consentId}/documents/{documentId} | Get organization consent document based on document id.



## DownloadConsentedDocumentById

> UserDocumentDownloadDto DownloadConsentedDocumentById(ctx, consentId, documentId).Execute()

Download a individuals consented document.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Consent id.
    documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Document id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.DownloadConsentedDocumentById(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.DownloadConsentedDocumentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadConsentedDocumentById`: UserDocumentDownloadDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.DownloadConsentedDocumentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Consent id. | 
**documentId** | **string** | Document id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadConsentedDocumentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserDocumentDownloadDto**](UserDocumentDownloadDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadOrgConsentedDocumentById

> OrganizationDocumentDownloadDto DownloadOrgConsentedDocumentById(ctx, consentId, documentId).Execute()

Download a organizations consented document.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Consent id.
    documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Document id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.DownloadOrgConsentedDocumentById(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.DownloadOrgConsentedDocumentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadOrgConsentedDocumentById`: OrganizationDocumentDownloadDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.DownloadOrgConsentedDocumentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Consent id. | 
**documentId** | **string** | Document id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadOrgConsentedDocumentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationDocumentDownloadDto**](OrganizationDocumentDownloadDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllConsentedDocuments

> DataConsentDocumentsDto GetAllConsentedDocuments(ctx, consentId).Execute()

Get the individual documents based on ConsentId.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Consent id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetAllConsentedDocuments(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetAllConsentedDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllConsentedDocuments`: DataConsentDocumentsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetAllConsentedDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Consent id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConsentedDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataConsentDocumentsDto**](DataConsentDocumentsDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllConsentedFinancialAccounts

> DataConsentFinancialsDto GetAllConsentedFinancialAccounts(ctx, consentId).Execute()

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Consent id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetAllConsentedFinancialAccounts(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetAllConsentedFinancialAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllConsentedFinancialAccounts`: DataConsentFinancialsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetAllConsentedFinancialAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Consent id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConsentedFinancialAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataConsentFinancialsDto**](DataConsentFinancialsDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllOrganizationConsentedDocuments

> DataConsentDocumentsDto GetAllOrganizationConsentedDocuments(ctx, consentId).Execute()

Get the organization documents based on ConsentId.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Consent id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetAllOrganizationConsentedDocuments(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetAllOrganizationConsentedDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllOrganizationConsentedDocuments`: DataConsentDocumentsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetAllOrganizationConsentedDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Consent id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllOrganizationConsentedDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataConsentDocumentsDto**](DataConsentDocumentsDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsentDetailsById

> DataConsentDetailsDto GetConsentDetailsById(ctx, consentId).Execute()

Get all individuals consent details by consent id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Consent id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetConsentDetailsById(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetConsentDetailsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsentDetailsById`: DataConsentDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetConsentDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Consent id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentDetailsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataConsentDetailsDto**](DataConsentDetailsDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsentFinancialAccounts

> DataConsentFinancialsDto GetConsentFinancialAccounts(ctx, consentId).Execute()

Get all organizational consented financial accounts.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Consent id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetConsentFinancialAccounts(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetConsentFinancialAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsentFinancialAccounts`: DataConsentFinancialsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetConsentFinancialAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Consent id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentFinancialAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataConsentFinancialsDto**](DataConsentFinancialsDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsentedAccountById

> FinancialAccount GetConsentedAccountById(ctx, consentId, accountId).Execute()

Get individual consented financial account details based on account id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Consent id.
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Account id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetConsentedAccountById(context.Background(), consentId, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetConsentedAccountById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsentedAccountById`: FinancialAccount
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetConsentedAccountById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Consent id. | 
**accountId** | **string** | Account id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentedAccountByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FinancialAccount**](FinancialAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsentedDocumentById

> UserDocumentDetailsDto GetConsentedDocumentById(ctx, consentId, documentId).Execute()

Get individuals consent document based on document id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Consent id.
    documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Document Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetConsentedDocumentById(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetConsentedDocumentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsentedDocumentById`: UserDocumentDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetConsentedDocumentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Consent id. | 
**documentId** | **string** | Document Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentedDocumentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserDocumentDetailsDto**](UserDocumentDetailsDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsentedFinancialAccount

> OrganizationFinancialAccountDto GetConsentedFinancialAccount(ctx, consentId, accountId).Execute()

Get organization consented financial account details based on account id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Consent id.
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Account id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetConsentedFinancialAccount(context.Background(), consentId, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetConsentedFinancialAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsentedFinancialAccount`: OrganizationFinancialAccountDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetConsentedFinancialAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Consent id. | 
**accountId** | **string** | Account id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentedFinancialAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationFinancialAccountDto**](OrganizationFinancialAccountDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsentedFinancialAccountTransactions

> UserAccountFinancialTransactionsDtoPaginatedList GetConsentedFinancialAccountTransactions(ctx, consentId, accountId).Filters(filters).FromDateTimeUtc(fromDateTimeUtc).ToDateTimeUtc(toDateTimeUtc).PageNo(pageNo).PageSize(pageSize).Execute()

Get individual consented financial account transactions of an individual based on accountId.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Consent id.
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Account id.
    filters := "filters_example" // string | Filters. (optional)
    fromDateTimeUtc := time.Now() // time.Time | From date time in utc timezone. (optional)
    toDateTimeUtc := time.Now() // time.Time | Til date time in utc timezone. (optional)
    pageNo := int32(56) // int32 | Page number. (optional) (default to 10)
    pageSize := int32(56) // int32 | Number of items to return. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetConsentedFinancialAccountTransactions(context.Background(), consentId, accountId).Filters(filters).FromDateTimeUtc(fromDateTimeUtc).ToDateTimeUtc(toDateTimeUtc).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetConsentedFinancialAccountTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsentedFinancialAccountTransactions`: UserAccountFinancialTransactionsDtoPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetConsentedFinancialAccountTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Consent id. | 
**accountId** | **string** | Account id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentedFinancialAccountTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filters** | **string** | Filters. | 
 **fromDateTimeUtc** | **time.Time** | From date time in utc timezone. | 
 **toDateTimeUtc** | **time.Time** | Til date time in utc timezone. | 
 **pageNo** | **int32** | Page number. | [default to 10]
 **pageSize** | **int32** | Number of items to return. | [default to 25]

### Return type

[**UserAccountFinancialTransactionsDtoPaginatedList**](UserAccountFinancialTransactionsDtoPaginatedList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsentsForOrganizations

> OrganizationDataConsentInfoDtoPaginatedList GetConsentsForOrganizations(ctx).Status(status).From(from).To(to).PageNo(pageNo).PageSize(pageSize).Execute()

Get the list of data consents sent for organizations.

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
    status := openapiclient.DataConsentStatus("Pending") // DataConsentStatus | Data consent status MyDataMyConsent.Domain.Entities.ConsentAggregate.Enums.DataConsentStatus. (optional)
    from := time.Now() // time.Time | From date time in utc timezone. (optional)
    to := time.Now() // time.Time | Til date time in utc timezone. (optional)
    pageNo := int32(56) // int32 | Page number. (optional) (default to 1)
    pageSize := int32(56) // int32 | Number of items to return. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetConsentsForOrganizations(context.Background()).Status(status).From(from).To(to).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetConsentsForOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsentsForOrganizations`: OrganizationDataConsentInfoDtoPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetConsentsForOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentsForOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**DataConsentStatus**](DataConsentStatus.md) | Data consent status MyDataMyConsent.Domain.Entities.ConsentAggregate.Enums.DataConsentStatus. | 
 **from** | **time.Time** | From date time in utc timezone. | 
 **to** | **time.Time** | Til date time in utc timezone. | 
 **pageNo** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Number of items to return. | [default to 25]

### Return type

[**OrganizationDataConsentInfoDtoPaginatedList**](OrganizationDataConsentInfoDtoPaginatedList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsentsSentToIndividuals

> UserDataConsentInfoDtoPaginatedList GetConsentsSentToIndividuals(ctx).Status(status).From(from).To(to).PageNo(pageNo).PageSize(pageSize).Execute()

Get the list of Consents Sent to Individuals.

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
    status := openapiclient.DataConsentStatus("Pending") // DataConsentStatus | Data consent status MyDataMyConsent.Domain.Entities.ConsentAggregate.Enums.DataConsentStatus. (optional)
    from := time.Now() // time.Time | From date time in utc timezone. (optional)
    to := time.Now() // time.Time | Til date time in utc timezone. (optional)
    pageNo := int32(56) // int32 | Page number. (optional) (default to 1)
    pageSize := int32(56) // int32 | Number of items to return. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetConsentsSentToIndividuals(context.Background()).Status(status).From(from).To(to).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetConsentsSentToIndividuals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsentsSentToIndividuals`: UserDataConsentInfoDtoPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetConsentsSentToIndividuals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentsSentToIndividualsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**DataConsentStatus**](DataConsentStatus.md) | Data consent status MyDataMyConsent.Domain.Entities.ConsentAggregate.Enums.DataConsentStatus. | 
 **from** | **time.Time** | From date time in utc timezone. | 
 **to** | **time.Time** | Til date time in utc timezone. | 
 **pageNo** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Number of items to return. | [default to 25]

### Return type

[**UserDataConsentInfoDtoPaginatedList**](UserDataConsentInfoDtoPaginatedList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgConsentedAccountTransactions

> OrganizationFinancialTransactionsDtoPaginatedList GetOrgConsentedAccountTransactions(ctx, consentId, accountId).Filters(filters).FromDateTimeUtc(fromDateTimeUtc).ToDateTimeUtc(toDateTimeUtc).PageNo(pageNo).PageSize(pageSize).Execute()

Get organization consented financial account transactions of an individual based on accountId.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Consent id.
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Account id.
    filters := "filters_example" // string | Filters. (optional)
    fromDateTimeUtc := time.Now() // time.Time | From date time in utc timezone. (optional)
    toDateTimeUtc := time.Now() // time.Time | Til date time in utc timezone. (optional)
    pageNo := int32(56) // int32 | Page number. (optional) (default to 1)
    pageSize := int32(56) // int32 | Number of items to return. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetOrgConsentedAccountTransactions(context.Background(), consentId, accountId).Filters(filters).FromDateTimeUtc(fromDateTimeUtc).ToDateTimeUtc(toDateTimeUtc).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetOrgConsentedAccountTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgConsentedAccountTransactions`: OrganizationFinancialTransactionsDtoPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetOrgConsentedAccountTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Consent id. | 
**accountId** | **string** | Account id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgConsentedAccountTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filters** | **string** | Filters. | 
 **fromDateTimeUtc** | **time.Time** | From date time in utc timezone. | 
 **toDateTimeUtc** | **time.Time** | Til date time in utc timezone. | 
 **pageNo** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Number of items to return. | [default to 25]

### Return type

[**OrganizationFinancialTransactionsDtoPaginatedList**](OrganizationFinancialTransactionsDtoPaginatedList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConsentDetailsById

> DataConsentDetailsDto GetOrganizationConsentDetailsById(ctx, consentId).Execute()

Get all organization consent details by consent id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Consent id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetOrganizationConsentDetailsById(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetOrganizationConsentDetailsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConsentDetailsById`: DataConsentDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetOrganizationConsentDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Consent id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConsentDetailsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataConsentDetailsDto**](DataConsentDetailsDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConsentedDocumentById

> OrganizationDocumentDetailsDto GetOrganizationConsentedDocumentById(ctx, consentId, documentId).Execute()

Get organization consent document based on document id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Consent id.
    documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Document Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetOrganizationConsentedDocumentById(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetOrganizationConsentedDocumentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConsentedDocumentById`: OrganizationDocumentDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetOrganizationConsentedDocumentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Consent id. | 
**documentId** | **string** | Document Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConsentedDocumentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationDocumentDetailsDto**](OrganizationDocumentDetailsDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

