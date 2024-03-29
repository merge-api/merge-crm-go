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

// NotesApiService NotesApi service
type NotesApiService service

type ApiNotesCreateRequest struct {
	ctx _context.Context
	ApiService *NotesApiService
	xAccountToken *string
	noteEndpointRequest *NoteEndpointRequest
	isDebugMode *bool
	runAsync *bool
}

func (r ApiNotesCreateRequest) XAccountToken(xAccountToken string) ApiNotesCreateRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiNotesCreateRequest) NoteEndpointRequest(noteEndpointRequest NoteEndpointRequest) ApiNotesCreateRequest {
	r.noteEndpointRequest = &noteEndpointRequest
	return r
}
func (r ApiNotesCreateRequest) IsDebugMode(isDebugMode bool) ApiNotesCreateRequest {
	r.isDebugMode = &isDebugMode
	return r
}
func (r ApiNotesCreateRequest) RunAsync(runAsync bool) ApiNotesCreateRequest {
	r.runAsync = &runAsync
	return r
}

func (r ApiNotesCreateRequest) Execute() (NoteResponse, *_nethttp.Response, error) {
	return r.ApiService.NotesCreateExecute(r)
}

/*
 * NotesCreate Method for NotesCreate
 * Creates a `Note` object with the given values.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiNotesCreateRequest
 */
func (a *NotesApiService) NotesCreate(ctx _context.Context) ApiNotesCreateRequest {
	return ApiNotesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return NoteResponse
 */
func (a *NotesApiService) NotesCreateExecute(r ApiNotesCreateRequest) (NoteResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotesApiService.NotesCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}
	if r.noteEndpointRequest == nil {
		return localVarReturnValue, nil, reportError("noteEndpointRequest is required and must be specified")
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
	localVarPostBody = r.noteEndpointRequest
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

type ApiNotesListRequest struct {
	ctx _context.Context
	ApiService *NotesApiService
	xAccountToken *string
	accountId *string
	contactId *string
	createdAfter *time.Time
	createdBefore *time.Time
	cursor *string
	includeDeletedData *bool
	includeRemoteData *bool
	includeRemoteFields *bool
	modifiedAfter *time.Time
	modifiedBefore *time.Time
	opportunityId *string
	ownerId *string
	pageSize *int32
	remoteId *string
}

func (r ApiNotesListRequest) XAccountToken(xAccountToken string) ApiNotesListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiNotesListRequest) AccountId(accountId string) ApiNotesListRequest {
	r.accountId = &accountId
	return r
}
func (r ApiNotesListRequest) ContactId(contactId string) ApiNotesListRequest {
	r.contactId = &contactId
	return r
}
func (r ApiNotesListRequest) CreatedAfter(createdAfter time.Time) ApiNotesListRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiNotesListRequest) CreatedBefore(createdBefore time.Time) ApiNotesListRequest {
	r.createdBefore = &createdBefore
	return r
}
func (r ApiNotesListRequest) Cursor(cursor string) ApiNotesListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiNotesListRequest) IncludeDeletedData(includeDeletedData bool) ApiNotesListRequest {
	r.includeDeletedData = &includeDeletedData
	return r
}
func (r ApiNotesListRequest) IncludeRemoteData(includeRemoteData bool) ApiNotesListRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiNotesListRequest) IncludeRemoteFields(includeRemoteFields bool) ApiNotesListRequest {
	r.includeRemoteFields = &includeRemoteFields
	return r
}
func (r ApiNotesListRequest) ModifiedAfter(modifiedAfter time.Time) ApiNotesListRequest {
	r.modifiedAfter = &modifiedAfter
	return r
}
func (r ApiNotesListRequest) ModifiedBefore(modifiedBefore time.Time) ApiNotesListRequest {
	r.modifiedBefore = &modifiedBefore
	return r
}
func (r ApiNotesListRequest) OpportunityId(opportunityId string) ApiNotesListRequest {
	r.opportunityId = &opportunityId
	return r
}
func (r ApiNotesListRequest) OwnerId(ownerId string) ApiNotesListRequest {
	r.ownerId = &ownerId
	return r
}
func (r ApiNotesListRequest) PageSize(pageSize int32) ApiNotesListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiNotesListRequest) RemoteId(remoteId string) ApiNotesListRequest {
	r.remoteId = &remoteId
	return r
}

func (r ApiNotesListRequest) Execute() (PaginatedNoteList, *_nethttp.Response, error) {
	return r.ApiService.NotesListExecute(r)
}

/*
 * NotesList Method for NotesList
 * Returns a list of `Note` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiNotesListRequest
 */
