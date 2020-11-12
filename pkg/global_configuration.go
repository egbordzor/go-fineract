package pkg

import (
	_context "context"
	"github.com/antihax/optional"
	"github.com/wondenge/go-fineract/pkg/models"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// GlobalConfigurationApiService GlobalConfigurationApi service
type GlobalConfigurationApiService service

// RetrieveConfigurationOpts Optional parameters for the method 'RetrieveConfiguration'
type RetrieveConfigurationOpts struct {
	Survey optional.Bool
}

/*
RetrieveConfiguration Retrieve Global Configuration | Retrieve Global Configuration for surveys
Returns the list global enable/disable configurations.  Example Requests:  configurations   Returns the list global enable/disable survey configurations.  Example Requests:  configurations/survey
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveConfigurationOpts - Optional Parameters:
 * @param "Survey" (optional.Bool) -  survey
@return []GetGlobalConfigurationsResponse
*/
func (a *GlobalConfigurationApiService) RetrieveConfiguration(ctx _context.Context, opts *RetrieveConfigurationOpts) ([]models.GetGlobalConfigurationsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  []models.GetGlobalConfigurationsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/configurations"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.Survey.IsSet() {
		queryParams.Add("survey", parameterToString(opts.Survey.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
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
			headerParams["fineract-platform-tenantid"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, path, httpMethod, postBody, headerParams, queryParams, formParams, formFileName, fileName, fileBytes)
	if err != nil {
		return returnValue, nil, err
	}

	res, err := a.client.callAPI(r)
	if err != nil || res == nil {
		return returnValue, res, err
	}

	payload, err := _ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return returnValue, res, err
	}

	if res.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  payload,
			error: res.Status,
		}
		return returnValue, res, newErr
	}

	err = a.client.decode(&returnValue, payload, res.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  payload,
			error: err.Error(),
		}
		return returnValue, res, newErr
	}

	return returnValue, res, nil
}

/*
RetrieveOne3 Retrieve Global Configuration
Returns a global enable/disable configurations.  Example Requests:  configurations/1
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param configId configId
@return GetGlobalConfigurationsResponse
*/
func (a *GlobalConfigurationApiService) RetrieveOne3(ctx _context.Context, configId int64) (models.GetGlobalConfigurationsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetGlobalConfigurationsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/configurations/{configId}"
	path = strings.Replace(path, "{"+"configId"+"}", _neturl.QueryEscape(parameterToString(configId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
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
			headerParams["fineract-platform-tenantid"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, path, httpMethod, postBody, headerParams, queryParams, formParams, formFileName, fileName, fileBytes)
	if err != nil {
		return returnValue, nil, err
	}

	res, err := a.client.callAPI(r)
	if err != nil || res == nil {
		return returnValue, res, err
	}

	payload, err := _ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return returnValue, res, err
	}

	if res.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  payload,
			error: res.Status,
		}
		return returnValue, res, newErr
	}

	err = a.client.decode(&returnValue, payload, res.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  payload,
			error: err.Error(),
		}
		return returnValue, res, newErr
	}

	return returnValue, res, nil
}

/*
UpdateConfiguration1 Update Global Configuration
Updates an enable/disable global configuration item.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param configId configId
 * @param putGlobalConfigurationsRequest
@return PutGlobalConfigurationsResponse
*/
func (a *GlobalConfigurationApiService) UpdateConfiguration1(ctx _context.Context, configId int64, putGlobalConfigurationsRequest models.PutGlobalConfigurationsRequest) (models.PutGlobalConfigurationsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPut
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PutGlobalConfigurationsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/configurations/{configId}"
	path = strings.Replace(path, "{"+"configId"+"}", _neturl.QueryEscape(parameterToString(configId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	// body params
	postBody = &putGlobalConfigurationsRequest
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			headerParams["fineract-platform-tenantid"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, path, httpMethod, postBody, headerParams, queryParams, formParams, formFileName, fileName, fileBytes)
	if err != nil {
		return returnValue, nil, err
	}

	res, err := a.client.callAPI(r)
	if err != nil || res == nil {
		return returnValue, res, err
	}

	payload, err := _ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return returnValue, res, err
	}

	if res.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  payload,
			error: res.Status,
		}
		return returnValue, res, newErr
	}

	err = a.client.decode(&returnValue, payload, res.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  payload,
			error: err.Error(),
		}
		return returnValue, res, newErr
	}

	return returnValue, res, nil
}
