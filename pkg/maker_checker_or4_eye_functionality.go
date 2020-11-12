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

// MakerCheckerOr4EyeFunctionalityApiService MakerCheckerOr4EyeFunctionalityApi service
type MakerCheckerOr4EyeFunctionalityApiService service

// ApproveMakerCheckerEntryOpts Optional parameters for the method 'ApproveMakerCheckerEntry'
type ApproveMakerCheckerEntryOpts struct {
	Command optional.String
}

/*
ApproveMakerCheckerEntry Approve Maker Checker Entry | Reject Maker Checker Entry
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param auditId auditId
 * @param optional nil or *ApproveMakerCheckerEntryOpts - Optional Parameters:
 * @param "Command" (optional.String) -  command
@return PostMakerCheckersResponse
*/
func (a *MakerCheckerOr4EyeFunctionalityApiService) ApproveMakerCheckerEntry(ctx _context.Context, auditId int64, opts *ApproveMakerCheckerEntryOpts) (models.PostMakerCheckersResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostMakerCheckersResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/makercheckers/{auditId}"
	path = strings.Replace(path, "{"+"auditId"+"}", _neturl.QueryEscape(parameterToString(auditId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.Command.IsSet() {
		queryParams.Add("command", parameterToString(opts.Command.Value(), ""))
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
DeleteMakerCheckerEntry Delete Maker Checker Entry
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param auditId auditId
@return PostMakerCheckersResponse
*/
func (a *MakerCheckerOr4EyeFunctionalityApiService) DeleteMakerCheckerEntry(ctx _context.Context, auditId int64) (models.PostMakerCheckersResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodDelete
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostMakerCheckersResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/makercheckers/{auditId}"
	path = strings.Replace(path, "{"+"auditId"+"}", _neturl.QueryEscape(parameterToString(auditId, "")), -1)

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
RetrieveAuditSearchTemplate1 Maker Checker Search Template
This is a convenience resource. It can be useful when building a Checker Inbox UI. \&quot;appUsers\&quot; are data scoped to the office/branch the requestor is associated with. \&quot;actionNames\&quot; and \&quot;entityNames\&quot; returned are those that the requestor has Checker approval permissions for.  Example Requests:  makercheckers/searchtemplate makercheckers/searchtemplate?fields&#x3D;entityNames
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return GetMakerCheckersSearchTemplateResponse
*/
func (a *MakerCheckerOr4EyeFunctionalityApiService) RetrieveAuditSearchTemplate1(ctx _context.Context) (models.GetMakerCheckersSearchTemplateResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetMakerCheckersSearchTemplateResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/makercheckers/searchtemplate"
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

// RetrieveCommandsOpts Optional parameters for the method 'RetrieveCommands'
type RetrieveCommandsOpts struct {
	ActionName        optional.String
	EntityName        optional.String
	ResourceId        optional.Int64
	MakerId           optional.Int64
	MakerDateTimeFrom optional.String
	MakerDateTimeTo   optional.String
	OfficeId          optional.Int32
	GroupId           optional.Int32
	ClientId          optional.Int32
	Loanid            optional.Int32
	SavingsAccountId  optional.Int32
}

/*
RetrieveCommands List Maker Checker Entries
Get a list of entries that can be checked by the requestor that match the criteria supplied.  Example Requests:  makercheckers  makercheckers?fields&#x3D;madeOnDate,maker,processingResult  makercheckers?makerDateTimeFrom&#x3D;2013-03-25 08:00:00&amp;makerDateTimeTo&#x3D;2013-04-04 18:00:00  makercheckers?officeId&#x3D;1  makercheckers?officeId&#x3D;1&amp;includeJson&#x3D;true
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveCommandsOpts - Optional Parameters:
 * @param "ActionName" (optional.String) -  actionName
 * @param "EntityName" (optional.String) -  entityName
 * @param "ResourceId" (optional.Int64) -  resourceId
 * @param "MakerId" (optional.Int64) -  makerId
 * @param "MakerDateTimeFrom" (optional.String) -  makerDateTimeFrom
 * @param "MakerDateTimeTo" (optional.String) -  makerDateTimeTo
 * @param "OfficeId" (optional.Int32) -  officeId
 * @param "GroupId" (optional.Int32) -  groupId
 * @param "ClientId" (optional.Int32) -  clientId
 * @param "Loanid" (optional.Int32) -  loanid
 * @param "SavingsAccountId" (optional.Int32) -  savingsAccountId
@return []GetMakerCheckerResponse
*/
func (a *MakerCheckerOr4EyeFunctionalityApiService) RetrieveCommands(ctx _context.Context, opts *RetrieveCommandsOpts) ([]models.GetMakerCheckerResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  []models.GetMakerCheckerResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/makercheckers"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.ActionName.IsSet() {
		queryParams.Add("actionName", parameterToString(opts.ActionName.Value(), ""))
	}
	if opts != nil && opts.EntityName.IsSet() {
		queryParams.Add("entityName", parameterToString(opts.EntityName.Value(), ""))
	}
	if opts != nil && opts.ResourceId.IsSet() {
		queryParams.Add("resourceId", parameterToString(opts.ResourceId.Value(), ""))
	}
	if opts != nil && opts.MakerId.IsSet() {
		queryParams.Add("makerId", parameterToString(opts.MakerId.Value(), ""))
	}
	if opts != nil && opts.MakerDateTimeFrom.IsSet() {
		queryParams.Add("makerDateTimeFrom", parameterToString(opts.MakerDateTimeFrom.Value(), ""))
	}
	if opts != nil && opts.MakerDateTimeTo.IsSet() {
		queryParams.Add("makerDateTimeTo", parameterToString(opts.MakerDateTimeTo.Value(), ""))
	}
	if opts != nil && opts.OfficeId.IsSet() {
		queryParams.Add("officeId", parameterToString(opts.OfficeId.Value(), ""))
	}
	if opts != nil && opts.GroupId.IsSet() {
		queryParams.Add("groupId", parameterToString(opts.GroupId.Value(), ""))
	}
	if opts != nil && opts.ClientId.IsSet() {
		queryParams.Add("clientId", parameterToString(opts.ClientId.Value(), ""))
	}
	if opts != nil && opts.Loanid.IsSet() {
		queryParams.Add("loanid", parameterToString(opts.Loanid.Value(), ""))
	}
	if opts != nil && opts.SavingsAccountId.IsSet() {
		queryParams.Add("savingsAccountId", parameterToString(opts.SavingsAccountId.Value(), ""))
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
