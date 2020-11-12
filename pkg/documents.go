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

// DocumentsApiService DocumentsApi service
type DocumentsApiService service

// CreateDocumentOpts Optional parameters for the method 'CreateDocument'
type CreateDocumentOpts struct {
	ContentLength optional.Int64
	File          optional.Interface
	Name          optional.String
	Description   optional.String
}

/*
CreateDocument Create a Document
Note: A document is created using a Multi-part form upload   Body Parts  name :  Name or summary of the document  description :  Description of the document  file :  The file to be uploaded  Mandatory Fields :  file and description
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entityType entityType
 * @param entityId entityId
 * @param optional nil or *CreateDocumentOpts - Optional Parameters:
 * @param "ContentLength" (optional.Int64) -  Content-Length
 * @param "File" (optional.Interface of FormDataBodyPart) -
 * @param "Name" (optional.String) -  name
 * @param "Description" (optional.String) -  description
@return PostEntityTypeEntityIdDocumentsResponse
*/
func (a *DocumentsApiService) CreateDocument(ctx _context.Context, entityType string, entityId int64, opts *CreateDocumentOpts) (models.PostEntityTypeEntityIdDocumentsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostEntityTypeEntityIdDocumentsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/{entityType}/{entityId}/documents"
	path = strings.Replace(path, "{"+"entityType"+"}", _neturl.QueryEscape(parameterToString(entityType, "")), -1)

	path = strings.Replace(path, "{"+"entityId"+"}", _neturl.QueryEscape(parameterToString(entityId, "")), -1)

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
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	if opts != nil && opts.ContentLength.IsSet() {
		headerParams["Content-Length"] = parameterToString(opts.ContentLength.Value(), "")
	}
	if opts != nil && opts.File.IsSet() {
		paramJson, err := parameterToJson(opts.File.Value())
		if err != nil {
			return returnValue, nil, err
		}
		formParams.Add("file", paramJson)
	}
	if opts != nil && opts.Name.IsSet() {
		formParams.Add("name", parameterToString(opts.Name.Value(), ""))
	}
	if opts != nil && opts.Description.IsSet() {
		formParams.Add("description", parameterToString(opts.Description.Value(), ""))
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
DeleteDocument Remove a Document
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entityType entityType
 * @param entityId entityId
 * @param documentId documentId
@return DeleteEntityTypeEntityIdDocumentsResponse
*/
func (a *DocumentsApiService) DeleteDocument(ctx _context.Context, entityType string, entityId int64, documentId int64) (models.DeleteEntityTypeEntityIdDocumentsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodDelete
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.DeleteEntityTypeEntityIdDocumentsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/{entityType}/{entityId}/documents/{documentId}"
	path = strings.Replace(path, "{"+"entityType"+"}", _neturl.QueryEscape(parameterToString(entityType, "")), -1)

	path = strings.Replace(path, "{"+"entityId"+"}", _neturl.QueryEscape(parameterToString(entityId, "")), -1)

	path = strings.Replace(path, "{"+"documentId"+"}", _neturl.QueryEscape(parameterToString(documentId, "")), -1)

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
DownloadFile Retrieve Binary File associated with Document
Request used to download the file associated with the document  Example Requests:  clients/1/documents/1/attachment   loans/1/documents/1/attachment
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entityType entityType
 * @param entityId entityId
 * @param documentId documentId
*/
func (a *DocumentsApiService) DownloadFile(ctx _context.Context, entityType string, entityId int64, documentId int64) (*_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/{entityType}/{entityId}/documents/{documentId}/attachment"
	path = strings.Replace(path, "{"+"entityType"+"}", _neturl.QueryEscape(parameterToString(entityType, "")), -1)

	path = strings.Replace(path, "{"+"entityId"+"}", _neturl.QueryEscape(parameterToString(entityId, "")), -1)

	path = strings.Replace(path, "{"+"documentId"+"}", _neturl.QueryEscape(parameterToString(documentId, "")), -1)

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
	headerAccepts := []string{}

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

/*
GetDocument Retrieve a Document
Example Requests:  clients/1/documents/1   loans/1/documents/1   client_identifiers/1/documents/1?fields&#x3D;name,description
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entityType entityType
 * @param entityId entityId
 * @param documentId documentId
@return GetEntityTypeEntityIdDocumentsResponse
*/
func (a *DocumentsApiService) GetDocument(ctx _context.Context, entityType string, entityId int64, documentId int64) (models.GetEntityTypeEntityIdDocumentsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetEntityTypeEntityIdDocumentsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/{entityType}/{entityId}/documents/{documentId}"
	path = strings.Replace(path, "{"+"entityType"+"}", _neturl.QueryEscape(parameterToString(entityType, "")), -1)

	path = strings.Replace(path, "{"+"entityId"+"}", _neturl.QueryEscape(parameterToString(entityId, "")), -1)

	path = strings.Replace(path, "{"+"documentId"+"}", _neturl.QueryEscape(parameterToString(documentId, "")), -1)

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
RetrieveAllDocuments List documents
Example Requests:  clients/1/documents  client_identifiers/1/documents  loans/1/documents?fields&#x3D;name,description
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entityType entityType
 * @param entityId entityId
@return []GetEntityTypeEntityIdDocumentsResponse
*/
func (a *DocumentsApiService) RetrieveAllDocuments(ctx _context.Context, entityType string, entityId int64) ([]models.GetEntityTypeEntityIdDocumentsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  []models.GetEntityTypeEntityIdDocumentsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/{entityType}/{entityId}/documents"
	path = strings.Replace(path, "{"+"entityType"+"}", _neturl.QueryEscape(parameterToString(entityType, "")), -1)

	path = strings.Replace(path, "{"+"entityId"+"}", _neturl.QueryEscape(parameterToString(entityId, "")), -1)

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

// UpdateDocumentOpts Optional parameters for the method 'UpdateDocument'
type UpdateDocumentOpts struct {
	ContentLength optional.Int64
	File          optional.Interface
	Name          optional.String
	Description   optional.String
}

/*
UpdateDocument Update a Document
Note: A document is updated using a Multi-part form upload  Body Parts name Name or summary of the document description Description of the document file The file to be uploaded
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entityType entityType
 * @param entityId entityId
 * @param documentId documentId
 * @param optional nil or *UpdateDocumentOpts - Optional Parameters:
 * @param "ContentLength" (optional.Int64) -  Content-Length
 * @param "File" (optional.Interface of FormDataBodyPart) -
 * @param "Name" (optional.String) -  name
 * @param "Description" (optional.String) -  description
@return PutEntityTypeEntityIdDocumentsResponse
*/
func (a *DocumentsApiService) UpdateDocument(ctx _context.Context, entityType string, entityId int64, documentId int64, opts *UpdateDocumentOpts) (models.PutEntityTypeEntityIdDocumentsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPut
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PutEntityTypeEntityIdDocumentsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/{entityType}/{entityId}/documents/{documentId}"
	path = strings.Replace(path, "{"+"entityType"+"}", _neturl.QueryEscape(parameterToString(entityType, "")), -1)

	path = strings.Replace(path, "{"+"entityId"+"}", _neturl.QueryEscape(parameterToString(entityId, "")), -1)

	path = strings.Replace(path, "{"+"documentId"+"}", _neturl.QueryEscape(parameterToString(documentId, "")), -1)

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
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	if opts != nil && opts.ContentLength.IsSet() {
		headerParams["Content-Length"] = parameterToString(opts.ContentLength.Value(), "")
	}
	if opts != nil && opts.File.IsSet() {
		paramJson, err := parameterToJson(opts.File.Value())
		if err != nil {
			return returnValue, nil, err
		}
		formParams.Add("file", paramJson)
	}
	if opts != nil && opts.Name.IsSet() {
		formParams.Add("name", parameterToString(opts.Name.Value(), ""))
	}
	if opts != nil && opts.Description.IsSet() {
		formParams.Add("description", parameterToString(opts.Description.Value(), ""))
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
