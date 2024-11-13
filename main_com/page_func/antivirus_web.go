package page_func

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	config_main "head/main_com/config"
)

type ResponseData struct {
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

type ResponseData_1 struct {
	Data struct {
		Attributes struct {
			LastAnalysisStats struct {
				Malicious int `json:"malicious"`
			} `json:"last_analysis_stats"`
		} `json:"attributes"`
	} `json:"data"`
}

type Return_data struct {
	DNS string `json:"dns"`
	SSL string `json:"ssl"`
}

func checkDomain(domain string) Return_data {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", config_main.Url_domains_virustotal+domain, nil)

	req.Header.Add("x-apikey", config_main.ApiKey_virustotal)

	resp, _ := client.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Помилка при запиті: %d\n", resp.StatusCode)
	}

	var result ResponseData
	json.NewDecoder(resp.Body).Decode(&result)

	var result_func Return_data

	if result.Data.Attributes.LastDNSRecordsDate != 0 {
		lastDNSDate := time.Unix(result.Data.Attributes.LastDNSRecordsDate, 0).UTC().Format("2006-01-02 15:04:05")
		result_func.DNS = lastDNSDate
	} else {
		result_func.DNS = "null"
	}

	if result.Data.Attributes.LastHTTPSCertificate.CertSignature.SignatureAlgorithm != "" {
		result_func.SSL = "Безпечний (сертифікат HTTPS знайдено)"
	} else {
		result_func.SSL = "Небезпечний (можливо, відсутній HTTPS)"
	}

	return result_func
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

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Помилка при запиті: %d\n", resp.StatusCode)
	}

	var result ResponseData_1
	json.NewDecoder(resp.Body).Decode(&result)

	maliciousCount := result.Data.Attributes.LastAnalysisStats.Malicious
	securityStatus := 1
	if maliciousCount > 0 {
		securityStatus = 0
	}

	return securityStatus
}

func CheckUrlInFile(url_ string) int {
	domain := extractDomain(url_)

	checkDomain(domain)

	checkURL(url_)

	return 0
}
