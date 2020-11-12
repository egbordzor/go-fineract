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

// LoanTransactionsApiService LoanTransactionsApi service
type LoanTransactionsApiService service

/*
AdjustLoanTransaction Adjust a Transaction
Note: there is no need to specify command&#x3D;{transactionType} parameter.  Mandatory Fields: transactionDate, transactionAmount
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param loanId loanId
 * @param transactionId transactionId
 * @param postLoansLoanIdTransactionsTransactionIdRequest
@return PostLoansLoanIdTransactionsTransactionIdResponse
*/
func (a *LoanTransactionsApiService) AdjustLoanTransaction(ctx _context.Context, loanId int64, transactionId int64, postLoansLoanIdTransactionsTransactionIdRequest models.PostLoansLoanIdTransactionsTransactionIdRequest) (models.PostLoansLoanIdTransactionsTransactionIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostLoansLoanIdTransactionsTransactionIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/{loanId}/transactions/{transactionId}"
	path = strings.Replace(path, "{"+"loanId"+"}", _neturl.QueryEscape(parameterToString(loanId, "")), -1)

	path = strings.Replace(path, "{"+"transactionId"+"}", _neturl.QueryEscape(parameterToString(transactionId, "")), -1)

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
	postBody = &postLoansLoanIdTransactionsTransactionIdRequest
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

// ExecuteLoanTransactionOpts Optional parameters for the method 'ExecuteLoanTransaction'
type ExecuteLoanTransactionOpts struct {
	Command optional.String
}

/*
ExecuteLoanTransaction Make a Repayment | Make a Refund of an Active Loan by Cash | Foreclosure of an Active Loan | Waive Interest | Write-off Loan | Make Recovery Payment | Undo Loan Write-off Transaction
Make Recovery Payment:  This API allows collecting recovery payments for written-off loans
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param loanId loanId
 * @param body
 * @param optional nil or *ExecuteLoanTransactionOpts - Optional Parameters:
 * @param "Command" (optional.String) -  command
@return PostLoansLoanIdTransactionsResponse
*/
func (a *LoanTransactionsApiService) ExecuteLoanTransaction(ctx _context.Context, loanId int64, body map[string]interface{}, opts *ExecuteLoanTransactionOpts) (models.PostLoansLoanIdTransactionsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostLoansLoanIdTransactionsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/{loanId}/transactions"
	path = strings.Replace(path, "{"+"loanId"+"}", _neturl.QueryEscape(parameterToString(loanId, "")), -1)

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

/*
RetrieveTransaction Retrieve a Transaction Details
Retrieves a Transaction Details  Example Request:  loans/5/transactions/3
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param loanId loanId
 * @param transactionId transactionId
@return GetLoansLoanIdTransactionsTransactionIdResponse
*/
func (a *LoanTransactionsApiService) RetrieveTransaction(ctx _context.Context, loanId int64, transactionId int64) (models.GetLoansLoanIdTransactionsTransactionIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetLoansLoanIdTransactionsTransactionIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/{loanId}/transactions/{transactionId}"
	path = strings.Replace(path, "{"+"loanId"+"}", _neturl.QueryEscape(parameterToString(loanId, "")), -1)

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

// RetrieveTransactionTemplateOpts Optional parameters for the method 'RetrieveTransactionTemplate'
type RetrieveTransactionTemplateOpts struct {
	Command         optional.String
	DateFormat      optional.String
	TransactionDate optional.Interface
	Locale          optional.String
}

/*
RetrieveTransactionTemplate Retrieve Loan Transaction Template
This is a convenience resource. It can be useful when building maintenance user interface screens for client applications. The template data returned consists of any or all of:  Field Defaults Allowed Value Lists  Example Request:  loans/1/transactions/template?command&#x3D;repayment
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param loanId loanId
 * @param optional nil or *RetrieveTransactionTemplateOpts - Optional Parameters:
 * @param "Command" (optional.String) -  command
 * @param "DateFormat" (optional.String) -  dateFormat
 * @param "TransactionDate" (optional.Interface of map[string]interface{}) -  transactionDate
 * @param "Locale" (optional.String) -  locale
@return GetLoansLoanIdTransactionsTemplateResponse
*/
func (a *LoanTransactionsApiService) RetrieveTransactionTemplate(ctx _context.Context, loanId int64, opts *RetrieveTransactionTemplateOpts) (models.GetLoansLoanIdTransactionsTemplateResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetLoansLoanIdTransactionsTemplateResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/{loanId}/transactions/template"
	path = strings.Replace(path, "{"+"loanId"+"}", _neturl.QueryEscape(parameterToString(loanId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.Command.IsSet() {
		queryParams.Add("command", parameterToString(opts.Command.Value(), ""))
	}
	if opts != nil && opts.DateFormat.IsSet() {
		queryParams.Add("dateFormat", parameterToString(opts.DateFormat.Value(), ""))
	}
	if opts != nil && opts.TransactionDate.IsSet() {
		queryParams.Add("transactionDate", parameterToString(opts.TransactionDate.Value(), ""))
	}
	if opts != nil && opts.Locale.IsSet() {
		queryParams.Add("locale", parameterToString(opts.Locale.Value(), ""))
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
