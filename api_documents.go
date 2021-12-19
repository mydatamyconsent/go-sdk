/*
My Data My Consent - Developer API

Unleashing the power of data consent by establishing trust. The Platform Core Developer API defines a set of capabilities that can be used to request, issue, manage and update data, documents and credentials by organizations. The API can be used to request, manage and update Decentralised Identifiers, Financial Data, Health Data issue Documents, Credentials directly or using OpenID Connect flows, and verify Messages signed with DIDs and much more.

API version: v1
Contact: support@mydatamyconsent.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mydatamyconsent

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"time"
)

// Linger please
var (
	_ _context.Context
)

// DocumentsApiService DocumentsApi service
type DocumentsApiService service

type ApiIssueDocumentRequest struct {
	ctx _context.Context
	ApiService *DocumentsApiService
	documentIssueRequest *DocumentIssueRequest
}

func (r ApiIssueDocumentRequest) DocumentIssueRequest(documentIssueRequest DocumentIssueRequest) ApiIssueDocumentRequest {
	r.documentIssueRequest = &documentIssueRequest
	return r
}

func (r ApiIssueDocumentRequest) Execute() (bool, *_nethttp.Response, error) {
	return r.ApiService.IssueDocumentExecute(r)
}

/*
IssueDocument Issue a new document.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIssueDocumentRequest
*/
func (a *DocumentsApiService) IssueDocument(ctx _context.Context) ApiIssueDocumentRequest {
	return ApiIssueDocumentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return bool
func (a *DocumentsApiService) IssueDocumentExecute(r ApiIssueDocumentRequest) (bool, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocumentsApiService.IssueDocument")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/documents/issue"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.documentIssueRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1DocumentsIssuedDocumentIdGetRequest struct {
	ctx _context.Context
	ApiService *DocumentsApiService
	documentId string
}


func (r ApiV1DocumentsIssuedDocumentIdGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1DocumentsIssuedDocumentIdGetExecute(r)
}

/*
V1DocumentsIssuedDocumentIdGet Get issued document.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param documentId Document id.
 @return ApiV1DocumentsIssuedDocumentIdGetRequest
*/
func (a *DocumentsApiService) V1DocumentsIssuedDocumentIdGet(ctx _context.Context, documentId string) ApiV1DocumentsIssuedDocumentIdGetRequest {
	return ApiV1DocumentsIssuedDocumentIdGetRequest{
		ApiService: a,
		ctx: ctx,
		documentId: documentId,
	}
}

// Execute executes the request
func (a *DocumentsApiService) V1DocumentsIssuedDocumentIdGetExecute(r ApiV1DocumentsIssuedDocumentIdGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocumentsApiService.V1DocumentsIssuedDocumentIdGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/documents/issued/{documentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"documentId"+"}", _neturl.PathEscape(parameterToString(r.documentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1DocumentsIssuedGetRequest struct {
	ctx _context.Context
	ApiService *DocumentsApiService
	documentTypeId *string
	fromDateTime *time.Time
	toDateTime *time.Time
	pageSize *int32
	pageNo *int32
}

func (r ApiV1DocumentsIssuedGetRequest) DocumentTypeId(documentTypeId string) ApiV1DocumentsIssuedGetRequest {
	r.documentTypeId = &documentTypeId
	return r
}
func (r ApiV1DocumentsIssuedGetRequest) FromDateTime(fromDateTime time.Time) ApiV1DocumentsIssuedGetRequest {
	r.fromDateTime = &fromDateTime
	return r
}
func (r ApiV1DocumentsIssuedGetRequest) ToDateTime(toDateTime time.Time) ApiV1DocumentsIssuedGetRequest {
	r.toDateTime = &toDateTime
	return r
}
func (r ApiV1DocumentsIssuedGetRequest) PageSize(pageSize int32) ApiV1DocumentsIssuedGetRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiV1DocumentsIssuedGetRequest) PageNo(pageNo int32) ApiV1DocumentsIssuedGetRequest {
	r.pageNo = &pageNo
	return r
}

func (r ApiV1DocumentsIssuedGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1DocumentsIssuedGetExecute(r)
}

/*
V1DocumentsIssuedGet Get issued documents.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1DocumentsIssuedGetRequest
*/
func (a *DocumentsApiService) V1DocumentsIssuedGet(ctx _context.Context) ApiV1DocumentsIssuedGetRequest {
	return ApiV1DocumentsIssuedGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DocumentsApiService) V1DocumentsIssuedGetExecute(r ApiV1DocumentsIssuedGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocumentsApiService.V1DocumentsIssuedGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/documents/issued"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.documentTypeId != nil {
		localVarQueryParams.Add("documentTypeId", parameterToString(*r.documentTypeId, ""))
	}
	if r.fromDateTime != nil {
		localVarQueryParams.Add("fromDateTime", parameterToString(*r.fromDateTime, ""))
	}
	if r.toDateTime != nil {
		localVarQueryParams.Add("toDateTime", parameterToString(*r.toDateTime, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.pageNo != nil {
		localVarQueryParams.Add("pageNo", parameterToString(*r.pageNo, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1DocumentsTypesGetRequest struct {
	ctx _context.Context
	ApiService *DocumentsApiService
	pageSize *int32
	pageNo *int32
}

func (r ApiV1DocumentsTypesGetRequest) PageSize(pageSize int32) ApiV1DocumentsTypesGetRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiV1DocumentsTypesGetRequest) PageNo(pageNo int32) ApiV1DocumentsTypesGetRequest {
	r.pageNo = &pageNo
	return r
}

func (r ApiV1DocumentsTypesGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1DocumentsTypesGetExecute(r)
}

/*
V1DocumentsTypesGet Get registered document types.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1DocumentsTypesGetRequest
*/
func (a *DocumentsApiService) V1DocumentsTypesGet(ctx _context.Context) ApiV1DocumentsTypesGetRequest {
	return ApiV1DocumentsTypesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DocumentsApiService) V1DocumentsTypesGetExecute(r ApiV1DocumentsTypesGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocumentsApiService.V1DocumentsTypesGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/documents/types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.pageNo != nil {
		localVarQueryParams.Add("pageNo", parameterToString(*r.pageNo, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
