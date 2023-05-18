package api_processing_data_formatter

import (
	"data-platform-function-invoice-confirmation-rmq-kube/DPFM_API_Caller/requests"
	dpfm_api_input_reader "data-platform-function-invoice-confirmation-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"

	"golang.org/x/xerrors"
)

// 1-1. 請求伝票明細データの取得, 2-1.請求伝票明細データの取得
func (psdc *SDC) ConvertToInvoiceDocumentItemKey(sdc *dpfm_api_input_reader.SDC) *InvoiceDocumentItemKey {
	pm := &requests.InvoiceDocumentItemKey{
		IsUpdated:              false,
		ItemIsCleared:          false,
		ItemBillingIsConfirmed: false,
	}

	data := pm
	res := InvoiceDocumentItemKey{
		InvoiceDocument:        data.InvoiceDocument,
		InvoiceDocumentItem:    data.InvoiceDocumentItem,
		IsUpdated:              data.IsUpdated,
		ItemIsCleared:          data.ItemIsCleared,
		ItemBillingIsConfirmed: data.ItemBillingIsConfirmed,
	}

	return &res
}

func (psdc *SDC) ConvertToInvoiceDocumentItem(rows *sql.Rows) ([]*InvoiceDocumentItem, error) {
	defer rows.Close()
	res := make([]*InvoiceDocumentItem, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.InvoiceDocumentItem{}

		err := rows.Scan(
			&pm.InvoiceDocument,
			&pm.InvoiceDocumentItem,
			&pm.IsUpdated,
			&pm.ItemIsCleared,
			&pm.ItemBillingIsConfirmed,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &InvoiceDocumentItem{
			InvoiceDocument:        data.InvoiceDocument,
			InvoiceDocumentItem:    data.InvoiceDocumentItem,
			IsUpdated:              data.IsUpdated,
			ItemIsCleared:          data.ItemIsCleared,
			ItemBillingIsConfirmed: data.ItemBillingIsConfirmed,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_invoice_document_item_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

// 1-2. 請求伝票ヘッダデータの取得
func (psdc *SDC) ConvertToInvoiceDocumentHeaderKey(sdc *dpfm_api_input_reader.SDC) *InvoiceDocumentHeaderKey {
	pm := &requests.InvoiceDocumentHeaderKey{
		IsUpdated:                false,
		HeaderIsCleared:          false,
		HeaderBillingIsConfirmed: false,
	}

	data := pm
	res := InvoiceDocumentHeaderKey{
		InvoiceDocument:          data.InvoiceDocument,
		IsUpdated:                data.IsUpdated,
		HeaderIsCleared:          data.HeaderIsCleared,
		HeaderBillingIsConfirmed: data.HeaderBillingIsConfirmed,
	}

	return &res
}

func (psdc *SDC) ConvertToInvoiceDocumentHeader(rows *sql.Rows) (*InvoiceDocumentHeader, error) {
	defer rows.Close()
	pm := &requests.InvoiceDocumentHeader{}

	i := 0
	for rows.Next() {
		i++
		err := rows.Scan(
			&pm.InvoiceDocument,
			&pm.IsUpdated,
			&pm.HeaderIsCleared,
			&pm.HeaderBillingIsConfirmed,
		)
		if err != nil {
			return nil, err
		}
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_invoice_document_header_data'テーブルに対象のレコードが存在しません。")
	}

	data := pm
	res := &InvoiceDocumentHeader{
		InvoiceDocument:          data.InvoiceDocument,
		IsUpdated:                data.IsUpdated,
		HeaderIsCleared:          data.HeaderIsCleared,
		HeaderBillingIsConfirmed: data.HeaderBillingIsConfirmed,
	}

	return res, nil
}

// 1-3. HeaderBillingIsConfirmedのセット
func (psdc *SDC) ConvertToHeaderBillingIsConfirmed(invoiceDocument int) *HeaderBillingIsConfirmed {
	pm := &requests.HeaderBillingIsConfirmed{
		HeaderBillingIsConfirmed: getBoolPtr(true),
	}

	pm.InvoiceDocument = invoiceDocument

	data := pm
	res := HeaderBillingIsConfirmed{
		InvoiceDocument:          data.InvoiceDocument,
		HeaderBillingIsConfirmed: data.HeaderBillingIsConfirmed,
	}

	return &res
}

// 1-4. ItemBillingIsConfirmedのセット
func (psdc *SDC) ConvertToItemBillingIsConfirmed(invoiceDocument, invoiceDocumentItem int) *ItemBillingIsConfirmed {
	pm := &requests.ItemBillingIsConfirmed{
		ItemBillingIsConfirmed: getBoolPtr(true),
	}

	pm.InvoiceDocument = invoiceDocument
	pm.InvoiceDocumentItem = invoiceDocumentItem

	data := pm
	res := ItemBillingIsConfirmed{
		InvoiceDocument:        data.InvoiceDocument,
		InvoiceDocumentItem:    data.InvoiceDocumentItem,
		ItemBillingIsConfirmed: data.ItemBillingIsConfirmed,
	}

	return &res
}

func getBoolPtr(b bool) *bool {
	return &b
}
