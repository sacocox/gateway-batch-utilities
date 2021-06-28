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

func main(){

	scope := "prod"

	refundReadUrl := utils.GetUrl("refund-read", scope)
	refundWriteUrl := utils.GetUrl("refund-write", scope)

	rs := refunds.NewRefundService(refundReadUrl, refundWriteUrl)

	//dupCheckerUrl := utils.GetUrl("dup-checker", scope)

	//dcs := dup_checker.NewDupCheckerService(dupCheckerUrl)

	lines, err := filereader.ReadFile("cmd/deleteRefundsDupKeys/resources/deleteRefundsDupKeys_" + scope +".txt")
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

		//if deleteRefundKey, err = deleteRefundsDupKeys(txId, refundId, dcs); err == nil {
			if removeStatus, err = removeStatusDetailG2(txId, refundId, rs); err != nil {
				removeStatus = err.Error()
			}
		//}else{
			//deleteRefundKey = err.Error()
		//}

		result += "deleteRefundKey: " + deleteRefundKey + " " + "removeStatus: " + removeStatus

		fmt.Println(result)
	}

}

func removeStatusDetailG2(txId string, refundId string, rs refunds.RefundService) (string, error){
	refund := refunds.Refund{RetryNumber: 0, StatusDetailG2: ""}

	resp, err := rs.UpdateRefund(txId, refundId, refund)

	println(resp)

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


