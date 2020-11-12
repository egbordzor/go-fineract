package pkg

import (
	_context "context"
	"github.com/wondenge/go-fineract/pkg/models"
	"github.com/wondenge/go-fineract/pkg/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// SelfUserDetailsApiService SelfUserDetailsApi service
type SelfUserDetailsApiService service

// FetchAuthenticatedUserData1Opts Optional parameters for the method 'FetchAuthenticatedUserData1'
type FetchAuthenticatedUserData1Opts struct {
	AccessToken optional.String
}

/*
FetchAuthenticatedUserData1 Fetch authenticated user details
Checks the Authentication and returns the set roles and permissions allowed  For more info visit this link - https://demo.fineract.dev/fineract-provider/api-docs/apiLive.htm#selfoauth
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *FetchAuthenticatedUserData1Opts - Optional Parameters:
 * @param "AccessToken" (optional.String) -  äccess_token
@return GetSelfUserDetailsResponse
*/
func (a *SelfUserDetailsApiService) FetchAuthenticatedUserData1(ctx _context.Context, opts *FetchAuthenticatedUserData1Opts) (models.GetSelfUserDetailsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetSelfUserDetailsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/self/userdetails"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.AccessToken.IsSet() {
		queryParams.Add("access_token", parameterToString(opts.AccessToken.Value(), ""))
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
