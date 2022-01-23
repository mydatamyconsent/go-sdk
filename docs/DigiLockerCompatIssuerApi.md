# \DigiLockerCompatIssuerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssuerIssuedoc1XmlPost**](DigiLockerCompatIssuerApi.md#IssuerIssuedoc1XmlPost) | **Post** /issuer/issuedoc/1/xml | Digilocker Compatible endpoint to Issue Documents.



## IssuerIssuedoc1XmlPost

> PushUriResponse IssuerIssuedoc1XmlPost(ctx).PushUriRequest(pushUriRequest).Execute()

Digilocker Compatible endpoint to Issue Documents.

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
    pushUriRequest := *openapiclient.NewPushUriRequest() // PushUriRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DigiLockerCompatIssuerApi.IssuerIssuedoc1XmlPost(context.Background()).PushUriRequest(pushUriRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigiLockerCompatIssuerApi.IssuerIssuedoc1XmlPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuerIssuedoc1XmlPost`: PushUriResponse
    fmt.Fprintf(os.Stdout, "Response from `DigiLockerCompatIssuerApi.IssuerIssuedoc1XmlPost`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

