package clock

import (
	"time"

	null "gopkg.in/guregu/null.v3"
)

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
	CloseAt *Time `json:"close_at"`

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
	Deposit bool `json:"deposit"`

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
	IssuerBillingInfo *BillingInfo `json:"issuer_billing_info"`

	// Receiver of a folio/invoice (Client)
	ContragentBillingInfo *BillingInfo `json:"contragent_billing_info"`

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
	VoidReason         string   `json:"void_reason"`
	PaymentTermsDate   Date     `json:"payment_terms_date"`
}

type BillingInfo struct {
	ID                    int          `json:"id"`
	Name                  string       `json:"name"`
	VAT                   string       `json:"vat"`
	SecondIdentification  string       `json:"second_identification"`
	Country               string       `json:"country"`
	Address               string       `json:"address"`
	City                  string       `json:"city"`
	State                 string       `json:"state"`
	ZipCode               string       `json:"zip_code"`
	PersonName            string       `json:"person_name"`
	AdditionalBillingInfo string       `json:"additional_billing_info"`
	CreatedAt             Time         `json:"created_at"`
	UpdatedAt             Time         `json:"updated_at"`
	CustomFields          CustomFields `json:"custom_fields"`
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

type CustomFields map[string]interface{}

type CustomAttributes interface{}

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

	CustomAttributes CustomAttributes `json:"custom_attributes"`
	GLCode           string           `json:"gl_code"`
	CustomFields     CustomFields     `json:"custom_fields"`

	PaymentDate     Date   `json:"payment_date"`
	LocalCurrency   string `json:"local_currency"`
	LocalValueCents int    `json:"local_value_cents"`
	Reference       string `json:"reference"`
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

type Company struct {
	ID                           int         `json:"id"`
	CreatedAt                    Time        `json:"created_at"`
	UpdatedAt                    Time        `json:"updated_at"`
	Notes                        interface{} `json:"notes"`
	SubscriptionID               int         `json:"subscription_id"`
	Tsv                          string      `json:"tsv"`
	Links                        interface{} `json:"links"`
	AcceptChargeTransfers        bool        `json:"accept_charge_transfers"`
	CommissionRate               interface{} `json:"commission_rate"`
	ContactInfoID                int         `json:"contact_info_id"`
	PaymentTermsDays             interface{} `json:"payment_terms_days"`
	CreditLimitCents             interface{} `json:"credit_limit_cents"`
	Currency                     string      `json:"currency"`
	ArchivedAt                   interface{} `json:"archived_at"`
	Name                         string      `json:"name"`
	VatNumber                    string      `json:"vat_number"`
	SecondIdentification         string      `json:"second_identification"`
	Country                      string      `json:"country"`
	Address                      string      `json:"address"`
	PersonName                   string      `json:"person_name"`
	AdditionalBillingInfo        interface{} `json:"additional_billing_info"`
	Branch                       interface{} `json:"branch"`
	UpdateControlEventAttributes struct {
	} `json:"update_control_event_attributes"`
	City                     interface{} `json:"city"`
	ZipCode                  interface{} `json:"zip_code"`
	State                    interface{} `json:"state"`
	ReportSegmentID          int         `json:"report_segment_id"`
	IATA                     interface{} `json:"iata"`
	ChannelManagerSearchCode interface{} `json:"channel_manager_search_code"`
	WBEAccessCode            interface{} `json:"wbe_access_code"`
	WBERole                  interface{} `json:"wbe_role"`
	RatesValidityDate        interface{} `json:"rates_validity_date"`
	WBEAccessCodeUniqueness  interface{} `json:"wbe_access_code_uniqueness"`
	FamilyID                 int         `json:"family_id"`
	ContactInfo              struct {
		ID                  int         `json:"id"`
		FirstName           string      `json:"first_name"`
		MiddleName          string      `json:"middle_name"`
		LastName            string      `json:"last_name"`
		Country             string      `json:"country"`
		City                string      `json:"city"`
		Address             string      `json:"address"`
		ZipCode             string      `json:"zip_code"`
		PhoneNumber         string      `json:"phone_number"`
		EMail               string      `json:"e_mail"`
		FaxNumber           string      `json:"fax_number"`
		CreatedAt           Time        `json:"created_at"`
		UpdatedAt           Time        `json:"updated_at"`
		TSV                 string      `json:"tsv"`
		Language            interface{} `json:"language"`
		ProfileUUID         interface{} `json:"profile_uuid"`
		ProfileKey          string      `json:"profile_key"`
		SubscriptionID      int         `json:"subscription_id"`
		Current             interface{} `json:"current"`
		State               string      `json:"state"`
		PersonTitleID       interface{} `json:"person_title_id"`
		MigrationPmsGuestID interface{} `json:"migration_pms_guest_id"`
	} `json:"contact_info"`
	BillingInfo struct {
		Name                  string       `json:"name"`
		Vat                   string       `json:"vat"`
		SecondIdentification  string       `json:"second_identification"`
		PersonName            string       `json:"person_name"`
		Country               string       `json:"country"`
		Address               string       `json:"address"`
		City                  interface{}  `json:"city"`
		State                 interface{}  `json:"state"`
		ZipCode               interface{}  `json:"zip_code"`
		AdditionalBillingInfo interface{}  `json:"additional_billing_info"`
		CustomFields          CustomFields `json:"custom_fields"`
	} `json:"billing_info"`
	ReportSegment struct {
		ID                int         `json:"id"`
		MarketingSource   interface{} `json:"marketing_source"`
		MarketingChannel  interface{} `json:"marketing_channel"`
		MarketingSegment  interface{} `json:"marketing_segment"`
		ReservationStatus interface{} `json:"reservation_status"`
	} `json:"report_segment"`
	CustomFields CustomFields `json:"custom_fields"`
}

type Event struct {
	ID                 int         `json:"id"`
	Name               string      `json:"name"`
	Note               string      `json:"note"`
	CompanyID          int         `json:"company_id"`
	CreatedAt          Time        `json:"created_at"`
	UpdatedAt          Time        `json:"updated_at"`
	Arrival            string      `json:"arrival"`
	Departure          string      `json:"departure"`
	AccountID          int         `json:"account_id"`
	Tsv                string      `json:"tsv"`
	Links              string      `json:"links"`
	BlockCutOffDays    interface{} `json:"block_cut_off_days"`
	BlockExpireDate    interface{} `json:"block_expire_date"`
	Number             string      `json:"number"`
	Color              string      `json:"color"`
	CanceledAt         interface{} `json:"canceled_at"`
	UserCreatedID      int         `json:"user_created_id"`
	UserUpdatedID      int         `json:"user_updated_id"`
	ReportSegmentID    int         `json:"report_segment_id"`
	Status             string      `json:"status"`
	CheckinStatus      string      `json:"checkin_status"`
	GuaranteeStatus    string      `json:"guarantee_status"`
	AssignedToUserID   int         `json:"assigned_to_user_id"`
	Tags               string      `json:"tags"`
	CustomerCostCenter string      `json:"customer_cost_center"`
	ReportSegment      struct {
		ID                int    `json:"id"`
		MarketingSource   string `json:"marketing_source"`
		MarketingChannel  string `json:"marketing_channel"`
		MarketingSegment  string `json:"marketing_segment"`
		ReservationStatus string `json:"reservation_status"`
	} `json:"report_segment"`
}

type Booking struct {
	ID                        int           `json:"id"`
	Arrival                   string        `json:"arrival"`
	Departure                 string        `json:"departure"`
	CreatedAt                 time.Time     `json:"created_at"`
	UpdatedAt                 time.Time     `json:"updated_at"`
	Status                    string        `json:"status"`
	BlockID                   int           `json:"block_id"`
	Adults                    int           `json:"adults"`
	Children                  int           `json:"children"`
	AccountID                 int           `json:"account_id"`
	ArrivalTime               string        `json:"arrival_time"`
	TransferArrival           string        `json:"transfer_arrival"`
	TransferDeparture         string        `json:"transfer_departure"`
	SourceOfBusinessCompanyID int           `json:"source_of_business_company_id"`
	DepartureTime             string        `json:"departure_time"`
	ReferenceNumber           string        `json:"reference_number"`
	ReferenceDate             Date          `json:"reference_date"`
	IsGuaranteed              bool          `json:"is_guaranteed"`
	FirstMealID               int           `json:"first_meal_id"`
	RateID                    int           `json:"rate_id"`
	EatingObjectID            int           `json:"eating_object_id"`
	EventID                   int           `json:"event_id"`
	StatusChangedAt           Time          `json:"status_changed_at"`
	UserCreatedID             int           `json:"user_created_id"`
	UserUpdatedID             int           `json:"user_updated_id"`
	ArrivalRoomTypeID         int           `json:"arrival_room_type_id"`
	ArrivalRoomID             int           `json:"arrival_room_id"`
	RegistrationCardsCount    int           `json:"registration_cards_count"`
	ConfirmationLogsCount     int           `json:"confirmation_logs_count"`
	ChildrenAges              []int         `json:"children_ages"`
	SelfServiceKey            string        `json:"self_service_key"`
	SelfServicePin            string        `json:"self_service_pin"`
	Tsv                       string        `json:"tsv"`
	Links                     string        `json:"links"`
	MarketingSource           string        `json:"marketing_source"`
	MarketingChannel          string        `json:"marketing_channel"`
	GuaranteePolicyID         int           `json:"guarantee_policy_id"`
	AcceptChargeTransfers     bool          `json:"accept_charge_transfers"`
	CommissionValueCents      float64       `json:"commission_value_cents"`
	CommissionCurrency        string        `json:"commission_currency"`
	CommissionPaymentDate     Date          `json:"commission_payment_date"`
	CommissionChecked         bool          `json:"commission_checked"`
	CommissionNotes           string        `json:"commission_notes"`
	BeforeArrivalMailSent     bool          `json:"before_arrival_mail_sent"`
	AfterDepartureMailSent    bool          `json:"after_departure_mail_sent"`
	CityTaxMode               string        `json:"city_tax_mode"`
	Number                    string        `json:"number"`
	CitytaxAdults             interface{}   `json:"citytax_adults"`
	CitytaxChildren           interface{}   `json:"citytax_children"`
	DisableRoomChange         bool          `json:"disable_room_change"`
	ConfirmationBodyHTML      interface{}   `json:"confirmation_body_html"`
	DisabledBookingMailerIds  []interface{} `json:"disabled_booking_mailer_ids"`
	RateChargesToCompanyEvent bool          `json:"rate_charges_to_company_event"`
	RequiredFieldsDone        bool          `json:"required_fields_done"`
	CommonDoorCodes           []interface{} `json:"common_door_codes"`
	SelfPreCheckedIn          interface{}   `json:"self_pre_checked_in"`
	CalendarColor             string        `json:"calendar_color"`
	MigrationPmsGuestID       interface{}   `json:"migration_pms_guest_id"`
	RequireRoomResource       bool          `json:"require_room_resource"`
	EarlyArrival              bool          `json:"early_arrival"`
	LateDeparture             bool          `json:"late_departure"`
	MarketingSegment          string        `json:"marketing_segment"`
	ContactPersonID           int           `json:"contact_person_id"`
	AgentID                   int           `json:"agent_id"`
	ReportSegmentID           int           `json:"report_segment_id"`
	RateChargesTransfer       string        `json:"rate_charges_transfer"`
	CustomerCostCenter        string        `json:"customer_cost_center"`
	NotesDestroyedAt          Time          `json:"notes_destroyed_at"`
	CommissionRecipient       string        `json:"commission_recipient"`
	ManualPricesAsText        string        `json:"manual_prices_as_text"`
	ManualCurrency            string        `json:"manual_currency"`
	Balance                   struct {
		Cents    int    `json:"cents"`
		Currency string `json:"currency"`
	} `json:"balance"`
	TotalBookingValue struct {
		Cents    int    `json:"cents"`
		Currency string `json:"currency"`
	} `json:"total_booking_value"`
	Ota             bool `json:"ota"`
	RateCalculation []struct {
		Date     string `json:"date"`
		Cents    int    `json:"cents"`
		Currency string `json:"currency"`
	} `json:"rate_calculation"`
	IdentityTags     []interface{} `json:"identity_tags"`
	MainBookingGuest struct {
		GuestID int `json:"guest_id"`
	} `json:"main_booking_guest"`
	BookingGuests []struct {
		GuestID int `json:"guest_id"`
	} `json:"booking_guests"`
	BookingRoomChanges      []interface{} `json:"booking_room_changes"`
	ActiveNotes             []interface{} `json:"active_notes"`
	ActiveHousekeepingNotes []interface{} `json:"active_housekeeping_notes"`
	ActiveMealsNotes        []interface{} `json:"active_meals_notes"`
	ActiveClientRequests    []interface{} `json:"active_client_requests"`
	Meals                   []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"meals"`
	CustomFields struct {
	} `json:"custom_fields"`
	GuestState           string `json:"guest_state"`
	GuestLanguage        string `json:"guest_language"`
	GuestCountry         string `json:"guest_country"`
	GuestAddress         string `json:"guest_address"`
	GuestLastName        string `json:"guest_last_name"`
	GuestFirstName       string `json:"guest_first_name"`
	GuestCity            string `json:"guest_city"`
	GuestZipCode         string `json:"guest_zip_code"`
	GuestPhoneNumber     string `json:"guest_phone_number"`
	GuestMiddleName      string `json:"guest_middle_name"`
	GuestEMail           string `json:"guest_e_mail"`
	GuestFaxNumber       string `json:"guest_fax_number"`
	GuestPersonTitleID   int    `json:"guest_person_title_id"`
	GuestPersonTitleName string `json:"guest_person_title_name"`
}
type Guest struct {
	ID int `json:"id"`
}

type FolioLedgers struct {
	Items []struct {
		FolioID  int `json:"folio_id"`
		Previous struct {
			Charges struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"charges"`
			ChargesFromDepositFolios struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"charges_from_deposit_folios"`
			Payments struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"payments"`
			Debts struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"debts"`
			DebtsOpenFolios struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"debts_open_folios"`
			DebtsClosedFolios struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"debts_closed_folios"`
			Deposits struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"deposits"`
			DepositsInConsumption struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"deposits_in_consumption"`
			DepositsInAdvance struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"deposits_in_advance"`
			DepositsInAdvanceFromDepositFolios struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"deposits_in_advance_from_deposit_folios"`
			DepositsInAdvanceFromNonDepositFolios struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"deposits_in_advance_from_non_deposit_folios"`
		} `json:"previous"`
		Date struct {
			Charges struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"charges"`
			ChargesFromDepositFolios struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"charges_from_deposit_folios"`
			Payments struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"payments"`
			Debts struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"debts"`
			DebtsOpenFolios struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"debts_open_folios"`
			DebtsClosedFolios struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"debts_closed_folios"`
			Deposits struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"deposits"`
			DepositsInConsumption struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"deposits_in_consumption"`
			DepositsInAdvance struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"deposits_in_advance"`
			DepositsInAdvanceFromDepositFolios struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"deposits_in_advance_from_deposit_folios"`
			DepositsInAdvanceFromNonDepositFolios struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"deposits_in_advance_from_non_deposit_folios"`
		} `json:"date"`
		EndOfDate struct {
			Charges struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"charges"`
			ChargesFromDepositFolios struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"charges_from_deposit_folios"`
			Payments struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"payments"`
			Debts struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"debts"`
			DebtsOpenFolios struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"debts_open_folios"`
			DebtsClosedFolios struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"debts_closed_folios"`
			Deposits struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"deposits"`
			DepositsInConsumption struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"deposits_in_consumption"`
			DepositsInAdvance struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"deposits_in_advance"`
			DepositsInAdvanceFromDepositFolios struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"deposits_in_advance_from_deposit_folios"`
			DepositsInAdvanceFromNonDepositFolios struct {
				Cents    int    `json:"cents"`
				Currency string `json:"currency"`
			} `json:"deposits_in_advance_from_non_deposit_folios"`
		} `json:"end_of_date"`
	} `json:"items"`
	PreviousTotals struct {
		Charges []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"charges"`
		ChargesFromDepositFolios []interface{} `json:"charges_from_deposit_folios"`
		Payments                 []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"payments"`
		Debts []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"debts"`
		DebtsOpenFolios []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"debts_open_folios"`
		DebtsClosedFolios []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"debts_closed_folios"`
		Deposits []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"deposits"`
		DepositsInConsumption []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"deposits_in_consumption"`
		DepositsInAdvance []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"deposits_in_advance"`
		DepositsInAdvanceFromDepositFolios []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"deposits_in_advance_from_deposit_folios"`
		DepositsInAdvanceFromNonDepositFolios []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"deposits_in_advance_from_non_deposit_folios"`
	} `json:"previous_totals"`
	DateTotals struct {
		Charges []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"charges"`
		ChargesFromDepositFolios []interface{} `json:"charges_from_deposit_folios"`
		Payments                 []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"payments"`
		Debts []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"debts"`
		DebtsOpenFolios []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"debts_open_folios"`
		DebtsClosedFolios []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"debts_closed_folios"`
		Deposits []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"deposits"`
		DepositsInConsumption []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"deposits_in_consumption"`
		DepositsInAdvance []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"deposits_in_advance"`
		DepositsInAdvanceFromDepositFolios []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"deposits_in_advance_from_deposit_folios"`
		DepositsInAdvanceFromNonDepositFolios []interface{} `json:"deposits_in_advance_from_non_deposit_folios"`
	} `json:"date_totals"`
	EndOfDateTotals struct {
		Charges []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"charges"`
		ChargesFromDepositFolios []interface{} `json:"charges_from_deposit_folios"`
		Payments                 []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"payments"`
		Debts []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"debts"`
		DebtsOpenFolios []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"debts_open_folios"`
		DebtsClosedFolios []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"debts_closed_folios"`
		Deposits []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"deposits"`
		DepositsInConsumption []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"deposits_in_consumption"`
		DepositsInAdvance []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"deposits_in_advance"`
		DepositsInAdvanceFromDepositFolios []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"deposits_in_advance_from_deposit_folios"`
		DepositsInAdvanceFromNonDepositFolios []struct {
			Cents    int    `json:"cents"`
			Currency string `json:"currency"`
		} `json:"deposits_in_advance_from_non_deposit_folios"`
	} `json:"end_of_date_totals"`
}

type ChargeTemplates []ChargeTemplate

type ChargeTemplate struct {
	ID                     int         `json:"id"`
	Text                   string      `json:"text"`
	RevenueGroup           string      `json:"revenue_group"`
	VisualGroupText        string      `json:"visual_group_text"`
	PlainPriceCents        int         `json:"plain_price_cents"`
	Currency               string      `json:"currency"`
	TaxRate                float64     `json:"tax_rate"`
	AccountID              int         `json:"account_id"`
	CreatedAt              time.Time   `json:"created_at"`
	UpdatedAt              time.Time   `json:"updated_at"`
	StoreID                int         `json:"store_id"`
	InventoryCode          interface{} `json:"inventory_code"`
	Tsv                    interface{} `json:"tsv"`
	RevenueCategory        string      `json:"revenue_category"`
	TaxCode                interface{} `json:"tax_code"`
	Qty                    float64     `json:"qty"`
	PrintText              string      `json:"print_text"`
	DefaultOrderGroup      interface{} `json:"default_order_group"`
	SortOrder              int         `json:"sort_order"`
	Color                  string      `json:"color"`
	CapacityPoolID         int         `json:"capacity_pool_id"`
	ArchivedAt             interface{} `json:"archived_at"`
	HideInPostingScreen    bool        `json:"hide_in_posting_screen"`
	HousekeepingTemplateID int         `json:"housekeeping_template_id"`
	FloatingPrice          bool        `json:"floating_price"`
	CustomFields           struct {
		AccountingCode               string      `json:"accounting_code"`
		Booking                      interface{} `json:"booking"`
		BrpIncludedGroup             interface{} `json:"brp_included_group"`
		BrpIncludedPrice             interface{} `json:"brp_included_price"`
		BrpUpgradableGroup           interface{} `json:"brp_upgradable_group"`
		ChargeOnlyIncludedFirstNight interface{} `json:"charge_only_included_first_night"`
		IncludedDays                 string      `json:"included_days"`
		MealCode                     string      `json:"meal_code"`
		MealCodeForUpgrade           interface{} `json:"meal_code_for_upgrade"`
		MealGroup                    interface{} `json:"meal_group"`
		MealSeatingGroup             string      `json:"meal_seating_group"`
		MealUpgradePrice             interface{} `json:"meal_upgrade_price"`
		ProfitCentre                 string      `json:"profit_centre"`
	} `json:"custom_fields"`
}

type DocumentTypes []DocumentType

type DocumentType struct {
	ID                          int           `json:"id"`
	TextPositiveDoc             string        `json:"text_positive_doc"`
	TextNegativeDoc             string        `json:"text_negative_doc"`
	TextPositiveCorrDoc         string        `json:"text_positive_corr_doc"`
	TextNegativeCorrDoc         string        `json:"text_negative_corr_doc"`
	PositiveDocumentGeneratorID int           `json:"positive_document_generator_id"`
	NegativeDocumentGeneratorID int           `json:"negative_document_generator_id"`
	CreatedAt                   time.Time     `json:"created_at"`
	UpdatedAt                   time.Time     `json:"updated_at"`
	FolioPrintTemplateID        interface{}   `json:"folio_print_template_id"`
	RequireFolioFiscalization   bool          `json:"require_folio_fiscalization"`
	PrintFooter                 []interface{} `json:"print_footer"`
}
