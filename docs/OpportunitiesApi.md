# \OpportunitiesApi

All URIs are relative to *https://api.merge.dev/api/crm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpportunitiesCreate**](OpportunitiesApi.md#OpportunitiesCreate) | **Post** /opportunities | 
[**OpportunitiesList**](OpportunitiesApi.md#OpportunitiesList) | **Get** /opportunities | 
[**OpportunitiesMetaPostRetrieve**](OpportunitiesApi.md#OpportunitiesMetaPostRetrieve) | **Get** /opportunities/meta/post | 
[**OpportunitiesRetrieve**](OpportunitiesApi.md#OpportunitiesRetrieve) | **Get** /opportunities/{id} | 



## OpportunitiesCreate

> OpportunityResponse OpportunitiesCreate(ctx).XAccountToken(xAccountToken).OpportunityEndpointRequest(opportunityEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()





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
    opportunityEndpointRequest := *openapiclient.NewOpportunityEndpointRequest(*openapiclient.NewOpportunityRequest()) // OpportunityEndpointRequest | 
    isDebugMode := true // bool | Whether to include debug fields (such as log file links) in the response. (optional)
    runAsync := true // bool | Whether or not third-party updates should be run asynchronously. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpportunitiesApi.OpportunitiesCreate(context.Background()).XAccountToken(xAccountToken).OpportunityEndpointRequest(opportunityEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesApi.OpportunitiesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpportunitiesCreate`: OpportunityResponse
    fmt.Fprintf(os.Stdout, "Response from `OpportunitiesApi.OpportunitiesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpportunitiesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **opportunityEndpointRequest** | [**OpportunityEndpointRequest**](OpportunityEndpointRequest.md) |  | 
 **isDebugMode** | **bool** | Whether to include debug fields (such as log file links) in the response. | 
 **runAsync** | **bool** | Whether or not third-party updates should be run asynchronously. | 

### Return type

[**OpportunityResponse**](OpportunityResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpportunitiesList

> PaginatedOpportunityList OpportunitiesList(ctx).XAccountToken(xAccountToken).AccountId(accountId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).OwnerId(ownerId).PageSize(pageSize).RemoteFields(remoteFields).RemoteId(remoteId).StageId(stageId).Status(status).Execute()





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
    accountId := "accountId_example" // string | If provided, will only return opportunities with this account. (optional)
    createdAfter := time.Now() // time.Time | If provided, will only return objects created after this datetime. (optional)
    createdBefore := time.Now() // time.Time | If provided, will only return objects created before this datetime. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    includeDeletedData := true // bool | Whether to include data that was marked as deleted by third party webhooks. (optional)
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    modifiedAfter := time.Now() // time.Time | If provided, will only return objects modified after this datetime. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, will only return objects modified before this datetime. (optional)
    ownerId := "ownerId_example" // string | If provided, will only return opportunities with this owner. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    remoteFields := "status" // string | Which fields should be returned in non-normalized form. (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)
    stageId := "stageId_example" // string | If provided, will only return opportunities with this stage. (optional)
    status := "status_example" // string | If provided, will only return opportunities with this status. Options: ('OPEN', 'WON', 'LOST') (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpportunitiesApi.OpportunitiesList(context.Background()).XAccountToken(xAccountToken).AccountId(accountId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).OwnerId(ownerId).PageSize(pageSize).RemoteFields(remoteFields).RemoteId(remoteId).StageId(stageId).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesApi.OpportunitiesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpportunitiesList`: PaginatedOpportunityList
    fmt.Fprintf(os.Stdout, "Response from `OpportunitiesApi.OpportunitiesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpportunitiesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **accountId** | **string** | If provided, will only return opportunities with this account. | 
 **createdAfter** | **time.Time** | If provided, will only return objects created after this datetime. | 
 **createdBefore** | **time.Time** | If provided, will only return objects created before this datetime. | 
 **cursor** | **string** | The pagination cursor value. | 
 **includeDeletedData** | **bool** | Whether to include data that was marked as deleted by third party webhooks. | 
 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **modifiedAfter** | **time.Time** | If provided, will only return objects modified after this datetime. | 
 **modifiedBefore** | **time.Time** | If provided, will only return objects modified before this datetime. | 
 **ownerId** | **string** | If provided, will only return opportunities with this owner. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **remoteFields** | **string** | Which fields should be returned in non-normalized form. | 
 **remoteId** | **string** | The API provider&#39;s ID for the given object. | 
 **stageId** | **string** | If provided, will only return opportunities with this stage. | 
 **status** | **string** | If provided, will only return opportunities with this status. Options: (&#39;OPEN&#39;, &#39;WON&#39;, &#39;LOST&#39;) | 

### Return type

[**PaginatedOpportunityList**](PaginatedOpportunityList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpportunitiesMetaPostRetrieve

> MetaResponse OpportunitiesMetaPostRetrieve(ctx).XAccountToken(xAccountToken).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpportunitiesApi.OpportunitiesMetaPostRetrieve(context.Background()).XAccountToken(xAccountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesApi.OpportunitiesMetaPostRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpportunitiesMetaPostRetrieve`: MetaResponse
    fmt.Fprintf(os.Stdout, "Response from `OpportunitiesApi.OpportunitiesMetaPostRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpportunitiesMetaPostRetrieveRequest struct via the builder pattern


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


## OpportunitiesRetrieve

> Opportunity OpportunitiesRetrieve(ctx, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).RemoteFields(remoteFields).Execute()





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
    remoteFields := "status" // string | Which fields should be returned in non-normalized form. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpportunitiesApi.OpportunitiesRetrieve(context.Background(), id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).RemoteFields(remoteFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesApi.OpportunitiesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpportunitiesRetrieve`: Opportunity
    fmt.Fprintf(os.Stdout, "Response from `OpportunitiesApi.OpportunitiesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpportunitiesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **remoteFields** | **string** | Which fields should be returned in non-normalized form. | 

### Return type

[**Opportunity**](Opportunity.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

