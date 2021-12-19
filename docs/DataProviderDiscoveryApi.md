# MyDataMyConsent\DataProviderDiscoveryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DataProvidersGet**](DataProviderDiscoveryApi.md#V1DataProvidersGet) | **Get** /v1/data-providers | Discover all data providers in My Data My Consent by country and filters.
[**V1DataProvidersProviderIdGet**](DataProviderDiscoveryApi.md#V1DataProvidersProviderIdGet) | **Get** /v1/data-providers/{providerId} | Get a Data Provider details.



## V1DataProvidersGet

> DataProviderPaginatedList V1DataProvidersGet(ctx).AccountType(accountType).DocumentType(documentType).OrganizationCategory(organizationCategory).PageNo(pageNo).PageSize(pageSize).Country(country).Execute()

Discover all data providers in My Data My Consent by country and filters.



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
    accountType := "accountType_example" // string | Account type. (optional)
    documentType := "documentType_example" // string | Document type. (optional)
    organizationCategory := "organizationCategory_example" // string | Organization category. (optional)
    pageNo := int32(56) // int32 | Page number. (optional)
    pageSize := int32(56) // int32 | Page size. (optional)
    country := "country_example" // string | ISO2 Country code. (optional) (default to "IN")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataProviderDiscoveryApi.V1DataProvidersGet(context.Background()).AccountType(accountType).DocumentType(documentType).OrganizationCategory(organizationCategory).PageNo(pageNo).PageSize(pageSize).Country(country).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProviderDiscoveryApi.V1DataProvidersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DataProvidersGet`: DataProviderPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataProviderDiscoveryApi.V1DataProvidersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DataProvidersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountType** | **string** | Account type. | 
 **documentType** | **string** | Document type. | 
 **organizationCategory** | **string** | Organization category. | 
 **pageNo** | **int32** | Page number. | 
 **pageSize** | **int32** | Page size. | 
 **country** | **string** | ISO2 Country code. | [default to &quot;IN&quot;]

### Return type

[**DataProviderPaginatedList**](DataProviderPaginatedList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataProvidersProviderIdGet

> DataProvider V1DataProvidersProviderIdGet(ctx, providerId).Execute()

Get a Data Provider details.



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
    providerId := "providerId_example" // string | Provider Id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataProviderDiscoveryApi.V1DataProvidersProviderIdGet(context.Background(), providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProviderDiscoveryApi.V1DataProvidersProviderIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DataProvidersProviderIdGet`: DataProvider
    fmt.Fprintf(os.Stdout, "Response from `DataProviderDiscoveryApi.V1DataProvidersProviderIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** | Provider Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DataProvidersProviderIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataProvider**](DataProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

