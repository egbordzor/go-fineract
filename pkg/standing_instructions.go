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

// StandingInstructionsApiService StandingInstructionsApi service
type StandingInstructionsApiService service

/*
Create4 Create new Standing Instruction
Ability to create new instruction for transfer of monetary funds from one account to another
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param postStandingInstructionsRequest
@return PostStandingInstructionsResponse
*/
func (a *StandingInstructionsApiService) Create4(ctx _context.Context, postStandingInstructionsRequest models.PostStandingInstructionsRequest) (models.PostStandingInstructionsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostStandingInstructionsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/standinginstructions"
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
	postBody = &postStandingInstructionsRequest
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

// RetrieveAll18Opts Optional parameters for the method 'RetrieveAll18'
type RetrieveAll18Opts struct {
	SqlSearch       optional.String
	ExternalId      optional.String
	Offset          optional.Int32
	Limit           optional.Int32
	OrderBy         optional.String
	SortOrder       optional.String
	TransferType    optional.Int32
	ClientName      optional.String
	ClientId        optional.Int64
	FromAccountId   optional.Int64
	FromAccountType optional.Int32
}

/*
RetrieveAll18 List Standing Instructions
Example Requests:  standinginstructions
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveAll18Opts - Optional Parameters:
 * @param "SqlSearch" (optional.String) -  sqlSearch
 * @param "ExternalId" (optional.String) -  externalId
 * @param "Offset" (optional.Int32) -  offset
 * @param "Limit" (optional.Int32) -  limit
 * @param "OrderBy" (optional.String) -  orderBy
 * @param "SortOrder" (optional.String) -  sortOrder
 * @param "TransferType" (optional.Int32) -  transferType
 * @param "ClientName" (optional.String) -  clientName
 * @param "ClientId" (optional.Int64) -  clientId
 * @param "FromAccountId" (optional.Int64) -  fromAccountId
 * @param "FromAccountType" (optional.Int32) -  fromAccountType
@return GetStandingInstructionsResponse
*/
func (a *StandingInstructionsApiService) RetrieveAll18(ctx _context.Context, opts *RetrieveAll18Opts) (models.GetStandingInstructionsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetStandingInstructionsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/standinginstructions"
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
	if opts != nil && opts.TransferType.IsSet() {
		queryParams.Add("transferType", parameterToString(opts.TransferType.Value(), ""))
	}
	if opts != nil && opts.ClientName.IsSet() {
		queryParams.Add("clientName", parameterToString(opts.ClientName.Value(), ""))
	}
	if opts != nil && opts.ClientId.IsSet() {
		queryParams.Add("clientId", parameterToString(opts.ClientId.Value(), ""))
	}
	if opts != nil && opts.FromAccountId.IsSet() {
		queryParams.Add("fromAccountId", parameterToString(opts.FromAccountId.Value(), ""))
	}
	if opts != nil && opts.FromAccountType.IsSet() {
		queryParams.Add("fromAccountType", parameterToString(opts.FromAccountType.Value(), ""))
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

// RetrieveOne9Opts Optional parameters for the method 'RetrieveOne9'
type RetrieveOne9Opts struct {
	SqlSearch  optional.String
	ExternalId optional.String
	Offset     optional.Int32
	Limit      optional.Int32
	OrderBy    optional.String
	SortOrder  optional.String
}

/*
RetrieveOne9 Retrieve Standing Instruction
Example Requests :  standinginstructions/1
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param standingInstructionId standingInstructionId
 * @param optional nil or *RetrieveOne9Opts - Optional Parameters:
 * @param "SqlSearch" (optional.String) -  sqlSearch
 * @param "ExternalId" (optional.String) -  externalId
 * @param "Offset" (optional.Int32) -  offset
 * @param "Limit" (optional.Int32) -  limit
 * @param "OrderBy" (optional.String) -  orderBy
 * @param "SortOrder" (optional.String) -  sortOrder
@return GetStandingInstructionsStandingInstructionIdResponse
*/
func (a *StandingInstructionsApiService) RetrieveOne9(ctx _context.Context, standingInstructionId int64, opts *RetrieveOne9Opts) (models.GetStandingInstructionsStandingInstructionIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetStandingInstructionsStandingInstructionIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/standinginstructions/{standingInstructionId}"
	path = strings.Replace(path, "{"+"standingInstructionId"+"}", _neturl.QueryEscape(parameterToString(standingInstructionId, "")), -1)

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

// Template6Opts Optional parameters for the method 'Template6'
type Template6Opts struct {
	FromOfficeId    optional.Int64
	FromClientId    optional.Int64
	FromAccountId   optional.Int64
	FromAccountType optional.Int32
	ToOfficeId      optional.Int64
	ToClientId      optional.Int64
	ToAccountId     optional.Int64
	ToAccountType   optional.Int32
	TransferType    optional.Int32
}

/*
Template6 Retrieve Standing Instruction Template
This is a convenience resource. It can be useful when building maintenance user interface screens for client applications. The template data returned consists of any or all of:  Field Defaults Allowed Value Lists Example Requests:  standinginstructions/template?fromAccountType&#x3D;2&amp;fromOfficeId&#x3D;1  standinginstructions/template?fromAccountType&#x3D;2&amp;fromOfficeId&#x3D;1&amp;fromClientId&#x3D;1&amp;transferType&#x3D;1  standinginstructions/template?fromClientId&#x3D;1&amp;fromAccountType&#x3D;2&amp;fromAccountId&#x3D;1&amp;transferType&#x3D;1
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *Template6Opts - Optional Parameters:
 * @param "FromOfficeId" (optional.Int64) -  fromOfficeId
 * @param "FromClientId" (optional.Int64) -  fromClientId
 * @param "FromAccountId" (optional.Int64) -  fromAccountId
 * @param "FromAccountType" (optional.Int32) -  fromAccountType
 * @param "ToOfficeId" (optional.Int64) -  toOfficeId
 * @param "ToClientId" (optional.Int64) -  toClientId
 * @param "ToAccountId" (optional.Int64) -  toAccountId
 * @param "ToAccountType" (optional.Int32) -  toAccountType
 * @param "TransferType" (optional.Int32) -  transferType
@return GetStandingInstructionsTemplateResponse
*/
func (a *StandingInstructionsApiService) Template6(ctx _context.Context, opts *Template6Opts) (models.GetStandingInstructionsTemplateResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetStandingInstructionsTemplateResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/standinginstructions/template"
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
	if opts != nil && opts.TransferType.IsSet() {
		queryParams.Add("transferType", parameterToString(opts.TransferType.Value(), ""))
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

// Update8Opts Optional parameters for the method 'Update8'
type Update8Opts struct {
	Command                                             optional.String
	PutStandingInstructionsStandingInstructionIdRequest optional.Interface
}

/*
Update8 Update Standing Instruction | Delete Standing Instruction
Ability to modify existing instruction for transfer of monetary funds from one account to another.  PUT https://DomainName/api/v1/standinginstructions/1?command&#x3D;update   Ability to modify existing instruction for transfer of monetary funds from one account to another.  PUT https://DomainName/api/v1/standinginstructions/1?command&#x3D;delete
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param standingInstructionId standingInstructionId
 * @param optional nil or *Update8Opts - Optional Parameters:
 * @param "Command" (optional.String) -  command
 * @param "PutStandingInstructionsStandingInstructionIdRequest" (optional.Interface of PutStandingInstructionsStandingInstructionIdRequest) -
@return PutStandingInstructionsStandingInstructionIdResponse
*/
func (a *StandingInstructionsApiService) Update8(ctx _context.Context, standingInstructionId int64, opts *Update8Opts) (models.PutStandingInstructionsStandingInstructionIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPut
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PutStandingInstructionsStandingInstructionIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/standinginstructions/{standingInstructionId}"
	path = strings.Replace(path, "{"+"standingInstructionId"+"}", _neturl.QueryEscape(parameterToString(standingInstructionId, "")), -1)

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
	if opts != nil && opts.PutStandingInstructionsStandingInstructionIdRequest.IsSet() {
		optPutStandingInstructionsStandingInstructionIdRequest, optPutStandingInstructionsStandingInstructionIdRequestok := opts.PutStandingInstructionsStandingInstructionIdRequest.Value().(models.PutStandingInstructionsStandingInstructionIdRequest)
		if !optPutStandingInstructionsStandingInstructionIdRequestok {
			return returnValue, nil, reportError("putStandingInstructionsStandingInstructionIdRequest should be PutStandingInstructionsStandingInstructionIdRequest")
		}
		postBody = &optPutStandingInstructionsStandingInstructionIdRequest
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
