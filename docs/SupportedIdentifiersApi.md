# \SupportedIdentifiersApi

All URIs are relative to *https://api.mydatamyconsent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SupportedIdentifiersCountryIso2CodeGet**](SupportedIdentifiersApi.md#V1SupportedIdentifiersCountryIso2CodeGet) | **Get** /v1/supported-identifiers/{country_iso2_code} | Get all supported identifiers by country.



## V1SupportedIdentifiersCountryIso2CodeGet

> SupportedIdentifier V1SupportedIdentifiersCountryIso2CodeGet(ctx, countryIso2Code).Execute()

Get all supported identifiers by country.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SupportedIdentifiersApi.V1SupportedIdentifiersCountryIso2CodeGet(context.Background(), countryIso2Code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportedIdentifiersApi.V1SupportedIdentifiersCountryIso2CodeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SupportedIdentifiersCountryIso2CodeGet`: SupportedIdentifier
    fmt.Fprintf(os.Stdout, "Response from `SupportedIdentifiersApi.V1SupportedIdentifiersCountryIso2CodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryIso2Code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SupportedIdentifiersCountryIso2CodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SupportedIdentifier**](SupportedIdentifier.md)

### Authorization

[OAuth2ClientCredentials](../README.md#OAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

