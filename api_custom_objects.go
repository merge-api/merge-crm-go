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

// CustomObjectsApiService CustomObjectsApi service
type CustomObjectsApiService service

type ApiCustomObjectClassesCustomObjectsCreateRequest struct {
	ctx _context.Context
	ApiService *CustomObjectsApiService
	xAccountToken *string
	customObjectClassId string
	cRMCustomObjectEndpointRequest *CRMCustomObjectEndpointRequest
	isDebugMode *bool
	runAsync *bool
}

func (r ApiCustomObjectClassesCustomObjectsCreateRequest) XAccountToken(xAccountToken string) ApiCustomObjectClassesCustomObjectsCreateRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiCustomObjectClassesCustomObjectsCreateRequest) CRMCustomObjectEndpointRequest(cRMCustomObjectEndpointRequest CRMCustomObjectEndpointRequest) ApiCustomObjectClassesCustomObjectsCreateRequest {
	r.cRMCustomObjectEndpointRequest = &cRMCustomObjectEndpointRequest
	return r
}
func (r ApiCustomObjectClassesCustomObjectsCreateRequest) IsDebugMode(isDebugMode bool) ApiCustomObjectClassesCustomObjectsCreateRequest {
	r.isDebugMode = &isDebugMode
	return r
}
func (r ApiCustomObjectClassesCustomObjectsCreateRequest) RunAsync(runAsync bool) ApiCustomObjectClassesCustomObjectsCreateRequest {
	r.runAsync = &runAsync
	return r
}

func (r ApiCustomObjectClassesCustomObjectsCreateRequest) Execute() (CRMCustomObjectResponse, *_nethttp.Response, error) {
	return r.ApiService.CustomObjectClassesCustomObjectsCreateExecute(r)
}

/*
 * CustomObjectClassesCustomObjectsCreate Method for CustomObjectClassesCustomObjectsCreate
 * Creates a `CustomObject` object with the given values.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customObjectClassId
 * @return ApiCustomObjectClassesCustomObjectsCreateRequest
 */
func (a *CustomObjectsApiService) CustomObjectClassesCustomObjectsCreate(ctx _context.Context, customObjectClassId string) ApiCustomObjectClassesCustomObjectsCreateRequest {
	return ApiCustomObjectClassesCustomObjectsCreateRequest{
		ApiService: a,
		ctx: ctx,
		customObjectClassId: customObjectClassId,
	}
}

/*
 * Execute executes the request
 * @return CRMCustomObjectResponse
 */
func (a *CustomObjectsApiService) CustomObjectClassesCustomObjectsCreateExecute(r ApiCustomObjectClassesCustomObjectsCreateRequest) (CRMCustomObjectResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CRMCustomObjectResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomObjectsApiService.CustomObjectClassesCustomObjectsCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom-object-classes/{custom_object_class_id}/custom-objects"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_object_class_id"+"}", _neturl.PathEscape(parameterToString(r.customObjectClassId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}
	if r.cRMCustomObjectEndpointRequest == nil {
		return localVarReturnValue, nil, reportError("cRMCustomObjectEndpointRequest is required and must be specified")
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
	localVarPostBody = r.cRMCustomObjectEndpointRequest
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

type ApiCustomObjectClassesCustomObjectsListRequest struct {
	ctx _context.Context
	ApiService *CustomObjectsApiService
	xAccountToken *string
	customObjectClassId string
	createdAfter *time.Time
	createdBefore *time.Time
	cursor *string
	includeDeletedData *bool
	includeRemoteData *bool
	includeRemoteFields *bool
	modifiedAfter *time.Time
	modifiedBefore *time.Time
	pageSize *int32
	remoteId *string
}

func (r ApiCustomObjectClassesCustomObjectsListRequest) XAccountToken(xAccountToken string) ApiCustomObjectClassesCustomObjectsListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiCustomObjectClassesCustomObjectsListRequest) CreatedAfter(createdAfter time.Time) ApiCustomObjectClassesCustomObjectsListRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiCustomObjectClassesCustomObjectsListRequest) CreatedBefore(createdBefore time.Time) ApiCustomObjectClassesCustomObjectsListRequest {
	r.createdBefore = &createdBefore
	return r
}
func (r ApiCustomObjectClassesCustomObjectsListRequest) Cursor(cursor string) ApiCustomObjectClassesCustomObjectsListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiCustomObjectClassesCustomObjectsListRequest) IncludeDeletedData(includeDeletedData bool) ApiCustomObjectClassesCustomObjectsListRequest {
	r.includeDeletedData = &includeDeletedData
	return r
}
func (r ApiCustomObjectClassesCustomObjectsListRequest) IncludeRemoteData(includeRemoteData bool) ApiCustomObjectClassesCustomObjectsListRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiCustomObjectClassesCustomObjectsListRequest) IncludeRemoteFields(includeRemoteFields bool) ApiCustomObjectClassesCustomObjectsListRequest {
	r.includeRemoteFields = &includeRemoteFields
	return r
}
func (r ApiCustomObjectClassesCustomObjectsListRequest) ModifiedAfter(modifiedAfter time.Time) ApiCustomObjectClassesCustomObjectsListRequest {
	r.modifiedAfter = &modifiedAfter
	return r
}
func (r ApiCustomObjectClassesCustomObjectsListRequest) ModifiedBefore(modifiedBefore time.Time) ApiCustomObjectClassesCustomObjectsListRequest {
	r.modifiedBefore = &modifiedBefore
	return r
}
func (r ApiCustomObjectClassesCustomObjectsListRequest) PageSize(pageSize int32) ApiCustomObjectClassesCustomObjectsListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiCustomObjectClassesCustomObjectsListRequest) RemoteId(remoteId string) ApiCustomObjectClassesCustomObjectsListRequest {
	r.remoteId = &remoteId
	return r
}

