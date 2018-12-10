package clock

import null "gopkg.in/guregu/null.v3"

type Folio struct {
	// ID of a folio
	ID int `json:"id"`

	// Currency of a folio/invoice
	Currency Currency `json:"currency"`

	// PMS: Folio number; POS: Bill number
	Number string `json:"number"`

	// Full number of a invoice as printed on document.
	InvoiceNumber string `json:"invoice_number"`

	// Data of folio closure / Invoice date. If not present, the folio is still
	// open.
	CloseDate Date `json:"close_date"`

	// DateTime of the folio closure. Use 'close_date' for finansial purposes
	// instead this field.
	CloseAt Time `json:"close_at"`

	// ID of the 'Company' issued to
	InvoiceToCompanyID int `json:"invoice_to_company_id"`

	// Date/Time of the folio/invoice VOID
	VoidedAt *Time `json:"voided_at"`

	// Date/time folio/invoice was closed
	ClosedAt *Time `json:"closed_at"`

	// Tax mode of this folio/invoice. Check the “Tax Methods” appendix
	TaxMode string `json:"tax_mode"`

	// Is the folio/invoice paid or not
	Payed bool `json:"payed"`

	// Normal (Client) folios/invoices are 'true'. Deposit folios/invoices are
	// 'false'
	Deposit bool `json:"payed"`

	// ID of the Document Type. Document Types are defined by clients for
	// customization puproses.
	DocumentTypeID int `json:"document_type_id"`

	// Due days after folio closure.
	PaymentTermsDays int `json:"payment_terms_days"`

	// Total value of a folio/invoice
	Value CurrencyValue `json:"value"`

	// Net (value without taxes) amount of the folio/invoice
	NetValue CurrencyValue `json:"net_value"`

	// Total tax value of a folio / invoice
	TotalVAT CurrencyValue `json:"total_vat"`

	// Unpaid amount of a folio / invoice
	Balance CurrencyValue `json:"balance"`

	// Issuer of a folio/invoice
	IssuesBillingInfo *BillingInfo `json:"issuer_billing_info"`

	// Receiver of a folio/invoice (Client)
	Contragent *BillingInfo `json:"contragent_billing_info:"`

	// PMS: Optional name of a folio; POS: Table of a bill
	Name string `json:"name"`

	// ID of a related object (Payer).
	// Polymorphic relation.
	PayerID null.Int `json:"payer_id"`

	// Type (Class) of a related object (Payer).
	// Polymorphic relation.
	PayerType string `json:"payer_type"`

	// User having closed a folio/invoice
	ClosedByID null.Int `json:"closed_by_id"`

	// if the folio is a correction folio the field is not null and contains the id of the 'master' folio.
	// The filed is empty (null) for 'master' folios.
	ClientFolioID null.Int `json:"client_folio_id"`

	// Text of the invoice description, If a folio/invoice is issued with “one
	// line text” function
	OneLineDescription string `json:"one_line_description"`

	// POS: Percent of the surcharge of a bill
	SurchargeRate float64 `json:"surcharge_rate"`

	UserCreatedID      null.Int `json:"user_created_id"`
	UserUpdatedID      null.Int `json:"user_updated_id"`
	CreatedAt          Time     `json:"created_at"`
	UpdatedAt          Time     `json:"updated_at"`
	Notes              string   `json:"notes"`
	FiscalRequestCode  string   `json:"fiscal_request_code"`
	FiscalResponseCode string   `json:"fiscal_response_code"`
	QRCodeText         string   `json:"qr_code_text"`
}

type BillingInfo struct {
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	VAT                   string `json:"vat"`
	SecondIdentification  string `json:"second_identification"`
	Country               string `json:"country"`
	Address               string `json"address`
	PersonName            string `json:"person_name"`
	AdditionalBillingInfo string `json:"additional_billing_info"`
	CreatedAt             Time   `json:"created_at"`
	UpdatedAt             Time   `json:"updated_at"`
}

type FolioCharges []FolioCharge

