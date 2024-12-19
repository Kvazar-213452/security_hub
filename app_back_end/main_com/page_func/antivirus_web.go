package page_func

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	config_main "head/main_com/config"
)

type return_func_data struct {
	DNS string `json:"dns"`
	SSL int    `json:"ssl"`
	URL int    `json:"url"`
}

type responseData struct {
	Data struct {
		Attributes struct {
			LastDNSRecordsDate   int64 `json:"last_dns_records_date"`
			LastHTTPSCertificate struct {
				CertSignature struct {
					SignatureAlgorithm string `json:"signature_algorithm"`
				} `json:"cert_signature"`
			} `json:"last_https_certificate"`
		} `json:"attributes"`
	} `json:"data"`
}

type responseData_1 struct {
	Data struct {
		Attributes struct {
			LastAnalysisStats struct {
				Malicious int `json:"malicious"`
			} `json:"last_analysis_stats"`
		} `json:"attributes"`
	} `json:"data"`
}

func checkDomain(domain string) (string, int) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", config_main.Url_domains_virustotal+domain, nil)

	req.Header.Add("x-apikey", config_main.ApiKey_virustotal)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	var result responseData
	json.NewDecoder(resp.Body).Decode(&result)

	data_1 := ""
	data_2 := 0

	if result.Data.Attributes.LastDNSRecordsDate != 0 {
		lastDNSDate := time.Unix(result.Data.Attributes.LastDNSRecordsDate, 0).UTC().Format("2006-01-02 15:04:05")
		data_1 = lastDNSDate
	} else {
		data_1 = "null"
	}

	if result.Data.Attributes.LastHTTPSCertificate.CertSignature.SignatureAlgorithm != "" {
		data_2 = 1
	} else {
		data_2 = 0
	}

	return data_1, data_2
}

func extractDomain(url string) string {
	if strings.HasPrefix(url, "http://") {
		url = url[len("http://"):]
	} else if strings.HasPrefix(url, "https://") {
		url = url[len("https://"):]
	}

	if idx := strings.Index(url, "/"); idx != -1 {
		url = url[:idx]
	}

	return url
}

func checkURL(url string) int {
	encodedURL := base64.URLEncoding.EncodeToString([]byte(url))
	encodedURL = strings.TrimRight(encodedURL, "=")

	client := &http.Client{}
	req, _ := http.NewRequest("GET", config_main.ApiURL_virustotal+encodedURL, nil)

	req.Header.Add("x-apikey", config_main.ApiKey_virustotal)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	var result responseData_1
	json.NewDecoder(resp.Body).Decode(&result)

	maliciousCount := result.Data.Attributes.LastAnalysisStats.Malicious
	securityStatus := 1
	if maliciousCount > 0 {
		securityStatus = 0
	}

	return securityStatus
}

func CheckUrlInFile(url_ string) return_func_data {
	domain := extractDomain(url_)
	var return_func return_func_data

	data_1, data_2 := checkDomain(domain)
	data_3 := checkURL(url_)

	return_func.DNS = data_1
	return_func.SSL = data_2
	return_func.URL = data_3

	return return_func
}
