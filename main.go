package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

var (
	auth       authentication
	IP_ADDRESS = ""
)

func logError(err error) {
	if err != nil {
		log.Printf("%s: %s \n", "Error", err)
	}
}

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

func getZoneDNSListURL(d domain) string {
	return fmt.Sprintf(
		"https://api.cloudflare.com/client/v4/zones/%s/dns_records?type=A&name=%s",
		d.Zone, d.Name)
}

func getZoneDNSList(d domain) zoneDNSList {
	client := &http.Client{}
	req, err := http.NewRequest("GET", getZoneDNSListURL(d), nil)
	logError(err)
	setHeaders(req, true)
	res, err := client.Do(req)
	logError(err)
	body, err := ioutil.ReadAll(res.Body)
	logError(err)
	zRecords := zoneDNSList{}
	err = json.Unmarshal(body, &zRecords)
	logError(err)
	if len(zRecords.Errors) > 0 {
		log.Println(zRecords.Errors[0].Message)
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
	logError(err)
	setHeaders(req, false)
	res, err := client.Do(req)
	logError(err)
	body, err := ioutil.ReadAll(res.Body)
	logError(err)
	updateResult := updateDNSResult{}
	err = json.Unmarshal(body, &updateResult)
	logError(err)
	if len(updateResult.Errors) > 0 {
		log.Println(updateResult.Errors[0].Message)
		logPanic(errors.New(updateResult.Errors[0].Message))
	}
}

func getIPv4Address() IPInfo {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://ipinfo.io", nil)
	logError(err)
	res, err := client.Do(req)
	logError(err)
	body, err := ioutil.ReadAll(res.Body)
	logError(err)
	ipInfo := IPInfo{}
	err = json.Unmarshal(body, &ipInfo)
	logError(err)
	return ipInfo
}

func updateDomains(ip string, domains []domain) {
	for _, d := range domains {
		records := getZoneDNSList(d)
		d.ID = records.Result[0].ID
		updateDNS(d, ip)
		log.Println("Updated", d.Name)
	}
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

func keepUpdated() {
	for {
		ipA := getIPv4Address()
		if ipA.IP != IP_ADDRESS {
			log.Printf(
				"IPV4 address changed from %s to %s, updating/n",
				IP_ADDRESS, ipA.IP)
			domains := getConfig()
			updateDomains(ipA.IP, domains)
			IP_ADDRESS = ipA.IP
		}
		time.Sleep(10 * time.Minute)
	}
}

func main() {
	keepUpdated()
}
