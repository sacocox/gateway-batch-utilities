package main

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/gateway_batch_utilities/internal/dup_checker"
	"github.com/mercadolibre/gateway_batch_utilities/internal/filereader"
	"github.com/mercadolibre/gateway_batch_utilities/internal/refunds"
	"github.com/mercadolibre/gateway_batch_utilities/internal/utils"
	"strings"
)

func main() {

	scope := "prod"

	refundReadUrl := utils.GetUrl("refund-read", scope)
	refundWriteUrl := utils.GetUrl("refund-write", scope)

	rs := refunds.NewRefundService(refundReadUrl, refundWriteUrl)

	dupCheckerUrl := utils.GetUrl("dup-checker", scope)

	dcs := dup_checker.NewDupCheckerService(dupCheckerUrl)

	lines, err := filereader.ReadFile("cmd/deleteRefundsDupKeys/resources/deleteRefundsDupKeys_" + scope + ".txt")
	if err != nil {
		return
	}

	for _, line := range lines {

		result := ""

		data := strings.Split(line, ",")

		txId := data[0]

		result += "TxId: " + txId + " "

		refundId := data[1]

		result += "RefundId: " + refundId + " "

		deleteRefundKey := ""

		removeStatus := ""

		retryStatus :=""

		if deleteRefundKey, err = deleteRefundsDupKeys(txId, refundId, dcs); err == nil {
			if removeStatus, err = resetRefund(txId, refundId, rs); err == nil {
				if  retryStatus,err = retryRefund(txId,refundId,rs);err != nil{
					retryStatus = err.Error()
				}
			}else {
				removeStatus = err.Error()
			}
		} else {
			deleteRefundKey = err.Error()
		}

		result += "deleteRefundKey: " + deleteRefundKey + " " + "removeStatus: " + removeStatus + "retryStatus: " + retryStatus

		fmt.Println(result)
	}

}

func resetRefund(txId string, refundId string, rs refunds.RefundService) (string, error) {
	refund := refunds.Refund{RetryNumber: 0, StatusDetailG2: ""}

	_, err := rs.UpdateRefund(txId, refundId, refund)

	if err != nil {

		return "", fmt.Errorf("no se pudo actualizar el key: %w", err)
	}

	return "OK", nil
}

func deleteRefundsDupKeys(txId string, refundId string, dcs dup_checker.DupCheckerService) (string, error) {

	refundKeyId := txId + "-refund-" + refundId

	dupCheckerKey, err := dcs.GetKey(refundKeyId)

	if err != nil {
		return "", fmt.Errorf("no se pudo obtener el key: %w", err)
	} else {

		key := dup_checker.Key{}
		err = json.Unmarshal([]byte(dupCheckerKey), &key)

		if err != nil {
			return "", fmt.Errorf("no se pudo formatear el key: %w", err)
		}

		deleteResp, err := dcs.DeleteKey(key)

		if err != nil {

			return "", fmt.Errorf("no se pudo borrar el key: %w", err)
		}

		fmt.Println(deleteResp)

	}

	return "OK", nil
}

func retryRefund(txId string, refundId string, rs refunds.RefundService) (string, error) {

	_, err := rs.RetryRefund(txId, refundId)

	if err != nil {

		return "", fmt.Errorf("no se pudo actualizar el key: %w", err)
	}

	return "OK", nil
}
