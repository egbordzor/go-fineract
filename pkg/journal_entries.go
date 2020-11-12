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

// JournalEntriesApiService JournalEntriesApi service
type JournalEntriesApiService service

// CreateGLJournalEntryOpts Optional parameters for the method 'CreateGLJournalEntry'
type CreateGLJournalEntryOpts struct {
	Command             optional.String
	JournalEntryCommand optional.Interface
}

/*
CreateGLJournalEntry Create \"Balanced\" Journal Entries
Note: A Balanced (simple) Journal entry would have atleast one \&quot;Debit\&quot; and one \&quot;Credit\&quot; entry whose amounts are equal  Compound Journal entries may have \&quot;n\&quot; debits and \&quot;m\&quot; credits where both \&quot;m\&quot; and \&quot;n\&quot; are greater than 0 and the net sum or all debits and credits are equal    Mandatory Fields officeId, transactionDate   credits- glAccountId, amount, comments    debits-  glAccountId, amount, comments    Optional Fields paymentTypeId, accountNumber, checkNumber, routingCode, receiptNumber, bankNumber
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CreateGLJournalEntryOpts - Optional Parameters:
 * @param "Command" (optional.String) -  command
 * @param "JournalEntryCommand" (optional.Interface of JournalEntryCommand) -
@return PostJournalEntriesResponse
*/
func (a *JournalEntriesApiService) CreateGLJournalEntry(ctx _context.Context, opts *CreateGLJournalEntryOpts) (models.PostJournalEntriesResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostJournalEntriesResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/journalentries"
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
	if opts != nil && opts.JournalEntryCommand.IsSet() {
		optJournalEntryCommand, optJournalEntryCommandok := opts.JournalEntryCommand.Value().(models.JournalEntryCommand)
		if !optJournalEntryCommandok {
			return returnValue, nil, reportError("journalEntryCommand should be JournalEntryCommand")
		}
		postBody = &optJournalEntryCommand
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

// CreateReversalJournalEntryOpts Optional parameters for the method 'CreateReversalJournalEntry'
type CreateReversalJournalEntryOpts struct {
	Command                                optional.String
	PostJournalEntriesTransactionIdRequest optional.Interface
}

/*
CreateReversalJournalEntry Update Running balances for Journal Entries
This API calculates the running balances for office. If office ID not provided this API calculates running balances for all offices.  Mandatory Fields officeId
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param transactionId transactionId
 * @param optional nil or *CreateReversalJournalEntryOpts - Optional Parameters:
 * @param "Command" (optional.String) -  command
 * @param "PostJournalEntriesTransactionIdRequest" (optional.Interface of PostJournalEntriesTransactionIdRequest) -
@return PostJournalEntriesTransactionIdResponse
*/
func (a *JournalEntriesApiService) CreateReversalJournalEntry(ctx _context.Context, transactionId string, opts *CreateReversalJournalEntryOpts) (models.PostJournalEntriesTransactionIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostJournalEntriesTransactionIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/journalentries/{transactionId}"
	path = strings.Replace(path, "{"+"transactionId"+"}", _neturl.QueryEscape(parameterToString(transactionId, "")), -1)

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
	if opts != nil && opts.PostJournalEntriesTransactionIdRequest.IsSet() {
		optPostJournalEntriesTransactionIdRequest, optPostJournalEntriesTransactionIdRequestok := opts.PostJournalEntriesTransactionIdRequest.Value().(models.PostJournalEntriesTransactionIdRequest)
		if !optPostJournalEntriesTransactionIdRequestok {
			return returnValue, nil, reportError("postJournalEntriesTransactionIdRequest should be PostJournalEntriesTransactionIdRequest")
		}
		postBody = &optPostJournalEntriesTransactionIdRequest
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

// GetJournalEntriesTemplateOpts Optional parameters for the method 'GetJournalEntriesTemplate'
type GetJournalEntriesTemplateOpts struct {
	OfficeId   optional.Int64
	DateFormat optional.String
}

/*
GetJournalEntriesTemplate Method for GetJournalEntriesTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetJournalEntriesTemplateOpts - Optional Parameters:
 * @param "OfficeId" (optional.Int64) -
 * @param "DateFormat" (optional.String) -
*/
func (a *JournalEntriesApiService) GetJournalEntriesTemplate(ctx _context.Context, opts *GetJournalEntriesTemplateOpts) (*_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/journalentries/downloadtemplate"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.OfficeId.IsSet() {
		queryParams.Add("officeId", parameterToString(opts.OfficeId.Value(), ""))
	}
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

// PostJournalEntriesTemplateOpts Optional parameters for the method 'PostJournalEntriesTemplate'
type PostJournalEntriesTemplateOpts struct {
	File       optional.Interface
	Locale     optional.String
	DateFormat optional.String
}

/*
PostJournalEntriesTemplate Method for PostJournalEntriesTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *PostJournalEntriesTemplateOpts - Optional Parameters:
 * @param "File" (optional.Interface of FormDataContentDisposition) -
 * @param "Locale" (optional.String) -
 * @param "DateFormat" (optional.String) -
@return string
*/
func (a *JournalEntriesApiService) PostJournalEntriesTemplate(ctx _context.Context, opts *PostJournalEntriesTemplateOpts) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/journalentries/uploadtemplate"
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

// RetreiveJournalEntryByIdOpts Optional parameters for the method 'RetreiveJournalEntryById'
type RetreiveJournalEntryByIdOpts struct {
	RunningBalance     optional.Bool
	TransactionDetails optional.Bool
}

/*
RetreiveJournalEntryById Retrieve a single Entry
Example Requests:  journalentries/1    journalentries/1?fields&#x3D;officeName,glAccountId,entryType,amount  journalentries/1?runningBalance&#x3D;true  journalentries/1?transactionDetails&#x3D;true
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param journalEntryId journalEntryId
 * @param optional nil or *RetreiveJournalEntryByIdOpts - Optional Parameters:
 * @param "RunningBalance" (optional.Bool) -  runningBalance
 * @param "TransactionDetails" (optional.Bool) -  transactionDetails
@return JournalEntryData
*/
func (a *JournalEntriesApiService) RetreiveJournalEntryById(ctx _context.Context, journalEntryId int64, opts *RetreiveJournalEntryByIdOpts) (models.JournalEntryData, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.JournalEntryData
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/journalentries/{journalEntryId}"
	path = strings.Replace(path, "{"+"journalEntryId"+"}", _neturl.QueryEscape(parameterToString(journalEntryId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.RunningBalance.IsSet() {
		queryParams.Add("runningBalance", parameterToString(opts.RunningBalance.Value(), ""))
	}
	if opts != nil && opts.TransactionDetails.IsSet() {
		queryParams.Add("transactionDetails", parameterToString(opts.TransactionDetails.Value(), ""))
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

// RetrieveAll1Opts Optional parameters for the method 'RetrieveAll1'
type RetrieveAll1Opts struct {
	OfficeId           optional.Int64
	GlAccountId        optional.Int64
	ManualEntriesOnly  optional.Bool
	FromDate           optional.Interface
	ToDate             optional.Interface
	TransactionId      optional.String
	EntityType         optional.Int32
	Offset             optional.Int32
	Limit              optional.Int32
	OrderBy            optional.String
	SortOrder          optional.String
	Locale             optional.String
	DateFormat         optional.String
	LoanId             optional.Int64
	SavingsId          optional.Int64
	RunningBalance     optional.Bool
	TransactionDetails optional.Bool
}

/*
RetrieveAll1 List Journal Entries
The list capability of journal entries can support pagination and sorting.  Example Requests:  journalentries  journalentries?transactionId&#x3D;PB37X8Y21EQUY4S  journalentries?officeId&#x3D;1&amp;manualEntriesOnly&#x3D;true&amp;fromDate&#x3D;1 July 2013&amp;toDate&#x3D;15 July 2013&amp;dateFormat&#x3D;dd MMMM yyyy&amp;locale&#x3D;en  journalentries?fields&#x3D;officeName,glAccountName,transactionDate  journalentries?offset&#x3D;10&amp;limit&#x3D;50  journalentries?orderBy&#x3D;transactionId&amp;sortOrder&#x3D;DESC  journalentries?runningBalance&#x3D;true  journalentries?transactionDetails&#x3D;true  journalentries?loanId&#x3D;12  journalentries?savingsId&#x3D;24
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveAll1Opts - Optional Parameters:
 * @param "OfficeId" (optional.Int64) -  officeId
 * @param "GlAccountId" (optional.Int64) -  glAccountId
 * @param "ManualEntriesOnly" (optional.Bool) -  manualEntriesOnly
 * @param "FromDate" (optional.Interface of map[string]interface{}) -  fromDate
 * @param "ToDate" (optional.Interface of map[string]interface{}) -  toDate
 * @param "TransactionId" (optional.String) -  transactionId
 * @param "EntityType" (optional.Int32) -  entityType
 * @param "Offset" (optional.Int32) -  offset
 * @param "Limit" (optional.Int32) -  limit
 * @param "OrderBy" (optional.String) -  orderBy
 * @param "SortOrder" (optional.String) -  sortOrder
 * @param "Locale" (optional.String) -  locale
 * @param "DateFormat" (optional.String) -  dateFormat
 * @param "LoanId" (optional.Int64) -  loanId
 * @param "SavingsId" (optional.Int64) -  savingsId
 * @param "RunningBalance" (optional.Bool) -  runningBalance
 * @param "TransactionDetails" (optional.Bool) -  transactionDetails
@return []JournalEntryData
*/
func (a *JournalEntriesApiService) RetrieveAll1(ctx _context.Context, opts *RetrieveAll1Opts) ([]models.JournalEntryData, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  []models.JournalEntryData
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/journalentries"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.OfficeId.IsSet() {
		queryParams.Add("officeId", parameterToString(opts.OfficeId.Value(), ""))
	}
	if opts != nil && opts.GlAccountId.IsSet() {
		queryParams.Add("glAccountId", parameterToString(opts.GlAccountId.Value(), ""))
	}
	if opts != nil && opts.ManualEntriesOnly.IsSet() {
		queryParams.Add("manualEntriesOnly", parameterToString(opts.ManualEntriesOnly.Value(), ""))
	}
	if opts != nil && opts.FromDate.IsSet() {
		queryParams.Add("fromDate", parameterToString(opts.FromDate.Value(), ""))
	}
	if opts != nil && opts.ToDate.IsSet() {
		queryParams.Add("toDate", parameterToString(opts.ToDate.Value(), ""))
	}
	if opts != nil && opts.TransactionId.IsSet() {
		queryParams.Add("transactionId", parameterToString(opts.TransactionId.Value(), ""))
	}
	if opts != nil && opts.EntityType.IsSet() {
		queryParams.Add("entityType", parameterToString(opts.EntityType.Value(), ""))
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
	if opts != nil && opts.Locale.IsSet() {
		queryParams.Add("locale", parameterToString(opts.Locale.Value(), ""))
	}
	if opts != nil && opts.DateFormat.IsSet() {
		queryParams.Add("dateFormat", parameterToString(opts.DateFormat.Value(), ""))
	}
	if opts != nil && opts.LoanId.IsSet() {
		queryParams.Add("loanId", parameterToString(opts.LoanId.Value(), ""))
	}
	if opts != nil && opts.SavingsId.IsSet() {
		queryParams.Add("savingsId", parameterToString(opts.SavingsId.Value(), ""))
	}
	if opts != nil && opts.RunningBalance.IsSet() {
		queryParams.Add("runningBalance", parameterToString(opts.RunningBalance.Value(), ""))
	}
	if opts != nil && opts.TransactionDetails.IsSet() {
		queryParams.Add("transactionDetails", parameterToString(opts.TransactionDetails.Value(), ""))
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

// RetrieveJournalEntriesOpts Optional parameters for the method 'RetrieveJournalEntries'
type RetrieveJournalEntriesOpts struct {
	Offset  optional.Int32
	Limit   optional.Int32
	EntryId optional.Int64
}

/*
RetrieveJournalEntries Method for RetrieveJournalEntries
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveJournalEntriesOpts - Optional Parameters:
 * @param "Offset" (optional.Int32) -
 * @param "Limit" (optional.Int32) -
 * @param "EntryId" (optional.Int64) -
@return string
*/
func (a *JournalEntriesApiService) RetrieveJournalEntries(ctx _context.Context, opts *RetrieveJournalEntriesOpts) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/journalentries/provisioning"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.Offset.IsSet() {
		queryParams.Add("offset", parameterToString(opts.Offset.Value(), ""))
	}
	if opts != nil && opts.Limit.IsSet() {
		queryParams.Add("limit", parameterToString(opts.Limit.Value(), ""))
	}
	if opts != nil && opts.EntryId.IsSet() {
		queryParams.Add("entryId", parameterToString(opts.EntryId.Value(), ""))
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

// RetrieveOpeningBalanceOpts Optional parameters for the method 'RetrieveOpeningBalance'
type RetrieveOpeningBalanceOpts struct {
	OfficeId     optional.Int64
	CurrencyCode optional.String
}

/*
RetrieveOpeningBalance Method for RetrieveOpeningBalance
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveOpeningBalanceOpts - Optional Parameters:
 * @param "OfficeId" (optional.Int64) -
 * @param "CurrencyCode" (optional.String) -
@return string
*/
func (a *JournalEntriesApiService) RetrieveOpeningBalance(ctx _context.Context, opts *RetrieveOpeningBalanceOpts) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/journalentries/openingbalance"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.OfficeId.IsSet() {
		queryParams.Add("officeId", parameterToString(opts.OfficeId.Value(), ""))
	}
	if opts != nil && opts.CurrencyCode.IsSet() {
		queryParams.Add("currencyCode", parameterToString(opts.CurrencyCode.Value(), ""))
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
