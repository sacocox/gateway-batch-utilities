package utils




var prodURLs = map[string]string{
	"dup-checker": "https://staging_g2-dup-checker.furyapps.io",
	"refund-read": "https://staging_gateway-apitransactions.furyapps.io",
	"refund-write": "https://staging_gateway-apitransactions.furyapps.io",
}

var stagingURLs = map[string]string{
	"dup-checker": "https://internal-api.mercadopago.com/g2/dup-checker",
	"refund-read": "https://refund-read_gateway-apitransactions.furyapps.io",
	"refund-write": "https://prod_gateway-apitransactions.furyapps.io",
}

func GetProductionURL(appName string){
	url,ok:= prodURLs[appName]; if ok {
		return url
	}
	return ""
}

func GetStagingURL(appName string){
	url,ok:= stagingURLs[appName]; if ok {
		return url
	}
	return ""
}