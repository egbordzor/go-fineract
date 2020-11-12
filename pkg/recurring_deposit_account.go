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

// RecurringDepositAccountApiService RecurringDepositAccountApi service
type RecurringDepositAccountApiService service

// AccountClosureTemplate1Opts Optional parameters for the method 'AccountClosureTemplate1'
type AccountClosureTemplate1Opts struct {
	Command optional.String
}

/*
AccountClosureTemplate1 Method for AccountClosureTemplate1
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId accountId
 * @param optional nil or *AccountClosureTemplate1Opts - Optional Parameters:
 * @param "Command" (optional.String) -  command
@return string
*/
func (a *RecurringDepositAccountApiService) AccountClosureTemplate1(ctx _context.Context, accountId int64, opts *AccountClosureTemplate1Opts) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts/{accountId}/template"
	path = strings.Replace(path, "{"+"accountId"+"}", _neturl.QueryEscape(parameterToString(accountId, "")), -1)

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
Delete16 Delete a recurring deposit application
At present we support hard delete of recurring deposit application so long as its in &#39;Submitted and pending approval&#39; state. One the application is moves past this state, it is not possible to do a &#39;hard&#39; delete of the application or the account. An API endpoint will be added to close/de-activate the recurring deposit account.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId accountId
@return DeleteRecurringDepositAccountsResponse
*/
func (a *RecurringDepositAccountApiService) Delete16(ctx _context.Context, accountId int64) (models.DeleteRecurringDepositAccountsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodDelete
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.DeleteRecurringDepositAccountsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts/{accountId}"
	path = strings.Replace(path, "{"+"accountId"+"}", _neturl.QueryEscape(parameterToString(accountId, "")), -1)

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

// GetRecurringDepositTemplateOpts Optional parameters for the method 'GetRecurringDepositTemplate'
type GetRecurringDepositTemplateOpts struct {
	OfficeId   optional.Int64
	StaffId    optional.Int64
	DateFormat optional.String
}

/*
GetRecurringDepositTemplate Method for GetRecurringDepositTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetRecurringDepositTemplateOpts - Optional Parameters:
 * @param "OfficeId" (optional.Int64) -
 * @param "StaffId" (optional.Int64) -
 * @param "DateFormat" (optional.String) -
*/
func (a *RecurringDepositAccountApiService) GetRecurringDepositTemplate(ctx _context.Context, opts *GetRecurringDepositTemplateOpts) (*_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts/downloadtemplate"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.OfficeId.IsSet() {
		queryParams.Add("officeId", parameterToString(opts.OfficeId.Value(), ""))
	}
	if opts != nil && opts.StaffId.IsSet() {
		queryParams.Add("staffId", parameterToString(opts.StaffId.Value(), ""))
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

// GetRecurringDepositTransactionTemplateOpts Optional parameters for the method 'GetRecurringDepositTransactionTemplate'
type GetRecurringDepositTransactionTemplateOpts struct {
	OfficeId   optional.Int64
	DateFormat optional.String
}

/*
GetRecurringDepositTransactionTemplate Method for GetRecurringDepositTransactionTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetRecurringDepositTransactionTemplateOpts - Optional Parameters:
 * @param "OfficeId" (optional.Int64) -
 * @param "DateFormat" (optional.String) -
*/
func (a *RecurringDepositAccountApiService) GetRecurringDepositTransactionTemplate(ctx _context.Context, opts *GetRecurringDepositTransactionTemplateOpts) (*_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts/transactions/downloadtemplate"
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

// HandleCommands5Opts Optional parameters for the method 'HandleCommands5'
type HandleCommands5Opts struct {
	Command optional.String
}

/*
HandleCommands5 Approve recurring deposit application | Undo approval recurring deposit application | Reject recurring deposit application | Withdraw recurring deposit application | Activate a recurring deposit account | Update the recommended deposit amount for a recurring deposit account | Close a recurring deposit account | Premature Close a recurring deposit account | Calculate Premature amount on Recurring deposit account | Calculate Interest on recurring Deposit Account | Post Interest on recurring Deposit Account
Approve recurring deposit application:  Approves recurring deposit application so long as its in &#39;Submitted and pending approval&#39; state.  Undo approval recurring deposit application:  Will move &#39;approved&#39; recurring deposit application back to &#39;Submitted and pending approval&#39; state.  Reject recurring deposit application  Rejects recurring deposit application so long as its in &#39;Submitted and pending approval&#39; state.  Withdraw recurring deposit application:  Used when an applicant withdraws from the recurring deposit application. It must be in &#39;Submitted and pending approval&#39; state.  Activate a recurring deposit account:  Results in an approved recurring deposit application being converted into an &#39;active&#39; recurring deposit account.  Update the recommended deposit amount for a recurring deposit account:  Updates the recommended deposit amount for a RD account as on the effective date.  Close a recurring deposit account  Results in a Matured recurring deposit account being converted into a &#39;closed&#39; recurring deposit account.  On account close allowed actions are.Premature Close a recurring deposit account:  Results in an Active recurring deposit account being converted into a &#39;Premature Closed&#39; recurring deposit account with options to withdraw prematured amount. (premature amount is calculated using interest rate chart applicable along with penal interest if any.)  On account premature closure allowed actions are.  Calculate Premature amount on Recurring deposit account:  Calculate premature amount on recurring deposit till premature close date. Premature amount is calculated based on interest chart and penal interest applicable if any.  Calculate Interest on recurring Deposit Account:  Calculates interest earned on a recurring deposit account based on todays date. It does not attempt to post or credit the interest on the account. That is responsibility of the Post Interest API that will likely be called by overnight process.  Post Interest on recurring Deposit Account:  Calculates and Posts interest earned on a recurring deposit account based on todays date and whether an interest posting or crediting event is due.  Showing request/response for &#39;Post Interest on recurring Deposit Account&#39;
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId accountId
 * @param body
 * @param optional nil or *HandleCommands5Opts - Optional Parameters:
 * @param "Command" (optional.String) -  command
@return PostRecurringDepositAccountsAccountIdResponse
*/
func (a *RecurringDepositAccountApiService) HandleCommands5(ctx _context.Context, accountId int64, body map[string]interface{}, opts *HandleCommands5Opts) (models.PostRecurringDepositAccountsAccountIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostRecurringDepositAccountsAccountIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts/{accountId}"
	path = strings.Replace(path, "{"+"accountId"+"}", _neturl.QueryEscape(parameterToString(accountId, "")), -1)

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
	postBody = &body
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

// PostRecurringDepositTemplateOpts Optional parameters for the method 'PostRecurringDepositTemplate'
type PostRecurringDepositTemplateOpts struct {
	File       optional.Interface
	Locale     optional.String
	DateFormat optional.String
}

/*
PostRecurringDepositTemplate Method for PostRecurringDepositTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *PostRecurringDepositTemplateOpts - Optional Parameters:
 * @param "File" (optional.Interface of FormDataContentDisposition) -
 * @param "Locale" (optional.String) -
 * @param "DateFormat" (optional.String) -
@return string
*/
func (a *RecurringDepositAccountApiService) PostRecurringDepositTemplate(ctx _context.Context, opts *PostRecurringDepositTemplateOpts) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts/uploadtemplate"
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

// PostRecurringDepositTransactionsTemplateOpts Optional parameters for the method 'PostRecurringDepositTransactionsTemplate'
type PostRecurringDepositTransactionsTemplateOpts struct {
	File       optional.Interface
	Locale     optional.String
	DateFormat optional.String
}

/*
PostRecurringDepositTransactionsTemplate Method for PostRecurringDepositTransactionsTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *PostRecurringDepositTransactionsTemplateOpts - Optional Parameters:
 * @param "File" (optional.Interface of FormDataContentDisposition) -
 * @param "Locale" (optional.String) -
 * @param "DateFormat" (optional.String) -
@return string
*/
func (a *RecurringDepositAccountApiService) PostRecurringDepositTransactionsTemplate(ctx _context.Context, opts *PostRecurringDepositTransactionsTemplateOpts) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts/transactions/uploadtemplate"
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

// RetrieveAll30Opts Optional parameters for the method 'RetrieveAll30'
type RetrieveAll30Opts struct {
	Paged     optional.Bool
	Offset    optional.Int32
	Limit     optional.Int32
	OrderBy   optional.String
	SortOrder optional.String
}

/*
RetrieveAll30 List Recurring deposit applications/accounts
Lists Recurring deposit applications/accounts  Example Requests:  recurringdepositaccounts   recurringdepositaccounts?fields&#x3D;name
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveAll30Opts - Optional Parameters:
 * @param "Paged" (optional.Bool) -  paged
 * @param "Offset" (optional.Int32) -  offset
 * @param "Limit" (optional.Int32) -  limit
 * @param "OrderBy" (optional.String) -  orderBy
 * @param "SortOrder" (optional.String) -  sortOrder
@return []GetRecurringDepositAccountsResponse
*/
func (a *RecurringDepositAccountApiService) RetrieveAll30(ctx _context.Context, opts *RetrieveAll30Opts) ([]models.GetRecurringDepositAccountsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  []models.GetRecurringDepositAccountsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

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

// RetrieveOne20Opts Optional parameters for the method 'RetrieveOne20'
type RetrieveOne20Opts struct {
	StaffInSelectedOfficeOnly optional.Bool
	ChargeStatus              optional.String
}

/*
RetrieveOne20 Retrieve a recurring deposit application/account
Retrieves a recurring deposit application/account  Example Requests :  recurringdepositaccounts/1   recurringdepositaccounts/1?associations&#x3D;all
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId accountId
 * @param optional nil or *RetrieveOne20Opts - Optional Parameters:
 * @param "StaffInSelectedOfficeOnly" (optional.Bool) -  staffInSelectedOfficeOnly
 * @param "ChargeStatus" (optional.String) -  chargeStatus
@return GetRecurringDepositAccountsAccountIdResponse
*/
func (a *RecurringDepositAccountApiService) RetrieveOne20(ctx _context.Context, accountId int64, opts *RetrieveOne20Opts) (models.GetRecurringDepositAccountsAccountIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetRecurringDepositAccountsAccountIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts/{accountId}"
	path = strings.Replace(path, "{"+"accountId"+"}", _neturl.QueryEscape(parameterToString(accountId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.StaffInSelectedOfficeOnly.IsSet() {
		queryParams.Add("staffInSelectedOfficeOnly", parameterToString(opts.StaffInSelectedOfficeOnly.Value(), ""))
	}
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
SubmitApplication1 Submit new recurring deposit application
Submits new recurring deposit application  Mandatory Fields: clientId or groupId, productId, submittedOnDate, depositAmount, depositPeriod, depositPeriodFrequencyId  Optional Fields: accountNo, externalId, fieldOfficerId,linkAccountId(if provided initial deposit amount will be collected from this account),transferInterestToSavings(By enabling this flag all interest postings will be transferred to linked saving account )  Inherited from Product (if not provided): interestCompoundingPeriodType, interestCalculationType, interestCalculationDaysInYearType, lockinPeriodFrequency, lockinPeriodFrequencyType, preClosurePenalApplicable, preClosurePenalInterest, preClosurePenalInterestOnTypeId, charts, withHoldTax
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param postRecurringDepositAccountsRequest
@return PostRecurringDepositAccountsResponse
*/
func (a *RecurringDepositAccountApiService) SubmitApplication1(ctx _context.Context, postRecurringDepositAccountsRequest models.PostRecurringDepositAccountsRequest) (models.PostRecurringDepositAccountsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostRecurringDepositAccountsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts"
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
	postBody = &postRecurringDepositAccountsRequest
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

// Template13Opts Optional parameters for the method 'Template13'
type Template13Opts struct {
	ClientId                  optional.Int64
	GroupId                   optional.Int64
	ProductId                 optional.Int64
	StaffInSelectedOfficeOnly optional.Bool
}

/*
Template13 Retrieve recurring Deposit Account Template
This is a convenience resource. It can be useful when building maintenance user interface screens for recurring deposit applications. The template data returned consists of any or all of:  Field Defaults Allowed Value Lists  Example Requests:  recurringdepositaccounts/template?clientId&#x3D;1   recurringdepositaccounts/template?clientId&#x3D;1&amp;productId&#x3D;1
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *Template13Opts - Optional Parameters:
 * @param "ClientId" (optional.Int64) -  clientId
 * @param "GroupId" (optional.Int64) -  groupId
 * @param "ProductId" (optional.Int64) -  productId
 * @param "StaffInSelectedOfficeOnly" (optional.Bool) -  staffInSelectedOfficeOnly
@return GetRecurringDepositAccountsTemplateResponse
*/
func (a *RecurringDepositAccountApiService) Template13(ctx _context.Context, opts *Template13Opts) (models.GetRecurringDepositAccountsTemplateResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetRecurringDepositAccountsTemplateResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts/template"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.ClientId.IsSet() {
		queryParams.Add("clientId", parameterToString(opts.ClientId.Value(), ""))
	}
	if opts != nil && opts.GroupId.IsSet() {
		queryParams.Add("groupId", parameterToString(opts.GroupId.Value(), ""))
	}
	if opts != nil && opts.ProductId.IsSet() {
		queryParams.Add("productId", parameterToString(opts.ProductId.Value(), ""))
	}
	if opts != nil && opts.StaffInSelectedOfficeOnly.IsSet() {
		queryParams.Add("staffInSelectedOfficeOnly", parameterToString(opts.StaffInSelectedOfficeOnly.Value(), ""))
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
Update16 Modify a recurring deposit application
Recurring deposit application can only be modified when in &#39;Submitted and pending approval&#39; state. Once the application is approved, the details cannot be changed using this method. Specific api endpoints will be created to allow change of interest detail such as rate, compounding period, posting period etc
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId accountId
 * @param putRecurringDepositAccountsAccountIdRequest
@return PutRecurringDepositAccountsAccountIdResponse
*/
func (a *RecurringDepositAccountApiService) Update16(ctx _context.Context, accountId int64, putRecurringDepositAccountsAccountIdRequest models.PutRecurringDepositAccountsAccountIdRequest) (models.PutRecurringDepositAccountsAccountIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPut
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PutRecurringDepositAccountsAccountIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts/{accountId}"
	path = strings.Replace(path, "{"+"accountId"+"}", _neturl.QueryEscape(parameterToString(accountId, "")), -1)

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
	postBody = &putRecurringDepositAccountsAccountIdRequest
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
