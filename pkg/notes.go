package pkg

import (
	_context "context"
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

// NotesApiService NotesApi service
type NotesApiService service

/*
AddNewNote Add a Resource Note
Adds a new note to a supported resource.  Example Requests:  clients/1/notes   groups/1/notes
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param resourceType resourceType
 * @param resourceId resourceId
 * @param postResourceTypeResourceIdNotesRequest
@return PostResourceTypeResourceIdNotesResponse
*/
func (a *NotesApiService) AddNewNote(ctx _context.Context, resourceType string, resourceId int64, postResourceTypeResourceIdNotesRequest models.PostResourceTypeResourceIdNotesRequest) (models.PostResourceTypeResourceIdNotesResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostResourceTypeResourceIdNotesResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/{resourceType}/{resourceId}/notes"
	path = strings.Replace(path, "{"+"resourceType"+"}", _neturl.QueryEscape(parameterToString(resourceType, "")), -1)

	path = strings.Replace(path, "{"+"resourceId"+"}", _neturl.QueryEscape(parameterToString(resourceId, "")), -1)

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
	postBody = &postResourceTypeResourceIdNotesRequest
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
DeleteNote Delete a Resource Note
Deletes a Resource Note
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param resourceType resourceType
 * @param resourceId resourceId
 * @param noteId noteId
@return DeleteResourceTypeResourceIdNotesNoteIdResponse
*/
func (a *NotesApiService) DeleteNote(ctx _context.Context, resourceType string, resourceId int64, noteId int64) (models.DeleteResourceTypeResourceIdNotesNoteIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodDelete
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.DeleteResourceTypeResourceIdNotesNoteIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/{resourceType}/{resourceId}/notes/{noteId}"
	path = strings.Replace(path, "{"+"resourceType"+"}", _neturl.QueryEscape(parameterToString(resourceType, "")), -1)

	path = strings.Replace(path, "{"+"resourceId"+"}", _neturl.QueryEscape(parameterToString(resourceId, "")), -1)

	path = strings.Replace(path, "{"+"noteId"+"}", _neturl.QueryEscape(parameterToString(noteId, "")), -1)

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
RetrieveNote Retrieve a Resource Note
Retrieves a Resource Note  Example Requests:  clients/1/notes/76   groups/1/notes/20   clients/1/notes/76?fields&#x3D;note,createdOn,createdByUsername   groups/1/notes/20?fields&#x3D;note,createdOn,createdByUsername
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param resourceType resourceType
 * @param resourceId resourceId
 * @param noteId noteId
@return GetResourceTypeResourceIdNotesNoteIdResponse
*/
func (a *NotesApiService) RetrieveNote(ctx _context.Context, resourceType string, resourceId int64, noteId int64) (models.GetResourceTypeResourceIdNotesNoteIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetResourceTypeResourceIdNotesNoteIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/{resourceType}/{resourceId}/notes/{noteId}"
	path = strings.Replace(path, "{"+"resourceType"+"}", _neturl.QueryEscape(parameterToString(resourceType, "")), -1)

	path = strings.Replace(path, "{"+"resourceId"+"}", _neturl.QueryEscape(parameterToString(resourceId, "")), -1)

	path = strings.Replace(path, "{"+"noteId"+"}", _neturl.QueryEscape(parameterToString(noteId, "")), -1)

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
RetrieveNotesByResource Retrieve a Resource's description
Retrieves a Resource&#39;s Notes  Note: Notes are returned in descending createOn order.  Example Requests:  clients/2/notes   groups/2/notes?fields&#x3D;note,createdOn,createdByUsername
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param resourceType resourceType
 * @param resourceId resourceId
@return []GetResourceTypeResourceIdNotesResponse
*/
func (a *NotesApiService) RetrieveNotesByResource(ctx _context.Context, resourceType string, resourceId int64) ([]models.GetResourceTypeResourceIdNotesResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  []models.GetResourceTypeResourceIdNotesResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/{resourceType}/{resourceId}/notes"
	path = strings.Replace(path, "{"+"resourceType"+"}", _neturl.QueryEscape(parameterToString(resourceType, "")), -1)

	path = strings.Replace(path, "{"+"resourceId"+"}", _neturl.QueryEscape(parameterToString(resourceId, "")), -1)

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
UpdateNote Update a Resource Note
Updates a Resource Note
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param resourceType resourceType
 * @param resourceId resourceId
 * @param noteId noteId
 * @param putResourceTypeResourceIdNotesNoteIdRequest
@return PutResourceTypeResourceIdNotesNoteIdResponse
*/
func (a *NotesApiService) UpdateNote(ctx _context.Context, resourceType string, resourceId int64, noteId int64, putResourceTypeResourceIdNotesNoteIdRequest models.PutResourceTypeResourceIdNotesNoteIdRequest) (models.PutResourceTypeResourceIdNotesNoteIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPut
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PutResourceTypeResourceIdNotesNoteIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/{resourceType}/{resourceId}/notes/{noteId}"
	path = strings.Replace(path, "{"+"resourceType"+"}", _neturl.QueryEscape(parameterToString(resourceType, "")), -1)

	path = strings.Replace(path, "{"+"resourceId"+"}", _neturl.QueryEscape(parameterToString(resourceId, "")), -1)

	path = strings.Replace(path, "{"+"noteId"+"}", _neturl.QueryEscape(parameterToString(noteId, "")), -1)

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
	postBody = &putResourceTypeResourceIdNotesNoteIdRequest
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
