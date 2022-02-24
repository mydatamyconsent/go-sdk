# \SupportedIdentifiersApi

All URIs are relative to *https://api.mydatamyconsent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllSupportedIdentifiers**](SupportedIdentifiersApi.md#GetAllSupportedIdentifiers) | **Get** /v1/supported-identifiers/{countryIso2Code} | Get all supported identifiers by country.



## GetAllSupportedIdentifiers

> SupportedIdentifier GetAllSupportedIdentifiers(ctx, countryIso2Code).Execute()

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
    countryIso2Code := "countryIso2Code_example" // string | Country ISO 2 code.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SupportedIdentifiersApi.GetAllSupportedIdentifiers(context.Background(), countryIso2Code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportedIdentifiersApi.GetAllSupportedIdentifiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSupportedIdentifiers`: SupportedIdentifier
    fmt.Fprintf(os.Stdout, "Response from `SupportedIdentifiersApi.GetAllSupportedIdentifiers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryIso2Code** | **string** | Country ISO 2 code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSupportedIdentifiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SupportedIdentifier**](SupportedIdentifier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

