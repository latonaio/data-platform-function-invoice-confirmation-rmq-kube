package function

import (
	dpfm_api_input_reader "data-platform-function-invoice-confirmation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-function-invoice-confirmation-rmq-kube/DPFM_API_Processing_Data_Formatter"
)

func (f *Function) InvoiceDocumentHeader(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) (*dpfm_api_processing_data_formatter.InvoiceDocumentHeader, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToInvoiceDocumentHeaderKey(sdc)

	dataKey.InvoiceDocument = sdc.Header.InvoiceDocument

	args = append(args, dataKey.InvoiceDocument, dataKey.IsUpdated, dataKey.HeaderIsCleared, dataKey.HeaderBillingIsConfirmed)

	rows, err := f.db.Query(
		`SELECT InvoiceDocument, IsUpdated, HeaderIsCleared, HeaderBillingIsConfirmed
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_header_data
		WHERE (InvoiceDocument, IsUpdated, HeaderIsCleared, HeaderBillingIsConfirmed) = (?, ?, ?, ?);`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToInvoiceDocumentHeader(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *Function) HeaderBillingIsConfirmed(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) *dpfm_api_processing_data_formatter.HeaderBillingIsConfirmed {
	header := psdc.InvoiceDocumentHeader
	invoiceDocument := header.InvoiceDocument

	data := psdc.ConvertToHeaderBillingIsConfirmed(invoiceDocument)

	return data
}
