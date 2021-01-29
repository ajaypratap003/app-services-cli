/*
 * Strimzi Kubernetes REST API
 *
 * An API to provide k8s REST endpoints for query
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package strimziadminclient

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type DefaultApi interface {

	/*
	 * CreateTopic Creates a new topic
	 * Creates a new topic for Kafka.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiCreateTopicRequest
	 */
	CreateTopic(ctx _context.Context) ApiCreateTopicRequest

	/*
	 * CreateTopicExecute executes the request
	 * @return Topic
	 */
	CreateTopicExecute(r ApiCreateTopicRequest) (Topic, *_nethttp.Response, GenericOpenAPIError)

	/*
	 * DeleteTopic Deletes a  topic
	 * Deletes the topic with the specified id.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param topicId Topic id to delete
	 * @return ApiDeleteTopicRequest
	 */
	DeleteTopic(ctx _context.Context, topicId string) ApiDeleteTopicRequest

	/*
	 * DeleteTopicExecute executes the request
	 */
	DeleteTopicExecute(r ApiDeleteTopicRequest) (*_nethttp.Response, GenericOpenAPIError)

	/*
	 * GetTopic Topic associated with the topic id
	 * Topic
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param topicId The id of the topic
	 * @return ApiGetTopicRequest
	 */
	GetTopic(ctx _context.Context, topicId string) ApiGetTopicRequest

	/*
	 * GetTopicExecute executes the request
	 * @return Topic
	 */
	GetTopicExecute(r ApiGetTopicRequest) (Topic, *_nethttp.Response, GenericOpenAPIError)

	/*
	 * GetTopicsList List of topics
	 * Returns a list of all of the available topics, or the list of topics that meet the users URL Query Parameters.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiGetTopicsListRequest
	 */
	GetTopicsList(ctx _context.Context) ApiGetTopicsListRequest

	/*
	 * GetTopicsListExecute executes the request
	 * @return TopicsList
	 */
	GetTopicsListExecute(r ApiGetTopicsListRequest) (TopicsList, *_nethttp.Response, GenericOpenAPIError)

	/*
	 * UpdateTopic Updates the topic with the specified id.
	 * updates the topic with the new data.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param topicId topic id to update
	 * @return ApiUpdateTopicRequest
	 */
	UpdateTopic(ctx _context.Context, topicId string) ApiUpdateTopicRequest

	/*
	 * UpdateTopicExecute executes the request
	 * @return Topic
	 */
	UpdateTopicExecute(r ApiUpdateTopicRequest) (Topic, *_nethttp.Response, GenericOpenAPIError)
}

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiCreateTopicRequest struct {
	ctx           _context.Context
	ApiService    DefaultApi
	topicSettings *TopicSettings
}

func (r ApiCreateTopicRequest) TopicSettings(topicSettings TopicSettings) ApiCreateTopicRequest {
	r.topicSettings = &topicSettings
	return r
}

func (r ApiCreateTopicRequest) Execute() (Topic, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.CreateTopicExecute(r)
}

/*
 * CreateTopic Creates a new topic
 * Creates a new topic for Kafka.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateTopicRequest
 */
