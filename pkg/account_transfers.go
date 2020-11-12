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

// AccountTransfersApiService AccountTransfersApi service
type AccountTransfersApiService service

/*
Create3 Create new Transfer
Ability to create new transfer of monetary funds from one account to another.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param postAccountTransfersRequest
@return PostAccountTransfersResponse
*/
func (a *AccountTransfersApiService) Create3(ctx _context.Context, postAccountTransfersRequest models.PostAccountTransfersRequest) (models.PostAccountTransfersResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostAccountTransfersResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/accounttransfers"
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
	postBody = &postAccountTransfersRequest
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

// RetrieveAll17Opts Optional parameters for the method 'RetrieveAll17'
type RetrieveAll17Opts struct {
	SqlSearch       optional.String
	ExternalId      optional.String
	Offset          optional.Int32
	Limit           optional.Int32
	OrderBy         optional.String
	SortOrder       optional.String
	AccountDetailId optional.Int64
}

/*
RetrieveAll17 List account transfers
Lists account&#39;s transfers  Example Requests:    accounttransfers
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveAll17Opts - Optional Parameters:
 * @param "SqlSearch" (optional.String) -  sqlSearch
 * @param "ExternalId" (optional.String) -  externalId
 * @param "Offset" (optional.Int32) -  offset
 * @param "Limit" (optional.Int32) -
 * @param "OrderBy" (optional.String) -  orderBy
 * @param "SortOrder" (optional.String) -  sortOrder
 * @param "AccountDetailId" (optional.Int64) -  accountDetailId
@return GetAccountTransfersResponse
*/
func (a *AccountTransfersApiService) RetrieveAll17(ctx _context.Context, opts *RetrieveAll17Opts) (models.GetAccountTransfersResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetAccountTransfersResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/accounttransfers"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.SqlSearch.IsSet() {
		queryParams.Add("sqlSearch", parameterToString(opts.SqlSearch.Value(), ""))
	}
	if opts != nil && opts.ExternalId.IsSet() {
		queryParams.Add("externalId", parameterToString(opts.ExternalId.Value(), ""))
	}
	if opts != nil && opts.Offset.IsSet() {
		queryParams.Add("offset", parameterToString(opts.Offset.Value(), ""))
	}
	if opts != nil && opts.Limit.IsSet() {
		queryParams.Add("limit", parameterToString(opts.Limit.Value(), ""))
	}
	if opts != nil && opts.OrderBy.IsSet() {
		queryParams.Add("orderBy", parameterToString(opts.OrderBy.Value(), ""))
	}
	if opts != nil && opts.SortOrder.IsSet() {
		queryParams.Add("sortOrder", parameterToString(opts.SortOrder.Value(), ""))
	}
	if opts != nil && opts.AccountDetailId.IsSet() {
		queryParams.Add("accountDetailId", parameterToString(opts.AccountDetailId.Value(), ""))
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
RetrieveOne8 Retrieve account transfer
Retrieves account transfer  Example Requests :    accounttransfers/1
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param transferId transferId
@return GetAccountTransfersPageItems
*/
func (a *AccountTransfersApiService) RetrieveOne8(ctx _context.Context, transferId int64) (models.GetAccountTransfersPageItems, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetAccountTransfersPageItems
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/accounttransfers/{transferId}"
	path = strings.Replace(path, "{"+"transferId"+"}", _neturl.QueryEscape(parameterToString(transferId, "")), -1)

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

// Template5Opts Optional parameters for the method 'Template5'
type Template5Opts struct {
	FromOfficeId    optional.Int64
	FromClientId    optional.Int64
	FromAccountId   optional.Int64
	FromAccountType optional.Int32
	ToOfficeId      optional.Int64
	ToClientId      optional.Int64
	ToAccountId     optional.Int64
	ToAccountType   optional.Int32
}

/*
Template5 Retrieve Account Transfer Template
This is a convenience resource. It can be useful when building maintenance user interface screens for client applications. The template data returned consists of any or all of:    Field Defaults  Allowed Value Lists  Example Requests:    accounttransfers/template?fromAccountType&#x3D;2&amp;fromOfficeId&#x3D;1    accounttransfers/template?fromAccountType&#x3D;2&amp;fromOfficeId&#x3D;1&amp;fromClientId&#x3D;1    accounttransfers/template?fromClientId&#x3D;1&amp;fromAccountType&#x3D;2&amp;fromAccountId&#x3D;1
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *Template5Opts - Optional Parameters:
 * @param "FromOfficeId" (optional.Int64) -  fromOfficeId
 * @param "FromClientId" (optional.Int64) -  fromClientId
 * @param "FromAccountId" (optional.Int64) -  fromAccountId
 * @param "FromAccountType" (optional.Int32) -  fromAccountType
 * @param "ToOfficeId" (optional.Int64) -  toOfficeId
 * @param "ToClientId" (optional.Int64) -  toClientId
 * @param "ToAccountId" (optional.Int64) -  toAccountId
 * @param "ToAccountType" (optional.Int32) -  toAccountType
@return GetAccountTransfersTemplateResponse
*/
func (a *AccountTransfersApiService) Template5(ctx _context.Context, opts *Template5Opts) (models.GetAccountTransfersTemplateResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetAccountTransfersTemplateResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/accounttransfers/template"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.FromOfficeId.IsSet() {
		queryParams.Add("fromOfficeId", parameterToString(opts.FromOfficeId.Value(), ""))
	}
	if opts != nil && opts.FromClientId.IsSet() {
		queryParams.Add("fromClientId", parameterToString(opts.FromClientId.Value(), ""))
	}
	if opts != nil && opts.FromAccountId.IsSet() {
		queryParams.Add("fromAccountId", parameterToString(opts.FromAccountId.Value(), ""))
	}
	if opts != nil && opts.FromAccountType.IsSet() {
		queryParams.Add("fromAccountType", parameterToString(opts.FromAccountType.Value(), ""))
	}
	if opts != nil && opts.ToOfficeId.IsSet() {
		queryParams.Add("toOfficeId", parameterToString(opts.ToOfficeId.Value(), ""))
	}
	if opts != nil && opts.ToClientId.IsSet() {
		queryParams.Add("toClientId", parameterToString(opts.ToClientId.Value(), ""))
	}
	if opts != nil && opts.ToAccountId.IsSet() {
		queryParams.Add("toAccountId", parameterToString(opts.ToAccountId.Value(), ""))
	}
	if opts != nil && opts.ToAccountType.IsSet() {
		queryParams.Add("toAccountType", parameterToString(opts.ToAccountType.Value(), ""))
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

// TemplateRefundByTransferOpts Optional parameters for the method 'TemplateRefundByTransfer'
type TemplateRefundByTransferOpts struct {
	FromOfficeId    optional.Int64
	FromClientId    optional.Int64
	FromAccountId   optional.Int64
	FromAccountType optional.Int32
	ToOfficeId      optional.Int64
	ToClientId      optional.Int64
	ToAccountId     optional.Int64
	ToAccountType   optional.Int32
}

/*
TemplateRefundByTransfer Retrieve Refund of an Active Loan by Transfer Template
Retrieves Refund of an Active Loan by Transfer TemplateExample Requests :    accounttransfers/templateRefundByTransfer?fromAccountId&#x3D;2&amp;fromAccountType&#x3D;1&amp; toAccountId&#x3D;1&amp;toAccountType&#x3D;2&amp;toClientId&#x3D;1&amp;toOfficeId&#x3D;1
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *TemplateRefundByTransferOpts - Optional Parameters:
 * @param "FromOfficeId" (optional.Int64) -  fromOfficeId
 * @param "FromClientId" (optional.Int64) -  fromClientId
 * @param "FromAccountId" (optional.Int64) -  fromAccountId
 * @param "FromAccountType" (optional.Int32) -  fromAccountType
 * @param "ToOfficeId" (optional.Int64) -  toOfficeId
 * @param "ToClientId" (optional.Int64) -  toClientId
 * @param "ToAccountId" (optional.Int64) -  toAccountId
 * @param "ToAccountType" (optional.Int32) -  toAccountType
@return GetAccountTransfersTemplateRefundByTransferResponse
*/
func (a *AccountTransfersApiService) TemplateRefundByTransfer(ctx _context.Context, opts *TemplateRefundByTransferOpts) (models.GetAccountTransfersTemplateRefundByTransferResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetAccountTransfersTemplateRefundByTransferResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/accounttransfers/templateRefundByTransfer"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.FromOfficeId.IsSet() {
		queryParams.Add("fromOfficeId", parameterToString(opts.FromOfficeId.Value(), ""))
	}
	if opts != nil && opts.FromClientId.IsSet() {
		queryParams.Add("fromClientId", parameterToString(opts.FromClientId.Value(), ""))
	}
	if opts != nil && opts.FromAccountId.IsSet() {
		queryParams.Add("fromAccountId", parameterToString(opts.FromAccountId.Value(), ""))
	}
	if opts != nil && opts.FromAccountType.IsSet() {
		queryParams.Add("fromAccountType", parameterToString(opts.FromAccountType.Value(), ""))
	}
	if opts != nil && opts.ToOfficeId.IsSet() {
		queryParams.Add("toOfficeId", parameterToString(opts.ToOfficeId.Value(), ""))
	}
	if opts != nil && opts.ToClientId.IsSet() {
		queryParams.Add("toClientId", parameterToString(opts.ToClientId.Value(), ""))
	}
	if opts != nil && opts.ToAccountId.IsSet() {
		queryParams.Add("toAccountId", parameterToString(opts.ToAccountId.Value(), ""))
	}
	if opts != nil && opts.ToAccountType.IsSet() {
		queryParams.Add("toAccountType", parameterToString(opts.ToAccountType.Value(), ""))
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
TemplateRefundByTransferPost Refund of an Active Loan by Transfer
Ability to refund an active loan by transferring to a savings account.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param postAccountTransfersRefundByTransferRequest
@return PostAccountTransfersRefundByTransferResponse
*/
func (a *AccountTransfersApiService) TemplateRefundByTransferPost(ctx _context.Context, postAccountTransfersRefundByTransferRequest models.PostAccountTransfersRefundByTransferRequest) (models.PostAccountTransfersRefundByTransferResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostAccountTransfersRefundByTransferResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/accounttransfers/refundByTransfer"
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
	postBody = &postAccountTransfersRefundByTransferRequest
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
