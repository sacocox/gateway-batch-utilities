package utils

func getStagingUrls() map[string]string{
	return map[string]string{
		"dup-checker": "https://staging_g2-dup-checker.furyapps.io",
		"refund-read": "https://staging_gateway-apitransactions.furyapps.io",
		"refund-write": "https://staging_gateway-apitransactions.furyapps.io",
	}
}

func getProdUrls() map[string]string{
	return map[string]string{
		"dup-checker": "https://internal-api.mercadopago.com/g2/dup-checker",
		"refund-read": "https://refund-read_gateway-apitransactions.furyapps.io",
		"refund-write": "https://prod_gateway-apitransactions.furyapps.io",
	}
}

func GetUrl(app string, scope string) string{
	var url = ""
	if scope == "staging" {
		url, _ = getStagingUrls()[app]
	} else if scope == "production" {
		url, _ = getProdUrls()[app]
	}
	return url
}