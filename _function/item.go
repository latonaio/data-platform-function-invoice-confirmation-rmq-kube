package function

import (
	dpfm_api_input_reader "data-platform-function-invoice-confirmation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-function-invoice-confirmation-rmq-kube/DPFM_API_Processing_Data_Formatter"
	"strings"
)

func (f *Function) InvoiceDocumentItem(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) ([]*dpfm_api_processing_data_formatter.InvoiceDocumentItem, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToInvoiceDocumentItemKey(sdc)

	dataKey.InvoiceDocument = sdc.Header.InvoiceDocument

	args = append(args, dataKey.InvoiceDocument, dataKey.IsUpdated, dataKey.ItemIsCleared, dataKey.ItemBillingIsConfirmed)

	rows, err := f.db.Query(
		`SELECT InvoiceDocument, InvoiceDocumentItem, IsUpdated, ItemIsCleared, ItemBillingIsConfirmed
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_item_data
		WHERE (InvoiceDocument, IsUpdated, ItemIsCleared, ItemBillingIsConfirmed) = (?, ?, ?, ?);`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToInvoiceDocumentItem(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *Function) InvoiceDocumentItemUsingItem(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) ([]*dpfm_api_processing_data_formatter.InvoiceDocumentItem, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToInvoiceDocumentItemKey(sdc)

	dataKey.InvoiceDocument = sdc.Header.InvoiceDocument

	for _, item := range sdc.Header.Item {
		if item.ItemBillingIsConfirmed == nil {
			continue
		}
		if *item.ItemBillingIsConfirmed {
			dataKey.InvoiceDocumentItem = append(dataKey.InvoiceDocumentItem, item.InvoiceDocumentItem)
		}
	}

	args = append(args, dataKey.InvoiceDocument, dataKey.IsUpdated, dataKey.ItemIsCleared, dataKey.ItemBillingIsConfirmed)

	repeat := strings.Repeat("?,", len(dataKey.InvoiceDocumentItem)-1) + "?"
	for _, v := range dataKey.InvoiceDocumentItem {
		args = append(args, v)
	}

	rows, err := f.db.Query(
		`SELECT InvoiceDocument, InvoiceDocumentItem, IsUpdated, ItemIsCleared, ItemBillingIsConfirmed
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_item_data
		WHERE (InvoiceDocument, IsUpdated, ItemIsCleared, ItemBillingIsConfirmed) = (?, ?, ?, ?)
		AND InvoiceDocumentItem IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToInvoiceDocumentItem(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *Function) ItemBillingIsConfirmed(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) []*dpfm_api_processing_data_formatter.ItemBillingIsConfirmed {
	data := make([]*dpfm_api_processing_data_formatter.ItemBillingIsConfirmed, 0)

	for _, item := range psdc.InvoiceDocumentItem {
		invoiceDocument := item.InvoiceDocument
		invoiceDocumentItem := item.InvoiceDocumentItem

		datum := psdc.ConvertToItemBillingIsConfirmed(invoiceDocument, invoiceDocumentItem)
		data = append(data, datum)
	}

	return data
}
