package main

import (
	"bufio"
	"fmt"
	"github.com/mercadolibre/gateway_batch_utilities/internal/filereader"
	restclient "github.com/mercadolibre/gateway_batch_utilities/internal/resclient"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const PROD = "production"
const STAG = "staging"

func main() {

	scope := PROD
	url := "https://prod_gateway-apitransactions.furyapps.io/gateway/transactions/g2/%s"
	expression := "channel_transport_certificate_x509_keypair_error"
	fileName:= strconv.FormatInt(time.Now().Unix(), 10)

	lines, err := filereader.ReadFile(fmt.Sprintf("cmd/getTransactionsAndFilter/resources/%s/transactions.txt", scope))
	if err != nil {
		fmt.Printf("Error reading file -> %s", err.Error())
	}

	for n, line := range lines {
		fmt.Printf("Filtering trx in line %d \n", n)
		resp, err := restclient.DoGet(fmt.Sprintf(url, line))
		if err != nil {
			fmt.Printf("Error getting transaction -> %s", err.Error())
		}

		//do filter
		if strings.Contains(resp, expression) {
			appendData(fmt.Sprintf("cmd/getTransactionsAndFilter/resources/%s/result/%s.txt",scope,fileName), line)
		}
	}
}

func appendData(filePath, data string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	dataWriter := bufio.NewWriter(file)

	_, _ = dataWriter.WriteString(data + "\n")

	dataWriter.Flush()
	file.Close()
}
