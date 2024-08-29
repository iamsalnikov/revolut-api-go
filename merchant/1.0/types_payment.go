package revolut_merchant

import (
	"encoding/json"
	"fmt"
	"time"
)

type Payment struct {
	// The ID of the payment.
	Id string `json:"id"`
	//Possible values: [pending, authentication_challenge, authentication_verified, authorisation_started, authorisation_passed, authorised, capture_started, captured, refund_validated, refund_started, cancellation_started, declining, completing, cancelling, failing, completed, declined, soft_declined, cancelled, failed]
	//
	//The status of the payment.
	State string `json:"state"`
	//Possible values: [high_risk, cardholder_name_missing, unknown_card, unknown_card, pick_up_card, invalid_card, invalid_card, expired_card, do_not_honour, invalid_email, invalid_email, invalid_amount, restricted_card, restricted_card, expired_card, insufficient_funds, rejected_by_customer, rejected_by_customer, cardholder_name_missing, withdrawal_limit_exceeded, withdrawal_limit_exceeded, pick_up_card, 3ds_challenge_failed_manually, invalid_amount, transaction_not_allowed_for_cardholder, issuer_not_available, invalid_expiry, invalid_cvv, invalid_pin, invalid_phone, invalid_address, invalid_country, invalid_merchant, customer_challenge_failed, customer_challenge_abandoned, customer_name_mismatch, technical_error]
	//
	//The reason for a failed or declined payment.
	//
	//A failed or declined payment can result from multiple reasons. To learn more, check our failure reasons.
	DeclineReason string `json:"decline_reason"`
	// The reason for a failed or declined payment, sent by the financial institution processing the payment.
	BankMessage string `json:"bank_message"`
	// The date and time the payment was created.
	CreatedAt time.Time `json:"created_at"`
	// The date and time the payment was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	//Temporary token of the payment used to fetch the reward offer during checkout and link user registrations to the given offers.
	//
	//The token is only valid for a limited time.
	Token string `json:"token"`
	// The total amount of the order in minor currency units. For example, 7034 represents €70.34.
	Amount int `json:"amount"`
	// ISO 4217 currency code in upper case.
	Currency string `json:"currency"`
	// The amount of the settled payment (minor currency unit). For example, 7034 stands for €70.34.
	SettledAmount int `json:"settled_amount"`
	// The currency of the settled payment. ISO 4217 currency code in upper case.
	SettledCurrency string `json:"settled_currency"`
	// The details of the payment method used to make the payment.
	PaymentMethod *PaymentMethod `json:"payment_method"`
	//Details about the authentication challenge that should be performed to complete the authentication process. For more information about Revolut's 3DS solution, see: 3D Secure overview.
	//
	//Only returned if the payment's state is authentication_challenge.
	AuthenticationChallenge *AuthenticationChallenge `json:"authentication_challenge"`
	// Object containing address details.
	BillingAddress *Address `json:"billing_address"`
	//Possible values: [low, high]
	//
	//The risk level of the card.
	//
	//If the risk level is high, the payment might be declined.
	RiskLevel string `json:"risk_level"`
	// The details of the order fee.
	Fees []Fee `json:"fees"`
	// The ID of the associated order.
	OrderId string `json:"order_id,omitempty"`
}

type AuthenticationChallenge struct {
	ThreeDs            *ThreeDs
	ThreeDsFingerPrint *ThreeDsFingerPrint
}

func (p *AuthenticationChallenge) UnmarshalJSON(bytes []byte) error {
	var temp map[string]interface{}
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return err
	}

	if typeStr, ok := temp["type"].(string); ok {
		switch typeStr {
		case "three_ds":
			p.ThreeDs = new(ThreeDs)
			return json.Unmarshal(bytes, p.ThreeDs)
		case "three_ds_fingerPrint":
			p.ThreeDsFingerPrint = new(ThreeDsFingerPrint)
			return json.Unmarshal(bytes, p.ThreeDsFingerPrint)
		default:
			return fmt.Errorf("unknown type %s", typeStr)
		}
	}

	return fmt.Errorf("missing or invalid 'type' field")
}

func (p AuthenticationChallenge) MarshalJSON() ([]byte, error) {
	if p.ThreeDs != nil {
		return json.Marshal(p.ThreeDs)
	}
	if p.ThreeDsFingerPrint != nil {
		return json.Marshal(p.ThreeDsFingerPrint)
	}
	return json.Marshal(nil)
}

type PaymentMethod struct {
	Card              *Card
	RevolutPayCard    *Card
	RevolutPayAccount *PayAccount
}

func (p *PaymentMethod) UnmarshalJSON(bytes []byte) error {
	var temp map[string]interface{}
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return err
	}

	if typeStr, ok := temp["type"].(string); ok {
		switch typeStr {
		case "apple_pay", "apple_tap_to_pay", "card", "cash", "google_pay":
			p.Card = new(Card)
			return json.Unmarshal(bytes, p.Card)
		case "revolut_pay_card":
			p.RevolutPayCard = new(Card)
			return json.Unmarshal(bytes, p.RevolutPayCard)
		case "revolut_pay_account":
			p.RevolutPayAccount = new(PayAccount)
			return json.Unmarshal(bytes, p.RevolutPayAccount)
		default:
			return fmt.Errorf("unknown type %s", typeStr)
		}
	}

	return fmt.Errorf("missing or invalid 'type' field")
}

func (p PaymentMethod) MarshalJSON() ([]byte, error) {
	if p.Card != nil {
		return json.Marshal(p.Card)
	}
	if p.RevolutPayCard != nil {
		return json.Marshal(p.RevolutPayCard)
	}
	if p.RevolutPayAccount != nil {
		return json.Marshal(p.RevolutPayAccount)
	}
	return json.Marshal(nil)
}

type Fee struct {
	//Possible values: [fx, acquiring]
	//
	//The type of the order fee.
	Type string `json:"type"`
	// The amount of the payment fee (minor currency unit). For example, enter 7034 for €70.34 in the field.
	Amount int `json:"amount"`
	// The currency of the payment fee. ISO 4217 currency code in upper case.
	Currency string `json:"currency"`
}

type RetrievePaymentDetailsResponse struct {
	Payment
}
