# \DataProvidersDiscoveryApi

All URIs are relative to *https://api.mydatamyconsent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DataProvidersGet**](DataProvidersDiscoveryApi.md#V1DataProvidersGet) | **Get** /v1/data-providers | Discover all data providers in my data my consent by country and filters.
[**V1DataProvidersIdGet**](DataProvidersDiscoveryApi.md#V1DataProvidersIdGet) | **Get** /v1/data-providers/{id} | Get a data provider details by provider id.



## V1DataProvidersGet

> PaginatedListOfDataProviders V1DataProvidersGet(ctx).CountryIso2Code(countryIso2Code).PageNo(pageNo).PageSize(pageSize).Execute()

Discover all data providers in my data my consent by country and filters.

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
    countryIso2Code := "countryIso2Code_example" // string | 
    pageNo := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProvidersDiscoveryApi.V1DataProvidersGet(context.Background()).CountryIso2Code(countryIso2Code).PageNo(pageNo).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProvidersDiscoveryApi.V1DataProvidersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DataProvidersGet`: PaginatedListOfDataProviders
    fmt.Fprintf(os.Stdout, "Response from `DataProvidersDiscoveryApi.V1DataProvidersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DataProvidersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryIso2Code** | **string** |  | 
 **pageNo** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**PaginatedListOfDataProviders**](PaginatedListOfDataProviders.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataProvidersIdGet

> DataProviderDetails V1DataProvidersIdGet(ctx, id).Execute()

Get a data provider details by provider id.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataProvidersDiscoveryApi.V1DataProvidersIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataProvidersDiscoveryApi.V1DataProvidersIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DataProvidersIdGet`: DataProviderDetails
    fmt.Fprintf(os.Stdout, "Response from `DataProvidersDiscoveryApi.V1DataProvidersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DataProvidersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataProviderDetails**](DataProviderDetails.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

