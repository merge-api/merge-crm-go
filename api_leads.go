/*
 * Merge CRM API
 *
 * The unified API for building rich integrations with multiple CRM platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_crm_client

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

// LeadsApiService LeadsApi service
type LeadsApiService service

type ApiLeadsCreateRequest struct {
	ctx _context.Context
	ApiService *LeadsApiService
	xAccountToken *string
	leadEndpointRequest *LeadEndpointRequest
	isDebugMode *bool
	runAsync *bool
}

func (r ApiLeadsCreateRequest) XAccountToken(xAccountToken string) ApiLeadsCreateRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiLeadsCreateRequest) LeadEndpointRequest(leadEndpointRequest LeadEndpointRequest) ApiLeadsCreateRequest {
	r.leadEndpointRequest = &leadEndpointRequest
	return r
}
func (r ApiLeadsCreateRequest) IsDebugMode(isDebugMode bool) ApiLeadsCreateRequest {
	r.isDebugMode = &isDebugMode
	return r
}
func (r ApiLeadsCreateRequest) RunAsync(runAsync bool) ApiLeadsCreateRequest {
	r.runAsync = &runAsync
	return r
}

func (r ApiLeadsCreateRequest) Execute() (LeadResponse, *_nethttp.Response, error) {
	return r.ApiService.LeadsCreateExecute(r)
}

/*
 * LeadsCreate Method for LeadsCreate
 * Creates a `Lead` object with the given values.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiLeadsCreateRequest
 */
func (a *LeadsApiService) LeadsCreate(ctx _context.Context) ApiLeadsCreateRequest {
	return ApiLeadsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return LeadResponse
 */
func (a *LeadsApiService) LeadsCreateExecute(r ApiLeadsCreateRequest) (LeadResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LeadResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LeadsApiService.LeadsCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/leads"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}
	if r.leadEndpointRequest == nil {
		return localVarReturnValue, nil, reportError("leadEndpointRequest is required and must be specified")
	}

	if r.isDebugMode != nil {
		localVarQueryParams.Add("is_debug_mode", parameterToString(*r.isDebugMode, ""))
	}
	if r.runAsync != nil {
		localVarQueryParams.Add("run_async", parameterToString(*r.runAsync, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

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
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	// body params
	localVarPostBody = r.leadEndpointRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
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

type ApiLeadsListRequest struct {
	ctx _context.Context
	ApiService *LeadsApiService
	xAccountToken *string
	convertedAccountId *string
	convertedContactId *string
	createdAfter *time.Time
	createdBefore *time.Time
	cursor *string
	includeDeletedData *bool
	includeRemoteData *bool
	modifiedAfter *time.Time
	modifiedBefore *time.Time
	ownerId *string
	pageSize *int32
	remoteId *string
}

func (r ApiLeadsListRequest) XAccountToken(xAccountToken string) ApiLeadsListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiLeadsListRequest) ConvertedAccountId(convertedAccountId string) ApiLeadsListRequest {
	r.convertedAccountId = &convertedAccountId
	return r
}
func (r ApiLeadsListRequest) ConvertedContactId(convertedContactId string) ApiLeadsListRequest {
	r.convertedContactId = &convertedContactId
	return r
}
func (r ApiLeadsListRequest) CreatedAfter(createdAfter time.Time) ApiLeadsListRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiLeadsListRequest) CreatedBefore(createdBefore time.Time) ApiLeadsListRequest {
	r.createdBefore = &createdBefore
	return r
}
func (r ApiLeadsListRequest) Cursor(cursor string) ApiLeadsListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiLeadsListRequest) IncludeDeletedData(includeDeletedData bool) ApiLeadsListRequest {
	r.includeDeletedData = &includeDeletedData
	return r
}
func (r ApiLeadsListRequest) IncludeRemoteData(includeRemoteData bool) ApiLeadsListRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiLeadsListRequest) ModifiedAfter(modifiedAfter time.Time) ApiLeadsListRequest {
	r.modifiedAfter = &modifiedAfter
	return r
}
func (r ApiLeadsListRequest) ModifiedBefore(modifiedBefore time.Time) ApiLeadsListRequest {
	r.modifiedBefore = &modifiedBefore
	return r
}
func (r ApiLeadsListRequest) OwnerId(ownerId string) ApiLeadsListRequest {
	r.ownerId = &ownerId
	return r
}
func (r ApiLeadsListRequest) PageSize(pageSize int32) ApiLeadsListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiLeadsListRequest) RemoteId(remoteId string) ApiLeadsListRequest {
	r.remoteId = &remoteId
	return r
}

