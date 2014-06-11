package braintree

type CreditCard struct {
	CustomerId                string             `schema:"-" json:"-" xml:"customer-id,omitempty"`
	Token                     string             `schema:"-" json:"-" xml:"token,omitempty"`
	Number                    string             `schema:"number,omitempty" json:"number,omitempty" xml:"number,omitempty"`
	ExpirationDate            string             `schema:"expiration_date" json:"expiration_date" xml:"expiration-date,omitempty"`
	ExpirationMonth           string             `schema:"-" json:"-" xml:"expiration-month,omitempty"`
	ExpirationYear            string             `schema:"-" json:"-" xml:"expiration-year,omitempty"`
	CVV                       string             `schema:"cvv" json:"cvv" xml:"cvv,omitempty"`
	VenmoSDKPaymentMethodCode string             `schema:"-" json:"-" xml:"venmo-sdk-payment-method-code,omitempty"`
	VenmoSDK                  bool               `schema:"-" json:"-" xml:"venmo-sdk,omitempty"`
	Options                   *CreditCardOptions `schema:"-" json:"-" xml:"options,omitempty"`
	CreatedAt                 string             `schema:"-" json:"-" xml:"created-at,omitempty"`
	UpdatedAt                 string             `schema:"-" json:"-" xml:"updated-at,omitempty"`
	Bin                       string             `schema:"-" json:"-" xml:"bin,omitempty"`
	CardType                  string             `schema:"-" json:"-" xml:"card-type,omitempty"`
	CardholderName            string             `schema:"-" json:"-" xml:"cardholder-name,omitempty"`
	CustomerLocation          string             `schema:"-" json:"-" xml:"customer-location,omitempty"`
	ImageURL                  string             `schema:"-" json:"-" xml:"image-url,omitempty"`
	Default                   bool               `schema:"-" json:"-" xml:"default,omitempty"`
	Expired                   bool               `schema:"-" json:"-" xml:"expired,omitempty"`
	Last4                     string             `schema:"last_4,omitempty" json:"last_4,omitempty" xml:"last-4,omitempty"`
	Commercial                string             `schema:"-" json:"-" xml:"commercial,omitempty"`
	Debit                     string             `schema:"-" json:"-" xml:"debit,omitempty"`
	DurbinRegulated           string             `schema:"-" json:"-" xml:"durbin-regulated,omitempty"`
	Healthcare                string             `schema:"-" json:"-" xml:"healthcare,omitempty"`
	Payroll                   string             `schema:"-" json:"-" xml:"payroll,omitempty"`
	Prepaid                   string             `schema:"-" json:"-" xml:"prepaid,omitempty"`
	CountryOfIssuance         string             `schema:"-" json:"-" xml:"country-of-issuance,omitempty"`
	IssuingBank               string             `schema:"-" json:"-" xml:"issuing-bank,omitempty"`
	UniqueNumberIdentifier    string             `schema:"-" json:"-" xml:"unique-number-identifier,omitempty"`
	BillingAddress            *Address           `schema:"-" json:"-" xml:"billing-address,omitempty"`
	Subscriptions             *Subscriptions     `schema:"-" json:"-" xml:"subscriptions,omitempty"`
}

type CreditCards struct {
	CreditCard []*CreditCard `schema:"-" json:"-" xml:"credit-card"`
}

type CreditCardOptions struct {
	VerifyCard                    bool   `schema:"-" json:"-" xml:"verify-card,omitempty"`
	VenmoSDKSession               string `schema:"-" json:"-" xml:"venmo-sdk-session,omitempty"`
	MakeDefault                   bool   `schema:"-" json:"-" xml:"make-default,omitempty"`
	FailOnDuplicatePaymentMethod  bool   `schema:"-" json:"-" xml:"fail-on-duplicate-payment-method,omitempty"`
	VerificationMerchantAccountId string `schema:"-" json:"-" xml:"verification-merchant-account-id,omitempty"`
	UpdateExistingToken           string `schema:"-" json:"-" xml:"update-existing-token,omitempty"`
}

// AllSubscriptions returns all subscriptions for this card, or nil if none present.
func (card *CreditCard) AllSubscriptions() []*Subscription {
	if card.Subscriptions != nil {
		subs := card.Subscriptions.Subscription
		if len(subs) > 0 {
			a := make([]*Subscription, 0, len(subs))
			for _, s := range subs {
				a = append(a, s)
			}
			return a
		}
	}
	return nil
}
