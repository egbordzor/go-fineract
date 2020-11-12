package pkg

import (
	_context "context"
	"github.com/wondenge/go-fineract/pkg/models"
	"github.com/wondenge/go-fineract/pkg/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// SavingsChargesApiService SavingsChargesApi service
type SavingsChargesApiService service

/*
AddSavingsAccountCharge Create a Savings account Charge
Creates a Savings account Charge  Mandatory Fields for Savings account Charges: chargeId, amount  chargeId, amount, dueDate, dateFormat, locale  chargeId, amount, feeOnMonthDay, monthDayFormat, locale
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param savingsAccountId savingsAccountId
 * @param postSavingsAccountsSavingsAccountIdChargesRequest
@return PostSavingsAccountsSavingsAccountIdChargesResponse
*/
func (a *SavingsChargesApiService) AddSavingsAccountCharge(ctx _context.Context, savingsAccountId int64, postSavingsAccountsSavingsAccountIdChargesRequest models.PostSavingsAccountsSavingsAccountIdChargesRequest) (models.PostSavingsAccountsSavingsAccountIdChargesResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostSavingsAccountsSavingsAccountIdChargesResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/savingsaccounts/{savingsAccountId}/charges"
	path = strings.Replace(path, "{"+"savingsAccountId"+"}", _neturl.QueryEscape(parameterToString(savingsAccountId, "")), -1)

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
	postBody = &postSavingsAccountsSavingsAccountIdChargesRequest
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
DeleteSavingsAccountCharge Delete a Savings account Charge
Note: Currently, A Savings account Charge may only be removed from Savings that are not yet approved.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param savingsAccountId savingsAccountId
 * @param savingsAccountChargeId savingsAccountChargeId
@return DeleteSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse
*/
func (a *SavingsChargesApiService) DeleteSavingsAccountCharge(ctx _context.Context, savingsAccountId int64, savingsAccountChargeId int64) (models.DeleteSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodDelete
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.DeleteSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/savingsaccounts/{savingsAccountId}/charges/{savingsAccountChargeId}"
	path = strings.Replace(path, "{"+"savingsAccountId"+"}", _neturl.QueryEscape(parameterToString(savingsAccountId, "")), -1)

	path = strings.Replace(path, "{"+"savingsAccountChargeId"+"}", _neturl.QueryEscape(parameterToString(savingsAccountChargeId, "")), -1)

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

// PayOrWaiveSavingsAccountChargeOpts Optional parameters for the method 'PayOrWaiveSavingsAccountCharge'
type PayOrWaiveSavingsAccountChargeOpts struct {
	Command optional.String
}

/*
PayOrWaiveSavingsAccountCharge Pay a Savings account Charge | Waive off a Savings account Charge | Inactivate a Savings account Charge
Pay a Savings account Charge:  An active charge will be paid when savings account is active and having sufficient balance.  Waive off a Savings account Charge:  Outstanding charge amount will be waived off.  Inactivate a Savings account Charge:  A charge will be allowed to inactivate when savings account is active and not having any dues as of today. If charge is overpaid, corresponding charge payment transactions will be reversed.  Showing request/response for &#39;Pay a Savings account Charge&#39;
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param savingsAccountId savingsAccountId
 * @param savingsAccountChargeId savingsAccountChargeId
 * @param postSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdRequest
 * @param optional nil or *PayOrWaiveSavingsAccountChargeOpts - Optional Parameters:
 * @param "Command" (optional.String) -  command
@return PostSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse
*/
func (a *SavingsChargesApiService) PayOrWaiveSavingsAccountCharge(ctx _context.Context, savingsAccountId int64, savingsAccountChargeId int64, postSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdRequest models.PostSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdRequest, opts *PayOrWaiveSavingsAccountChargeOpts) (models.PostSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/savingsaccounts/{savingsAccountId}/charges/{savingsAccountChargeId}"
	path = strings.Replace(path, "{"+"savingsAccountId"+"}", _neturl.QueryEscape(parameterToString(savingsAccountId, "")), -1)

	path = strings.Replace(path, "{"+"savingsAccountChargeId"+"}", _neturl.QueryEscape(parameterToString(savingsAccountChargeId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.Command.IsSet() {
		queryParams.Add("command", parameterToString(opts.Command.Value(), ""))
	}
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
	postBody = &postSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdRequest
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

// RetrieveAllSavingsAccountChargesOpts Optional parameters for the method 'RetrieveAllSavingsAccountCharges'
type RetrieveAllSavingsAccountChargesOpts struct {
	ChargeStatus optional.String
}

/*
RetrieveAllSavingsAccountCharges List Savings Charges
Lists Savings Charges  Example Requests:  savingsaccounts/1/charges  savingsaccounts/1/charges?chargeStatus&#x3D;all  savingsaccounts/1/charges?chargeStatus&#x3D;inactive  savingsaccounts/1/charges?chargeStatus&#x3D;active  savingsaccounts/1/charges?fields&#x3D;name,amountOrPercentage
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param savingsAccountId savingsAccountId
 * @param optional nil or *RetrieveAllSavingsAccountChargesOpts - Optional Parameters:
 * @param "ChargeStatus" (optional.String) -  chargeStatus
@return []GetSavingsAccountsSavingsAccountIdChargesResponse
*/
func (a *SavingsChargesApiService) RetrieveAllSavingsAccountCharges(ctx _context.Context, savingsAccountId int64, opts *RetrieveAllSavingsAccountChargesOpts) ([]models.GetSavingsAccountsSavingsAccountIdChargesResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  []models.GetSavingsAccountsSavingsAccountIdChargesResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/savingsaccounts/{savingsAccountId}/charges"
	path = strings.Replace(path, "{"+"savingsAccountId"+"}", _neturl.QueryEscape(parameterToString(savingsAccountId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.ChargeStatus.IsSet() {
		queryParams.Add("chargeStatus", parameterToString(opts.ChargeStatus.Value(), ""))
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
RetrieveSavingsAccountCharge Retrieve a Savings account Charge
Retrieves a Savings account Charge  Example Requests:  /savingsaccounts/1/charges/5   /savingsaccounts/1/charges/5?fields&#x3D;name,amountOrPercentage
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param savingsAccountId savingsAccountId
 * @param savingsAccountChargeId savingsAccountChargeId
@return GetSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse
*/
func (a *SavingsChargesApiService) RetrieveSavingsAccountCharge(ctx _context.Context, savingsAccountId int64, savingsAccountChargeId int64) (models.GetSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/savingsaccounts/{savingsAccountId}/charges/{savingsAccountChargeId}"
	path = strings.Replace(path, "{"+"savingsAccountId"+"}", _neturl.QueryEscape(parameterToString(savingsAccountId, "")), -1)

	path = strings.Replace(path, "{"+"savingsAccountChargeId"+"}", _neturl.QueryEscape(parameterToString(savingsAccountChargeId, "")), -1)

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
RetrieveTemplate17 Retrieve Savings Charges Template
This is a convenience resource. It can be useful when building maintenance user interface screens for client applications. The template data returned consists of any or all of:  Field Defaults Allowed description Lists Example Request:  savingsaccounts/1/charges/template
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param savingsAccountId savingsAccountId
@return GetSavingsAccountsSavingsAccountIdChargesTemplateResponse
*/
func (a *SavingsChargesApiService) RetrieveTemplate17(ctx _context.Context, savingsAccountId int64) (models.GetSavingsAccountsSavingsAccountIdChargesTemplateResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetSavingsAccountsSavingsAccountIdChargesTemplateResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/savingsaccounts/{savingsAccountId}/charges/template"
	path = strings.Replace(path, "{"+"savingsAccountId"+"}", _neturl.QueryEscape(parameterToString(savingsAccountId, "")), -1)

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
UpdateSavingsAccountCharge Update a Savings account Charge
Currently Savings account Charges may be updated only if the Savings account is not yet approved.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param savingsAccountId savingsAccountId
 * @param savingsAccountChargeId savingsAccountChargeId
 * @param putSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdRequest
@return PutSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse
*/
func (a *SavingsChargesApiService) UpdateSavingsAccountCharge(ctx _context.Context, savingsAccountId int64, savingsAccountChargeId int64, putSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdRequest models.PutSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdRequest) (models.PutSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPut
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PutSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/savingsaccounts/{savingsAccountId}/charges/{savingsAccountChargeId}"
	path = strings.Replace(path, "{"+"savingsAccountId"+"}", _neturl.QueryEscape(parameterToString(savingsAccountId, "")), -1)

	path = strings.Replace(path, "{"+"savingsAccountChargeId"+"}", _neturl.QueryEscape(parameterToString(savingsAccountChargeId, "")), -1)

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
	postBody = &putSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdRequest
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
