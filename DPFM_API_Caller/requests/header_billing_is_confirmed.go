package requests

type HeaderBillingIsConfirmed struct {
	InvoiceDocument          int   `json:"InvoiceDocument"`
	HeaderBillingIsConfirmed *bool `json:"HeaderBillingIsConfirmed"`
}
