# \DigiLockerCompatIssuerApi

All URIs are relative to *https://api.mydatamyconsent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DigilockerCompatIssueDocument**](DigiLockerCompatIssuerApi.md#DigilockerCompatIssueDocument) | **Post** /issuer/issuedoc/1/xml | Digilocker Compatible endpoint to issue document.



## DigilockerCompatIssueDocument

> PushUriResponse DigilockerCompatIssueDocument(ctx).PushUriRequest(pushUriRequest).Execute()

Digilocker Compatible endpoint to issue document.

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
    pushUriRequest := *openapiclient.NewPushUriRequest(*openapiclient.NewUriDetails("Aadhaar_example", "Uri_example", "DocType_example", "DocName_example", "DocId_example", "IssuedOn_example", "ValidFrom_example")) // PushUriRequest | Push URI request payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigiLockerCompatIssuerApi.DigilockerCompatIssueDocument(context.Background()).PushUriRequest(pushUriRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigiLockerCompatIssuerApi.DigilockerCompatIssueDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DigilockerCompatIssueDocument`: PushUriResponse
    fmt.Fprintf(os.Stdout, "Response from `DigiLockerCompatIssuerApi.DigilockerCompatIssueDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDigilockerCompatIssueDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushUriRequest** | [**PushUriRequest**](PushUriRequest.md) | Push URI request payload | 

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

