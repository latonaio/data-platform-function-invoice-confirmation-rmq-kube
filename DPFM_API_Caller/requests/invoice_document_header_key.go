package requests

type InvoiceDocumentHeaderKey struct {
	InvoiceDocument          int  `json:"InvoiceDocument"`
	IsUpdated                bool `json:"IsUpdated"`
	HeaderIsCleared          bool `json:"HeaderIsCleared"`
	HeaderBillingIsConfirmed bool `json:"HeaderBillingIsConfirmed"`
}
