package refunds

import (
	"encoding/json"
	"errors"
	restclient "github.com/mercadolibre/gateway_batch_utilities/internal/resclient"
)

type Refund struct {
	Id int `json:"id"`
	Status string `json:"status"`
	StatusDetail string `json:"status_detail"`
	StatusDetailG2 string `json:"status_detail_g2"`
	RetryNumber int `json:"retry_number"`
}

type RefundServiceI interface {
	GetRefund(txId string, refundId string) (Refund, error)
	UpdateRefund(txId string, refundId string, refund Refund) (string, error)
}

type RefundService struct {
	readUrl string
	writeUrl string
}

func NewRefundService(readUrl string, writeUrl string) RefundService {
	return RefundService{readUrl: readUrl, writeUrl: writeUrl}
}

func (rs RefundService) getUpdateUrl(txId string, refundId string) string{
	return rs.writeUrl + "/gateway/transactions/" + txId + "/refund/" + refundId
}

func (rs RefundService) getReadUrl(txId string, refundId string) string{
	return rs.readUrl + "/gateway/transactions/" + txId + "/refund/" + refundId
}

func (rs RefundService) GetRefund(txId string, refundId string) (*Refund, error){

	url := rs.getReadUrl(txId, refundId)

	resp, err := restclient.DoGet(url)

	if err != nil {
		return nil, errors.New("Error obteniendo el refund: " + err.Error())
	}

	refund := &Refund{}

	err = json.Unmarshal([]byte(resp), refund)

	if err != nil {
		return nil, errors.New("Error formateando el body: " + err.Error())
	}

	return refund, nil
}

func (rs RefundService) UpdateRefund(txId string, refundId string, refund Refund) (string, error) {

	url := rs.getUpdateUrl(txId, refundId)

	body, err := json.Marshal(refund)
	if err != nil {
		return "", errors.New("Error formateando el body: " + err.Error())
	}

	resp, err := restclient.DoPut(url, body)

	if err != nil {
		return "", errors.New("Error consumiendo el endpoint: " + err.Error())
	}
	return resp, nil

}