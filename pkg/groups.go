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

// GroupsApiService GroupsApi service
type GroupsApiService service

// ActivateOrGenerateCollectionSheetOpts Optional parameters for the method 'ActivateOrGenerateCollectionSheet'
type ActivateOrGenerateCollectionSheetOpts struct {
	Command optional.String
	RoleId  optional.Int64
}

/*
ActivateOrGenerateCollectionSheet Activate a Group | Associate Clients | Disassociate Clients | Transfer Clients across groups | Generate Collection Sheet | Save Collection Sheet | Unassign a Staff | Assign a Staff | Close a Group | Unassign a Role | Update a Role
Activate a Group:  Groups can be created in a Pending state. This API exists to enable group activation.    If the group happens to be already active this API will result in an error.  Mandatory Fields: activationDate  Associate Clients:  This API allows to associate existing clients to a group.    The clients are listed from the office to which the group is associated.    If client(s) is already associated with group then API will result in an error.  Mandatory Fields: clientMembers  Disassociate Clients:  This API allows to disassociate clients from a group.    Disassociating a client with active joint liability group loans results in an error.  Mandatory Fields: clientMembers  Transfer Clients across groups:  This API allows to transfer clients from one group to another  Mandatory Fields: destinationGroupId and clients  Optional Fields: inheritDestinationGroupLoanOfficer (defaults to true) and transferActiveLoans (defaults to true)  Generate Collection Sheet:  This API retrieves repayment details of all jlg loans of all members of a group on a specified meeting date.  Mandatory Fields: calendarId and transactionDate  Save Collection Sheet:  This api allows the loan officer to perform bulk repayments of JLG loans for a group on its meeting date.  Mandatory Fields: calendarId, transactionDate, actualDisbursementDate  Optional Fields: clientsAttendance, bulkRepaymentTransaction, bulkDisbursementTransactions  Unassign a Staff:  Allows you to unassign the Staff.  Mandatory Fields: staffId  Assign a Staff:  Allows you to assign Staff to an existing Group.    The selected Staff should be belong to the same office (or an office higher up in the hierarchy) as this groupMandatory Fields: staffId  Optional Fields: inheritStaffForClientAccounts (Optional: Boolean if true all members of the group (i.e all clients with active loans and savings ) will inherit the staffId)  Close a Group:  This API exists to close a group. Groups can be closed if they don&#39;t have any non-closed clients/loans/savingsAccounts.    If the group has any active clients/loans/savingsAccount, this API will result in an error.Assign a Role:  Allows you to assign a Role to an existing member of a group.    We can define the different roles applicable to group members by adding code values to the pre-defined system code GROUPROLE. Example:Group leader etc.  Mandatory Fields: clientId, role  Unassign a Role:  Allows you to unassign Roles associated tp Group members.  Update a Role:  Allows you to update the member Role.  Mandatory Fields: role  Showing request/response for Transfer Clients across groups
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param groupId groupId
 * @param postGroupsGroupIdRequest
 * @param optional nil or *ActivateOrGenerateCollectionSheetOpts - Optional Parameters:
 * @param "Command" (optional.String) -  command
 * @param "RoleId" (optional.Int64) -  roleId
@return PostGroupsGroupIdResponse
*/
func (a *GroupsApiService) ActivateOrGenerateCollectionSheet(ctx _context.Context, groupId int64, postGroupsGroupIdRequest models.PostGroupsGroupIdRequest, opts *ActivateOrGenerateCollectionSheetOpts) (models.PostGroupsGroupIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostGroupsGroupIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/groups/{groupId}"
	path = strings.Replace(path, "{"+"groupId"+"}", _neturl.QueryEscape(parameterToString(groupId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.Command.IsSet() {
		queryParams.Add("command", parameterToString(opts.Command.Value(), ""))
	}
	if opts != nil && opts.RoleId.IsSet() {
		queryParams.Add("roleId", parameterToString(opts.RoleId.Value(), ""))
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
	postBody = &postGroupsGroupIdRequest
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
Create7 Create a Group
Creates a Group  Mandatory Fields: name, officeId, active, activationDate (if active&#x3D;true)  Optional Fields: externalId, staffId, clientMembers
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param postGroupsRequest
@return PostGroupsResponse
*/
func (a *GroupsApiService) Create7(ctx _context.Context, postGroupsRequest models.PostGroupsRequest) (models.PostGroupsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostGroupsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/groups"
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
	postBody = &postGroupsRequest
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
Delete11 Delete a Group
A group can be deleted if it is in pending state and has no associations - clients, loans or savings
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param groupId groupId
@return DeleteGroupsGroupIdResponse
*/
func (a *GroupsApiService) Delete11(ctx _context.Context, groupId int64) (models.DeleteGroupsGroupIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodDelete
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.DeleteGroupsGroupIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/groups/{groupId}"
	path = strings.Replace(path, "{"+"groupId"+"}", _neturl.QueryEscape(parameterToString(groupId, "")), -1)

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

// GetGroupsTemplateOpts Optional parameters for the method 'GetGroupsTemplate'
type GetGroupsTemplateOpts struct {
	OfficeId   optional.Int64
	StaffId    optional.Int64
	DateFormat optional.String
}

/*
GetGroupsTemplate Method for GetGroupsTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetGroupsTemplateOpts - Optional Parameters:
 * @param "OfficeId" (optional.Int64) -
 * @param "StaffId" (optional.Int64) -
 * @param "DateFormat" (optional.String) -
*/
func (a *GroupsApiService) GetGroupsTemplate(ctx _context.Context, opts *GetGroupsTemplateOpts) (*_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/groups/downloadtemplate"
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

// PostGroupTemplateOpts Optional parameters for the method 'PostGroupTemplate'
type PostGroupTemplateOpts struct {
	File       optional.Interface
	Locale     optional.String
	DateFormat optional.String
}

/*
PostGroupTemplate Method for PostGroupTemplate
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *PostGroupTemplateOpts - Optional Parameters:
 * @param "File" (optional.Interface of FormDataContentDisposition) -
 * @param "Locale" (optional.String) -
 * @param "DateFormat" (optional.String) -
@return string
*/
func (a *GroupsApiService) PostGroupTemplate(ctx _context.Context, opts *PostGroupTemplateOpts) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/groups/uploadtemplate"
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

/*
RetrieveAccounts Retrieve Group accounts overview
Retrieves details of all Loan and Savings accounts associated with this group.    Example Requests:    groups/1/accounts      groups/1/accounts?fields&#x3D;loanAccounts,savingsAccounts,memberLoanAccounts,  memberSavingsAccounts
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param groupId groupId
@return GetGroupsGroupIdAccountsResponse
*/
func (a *GroupsApiService) RetrieveAccounts(ctx _context.Context, groupId int64) (models.GetGroupsGroupIdAccountsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetGroupsGroupIdAccountsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/groups/{groupId}/accounts"
	path = strings.Replace(path, "{"+"groupId"+"}", _neturl.QueryEscape(parameterToString(groupId, "")), -1)

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

// RetrieveAll23Opts Optional parameters for the method 'RetrieveAll23'
type RetrieveAll23Opts struct {
	OfficeId       optional.Int64
	StaffId        optional.Int64
	ExternalId     optional.String
	Name           optional.String
	UnderHierarchy optional.String
	Paged          optional.Bool
	Offset         optional.Int32
	Limit          optional.Int32
	OrderBy        optional.String
	SortOrder      optional.String
	OrphansOnly    optional.Bool
}

/*
RetrieveAll23 List Groups
The default implementation of listing Groups returns 200 entries with support for pagination and sorting. Using the parameter limit with description -1 returns all entries.  Example Requests:    groups    groups?fields&#x3D;name,officeName,joinedDate    groups?offset&#x3D;10&amp;limit&#x3D;50    groups?orderBy&#x3D;name&amp;sortOrder&#x3D;DESC
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveAll23Opts - Optional Parameters:
 * @param "OfficeId" (optional.Int64) -  officeId
 * @param "StaffId" (optional.Int64) -  staffId
 * @param "ExternalId" (optional.String) -  externalId
 * @param "Name" (optional.String) -  name
 * @param "UnderHierarchy" (optional.String) -  underHierarchy
 * @param "Paged" (optional.Bool) -  paged
 * @param "Offset" (optional.Int32) -  offset
 * @param "Limit" (optional.Int32) -  limit
 * @param "OrderBy" (optional.String) -  orderBy
 * @param "SortOrder" (optional.String) -  sortOrder
 * @param "OrphansOnly" (optional.Bool) -  orphansOnly
@return GetGroupsResponse
*/
func (a *GroupsApiService) RetrieveAll23(ctx _context.Context, opts *RetrieveAll23Opts) (models.GetGroupsResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetGroupsResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/groups"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.OfficeId.IsSet() {
		queryParams.Add("officeId", parameterToString(opts.OfficeId.Value(), ""))
	}
	if opts != nil && opts.StaffId.IsSet() {
		queryParams.Add("staffId", parameterToString(opts.StaffId.Value(), ""))
	}
	if opts != nil && opts.ExternalId.IsSet() {
		queryParams.Add("externalId", parameterToString(opts.ExternalId.Value(), ""))
	}
	if opts != nil && opts.Name.IsSet() {
		queryParams.Add("name", parameterToString(opts.Name.Value(), ""))
	}
	if opts != nil && opts.UnderHierarchy.IsSet() {
		queryParams.Add("underHierarchy", parameterToString(opts.UnderHierarchy.Value(), ""))
	}
	if opts != nil && opts.Paged.IsSet() {
		queryParams.Add("paged", parameterToString(opts.Paged.Value(), ""))
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
	if opts != nil && opts.OrphansOnly.IsSet() {
		queryParams.Add("orphansOnly", parameterToString(opts.OrphansOnly.Value(), ""))
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

// RetrieveGsimAccountsOpts Optional parameters for the method 'RetrieveGsimAccounts'
type RetrieveGsimAccountsOpts struct {
	ParentGSIMAccountNo optional.String
	ParentGSIMId        optional.Int64
}

/*
RetrieveGsimAccounts Method for RetrieveGsimAccounts
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param groupId
 * @param optional nil or *RetrieveGsimAccountsOpts - Optional Parameters:
 * @param "ParentGSIMAccountNo" (optional.String) -
 * @param "ParentGSIMId" (optional.Int64) -
@return string
*/
func (a *GroupsApiService) RetrieveGsimAccounts(ctx _context.Context, groupId int64, opts *RetrieveGsimAccountsOpts) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/groups/{groupId}/gsimaccounts"
	path = strings.Replace(path, "{"+"groupId"+"}", _neturl.QueryEscape(parameterToString(groupId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.ParentGSIMAccountNo.IsSet() {
		queryParams.Add("parentGSIMAccountNo", parameterToString(opts.ParentGSIMAccountNo.Value(), ""))
	}
	if opts != nil && opts.ParentGSIMId.IsSet() {
		queryParams.Add("parentGSIMId", parameterToString(opts.ParentGSIMId.Value(), ""))
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

// RetrieveOne13Opts Optional parameters for the method 'RetrieveOne13'
type RetrieveOne13Opts struct {
	StaffInSelectedOfficeOnly optional.Bool
	RoleId                    optional.Int64
}

/*
RetrieveOne13 Retrieve a Group
Retrieve group information.  Example Requests:    groups/1    groups/1?associations&#x3D;clientMembers
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param groupId groupId
 * @param optional nil or *RetrieveOne13Opts - Optional Parameters:
 * @param "StaffInSelectedOfficeOnly" (optional.Bool) -  staffInSelectedOfficeOnly
 * @param "RoleId" (optional.Int64) -  roleId
@return GetGroupsGroupIdResponse
*/
func (a *GroupsApiService) RetrieveOne13(ctx _context.Context, groupId int64, opts *RetrieveOne13Opts) (models.GetGroupsGroupIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetGroupsGroupIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/groups/{groupId}"
	path = strings.Replace(path, "{"+"groupId"+"}", _neturl.QueryEscape(parameterToString(groupId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.StaffInSelectedOfficeOnly.IsSet() {
		queryParams.Add("staffInSelectedOfficeOnly", parameterToString(opts.StaffInSelectedOfficeOnly.Value(), ""))
	}
	if opts != nil && opts.RoleId.IsSet() {
		queryParams.Add("roleId", parameterToString(opts.RoleId.Value(), ""))
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

// RetrieveTemplate7Opts Optional parameters for the method 'RetrieveTemplate7'
type RetrieveTemplate7Opts struct {
	OfficeId                  optional.Int64
	Center                    optional.Bool
	CenterId                  optional.Int64
	Command                   optional.String
	StaffInSelectedOfficeOnly optional.Bool
}

/*
RetrieveTemplate7 Retrieve Group Template
This is a convenience resource. It can be useful when building maintenance user interface screens for client applications. The template data returned consists of any or all of:    Field Defaults  Allowed Value Lists  Example Requests:    groups/template    groups/template?officeId&#x3D;2    groups/template?centerId&#x3D;1    groups/template?centerId&#x3D;1&amp;staffInSelectedOfficeOnly&#x3D;true
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RetrieveTemplate7Opts - Optional Parameters:
 * @param "OfficeId" (optional.Int64) -  officeId
 * @param "Center" (optional.Bool) -  center
 * @param "CenterId" (optional.Int64) -  centerId
 * @param "Command" (optional.String) -  command
 * @param "StaffInSelectedOfficeOnly" (optional.Bool) -  staffInSelectedOfficeOnly
@return GetGroupsTemplateResponse
*/
func (a *GroupsApiService) RetrieveTemplate7(ctx _context.Context, opts *RetrieveTemplate7Opts) (models.GetGroupsTemplateResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.GetGroupsTemplateResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/groups/template"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.OfficeId.IsSet() {
		queryParams.Add("officeId", parameterToString(opts.OfficeId.Value(), ""))
	}
	if opts != nil && opts.Center.IsSet() {
		queryParams.Add("center", parameterToString(opts.Center.Value(), ""))
	}
	if opts != nil && opts.CenterId.IsSet() {
		queryParams.Add("centerId", parameterToString(opts.CenterId.Value(), ""))
	}
	if opts != nil && opts.Command.IsSet() {
		queryParams.Add("command", parameterToString(opts.Command.Value(), ""))
	}
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

// RetrieveglimAccountsOpts Optional parameters for the method 'RetrieveglimAccounts'
type RetrieveglimAccountsOpts struct {
	ParentLoanAccountNo optional.String
}

/*
RetrieveglimAccounts Method for RetrieveglimAccounts
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param groupId
 * @param optional nil or *RetrieveglimAccountsOpts - Optional Parameters:
 * @param "ParentLoanAccountNo" (optional.String) -
@return string
*/
func (a *GroupsApiService) RetrieveglimAccounts(ctx _context.Context, groupId int64, opts *RetrieveglimAccountsOpts) (string, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodGet
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  string
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/groups/{groupId}/glimaccounts"
	path = strings.Replace(path, "{"+"groupId"+"}", _neturl.QueryEscape(parameterToString(groupId, "")), -1)

	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	if opts != nil && opts.ParentLoanAccountNo.IsSet() {
		queryParams.Add("parentLoanAccountNo", parameterToString(opts.ParentLoanAccountNo.Value(), ""))
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

/*
UnassignLoanOfficer Unassign a Staff
Allows you to unassign the Staff.  Mandatory Fields: staffId
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param groupId groupId
 * @param postGroupsGroupIdCommandUnassignStaffRequest
@return PostGroupsGroupIdCommandUnassignStaffResponse
*/
func (a *GroupsApiService) UnassignLoanOfficer(ctx _context.Context, groupId int64, postGroupsGroupIdCommandUnassignStaffRequest models.PostGroupsGroupIdCommandUnassignStaffRequest) (models.PostGroupsGroupIdCommandUnassignStaffResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PostGroupsGroupIdCommandUnassignStaffResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/groups/{groupId}/command/unassign_staff"
	path = strings.Replace(path, "{"+"groupId"+"}", _neturl.QueryEscape(parameterToString(groupId, "")), -1)

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
	postBody = &postGroupsGroupIdCommandUnassignStaffRequest
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
Update11 Update a Group
Updates a Group
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param groupId groupId
 * @param putGroupsGroupIdRequest
@return PutGroupsGroupIdResponse
*/
func (a *GroupsApiService) Update11(ctx _context.Context, groupId int64, putGroupsGroupIdRequest models.PutGroupsGroupIdRequest) (models.PutGroupsGroupIdResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPut
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  models.PutGroupsGroupIdResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/groups/{groupId}"
	path = strings.Replace(path, "{"+"groupId"+"}", _neturl.QueryEscape(parameterToString(groupId, "")), -1)

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
	postBody = &putGroupsGroupIdRequest
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
