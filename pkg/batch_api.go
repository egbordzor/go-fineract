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

// BatchAPIApiService BatchAPIApi service
type BatchAPIApiService service

// HandleBatchRequestsOpts Optional parameters for the method 'HandleBatchRequests'
type HandleBatchRequestsOpts struct {
	EnclosingTransaction optional.Bool
}

/*
HandleBatchRequests Batch requests in a single transaction
The Apache Fineract Batch API is also capable of executing all the requests in a single transaction, by setting a Query Parameter, \&quot;enclosingTransaction&#x3D;true\&quot;. So, if one or more of the requests in a batch returns an erroneous response all of the Data base transactions made by other successful requests will be rolled back.  If there has been a rollback in a transaction then a single response will be provided, with a &#39;400&#39; status code and a body consisting of the error details of the first failed request.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param postBatchesRequest
 * @param optional nil or *HandleBatchRequestsOpts - Optional Parameters:
 * @param "EnclosingTransaction" (optional.Bool) -  enclosingTransaction
@return BatchResponse
*/
func (a *BatchAPIApiService) HandleBatchRequests(ctx _context.Context, postBatchesRequest models.PostBatchesRequest, opts *HandleBatchRequestsOpts) (models.BatchResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.BatchResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/batches"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.EnclosingTransaction.IsSet() {
		queryParams.Add("enclosingTransaction", parameterToString(opts.EnclosingTransaction.Value(), ""))
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
	postBody = &postBatchesRequest
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
