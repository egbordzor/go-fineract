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

// GeneralLedgerAccountApiService GeneralLedgerAccountApi service
type GeneralLedgerAccountApiService service

// CreateGLAccount1Opts Optional parameters for the method 'CreateGLAccount1'
type CreateGLAccount1Opts struct {
	PostGlAccountsRequest optional.Interface
}

/*
CreateGLAccount1 Create a General Ledger Account
Note: You may optionally create Hierarchical Chart of Accounts by using the \&quot;parentId\&quot; property of an Account Mandatory Fields:  name, glCode, type, usage and manualEntriesAllowed
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CreateGLAccount1Opts - Optional Parameters:
 * @param "PostGlAccountsRequest" (optional.Interface of PostGlAccountsRequest) -
@return PostGlAccountsResponse
*/
func (a *GeneralLedgerAccountApiService) CreateGLAccount1(ctx _context.Context, opts *CreateGLAccount1Opts) (models.PostGlAccountsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostGlAccountsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/glaccounts"
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
	if opts != nil && opts.PostGlAccountsRequest.IsSet() {
		optPostGlAccountsRequest, optPostGlAccountsRequestok := opts.PostGlAccountsRequest.Value().(models.PostGlAccountsRequest)
		if !optPostGlAccountsRequestok {
			return returnValue, nil, reportError("postGlAccountsRequest should be PostGlAccountsRequest")
		}
		postBody = &optPostGlAccountsRequest
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
DeleteGLAccount1 Delete an accounting closure
Note: Only the latest accounting closure associated with a branch may be deleted.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param glAccountId glAccountId
@return DeleteGlAccountsRequest
*/
func (a *GeneralLedgerAccountApiService) DeleteGLAccount1(ctx _context.Context, glAccountId int64) (models.DeleteGlAccountsRequest, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodDelete
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.DeleteGlAccountsRequest
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/glaccounts/{glAccountId}"
	path = strings.Replace(path, "{"+"glAccountId"+"}", _neturl.QueryEscape(parameterToString(glAccountId, "")), -1)

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

// GetGlAccountsTemplateOpts Optional parameters for the method 'GetGlAccountsTemplate'
type GetGlAccountsTemplateOpts struct {
	DateFormat optional.String
}

/*
GetGlAccountsTemplate Method for GetGlAccountsTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetGlAccountsTemplateOpts - Optional Parameters:
 * @param "DateFormat" (optional.String) -
*/
func (a *GeneralLedgerAccountApiService) GetGlAccountsTemplate(ctx _context.Context, opts *GetGlAccountsTemplateOpts) (*_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/glaccounts/downloadtemplate"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.DateFormat.IsSet() {
		queryParams.Add("dateFormat", parameterToString(opts.DateFormat.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/vnd.ms-excel"}

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
		return nil, err
	}

	res, err := a.client.callAPI(r)
	if err != nil || res == nil {
		return res, err
	}

	payload, err := _ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return res, err
	}

	if res.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  payload,
			error: res.Status,
		}
		return res, newErr
	}

	return res, nil
}

// PostGlAccountsTemplateOpts Optional parameters for the method 'PostGlAccountsTemplate'
type PostGlAccountsTemplateOpts struct {
	File       optional.Interface
	Locale     optional.String
	DateFormat optional.String
}

/*
PostGlAccountsTemplate Method for PostGlAccountsTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *PostGlAccountsTemplateOpts - Optional Parameters:
 * @param "File" (optional.Interface of FormDataContentDisposition) -
 * @param "Locale" (optional.String) -
 * @param "DateFormat" (optional.String) -
@return string
*/
func (a *GeneralLedgerAccountApiService) PostGlAccountsTemplate(ctx _context.Context, opts *PostGlAccountsTemplateOpts) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/glaccounts/uploadtemplate"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"*/*"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	if opts != nil && opts.File.IsSet() {
		paramJson, err := parameterToJson(opts.File.Value())
		if err != nil {
			return returnValue, nil, err
		}
		formParams.Add("file", paramJson)
	}
	if opts != nil && opts.Locale.IsSet() {
		formParams.Add("locale", parameterToString(opts.Locale.Value(), ""))
	}
	if opts != nil && opts.DateFormat.IsSet() {
		formParams.Add("dateFormat", parameterToString(opts.DateFormat.Value(), ""))
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

// RetreiveAccountOpts Optional parameters for the method 'RetreiveAccount'
type RetreiveAccountOpts struct {
	FetchRunningBalance optional.Bool
}

/*
RetreiveAccount Retrieve a General Ledger Account
Example Requests:  glaccounts/1   glaccounts/1?template&#x3D;true   glaccounts/1?fields&#x3D;name,glCode   glaccounts/1?fetchRunningBalance&#x3D;true
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param glAccountId glAccountId
 * @param optional nil or *RetreiveAccountOpts - Optional Parameters:
 * @param "FetchRunningBalance" (optional.Bool) -  fetchRunningBalance
@return GetGlAccountsResponse
*/
func (a *GeneralLedgerAccountApiService) RetreiveAccount(ctx _context.Context, glAccountId int64, opts *RetreiveAccountOpts) (models.GetGlAccountsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetGlAccountsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/glaccounts/{glAccountId}"
	path = strings.Replace(path, "{"+"glAccountId"+"}", _neturl.QueryEscape(parameterToString(glAccountId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.FetchRunningBalance.IsSet() {
		queryParams.Add("fetchRunningBalance", parameterToString(opts.FetchRunningBalance.Value(), ""))
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

// RetrieveAllAccountsOpts Optional parameters for the method 'RetrieveAllAccounts'
type RetrieveAllAccountsOpts struct {
	Type_                optional.Int32
	SearchParam          optional.String
	Usage                optional.Int32
	ManualEntriesAllowed optional.Bool
	Disabled             optional.Bool
	FetchRunningBalance  optional.Bool
}

/*
RetrieveAllAccounts List General Ledger Accounts
ARGUMENTS type Integer optional manualEntriesAllowed boolean optional usage Integer optional disabled boolean optional parentId Long optional tagId Long optional Example Requests:  glaccounts   glaccounts?type&#x3D;1&amp;manualEntriesAllowed&#x3D;true&amp;usage&#x3D;1&amp;disabled&#x3D;false  glaccounts?fetchRunningBalance&#x3D;true
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveAllAccountsOpts - Optional Parameters:
 * @param "Type_" (optional.Int32) -  type
 * @param "SearchParam" (optional.String) -  searchParam
 * @param "Usage" (optional.Int32) -  usage
 * @param "ManualEntriesAllowed" (optional.Bool) -  manualEntriesAllowed
 * @param "Disabled" (optional.Bool) -  disabled
 * @param "FetchRunningBalance" (optional.Bool) -  fetchRunningBalance
@return []GetGlAccountsResponse
*/
func (a *GeneralLedgerAccountApiService) RetrieveAllAccounts(ctx _context.Context, opts *RetrieveAllAccountsOpts) ([]models.GetGlAccountsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  []models.GetGlAccountsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/glaccounts"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.Type_.IsSet() {
		queryParams.Add("type", parameterToString(opts.Type_.Value(), ""))
	}
	if opts != nil && opts.SearchParam.IsSet() {
		queryParams.Add("searchParam", parameterToString(opts.SearchParam.Value(), ""))
	}
	if opts != nil && opts.Usage.IsSet() {
		queryParams.Add("usage", parameterToString(opts.Usage.Value(), ""))
	}
	if opts != nil && opts.ManualEntriesAllowed.IsSet() {
		queryParams.Add("manualEntriesAllowed", parameterToString(opts.ManualEntriesAllowed.Value(), ""))
	}
	if opts != nil && opts.Disabled.IsSet() {
		queryParams.Add("disabled", parameterToString(opts.Disabled.Value(), ""))
	}
	if opts != nil && opts.FetchRunningBalance.IsSet() {
		queryParams.Add("fetchRunningBalance", parameterToString(opts.FetchRunningBalance.Value(), ""))
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

// RetrieveNewAccountDetailsOpts Optional parameters for the method 'RetrieveNewAccountDetails'
type RetrieveNewAccountDetailsOpts struct {
	Type_ optional.Int32
}

/*
RetrieveNewAccountDetails Retrieve GL Accounts Template
This is a convenience resource. It can be useful when building maintenance user interface screens for client applications. The template data returned consists of any or all of:  Field Defaults Allowed Value Lists Example Request:  glaccounts/template glaccounts/template?type&#x3D;1  type is optional and integer value from 1 to 5.  1.Assets  2.Liabilities  3.Equity  4.Income  5.Expenses
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveNewAccountDetailsOpts - Optional Parameters:
 * @param "Type_" (optional.Int32) -  type
@return GetGlAccountsTemplateResponse
*/
func (a *GeneralLedgerAccountApiService) RetrieveNewAccountDetails(ctx _context.Context, opts *RetrieveNewAccountDetailsOpts) (models.GetGlAccountsTemplateResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetGlAccountsTemplateResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/glaccounts/template"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.Type_.IsSet() {
		queryParams.Add("type", parameterToString(opts.Type_.Value(), ""))
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

// UpdateGLAccount1Opts Optional parameters for the method 'UpdateGLAccount1'
type UpdateGLAccount1Opts struct {
	PutGlAccountsRequest optional.Interface
}

/*
UpdateGLAccount1 Update an Accounting closure
Once an accounting closure is created, only the comments associated with it may be edited
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param glAccountId glAccountId
 * @param optional nil or *UpdateGLAccount1Opts - Optional Parameters:
 * @param "PutGlAccountsRequest" (optional.Interface of PutGlAccountsRequest) -
@return PutGlAccountsResponse
*/
func (a *GeneralLedgerAccountApiService) UpdateGLAccount1(ctx _context.Context, glAccountId int64, opts *UpdateGLAccount1Opts) (models.PutGlAccountsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPut
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PutGlAccountsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/glaccounts/{glAccountId}"
	path = strings.Replace(path, "{"+"glAccountId"+"}", _neturl.QueryEscape(parameterToString(glAccountId, "")), -1)

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
	if opts != nil && opts.PutGlAccountsRequest.IsSet() {
		optPutGlAccountsRequest, optPutGlAccountsRequestok := opts.PutGlAccountsRequest.Value().(models.PutGlAccountsRequest)
		if !optPutGlAccountsRequestok {
			return returnValue, nil, reportError("putGlAccountsRequest should be PutGlAccountsRequest")
		}
		postBody = &optPutGlAccountsRequest
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
