package meli

import (
	"net/http"
	"time"
)

// Doer defines the method required for an HTTP client to execute requests.
// An *http.Client satisfies this interface.
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// OauthTokenResponse ...
type OauthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	ReceivedAt   int
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

// Category ...
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ClientError ...
type ClientError struct {
	isTemporary bool
	message     string
	statusCode  int
}

func (r ClientError) Temporary() bool {
	return r.isTemporary
}

func (r ClientError) Error() string {
	return r.message
}

func (r ClientError) StatusCode() int {
	return r.statusCode
}

type CategorySettings struct {
	AdultContent            bool          `json:"adult_content"`
	BuyingAllowed           bool          `json:"buying_allowed"`
	BuyingModes             []string      `json:"buying_modes"`
	CatalogDomain           string        `json:"catalog_domain"`
	CoverageAreas           string        `json:"coverage_areas"`
	Currencies              []string      `json:"currencies"`
	Fragile                 bool          `json:"fragile"`
	ImmediatePayment        string        `json:"immediate_payment"`
	ItemConditions          []string      `json:"item_conditions"`
	ItemsReviewsAllowed     bool          `json:"items_reviews_allowed"`
	ListingAllowed          bool          `json:"listing_allowed"`
	MaxDescriptionLength    int           `json:"max_description_length"`
	MaxPicturesPerItem      int           `json:"max_pictures_per_item"`
	MaxPicturesPerItemVar   int           `json:"max_pictures_per_item_var"`
	MaxSubTitleLength       int           `json:"max_sub_title_length"`
	MaxTitleLength          int           `json:"max_title_length"`
	MaximumPrice            interface{}   `json:"maximum_price"`
	MinimumPrice            int           `json:"minimum_price"`
	MirrorCategory          interface{}   `json:"mirror_category"`
	MirrorMasterCategory    interface{}   `json:"mirror_master_category"`
	MirrorSlaveCategories   []interface{} `json:"mirror_slave_categories"`
	Price                   string        `json:"price"`
	ReservationAllowed      string        `json:"reservation_allowed"`
	Restrictions            []interface{} `json:"restrictions"`
	RoundedAddress          bool          `json:"rounded_address"`
	SellerContact           string        `json:"seller_contact"`
	ShippingModes           []string      `json:"shipping_modes"`
	ShippingOptions         []string      `json:"shipping_options"`
	ShippingProfile         string        `json:"shipping_profile"`
	ShowContactInformation  bool          `json:"show_contact_information"`
	SimpleShipping          string        `json:"simple_shipping"`
	Stock                   string        `json:"stock"`
	SubVertical             interface{}   `json:"sub_vertical"`
	Subscribable            bool          `json:"subscribable"`
	Tags                    []interface{} `json:"tags"`
	Vertical                interface{}   `json:"vertical"`
	VipSubdomain            string        `json:"vip_subdomain"`
	BuyerProtectionPrograms []interface{} `json:"buyer_protection_programs"`
	Status                  string        `json:"status"`
}

// CategoryDetails ...
type CategoryDetails struct {
	ID                       string     `json:"id"`
	Name                     string     `json:"name"`
	Picture                  string     `json:"picture"`
	Permalink                string     `json:"permalink"`
	TotalItemsInThisCategory int        `json:"total_items_in_this_category"`
	PathFromRoot             []Category `json:"path_from_root"`
	ChildrenCategories       []struct {
		ID                       string `json:"id"`
		Name                     string `json:"name"`
		TotalItemsInThisCategory int    `json:"total_items_in_this_category"`
	} `json:"children_categories"`
	AttributeTypes string           `json:"attribute_types"`
	Settings       CategorySettings `json:"settings"`
	MetaCategID    interface{}      `json:"meta_categ_id"`
	Attributable   bool             `json:"attributable"`
	DateCreated    time.Time        `json:"date_created"`
}

// ListingType ...
type ListingType struct {
	SiteID string `json:"site_id"`
	ID     string `json:"id"`
	Name   string `json:"name"`
}

// Site ...
type Site struct {
	DefaultCurrencyID string `json:"default_currency_id"`
	ID                string `json:"id"`
	Name              string `json:"name"`
}

// ListingPrice ...
type ListingPrice struct {
	ListingTypeID    string  `json:"listing_type_id"`
	ListingTypeName  string  `json:"listing_type_name"`
	ListingExposure  string  `json:"listing_exposure"`
	RequiresPicture  bool    `json:"requires_picture"`
	CurrencyID       string  `json:"currency_id"`
	ListingFeeAmount int     `json:"listing_fee_amount"`
	SaleFeeAmount    float64 `json:"sale_fee_amount"`
	FreeRelist       bool    `json:"free_relist"`
	StopTime         string  `json:"stop_time"`
}