func (r ApiLeadsListRequest) Execute() (PaginatedLeadList, *_nethttp.Response, error) {
	return r.ApiService.LeadsListExecute(r)
}

/*
 * LeadsList Method for LeadsList
 * Returns a list of `Lead` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiLeadsListRequest
 */
func (a *LeadsApiService) LeadsList(ctx _context.Context) ApiLeadsListRequest {
	return ApiLeadsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedLeadList
 */
func (a *LeadsApiService) LeadsListExecute(r ApiLeadsListRequest) (PaginatedLeadList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedLeadList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LeadsApiService.LeadsList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/leads"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.convertedAccountId != nil {
		localVarQueryParams.Add("converted_account_id", parameterToString(*r.convertedAccountId, ""))
	}
	if r.convertedContactId != nil {
		localVarQueryParams.Add("converted_contact_id", parameterToString(*r.convertedContactId, ""))
	}
	if r.createdAfter != nil {
		localVarQueryParams.Add("created_after", parameterToString(*r.createdAfter, ""))
	}
	if r.createdBefore != nil {
		localVarQueryParams.Add("created_before", parameterToString(*r.createdBefore, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.includeDeletedData != nil {
		localVarQueryParams.Add("include_deleted_data", parameterToString(*r.includeDeletedData, ""))
	}
	if r.includeRemoteData != nil {
		localVarQueryParams.Add("include_remote_data", parameterToString(*r.includeRemoteData, ""))
	}
	if r.modifiedAfter != nil {
		localVarQueryParams.Add("modified_after", parameterToString(*r.modifiedAfter, ""))
	}
	if r.modifiedBefore != nil {
		localVarQueryParams.Add("modified_before", parameterToString(*r.modifiedBefore, ""))
	}
	if r.ownerId != nil {
		localVarQueryParams.Add("owner_id", parameterToString(*r.ownerId, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.remoteId != nil {
		localVarQueryParams.Add("remote_id", parameterToString(*r.remoteId, ""))
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
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
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

type ApiLeadsMetaPostRetrieveRequest struct {
	ctx _context.Context
	ApiService *LeadsApiService
	xAccountToken *string
}

func (r ApiLeadsMetaPostRetrieveRequest) XAccountToken(xAccountToken string) ApiLeadsMetaPostRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}

func (r ApiLeadsMetaPostRetrieveRequest) Execute() (MetaResponse, *_nethttp.Response, error) {
	return r.ApiService.LeadsMetaPostRetrieveExecute(r)
}

/*
 * LeadsMetaPostRetrieve Method for LeadsMetaPostRetrieve
 * Returns metadata for `Lead` POSTs.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiLeadsMetaPostRetrieveRequest
 */
func (a *LeadsApiService) LeadsMetaPostRetrieve(ctx _context.Context) ApiLeadsMetaPostRetrieveRequest {
	return ApiLeadsMetaPostRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return MetaResponse
 */
func (a *LeadsApiService) LeadsMetaPostRetrieveExecute(r ApiLeadsMetaPostRetrieveRequest) (MetaResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MetaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LeadsApiService.LeadsMetaPostRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/leads/meta/post"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
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
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
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

type ApiLeadsRetrieveRequest struct {
	ctx _context.Context
	ApiService *LeadsApiService
	xAccountToken *string
	id string
	includeRemoteData *bool
}

func (r ApiLeadsRetrieveRequest) XAccountToken(xAccountToken string) ApiLeadsRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiLeadsRetrieveRequest) IncludeRemoteData(includeRemoteData bool) ApiLeadsRetrieveRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}

func (r ApiLeadsRetrieveRequest) Execute() (Lead, *_nethttp.Response, error) {
	return r.ApiService.LeadsRetrieveExecute(r)
}

/*
 * LeadsRetrieve Method for LeadsRetrieve
 * Returns a `Lead` object with the given `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiLeadsRetrieveRequest
 */
func (a *LeadsApiService) LeadsRetrieve(ctx _context.Context, id string) ApiLeadsRetrieveRequest {
	return ApiLeadsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return Lead
 */
func (a *LeadsApiService) LeadsRetrieveExecute(r ApiLeadsRetrieveRequest) (Lead, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Lead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LeadsApiService.LeadsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/leads/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.includeRemoteData != nil {
		localVarQueryParams.Add("include_remote_data", parameterToString(*r.includeRemoteData, ""))
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
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
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