func (a *NotesApiService) NotesList(ctx _context.Context) ApiNotesListRequest {
	return ApiNotesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedNoteList
 */
func (a *NotesApiService) NotesListExecute(r ApiNotesListRequest) (PaginatedNoteList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedNoteList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotesApiService.NotesList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.accountId != nil {
		localVarQueryParams.Add("account_id", parameterToString(*r.accountId, ""))
	}
	if r.contactId != nil {
		localVarQueryParams.Add("contact_id", parameterToString(*r.contactId, ""))
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
	if r.opportunityId != nil {
		localVarQueryParams.Add("opportunity_id", parameterToString(*r.opportunityId, ""))
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

type ApiNotesMetaPostRetrieveRequest struct {
	ctx _context.Context
	ApiService *NotesApiService
	xAccountToken *string
}

func (r ApiNotesMetaPostRetrieveRequest) XAccountToken(xAccountToken string) ApiNotesMetaPostRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}

func (r ApiNotesMetaPostRetrieveRequest) Execute() (MetaResponse, *_nethttp.Response, error) {
	return r.ApiService.NotesMetaPostRetrieveExecute(r)
}

/*
 * NotesMetaPostRetrieve Method for NotesMetaPostRetrieve
 * Returns metadata for `Note` POSTs.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiNotesMetaPostRetrieveRequest
 */
func (a *NotesApiService) NotesMetaPostRetrieve(ctx _context.Context) ApiNotesMetaPostRetrieveRequest {
	return ApiNotesMetaPostRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return MetaResponse
 */
func (a *NotesApiService) NotesMetaPostRetrieveExecute(r ApiNotesMetaPostRetrieveRequest) (MetaResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MetaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotesApiService.NotesMetaPostRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notes/meta/post"

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

type ApiNotesRemoteFieldClassesListRequest struct {
	ctx _context.Context
	ApiService *NotesApiService
	xAccountToken *string
	cursor *string
	includeDeletedData *bool
	includeRemoteData *bool
	includeRemoteFields *bool
	pageSize *int32
}

func (r ApiNotesRemoteFieldClassesListRequest) XAccountToken(xAccountToken string) ApiNotesRemoteFieldClassesListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiNotesRemoteFieldClassesListRequest) Cursor(cursor string) ApiNotesRemoteFieldClassesListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiNotesRemoteFieldClassesListRequest) IncludeDeletedData(includeDeletedData bool) ApiNotesRemoteFieldClassesListRequest {
	r.includeDeletedData = &includeDeletedData
	return r
}
func (r ApiNotesRemoteFieldClassesListRequest) IncludeRemoteData(includeRemoteData bool) ApiNotesRemoteFieldClassesListRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiNotesRemoteFieldClassesListRequest) IncludeRemoteFields(includeRemoteFields bool) ApiNotesRemoteFieldClassesListRequest {
	r.includeRemoteFields = &includeRemoteFields
	return r
}
func (r ApiNotesRemoteFieldClassesListRequest) PageSize(pageSize int32) ApiNotesRemoteFieldClassesListRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiNotesRemoteFieldClassesListRequest) Execute() (PaginatedRemoteFieldClassList, *_nethttp.Response, error) {
	return r.ApiService.NotesRemoteFieldClassesListExecute(r)
}

/*
 * NotesRemoteFieldClassesList Method for NotesRemoteFieldClassesList
 * Returns a list of `RemoteFieldClass` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiNotesRemoteFieldClassesListRequest
 */
func (a *NotesApiService) NotesRemoteFieldClassesList(ctx _context.Context) ApiNotesRemoteFieldClassesListRequest {
	return ApiNotesRemoteFieldClassesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedRemoteFieldClassList
 */
func (a *NotesApiService) NotesRemoteFieldClassesListExecute(r ApiNotesRemoteFieldClassesListRequest) (PaginatedRemoteFieldClassList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedRemoteFieldClassList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotesApiService.NotesRemoteFieldClassesList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notes/remote-field-classes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
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
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
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

type ApiNotesRetrieveRequest struct {
	ctx _context.Context
	ApiService *NotesApiService
	xAccountToken *string
	id string
	includeRemoteData *bool
	includeRemoteFields *bool
}

func (r ApiNotesRetrieveRequest) XAccountToken(xAccountToken string) ApiNotesRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiNotesRetrieveRequest) IncludeRemoteData(includeRemoteData bool) ApiNotesRetrieveRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiNotesRetrieveRequest) IncludeRemoteFields(includeRemoteFields bool) ApiNotesRetrieveRequest {
	r.includeRemoteFields = &includeRemoteFields
	return r
}

func (r ApiNotesRetrieveRequest) Execute() (Note, *_nethttp.Response, error) {
	return r.ApiService.NotesRetrieveExecute(r)
}

/*
 * NotesRetrieve Method for NotesRetrieve
 * Returns a `Note` object with the given `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiNotesRetrieveRequest
 */
func (a *NotesApiService) NotesRetrieve(ctx _context.Context, id string) ApiNotesRetrieveRequest {
	return ApiNotesRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return Note
 */
func (a *NotesApiService) NotesRetrieveExecute(r ApiNotesRetrieveRequest) (Note, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Note
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotesApiService.NotesRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notes/{id}"
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
