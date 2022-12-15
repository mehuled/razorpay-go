# {{classname}}

All URIs are relative to *https://api.razorpay.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrdersPost**](OrdersApi.md#OrdersPost) | **Post** /orders | Create a new Order

# **OrdersPost**
> Order OrdersPost(ctx, body)
Create a new Order

Create a new orders for a Payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrdersBody**](OrdersBody.md)| Create a new order for a Payment | 

### Return type

[**Order**](Order.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

