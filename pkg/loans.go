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

// LoansApiService LoansApi service
type LoansApiService service

// CalculateLoanScheduleOrSubmitLoanApplicationOpts Optional parameters for the method 'CalculateLoanScheduleOrSubmitLoanApplication'
type CalculateLoanScheduleOrSubmitLoanApplicationOpts struct {
	Command optional.String
}

/*
CalculateLoanScheduleOrSubmitLoanApplication Calculate loan repayment schedule | Submit a new Loan Application
It calculates the loan repayment Schedule Submits a new loan application Mandatory Fields: clientId, productId, principal, loanTermFrequency, loanTermFrequencyType, loanType, numberOfRepayments, repaymentEvery, repaymentFrequencyType, interestRatePerPeriod, amortizationType, interestType, interestCalculationPeriodType, transactionProcessingStrategyId, expectedDisbursementDate, submittedOnDate, loanType Optional Fields: graceOnPrincipalPayment, graceOnInterestPayment, graceOnInterestCharged, linkAccountId, allowPartialPeriodInterestCalcualtion, fixedEmiAmount, maxOutstandingLoanBalance, disbursementData, graceOnArrearsAgeing, createStandingInstructionAtDisbursement (requires linkedAccountId if set to true) Additional Mandatory Fields if interest recalculation is enabled for product and Rest frequency not same as repayment period: recalculationRestFrequencyDate Additional Mandatory Fields if interest recalculation with interest/fee compounding is enabled for product and compounding frequency not same as repayment period: recalculationCompoundingFrequencyDate Additional Mandatory Field if Entity-Datatable Check is enabled for the entity of type loan: datatables
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param postLoansRequest
 * @param optional nil or *CalculateLoanScheduleOrSubmitLoanApplicationOpts - Optional Parameters:
 * @param "Command" (optional.String) -  command
@return PostLoansResponse
*/
func (a *LoansApiService) CalculateLoanScheduleOrSubmitLoanApplication(ctx _context.Context, postLoansRequest models.PostLoansRequest, opts *CalculateLoanScheduleOrSubmitLoanApplicationOpts) (models.PostLoansResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostLoansResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans"
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
	postBody = &postLoansRequest
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
DeleteLoanApplication Delete a Loan Application
Note: Only loans in \&quot;Submitted and awaiting approval\&quot; status can be deleted.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param loanId loanId
@return DeleteLoansLoanIdResponse
*/
func (a *LoansApiService) DeleteLoanApplication(ctx _context.Context, loanId int64) (models.DeleteLoansLoanIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodDelete
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.DeleteLoansLoanIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/{loanId}"
	path = strings.Replace(path, "{"+"loanId"+"}", _neturl.QueryEscape(parameterToString(loanId, "")), -1)

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
GetGlimRepaymentTemplate Method for GetGlimRepaymentTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param glimId
@return string
*/
func (a *LoansApiService) GetGlimRepaymentTemplate(ctx _context.Context, glimId int64) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/glimAccount/{glimId}"
	path = strings.Replace(path, "{"+"glimId"+"}", _neturl.QueryEscape(parameterToString(glimId, "")), -1)

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

// GetLoanRepaymentTemplateOpts Optional parameters for the method 'GetLoanRepaymentTemplate'
type GetLoanRepaymentTemplateOpts struct {
	OfficeId   optional.Int64
	DateFormat optional.String
}

/*
GetLoanRepaymentTemplate Method for GetLoanRepaymentTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetLoanRepaymentTemplateOpts - Optional Parameters:
 * @param "OfficeId" (optional.Int64) -
 * @param "DateFormat" (optional.String) -
*/
func (a *LoansApiService) GetLoanRepaymentTemplate(ctx _context.Context, opts *GetLoanRepaymentTemplateOpts) (*_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/repayments/downloadtemplate"
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

// GetLoansTemplateOpts Optional parameters for the method 'GetLoansTemplate'
type GetLoansTemplateOpts struct {
	OfficeId   optional.Int64
	StaffId    optional.Int64
	DateFormat optional.String
}

/*
GetLoansTemplate Method for GetLoansTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetLoansTemplateOpts - Optional Parameters:
 * @param "OfficeId" (optional.Int64) -
 * @param "StaffId" (optional.Int64) -
 * @param "DateFormat" (optional.String) -
*/
func (a *LoansApiService) GetLoansTemplate(ctx _context.Context, opts *GetLoansTemplateOpts) (*_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/downloadtemplate"
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

// GlimStateTransitionsOpts Optional parameters for the method 'GlimStateTransitions'
type GlimStateTransitionsOpts struct {
	Command optional.String
}

/*
GlimStateTransitions Approve GLIM Application | Undo GLIM Application Approval | Reject GLIM Application | Disburse Loan Disburse Loan To Savings Account | Undo Loan Disbursal
Approve GLIM Application: Mandatory Fields: approvedOnDate Optional Fields: approvedLoanAmount and expectedDisbursementDate Approves the GLIM application  Undo GLIM Application Approval: Undoes the GLIM Application Approval  Reject GLIM Application: Mandatory Fields: rejectedOnDate Allows you to reject the GLIM application  Disburse Loan: Mandatory Fields: actualDisbursementDate Optional Fields: transactionAmount and fixedEmiAmount Disburses the Loan  Disburse Loan To Savings Account: Mandatory Fields: actualDisbursementDate Optional Fields: transactionAmount and fixedEmiAmount Disburses the loan to Saving Account  Undo Loan Disbursal: Undoes the Loan Disbursal
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param glimId
 * @param postLoansLoanIdRequest
 * @param optional nil or *GlimStateTransitionsOpts - Optional Parameters:
 * @param "Command" (optional.String) -
@return PostLoansLoanIdResponse
*/
func (a *LoansApiService) GlimStateTransitions(ctx _context.Context, glimId int64, postLoansLoanIdRequest models.PostLoansLoanIdRequest, opts *GlimStateTransitionsOpts) (models.PostLoansLoanIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostLoansLoanIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/glimAccount/{glimId}"
	path = strings.Replace(path, "{"+"glimId"+"}", _neturl.QueryEscape(parameterToString(glimId, "")), -1)

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
	postBody = &postLoansLoanIdRequest
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
ModifyLoanApplication Modify a loan application
Loan application can only be modified when in &#39;Submitted and pending approval&#39; state. Once the application is approved, the details cannot be changed using this method.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param loanId loanId
 * @param putLoansLoanIdRequest
@return PutLoansLoanIdResponse
*/
func (a *LoansApiService) ModifyLoanApplication(ctx _context.Context, loanId int64, putLoansLoanIdRequest models.PutLoansLoanIdRequest) (models.PutLoansLoanIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPut
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PutLoansLoanIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/{loanId}"
	path = strings.Replace(path, "{"+"loanId"+"}", _neturl.QueryEscape(parameterToString(loanId, "")), -1)

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
	postBody = &putLoansLoanIdRequest
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

// PostLoanRepaymentTemplateOpts Optional parameters for the method 'PostLoanRepaymentTemplate'
type PostLoanRepaymentTemplateOpts struct {
	File       optional.Interface
	Locale     optional.String
	DateFormat optional.String
}

/*
PostLoanRepaymentTemplate Method for PostLoanRepaymentTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *PostLoanRepaymentTemplateOpts - Optional Parameters:
 * @param "File" (optional.Interface of FormDataContentDisposition) -
 * @param "Locale" (optional.String) -
 * @param "DateFormat" (optional.String) -
@return string
*/
func (a *LoansApiService) PostLoanRepaymentTemplate(ctx _context.Context, opts *PostLoanRepaymentTemplateOpts) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/repayments/uploadtemplate"
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

// PostLoanTemplateOpts Optional parameters for the method 'PostLoanTemplate'
type PostLoanTemplateOpts struct {
	File       optional.Interface
	Locale     optional.String
	DateFormat optional.String
}

/*
PostLoanTemplate Method for PostLoanTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *PostLoanTemplateOpts - Optional Parameters:
 * @param "File" (optional.Interface of FormDataContentDisposition) -
 * @param "Locale" (optional.String) -
 * @param "DateFormat" (optional.String) -
@return string
*/
func (a *LoansApiService) PostLoanTemplate(ctx _context.Context, opts *PostLoanTemplateOpts) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/uploadtemplate"
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

// RetrieveAll26Opts Optional parameters for the method 'RetrieveAll26'
type RetrieveAll26Opts struct {
	SqlSearch  optional.String
	ExternalId optional.String
	Offset     optional.Int32
	Limit      optional.Int32
	OrderBy    optional.String
	SortOrder  optional.String
	AccountNo  optional.String
}

/*
RetrieveAll26 List Loans
The list capability of loans can support pagination and sorting. Example Requests:  loans  loans?fields&#x3D;accountNo  loans?offset&#x3D;10&amp;limit&#x3D;50  loans?orderBy&#x3D;accountNo&amp;sortOrder&#x3D;DESC
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveAll26Opts - Optional Parameters:
 * @param "SqlSearch" (optional.String) -  sqlSearch
 * @param "ExternalId" (optional.String) -  externalId
 * @param "Offset" (optional.Int32) -  offset
 * @param "Limit" (optional.Int32) -  limit
 * @param "OrderBy" (optional.String) -  orderBy
 * @param "SortOrder" (optional.String) -  sortOrder
 * @param "AccountNo" (optional.String) -  accountNo
@return GetLoansResponse
*/
func (a *LoansApiService) RetrieveAll26(ctx _context.Context, opts *RetrieveAll26Opts) (models.GetLoansResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetLoansResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans"
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
	if opts != nil && opts.AccountNo.IsSet() {
		queryParams.Add("accountNo", parameterToString(opts.AccountNo.Value(), ""))
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

// RetrieveApprovalTemplateOpts Optional parameters for the method 'RetrieveApprovalTemplate'
type RetrieveApprovalTemplateOpts struct {
	TemplateType optional.String
}

/*
RetrieveApprovalTemplate Method for RetrieveApprovalTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param loanId loanId
 * @param optional nil or *RetrieveApprovalTemplateOpts - Optional Parameters:
 * @param "TemplateType" (optional.String) -  templateType
@return string
*/
func (a *LoansApiService) RetrieveApprovalTemplate(ctx _context.Context, loanId int64, opts *RetrieveApprovalTemplateOpts) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/{loanId}/template"
	path = strings.Replace(path, "{"+"loanId"+"}", _neturl.QueryEscape(parameterToString(loanId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.TemplateType.IsSet() {
		queryParams.Add("templateType", parameterToString(opts.TemplateType.Value(), ""))
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

// RetrieveLoanOpts Optional parameters for the method 'RetrieveLoan'
type RetrieveLoanOpts struct {
	StaffInSelectedOfficeOnly optional.Bool
}

/*
RetrieveLoan Retrieve a Loan
Note: template&#x3D;true parameter doesn&#39;t apply to this resource.Example Requests:  loans/1   loans/1?fields&#x3D;id,principal,annualInterestRate   loans/1?associations&#x3D;all  loans/1?associations&#x3D;all&amp;exclude&#x3D;guarantors   loans/1?fields&#x3D;id,principal,annualInterestRate&amp;associations&#x3D;repaymentSchedule,transactions
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param loanId loanId
 * @param optional nil or *RetrieveLoanOpts - Optional Parameters:
 * @param "StaffInSelectedOfficeOnly" (optional.Bool) -  staffInSelectedOfficeOnly
@return GetLoansLoanIdResponse
*/
func (a *LoansApiService) RetrieveLoan(ctx _context.Context, loanId int64, opts *RetrieveLoanOpts) (models.GetLoansLoanIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetLoansLoanIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/{loanId}"
	path = strings.Replace(path, "{"+"loanId"+"}", _neturl.QueryEscape(parameterToString(loanId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

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

// StateTransitionsOpts Optional parameters for the method 'StateTransitions'
type StateTransitionsOpts struct {
	Command optional.String
}

/*
StateTransitions Approve Loan Application | Recover Loan Guarantee | Undo Loan Application Approval | Assign a Loan Officer | Unassign a Loan Officer | Reject Loan Application | Applicant Withdraws from Loan Application | Disburse Loan Disburse Loan To Savings Account | Undo Loan Disbursal
Approve Loan Application: Mandatory Fields: approvedOnDate Optional Fields: approvedLoanAmount and expectedDisbursementDate Approves the loan application  Recover Loan Guarantee: Recovers the loan guarantee  Undo Loan Application Approval: Undoes the Loan Application Approval  Assign a Loan Officer: Allows you to assign Loan Officer for existing Loan.  Unassign a Loan Officer: Allows you to unassign the Loan Officer.  Reject Loan Application: Mandatory Fields: rejectedOnDate Allows you to reject the loan application  Applicant Withdraws from Loan Application: Mandatory Fields: withdrawnOnDate Allows the applicant to withdraw the loan application  Disburse Loan: Mandatory Fields: actualDisbursementDate Optional Fields: transactionAmount and fixedEmiAmount Disburses the Loan  Disburse Loan To Savings Account: Mandatory Fields: actualDisbursementDate Optional Fields: transactionAmount and fixedEmiAmount Disburses the loan to Saving Account  Undo Loan Disbursal: Undoes the Loan Disbursal Showing request and response for Assign a Loan Officer
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param loanId loanId
 * @param postLoansLoanIdRequest
 * @param optional nil or *StateTransitionsOpts - Optional Parameters:
 * @param "Command" (optional.String) -  command
@return PostLoansLoanIdResponse
*/
func (a *LoansApiService) StateTransitions(ctx _context.Context, loanId int64, postLoansLoanIdRequest models.PostLoansLoanIdRequest, opts *StateTransitionsOpts) (models.PostLoansLoanIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostLoansLoanIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/{loanId}"
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
	postBody = &postLoansLoanIdRequest
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

// Template10Opts Optional parameters for the method 'Template10'
type Template10Opts struct {
	ClientId                  optional.Int64
	GroupId                   optional.Int64
	ProductId                 optional.Int64
	TemplateType              optional.String
	StaffInSelectedOfficeOnly optional.Bool
	ActiveOnly                optional.Bool
}

/*
Template10 Retrieve Loan Details Template
This is a convenience resource. It can be useful when building maintenance user interface screens for client applications. The template data returned consists of any or all of:  Field Defaults Allowed description Lists Example Requests:  loans/template?templateType&#x3D;individual&amp;clientId&#x3D;1   loans/template?templateType&#x3D;individual&amp;clientId&#x3D;1&amp;productId&#x3D;1
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *Template10Opts - Optional Parameters:
 * @param "ClientId" (optional.Int64) -  clientId
 * @param "GroupId" (optional.Int64) -  groupId
 * @param "ProductId" (optional.Int64) -  productId
 * @param "TemplateType" (optional.String) -  templateType
 * @param "StaffInSelectedOfficeOnly" (optional.Bool) -  staffInSelectedOfficeOnly
 * @param "ActiveOnly" (optional.Bool) -  activeOnly
@return GetLoansTemplateResponse
*/
func (a *LoansApiService) Template10(ctx _context.Context, opts *Template10Opts) (models.GetLoansTemplateResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetLoansTemplateResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/loans/template"
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
	if opts != nil && opts.TemplateType.IsSet() {
		queryParams.Add("templateType", parameterToString(opts.TemplateType.Value(), ""))
	}
	if opts != nil && opts.StaffInSelectedOfficeOnly.IsSet() {
		queryParams.Add("staffInSelectedOfficeOnly", parameterToString(opts.StaffInSelectedOfficeOnly.Value(), ""))
	}
	if opts != nil && opts.ActiveOnly.IsSet() {
		queryParams.Add("activeOnly", parameterToString(opts.ActiveOnly.Value(), ""))
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
