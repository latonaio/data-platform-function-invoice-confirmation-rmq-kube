package function

import (
	dpfm_api_output_formatter "data-platform-function-invoice-confirmation-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "data-platform-function-invoice-confirmation-rmq-kube/DPFM_API_Processing_Data_Formatter"
)

func (f *Function) SetValue(
	psdc *dpfm_api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) {
	header := dpfm_api_output_formatter.ConvertToHeader(psdc)
	item := dpfm_api_output_formatter.ConvertToItem(psdc)

	osdc.Message = dpfm_api_output_formatter.Message{
		Header: header,
		Item:   item,
	}
}