func (r ApiCustomObjectClassesCustomObjectsListRequest) Execute() (PaginatedCustomObjectList, *_nethttp.Response, error) {
	return r.ApiService.CustomObjectClassesCustomObjectsListExecute(r)
}

/*
 * CustomObjectClassesCustomObjectsList Method for CustomObjectClassesCustomObjectsList
 * Returns a list of `CustomObject` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customObjectClassId
 * @return ApiCustomObjectClassesCustomObjectsListRequest
 */
func (a *CustomObjectsApiService) CustomObjectClassesCustomObjectsList(ctx _context.Context, customObjectClassId string) ApiCustomObjectClassesCustomObjectsListRequest {
	return ApiCustomObjectClassesCustomObjectsListRequest{
		ApiService: a,
		ctx: ctx,
		customObjectClassId: customObjectClassId,
	}
}

/*
 * Execute executes the request
 * @return PaginatedCustomObjectList
 */
func (a *CustomObjectsApiService) CustomObjectClassesCustomObjectsListExecute(r ApiCustomObjectClassesCustomObjectsListRequest) (PaginatedCustomObjectList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedCustomObjectList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomObjectsApiService.CustomObjectClassesCustomObjectsList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom-object-classes/{custom_object_class_id}/custom-objects"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_object_class_id"+"}", _neturl.PathEscape(parameterToString(r.customObjectClassId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
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
	if r.includeRemoteFields != nil {
		localVarQueryParams.Add("include_remote_fields", parameterToString(*r.includeRemoteFields, ""))
	}
	if r.modifiedAfter != nil {
		localVarQueryParams.Add("modified_after", parameterToString(*r.modifiedAfter, ""))
	}
	if r.modifiedBefore != nil {
		localVarQueryParams.Add("modified_before", parameterToString(*r.modifiedBefore, ""))
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

type ApiCustomObjectClassesCustomObjectsMetaPatchRetrieveRequest struct {
	ctx _context.Context
	ApiService *CustomObjectsApiService
	xAccountToken *string
	customObjectClassId string
	id string
}

func (r ApiCustomObjectClassesCustomObjectsMetaPatchRetrieveRequest) XAccountToken(xAccountToken string) ApiCustomObjectClassesCustomObjectsMetaPatchRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}

func (r ApiCustomObjectClassesCustomObjectsMetaPatchRetrieveRequest) Execute() (MetaResponse, *_nethttp.Response, error) {
	return r.ApiService.CustomObjectClassesCustomObjectsMetaPatchRetrieveExecute(r)
}

/*
 * CustomObjectClassesCustomObjectsMetaPatchRetrieve Method for CustomObjectClassesCustomObjectsMetaPatchRetrieve
 * Returns metadata for `CRMCustomObject` PATCHs.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customObjectClassId
 * @param id
 * @return ApiCustomObjectClassesCustomObjectsMetaPatchRetrieveRequest
 */
func (a *CustomObjectsApiService) CustomObjectClassesCustomObjectsMetaPatchRetrieve(ctx _context.Context, customObjectClassId string, id string) ApiCustomObjectClassesCustomObjectsMetaPatchRetrieveRequest {
	return ApiCustomObjectClassesCustomObjectsMetaPatchRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		customObjectClassId: customObjectClassId,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return MetaResponse
 */
func (a *CustomObjectsApiService) CustomObjectClassesCustomObjectsMetaPatchRetrieveExecute(r ApiCustomObjectClassesCustomObjectsMetaPatchRetrieveRequest) (MetaResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MetaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomObjectsApiService.CustomObjectClassesCustomObjectsMetaPatchRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom-object-classes/{custom_object_class_id}/custom-objects/meta/patch/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_object_class_id"+"}", _neturl.PathEscape(parameterToString(r.customObjectClassId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiCustomObjectClassesCustomObjectsMetaPostRetrieveRequest struct {
	ctx _context.Context
	ApiService *CustomObjectsApiService
	xAccountToken *string
	customObjectClassId string
}

func (r ApiCustomObjectClassesCustomObjectsMetaPostRetrieveRequest) XAccountToken(xAccountToken string) ApiCustomObjectClassesCustomObjectsMetaPostRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}

func (r ApiCustomObjectClassesCustomObjectsMetaPostRetrieveRequest) Execute() (MetaResponse, *_nethttp.Response, error) {
	return r.ApiService.CustomObjectClassesCustomObjectsMetaPostRetrieveExecute(r)
}

/*
 * CustomObjectClassesCustomObjectsMetaPostRetrieve Method for CustomObjectClassesCustomObjectsMetaPostRetrieve
 * Returns metadata for `CRMCustomObject` POSTs.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customObjectClassId
 * @return ApiCustomObjectClassesCustomObjectsMetaPostRetrieveRequest
 */
func (a *CustomObjectsApiService) CustomObjectClassesCustomObjectsMetaPostRetrieve(ctx _context.Context, customObjectClassId string) ApiCustomObjectClassesCustomObjectsMetaPostRetrieveRequest {
	return ApiCustomObjectClassesCustomObjectsMetaPostRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		customObjectClassId: customObjectClassId,
	}
}

/*
 * Execute executes the request
 * @return MetaResponse
 */
func (a *CustomObjectsApiService) CustomObjectClassesCustomObjectsMetaPostRetrieveExecute(r ApiCustomObjectClassesCustomObjectsMetaPostRetrieveRequest) (MetaResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MetaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomObjectsApiService.CustomObjectClassesCustomObjectsMetaPostRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom-object-classes/{custom_object_class_id}/custom-objects/meta/post"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_object_class_id"+"}", _neturl.PathEscape(parameterToString(r.customObjectClassId, "")), -1)

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

type ApiCustomObjectClassesCustomObjectsPartialUpdateRequest struct {
	ctx _context.Context
	ApiService *CustomObjectsApiService
	xAccountToken *string
	customObjectClassId string
	id string
	patchedCRMCustomObjectEndpointRequest *PatchedCRMCustomObjectEndpointRequest
	isDebugMode *bool
	runAsync *bool
}

func (r ApiCustomObjectClassesCustomObjectsPartialUpdateRequest) XAccountToken(xAccountToken string) ApiCustomObjectClassesCustomObjectsPartialUpdateRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiCustomObjectClassesCustomObjectsPartialUpdateRequest) PatchedCRMCustomObjectEndpointRequest(patchedCRMCustomObjectEndpointRequest PatchedCRMCustomObjectEndpointRequest) ApiCustomObjectClassesCustomObjectsPartialUpdateRequest {
	r.patchedCRMCustomObjectEndpointRequest = &patchedCRMCustomObjectEndpointRequest
	return r
}
func (r ApiCustomObjectClassesCustomObjectsPartialUpdateRequest) IsDebugMode(isDebugMode bool) ApiCustomObjectClassesCustomObjectsPartialUpdateRequest {
	r.isDebugMode = &isDebugMode
	return r
}
func (r ApiCustomObjectClassesCustomObjectsPartialUpdateRequest) RunAsync(runAsync bool) ApiCustomObjectClassesCustomObjectsPartialUpdateRequest {
	r.runAsync = &runAsync
	return r
}

