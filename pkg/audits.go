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

// AuditsApiService AuditsApi service
type AuditsApiService service

// RetrieveAuditEntriesOpts Optional parameters for the method 'RetrieveAuditEntries'
type RetrieveAuditEntriesOpts struct {
	ActionName          optional.String
	EntityName          optional.String
	ResourceId          optional.Int64
	MakerId             optional.Int64
	MakerDateTimeFrom   optional.String
	MakerDateTimeTo     optional.String
	CheckerId           optional.Int64
	CheckerDateTimeFrom optional.String
	CheckerDateTimeTo   optional.String
	ProcessingResult    optional.Int32
	OfficeId            optional.Int32
	GroupId             optional.Int32
	ClientId            optional.Int32
	Loanid              optional.Int32
	SavingsAccountId    optional.Int32
	Paged               optional.Bool
	Offset              optional.Int32
	Limit               optional.Int32
	OrderBy             optional.String
	SortOrder           optional.String
}

/*
RetrieveAuditEntries List Audits
Get a 200 list of audits that match the criteria supplied and sorted by audit id in descending order, and are within the requestors&#39; data scope. Also it supports pagination and sorting  Example Requests:  audits  audits?fields&#x3D;madeOnDate,maker,processingResult  audits?makerDateTimeFrom&#x3D;2013-03-25 08:00:00&amp;makerDateTimeTo&#x3D;2013-04-04 18:00:00  audits?officeId&#x3D;1  audits?officeId&#x3D;1&amp;includeJson&#x3D;true
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveAuditEntriesOpts - Optional Parameters:
 * @param "ActionName" (optional.String) -  actionName
 * @param "EntityName" (optional.String) -  entityName
 * @param "ResourceId" (optional.Int64) -  resourceId
 * @param "MakerId" (optional.Int64) -  makerId
 * @param "MakerDateTimeFrom" (optional.String) -  makerDateTimeFrom
 * @param "MakerDateTimeTo" (optional.String) -  makerDateTimeTo
 * @param "CheckerId" (optional.Int64) -  checkerId
 * @param "CheckerDateTimeFrom" (optional.String) -  checkerDateTimeFrom
 * @param "CheckerDateTimeTo" (optional.String) -  checkerDateTimeTo
 * @param "ProcessingResult" (optional.Int32) -  processingResult
 * @param "OfficeId" (optional.Int32) -  officeId
 * @param "GroupId" (optional.Int32) -  groupId
 * @param "ClientId" (optional.Int32) -  clientId
 * @param "Loanid" (optional.Int32) -  loanid
 * @param "SavingsAccountId" (optional.Int32) -  savingsAccountId
 * @param "Paged" (optional.Bool) -  paged
 * @param "Offset" (optional.Int32) -  offset
 * @param "Limit" (optional.Int32) -  limit
 * @param "OrderBy" (optional.String) -  orderBy
 * @param "SortOrder" (optional.String) -  sortOrder
@return []GetMakerCheckerResponse
*/
func (a *AuditsApiService) RetrieveAuditEntries(ctx _context.Context, opts *RetrieveAuditEntriesOpts) ([]models.GetMakerCheckerResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  []models.GetMakerCheckerResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/audits"
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
	if opts != nil && opts.CheckerId.IsSet() {
		queryParams.Add("checkerId", parameterToString(opts.CheckerId.Value(), ""))
	}
	if opts != nil && opts.CheckerDateTimeFrom.IsSet() {
		queryParams.Add("checkerDateTimeFrom", parameterToString(opts.CheckerDateTimeFrom.Value(), ""))
	}
	if opts != nil && opts.CheckerDateTimeTo.IsSet() {
		queryParams.Add("checkerDateTimeTo", parameterToString(opts.CheckerDateTimeTo.Value(), ""))
	}
	if opts != nil && opts.ProcessingResult.IsSet() {
		queryParams.Add("processingResult", parameterToString(opts.ProcessingResult.Value(), ""))
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
	if opts != nil && opts.Paged.IsSet() {
		queryParams.Add("paged", parameterToString(opts.Paged.Value(), ""))
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

/*
RetrieveAuditEntry Retrieve an Audit Entry
Example Requests:  audits/20 audits/20?fields&#x3D;madeOnDate,maker,processingResult
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param auditId auditId
@return GetMakerCheckerResponse
*/
func (a *AuditsApiService) RetrieveAuditEntry(ctx _context.Context, auditId int64) (models.GetMakerCheckerResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetMakerCheckerResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/audits/{auditId}"
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
RetrieveAuditSearchTemplate Audit Search Template
This is a convenience resource. It can be useful when building an Audit Search UI. \&quot;appUsers\&quot; are data scoped to the office/branch the requestor is associated with.  Example Requests:  audits/searchtemplate audits/searchtemplate?fields&#x3D;actionNames
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return GetMakerCheckersSearchTemplateResponse
*/
func (a *AuditsApiService) RetrieveAuditSearchTemplate(ctx _context.Context) (models.GetMakerCheckersSearchTemplateResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetMakerCheckersSearchTemplateResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/audits/searchtemplate"
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
