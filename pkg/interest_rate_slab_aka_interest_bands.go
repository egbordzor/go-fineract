package pkg

import (
	_context "context"
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

// InterestRateSlabAKAInterestBandsApiService InterestRateSlabAKAInterestBandsApi service
type InterestRateSlabAKAInterestBandsApiService service

/*
Create8 Create a Slab
Creates a new interest rate slab for an interest rate chart. Mandatory Fields periodType, fromPeriod, annualInterestRate Optional Fields toPeriod and description Example Requests:  interestratecharts/1/chartslabs
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param chartId chartId
 * @param postInterestRateChartsChartIdChartSlabsRequest
@return PostInterestRateChartsChartIdChartSlabsResponse
*/
func (a *InterestRateSlabAKAInterestBandsApiService) Create8(ctx _context.Context, chartId int64, postInterestRateChartsChartIdChartSlabsRequest models.PostInterestRateChartsChartIdChartSlabsRequest) (models.PostInterestRateChartsChartIdChartSlabsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostInterestRateChartsChartIdChartSlabsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/interestratecharts/{chartId}/chartslabs"
	path = strings.Replace(path, "{"+"chartId"+"}", _neturl.QueryEscape(parameterToString(chartId, "")), -1)

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
	postBody = &postInterestRateChartsChartIdChartSlabsRequest
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
Delete12 Delete a Slab
Delete a Slab from a chart
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param chartId chartId
 * @param chartSlabId chartSlabId
@return DeleteInterestRateChartsChartIdChartSlabsResponse
*/
func (a *InterestRateSlabAKAInterestBandsApiService) Delete12(ctx _context.Context, chartId int64, chartSlabId int64) (models.DeleteInterestRateChartsChartIdChartSlabsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodDelete
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.DeleteInterestRateChartsChartIdChartSlabsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/interestratecharts/{chartId}/chartslabs/{chartSlabId}"
	path = strings.Replace(path, "{"+"chartId"+"}", _neturl.QueryEscape(parameterToString(chartId, "")), -1)

	path = strings.Replace(path, "{"+"chartSlabId"+"}", _neturl.QueryEscape(parameterToString(chartSlabId, "")), -1)

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
RetrieveAll24 Retrieve all Slabs
Retrieve list of slabs associated with a chart  Example Requests:  interestratecharts/1/chartslabs
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param chartId chartId
@return []GetInterestRateChartsChartIdChartSlabsResponse
*/
func (a *InterestRateSlabAKAInterestBandsApiService) RetrieveAll24(ctx _context.Context, chartId int64) ([]models.GetInterestRateChartsChartIdChartSlabsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  []models.GetInterestRateChartsChartIdChartSlabsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/interestratecharts/{chartId}/chartslabs"
	path = strings.Replace(path, "{"+"chartId"+"}", _neturl.QueryEscape(parameterToString(chartId, "")), -1)

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
RetrieveOne14 Retrieve a Slab
Retrieve a slab associated with an Interest rate chart  Example Requests:  interestratecharts/1/chartslabs/1
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param chartId chartId
 * @param chartSlabId chartSlabId
@return GetInterestRateChartsChartIdChartSlabsResponse
*/
func (a *InterestRateSlabAKAInterestBandsApiService) RetrieveOne14(ctx _context.Context, chartId int64, chartSlabId int64) (models.GetInterestRateChartsChartIdChartSlabsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetInterestRateChartsChartIdChartSlabsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/interestratecharts/{chartId}/chartslabs/{chartSlabId}"
	path = strings.Replace(path, "{"+"chartId"+"}", _neturl.QueryEscape(parameterToString(chartId, "")), -1)

	path = strings.Replace(path, "{"+"chartSlabId"+"}", _neturl.QueryEscape(parameterToString(chartSlabId, "")), -1)

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
Template8 Method for Template8
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param chartId chartId
@return string
*/
func (a *InterestRateSlabAKAInterestBandsApiService) Template8(ctx _context.Context, chartId int64) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/interestratecharts/{chartId}/chartslabs/template"
	path = strings.Replace(path, "{"+"chartId"+"}", _neturl.QueryEscape(parameterToString(chartId, "")), -1)

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
		var v string
		err = a.client.decode(&v, payload, res.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return returnValue, res, newErr
		}
		newErr.model = v
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
Update12 Update a Slab
It updates the Slab from chart
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param chartId chartId
 * @param chartSlabId chartSlabId
 * @param putInterestRateChartsChartIdChartSlabsChartSlabIdRequest
@return PutInterestRateChartsChartIdChartSlabsChartSlabIdResponse
*/
func (a *InterestRateSlabAKAInterestBandsApiService) Update12(ctx _context.Context, chartId int64, chartSlabId int64, putInterestRateChartsChartIdChartSlabsChartSlabIdRequest models.PutInterestRateChartsChartIdChartSlabsChartSlabIdRequest) (models.PutInterestRateChartsChartIdChartSlabsChartSlabIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPut
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PutInterestRateChartsChartIdChartSlabsChartSlabIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/interestratecharts/{chartId}/chartslabs/{chartSlabId}"
	path = strings.Replace(path, "{"+"chartId"+"}", _neturl.QueryEscape(parameterToString(chartId, "")), -1)

	path = strings.Replace(path, "{"+"chartSlabId"+"}", _neturl.QueryEscape(parameterToString(chartSlabId, "")), -1)

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
	postBody = &putInterestRateChartsChartIdChartSlabsChartSlabIdRequest
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