func (r ApiCustomObjectClassesCustomObjectsPartialUpdateRequest) Execute() (CRMCustomObjectResponse, *_nethttp.Response, error) {
	return r.ApiService.CustomObjectClassesCustomObjectsPartialUpdateExecute(r)
}

/*
 * CustomObjectClassesCustomObjectsPartialUpdate Method for CustomObjectClassesCustomObjectsPartialUpdate
 * Updates a `CustomObject` object with the given `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customObjectClassId
 * @param id
 * @return ApiCustomObjectClassesCustomObjectsPartialUpdateRequest
 */
func (a *CustomObjectsApiService) CustomObjectClassesCustomObjectsPartialUpdate(ctx _context.Context, customObjectClassId string, id string) ApiCustomObjectClassesCustomObjectsPartialUpdateRequest {
	return ApiCustomObjectClassesCustomObjectsPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		customObjectClassId: customObjectClassId,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return CRMCustomObjectResponse
 */
func (a *CustomObjectsApiService) CustomObjectClassesCustomObjectsPartialUpdateExecute(r ApiCustomObjectClassesCustomObjectsPartialUpdateRequest) (CRMCustomObjectResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CRMCustomObjectResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomObjectsApiService.CustomObjectClassesCustomObjectsPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom-object-classes/{custom_object_class_id}/custom-objects/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_object_class_id"+"}", _neturl.PathEscape(parameterToString(r.customObjectClassId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}
	if r.patchedCRMCustomObjectEndpointRequest == nil {
		return localVarReturnValue, nil, reportError("patchedCRMCustomObjectEndpointRequest is required and must be specified")
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
	localVarPostBody = r.patchedCRMCustomObjectEndpointRequest
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

type ApiCustomObjectClassesCustomObjectsRetrieveRequest struct {
	ctx _context.Context
	ApiService *CustomObjectsApiService
	xAccountToken *string
	customObjectClassId string
	id string
	includeRemoteData *bool
	includeRemoteFields *bool
}

func (r ApiCustomObjectClassesCustomObjectsRetrieveRequest) XAccountToken(xAccountToken string) ApiCustomObjectClassesCustomObjectsRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiCustomObjectClassesCustomObjectsRetrieveRequest) IncludeRemoteData(includeRemoteData bool) ApiCustomObjectClassesCustomObjectsRetrieveRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiCustomObjectClassesCustomObjectsRetrieveRequest) IncludeRemoteFields(includeRemoteFields bool) ApiCustomObjectClassesCustomObjectsRetrieveRequest {
	r.includeRemoteFields = &includeRemoteFields
	return r
}

