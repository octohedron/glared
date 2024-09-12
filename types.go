package main

import "time"

type IPInfo struct {
	IP string `json:"ip"`
}

type domain struct {
	Domain  string
	Name    string
	Zone    string
	ID      string
	All     bool
	Proxied bool
	Exclude []string
}

type authentication struct {
	Key   string
	Email string
}

type updateDNSResult struct {
	Result struct {
		ID        string `json:"id"`
		ZoneID    string `json:"zone_id"`
		ZoneName  string `json:"zone_name"`
		Name      string `json:"name"`
		Type      string `json:"type"`
		Content   string `json:"content"`
		Proxiable bool   `json:"proxiable"`
		Proxied   bool   `json:"proxied"`
		TTL       int    `json:"ttl"`
		Locked    bool   `json:"locked"`
		Meta      struct {
			AutoAdded           bool   `json:"auto_added"`
			ManagedByApps       bool   `json:"managed_by_apps"`
			ManagedByArgoTunnel bool   `json:"managed_by_argo_tunnel"`
			Source              string `json:"source"`
		} `json:"meta"`
		CreatedOn  time.Time `json:"created_on"`
		ModifiedOn time.Time `json:"modified_on"`
	} `json:"result"`
	Success bool `json:"success"`
	Errors  []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
	Messages []interface{} `json:"messages"`
}
type Result struct {
	ID        string `json:"id"`
	ZoneID    string `json:"zone_id"`
	ZoneName  string `json:"zone_name"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	Proxiable bool   `json:"proxiable"`
	Proxied   bool   `json:"proxied"`
	TTL       int    `json:"ttl"`
	Locked    bool   `json:"locked"`
	Meta      struct {
		AutoAdded           bool   `json:"auto_added"`
		ManagedByApps       bool   `json:"managed_by_apps"`
		ManagedByArgoTunnel bool   `json:"managed_by_argo_tunnel"`
		Source              string `json:"source"`
	} `json:"meta"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
}

type zoneDNSList struct {
	Results []Result `json:"result"`
	Success bool     `json:"success"`
	Errors  []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
	Messages   []interface{} `json:"messages"`
	ResultInfo struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		Count      int `json:"count"`
		TotalCount int `json:"total_count"`
		TotalPages int `json:"total_pages"`
	} `json:"result_info"`
}

type DNSRecord struct {
	Domain     string      `json:"domain"`
	RecordType string      `json:"type"`
	Name       string      `json:"name"`
	Content    string      `json:"content"`
	Ttl        interface{} `json:"ttl"`
	Proxied    bool        `json:"proxied"`
}
