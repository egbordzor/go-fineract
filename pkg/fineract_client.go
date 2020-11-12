package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// APIClient manages communication with the Apache Fineract API v1.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	AccountNumberFormatApi *AccountNumberFormatApiService

	AccountTransfersApi *AccountTransfersApiService

	AccountingClosureApi *AccountingClosureApiService

	AccountingRulesApi *AccountingRulesApiService

	AdhocQueryApiApi *AdhocQueryApiApiService

	AuditsApi *AuditsApiService

	AuthenticationHTTPBasicApi *AuthenticationHTTPBasicApiService

	BatchAPIApi *BatchAPIApiService

	CacheApi *CacheApiService

	CashierJournalsApi *CashierJournalsApiService

	CashiersApi *CashiersApiService

	CentersApi *CentersApiService

	ChargesApi *ChargesApiService

	ClientApi *ClientApiService

	ClientChargesApi *ClientChargesApiService

	ClientIdentifierApi *ClientIdentifierApiService

	ClientTransactionApi *ClientTransactionApiService

	ClientsAddressApi *ClientsAddressApiService

	CodeValuesApi *CodeValuesApiService

	CodesApi *CodesApiService

	CurrencyApi *CurrencyApiService

	DataTablesApi *DataTablesApiService

	DefaultApi *DefaultApiService

	DocumentsApi *DocumentsApiService

	EntityDataTableApi *EntityDataTableApiService

	EntityFieldConfigurationApi *EntityFieldConfigurationApiService

	ExternalServicesApi *ExternalServicesApiService

	FetchAuthenticatedUserDetailsApi *FetchAuthenticatedUserDetailsApiService

	FixedDepositAccountApi *FixedDepositAccountApiService

	FixedDepositProductApi *FixedDepositProductApiService

	FloatingRatesApi *FloatingRatesApiService

	GeneralLedgerAccountApi *GeneralLedgerAccountApiService

	GlobalConfigurationApi *GlobalConfigurationApiService

	GroupsApi *GroupsApiService

	HolidaysApi *HolidaysApiService

	HooksApi *HooksApiService

	InterestRateChartApi *InterestRateChartApiService

	InterestRateSlabAKAInterestBandsApi *InterestRateSlabAKAInterestBandsApiService

	JournalEntriesApi *JournalEntriesApiService

	ListReportMailingJobHistoryApi *ListReportMailingJobHistoryApiService

	LoanChargesApi *LoanChargesApiService

	LoanCollateralApi *LoanCollateralApiService

	LoanProductsApi *LoanProductsApiService

	LoanReschedulingApi *LoanReschedulingApiService

	LoanTransactionsApi *LoanTransactionsApiService

	LoansApi *LoansApiService

	MIFOSXBATCHJOBSApi *MIFOSXBATCHJOBSApiService

	MakerCheckerOr4EyeFunctionalityApi *MakerCheckerOr4EyeFunctionalityApiService

	MappingFinancialActivitiesToAccountsApi *MappingFinancialActivitiesToAccountsApiService

	MixMappingApi *MixMappingApiService

	MixReportApi *MixReportApiService

	MixTaxonomyApi *MixTaxonomyApiService

	NotesApi *NotesApiService

	NotificationApi *NotificationApiService

	OfficesApi *OfficesApiService

	PasswordPreferencesApi *PasswordPreferencesApiService

	PaymentTypeApi *PaymentTypeApiService

	PeriodicAccrualAccountingApi *PeriodicAccrualAccountingApiService

	PermissionsApi *PermissionsApiService

	PocketApi *PocketApiService

	ProvisioningCategoryApi *ProvisioningCategoryApiService

	ProvisioningCriteriaApi *ProvisioningCriteriaApiService

	ProvisioningEntriesApi *ProvisioningEntriesApiService

	RecurringDepositAccountApi *RecurringDepositAccountApiService

	RecurringDepositAccountTransactionsApi *RecurringDepositAccountTransactionsApiService

	RecurringDepositProductApi *RecurringDepositProductApiService

	ReportMailingJobsApi *ReportMailingJobsApiService

	ReportsApi *ReportsApiService

	RolesApi *RolesApiService

	RunReportsApi *RunReportsApiService

	SPMAPILookUpTableApi *SPMAPILookUpTableApiService

	SavingsAccountApi *SavingsAccountApiService

	SavingsChargesApi *SavingsChargesApiService

	SavingsProductApi *SavingsProductApiService

	SchedulerApi *SchedulerApiService

	ScoreCardApi *ScoreCardApiService

	SearchAPIApi *SearchAPIApiService

	SelfAccountTransferApi *SelfAccountTransferApiService

	SelfAuthenticationApi *SelfAuthenticationApiService

	SelfClientApi *SelfClientApiService

	SelfDividendApi *SelfDividendApiService

	SelfLoanProductsApi *SelfLoanProductsApiService

	SelfLoansApi *SelfLoansApiService

	SelfRunReportApi *SelfRunReportApiService

	SelfSavingsAccountApi *SelfSavingsAccountApiService

	SelfScoreCardApi *SelfScoreCardApiService

	SelfServiceRegistrationApi *SelfServiceRegistrationApiService

	SelfShareAccountsApi *SelfShareAccountsApiService

	SelfSpmApi *SelfSpmApiService

	SelfThirdPartyTransferApi *SelfThirdPartyTransferApiService

	SelfUserApi *SelfUserApiService

	SelfUserDetailsApi *SelfUserDetailsApiService

	ShareAccountApi *ShareAccountApiService

	SpmSurveysApi *SpmSurveysApiService

	StaffApi *StaffApiService

	StandingInstructionsApi *StandingInstructionsApiService

	StandingInstructionsHistoryApi *StandingInstructionsHistoryApiService

	TaxComponentsApi *TaxComponentsApiService

	TaxGroupApi *TaxGroupApiService

	TellerCashManagementApi *TellerCashManagementApiService

	UserGeneratedDocumentsApi *UserGeneratedDocumentsApiService

	UsersApi *UsersApiService

	WorkingDaysApi *WorkingDaysApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AccountNumberFormatApi = (*AccountNumberFormatApiService)(&c.common)
	c.AccountTransfersApi = (*AccountTransfersApiService)(&c.common)
	c.AccountingClosureApi = (*AccountingClosureApiService)(&c.common)
	c.AccountingRulesApi = (*AccountingRulesApiService)(&c.common)
	c.AdhocQueryApiApi = (*AdhocQueryApiApiService)(&c.common)
	c.AuditsApi = (*AuditsApiService)(&c.common)
	c.AuthenticationHTTPBasicApi = (*AuthenticationHTTPBasicApiService)(&c.common)
	c.BatchAPIApi = (*BatchAPIApiService)(&c.common)
	c.CacheApi = (*CacheApiService)(&c.common)
	c.CashierJournalsApi = (*CashierJournalsApiService)(&c.common)
	c.CashiersApi = (*CashiersApiService)(&c.common)
	c.CentersApi = (*CentersApiService)(&c.common)
	c.ChargesApi = (*ChargesApiService)(&c.common)
	c.ClientApi = (*ClientApiService)(&c.common)
	c.ClientChargesApi = (*ClientChargesApiService)(&c.common)
	c.ClientIdentifierApi = (*ClientIdentifierApiService)(&c.common)
	c.ClientTransactionApi = (*ClientTransactionApiService)(&c.common)
	c.ClientsAddressApi = (*ClientsAddressApiService)(&c.common)
	c.CodeValuesApi = (*CodeValuesApiService)(&c.common)
	c.CodesApi = (*CodesApiService)(&c.common)
	c.CurrencyApi = (*CurrencyApiService)(&c.common)
	c.DataTablesApi = (*DataTablesApiService)(&c.common)
	c.DefaultApi = (*DefaultApiService)(&c.common)
	c.DocumentsApi = (*DocumentsApiService)(&c.common)
	c.EntityDataTableApi = (*EntityDataTableApiService)(&c.common)
	c.EntityFieldConfigurationApi = (*EntityFieldConfigurationApiService)(&c.common)
	c.ExternalServicesApi = (*ExternalServicesApiService)(&c.common)
	c.FetchAuthenticatedUserDetailsApi = (*FetchAuthenticatedUserDetailsApiService)(&c.common)
	c.FixedDepositAccountApi = (*FixedDepositAccountApiService)(&c.common)
	c.FixedDepositProductApi = (*FixedDepositProductApiService)(&c.common)
	c.FloatingRatesApi = (*FloatingRatesApiService)(&c.common)
	c.GeneralLedgerAccountApi = (*GeneralLedgerAccountApiService)(&c.common)
	c.GlobalConfigurationApi = (*GlobalConfigurationApiService)(&c.common)
	c.GroupsApi = (*GroupsApiService)(&c.common)
	c.HolidaysApi = (*HolidaysApiService)(&c.common)
	c.HooksApi = (*HooksApiService)(&c.common)
	c.InterestRateChartApi = (*InterestRateChartApiService)(&c.common)
	c.InterestRateSlabAKAInterestBandsApi = (*InterestRateSlabAKAInterestBandsApiService)(&c.common)
	c.JournalEntriesApi = (*JournalEntriesApiService)(&c.common)
	c.ListReportMailingJobHistoryApi = (*ListReportMailingJobHistoryApiService)(&c.common)
	c.LoanChargesApi = (*LoanChargesApiService)(&c.common)
	c.LoanCollateralApi = (*LoanCollateralApiService)(&c.common)
	c.LoanProductsApi = (*LoanProductsApiService)(&c.common)
	c.LoanReschedulingApi = (*LoanReschedulingApiService)(&c.common)
	c.LoanTransactionsApi = (*LoanTransactionsApiService)(&c.common)
	c.LoansApi = (*LoansApiService)(&c.common)
	c.MIFOSXBATCHJOBSApi = (*MIFOSXBATCHJOBSApiService)(&c.common)
	c.MakerCheckerOr4EyeFunctionalityApi = (*MakerCheckerOr4EyeFunctionalityApiService)(&c.common)
	c.MappingFinancialActivitiesToAccountsApi = (*MappingFinancialActivitiesToAccountsApiService)(&c.common)
	c.MixMappingApi = (*MixMappingApiService)(&c.common)
	c.MixReportApi = (*MixReportApiService)(&c.common)
	c.MixTaxonomyApi = (*MixTaxonomyApiService)(&c.common)
	c.NotesApi = (*NotesApiService)(&c.common)
	c.NotificationApi = (*NotificationApiService)(&c.common)
	c.OfficesApi = (*OfficesApiService)(&c.common)
	c.PasswordPreferencesApi = (*PasswordPreferencesApiService)(&c.common)
	c.PaymentTypeApi = (*PaymentTypeApiService)(&c.common)
	c.PeriodicAccrualAccountingApi = (*PeriodicAccrualAccountingApiService)(&c.common)
	c.PermissionsApi = (*PermissionsApiService)(&c.common)
	c.PocketApi = (*PocketApiService)(&c.common)
	c.ProvisioningCategoryApi = (*ProvisioningCategoryApiService)(&c.common)
	c.ProvisioningCriteriaApi = (*ProvisioningCriteriaApiService)(&c.common)
	c.ProvisioningEntriesApi = (*ProvisioningEntriesApiService)(&c.common)
	c.RecurringDepositAccountApi = (*RecurringDepositAccountApiService)(&c.common)
	c.RecurringDepositAccountTransactionsApi = (*RecurringDepositAccountTransactionsApiService)(&c.common)
	c.RecurringDepositProductApi = (*RecurringDepositProductApiService)(&c.common)
	c.ReportMailingJobsApi = (*ReportMailingJobsApiService)(&c.common)
	c.ReportsApi = (*ReportsApiService)(&c.common)
	c.RolesApi = (*RolesApiService)(&c.common)
	c.RunReportsApi = (*RunReportsApiService)(&c.common)
	c.SPMAPILookUpTableApi = (*SPMAPILookUpTableApiService)(&c.common)
	c.SavingsAccountApi = (*SavingsAccountApiService)(&c.common)
	c.SavingsChargesApi = (*SavingsChargesApiService)(&c.common)
	c.SavingsProductApi = (*SavingsProductApiService)(&c.common)
	c.SchedulerApi = (*SchedulerApiService)(&c.common)
	c.ScoreCardApi = (*ScoreCardApiService)(&c.common)
	c.SearchAPIApi = (*SearchAPIApiService)(&c.common)
	c.SelfAccountTransferApi = (*SelfAccountTransferApiService)(&c.common)
	c.SelfAuthenticationApi = (*SelfAuthenticationApiService)(&c.common)
	c.SelfClientApi = (*SelfClientApiService)(&c.common)
	c.SelfDividendApi = (*SelfDividendApiService)(&c.common)
	c.SelfLoanProductsApi = (*SelfLoanProductsApiService)(&c.common)
	c.SelfLoansApi = (*SelfLoansApiService)(&c.common)
	c.SelfRunReportApi = (*SelfRunReportApiService)(&c.common)
	c.SelfSavingsAccountApi = (*SelfSavingsAccountApiService)(&c.common)
	c.SelfScoreCardApi = (*SelfScoreCardApiService)(&c.common)
	c.SelfServiceRegistrationApi = (*SelfServiceRegistrationApiService)(&c.common)
	c.SelfShareAccountsApi = (*SelfShareAccountsApiService)(&c.common)
	c.SelfSpmApi = (*SelfSpmApiService)(&c.common)
	c.SelfThirdPartyTransferApi = (*SelfThirdPartyTransferApiService)(&c.common)
	c.SelfUserApi = (*SelfUserApiService)(&c.common)
	c.SelfUserDetailsApi = (*SelfUserDetailsApiService)(&c.common)
	c.ShareAccountApi = (*ShareAccountApiService)(&c.common)
	c.SpmSurveysApi = (*SpmSurveysApiService)(&c.common)
	c.StaffApi = (*StaffApiService)(&c.common)
	c.StandingInstructionsApi = (*StandingInstructionsApiService)(&c.common)
	c.StandingInstructionsHistoryApi = (*StandingInstructionsHistoryApiService)(&c.common)
	c.TaxComponentsApi = (*TaxComponentsApiService)(&c.common)
	c.TaxGroupApi = (*TaxGroupApiService)(&c.common)
	c.TellerCashManagementApi = (*TellerCashManagementApiService)(&c.common)
	c.UserGeneratedDocumentsApi = (*UserGeneratedDocumentsApiService)(&c.common)
	c.UsersApi = (*UsersApiService)(&c.common)
	c.WorkingDaysApi = (*WorkingDaysApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	return resp, err
}

// ChangeBasePath changes base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.cfg.BasePath = path
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFileName string,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile(formFileName, filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}

	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(**os.File); ok {
		*f, err = ioutil.TempFile("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	if xmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if jsonCheck.MatchString(contentType) {
		if err = json.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}
