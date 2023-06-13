# \AssociationsApi

All URIs are relative to *https://api.merge.dev/api/crm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomObjectClassesCustomObjectsAssociationsList**](AssociationsApi.md#CustomObjectClassesCustomObjectsAssociationsList) | **Get** /custom-object-classes/{custom_object_class_id}/custom-objects/{object_id}/associations | 
[**CustomObjectClassesCustomObjectsAssociationsUpdate**](AssociationsApi.md#CustomObjectClassesCustomObjectsAssociationsUpdate) | **Put** /custom-object-classes/{source_class_id}/custom-objects/{source_object_id}/associations/{target_class_id}/{target_object_id}/{association_type_id} | 



## CustomObjectClassesCustomObjectsAssociationsList

> PaginatedAssociationList CustomObjectClassesCustomObjectsAssociationsList(ctx, customObjectClassId, objectId).XAccountToken(xAccountToken).AssociationTypeId(associationTypeId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()





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
    objectId := TODO // string | 
    associationTypeId := "associationTypeId_example" // string | If provided, will only return opportunities with this association_type. (optional)
    createdAfter := time.Now() // time.Time | If provided, will only return objects created after this datetime. (optional)
    createdBefore := time.Now() // time.Time | If provided, will only return objects created before this datetime. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    includeDeletedData := true // bool | Whether to include data that was marked as deleted by third party webhooks. (optional)
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    modifiedAfter := time.Now() // time.Time | If provided, only objects synced by Merge after this date time will be returned. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, only objects synced by Merge before this date time will be returned. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssociationsApi.CustomObjectClassesCustomObjectsAssociationsList(context.Background(), customObjectClassId, objectId).XAccountToken(xAccountToken).AssociationTypeId(associationTypeId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.CustomObjectClassesCustomObjectsAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomObjectClassesCustomObjectsAssociationsList`: PaginatedAssociationList
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.CustomObjectClassesCustomObjectsAssociationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customObjectClassId** | [**string**](.md) |  | 
**objectId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomObjectClassesCustomObjectsAssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 


 **associationTypeId** | **string** | If provided, will only return opportunities with this association_type. | 
 **createdAfter** | **time.Time** | If provided, will only return objects created after this datetime. | 
 **createdBefore** | **time.Time** | If provided, will only return objects created before this datetime. | 
 **cursor** | **string** | The pagination cursor value. | 
 **includeDeletedData** | **bool** | Whether to include data that was marked as deleted by third party webhooks. | 
 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **modifiedAfter** | **time.Time** | If provided, only objects synced by Merge after this date time will be returned. | 
 **modifiedBefore** | **time.Time** | If provided, only objects synced by Merge before this date time will be returned. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **remoteId** | **string** | The API provider&#39;s ID for the given object. | 

### Return type

[**PaginatedAssociationList**](PaginatedAssociationList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomObjectClassesCustomObjectsAssociationsUpdate

> Association CustomObjectClassesCustomObjectsAssociationsUpdate(ctx, associationTypeId, sourceClassId, sourceObjectId, targetClassId, targetObjectId).XAccountToken(xAccountToken).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()





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
    associationTypeId := TODO // string | 
    sourceClassId := TODO // string | 
    sourceObjectId := TODO // string | 
    targetClassId := "targetClassId_example" // string | 
    targetObjectId := TODO // string | 
    isDebugMode := true // bool | Whether to include debug fields (such as log file links) in the response. (optional)
    runAsync := true // bool | Whether or not third-party updates should be run asynchronously. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssociationsApi.CustomObjectClassesCustomObjectsAssociationsUpdate(context.Background(), associationTypeId, sourceClassId, sourceObjectId, targetClassId, targetObjectId).XAccountToken(xAccountToken).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.CustomObjectClassesCustomObjectsAssociationsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomObjectClassesCustomObjectsAssociationsUpdate`: Association
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.CustomObjectClassesCustomObjectsAssociationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**associationTypeId** | [**string**](.md) |  | 
**sourceClassId** | [**string**](.md) |  | 
**sourceObjectId** | [**string**](.md) |  | 
**targetClassId** | **string** |  | 
**targetObjectId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomObjectClassesCustomObjectsAssociationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 





 **isDebugMode** | **bool** | Whether to include debug fields (such as log file links) in the response. | 
 **runAsync** | **bool** | Whether or not third-party updates should be run asynchronously. | 

### Return type

[**Association**](Association.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