func (r ApiCustomObjectClassesCustomObjectsRetrieveRequest) Execute() (CustomObject, *_nethttp.Response, error) {
	return r.ApiService.CustomObjectClassesCustomObjectsRetrieveExecute(r)
}

/*
 * CustomObjectClassesCustomObjectsRetrieve Method for CustomObjectClassesCustomObjectsRetrieve
 * Returns a `CustomObject` object with the given `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customObjectClassId
 * @param id
 * @return ApiCustomObjectClassesCustomObjectsRetrieveRequest
 */
func (a *CustomObjectsApiService) CustomObjectClassesCustomObjectsRetrieve(ctx _context.Context, customObjectClassId string, id string) ApiCustomObjectClassesCustomObjectsRetrieveRequest {
	return ApiCustomObjectClassesCustomObjectsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		customObjectClassId: customObjectClassId,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return CustomObject
 */
func (a *CustomObjectsApiService) CustomObjectClassesCustomObjectsRetrieveExecute(r ApiCustomObjectClassesCustomObjectsRetrieveRequest) (CustomObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CustomObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomObjectsApiService.CustomObjectClassesCustomObjectsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom-object-classes/{custom_object_class_id}/custom-objects/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_object_class_id"+"}", _neturl.PathEscape(parameterToString(r.customObjectClassId, "")), -1)
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
	if r.includeRemoteFields != nil {
		localVarQueryParams.Add("include_remote_fields", parameterToString(*r.includeRemoteFields, ""))
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