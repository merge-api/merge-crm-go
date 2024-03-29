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

// AssociationsApiService AssociationsApi service
type AssociationsApiService service

type ApiCustomObjectClassesCustomObjectsAssociationsListRequest struct {
	ctx _context.Context
	ApiService *AssociationsApiService
	xAccountToken *string
	customObjectClassId string
	objectId string
	associationTypeId *string
	createdAfter *time.Time
	createdBefore *time.Time
	cursor *string
	includeDeletedData *bool
	includeRemoteData *bool
	modifiedAfter *time.Time
	modifiedBefore *time.Time
	pageSize *int32
	remoteId *string
}

func (r ApiCustomObjectClassesCustomObjectsAssociationsListRequest) XAccountToken(xAccountToken string) ApiCustomObjectClassesCustomObjectsAssociationsListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiCustomObjectClassesCustomObjectsAssociationsListRequest) AssociationTypeId(associationTypeId string) ApiCustomObjectClassesCustomObjectsAssociationsListRequest {
	r.associationTypeId = &associationTypeId
	return r
}
func (r ApiCustomObjectClassesCustomObjectsAssociationsListRequest) CreatedAfter(createdAfter time.Time) ApiCustomObjectClassesCustomObjectsAssociationsListRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiCustomObjectClassesCustomObjectsAssociationsListRequest) CreatedBefore(createdBefore time.Time) ApiCustomObjectClassesCustomObjectsAssociationsListRequest {
	r.createdBefore = &createdBefore
	return r
}
func (r ApiCustomObjectClassesCustomObjectsAssociationsListRequest) Cursor(cursor string) ApiCustomObjectClassesCustomObjectsAssociationsListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiCustomObjectClassesCustomObjectsAssociationsListRequest) IncludeDeletedData(includeDeletedData bool) ApiCustomObjectClassesCustomObjectsAssociationsListRequest {
	r.includeDeletedData = &includeDeletedData
	return r
}
func (r ApiCustomObjectClassesCustomObjectsAssociationsListRequest) IncludeRemoteData(includeRemoteData bool) ApiCustomObjectClassesCustomObjectsAssociationsListRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiCustomObjectClassesCustomObjectsAssociationsListRequest) ModifiedAfter(modifiedAfter time.Time) ApiCustomObjectClassesCustomObjectsAssociationsListRequest {
	r.modifiedAfter = &modifiedAfter
	return r
}
func (r ApiCustomObjectClassesCustomObjectsAssociationsListRequest) ModifiedBefore(modifiedBefore time.Time) ApiCustomObjectClassesCustomObjectsAssociationsListRequest {
	r.modifiedBefore = &modifiedBefore
	return r
}
func (r ApiCustomObjectClassesCustomObjectsAssociationsListRequest) PageSize(pageSize int32) ApiCustomObjectClassesCustomObjectsAssociationsListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiCustomObjectClassesCustomObjectsAssociationsListRequest) RemoteId(remoteId string) ApiCustomObjectClassesCustomObjectsAssociationsListRequest {
	r.remoteId = &remoteId
	return r
}

func (r ApiCustomObjectClassesCustomObjectsAssociationsListRequest) Execute() (PaginatedAssociationList, *_nethttp.Response, error) {
	return r.ApiService.CustomObjectClassesCustomObjectsAssociationsListExecute(r)
}

/*
 * CustomObjectClassesCustomObjectsAssociationsList Method for CustomObjectClassesCustomObjectsAssociationsList
 * Returns a list of `Association` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customObjectClassId
 * @param objectId
 * @return ApiCustomObjectClassesCustomObjectsAssociationsListRequest
 */
func (a *AssociationsApiService) CustomObjectClassesCustomObjectsAssociationsList(ctx _context.Context, customObjectClassId string, objectId string) ApiCustomObjectClassesCustomObjectsAssociationsListRequest {
	return ApiCustomObjectClassesCustomObjectsAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		customObjectClassId: customObjectClassId,
		objectId: objectId,
	}
}

/*
 * Execute executes the request
 * @return PaginatedAssociationList
 */
