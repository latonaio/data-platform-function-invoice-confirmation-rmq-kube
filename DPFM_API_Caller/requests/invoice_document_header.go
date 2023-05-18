package requests

type InvoiceDocumentHeader struct {
	InvoiceDocument          int   `json:"InvoiceDocument"`
	IsUpdated                *bool `json:"IsUpdated"`
	HeaderIsCleared          *bool `json:"HeaderIsCleared"`
	HeaderBillingIsConfirmed *bool `json:"HeaderBillingIsConfirmed"`
}
