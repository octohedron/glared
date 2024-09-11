package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

var (
	auth authentication
)

func logPanic(err error) {
	if err != nil {
		log.Panicf("%s: %s \n", "Error", err)
	}
}

func setHeaders(req *http.Request, json bool) {
	req.Header.Set("X-Auth-Email", auth.Email)
	req.Header.Set("X-Auth-Key", auth.Key)
	if json {
		req.Header.Set("Content-Type", "application/json")
	}
}

func getZoneDNSList(d domain) zoneDNSList {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(
		"https://api.cloudflare.com/client/v4/zones/%s/dns_records?type=A",
		d.Zone), nil)
	logPanic(err)
	setHeaders(req, true)
	res, err := client.Do(req)
	logPanic(err)
	body, err := io.ReadAll(res.Body)
	logPanic(err)
	zRecords := zoneDNSList{}
	err = json.Unmarshal(body, &zRecords)
	logPanic(err)
	if len(zRecords.Errors) > 0 {
		logPanic(errors.New(zRecords.Errors[0].Message))
	}
	return zRecords
}

func updateDNS(d domain, ip string) {
	client := &http.Client{}
	newDNS := DNSRecord{
		RecordType: "A",
		Name:       d.Name,
		Content:    ip,
		Ttl:        1,
		Proxied:    d.Proxied,
	}
	updateDNSURL := fmt.Sprintf(
		"https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s",
		d.Zone, d.ID)
	jsonValue, _ := json.Marshal(newDNS)
	req, err := http.NewRequest("PUT", updateDNSURL, bytes.NewBuffer(jsonValue))
	logPanic(err)
	setHeaders(req, false)
	res, err := client.Do(req)
	logPanic(err)
	body, err := io.ReadAll(res.Body)
	logPanic(err)
	updateResult := updateDNSResult{}
	err = json.Unmarshal(body, &updateResult)
	logPanic(err)
	if len(updateResult.Errors) > 0 {
		logPanic(errors.New(updateResult.Errors[0].Message))
	}
}

func getIPv4Address() IPInfo {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://ipinfo.io", nil)
	logPanic(err)
	res, err := client.Do(req)
	logPanic(err)
	body, err := io.ReadAll(res.Body)
	logPanic(err)
	ipInfo := IPInfo{}
	err = json.Unmarshal(body, &ipInfo)
	logPanic(err)
	return ipInfo
}

func getConfig() []domain {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	logPanic(err)
	var domains []domain
	err = viper.UnmarshalKey("domains", &domains)
	logPanic(err)
	err = viper.UnmarshalKey("auth", &auth)
	logPanic(err)
	return domains
}

func main() {
	domains := getConfig()
	ipInfo := getIPv4Address()
	for _, d := range domains {
		dnsRecords := getZoneDNSList(d)
		for _, r := range dnsRecords.Result {
			if r.ZoneName == d.Domain &&
				(r.Name == d.Name || r.Name == d.Name+"."+d.Domain) {
				if r.Type == "A" && ipInfo.IP != r.Content {
					log.Printf(
						"IPV4 address changed from %s to %s in %s, updating\n",
						r.Content, ipInfo.IP, r.Name)
					updateDNS(d, ipInfo.IP)
					r.Content = ipInfo.IP
				} else {
					log.Printf(
						"IPV4 address hasn't changed, %s = %s in %s",
						r.Content, ipInfo.IP, r.Name)
				}
			}
		}
	}
}
