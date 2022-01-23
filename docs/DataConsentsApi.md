# \DataConsentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ConsentsConsentIdAccountsAccountIdGet**](DataConsentsApi.md#V1ConsentsConsentIdAccountsAccountIdGet) | **Get** /v1/consents/{consentId}/accounts/{accountId} | Get consented financial account details.
[**V1ConsentsConsentIdAccountsAccountIdInsightsGet**](DataConsentsApi.md#V1ConsentsConsentIdAccountsAccountIdInsightsGet) | **Get** /v1/consents/{consentId}/accounts/{accountId}/insights | Get consented financial account insights.
[**V1ConsentsConsentIdAccountsAccountIdTransactionsGet**](DataConsentsApi.md#V1ConsentsConsentIdAccountsAccountIdTransactionsGet) | **Get** /v1/consents/{consentId}/accounts/{accountId}/transactions | Get consented financial account transactions.
[**V1ConsentsConsentIdAccountsGet**](DataConsentsApi.md#V1ConsentsConsentIdAccountsGet) | **Get** /v1/consents/{consentId}/accounts | Get all accounts in a consent.
[**V1ConsentsConsentIdDocumentsDocumentIdAnalysisGet**](DataConsentsApi.md#V1ConsentsConsentIdDocumentsDocumentIdAnalysisGet) | **Get** /v1/consents/{consentId}/documents/{documentId}/analysis | Get analysis of a consented document.
[**V1ConsentsConsentIdDocumentsDocumentIdDownloadGet**](DataConsentsApi.md#V1ConsentsConsentIdDocumentsDocumentIdDownloadGet) | **Get** /v1/consents/{consentId}/documents/{documentId}/download | Download a consented document.
[**V1ConsentsConsentIdDocumentsDocumentIdGet**](DataConsentsApi.md#V1ConsentsConsentIdDocumentsDocumentIdGet) | **Get** /v1/consents/{consentId}/documents/{documentId} | Get consented document details.
[**V1ConsentsConsentIdDocumentsGet**](DataConsentsApi.md#V1ConsentsConsentIdDocumentsGet) | **Get** /v1/consents/{consentId}/documents | Get all documents in a consent.
[**V1ConsentsConsentIdGet**](DataConsentsApi.md#V1ConsentsConsentIdGet) | **Get** /v1/consents/{consentId} | Get consent details by consent id.
[**V1ConsentsGet**](DataConsentsApi.md#V1ConsentsGet) | **Get** /v1/consents | Get all consents filtered by status and time.



## V1ConsentsConsentIdAccountsAccountIdGet

> map[string]interface{} V1ConsentsConsentIdAccountsAccountIdGet(ctx, consentId, accountId).Execute()

Get consented financial account details.

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataConsentsApi.V1ConsentsConsentIdAccountsAccountIdGet(context.Background(), consentId, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsConsentIdAccountsAccountIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsConsentIdAccountsAccountIdGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsConsentIdAccountsAccountIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsConsentIdAccountsAccountIdGetRequest struct via the builder pattern


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


## V1ConsentsConsentIdAccountsAccountIdInsightsGet

> map[string]interface{} V1ConsentsConsentIdAccountsAccountIdInsightsGet(ctx, consentId, accountId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataConsentsApi.V1ConsentsConsentIdAccountsAccountIdInsightsGet(context.Background(), consentId, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsConsentIdAccountsAccountIdInsightsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsConsentIdAccountsAccountIdInsightsGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsConsentIdAccountsAccountIdInsightsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsConsentIdAccountsAccountIdInsightsGetRequest struct via the builder pattern


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


## V1ConsentsConsentIdAccountsAccountIdTransactionsGet

> map[string]interface{} V1ConsentsConsentIdAccountsAccountIdTransactionsGet(ctx, consentId, accountId).Filters(filters).FromDate(fromDate).ToDate(toDate).Execute()

Get consented financial account transactions.

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
    filters := "filters_example" // string |  (optional)
    fromDate := time.Now() // time.Time |  (optional)
    toDate := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataConsentsApi.V1ConsentsConsentIdAccountsAccountIdTransactionsGet(context.Background(), consentId, accountId).Filters(filters).FromDate(fromDate).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsConsentIdAccountsAccountIdTransactionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsConsentIdAccountsAccountIdTransactionsGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsConsentIdAccountsAccountIdTransactionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsConsentIdAccountsAccountIdTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filters** | **string** |  | 
 **fromDate** | **time.Time** |  | 
 **toDate** | **time.Time** |  | 

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


## V1ConsentsConsentIdAccountsGet

> map[string]interface{} V1ConsentsConsentIdAccountsGet(ctx, consentId).Execute()

Get all accounts in a consent.

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataConsentsApi.V1ConsentsConsentIdAccountsGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsConsentIdAccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsConsentIdAccountsGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsConsentIdAccountsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsConsentIdAccountsGetRequest struct via the builder pattern


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


## V1ConsentsConsentIdDocumentsDocumentIdAnalysisGet

> map[string]interface{} V1ConsentsConsentIdDocumentsDocumentIdAnalysisGet(ctx, consentId, documentId).Execute()

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
    consentId := "consentId_example" // string | 
    documentId := "documentId_example" // string | Document Id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataConsentsApi.V1ConsentsConsentIdDocumentsDocumentIdAnalysisGet(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsConsentIdDocumentsDocumentIdAnalysisGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsConsentIdDocumentsDocumentIdAnalysisGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsConsentIdDocumentsDocumentIdAnalysisGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**documentId** | **string** | Document Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsConsentIdDocumentsDocumentIdAnalysisGetRequest struct via the builder pattern


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


## V1ConsentsConsentIdDocumentsDocumentIdDownloadGet

> map[string]interface{} V1ConsentsConsentIdDocumentsDocumentIdDownloadGet(ctx, consentId, documentId).Execute()

Download a consented document.

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
    documentId := "documentId_example" // string | Document Id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataConsentsApi.V1ConsentsConsentIdDocumentsDocumentIdDownloadGet(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsConsentIdDocumentsDocumentIdDownloadGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsConsentIdDocumentsDocumentIdDownloadGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsConsentIdDocumentsDocumentIdDownloadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**documentId** | **string** | Document Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsConsentIdDocumentsDocumentIdDownloadGetRequest struct via the builder pattern


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


## V1ConsentsConsentIdDocumentsDocumentIdGet

> map[string]interface{} V1ConsentsConsentIdDocumentsDocumentIdGet(ctx, consentId, documentId).Execute()

Get consented document details.

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
    documentId := "documentId_example" // string | Document Id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataConsentsApi.V1ConsentsConsentIdDocumentsDocumentIdGet(context.Background(), consentId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsConsentIdDocumentsDocumentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsConsentIdDocumentsDocumentIdGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsConsentIdDocumentsDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 
**documentId** | **string** | Document Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsConsentIdDocumentsDocumentIdGetRequest struct via the builder pattern


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


## V1ConsentsConsentIdDocumentsGet

> map[string]interface{} V1ConsentsConsentIdDocumentsGet(ctx, consentId).Execute()

Get all documents in a consent.

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataConsentsApi.V1ConsentsConsentIdDocumentsGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsConsentIdDocumentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsConsentIdDocumentsGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsConsentIdDocumentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsConsentIdDocumentsGetRequest struct via the builder pattern


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


## V1ConsentsConsentIdGet

> map[string]interface{} V1ConsentsConsentIdGet(ctx, consentId).Execute()

Get consent details by consent id.

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataConsentsApi.V1ConsentsConsentIdGet(context.Background(), consentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsConsentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsConsentIdGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsConsentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsConsentIdGetRequest struct via the builder pattern


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


## V1ConsentsGet

> map[string]interface{} V1ConsentsGet(ctx).Status(status).StartDate(startDate).EndDate(endDate).Execute()

Get all consents filtered by status and time.

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
    status := openapiclient.DataConsentStatus("Pending") // DataConsentStatus | MyDataMyConsent.Domain.Entities.ConsentAggregate.Enums.DataConsentStatus. (optional)
    startDate := time.Now() // time.Time | System.DateTime. (optional)
    endDate := time.Now() // time.Time | till dateSystem.DateTime. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataConsentsApi.V1ConsentsGet(context.Background()).Status(status).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataConsentsApi.V1ConsentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ConsentsGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataConsentsApi.V1ConsentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ConsentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**DataConsentStatus**](DataConsentStatus.md) | MyDataMyConsent.Domain.Entities.ConsentAggregate.Enums.DataConsentStatus. | 
 **startDate** | **time.Time** | System.DateTime. | 
 **endDate** | **time.Time** | till dateSystem.DateTime. | 

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

