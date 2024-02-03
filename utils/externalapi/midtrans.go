package externalapi

import (
	"MyEcommerce/app/config"
	"MyEcommerce/features/order"
	"errors"
	"strconv"

	mid "github.com/midtrans/midtrans-go"

	"github.com/midtrans/midtrans-go/coreapi"
)

type MidtransInterface interface {
	NewOrderPayment(data order.OrderCore) (*order.OrderCore, error)
	CancelOrderPayment(order_id string) error
}

type midtrans struct {
	client      coreapi.Client
	environment mid.EnvironmentType
}

func New() MidtransInterface {
	environment := mid.Sandbox
	var client coreapi.Client
	client.New(config.MID_KEY, environment)

	return &midtrans{
		client: client,
	}
}

// NewOrderPayment implements Midtrans.
func (pay *midtrans) NewOrderPayment(data order.OrderCore) (*order.OrderCore, error) {
	req := new(coreapi.ChargeReq)
	req.TransactionDetails = mid.TransactionDetails{
		OrderID:  data.ID,
		GrossAmt: int64(data.GrossAmount),
	}

	if data.PaymentType == "" {
		data.PaymentType = "bank_transfer"
	}

	switch data.Bank {
	case "bca":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mid.BankBca,
		}
	case "bni":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mid.BankBni,
		}
	case "bri":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mid.BankBri,
		}

	default:
		return nil, errors.New("payment not support")

	}

	res, err := pay.client.ChargeTransaction(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != "201" {
		return nil, errors.New(res.StatusMessage)
	}

	// response
	data.VaNumber, _ = strconv.Atoi(res.VaNumbers[0].VANumber)
	data.PaymentType = res.PaymentType
	data.Status = res.TransactionStatus
	data.Payment.StatusCode = res.StatusCode
	data.Payment.StatusMessage = res.StatusMessage
	data.Payment.TransactionId = res.TransactionID
	data.Payment.Currency = res.Currency
	data.Payment.TransactionTime = res.TransactionTime
	data.Payment.FraudStatus = res.FraudStatus
	data.Payment.ExpiredAt = res.ExpiryTime

	return &data, nil
}

func (pay *midtrans) CancelOrderPayment(orderId string) error {
	res, _ := pay.client.CancelTransaction(orderId)
	if res.StatusCode != "200" && res.StatusCode != "412" {
		return errors.New(res.StatusMessage)
	}

	return nil
}
