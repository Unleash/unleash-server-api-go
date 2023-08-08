/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// ImportExportApiService ImportExportApi service
type ImportExportApiService service

type ApiCallImportRequest struct {
	ctx         context.Context
	ApiService  *ImportExportApiService
	requestBody *map[string]interface{}
}

// stateSchema
func (r ApiCallImportRequest) RequestBody(requestBody map[string]interface{}) ApiCallImportRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiCallImportRequest) Execute() (*http.Response, error) {
	return r.ApiService.CallImportExecute(r)
}

/*
CallImport Import state (deprecated)

Imports state into the system. Deprecated in favor of /api/admin/features-batch/import

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCallImportRequest

Deprecated
*/
func (a *ImportExportApiService) CallImport(ctx context.Context) ApiCallImportRequest {
	return ApiCallImportRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
// Deprecated
func (a *ImportExportApiService) CallImportExecute(r ApiCallImportRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportExportApiService.CallImport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/admin/state/import"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestBody == nil {
		return nil, reportError("requestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.requestBody
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiExportRequest struct {
	ctx            context.Context
	ApiService     *ImportExportApiService
	format         *string
	download       *ExportDownloadParameter
	strategies     *ExportStrategiesParameter
	featureToggles *ExportStrategiesParameter
	projects       *ExportStrategiesParameter
	tags           *ExportStrategiesParameter
	environments   *ExportStrategiesParameter
}

// Desired export format. Must be either &#x60;json&#x60; or &#x60;yaml&#x60;.
func (r ApiExportRequest) Format(format string) ApiExportRequest {
	r.format = &format
	return r
}

// Whether exported data should be downloaded as a file.
func (r ApiExportRequest) Download(download ExportDownloadParameter) ApiExportRequest {
	r.download = &download
	return r
}

// Whether strategies should be included in the exported data.
func (r ApiExportRequest) Strategies(strategies ExportStrategiesParameter) ApiExportRequest {
	r.strategies = &strategies
	return r
}

// Whether feature toggles should be included in the exported data.
func (r ApiExportRequest) FeatureToggles(featureToggles ExportStrategiesParameter) ApiExportRequest {
	r.featureToggles = &featureToggles
	return r
}

// Whether projects should be included in the exported data.
func (r ApiExportRequest) Projects(projects ExportStrategiesParameter) ApiExportRequest {
	r.projects = &projects
	return r
}

// Whether tag types, tags, and feature_tags should be included in the exported data.
func (r ApiExportRequest) Tags(tags ExportStrategiesParameter) ApiExportRequest {
	r.tags = &tags
	return r
}

// Whether environments should be included in the exported data.
func (r ApiExportRequest) Environments(environments ExportStrategiesParameter) ApiExportRequest {
	r.environments = &environments
	return r
}

func (r ApiExportRequest) Execute() (*StateSchema, *http.Response, error) {
	return r.ApiService.ExportExecute(r)
}

/*
Export Export state (deprecated)

Exports the current state of the system. Deprecated in favor of /api/admin/features-batch/export

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExportRequest

Deprecated
*/
func (a *ImportExportApiService) Export(ctx context.Context) ApiExportRequest {
	return ApiExportRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return StateSchema
// Deprecated
func (a *ImportExportApiService) ExportExecute(r ApiExportRequest) (*StateSchema, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StateSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportExportApiService.Export")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/admin/state/export"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "")
	}
	if r.download != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "download", r.download, "")
	}
	if r.strategies != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "strategies", r.strategies, "")
	}
	if r.featureToggles != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "featureToggles", r.featureToggles, "")
	}
	if r.projects != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "projects", r.projects, "")
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tags", r.tags, "")
	}
	if r.environments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "environments", r.environments, "")
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
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

type ApiExportFeaturesRequest struct {
	ctx         context.Context
	ApiService  *ImportExportApiService
	requestBody *map[string]interface{}
}

