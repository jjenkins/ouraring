/*
Oura API Documentation

# Overview   The Oura API allows Oura users and partner applications to improve their user experience with Oura data.  This document describes the Oura API Version 2 (V2), which supports access to the latest Oura Ring data types.  For access to other data types—which have not yet migrated to V2—refer to the [Oura API Version 1 (V1)](https://cloud.ouraring.com/docs/) documentation.  # Data Access  Individual Oura users can access their own data through the API by using a [Personal Access Token](https://cloud.ouraring.com/personal-access-tokens).  If you want to retrieve data for multiple users, a registered [API Application](https://cloud.ouraring.com/oauth/applications) is required.  API Applications are limited to **10** users before requiring approval from Oura. There is no limit once an application is approved.  Additionally, Oura users **must provide consent** to share each data type an API Application has access to.  All data access requests through the Oura API require [Authentication](https://cloud.ouraring.com/docs/authentication).  Additionally, we recommend that Oura users keep their mobile app updated to support API access for the latest data types.  # Authentication  The Oura API provides two methods for Authentication: (1) the OAuth2 protocol and (2) Personal Access Tokens. For more information on the OAuth2 flow, see our [Authentication instructions](https://cloud.ouraring.com/docs/authentication).  Access tokens must be included in the request header as follows: ```http GET /v2/usercollection/personal_info HTTP/1.1 Host: api.ouraring.com Authorization: Bearer <token> ```  # Oura HTTP Response Codes  | Response Code                        | Description | | ------------------------------------ | - | | 200 OK                               | Successful Response         | | 400 Query Parameter Validation Error | The request contains query parameters that are invalid or incorrectly formatted. | | 429 Request Rate Limit Exceeded        | The API is rate limited to 5000 requests in a 5 minute period. You will receive a 429 error code if you exceed this limit. [Contact us](mailto:api-support@ouraring.com) if you expect your usage to exceed this limit.| 

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oura

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// SleepRoutesAPIService SleepRoutesAPI service
type SleepRoutesAPIService service

type ApiMultipleSleepDocumentsV2UsercollectionSleepGetRequest struct {
	ctx context.Context
	ApiService *SleepRoutesAPIService
	startDate *StartDate
	endDate *EndDate1
	nextToken *string
}

func (r ApiMultipleSleepDocumentsV2UsercollectionSleepGetRequest) StartDate(startDate StartDate) ApiMultipleSleepDocumentsV2UsercollectionSleepGetRequest {
	r.startDate = &startDate
	return r
}

func (r ApiMultipleSleepDocumentsV2UsercollectionSleepGetRequest) EndDate(endDate EndDate1) ApiMultipleSleepDocumentsV2UsercollectionSleepGetRequest {
	r.endDate = &endDate
	return r
}

func (r ApiMultipleSleepDocumentsV2UsercollectionSleepGetRequest) NextToken(nextToken string) ApiMultipleSleepDocumentsV2UsercollectionSleepGetRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiMultipleSleepDocumentsV2UsercollectionSleepGetRequest) Execute() (*MultiDocumentResponseSleepModel, *http.Response, error) {
	return r.ApiService.MultipleSleepDocumentsV2UsercollectionSleepGetExecute(r)
}

/*
MultipleSleepDocumentsV2UsercollectionSleepGet Multiple Sleep Documents

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMultipleSleepDocumentsV2UsercollectionSleepGetRequest
*/
func (a *SleepRoutesAPIService) MultipleSleepDocumentsV2UsercollectionSleepGet(ctx context.Context) ApiMultipleSleepDocumentsV2UsercollectionSleepGetRequest {
	return ApiMultipleSleepDocumentsV2UsercollectionSleepGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MultiDocumentResponseSleepModel
func (a *SleepRoutesAPIService) MultipleSleepDocumentsV2UsercollectionSleepGetExecute(r ApiMultipleSleepDocumentsV2UsercollectionSleepGetRequest) (*MultiDocumentResponseSleepModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MultiDocumentResponseSleepModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SleepRoutesAPIService.MultipleSleepDocumentsV2UsercollectionSleepGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/usercollection/sleep"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "next_token", r.nextToken, "form", "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSingleSleepDocumentV2UsercollectionSleepDocumentIdGetRequest struct {
	ctx context.Context
	ApiService *SleepRoutesAPIService
	documentId string
}

func (r ApiSingleSleepDocumentV2UsercollectionSleepDocumentIdGetRequest) Execute() (*SleepModel, *http.Response, error) {
	return r.ApiService.SingleSleepDocumentV2UsercollectionSleepDocumentIdGetExecute(r)
}

/*
SingleSleepDocumentV2UsercollectionSleepDocumentIdGet Single Sleep Document

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param documentId
 @return ApiSingleSleepDocumentV2UsercollectionSleepDocumentIdGetRequest
*/
func (a *SleepRoutesAPIService) SingleSleepDocumentV2UsercollectionSleepDocumentIdGet(ctx context.Context, documentId string) ApiSingleSleepDocumentV2UsercollectionSleepDocumentIdGetRequest {
	return ApiSingleSleepDocumentV2UsercollectionSleepDocumentIdGetRequest{
		ApiService: a,
		ctx: ctx,
		documentId: documentId,
	}
}

// Execute executes the request
//  @return SleepModel
func (a *SleepRoutesAPIService) SingleSleepDocumentV2UsercollectionSleepDocumentIdGetExecute(r ApiSingleSleepDocumentV2UsercollectionSleepDocumentIdGetRequest) (*SleepModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SleepModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SleepRoutesAPIService.SingleSleepDocumentV2UsercollectionSleepDocumentIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/usercollection/sleep/{document_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"document_id"+"}", url.PathEscape(parameterValueToString(r.documentId, "documentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
