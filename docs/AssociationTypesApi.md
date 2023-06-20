# \AssociationTypesApi

All URIs are relative to *https://api.merge.dev/api/crm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomObjectClassesAssociationTypesCreate**](AssociationTypesApi.md#CustomObjectClassesAssociationTypesCreate) | **Post** /custom-object-classes/{custom_object_class_id}/association-types | 
[**CustomObjectClassesAssociationTypesList**](AssociationTypesApi.md#CustomObjectClassesAssociationTypesList) | **Get** /custom-object-classes/{custom_object_class_id}/association-types | 
[**CustomObjectClassesAssociationTypesMetaPostRetrieve**](AssociationTypesApi.md#CustomObjectClassesAssociationTypesMetaPostRetrieve) | **Get** /custom-object-classes/{custom_object_class_id}/association-types/meta/post | 
[**CustomObjectClassesAssociationTypesRetrieve**](AssociationTypesApi.md#CustomObjectClassesAssociationTypesRetrieve) | **Get** /custom-object-classes/{custom_object_class_id}/association-types/{id} | 



## CustomObjectClassesAssociationTypesCreate

> CRMAssociationTypeResponse CustomObjectClassesAssociationTypesCreate(ctx, customObjectClassId).XAccountToken(xAccountToken).CRMAssociationTypeEndpointRequest(cRMAssociationTypeEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()





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
    cRMAssociationTypeEndpointRequest := *openapiclient.NewCRMAssociationTypeEndpointRequest(*openapiclient.NewAssociationTypeRequestRequest(*openapiclient.NewObjectClassDescriptionRequest("Id_example", openapiclient.OriginTypeEnum("CUSTOM_OBJECT")), []openapiclient.ObjectClassDescriptionRequest{*openapiclient.NewObjectClassDescriptionRequest("Id_example", openapiclient.OriginTypeEnum("CUSTOM_OBJECT"))}, "RemoteKeyName_example")) // CRMAssociationTypeEndpointRequest | 
    isDebugMode := true // bool | Whether to include debug fields (such as log file links) in the response. (optional)
    runAsync := true // bool | Whether or not third-party updates should be run asynchronously. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssociationTypesApi.CustomObjectClassesAssociationTypesCreate(context.Background(), customObjectClassId).XAccountToken(xAccountToken).CRMAssociationTypeEndpointRequest(cRMAssociationTypeEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationTypesApi.CustomObjectClassesAssociationTypesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomObjectClassesAssociationTypesCreate`: CRMAssociationTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `AssociationTypesApi.CustomObjectClassesAssociationTypesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customObjectClassId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomObjectClassesAssociationTypesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **cRMAssociationTypeEndpointRequest** | [**CRMAssociationTypeEndpointRequest**](CRMAssociationTypeEndpointRequest.md) |  | 
 **isDebugMode** | **bool** | Whether to include debug fields (such as log file links) in the response. | 
 **runAsync** | **bool** | Whether or not third-party updates should be run asynchronously. | 

### Return type

[**CRMAssociationTypeResponse**](CRMAssociationTypeResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomObjectClassesAssociationTypesList

> PaginatedAssociationTypeList CustomObjectClassesAssociationTypesList(ctx, customObjectClassId).XAccountToken(xAccountToken).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()





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
    modifiedAfter := time.Now() // time.Time | If provided, only objects synced by Merge after this date time will be returned. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, only objects synced by Merge before this date time will be returned. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssociationTypesApi.CustomObjectClassesAssociationTypesList(context.Background(), customObjectClassId).XAccountToken(xAccountToken).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationTypesApi.CustomObjectClassesAssociationTypesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomObjectClassesAssociationTypesList`: PaginatedAssociationTypeList
    fmt.Fprintf(os.Stdout, "Response from `AssociationTypesApi.CustomObjectClassesAssociationTypesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customObjectClassId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomObjectClassesAssociationTypesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

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

[**PaginatedAssociationTypeList**](PaginatedAssociationTypeList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomObjectClassesAssociationTypesMetaPostRetrieve

> MetaResponse CustomObjectClassesAssociationTypesMetaPostRetrieve(ctx, customObjectClassId).XAccountToken(xAccountToken).Execute()





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
    resp, r, err := api_client.AssociationTypesApi.CustomObjectClassesAssociationTypesMetaPostRetrieve(context.Background(), customObjectClassId).XAccountToken(xAccountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationTypesApi.CustomObjectClassesAssociationTypesMetaPostRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomObjectClassesAssociationTypesMetaPostRetrieve`: MetaResponse
    fmt.Fprintf(os.Stdout, "Response from `AssociationTypesApi.CustomObjectClassesAssociationTypesMetaPostRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customObjectClassId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomObjectClassesAssociationTypesMetaPostRetrieveRequest struct via the builder pattern


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


## CustomObjectClassesAssociationTypesRetrieve

> AssociationType CustomObjectClassesAssociationTypesRetrieve(ctx, customObjectClassId, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssociationTypesApi.CustomObjectClassesAssociationTypesRetrieve(context.Background(), customObjectClassId, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationTypesApi.CustomObjectClassesAssociationTypesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomObjectClassesAssociationTypesRetrieve`: AssociationType
    fmt.Fprintf(os.Stdout, "Response from `AssociationTypesApi.CustomObjectClassesAssociationTypesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customObjectClassId** | [**string**](.md) |  | 
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomObjectClassesAssociationTypesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 


 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 

### Return type

[**AssociationType**](AssociationType.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