type FolioCharge struct {
	// ID of a charge
	ID int `json:"id"`

	// Financial (revenue) date
	RevenueDate *Date `json:"revenue_date"`

	// Text/Description of a charge
	Text string `json:"text"`

	// The date service to be rendered to a client
	ServiceDate *Date `json:"service_date"`

	ArchivedAt Time `json:"archived_at"`
	CreatedAt  Time `json:"created_at"`
	UpdatedAt  Time `json:"updated_at"`

	// Original price of a charge as user posted it
	PriceCents int `json:"price_cents"`

	// Currency of the "price_cents" field
	Currency Currency

	// Revenue group of a charge. See the Revenue Groups appendix
	RevenueGroup string `json:"revenue_group"`

	RevenueCategory string `json:"revenue_category"`

	// Value of the original price converted to the folio currency
	FixedValueCents int `json:"fixed_value_cents"`

	FixedValueCurrency Currency `json:"fixed_value_currency"`

	// Tax percent (VAT, GST, etc)
	TaxRate float64 `json:"tax_rate"`

	// Value before taxation
	NetValueCents int `json:"net_value_cents"`

	// Tax value
	TaxValueCents int `json:"tax_value_cents"`

	// Value with taxes
	GrossValueCents int `json:"gross_value_cents"`

	// Currency of the net_value_cents, tax_value_cents, gross_value_cents. Same
	// as folio currency
	TaxCurrency Currency `json:"tax_currency"`

	InvoiceToID               int    `json:"invoice_to_id"`
	InvoiceToType             string `json:"invoice_to_type"`
	SourceOfBusinessCompanyID int    `json:"source_of_business_company_id"`
	UserCreatedID             int    `json:"user_created_id"`
	UserUpdatedID             int    `json:"user_updated_id"`

	// Quantity posted
	Qty float64 `json:"qty"`

	StoreID                 *int         `json:"store_id"`
	InventoryCode           *string      `json:"inventory_code"`
	TaxMode                 string       `json:"tax_mode"`
	VisualGroupText         string       `json:"visual_group_text"`
	Service                 bool         `json:"service"`
	Version                 int          `json:"version"`
	VoidedAt                Time         `json:"voided_at"`
	UserFixedValueCents     *int         `json:"user_fixed_value_cents"`
	UserFixedValueCurrency  *Currency    `json:"user_fixed_value_currency"`
	SourceID                int          `json:"source_id"`
	SourceType              string       `json:"source_type"`
	FolioID                 int          `json:"folio_id"`
	TaxCode                 string       `json:"tax_code"`
	TransferSourceAccountID *int         `json:"transfer_source_account_id"`
	ValueCent               int          `json:"value_cents"`
	PrintText               string       `json:"print_text"`
	VersionByID             int          `json:"version_by_id"`
	PBalanceValueCents      int          `json:"p_balance_value_cents"`
	VoidReason              string       `json:"void_reason"`
	AccountID               int          `json:"account_id"`
	CapacityPoolID          *int         `json:"capacity_pool_id"`
	CustomFields            CustomFields `json:"custom_fields"`

	Modifiers []struct {
		Text          string   `json:"text"`
		InventoryCode string   `json:"inventory_code"`
		PriceCents    int      `json:"price_cents"`
		Currency      Currency `json:"currency"`
		ChargeID      int      `json:"charge_id"`
		VisualGroup   string   `json:"visual_group"`
		CreatedAt     Time     `json:"created_at"`
		UpdatedAt     Time     `json:"updated_at"`
	} `json:"modifiers"`
}

type CustomFields struct {
}

type FolioCredits []FolioCredit

type FolioCredit struct {
	// ID of a payment
	ID int `json:"id"`

	// Financial (revenue) date
	RevenueDate Date `json:"revenue_date"`

	// Description (optional)
	Text string `json:"text"`

	// Currency of a payment
	Currency Currency `json:"currency"`

	// Value of a payment
	ValueCents int `json:"value_cents"`

	// See the appendig
	PaymentType string `json:"payment_type"`

	PaymentSubType string `json:"payment_sub_type"`
	ArchivedAt     Time   `json:"archived_at"`
	CreatedAt      Time   `json:"created_at"`
	UpdatedAt      Time   `json:"updated_at"`

	// Value of a payment converted to a folio currency
	FixedValueCents int `json:"fixed_value_cents"`

	FixedValueCurrency Currency `json:"fixed_value_currency"`

	UserCreatedID int `json:"user_created_id"`
	UserUpdatedID int `json:"user_updated_id"`

	// Some transaction identification from the external payment system
	ReferenceID string `json:"reference_id"`

	// Some descriptions from the external payment system
	ReferenceText string `json:"reference_text"`

	FolioID                int       `json:"folio_id"`
	UserFixedValueCents    *int      `json:"user_fixed_value_cents"`
	UserFixedValueCurrency *Currency `json:"user_fixed_value_currency"`

	Version                       int      `json:"version"`
	VoidedAt                      *Time    `json:"voided_at"`
	VersionAt                     Time     `json:"version_at"`
	VersionByID                   int      `json:"version_by_id"`
	VersionEffectiveValueCents    int      `json:"version_effective_value_cents"`
	VersionEffectiveValueCurrency Currency `json:"version_effective_value_currency"`
	VoidReason                    string   `json:"void_reason"`
	AccountID                     int      `json:"account_id"`
	Number                        string   `json:"number"`
}

type Users []User

type User struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	LoginID            string `json:"login_id"`
	Email              string `json:"email"`
	AuthorizedAccounts []int  `json:"authorized_accounts"`
	FiscalCode         string `json:"fiscal_code"`
}
