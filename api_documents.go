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
	"time"
	"os"
)


// DocumentsApiService DocumentsApi service
type DocumentsApiService service

type ApiGetIssuedDocumentByIdRequest struct {
	ctx context.Context
	ApiService *DocumentsApiService
	documentId string
}

func (r ApiGetIssuedDocumentByIdRequest) Execute() (*IssuedDocumentDetails, *http.Response, error) {
	return r.ApiService.GetIssuedDocumentByIdExecute(r)
}

/*
GetIssuedDocumentById Get issued document.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param documentId Document id.
 @return ApiGetIssuedDocumentByIdRequest
*/
func (a *DocumentsApiService) GetIssuedDocumentById(ctx context.Context, documentId string) ApiGetIssuedDocumentByIdRequest {
	return ApiGetIssuedDocumentByIdRequest{
		ApiService: a,
		ctx: ctx,
		documentId: documentId,
	}
}

// Execute executes the request
//  @return IssuedDocumentDetails
func (a *DocumentsApiService) GetIssuedDocumentByIdExecute(r ApiGetIssuedDocumentByIdRequest) (*IssuedDocumentDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IssuedDocumentDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocumentsApiService.GetIssuedDocumentById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/documents/issued/{documentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"documentId"+"}", url.PathEscape(parameterToString(r.documentId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
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

type ApiGetIssuedDocumentsRequest struct {
	ctx context.Context
	ApiService *DocumentsApiService
	documentTypeId *string
	fromDateTime *time.Time
	toDateTime *time.Time
	pageNo *int32
	pageSize *int32
}

// Document type id.
func (r ApiGetIssuedDocumentsRequest) DocumentTypeId(documentTypeId string) ApiGetIssuedDocumentsRequest {
	r.documentTypeId = &documentTypeId
	return r
}

// From DateTime in UTC timezone.
func (r ApiGetIssuedDocumentsRequest) FromDateTime(fromDateTime time.Time) ApiGetIssuedDocumentsRequest {
	r.fromDateTime = &fromDateTime
	return r
}

// To DateTime in UTC timezone.
func (r ApiGetIssuedDocumentsRequest) ToDateTime(toDateTime time.Time) ApiGetIssuedDocumentsRequest {
	r.toDateTime = &toDateTime
	return r
}

// Page number.
func (r ApiGetIssuedDocumentsRequest) PageNo(pageNo int32) ApiGetIssuedDocumentsRequest {
	r.pageNo = &pageNo
	return r
}

// Number of items to return.
func (r ApiGetIssuedDocumentsRequest) PageSize(pageSize int32) ApiGetIssuedDocumentsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetIssuedDocumentsRequest) Execute() (*IssuedDocumentPaginatedList, *http.Response, error) {
	return r.ApiService.GetIssuedDocumentsExecute(r)
}

/*
GetIssuedDocuments Get paginated list of issued documents of given document type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetIssuedDocumentsRequest
*/
func (a *DocumentsApiService) GetIssuedDocuments(ctx context.Context) ApiGetIssuedDocumentsRequest {
	return ApiGetIssuedDocumentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IssuedDocumentPaginatedList
func (a *DocumentsApiService) GetIssuedDocumentsExecute(r ApiGetIssuedDocumentsRequest) (*IssuedDocumentPaginatedList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IssuedDocumentPaginatedList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocumentsApiService.GetIssuedDocuments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/documents/issued"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.documentTypeId != nil {
		localVarQueryParams.Add("documentTypeId", parameterToString(*r.documentTypeId, ""))
	}
	if r.fromDateTime != nil {
		localVarQueryParams.Add("fromDateTime", parameterToString(*r.fromDateTime, ""))
	}
	if r.toDateTime != nil {
		localVarQueryParams.Add("toDateTime", parameterToString(*r.toDateTime, ""))
	}
	if r.pageNo != nil {
		localVarQueryParams.Add("pageNo", parameterToString(*r.pageNo, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
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

type ApiGetRegisteredDocumentTypesRequest struct {
	ctx context.Context
	ApiService *DocumentsApiService
	supportedEntityType *SupportedEntityType
	pageNo *int32
	pageSize *int32
}

// Supported entity type.
func (r ApiGetRegisteredDocumentTypesRequest) SupportedEntityType(supportedEntityType SupportedEntityType) ApiGetRegisteredDocumentTypesRequest {
	r.supportedEntityType = &supportedEntityType
	return r
}

// Page number.
func (r ApiGetRegisteredDocumentTypesRequest) PageNo(pageNo int32) ApiGetRegisteredDocumentTypesRequest {
	r.pageNo = &pageNo
	return r
}

// Number of items to return.
func (r ApiGetRegisteredDocumentTypesRequest) PageSize(pageSize int32) ApiGetRegisteredDocumentTypesRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetRegisteredDocumentTypesRequest) Execute() (*DocumentTypePaginatedList, *http.Response, error) {
	return r.ApiService.GetRegisteredDocumentTypesExecute(r)
}

/*
GetRegisteredDocumentTypes Get paginated list of registered document types.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRegisteredDocumentTypesRequest
*/
func (a *DocumentsApiService) GetRegisteredDocumentTypes(ctx context.Context) ApiGetRegisteredDocumentTypesRequest {
	return ApiGetRegisteredDocumentTypesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DocumentTypePaginatedList
func (a *DocumentsApiService) GetRegisteredDocumentTypesExecute(r ApiGetRegisteredDocumentTypesRequest) (*DocumentTypePaginatedList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DocumentTypePaginatedList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocumentsApiService.GetRegisteredDocumentTypes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/documents/types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.supportedEntityType != nil {
		localVarQueryParams.Add("supportedEntityType", parameterToString(*r.supportedEntityType, ""))
	}
	if r.pageNo != nil {
		localVarQueryParams.Add("pageNo", parameterToString(*r.pageNo, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
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

type ApiIssueDocumentToIndividualRequest struct {
	ctx context.Context
	ApiService *DocumentsApiService
	documentIssueRequest *DocumentIssueRequest
}

// Document issue request payload
func (r ApiIssueDocumentToIndividualRequest) DocumentIssueRequest(documentIssueRequest DocumentIssueRequest) ApiIssueDocumentToIndividualRequest {
	r.documentIssueRequest = &documentIssueRequest
	return r
}

func (r ApiIssueDocumentToIndividualRequest) Execute() (*DocumentIssueRequestDetails, *http.Response, error) {
	return r.ApiService.IssueDocumentToIndividualExecute(r)
}

/*
IssueDocumentToIndividual Issue a new document to an individual user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIssueDocumentToIndividualRequest
*/
func (a *DocumentsApiService) IssueDocumentToIndividual(ctx context.Context) ApiIssueDocumentToIndividualRequest {
	return ApiIssueDocumentToIndividualRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DocumentIssueRequestDetails
func (a *DocumentsApiService) IssueDocumentToIndividualExecute(r ApiIssueDocumentToIndividualRequest) (*DocumentIssueRequestDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DocumentIssueRequestDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocumentsApiService.IssueDocumentToIndividual")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/documents/issue/individual"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.documentIssueRequest == nil {
		return localVarReturnValue, nil, reportError("documentIssueRequest is required and must be specified")
	}

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
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
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

type ApiIssueDocumentToOrganizationRequest struct {
	ctx context.Context
	ApiService *DocumentsApiService
	documentIssueRequest *DocumentIssueRequest
}

// Document issue request payload
func (r ApiIssueDocumentToOrganizationRequest) DocumentIssueRequest(documentIssueRequest DocumentIssueRequest) ApiIssueDocumentToOrganizationRequest {
	r.documentIssueRequest = &documentIssueRequest
	return r
}

func (r ApiIssueDocumentToOrganizationRequest) Execute() (*DocumentIssueRequestDetails, *http.Response, error) {
	return r.ApiService.IssueDocumentToOrganizationExecute(r)
}

/*
IssueDocumentToOrganization Issue a new document to an organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIssueDocumentToOrganizationRequest
*/
func (a *DocumentsApiService) IssueDocumentToOrganization(ctx context.Context) ApiIssueDocumentToOrganizationRequest {
	return ApiIssueDocumentToOrganizationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DocumentIssueRequestDetails
func (a *DocumentsApiService) IssueDocumentToOrganizationExecute(r ApiIssueDocumentToOrganizationRequest) (*DocumentIssueRequestDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DocumentIssueRequestDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocumentsApiService.IssueDocumentToOrganization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/documents/issue/organization"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.documentIssueRequest == nil {
		return localVarReturnValue, nil, reportError("documentIssueRequest is required and must be specified")
	}

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
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
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

type ApiUploadDocumentForIndividualRequest struct {
	ctx context.Context
	ApiService *DocumentsApiService
	issueRequestId string
	formFile **os.File
}

func (r ApiUploadDocumentForIndividualRequest) FormFile(formFile *os.File) ApiUploadDocumentForIndividualRequest {
	r.formFile = &formFile
	return r
}

func (r ApiUploadDocumentForIndividualRequest) Execute() (*http.Response, error) {
	return r.ApiService.UploadDocumentForIndividualExecute(r)
}

/*
UploadDocumentForIndividual Upload a document for issuance request of individual.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param issueRequestId Document issue request id.
 @return ApiUploadDocumentForIndividualRequest
*/
func (a *DocumentsApiService) UploadDocumentForIndividual(ctx context.Context, issueRequestId string) ApiUploadDocumentForIndividualRequest {
	return ApiUploadDocumentForIndividualRequest{
		ApiService: a,
		ctx: ctx,
		issueRequestId: issueRequestId,
	}
}

// Execute executes the request
func (a *DocumentsApiService) UploadDocumentForIndividualExecute(r ApiUploadDocumentForIndividualRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocumentsApiService.UploadDocumentForIndividual")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/documents/issue/individual/upload/{issueRequestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"issueRequestId"+"}", url.PathEscape(parameterToString(r.issueRequestId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.formFile == nil {
		return nil, reportError("formFile is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var formFileLocalVarFormFileName string
	var formFileLocalVarFileName     string
	var formFileLocalVarFileBytes    []byte

	formFileLocalVarFormFileName = "formFile"

	formFileLocalVarFile := *r.formFile
	if formFileLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(formFileLocalVarFile)
		formFileLocalVarFileBytes = fbs
		formFileLocalVarFileName = formFileLocalVarFile.Name()
		formFileLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: formFileLocalVarFileBytes, fileName: formFileLocalVarFileName, formFileName: formFileLocalVarFormFileName})
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUploadDocumentForOrganizationRequest struct {
	ctx context.Context
	ApiService *DocumentsApiService
	issueRequestId string
	formFile **os.File
}

func (r ApiUploadDocumentForOrganizationRequest) FormFile(formFile *os.File) ApiUploadDocumentForOrganizationRequest {
	r.formFile = &formFile
	return r
}

func (r ApiUploadDocumentForOrganizationRequest) Execute() (*http.Response, error) {
	return r.ApiService.UploadDocumentForOrganizationExecute(r)
}

/*
UploadDocumentForOrganization Upload a document for issuance request of organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param issueRequestId Document issue request id System.Guid.
 @return ApiUploadDocumentForOrganizationRequest
*/
func (a *DocumentsApiService) UploadDocumentForOrganization(ctx context.Context, issueRequestId string) ApiUploadDocumentForOrganizationRequest {
	return ApiUploadDocumentForOrganizationRequest{
		ApiService: a,
		ctx: ctx,
		issueRequestId: issueRequestId,
	}
}

// Execute executes the request
func (a *DocumentsApiService) UploadDocumentForOrganizationExecute(r ApiUploadDocumentForOrganizationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocumentsApiService.UploadDocumentForOrganization")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/documents/issue/organization/upload/{issueRequestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"issueRequestId"+"}", url.PathEscape(parameterToString(r.issueRequestId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.formFile == nil {
		return nil, reportError("formFile is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var formFileLocalVarFormFileName string
	var formFileLocalVarFileName     string
	var formFileLocalVarFileBytes    []byte

	formFileLocalVarFormFileName = "formFile"

	formFileLocalVarFile := *r.formFile
	if formFileLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(formFileLocalVarFile)
		formFileLocalVarFileBytes = fbs
		formFileLocalVarFileName = formFileLocalVarFile.Name()
		formFileLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: formFileLocalVarFileBytes, fileName: formFileLocalVarFileName, formFileName: formFileLocalVarFormFileName})
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
