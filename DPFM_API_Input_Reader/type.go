package dpfm_api_input_reader

type SDC struct {
	ConnectionKey    string          `json:"connection_key"`
	Result           bool            `json:"result"`
	RedisKey         string          `json:"redis_key"`
	Filepath         string          `json:"filepath"`
	APIStatusCode    int             `json:"api_status_code"`
	RuntimeSessionID string          `json:"runtime_session_id"`
	BusinessPartner  int             `json:"business_partner"`
	ServiceLabel     string          `json:"service_label"`
	APIType          string          `json:"api_type"`
	InvoiceDocument  InvoiceDocument `json:"InvoiceDocument"`
	APISchema        string          `json:"api_schema"`
	Accepter         []string        `json:"accepter"`
	Deleted          bool            `json:"deleted"`
}
type InvoiceDocument struct {
	InvoiceDocument          int    `json:"InvoiceDocument"`
	HeaderBillingIsConfirmed *bool  `json:"HeaderBillingIsConfirmed"`
	Item                     []Item `json:"Item"`
}
type Item struct {
	InvoiceDocumentItem    int   `json:"InvoiceDocumentItem"`
	ItemBillingIsConfirmed *bool `json:"ItemBillingIsConfirmed"`
}
