/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nodeum

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type TaskOptionsApiService service

/* 
TaskOptionsApiService Creates a new task option.
**API Key Scope**: task_options / create
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param taskId Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID.
 * @param taskOptionBody

@return TaskOption
*/
func (a *TaskOptionsApiService) CreateTaskOption(ctx context.Context, taskId string, taskOptionBody TaskOption) (TaskOption, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TaskOption
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tasks/{task_id}/task_options"
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &taskOptionBody
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 201 {
			var v TaskOption
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 422 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
TaskOptionsApiService Destroys a specific task option.
**API Key Scope**: task_options / destroy
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param taskId Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID.
 * @param taskOptionId Numeric ID of task option.


*/
func (a *TaskOptionsApiService) DestroyTaskOption(ctx context.Context, taskId string, taskOptionId int32) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tasks/{task_id}/task_options/{task_option_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_option_id"+"}", fmt.Sprintf("%v", taskOptionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/* 
TaskOptionsApiService Lists all task options.
**API Key Scope**: task_options / index
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param taskId Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID.
 * @param optional nil or *IndexTaskOptionsOpts - Optional Parameters:
     * @param "Limit" (optional.Int32) -  The number of items to display for pagination.
     * @param "Offset" (optional.Int32) -  The number of items to skip for pagination.
     * @param "SortBy" (optional.Interface of []string) -  Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;.
     * @param "Id" (optional.String) -  Filter on id
     * @param "Type_" (optional.String) -  Filter on type
     * @param "Value" (optional.String) -  Filter on value

@return TaskOptionCollection
*/

type IndexTaskOptionsOpts struct { 
	Limit optional.Int32
	Offset optional.Int32
	SortBy optional.Interface
	Id optional.String
	Type_ optional.String
	Value optional.String
}

func (a *TaskOptionsApiService) IndexTaskOptions(ctx context.Context, taskId string, localVarOptionals *IndexTaskOptionsOpts) (TaskOptionCollection, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TaskOptionCollection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tasks/{task_id}/task_options"
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortBy.IsSet() {
		localVarQueryParams.Add("sort_by", parameterToString(localVarOptionals.SortBy.Value(), "pipes"))
	}
	if localVarOptionals != nil && localVarOptionals.Id.IsSet() {
		localVarQueryParams.Add("id", parameterToString(localVarOptionals.Id.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Value.IsSet() {
		localVarQueryParams.Add("value", parameterToString(localVarOptionals.Value.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v TaskOptionCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
TaskOptionsApiService Displays a specific task option.
**API Key Scope**: task_options / show
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param taskId Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID.
 * @param taskOptionId Numeric ID of task option.

@return TaskOption
*/
func (a *TaskOptionsApiService) ShowTaskOption(ctx context.Context, taskId string, taskOptionId int32) (TaskOption, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TaskOption
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tasks/{task_id}/task_options/{task_option_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_option_id"+"}", fmt.Sprintf("%v", taskOptionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v TaskOption
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
TaskOptionsApiService Updates a specific task option.
**API Key Scope**: task_options / update
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param taskId Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID.
 * @param taskOptionId Numeric ID of task option.
 * @param taskOptionBody

@return TaskOption
*/
func (a *TaskOptionsApiService) UpdateTaskOption(ctx context.Context, taskId string, taskOptionId int32, taskOptionBody TaskOption) (TaskOption, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TaskOption
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tasks/{task_id}/task_options/{task_option_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_option_id"+"}", fmt.Sprintf("%v", taskOptionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &taskOptionBody
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v TaskOption
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 422 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
