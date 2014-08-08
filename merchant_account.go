package braintree

type MerchantAccount struct {
	XMLName                 string                         `xml:"merchant-account,omitempty" json:"-" schema:"-"`
	Id                      string                         `xml:"id,omitempty" json:"id" schema:"id"`
	MasterMerchantAccountId string                         `xml:"master-merchant-account-id,omitempty" json:"-" schema:"-"`
	TOSAccepted             bool                           `xml:"tos_accepted,omitempty" json: "-"`
	Individual              *MerchantAccountPerson         `xml:"individual,omitempty" json:"individual" schema:"individual"`
	Business                *MerchantAccountBusiness       `xml:"business,omitempty" json:"business" schema:"business"`
	FundingOptions          *MerchantAccountFundingOptions `xml:"funding,omitempty" json:"funding" schema:"funding"`
	Status                  string                         `xml:"status,omitempty"  json:"status,omitempty" schema:"status,omitempty"`
	Primary                 bool                           `xml:"-" json:"primary,omitempty" schema:"primary,omitempty"`
}

type MerchantAccountPerson struct {
	FirstName   string   `xml:"first-name,omitempty" json:"first_name,omitempty" schema:"first_name,omitempty"`
	LastName    string   `xml:"last-name,omitempty" json:"last_name,omitempty" schema:"last_name,omitempty"`
	Email       string   `xml:"email,omitempty" json:"email,omitempty" schema:"email,omitempty"`
	Phone       string   `xml:"phone,omitempty" json:"phone,omitempty" schema:"phone,omitempty"`
	DateOfBirth string   `xml:"date-of-birth,omitempty" json:"dob,omitempty" schema:"dob,omitempty"`
	SSN         string   `xml:"ssn,omitempty" json:"-" schema:"ssn,omitempty"`
	SSNLast4    string   `xml:"ssn-last-4,omitempty" json:"ssn_last_4,omitempty" schema:"-"`
	Address     *Address `xml:"address,omitempty" json:"address,omitempty" schema:"address,omitempty"`
}

type MerchantAccountBusiness struct {
	LegalName string   `xml:"legal-name,omitempty" json:"legal_name,omitempty" schema:"legal_name,omitempty"`
	DbaName   string   `xml:"dba-name,omitempty" json:"dba_name,omitempty" schema:"dba_name,omitempty"`
	TaxId     string   `xml:"tax-id,omitempty" json:"tax_id,omitempty" schema:"tax_id,omitempty"`
	Address   *Address `xml:"address,omitempty" json:"address,omitempty" schema:"address,omitempty"`
}

type MerchantAccountFundingOptions struct {
	Destination        string `xml:"destination,omitempty" json:"destination,omitempty" schema:"destination,omitempty"`
	Email              string `xml:"email,omitempty" json:"email,omitempty" schema:"email,omitempty"`
	MobilePhone        string `xml:"mobile-phone,omitempty" json:"mobile_phone,omitempty" schema:"mobile_phone,omitempty"`
	AccountNumber      string `xml:"account-number,omitempty" json:"-" schema:"account_number,omitempty"`
	AccountNumberLast4 string `xml:"account-number-last-4,omitempty" json:"account_number_last_4,omitempty" schema:"-"`
	RoutingNumber      string `xml:"routing-number,omitempty" json:"routing_number,omitempty" schema:"routing_number,omitempty"`
}

const (
	FUNDING_DEST_BANK         = "bank"
	FUNDING_DEST_MOBILE_PHONE = "mobile_phone"
	FUNDING_DEST_EMAIL        = "email"
)
