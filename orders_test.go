package meli_test

import (
	"context"
	meli "github.com/hyeomans/go_meli"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOrdersEndpoint_Get(t *testing.T) {
	ctx := context.Background()
	accessToken := "accesstoken"

	mockServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		t.Logf("Sending fake request to %v", req.URL)
		assert.Equal(t, accessToken, req.URL.Query().Get("access_token"))
		rw.WriteHeader(200)
		rw.Write([]byte(`
		{
  "application_id": "7092",
  "buyer": {
    "alternative_phone": {
      "area_code": "41",
      "extension": "",
      "number": "30576339"
    },
    "billing_info": {
      "doc_number": "02222111110",
      "doc_type": "CPF"
    },
    "email": "vsch@mail.com",
    "first_name": "Hugo",
    "id": 89660613,
    "last_name": "Schberger",
    "nickname": "HUGO",
    "phone": {
      "area_code": "41",
      "extension": "",
      "number": "99911112",
      "verified": false
    }
  },
  "buying_mode": "buy_equals_pay",
  "comments": null,
  "coupon": {
    "amount": 0,
    "id": null
  },
  "currency_id": "BRL",
  "date_closed": "2019-05-22T03:51:07.000-04:00",
  "date_created": "2019-05-22T03:51:05.000-04:00",
  "expiration_date": "2019-06-19T03:51:07.000-04:00",
  "feedback": {
    "purchase": null,
    "sale": null
  },
  "fulfilled": true,
  "hidden_for_seller": false,
  "id": 2032217210,
  "internal_tags": [],
  "last_updated": "2019-05-28T15:16:04.000-04:00",
  "manufacturing_ending_date": null,
  "mediations": [],
  "order_items": [
    {
      "base_currency_id": null,
      "base_exchange_rate": null,
      "currency_id": "BRL",
      "differential_pricing_id": null,
      "full_unit_price": 129.95,
      "item": {
        "category_id": "MLB33383",
        "condition": "new",
        "id": "MLB1054990648",
        "seller_custom_field": null,
        "seller_sku": null,
        "title": "Kit Com 03 Adesivo Spray 3m 75 Cola Silk Sublimação 300g",
        "variation_attributes": [],
        "variation_id": null,
        "warranty": "Garantia de 1 ano fabricante"
      },
      "listing_type_id": "gold_special",
      "manufacturing_days": null,
      "quantity": 1,
      "sale_fee": 14.29,
      "unit_price": 129.95
    }
  ],
  "order_request": {
    "change": null,
    "return": null
  },
  "pack_id": null,
  "paid_amount": 129.95,
  "payments": [
    {
      "activation_uri": null,
      "atm_transfer_reference": {
        "company_id": null,
        "transaction_id": "135292"
      },
      "authorization_code": "008877",
      "available_actions": [
        "refund"
      ],
      "card_id": 203453778,
      "collector": {
        "id": 239432672
      },
      "coupon_amount": 0,
      "coupon_id": null,
      "currency_id": "BRL",
      "date_approved": "2019-05-22T03:51:07.000-04:00",
      "date_created": "2019-05-22T03:51:05.000-04:00",
      "date_last_modified": "2019-05-22T03:51:07.000-04:00",
      "deferred_period": null,
      "id": 4792155710,
      "installment_amount": 129.95,
      "installments": 1,
      "issuer_id": "24",
      "marketplace_fee": 14.290000000000001,
      "operation_type": "regular_payment",
      "order_id": 2032217210,
      "overpaid_amount": 0,
      "payer_id": 89660613,
      "payment_method_id": "master",
      "payment_type": "credit_card",
      "reason": "Kit Com 03 Adesivo Spray 3m 75 Cola Silk Sublimação 300g",
      "shipping_cost": 0,
      "site_id": "MLB",
      "status": "approved",
      "status_code": null,
      "status_detail": "accredited",
      "taxes_amount": 0,
      "total_paid_amount": 129.95,
      "transaction_amount": 129.95,
      "transaction_order_id": null
    }
  ],
  "pickup_id": null,
  "seller": {
    "alternative_phone": {
      "area_code": "",
      "extension": "",
      "number": ""
    },
    "email": "vsch@mail.com",
    "first_name": "Luis",
    "id": 239432672,
    "last_name": "Cheracomo",
    "nickname": "VENDASDKMB",
    "phone": {
      "area_code": null,
      "extension": "",
      "number": "11971427863",
      "verified": false
    }
  },
  "shipping": {
    "id": 27968238880
  },
  "shipping_cost": 0,
  "status": "paid",
  "status_detail": null,
  "tags": [
    "delivered",
    "paid"
  ],
  "taxes": {
    "amount": null,
    "currency_id": null
  },
  "total_amount": 129.95
}
		`))
	}))
	defer mockServer.Close()
	c := meli.Config{
		HTTP:        mockServer.Client(),
		BaseURL:     mockServer.URL,
		CallbackURL: "https://localhost:3000/callback/mercadolibre?something=token",
		ClientID:    "clientID",
		Secret:      "clientSecret",
	}
	client, err := meli.New(c)

	assert.Nil(t, err)

	r, err := client.Orders.Get(ctx, accessToken, 1)

	assert.Nil(t, err)
	assert.NotNil(t, r)
}