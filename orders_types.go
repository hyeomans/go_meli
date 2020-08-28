package meli

type OrderPhone struct {
	AreaCode  *string `json:"area_code"`
	Extension string  `json:"extension"`
	Number    string  `json:"number"`
	Verified  bool    `json:"verified"`
}

type AlternativePhone struct {
	AreaCode  string `json:"area_code"`
	Extension string `json:"extension"`
	Number    string `json:"number"`
}

type BillingInfo struct {
	DocNumber string `json:"doc_number"`
	DocType   string `json:"doc_type"`
}

type Seller struct {
	AlternativePhone AlternativePhone `json:"alternative_phone"`
	Email            string           `json:"email"`
	FirstName        string           `json:"first_name"`
	ID               int              `json:"id"`
	LastName         string           `json:"last_name"`
	Nickname         string           `json:"nickname"`
	Phone            OrderPhone       `json:"phone"`
}

type Buyer struct {
	AlternativePhone AlternativePhone `json:"alternative_phone"`
	BillingInfo      BillingInfo      `json:"billing_info"`
	Email            string           `json:"email"`
	FirstName        string           `json:"first_name"`
	ID               int              `json:"id"`
	LastName         string           `json:"last_name"`
	Nickname         string           `json:"nickname"`
	Phone            OrderPhone       `json:"phone"`
}

type Coupon struct {
	Amount int  `json:"amount"`
	ID     *int `json:"id"`
}

type Feedback struct {
	Purchase interface{} `json:"purchase"`
	Sale     interface{} `json:"sale"`
}

type OrderItem struct {
	BaseCurrencyID        interface{} `json:"base_currency_id"`
	BaseExchangeRate      interface{} `json:"base_exchange_rate"`
	CurrencyID            string      `json:"currency_id"`
	DifferentialPricingID interface{} `json:"differential_pricing_id"`
	FullUnitPrice         float64     `json:"full_unit_price"`
	Item                  struct {
		CategoryID          string        `json:"category_id"`
		Condition           string        `json:"condition"`
		ID                  string        `json:"id"`
		SellerCustomField   interface{}   `json:"seller_custom_field"`
		SellerSku           interface{}   `json:"seller_sku"`
		Title               string        `json:"title"`
		VariationAttributes []interface{} `json:"variation_attributes"`
		VariationID         interface{}   `json:"variation_id"`
		Warranty            string        `json:"warranty"`
	} `json:"item"`
	ListingTypeID     string      `json:"listing_type_id"`
	ManufacturingDays interface{} `json:"manufacturing_days"`
	Quantity          int         `json:"quantity"`
	SaleFee           float64     `json:"sale_fee"`
	UnitPrice         float64     `json:"unit_price"`
}

type OrderRequest struct {
	Change interface{} `json:"change"`
	Return interface{} `json:"return"`
}

type Payment struct {
	ActivationURI        interface{} `json:"activation_uri"`
	AtmTransferReference struct {
		CompanyID     interface{} `json:"company_id"`
		TransactionID string      `json:"transaction_id"`
	} `json:"atm_transfer_reference"`
	AuthorizationCode string   `json:"authorization_code"`
	AvailableActions  []string `json:"available_actions"`
	CardID            int      `json:"card_id"`
	Collector         struct {
		ID int `json:"id"`
	} `json:"collector"`
	CouponAmount       int         `json:"coupon_amount"`
	CouponID           interface{} `json:"coupon_id"`
	CurrencyID         string      `json:"currency_id"`
	DateApproved       string      `json:"date_approved"`
	DateCreated        string      `json:"date_created"`
	DateLastModified   string      `json:"date_last_modified"`
	DeferredPeriod     interface{} `json:"deferred_period"`
	ID                 int64       `json:"id"`
	InstallmentAmount  float64     `json:"installment_amount"`
	Installments       int         `json:"installments"`
	IssuerID           string      `json:"issuer_id"`
	MarketplaceFee     float64     `json:"marketplace_fee"`
	OperationType      string      `json:"operation_type"`
	OrderID            int         `json:"order_id"`
	OverpaidAmount     int         `json:"overpaid_amount"`
	PayerID            int         `json:"payer_id"`
	PaymentMethodID    string      `json:"payment_method_id"`
	PaymentType        string      `json:"payment_type"`
	Reason             string      `json:"reason"`
	ShippingCost       int         `json:"shipping_cost"`
	SiteID             string      `json:"site_id"`
	Status             string      `json:"status"`
	StatusCode         interface{} `json:"status_code"`
	StatusDetail       string      `json:"status_detail"`
	TaxesAmount        int         `json:"taxes_amount"`
	TotalPaidAmount    float64     `json:"total_paid_amount"`
	TransactionAmount  float64     `json:"transaction_amount"`
	TransactionOrderID interface{} `json:"transaction_order_id"`
}

type Shipping struct {
	ID int64 `json:"id"`
}

type Taxes struct {
	Amount     interface{} `json:"amount"`
	CurrencyID interface{} `json:"currency_id"`
}

type Order struct {
	ApplicationID           string        `json:"application_id"`
	Buyer                   Buyer         `json:"buyer"`
	BuyingMode              string        `json:"buying_mode"`
	Comments                *string       `json:"comments"`
	Coupon                  Coupon        `json:"coupon"`
	CurrencyID              string        `json:"currency_id"`
	DateClosed              string        `json:"date_closed"`
	DateCreated             string        `json:"date_created"`
	ExpirationDate          string        `json:"expiration_date"`
	Feedback                Feedback      `json:"feedback"`
	Fulfilled               bool          `json:"fulfilled"`
	HiddenForSeller         bool          `json:"hidden_for_seller"`
	ID                      int           `json:"id"`
	InternalTags            []string      `json:"internal_tags"`
	LastUpdated             string        `json:"last_updated"`
	ManufacturingEndingDate interface{}   `json:"manufacturing_ending_date"`
	Mediations              []interface{} `json:"mediations"`
	OrderItems              []OrderItem   `json:"order_items"`
	OrderRequest            OrderRequest  `json:"order_request"`
	PackID                  interface{}   `json:"pack_id"`
	PaidAmount              float64       `json:"paid_amount"`
	Payments                []Payment     `json:"payments"`
	PickupID                interface{}   `json:"pickup_id"`
	Seller                  Seller        `json:"seller"`
	Shipping                Shipping      `json:"shipping"`
	ShippingCost            int           `json:"shipping_cost"`
	Status                  string        `json:"status"`
	StatusDetail            interface{}   `json:"status_detail"`
	Tags                    []string      `json:"tags"`
	Taxes                   Taxes         `json:"taxes"`
	TotalAmount             float64       `json:"total_amount"`
}