func (a *DefaultApiService) CreateTopic(ctx _context.Context) ApiCreateTopicRequest {
	return ApiCreateTopicRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Topic
 */
func (a *DefaultApiService) CreateTopicExecute(r ApiCreateTopicRequest) (Topic, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  Topic
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreateTopic")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/topics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.topicSettings == nil {
		executionError.error = "topicSettings is required and must be specified"
		return localVarReturnValue, nil, executionError
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.topicSettings
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
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

	return localVarReturnValue, localVarHTTPResponse, executionError
}

type ApiDeleteTopicRequest struct {
	ctx        _context.Context
	ApiService DefaultApi
	topicId    string
}

func (r ApiDeleteTopicRequest) Execute() (*_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.DeleteTopicExecute(r)
}

/*
 * DeleteTopic Deletes a  topic
 * Deletes the topic with the specified id.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param topicId Topic id to delete
 * @return ApiDeleteTopicRequest
 */
func (a *DefaultApiService) DeleteTopic(ctx _context.Context, topicId string) ApiDeleteTopicRequest {
	return ApiDeleteTopicRequest{
		ApiService: a,
		ctx:        ctx,
		topicId:    topicId,
	}
}

/*
 * Execute executes the request
 */
func (a *DefaultApiService) DeleteTopicExecute(r ApiDeleteTopicRequest) (*_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DeleteTopic")
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarPath := localBasePath + "/topics/{topicId}"
	localVarPath = strings.Replace(localVarPath, "{"+"topicId"+"}", _neturl.PathEscape(parameterToString(r.topicId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, executionError
}

type ApiGetTopicRequest struct {
	ctx        _context.Context
	ApiService DefaultApi
	topicId    string
}

func (r ApiGetTopicRequest) Execute() (Topic, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.GetTopicExecute(r)
}

/*
 * GetTopic Topic associated with the topic id
 * Topic
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param topicId The id of the topic
 * @return ApiGetTopicRequest
 */
func (a *DefaultApiService) GetTopic(ctx _context.Context, topicId string) ApiGetTopicRequest {
	return ApiGetTopicRequest{
		ApiService: a,
		ctx:        ctx,
		topicId:    topicId,
	}
}

/*
 * Execute executes the request
 * @return Topic
 */
func (a *DefaultApiService) GetTopicExecute(r ApiGetTopicRequest) (Topic, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  Topic
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetTopic")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/topics/{topicId}"
	localVarPath = strings.Replace(localVarPath, "{"+"topicId"+"}", _neturl.PathEscape(parameterToString(r.topicId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
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

	return localVarReturnValue, localVarHTTPResponse, executionError
}

type ApiGetTopicsListRequest struct {
	ctx        _context.Context
	ApiService DefaultApi
	limit      *int32
	filter     *string
	offset     *int32
}

func (r ApiGetTopicsListRequest) Limit(limit int32) ApiGetTopicsListRequest {
	r.limit = &limit
	return r
}
func (r ApiGetTopicsListRequest) Filter(filter string) ApiGetTopicsListRequest {
	r.filter = &filter
	return r
}
func (r ApiGetTopicsListRequest) Offset(offset int32) ApiGetTopicsListRequest {
	r.offset = &offset
	return r
}

func (r ApiGetTopicsListRequest) Execute() (TopicsList, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.GetTopicsListExecute(r)
}

/*
 * GetTopicsList List of topics
 * Returns a list of all of the available topics, or the list of topics that meet the users URL Query Parameters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetTopicsListRequest
 */
func (a *DefaultApiService) GetTopicsList(ctx _context.Context) ApiGetTopicsListRequest {
	return ApiGetTopicsListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return TopicsList
 */
func (a *DefaultApiService) GetTopicsListExecute(r ApiGetTopicsListRequest) (TopicsList, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  TopicsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetTopicsList")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/topics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
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

	return localVarReturnValue, localVarHTTPResponse, executionError
}

type ApiUpdateTopicRequest struct {
	ctx           _context.Context
	ApiService    DefaultApi
	topicId       string
	topicSettings *TopicSettings
}

func (r ApiUpdateTopicRequest) TopicSettings(topicSettings TopicSettings) ApiUpdateTopicRequest {
	r.topicSettings = &topicSettings
	return r
}

func (r ApiUpdateTopicRequest) Execute() (Topic, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.UpdateTopicExecute(r)
}

/*
 * UpdateTopic Updates the topic with the specified id.
 * updates the topic with the new data.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param topicId topic id to update
 * @return ApiUpdateTopicRequest
 */
func (a *DefaultApiService) UpdateTopic(ctx _context.Context, topicId string) ApiUpdateTopicRequest {
	return ApiUpdateTopicRequest{
		ApiService: a,
		ctx:        ctx,
		topicId:    topicId,
	}
}

/*
 * Execute executes the request
 * @return Topic
 */
func (a *DefaultApiService) UpdateTopicExecute(r ApiUpdateTopicRequest) (Topic, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  Topic
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.UpdateTopic")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/topics/{topicId}"
	localVarPath = strings.Replace(localVarPath, "{"+"topicId"+"}", _neturl.PathEscape(parameterToString(r.topicId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.topicSettings == nil {
		executionError.error = "topicSettings is required and must be specified"
		return localVarReturnValue, nil, executionError
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.topicSettings
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
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

	return localVarReturnValue, localVarHTTPResponse, executionError
}