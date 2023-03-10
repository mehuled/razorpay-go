/*
 * Razorpay APIs - OpenAPI 3.0
 *
 * Razorpay is an Indian payments solution provider that allows businesses to accept, process and disburse payments with its product suite. 
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Order struct {
	Id string `json:"id,omitempty"`
	Entity string `json:"entity,omitempty"`
	Amount int64 `json:"amount,omitempty"`
	AmountPaid int64 `json:"amount_paid,omitempty"`
	Currency string `json:"currency,omitempty"`
	Receipt string `json:"receipt,omitempty"`
	OfferId string `json:"offer_id,omitempty"`
	Status string `json:"status,omitempty"`
	Attempts int32 `json:"attempts,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
