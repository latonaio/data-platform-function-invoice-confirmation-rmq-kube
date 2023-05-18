package requests

type ItemBillingIsConfirmed struct {
	InvoiceDocument        int   `json:"InvoiceDocument"`
	InvoiceDocumentItem    int   `json:"InvoiceDocumentItem"`
	ItemBillingIsConfirmed *bool `json:"ItemBillingIsConfirmed"`
}
