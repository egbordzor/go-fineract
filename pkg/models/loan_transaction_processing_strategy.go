package models

// LoanTransactionProcessingStrategy struct for LoanTransactionProcessingStrategy
type LoanTransactionProcessingStrategy struct {
	Id                                          int64 `json:"id,omitempty"`
	StandardStrategy                            bool  `json:"standardStrategy,omitempty"`
	HeavensfamilyStrategy                       bool  `json:"heavensfamilyStrategy,omitempty"`
	EarlyPaymentStrategy                        bool  `json:"earlyPaymentStrategy,omitempty"`
	CreocoreStrategy                            bool  `json:"creocoreStrategy,omitempty"`
	IndianRBIStrategy                           bool  `json:"indianRBIStrategy,omitempty"`
	PrincipalInterestPenaltiesFeesOrderStrategy bool  `json:"principalInterestPenaltiesFeesOrderStrategy,omitempty"`
	InterestPrincipalPenaltiesFeesOrderStrategy bool  `json:"interestPrincipalPenaltiesFeesOrderStrategy,omitempty"`
	New                                         bool  `json:"new,omitempty"`
}
