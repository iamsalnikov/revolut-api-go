package customer_types

type CreateCustomerPayload struct {
	//The full name of the customer.
	FullName string `json:"full_name,omitempty"`
	//The name of the customer's business.
	BusinessName string `json:"business_name,omitempty"`
	//The email address of the customer.
	Email string `json:"email"`
	//The phone number of the customer in E.164 format.
	Phone string `json:"phone"`
}

type CreateCustomerResponse struct {
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
}
