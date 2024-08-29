package revolut_merchant

import "time"

type Customer struct {
	//Permanent customer ID used to retrieve, update, and delete a customer.
	Id string `json:"id"`
	//The full name of the customer.
	FullName string `json:"full_name"`
	//The name of the customer's business.
	BusinessName string `json:"business_name"`
	//The phone number of the customer in E.164 format.
	Email string `json:"email"`
	//The date and time the customer was created.
	Phone string `json:"phone"`
	//The data and time the customer was last updated.
	CreatedAt string `json:"created_at"`
	//The email address of the customer.
	UpdatedAt string `json:"updated_at"`
	//The details of the payment method.
	PaymentMethods []CustomerPaymentMethod `json:"payment_methods"`
}

type MethodDetails struct {
	//Possible values: >= 6 characters and <= 6 characters
	//
	//The BIN of the payment card.
	Bin string `json:"bin"`
	//Possible values: >= 4 characters and <= 4 characters
	//
	//The last four digits of the payment card.
	Last4 string `json:"last4"`
	//The expiry month of the payment card.
	ExpiryMonth int `json:"expiry_month"`
	//The expiry year of the payment card.
	ExpiryYear int `json:"expiry_year"`
	//The name of the cardholder.
	CardholderName string `json:"cardholder_name"`
	//The billing address of the payment method.
	BillingAddress Address `json:"billing_address"`
	//Possible values: [VISA, MASTERCARD, MAESTRO]
	//
	//The brand of the payment card.
	Brand string `json:"brand"`
	//Possible values: [DEBIT, CREDIT, PREPAID, DEFERRED_DEBIT, CHARGE]
	//
	//The funding type of the payment card.
	Funding string `json:"funding"`
	//The issuer of the payment card.
	Issuer string `json:"issuer"`
	//Two-letter country code of the country where the payment card was issued.
	IssuerCountry string `json:"issuer_country"`
	//The date and time the payment card was added.
	CreatedAt time.Time `json:"created_at"`
}

type CustomerPaymentMethod struct {
	//The ID of the payment method.
	Id string `json:"id"`
	//Possible values: [CARD, REVOLUT_PAY]
	//
	//The type of the payment method.
	Type string `json:"type"`
	//Possible values: [CUSTOMER, MERCHANT]
	//
	//Indicates in which case this saved payment method can be used for payments.
	//
	//CUSTOMER: This payment method can be used only when the customer is on the checkout page.
	//MERCHANT: This payment method can be used without the customer being on the checkout page, and the merchant can initiate transactions, for example, to take payments for recurring transactions.
	SavedFor string `json:"saved_for"`
	//The details of the payment method.
	MethodDetails *MethodDetails `json:"method_details"`
}

type CreateACustomerPayload struct {
	//The full name of the customer.
	FullName string `json:"full_name,omitempty"`
	//The name of the customer's business.
	BusinessName string `json:"business_name,omitempty"`
	//The email address of the customer.
	Email string `json:"email"`
	//The phone number of the customer in E.164 format.
	Phone string `json:"phone"`
}

type CreateACustomerResponse struct {
	Customer
}

type RetrieveACustomerListResponse struct {
	Customer
}

type RetrieveACustomerResponse struct {
	Customer
}

type UpdateACustomerPayload struct {
	//The full name of the customer.
	FullName string `json:"full_name,omitempty"`
	//The name of the customer's business.
	BusinessName string `json:"business_name,omitempty"`
	//The email address of the customer.
	Email string `json:"email,omitempty"`
	//The phone number of the customer in E.164 format.
	Phone string `json:"phone,omitempty"`
}

type UpdateACustomerResponse struct {
	Customer
}

type RetrieveAllPaymentMethodOfACustomerResponse []CustomerPaymentMethod

type RetrieveACustomerPaymentMethodResponse struct {
	CustomerPaymentMethod
}

type UpdateACustomerPaymentMethodPayload struct {
	//Possible values: [CUSTOMER]
	//
	//Update the value of saved_for from MERCHANT to CUSTOMER.
	//
	//This indicates that the updated payment method can't be used for merchant initiated transactions (MIT) any more.
	SavedFor string `json:"saved_for"`
}

type UpdateACustomerPaymentMethodResponse struct {
	CustomerPaymentMethod
}
