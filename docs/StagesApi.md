# \StagesApi

All URIs are relative to *https://api.merge.dev/api/crm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StagesList**](StagesApi.md#StagesList) | **Get** /stages | 
[**StagesRemoteFieldClassesList**](StagesApi.md#StagesRemoteFieldClassesList) | **Get** /stages/remote-field-classes | 
[**StagesRetrieve**](StagesApi.md#StagesRetrieve) | **Get** /stages/{id} | 



## StagesList

> PaginatedStageList StagesList(ctx).XAccountToken(xAccountToken).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).IncludeRemoteFields(includeRemoteFields).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()





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
    resp, r, err := api_client.StagesApi.StagesList(context.Background()).XAccountToken(xAccountToken).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).IncludeRemoteFields(includeRemoteFields).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesList`: PaginatedStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesListRequest struct via the builder pattern


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

[**PaginatedStageList**](PaginatedStageList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesRemoteFieldClassesList

> PaginatedRemoteFieldClassList StagesRemoteFieldClassesList(ctx).XAccountToken(xAccountToken).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).IncludeRemoteFields(includeRemoteFields).PageSize(pageSize).Execute()





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
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    includeDeletedData := true // bool | Whether to include data that was marked as deleted by third party webhooks. (optional)
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    includeRemoteFields := true // bool | Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StagesApi.StagesRemoteFieldClassesList(context.Background()).XAccountToken(xAccountToken).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).IncludeRemoteFields(includeRemoteFields).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesRemoteFieldClassesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesRemoteFieldClassesList`: PaginatedRemoteFieldClassList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesRemoteFieldClassesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesRemoteFieldClassesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **cursor** | **string** | The pagination cursor value. | 
 **includeDeletedData** | **bool** | Whether to include data that was marked as deleted by third party webhooks. | 
 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **includeRemoteFields** | **bool** | Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedRemoteFieldClassList**](PaginatedRemoteFieldClassList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesRetrieve

> Stage StagesRetrieve(ctx, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).IncludeRemoteFields(includeRemoteFields).Execute()





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
    id := TODO // string | 
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    includeRemoteFields := true // bool | Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StagesApi.StagesRetrieve(context.Background(), id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).IncludeRemoteFields(includeRemoteFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesRetrieve`: Stage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **includeRemoteFields** | **bool** | Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format. | 

### Return type

[**Stage**](Stage.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

