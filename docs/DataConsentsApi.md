# \DataConsentsApi

All URIs are relative to *https://api.mydatamyconsent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadConsentedDocumentAnalysis**](DataConsentsApi.md#DownloadConsentedDocumentAnalysis) | **Get** /v1/consents/{consentId}/documents/{documentId}/analysis | Get analysis of a consented document.
[**DownloadIndividualConsentedDocumentById**](DataConsentsApi.md#DownloadIndividualConsentedDocumentById) | **Get** /v1/consents/individuals/{consentId}/documents/{documentId}/download | Download individual consented document by document id.
[**DownloadOrganizationConsentedDocumentById**](DataConsentsApi.md#DownloadOrganizationConsentedDocumentById) | **Get** /v1/consents/organizations/{consentId}/documents/{documentId}/download | Download organization consent document based on document id.
[**GetAllConsentedFinancialAccounts**](DataConsentsApi.md#GetAllConsentedFinancialAccounts) | **Get** /v1/consents/individuals/{consentId}/financial-accounts | Get all individual consented financial accounts.
[**GetConsentFinancialAccounts**](DataConsentsApi.md#GetConsentFinancialAccounts) | **Get** /v1/consents/organizations/{consentId}/financial-accounts | Get all organizational consented financial accounts.
[**GetConsentedAccountById**](DataConsentsApi.md#GetConsentedAccountById) | **Get** /v1/consents/individuals/{consentId}/financial-accounts/{accountId} | Get individual consented financial account details based on account id.
[**GetConsentedDocumentById**](DataConsentsApi.md#GetConsentedDocumentById) | **Get** /v1/consents/individuals/{consentId}/documents/{documentId} | Get individual consented document by document id.
[**GetConsentedFinancialAccount**](DataConsentsApi.md#GetConsentedFinancialAccount) | **Get** /v1/consents/organizations/{consentId}/financial-accounts/{accountId} | Get organization consented financial account details based on account id.
[**GetConsentedFinancialAccountInsights**](DataConsentsApi.md#GetConsentedFinancialAccountInsights) | **Get** /v1/consents/{consentId}/financial-accounts/{accountId}/insights | Get consented financial account insights.
[**GetConsentedFinancialAccountTransactions**](DataConsentsApi.md#GetConsentedFinancialAccountTransactions) | **Get** /v1/consents/individuals/{consentId}/financial-accounts/{accountId}/transactions | Get individual consented financial account transactions of an individual based on accountId.
[**GetConsents**](DataConsentsApi.md#GetConsents) | **Get** /v1/consents/individuals | Get the paginated list of individual data consents.
[**GetIndividualConsentedDocuments**](DataConsentsApi.md#GetIndividualConsentedDocuments) | **Get** /v1/consents/individuals/{consentId}/documents | Get individual consented documents by consent id.
[**GetIndividualDataConsentById**](DataConsentsApi.md#GetIndividualDataConsentById) | **Get** /v1/consents/individuals/{consentId} | Get individuals data consent details by consent id.
[**GetOrgConsentedAccountTransactions**](DataConsentsApi.md#GetOrgConsentedAccountTransactions) | **Get** /v1/consents/organizations/{consentId}/financial-accounts/{accountId}/transactions | Get organization consented financial account transactions of an individual based on accountId.
[**GetOrganizationConsentedDocumentById**](DataConsentsApi.md#GetOrganizationConsentedDocumentById) | **Get** /v1/consents/organizations/{consentId}/documents/{documentId} | Get organization consent document based on document id.
[**GetOrganizationConsentedDocuments**](DataConsentsApi.md#GetOrganizationConsentedDocuments) | **Get** /v1/consents/organizations/{consentId}/documents | Get organization consented documents by consent id.
[**GetOrganizationDataConsentById**](DataConsentsApi.md#GetOrganizationDataConsentById) | **Get** /v1/consents/organizations/{consentId} | Get organizations data consent details by consent id.
[**GetOrganizationDataConsents**](DataConsentsApi.md#GetOrganizationDataConsents) | **Get** /v1/consents/organizations | Get the paginated list of organization data consents.



## DownloadConsentedDocumentAnalysis

> map[string]interface{} DownloadConsentedDocumentAnalysis(ctx, consentId, documentId).Execute()

Get analysis of a consented document.

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
    consentId := "consentId_example" // string | Data consent id.
    documentId := "documentId_example" // string | Consented document Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.DownloadConsentedDocumentAnalysis(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.DownloadConsentedDocumentAnalysis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadConsentedDocumentAnalysis`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.DownloadConsentedDocumentAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Data consent id. | 
**documentId** | **string** | Consented document Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadConsentedDocumentAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadIndividualConsentedDocumentById

> map[string]interface{} DownloadIndividualConsentedDocumentById(ctx, consentId, documentId).Execute()

Download individual consented document by document id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Individual data consent id.
    documentId := "documentId_example" // string | Consented document id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.DownloadIndividualConsentedDocumentById(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.DownloadIndividualConsentedDocumentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadIndividualConsentedDocumentById`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.DownloadIndividualConsentedDocumentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Individual data consent id. | 
**documentId** | **string** | Consented document id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadIndividualConsentedDocumentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadOrganizationConsentedDocumentById

> map[string]interface{} DownloadOrganizationConsentedDocumentById(ctx, consentId, documentId).Execute()

Download organization consent document based on document id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization data consent id.
    documentId := "documentId_example" // string | Organization consented document Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.DownloadOrganizationConsentedDocumentById(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.DownloadOrganizationConsentedDocumentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadOrganizationConsentedDocumentById`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.DownloadOrganizationConsentedDocumentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Organization data consent id. | 
**documentId** | **string** | Organization consented document Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadOrganizationConsentedDocumentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

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

> map[string]interface{} GetConsentedDocumentById(ctx, consentId, documentId).Execute()

Get individual consented document by document id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Individual data consent id.
    documentId := "documentId_example" // string | Consented document id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetConsentedDocumentById(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetConsentedDocumentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsentedDocumentById`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetConsentedDocumentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Individual data consent id. | 
**documentId** | **string** | Consented document id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentedDocumentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

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


## GetConsentedFinancialAccountInsights

> GetConsentedFinancialAccountInsights(ctx, consentId, accountId).Execute()

Get consented financial account insights.

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
    resp, r, err := apiClient.DataConsentsApi.GetConsentedFinancialAccountInsights(context.Background(), consentId, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetConsentedFinancialAccountInsights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentedFinancialAccountInsightsRequest struct via the builder pattern


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


## GetConsents

> map[string]interface{} GetConsents(ctx).Status(status).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()

Get the paginated list of individual data consents.



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
    fromDateTime := time.Now() // time.Time | From datetime in UTC timezone. (optional)
    toDateTime := time.Now() // time.Time | To datetime in UTC timezone. (optional)
    pageNo := int32(56) // int32 | Page number. (optional) (default to 1)
    pageSize := int32(56) // int32 | Number of items to return. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetConsents(context.Background()).Status(status).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetConsents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsents`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetConsents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**DataConsentStatus**](DataConsentStatus.md) | Data consent status MyDataMyConsent.Domain.Entities.ConsentAggregate.Enums.DataConsentStatus. | 
 **fromDateTime** | **time.Time** | From datetime in UTC timezone. | 
 **toDateTime** | **time.Time** | To datetime in UTC timezone. | 
 **pageNo** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Number of items to return. | [default to 25]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndividualConsentedDocuments

> map[string]interface{} GetIndividualConsentedDocuments(ctx, consentId).Execute()

Get individual consented documents by consent id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Individual data consent id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetIndividualConsentedDocuments(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetIndividualConsentedDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndividualConsentedDocuments`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetIndividualConsentedDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Individual data consent id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndividualConsentedDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndividualDataConsentById

> map[string]interface{} GetIndividualDataConsentById(ctx, consentId).Execute()

Get individuals data consent details by consent id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Individual data consent id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetIndividualDataConsentById(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetIndividualDataConsentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndividualDataConsentById`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetIndividualDataConsentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Individual data consent id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndividualDataConsentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

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


## GetOrganizationConsentedDocumentById

> map[string]interface{} GetOrganizationConsentedDocumentById(ctx, consentId, documentId).Execute()

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization data consent id.
    documentId := "documentId_example" // string | Organization consented document Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetOrganizationConsentedDocumentById(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetOrganizationConsentedDocumentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConsentedDocumentById`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetOrganizationConsentedDocumentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Organization data consent id. | 
**documentId** | **string** | Organization consented document Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConsentedDocumentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConsentedDocuments

> map[string]interface{} GetOrganizationConsentedDocuments(ctx, consentId).Execute()

Get organization consented documents by consent id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization data consent id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetOrganizationConsentedDocuments(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetOrganizationConsentedDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConsentedDocuments`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetOrganizationConsentedDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Organization data consent id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConsentedDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDataConsentById

> map[string]interface{} GetOrganizationDataConsentById(ctx, consentId).Execute()

Get organizations data consent details by consent id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization data consent id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetOrganizationDataConsentById(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetOrganizationDataConsentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDataConsentById`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetOrganizationDataConsentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | Organization data consent id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDataConsentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDataConsents

> map[string]interface{} GetOrganizationDataConsents(ctx).Status(status).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()

Get the paginated list of organization data consents.

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
    fromDateTime := time.Now() // time.Time | From datetime in UTC timezone. (optional)
    toDateTime := time.Now() // time.Time | To datetime in UTC timezone. (optional)
    pageNo := int32(56) // int32 | Page number. (optional) (default to 1)
    pageSize := int32(56) // int32 | Number of items to return. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.GetOrganizationDataConsents(context.Background()).Status(status).FromDateTime(fromDateTime).ToDateTime(toDateTime).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.GetOrganizationDataConsents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDataConsents`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.GetOrganizationDataConsents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDataConsentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**DataConsentStatus**](DataConsentStatus.md) | Data consent status MyDataMyConsent.Domain.Entities.ConsentAggregate.Enums.DataConsentStatus. | 
 **fromDateTime** | **time.Time** | From datetime in UTC timezone. | 
 **toDateTime** | **time.Time** | To datetime in UTC timezone. | 
 **pageNo** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Number of items to return. | [default to 25]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

