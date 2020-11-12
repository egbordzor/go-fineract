package models

// CommandWrapper struct for CommandWrapper
type CommandWrapper struct {
	GroupId                     int64  `json:"groupId,omitempty"`
	ClientId                    int64  `json:"clientId,omitempty"`
	LoanId                      int64  `json:"loanId,omitempty"`
	SavingsId                   int64  `json:"savingsId,omitempty"`
	EntityName                  string `json:"entityName,omitempty"`
	TaskPermissionName          string `json:"taskPermissionName,omitempty"`
	EntityId                    int64  `json:"entityId,omitempty"`
	SubentityId                 int64  `json:"subentityId,omitempty"`
	Href                        string `json:"href,omitempty"`
	Json                        string `json:"json,omitempty"`
	TransactionId               string `json:"transactionId,omitempty"`
	ProductId                   int64  `json:"productId,omitempty"`
	CreditBureauId              int64  `json:"creditBureauId,omitempty"`
	OrganisationCreditBureauId  int64  `json:"organisationCreditBureauId,omitempty"`
	UpdateOperation             bool   `json:"updateOperation,omitempty"`
	Delete                      bool   `json:"delete,omitempty"`
	DeleteOperation             bool   `json:"deleteOperation,omitempty"`
	SurveyResource              bool   `json:"surveyResource,omitempty"`
	RegisterSurvey              bool   `json:"registerSurvey,omitempty"`
	FullFilSurvey               bool   `json:"fullFilSurvey,omitempty"`
	WorkingDaysResource         bool   `json:"workingDaysResource,omitempty"`
	PasswordPreferencesResource bool   `json:"passwordPreferencesResource,omitempty"`
	PermissionResource          bool   `json:"permissionResource,omitempty"`
	UserResource                bool   `json:"userResource,omitempty"`
	CurrencyResource            bool   `json:"currencyResource,omitempty"`
	LoanDisburseDetailResource  bool   `json:"loanDisburseDetailResource,omitempty"`
	UpdateDisbursementDate      bool   `json:"updateDisbursementDate,omitempty"`
	Create                      bool   `json:"create,omitempty"`
	CreateDatatable             bool   `json:"createDatatable,omitempty"`
	DeleteDatatable             bool   `json:"deleteDatatable,omitempty"`
	UpdateDatatable             bool   `json:"updateDatatable,omitempty"`
	DatatableResource           bool   `json:"datatableResource,omitempty"`
	DeleteOneToOne              bool   `json:"deleteOneToOne,omitempty"`
	DeleteMultiple              bool   `json:"deleteMultiple,omitempty"`
	UpdateOneToOne              bool   `json:"updateOneToOne,omitempty"`
	UpdateMultiple              bool   `json:"updateMultiple,omitempty"`
	RegisterDatatable           bool   `json:"registerDatatable,omitempty"`
	NoteResource                bool   `json:"noteResource,omitempty"`
	CacheResource               bool   `json:"cacheResource,omitempty"`
	Update                      bool   `json:"update,omitempty"`
}
