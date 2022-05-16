package payment

import (
	"order_kafe/config"
	"order_kafe/user"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

type Service interface {
	GetPaymentUrl(transaction Transaction, user user.User) (string, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentUrl(transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = config.InitConfiguration().VT_SERVER_KEY
	midclient.ClientKey = config.InitConfiguration().VT_CLIENT_KEY
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Fullname,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
