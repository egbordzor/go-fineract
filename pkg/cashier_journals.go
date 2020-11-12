package pkg

import (
	_context "context"
	"github.com/wondenge/go-fineract/pkg/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// CashierJournalsApiService CashierJournalsApi service
type CashierJournalsApiService service

// GetJournalData1Opts Optional parameters for the method 'GetJournalData1'
type GetJournalData1Opts struct {
	OfficeId  optional.Int64
	TellerId  optional.Int64
	CashierId optional.Int64
	DateRange optional.String
}

/*
GetJournalData1 Method for GetJournalData1
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetJournalData1Opts - Optional Parameters:
 * @param "OfficeId" (optional.Int64) -
 * @param "TellerId" (optional.Int64) -
 * @param "CashierId" (optional.Int64) -
 * @param "DateRange" (optional.String) -
@return string
*/
func (a *CashierJournalsApiService) GetJournalData1(ctx _context.Context, opts *GetJournalData1Opts) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/cashiersjournal"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.OfficeId.IsSet() {
		queryParams.Add("officeId", parameterToString(opts.OfficeId.Value(), ""))
	}
	if opts != nil && opts.TellerId.IsSet() {
		queryParams.Add("tellerId", parameterToString(opts.TellerId.Value(), ""))
	}
	if opts != nil && opts.CashierId.IsSet() {
		queryParams.Add("cashierId", parameterToString(opts.CashierId.Value(), ""))
	}
	if opts != nil && opts.DateRange.IsSet() {
		queryParams.Add("dateRange", parameterToString(opts.DateRange.Value(), ""))
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
