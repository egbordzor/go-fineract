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

// ProvisioningEntriesApiService ProvisioningEntriesApi service
type ProvisioningEntriesApiService service

// CreateProvisioningEntriesOpts Optional parameters for the method 'CreateProvisioningEntries'
type CreateProvisioningEntriesOpts struct {
	PostProvisioningEntriesRequest optional.Interface
}

/*
CreateProvisioningEntries Create new Provisioning Entries
Creates a new Provisioning Entries  Mandatory Fields date dateFormat locale Optional Fields createjournalentries
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CreateProvisioningEntriesOpts - Optional Parameters:
 * @param "PostProvisioningEntriesRequest" (optional.Interface of PostProvisioningEntriesRequest) -
@return PostProvisioningEntriesResponse
*/
func (a *ProvisioningEntriesApiService) CreateProvisioningEntries(ctx _context.Context, opts *CreateProvisioningEntriesOpts) (models.PostProvisioningEntriesResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostProvisioningEntriesResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/provisioningentries"
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
	if opts != nil && opts.PostProvisioningEntriesRequest.IsSet() {
		optPostProvisioningEntriesRequest, optPostProvisioningEntriesRequestok := opts.PostProvisioningEntriesRequest.Value().(models.PostProvisioningEntriesRequest)
		if !optPostProvisioningEntriesRequestok {
			return returnValue, nil, reportError("postProvisioningEntriesRequest should be PostProvisioningEntriesRequest")
		}
		postBody = &optPostProvisioningEntriesRequest
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

// ModifyProvisioningEntryOpts Optional parameters for the method 'ModifyProvisioningEntry'
type ModifyProvisioningEntryOpts struct {
	Command                       optional.String
	PutProvisioningEntriesRequest optional.Interface
}

/*
ModifyProvisioningEntry Recreates Provisioning Entry
Recreates Provisioning Entry | createjournalentry.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entryId entryId
 * @param optional nil or *ModifyProvisioningEntryOpts - Optional Parameters:
 * @param "Command" (optional.String) -  command=createjournalentry command=recreateprovisioningentry
 * @param "PutProvisioningEntriesRequest" (optional.Interface of PutProvisioningEntriesRequest) -
@return PutProvisioningEntriesResponse
*/
func (a *ProvisioningEntriesApiService) ModifyProvisioningEntry(ctx _context.Context, entryId int64, opts *ModifyProvisioningEntryOpts) (models.PutProvisioningEntriesResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PutProvisioningEntriesResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/provisioningentries/{entryId}"
	path = strings.Replace(path, "{"+"entryId"+"}", _neturl.QueryEscape(parameterToString(entryId, "")), -1)

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
	if opts != nil && opts.PutProvisioningEntriesRequest.IsSet() {
		optPutProvisioningEntriesRequest, optPutProvisioningEntriesRequestok := opts.PutProvisioningEntriesRequest.Value().(models.PutProvisioningEntriesRequest)
		if !optPutProvisioningEntriesRequestok {
			return returnValue, nil, reportError("putProvisioningEntriesRequest should be PutProvisioningEntriesRequest")
		}
		postBody = &optPutProvisioningEntriesRequest
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

// RetrieveAllProvisioningEntriesOpts Optional parameters for the method 'RetrieveAllProvisioningEntries'
type RetrieveAllProvisioningEntriesOpts struct {
	Offset optional.Int32
	Limit  optional.Int32
}

/*
RetrieveAllProvisioningEntries List all Provisioning Entries
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveAllProvisioningEntriesOpts - Optional Parameters:
 * @param "Offset" (optional.Int32) -  offset
 * @param "Limit" (optional.Int32) -  limit
@return []ProvisioningEntryData
*/
func (a *ProvisioningEntriesApiService) RetrieveAllProvisioningEntries(ctx _context.Context, opts *RetrieveAllProvisioningEntriesOpts) ([]models.ProvisioningEntryData, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  []models.ProvisioningEntryData
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/provisioningentries"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.Offset.IsSet() {
		queryParams.Add("offset", parameterToString(opts.Offset.Value(), ""))
	}
	if opts != nil && opts.Limit.IsSet() {
		queryParams.Add("limit", parameterToString(opts.Limit.Value(), ""))
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

// RetrieveProviioningEntriesOpts Optional parameters for the method 'RetrieveProviioningEntries'
type RetrieveProviioningEntriesOpts struct {
	EntryId    optional.Int64
	Offset     optional.Int32
	Limit      optional.Int32
	OfficeId   optional.Int64
	ProductId  optional.Int64
	CategoryId optional.Int64
}

/*
RetrieveProviioningEntries Method for RetrieveProviioningEntries
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveProviioningEntriesOpts - Optional Parameters:
 * @param "EntryId" (optional.Int64) -
 * @param "Offset" (optional.Int32) -
 * @param "Limit" (optional.Int32) -
 * @param "OfficeId" (optional.Int64) -
 * @param "ProductId" (optional.Int64) -
 * @param "CategoryId" (optional.Int64) -
@return LoanProductProvisioningEntryData
*/
func (a *ProvisioningEntriesApiService) RetrieveProviioningEntries(ctx _context.Context, opts *RetrieveProviioningEntriesOpts) (models.LoanProductProvisioningEntryData, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.LoanProductProvisioningEntryData
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/provisioningentries/entries"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.EntryId.IsSet() {
		queryParams.Add("entryId", parameterToString(opts.EntryId.Value(), ""))
	}
	if opts != nil && opts.Offset.IsSet() {
		queryParams.Add("offset", parameterToString(opts.Offset.Value(), ""))
	}
	if opts != nil && opts.Limit.IsSet() {
		queryParams.Add("limit", parameterToString(opts.Limit.Value(), ""))
	}
	if opts != nil && opts.OfficeId.IsSet() {
		queryParams.Add("officeId", parameterToString(opts.OfficeId.Value(), ""))
	}
	if opts != nil && opts.ProductId.IsSet() {
		queryParams.Add("productId", parameterToString(opts.ProductId.Value(), ""))
	}
	if opts != nil && opts.CategoryId.IsSet() {
		queryParams.Add("categoryId", parameterToString(opts.CategoryId.Value(), ""))
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
RetrieveProvisioningEntry Retrieves a Provisioning Entry
Returns the details of a generated Provisioning Entry.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entryId entryId
@return ProvisioningEntryData
*/
func (a *ProvisioningEntriesApiService) RetrieveProvisioningEntry(ctx _context.Context, entryId int64) (models.ProvisioningEntryData, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.ProvisioningEntryData
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/provisioningentries/{entryId}"
	path = strings.Replace(path, "{"+"entryId"+"}", _neturl.QueryEscape(parameterToString(entryId, "")), -1)

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
