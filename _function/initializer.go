package function

import (
	"context"
	dpfm_api_input_reader "data-platform-function-invoice-confirmation-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-function-invoice-confirmation-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "data-platform-function-invoice-confirmation-rmq-kube/DPFM_API_Processing_Data_Formatter"
	"sync"

	database "github.com/latonaio/golang-mysql-network-connector"
	"golang.org/x/xerrors"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type Function struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewFunction(ctx context.Context, db *database.Mysql, l *logger.Logger) *Function {
	return &Function{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (f *Function) CreateSdc(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error

	if sdc.Header.HeaderBillingIsConfirmed == nil {
		return xerrors.New("HeaderBillingIsConfirmedがnullです。")
	}

	if *sdc.Header.HeaderBillingIsConfirmed {
		err = f.HeaderInvoiceConfirmation(sdc, psdc)
	} else {
		itemBillingIsConfirmed := false
		for _, item := range sdc.Header.Item {
			if item.ItemBillingIsConfirmed == nil {
				continue
			}
			if *item.ItemBillingIsConfirmed {
				itemBillingIsConfirmed = true
			}
		}

		if itemBillingIsConfirmed {
			err = f.ItemInvoiceConfirmation(sdc, psdc)
		} else {
			err = xerrors.New("ItemBillingIsConfirmedがtrueのItemが存在しません。")
		}
	}
	if err != nil {
		return err
	}

	f.l.Info(psdc)

	f.SetValue(psdc, osdc)

	return nil
}

// 1. ヘッダ請求照合
func (f *Function) HeaderInvoiceConfirmation(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-1. 請求伝票明細データの取得
		psdc.InvoiceDocumentItem, e = f.InvoiceDocumentItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-4. ItemBillingIsConfirmedのセット  //1-1
		psdc.ItemBillingIsConfirmed = f.ItemBillingIsConfirmed(sdc, psdc)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-2. 請求伝票ヘッダデータの取得
		psdc.InvoiceDocumentHeader, e = f.InvoiceDocumentHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-3. HeaderBillingIsConfirmedのセット  //1-2
		psdc.HeaderBillingIsConfirmed = f.HeaderBillingIsConfirmed(sdc, psdc)
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	return nil
}

// 2. 明細請求照合 (請求伝票明細の更新)
func (f *Function) ItemInvoiceConfirmation(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-1. 請求伝票明細データの取得
		psdc.InvoiceDocumentItem, e = f.InvoiceDocumentItemUsingItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-4. ItemBillingIsConfirmedのセット  //1-1
		psdc.ItemBillingIsConfirmed = f.ItemBillingIsConfirmed(sdc, psdc)
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	return nil
}
