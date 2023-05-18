package requests

type InvoiceDocumentItem struct {
	InvoiceDocument        int   `json:"InvoiceDocument"`
	InvoiceDocumentItem    int   `json:"InvoiceDocumentItem"`
	IsUpdated              *bool `json:"IsUpdated"`
	ItemIsCleared          *bool `json:"ItemIsCleared"`
	ItemBillingIsConfirmed *bool `json:"ItemBillingIsConfirmed"`
}
