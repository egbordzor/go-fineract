package pkg

import (
	_context "context"
	"github.com/antihax/optional"
	"github.com/wondenge/go-fineract/pkg/models"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// StandingInstructionsHistoryApiService StandingInstructionsHistoryApi service
type StandingInstructionsHistoryApiService service

// RetrieveAll19Opts Optional parameters for the method 'RetrieveAll19'
type RetrieveAll19Opts struct {
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
	Locale          optional.String
	DateFormat      optional.String
	FromDate        optional.Interface
	ToDate          optional.Interface
}

/*
RetrieveAll19 Standing Instructions Logged History
The list capability of history can support pagination and sorting   Example Requests :  standinginstructionrunhistory  standinginstructionrunhistory?orderBy&#x3D;name&amp;sortOrder&#x3D;DESC  standinginstructionrunhistory?offset&#x3D;10&amp;limit&#x3D;50
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveAll19Opts - Optional Parameters:
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
 * @param "Locale" (optional.String) -  locale
 * @param "DateFormat" (optional.String) -  dateFormat
 * @param "FromDate" (optional.Interface of map[string]interface{}) -  fromDate
 * @param "ToDate" (optional.Interface of map[string]interface{}) -  toDate
@return GetStandingInstructionRunHistoryResponse
*/
func (a *StandingInstructionsHistoryApiService) RetrieveAll19(ctx _context.Context, opts *RetrieveAll19Opts) (models.GetStandingInstructionRunHistoryResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetStandingInstructionRunHistoryResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/standinginstructionrunhistory"
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
	if opts != nil && opts.Locale.IsSet() {
		queryParams.Add("locale", parameterToString(opts.Locale.Value(), ""))
	}
	if opts != nil && opts.DateFormat.IsSet() {
		queryParams.Add("dateFormat", parameterToString(opts.DateFormat.Value(), ""))
	}
	if opts != nil && opts.FromDate.IsSet() {
		queryParams.Add("fromDate", parameterToString(opts.FromDate.Value(), ""))
	}
	if opts != nil && opts.ToDate.IsSet() {
		queryParams.Add("toDate", parameterToString(opts.ToDate.Value(), ""))
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
