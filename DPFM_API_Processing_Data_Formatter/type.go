package api_processing_data_formatter

type SDC struct {
	InvoiceDocumentItem      []*InvoiceDocumentItem    `json:"InvoiceDocumentItem"`
	InvoiceDocumentHeader    *InvoiceDocumentHeader    `json:"InvoiceDocumentHeader"`
	HeaderBillingIsConfirmed *HeaderBillingIsConfirmed `json:"HeaderBillingIsConfirmed"`
	ItemBillingIsConfirmed   []*ItemBillingIsConfirmed `json:"ItemBillingIsConfirmed"`
}

// 1-1. 請求伝票明細データの取得, 2-1.請求伝票明細データの取得
type InvoiceDocumentItemKey struct {
	InvoiceDocument        int   `json:"InvoiceDocument"`
	InvoiceDocumentItem    []int `json:"InvoiceDocumentItem"`
	IsUpdated              bool  `json:"IsUpdated"`
	ItemIsCleared          bool  `json:"ItemIsCleared"`
	ItemBillingIsConfirmed bool  `json:"ItemBillingIsConfirmed"`
}

type InvoiceDocumentItem struct {
	InvoiceDocument        int   `json:"InvoiceDocument"`
	InvoiceDocumentItem    int   `json:"InvoiceDocumentItem"`
	IsUpdated              *bool `json:"IsUpdated"`
	ItemIsCleared          *bool `json:"ItemIsCleared"`
	ItemBillingIsConfirmed *bool `json:"ItemBillingIsConfirmed"`
}

// 1-2. 請求伝票ヘッダデータの取得
type InvoiceDocumentHeaderKey struct {
	InvoiceDocument          int  `json:"InvoiceDocument"`
	IsUpdated                bool `json:"IsUpdated"`
	HeaderIsCleared          bool `json:"HeaderIsCleared"`
	HeaderBillingIsConfirmed bool `json:"HeaderBillingIsConfirmed"`
}

type InvoiceDocumentHeader struct {
	InvoiceDocument          int   `json:"InvoiceDocument"`
	IsUpdated                *bool `json:"IsUpdated"`
	HeaderIsCleared          *bool `json:"HeaderIsCleared"`
	HeaderBillingIsConfirmed *bool `json:"HeaderBillingIsConfirmed"`
}

// 1-3. HeaderBillingIsConfirmedのセット
type HeaderBillingIsConfirmed struct {
	InvoiceDocument          int   `json:"InvoiceDocument"`
	HeaderBillingIsConfirmed *bool `json:"HeaderBillingIsConfirmed"`
}

// 1-4. ItemBillingIsConfirmedのセット
type ItemBillingIsConfirmed struct {
	InvoiceDocument        int   `json:"InvoiceDocument"`
	InvoiceDocumentItem    int   `json:"InvoiceDocumentItem"`
	ItemBillingIsConfirmed *bool `json:"ItemBillingIsConfirmed"`
}