// exportQuerySchema
func (r ApiExportFeaturesRequest) RequestBody(requestBody map[string]interface{}) ApiExportFeaturesRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiExportFeaturesRequest) Execute() (*ExportResultSchema, *http.Response, error) {
	return r.ApiService.ExportFeaturesExecute(r)
}

/*
ExportFeatures Export feature toggles from an environment

Exports all features listed in the `features` property from the environment specified in the request body. If set to `true`, the `downloadFile` property will let you download a file with the exported data. Otherwise, the export data is returned directly as JSON. Refer to the documentation for more information about [Unleash's export functionality](https://docs.getunleash.io/reference/deploy/environment-import-export#export).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExportFeaturesRequest
*/
func (a *ImportExportApiService) ExportFeatures(ctx context.Context) ApiExportFeaturesRequest {
	return ApiExportFeaturesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return ExportResultSchema
func (a *ImportExportApiService) ExportFeaturesExecute(r ApiExportFeaturesRequest) (*ExportResultSchema, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExportResultSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportExportApiService.ExportFeatures")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/admin/features-batch/export"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestBody == nil {
		return localVarReturnValue, nil, reportError("requestBody is required and must be specified")
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
	localVarPostBody = r.requestBody
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v GetGroup404Response
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

type ApiImportTogglesRequest struct {
	ctx                 context.Context
	ApiService          *ImportExportApiService
	importTogglesSchema *ImportTogglesSchema
}

// importTogglesSchema
func (r ApiImportTogglesRequest) ImportTogglesSchema(importTogglesSchema ImportTogglesSchema) ApiImportTogglesRequest {
	r.importTogglesSchema = &importTogglesSchema
	return r
}

func (r ApiImportTogglesRequest) Execute() (*http.Response, error) {
	return r.ApiService.ImportTogglesExecute(r)
}

/*
ImportToggles Import feature toggles

[Import feature toggles](https://docs.getunleash.io/reference/deploy/environment-import-export#import) into a specific project and environment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiImportTogglesRequest
*/
func (a *ImportExportApiService) ImportToggles(ctx context.Context) ApiImportTogglesRequest {
	return ApiImportTogglesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ImportExportApiService) ImportTogglesExecute(r ApiImportTogglesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportExportApiService.ImportToggles")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/admin/features-batch/import"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.importTogglesSchema == nil {
		return nil, reportError("importTogglesSchema is required and must be specified")
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
	localVarPostBody = r.importTogglesSchema
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GetGroup404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiValidateImportRequest struct {
	ctx                 context.Context
	ApiService          *ImportExportApiService
	importTogglesSchema *ImportTogglesSchema
}

// importTogglesSchema
func (r ApiValidateImportRequest) ImportTogglesSchema(importTogglesSchema ImportTogglesSchema) ApiValidateImportRequest {
	r.importTogglesSchema = &importTogglesSchema
	return r
}

func (r ApiValidateImportRequest) Execute() (*ImportTogglesValidateSchema, *http.Response, error) {
	return r.ApiService.ValidateImportExecute(r)
}

/*
ValidateImport Validate feature import data

Validates a feature toggle data set. Checks whether the data can be imported into the specified project and environment. The returned value is an object that contains errors, warnings, and permissions required to perform the import, as described in the [import documentation](https://docs.getunleash.io/reference/deploy/environment-import-export#import).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValidateImportRequest
*/
func (a *ImportExportApiService) ValidateImport(ctx context.Context) ApiValidateImportRequest {
	return ApiValidateImportRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return ImportTogglesValidateSchema
func (a *ImportExportApiService) ValidateImportExecute(r ApiValidateImportRequest) (*ImportTogglesValidateSchema, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ImportTogglesValidateSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportExportApiService.ValidateImport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/admin/features-batch/validate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.importTogglesSchema == nil {
		return localVarReturnValue, nil, reportError("importTogglesSchema is required and must be specified")
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
	localVarPostBody = r.importTogglesSchema
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v GetGroup404Response
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
