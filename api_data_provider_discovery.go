/*
My Data My Consent - Developer API

Unleashing the power of data consent by establishing trust. The Platform Core Developer API defines a set of capabilities that can be used to request, issue, manage and update data, documents and credentials by organizations. The API can be used to request, manage and update Decentralised Identifiers, Financial Data, Health Data issue Documents, Credentials directly or using OpenID Connect flows, and verify Messages signed with DIDs and much more.

API version: v1
Contact: support@mydatamyconsent.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/mydatamyconsent/sdk

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// DataProviderDiscoveryApiService DataProviderDiscoveryApi service
type DataProviderDiscoveryApiService service

type ApiGetDataProviderByIdRequest struct {
	ctx context.Context
	ApiService *DataProviderDiscoveryApiService
	providerId string
}


func (r ApiGetDataProviderByIdRequest) Execute() (*DataProvider, *http.Response, error) {
	return r.ApiService.GetDataProviderByIdExecute(r)
}

/*
GetDataProviderById Get a Data Provider details based on provider id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param providerId Provider id.
 @return ApiGetDataProviderByIdRequest
*/
func (a *DataProviderDiscoveryApiService) GetDataProviderById(ctx context.Context, providerId string) ApiGetDataProviderByIdRequest {
	return ApiGetDataProviderByIdRequest{
		ApiService: a,
		ctx: ctx,
		providerId: providerId,
	}
}

// Execute executes the request
//  @return DataProvider
func (a *DataProviderDiscoveryApiService) GetDataProviderByIdExecute(r ApiGetDataProviderByIdRequest) (*DataProvider, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DataProvider
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataProviderDiscoveryApiService.GetDataProviderById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/data-providers/{providerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"providerId"+"}", url.PathEscape(parameterToString(r.providerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetDataProvidersRequest struct {
	ctx context.Context
	ApiService *DataProviderDiscoveryApiService
	accountType *string
	documentType *string
	organizationCategory *string
	pageNo *int32
	pageSize *int32
	country *string
}

// Account type.
func (r ApiGetDataProvidersRequest) AccountType(accountType string) ApiGetDataProvidersRequest {
	r.accountType = &accountType
	return r
}
// Document type.
func (r ApiGetDataProvidersRequest) DocumentType(documentType string) ApiGetDataProvidersRequest {
	r.documentType = &documentType
	return r
}
// Organization category.
func (r ApiGetDataProvidersRequest) OrganizationCategory(organizationCategory string) ApiGetDataProvidersRequest {
	r.organizationCategory = &organizationCategory
	return r
}
// Page number.
func (r ApiGetDataProvidersRequest) PageNo(pageNo int32) ApiGetDataProvidersRequest {
	r.pageNo = &pageNo
	return r
}
// Number of items to return.
func (r ApiGetDataProvidersRequest) PageSize(pageSize int32) ApiGetDataProvidersRequest {
	r.pageSize = &pageSize
	return r
}
// ISO2 Country code.
func (r ApiGetDataProvidersRequest) Country(country string) ApiGetDataProvidersRequest {
	r.country = &country
	return r
}

func (r ApiGetDataProvidersRequest) Execute() (*DataProviderPaginatedList, *http.Response, error) {
	return r.ApiService.GetDataProvidersExecute(r)
}

/*
GetDataProviders Discover all data providers in My Data My Consent by country and filters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDataProvidersRequest
*/
func (a *DataProviderDiscoveryApiService) GetDataProviders(ctx context.Context) ApiGetDataProvidersRequest {
	return ApiGetDataProvidersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DataProviderPaginatedList
func (a *DataProviderDiscoveryApiService) GetDataProvidersExecute(r ApiGetDataProvidersRequest) (*DataProviderPaginatedList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DataProviderPaginatedList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataProviderDiscoveryApiService.GetDataProviders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/data-providers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.accountType != nil {
		localVarQueryParams.Add("accountType", parameterToString(*r.accountType, ""))
	}
	if r.documentType != nil {
		localVarQueryParams.Add("documentType", parameterToString(*r.documentType, ""))
	}
	if r.organizationCategory != nil {
		localVarQueryParams.Add("organizationCategory", parameterToString(*r.organizationCategory, ""))
	}
	if r.pageNo != nil {
		localVarQueryParams.Add("pageNo", parameterToString(*r.pageNo, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.country != nil {
		localVarQueryParams.Add("country", parameterToString(*r.country, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
