# \CustomObjectsApi

All URIs are relative to *https://api.merge.dev/api/crm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomObjectClassesCustomObjectsCreate**](CustomObjectsApi.md#CustomObjectClassesCustomObjectsCreate) | **Post** /custom-object-classes/{custom_object_class_id}/custom-objects | 
[**CustomObjectClassesCustomObjectsList**](CustomObjectsApi.md#CustomObjectClassesCustomObjectsList) | **Get** /custom-object-classes/{custom_object_class_id}/custom-objects | 
[**CustomObjectClassesCustomObjectsMetaPatchRetrieve**](CustomObjectsApi.md#CustomObjectClassesCustomObjectsMetaPatchRetrieve) | **Get** /custom-object-classes/{custom_object_class_id}/custom-objects/meta/patch/{id} | 
[**CustomObjectClassesCustomObjectsMetaPostRetrieve**](CustomObjectsApi.md#CustomObjectClassesCustomObjectsMetaPostRetrieve) | **Get** /custom-object-classes/{custom_object_class_id}/custom-objects/meta/post | 
[**CustomObjectClassesCustomObjectsPartialUpdate**](CustomObjectsApi.md#CustomObjectClassesCustomObjectsPartialUpdate) | **Patch** /custom-object-classes/{custom_object_class_id}/custom-objects/{id} | 
[**CustomObjectClassesCustomObjectsRetrieve**](CustomObjectsApi.md#CustomObjectClassesCustomObjectsRetrieve) | **Get** /custom-object-classes/{custom_object_class_id}/custom-objects/{id} | 



## CustomObjectClassesCustomObjectsCreate

> CRMCustomObjectResponse CustomObjectClassesCustomObjectsCreate(ctx, customObjectClassId).XAccountToken(xAccountToken).CRMCustomObjectEndpointRequest(cRMCustomObjectEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()





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
    xAccountToken := "xAccountToken_example" // string | Token identifying the end user.
    customObjectClassId := TODO // string | 
    cRMCustomObjectEndpointRequest := *openapiclient.NewCRMCustomObjectEndpointRequest(*openapiclient.NewCustomObjectRequest(map[string]interface{}{"key": interface{}(123)})) // CRMCustomObjectEndpointRequest | 
    isDebugMode := true // bool | Whether to include debug fields (such as log file links) in the response. (optional)
    runAsync := true // bool | Whether or not third-party updates should be run asynchronously. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomObjectsApi.CustomObjectClassesCustomObjectsCreate(context.Background(), customObjectClassId).XAccountToken(xAccountToken).CRMCustomObjectEndpointRequest(cRMCustomObjectEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.CustomObjectClassesCustomObjectsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomObjectClassesCustomObjectsCreate`: CRMCustomObjectResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomObjectsApi.CustomObjectClassesCustomObjectsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customObjectClassId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomObjectClassesCustomObjectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **cRMCustomObjectEndpointRequest** | [**CRMCustomObjectEndpointRequest**](CRMCustomObjectEndpointRequest.md) |  | 
 **isDebugMode** | **bool** | Whether to include debug fields (such as log file links) in the response. | 
 **runAsync** | **bool** | Whether or not third-party updates should be run asynchronously. | 

### Return type

[**CRMCustomObjectResponse**](CRMCustomObjectResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomObjectClassesCustomObjectsList

> PaginatedCustomObjectList CustomObjectClassesCustomObjectsList(ctx, customObjectClassId).XAccountToken(xAccountToken).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).IncludeRemoteFields(includeRemoteFields).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    xAccountToken := "xAccountToken_example" // string | Token identifying the end user.
    customObjectClassId := TODO // string | 
    createdAfter := time.Now() // time.Time | If provided, will only return objects created after this datetime. (optional)
    createdBefore := time.Now() // time.Time | If provided, will only return objects created before this datetime. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    includeDeletedData := true // bool | Whether to include data that was marked as deleted by third party webhooks. (optional)
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    includeRemoteFields := true // bool | Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format. (optional)
    modifiedAfter := time.Now() // time.Time | If provided, only objects synced by Merge after this date time will be returned. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, only objects synced by Merge before this date time will be returned. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomObjectsApi.CustomObjectClassesCustomObjectsList(context.Background(), customObjectClassId).XAccountToken(xAccountToken).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).IncludeRemoteFields(includeRemoteFields).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.CustomObjectClassesCustomObjectsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomObjectClassesCustomObjectsList`: PaginatedCustomObjectList
    fmt.Fprintf(os.Stdout, "Response from `CustomObjectsApi.CustomObjectClassesCustomObjectsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customObjectClassId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomObjectClassesCustomObjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **createdAfter** | **time.Time** | If provided, will only return objects created after this datetime. | 
 **createdBefore** | **time.Time** | If provided, will only return objects created before this datetime. | 
 **cursor** | **string** | The pagination cursor value. | 
 **includeDeletedData** | **bool** | Whether to include data that was marked as deleted by third party webhooks. | 
 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **includeRemoteFields** | **bool** | Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format. | 
 **modifiedAfter** | **time.Time** | If provided, only objects synced by Merge after this date time will be returned. | 
 **modifiedBefore** | **time.Time** | If provided, only objects synced by Merge before this date time will be returned. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **remoteId** | **string** | The API provider&#39;s ID for the given object. | 

### Return type

[**PaginatedCustomObjectList**](PaginatedCustomObjectList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomObjectClassesCustomObjectsMetaPatchRetrieve

> MetaResponse CustomObjectClassesCustomObjectsMetaPatchRetrieve(ctx, customObjectClassId, id).XAccountToken(xAccountToken).Execute()





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
    xAccountToken := "xAccountToken_example" // string | Token identifying the end user.
    customObjectClassId := TODO // string | 
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomObjectsApi.CustomObjectClassesCustomObjectsMetaPatchRetrieve(context.Background(), customObjectClassId, id).XAccountToken(xAccountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.CustomObjectClassesCustomObjectsMetaPatchRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomObjectClassesCustomObjectsMetaPatchRetrieve`: MetaResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomObjectsApi.CustomObjectClassesCustomObjectsMetaPatchRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customObjectClassId** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomObjectClassesCustomObjectsMetaPatchRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 



### Return type

[**MetaResponse**](MetaResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomObjectClassesCustomObjectsMetaPostRetrieve

> MetaResponse CustomObjectClassesCustomObjectsMetaPostRetrieve(ctx, customObjectClassId).XAccountToken(xAccountToken).Execute()





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
    xAccountToken := "xAccountToken_example" // string | Token identifying the end user.
    customObjectClassId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomObjectsApi.CustomObjectClassesCustomObjectsMetaPostRetrieve(context.Background(), customObjectClassId).XAccountToken(xAccountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.CustomObjectClassesCustomObjectsMetaPostRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomObjectClassesCustomObjectsMetaPostRetrieve`: MetaResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomObjectsApi.CustomObjectClassesCustomObjectsMetaPostRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customObjectClassId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomObjectClassesCustomObjectsMetaPostRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 


### Return type

[**MetaResponse**](MetaResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomObjectClassesCustomObjectsPartialUpdate

> CRMCustomObjectResponse CustomObjectClassesCustomObjectsPartialUpdate(ctx, customObjectClassId, id).XAccountToken(xAccountToken).PatchedCRMCustomObjectEndpointRequest(patchedCRMCustomObjectEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()





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
    xAccountToken := "xAccountToken_example" // string | Token identifying the end user.
    customObjectClassId := TODO // string | 
    id := TODO // string | 
    patchedCRMCustomObjectEndpointRequest := *openapiclient.NewPatchedCRMCustomObjectEndpointRequest(*openapiclient.NewCustomObjectRequest(map[string]interface{}{"key": interface{}(123)})) // PatchedCRMCustomObjectEndpointRequest | 
    isDebugMode := true // bool | Whether to include debug fields (such as log file links) in the response. (optional)
    runAsync := true // bool | Whether or not third-party updates should be run asynchronously. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomObjectsApi.CustomObjectClassesCustomObjectsPartialUpdate(context.Background(), customObjectClassId, id).XAccountToken(xAccountToken).PatchedCRMCustomObjectEndpointRequest(patchedCRMCustomObjectEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.CustomObjectClassesCustomObjectsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomObjectClassesCustomObjectsPartialUpdate`: CRMCustomObjectResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomObjectsApi.CustomObjectClassesCustomObjectsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customObjectClassId** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomObjectClassesCustomObjectsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 


 **patchedCRMCustomObjectEndpointRequest** | [**PatchedCRMCustomObjectEndpointRequest**](PatchedCRMCustomObjectEndpointRequest.md) |  | 
 **isDebugMode** | **bool** | Whether to include debug fields (such as log file links) in the response. | 
 **runAsync** | **bool** | Whether or not third-party updates should be run asynchronously. | 

### Return type

[**CRMCustomObjectResponse**](CRMCustomObjectResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomObjectClassesCustomObjectsRetrieve

> CustomObject CustomObjectClassesCustomObjectsRetrieve(ctx, customObjectClassId, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).IncludeRemoteFields(includeRemoteFields).Execute()





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
    xAccountToken := "xAccountToken_example" // string | Token identifying the end user.
    customObjectClassId := TODO // string | 
    id := TODO // string | 
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    includeRemoteFields := true // bool | Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomObjectsApi.CustomObjectClassesCustomObjectsRetrieve(context.Background(), customObjectClassId, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).IncludeRemoteFields(includeRemoteFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.CustomObjectClassesCustomObjectsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomObjectClassesCustomObjectsRetrieve`: CustomObject
    fmt.Fprintf(os.Stdout, "Response from `CustomObjectsApi.CustomObjectClassesCustomObjectsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customObjectClassId** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomObjectClassesCustomObjectsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 


 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **includeRemoteFields** | **bool** | Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format. | 

### Return type

[**CustomObject**](CustomObject.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

