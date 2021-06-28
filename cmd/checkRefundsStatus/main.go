package main

import (
	"fmt"
	"github.com/mercadolibre/gateway_batch_utilities/internal/filereader"
	"github.com/mercadolibre/gateway_batch_utilities/internal/refunds"
	"github.com/mercadolibre/gateway_batch_utilities/internal/utils"
	"strings"
)

func main()  {

	scope := "prod"

	lines, err := filereader.ReadFile("cmd/checkRefundsStatus/resources/checkRefunds_" + scope +".txt")
	if err != nil {
		return
	}

	refundReadUrl := utils.GetUrl("refund-read", scope)
	refundWriteUrl := utils.GetUrl("refund-write", scope)

	rs := refunds.NewRefundService(refundReadUrl, refundWriteUrl)

	for _, line := range lines {
		result := ""

		data := strings.Split(line, ",")

		txId := data[0]

		result += "TxId: " + txId + " "

		refundId := data[1]

		result += "RefundId: " + refundId + " "

		refund, err := rs.GetRefund(txId, refundId)

		if err != nil {
			result += "refund not found: " + err.Error()
		}else{
			result += "status: " + refund.Status
		}

		fmt.Println(result)
	}

}