func (a *AssociationsApiService) CustomObjectClassesCustomObjectsAssociationsListExecute(r ApiCustomObjectClassesCustomObjectsAssociationsListRequest) (PaginatedAssociationList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedAssociationList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.CustomObjectClassesCustomObjectsAssociationsList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom-object-classes/{custom_object_class_id}/custom-objects/{object_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_object_class_id"+"}", _neturl.PathEscape(parameterToString(r.customObjectClassId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"object_id"+"}", _neturl.PathEscape(parameterToString(r.objectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.associationTypeId != nil {
		localVarQueryParams.Add("association_type_id", parameterToString(*r.associationTypeId, ""))
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

type ApiCustomObjectClassesCustomObjectsAssociationsUpdateRequest struct {
	ctx _context.Context
	ApiService *AssociationsApiService
	xAccountToken *string
	associationTypeId string
	sourceClassId string
	sourceObjectId string
	targetClassId string
	targetObjectId string
	isDebugMode *bool
	runAsync *bool
}

func (r ApiCustomObjectClassesCustomObjectsAssociationsUpdateRequest) XAccountToken(xAccountToken string) ApiCustomObjectClassesCustomObjectsAssociationsUpdateRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiCustomObjectClassesCustomObjectsAssociationsUpdateRequest) IsDebugMode(isDebugMode bool) ApiCustomObjectClassesCustomObjectsAssociationsUpdateRequest {
	r.isDebugMode = &isDebugMode
	return r
}
func (r ApiCustomObjectClassesCustomObjectsAssociationsUpdateRequest) RunAsync(runAsync bool) ApiCustomObjectClassesCustomObjectsAssociationsUpdateRequest {
	r.runAsync = &runAsync
	return r
}

func (r ApiCustomObjectClassesCustomObjectsAssociationsUpdateRequest) Execute() (Association, *_nethttp.Response, error) {
	return r.ApiService.CustomObjectClassesCustomObjectsAssociationsUpdateExecute(r)
}

/*
 * CustomObjectClassesCustomObjectsAssociationsUpdate Method for CustomObjectClassesCustomObjectsAssociationsUpdate
 * Creates an Association between `source_object_id` and `target_object_id` of type `association_type_id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param associationTypeId
 * @param sourceClassId
 * @param sourceObjectId
 * @param targetClassId
 * @param targetObjectId
 * @return ApiCustomObjectClassesCustomObjectsAssociationsUpdateRequest
 */
func (a *AssociationsApiService) CustomObjectClassesCustomObjectsAssociationsUpdate(ctx _context.Context, associationTypeId string, sourceClassId string, sourceObjectId string, targetClassId string, targetObjectId string) ApiCustomObjectClassesCustomObjectsAssociationsUpdateRequest {
	return ApiCustomObjectClassesCustomObjectsAssociationsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		associationTypeId: associationTypeId,
		sourceClassId: sourceClassId,
		sourceObjectId: sourceObjectId,
		targetClassId: targetClassId,
		targetObjectId: targetObjectId,
	}
}

/*
 * Execute executes the request
 * @return Association
 */
func (a *AssociationsApiService) CustomObjectClassesCustomObjectsAssociationsUpdateExecute(r ApiCustomObjectClassesCustomObjectsAssociationsUpdateRequest) (Association, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Association
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.CustomObjectClassesCustomObjectsAssociationsUpdate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom-object-classes/{source_class_id}/custom-objects/{source_object_id}/associations/{target_class_id}/{target_object_id}/{association_type_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"association_type_id"+"}", _neturl.PathEscape(parameterToString(r.associationTypeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"source_class_id"+"}", _neturl.PathEscape(parameterToString(r.sourceClassId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"source_object_id"+"}", _neturl.PathEscape(parameterToString(r.sourceObjectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"target_class_id"+"}", _neturl.PathEscape(parameterToString(r.targetClassId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"target_object_id"+"}", _neturl.PathEscape(parameterToString(r.targetObjectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.isDebugMode != nil {
		localVarQueryParams.Add("is_debug_mode", parameterToString(*r.isDebugMode, ""))
	}
	if r.runAsync != nil {
		localVarQueryParams.Add("run_async", parameterToString(*r.runAsync, ""))
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
