# \DataProviderDiscoveryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataProviderById**](DataProviderDiscoveryApi.md#GetDataProviderById) | **Get** /v1/data-providers/{providerId} | Get a Data Provider details based on provider id.
[**GetDataProviders**](DataProviderDiscoveryApi.md#GetDataProviders) | **Get** /v1/data-providers | Discover all data providers in My Data My Consent by country and filters.



## GetDataProviderById

> DataProvider GetDataProviderById(ctx, providerId).Execute()

Get a Data Provider details based on provider id.

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
    providerId := "providerId_example" // string | Provider id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProviderDiscoveryApi.GetDataProviderById(context.Background(), providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProviderDiscoveryApi.GetDataProviderById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataProviderById`: DataProvider
    fmt.Fprintf(os.Stdout, "Response from `DataProviderDiscoveryApi.GetDataProviderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** | Provider id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataProviderByIdRequest struct via the builder pattern


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


## GetDataProviders

> DataProviderPaginatedList GetDataProviders(ctx).AccountType(accountType).DocumentType(documentType).OrganizationCategory(organizationCategory).PageNo(pageNo).PageSize(pageSize).Country(country).Execute()

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
    pageNo := int32(56) // int32 | Page number. (optional) (default to 1)
    pageSize := int32(56) // int32 | Number of items to return. (optional) (default to 25)
    country := "country_example" // string | ISO2 Country code. (optional) (default to "IN")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProviderDiscoveryApi.GetDataProviders(context.Background()).AccountType(accountType).DocumentType(documentType).OrganizationCategory(organizationCategory).PageNo(pageNo).PageSize(pageSize).Country(country).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProviderDiscoveryApi.GetDataProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataProviders`: DataProviderPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DataProviderDiscoveryApi.GetDataProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountType** | **string** | Account type. | 
 **documentType** | **string** | Document type. | 
 **organizationCategory** | **string** | Organization category. | 
 **pageNo** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Number of items to return. | [default to 25]
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

