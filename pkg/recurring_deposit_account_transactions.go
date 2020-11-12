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

// RecurringDepositAccountTransactionsApiService RecurringDepositAccountTransactionsApi service
type RecurringDepositAccountTransactionsApiService service

// HandleTransactionCommandsOpts Optional parameters for the method 'HandleTransactionCommands'
type HandleTransactionCommandsOpts struct {
	Command optional.String
}

/*
HandleTransactionCommands Adjust Transaction | Undo transaction
Adjust Transaction:  This command modifies the given transaction.  Undo transaction:  This command reverses the given transaction.  Showing request/response for &#39;Adjust Transaction&#39;
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param recurringDepositAccountId recurringDepositAccountId
 * @param transactionId transactionId
 * @param postRecurringDepositAccountsRecurringDepositAccountIdTransactionsRequest
 * @param optional nil or *HandleTransactionCommandsOpts - Optional Parameters:
 * @param "Command" (optional.String) -  command
@return PostRecurringDepositAccountsRecurringDepositAccountIdTransactionsTransactionIdResponse
*/
func (a *RecurringDepositAccountTransactionsApiService) HandleTransactionCommands(ctx _context.Context, recurringDepositAccountId int64, transactionId int64, postRecurringDepositAccountsRecurringDepositAccountIdTransactionsRequest models.PostRecurringDepositAccountsRecurringDepositAccountIdTransactionsRequest, opts *HandleTransactionCommandsOpts) (models.PostRecurringDepositAccountsRecurringDepositAccountIdTransactionsTransactionIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostRecurringDepositAccountsRecurringDepositAccountIdTransactionsTransactionIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts/{recurringDepositAccountId}/transactions/{transactionId}"
	path = strings.Replace(path, "{"+"recurringDepositAccountId"+"}", _neturl.QueryEscape(parameterToString(recurringDepositAccountId, "")), -1)

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
	postBody = &postRecurringDepositAccountsRecurringDepositAccountIdTransactionsRequest
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
RetrieveOne19 Retrieve Recurring Deposit Account Transaction
Retrieves Recurring Deposit Account Transaction  Example Requests:  recurringdepositaccounts/1/transactions/1
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param recurringDepositAccountId recurringDepositAccountId
 * @param transactionId transactionId
@return GetRecurringDepositAccountsRecurringDepositAccountIdTransactionsTransactionIdResponse
*/
func (a *RecurringDepositAccountTransactionsApiService) RetrieveOne19(ctx _context.Context, recurringDepositAccountId int64, transactionId int64) (models.GetRecurringDepositAccountsRecurringDepositAccountIdTransactionsTransactionIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetRecurringDepositAccountsRecurringDepositAccountIdTransactionsTransactionIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts/{recurringDepositAccountId}/transactions/{transactionId}"
	path = strings.Replace(path, "{"+"recurringDepositAccountId"+"}", _neturl.QueryEscape(parameterToString(recurringDepositAccountId, "")), -1)

	path = strings.Replace(path, "{"+"transactionId"+"}", _neturl.QueryEscape(parameterToString(transactionId, "")), -1)

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

// RetrieveTemplate15Opts Optional parameters for the method 'RetrieveTemplate15'
type RetrieveTemplate15Opts struct {
	Command optional.String
}

/*
RetrieveTemplate15 Retrieve Recurring Deposit Account Transaction Template
This is a convenience resource. It can be useful when building maintenance user interface screens for client applications. The template data returned consists of any or all of:  Field Defaults Allowed Value Lists Example Requests:  recurringdepositaccounts/1/transactions/template?command&#x3D;deposit  recurringdepositaccounts/1/transactions/template?command&#x3D;withdrawal
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param recurringDepositAccountId recurringDepositAccountId
 * @param optional nil or *RetrieveTemplate15Opts - Optional Parameters:
 * @param "Command" (optional.String) -  command
@return GetRecurringDepositAccountsRecurringDepositAccountIdTransactionsTemplateResponse
*/
func (a *RecurringDepositAccountTransactionsApiService) RetrieveTemplate15(ctx _context.Context, recurringDepositAccountId int64, opts *RetrieveTemplate15Opts) (models.GetRecurringDepositAccountsRecurringDepositAccountIdTransactionsTemplateResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetRecurringDepositAccountsRecurringDepositAccountIdTransactionsTemplateResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts/{recurringDepositAccountId}/transactions/template"
	path = strings.Replace(path, "{"+"recurringDepositAccountId"+"}", _neturl.QueryEscape(parameterToString(recurringDepositAccountId, "")), -1)

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

// Transaction1Opts Optional parameters for the method 'Transaction1'
type Transaction1Opts struct {
	Command optional.String
}

/*
Transaction1 Deposit Transaction | Withdrawal Transaction
Deposit Transaction:  Used for a deposit transaction  Withdrawal Transaction:  Used for a Withdrawal Transaction  Showing request/response for Deposit Transaction
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param recurringDepositAccountId recurringDepositAccountId
 * @param postRecurringDepositAccountsRecurringDepositAccountIdTransactionsRequest
 * @param optional nil or *Transaction1Opts - Optional Parameters:
 * @param "Command" (optional.String) -  command
@return PostRecurringDepositAccountsRecurringDepositAccountIdTransactionsResponse
*/
func (a *RecurringDepositAccountTransactionsApiService) Transaction1(ctx _context.Context, recurringDepositAccountId int64, postRecurringDepositAccountsRecurringDepositAccountIdTransactionsRequest models.PostRecurringDepositAccountsRecurringDepositAccountIdTransactionsRequest, opts *Transaction1Opts) (models.PostRecurringDepositAccountsRecurringDepositAccountIdTransactionsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostRecurringDepositAccountsRecurringDepositAccountIdTransactionsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/recurringdepositaccounts/{recurringDepositAccountId}/transactions"
	path = strings.Replace(path, "{"+"recurringDepositAccountId"+"}", _neturl.QueryEscape(parameterToString(recurringDepositAccountId, "")), -1)

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
	postBody = &postRecurringDepositAccountsRecurringDepositAccountIdTransactionsRequest
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
