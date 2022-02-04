# \DataConsentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ConsentsIndividualsConsentIdAccountsAccountIdGet**](DataConsentsApi.md#V1ConsentsIndividualsConsentIdAccountsAccountIdGet) | **Get** /v1/consents/individuals/{consentId}/accounts/{accountId} | Get individual consented financial account details based on account id.
[**V1ConsentsIndividualsConsentIdAccountsAccountIdTransactionsGet**](DataConsentsApi.md#V1ConsentsIndividualsConsentIdAccountsAccountIdTransactionsGet) | **Get** /v1/consents/individuals/{consentId}/accounts/{accountId}/transactions | Get consented financial account transactions of an individual based on accountId.
[**V1ConsentsIndividualsConsentIdAccountsGet**](DataConsentsApi.md#V1ConsentsIndividualsConsentIdAccountsGet) | **Get** /v1/consents/individuals/{consentId}/accounts | Get all individual financial accounts in a consent.
[**V1ConsentsIndividualsConsentIdDocumentsDocumentIdDownloadGet**](DataConsentsApi.md#V1ConsentsIndividualsConsentIdDocumentsDocumentIdDownloadGet) | **Get** /v1/consents/individuals/{consentId}/documents/{documentId}/download | Download a individuals consented document.
[**V1ConsentsIndividualsConsentIdDocumentsDocumentIdGet**](DataConsentsApi.md#V1ConsentsIndividualsConsentIdDocumentsDocumentIdGet) | **Get** /v1/consents/individuals/{consentId}/documents/{documentId} | Get individuals consent document based on document id.
[**V1ConsentsIndividualsConsentIdDocumentsGet**](DataConsentsApi.md#V1ConsentsIndividualsConsentIdDocumentsGet) | **Get** /v1/consents/individuals/{consentId}/documents | Get the individual documents based on ConsentId.
[**V1ConsentsIndividualsConsentIdGet**](DataConsentsApi.md#V1ConsentsIndividualsConsentIdGet) | **Get** /v1/consents/individuals/{consentId} | Get individuals consent details by consent id.
[**V1ConsentsIndividualsGet**](DataConsentsApi.md#V1ConsentsIndividualsGet) | **Get** /v1/consents/individuals | Get the list of Consents Sent to Individuals.
[**V1ConsentsOrganizationsConsentIdAccountsAccountIdGet**](DataConsentsApi.md#V1ConsentsOrganizationsConsentIdAccountsAccountIdGet) | **Get** /v1/consents/organizations/{consentId}/accounts/{accountId} | Get orgnization consented financial account details based on account id.
[**V1ConsentsOrganizationsConsentIdAccountsAccountIdTransactionsGet**](DataConsentsApi.md#V1ConsentsOrganizationsConsentIdAccountsAccountIdTransactionsGet) | **Get** /v1/consents/organizations/{consentId}/accounts/{accountId}/transactions | Get consented financial account transactions of an organization based on accountId.
[**V1ConsentsOrganizationsConsentIdAccountsGet**](DataConsentsApi.md#V1ConsentsOrganizationsConsentIdAccountsGet) | **Get** /v1/consents/organizations/{consentId}/accounts | Get all organizational financial accounts in a consent.
[**V1ConsentsOrganizationsConsentIdDocumentsDocumentIdDownloadGet**](DataConsentsApi.md#V1ConsentsOrganizationsConsentIdDocumentsDocumentIdDownloadGet) | **Get** /v1/consents/organizations/{consentId}/documents/{documentId}/download | Download organizations consented document.
[**V1ConsentsOrganizationsConsentIdDocumentsDocumentIdGet**](DataConsentsApi.md#V1ConsentsOrganizationsConsentIdDocumentsDocumentIdGet) | **Get** /v1/consents/organizations/{consentId}/documents/{documentId} | Get organizations consent document based on document id.
[**V1ConsentsOrganizationsConsentIdDocumentsGet**](DataConsentsApi.md#V1ConsentsOrganizationsConsentIdDocumentsGet) | **Get** /v1/consents/organizations/{consentId}/documents | Get the organizations documents based on ConsentId.
[**V1ConsentsOrganizationsConsentIdGet**](DataConsentsApi.md#V1ConsentsOrganizationsConsentIdGet) | **Get** /v1/consents/organizations/{consentId} | Get organizations consent details by consent id.
[**V1ConsentsOrganizationsGet**](DataConsentsApi.md#V1ConsentsOrganizationsGet) | **Get** /v1/consents/organizations | Get the list of data consents sent for organizations.



## V1ConsentsIndividualsConsentIdAccountsAccountIdGet

> FinancialAccount V1ConsentsIndividualsConsentIdAccountsAccountIdGet(ctx, consentId, accountId).Execute()

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsIndividualsConsentIdAccountsAccountIdGet(context.Background(), consentId, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsIndividualsConsentIdAccountsAccountIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsIndividualsConsentIdAccountsAccountIdGet`: FinancialAccount
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsIndividualsConsentIdAccountsAccountIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsIndividualsConsentIdAccountsAccountIdGetRequest struct via the builder pattern


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


## V1ConsentsIndividualsConsentIdAccountsAccountIdTransactionsGet

> UserAccountFinancialTransactionsDtoPaginatedList V1ConsentsIndividualsConsentIdAccountsAccountIdTransactionsGet(ctx, consentId, accountId).Filters(filters).PageNo(pageNo).PageSize(pageSize).FromDate(fromDate).ToDate(toDate).Execute()

Get consented financial account transactions of an individual based on accountId.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    filters := "filters_example" // string |  (optional)
    pageNo := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    fromDate := time.Now() // time.Time |  (optional)
    toDate := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsIndividualsConsentIdAccountsAccountIdTransactionsGet(context.Background(), consentId, accountId).Filters(filters).PageNo(pageNo).PageSize(pageSize).FromDate(fromDate).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsIndividualsConsentIdAccountsAccountIdTransactionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsIndividualsConsentIdAccountsAccountIdTransactionsGet`: UserAccountFinancialTransactionsDtoPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsIndividualsConsentIdAccountsAccountIdTransactionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsIndividualsConsentIdAccountsAccountIdTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filters** | **string** |  | 
 **pageNo** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **fromDate** | **time.Time** |  | 
 **toDate** | **time.Time** |  | 

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


## V1ConsentsIndividualsConsentIdAccountsGet

> DataConsentFinancialsDto V1ConsentsIndividualsConsentIdAccountsGet(ctx, consentId).Execute()

Get all individual financial accounts in a consent.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsIndividualsConsentIdAccountsGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsIndividualsConsentIdAccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsIndividualsConsentIdAccountsGet`: DataConsentFinancialsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsIndividualsConsentIdAccountsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsIndividualsConsentIdAccountsGetRequest struct via the builder pattern


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


## V1ConsentsIndividualsConsentIdDocumentsDocumentIdDownloadGet

> UserDocumentDownloadDto V1ConsentsIndividualsConsentIdDocumentsDocumentIdDownloadGet(ctx, consentId, documentId).Execute()

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | consentId.
    documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | documentId.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsIndividualsConsentIdDocumentsDocumentIdDownloadGet(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsIndividualsConsentIdDocumentsDocumentIdDownloadGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsIndividualsConsentIdDocumentsDocumentIdDownloadGet`: UserDocumentDownloadDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsIndividualsConsentIdDocumentsDocumentIdDownloadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** | consentId. | 
**documentId** | **string** | documentId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsIndividualsConsentIdDocumentsDocumentIdDownloadGetRequest struct via the builder pattern


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


## V1ConsentsIndividualsConsentIdDocumentsDocumentIdGet

> UserDocumentDetailsDto V1ConsentsIndividualsConsentIdDocumentsDocumentIdGet(ctx, consentId, documentId).Execute()

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Document Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsIndividualsConsentIdDocumentsDocumentIdGet(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsIndividualsConsentIdDocumentsDocumentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsIndividualsConsentIdDocumentsDocumentIdGet`: UserDocumentDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsIndividualsConsentIdDocumentsDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**documentId** | **string** | Document Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsIndividualsConsentIdDocumentsDocumentIdGetRequest struct via the builder pattern


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


## V1ConsentsIndividualsConsentIdDocumentsGet

> DataConsentDocumentsDto V1ConsentsIndividualsConsentIdDocumentsGet(ctx, consentId).Execute()

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsIndividualsConsentIdDocumentsGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsIndividualsConsentIdDocumentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsIndividualsConsentIdDocumentsGet`: DataConsentDocumentsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsIndividualsConsentIdDocumentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsIndividualsConsentIdDocumentsGetRequest struct via the builder pattern


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


## V1ConsentsIndividualsConsentIdGet

> DataConsentDetailsDto V1ConsentsIndividualsConsentIdGet(ctx, consentId).Execute()

Get individuals consent details by consent id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsIndividualsConsentIdGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsIndividualsConsentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsIndividualsConsentIdGet`: DataConsentDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsIndividualsConsentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsIndividualsConsentIdGetRequest struct via the builder pattern


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


## V1ConsentsIndividualsGet

> UserDataConsentInfoDtoPaginatedList V1ConsentsIndividualsGet(ctx).PageNo(pageNo).PageSize(pageSize).Status(status).StartDate(startDate).EndDate(endDate).Execute()

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
    pageNo := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    status := openapiclient.DataConsentStatus("Pending") // DataConsentStatus |  (optional)
    startDate := time.Now() // time.Time |  (optional)
    endDate := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsIndividualsGet(context.Background()).PageNo(pageNo).PageSize(pageSize).Status(status).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsIndividualsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsIndividualsGet`: UserDataConsentInfoDtoPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsIndividualsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsIndividualsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNo** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 

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


## V1ConsentsOrganizationsConsentIdAccountsAccountIdGet

> OrganizationFinancialAccountDto V1ConsentsOrganizationsConsentIdAccountsAccountIdGet(ctx, consentId, accountId).Execute()

Get orgnization consented financial account details based on account id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsOrganizationsConsentIdAccountsAccountIdGet(context.Background(), consentId, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsOrganizationsConsentIdAccountsAccountIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsOrganizationsConsentIdAccountsAccountIdGet`: OrganizationFinancialAccountDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsOrganizationsConsentIdAccountsAccountIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsOrganizationsConsentIdAccountsAccountIdGetRequest struct via the builder pattern


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


## V1ConsentsOrganizationsConsentIdAccountsAccountIdTransactionsGet

> OrganizationFinancialTransactionsDtoPaginatedList V1ConsentsOrganizationsConsentIdAccountsAccountIdTransactionsGet(ctx, consentId, accountId).Filters(filters).PageNo(pageNo).PageSize(pageSize).FromDate(fromDate).ToDate(toDate).Execute()

Get consented financial account transactions of an organization based on accountId.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    filters := "filters_example" // string |  (optional)
    pageNo := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    fromDate := time.Now() // time.Time |  (optional)
    toDate := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsOrganizationsConsentIdAccountsAccountIdTransactionsGet(context.Background(), consentId, accountId).Filters(filters).PageNo(pageNo).PageSize(pageSize).FromDate(fromDate).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsOrganizationsConsentIdAccountsAccountIdTransactionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsOrganizationsConsentIdAccountsAccountIdTransactionsGet`: OrganizationFinancialTransactionsDtoPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsOrganizationsConsentIdAccountsAccountIdTransactionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsOrganizationsConsentIdAccountsAccountIdTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filters** | **string** |  | 
 **pageNo** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **fromDate** | **time.Time** |  | 
 **toDate** | **time.Time** |  | 

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


## V1ConsentsOrganizationsConsentIdAccountsGet

> DataConsentFinancialsDto V1ConsentsOrganizationsConsentIdAccountsGet(ctx, consentId).Execute()

Get all organizational financial accounts in a consent.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsOrganizationsConsentIdAccountsGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsOrganizationsConsentIdAccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsOrganizationsConsentIdAccountsGet`: DataConsentFinancialsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsOrganizationsConsentIdAccountsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsOrganizationsConsentIdAccountsGetRequest struct via the builder pattern


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


## V1ConsentsOrganizationsConsentIdDocumentsDocumentIdDownloadGet

> OrganizationDocumentDownloadDto V1ConsentsOrganizationsConsentIdDocumentsDocumentIdDownloadGet(ctx, consentId, documentId).Execute()

Download organizations consented document.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsOrganizationsConsentIdDocumentsDocumentIdDownloadGet(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsOrganizationsConsentIdDocumentsDocumentIdDownloadGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsOrganizationsConsentIdDocumentsDocumentIdDownloadGet`: OrganizationDocumentDownloadDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsOrganizationsConsentIdDocumentsDocumentIdDownloadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsOrganizationsConsentIdDocumentsDocumentIdDownloadGetRequest struct via the builder pattern


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


## V1ConsentsOrganizationsConsentIdDocumentsDocumentIdGet

> OrganizationDocumentDetailsDto V1ConsentsOrganizationsConsentIdDocumentsDocumentIdGet(ctx, consentId, documentId).Execute()

Get organizations consent document based on document id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsOrganizationsConsentIdDocumentsDocumentIdGet(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsOrganizationsConsentIdDocumentsDocumentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsOrganizationsConsentIdDocumentsDocumentIdGet`: OrganizationDocumentDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsOrganizationsConsentIdDocumentsDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsOrganizationsConsentIdDocumentsDocumentIdGetRequest struct via the builder pattern


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


## V1ConsentsOrganizationsConsentIdDocumentsGet

> DataConsentDocumentsDto V1ConsentsOrganizationsConsentIdDocumentsGet(ctx, consentId).Execute()

Get the organizations documents based on ConsentId.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsOrganizationsConsentIdDocumentsGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsOrganizationsConsentIdDocumentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsOrganizationsConsentIdDocumentsGet`: DataConsentDocumentsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsOrganizationsConsentIdDocumentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsOrganizationsConsentIdDocumentsGetRequest struct via the builder pattern


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


## V1ConsentsOrganizationsConsentIdGet

> DataConsentDetailsDto V1ConsentsOrganizationsConsentIdGet(ctx, consentId).Execute()

Get organizations consent details by consent id.

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
    consentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsOrganizationsConsentIdGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsOrganizationsConsentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsOrganizationsConsentIdGet`: DataConsentDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsOrganizationsConsentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsOrganizationsConsentIdGetRequest struct via the builder pattern


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


## V1ConsentsOrganizationsGet

> OrganizationDataConsentInfoDtoPaginatedList V1ConsentsOrganizationsGet(ctx).PageNo(pageNo).PageSize(pageSize).Status(status).StartDate(startDate).EndDate(endDate).Execute()

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
    pageNo := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    status := openapiclient.DataConsentStatus("Pending") // DataConsentStatus |  (optional)
    startDate := time.Now() // time.Time |  (optional)
    endDate := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataConsentsApi.V1ConsentsOrganizationsGet(context.Background()).PageNo(pageNo).PageSize(pageSize).Status(status).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsOrganizationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsOrganizationsGet`: OrganizationDataConsentInfoDtoPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsOrganizationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsOrganizationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNo** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 

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

