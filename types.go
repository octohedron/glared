package main

import "time"

type IPInfo struct {
	IP string `json:"ip"`
}

type domain struct {
	Name    string
	Zone    string
	ID      string
	Proxied bool
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

type zoneDNSList struct {
	Result []struct {
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
	RecordType string      `json:"type"`
	Name       string      `json:"name"`
	Content    string      `json:"content"`
	Ttl        interface{} `json:"ttl"`
	Proxied    bool        `json:"proxied"`
}

type TestType struct {
	ID               uint   `json:"id,omitempty" es:"id"`
	CreatedAt        string `json:"created_at,omitempty" es:"created_at"`
	UpdatedAt        string `json:"updated_at,omitempty" es:"updated_at"`
	DeletedAt        string `json:"deleted_at,omitempty" es:"deleted_at"`
	Investments      string `json:"investments,omitempty" custom:"investments"`
	Startups         string `json:"startups,omitempty" es:"startups"`
	Starred          string `json:"starred,omitempty" es:"starred"`
	Markets          string `json:"markets,omitempty" es:"markets"`
	UserTokens       string `json:"user_tokens,omitempty" es:"user_tokens"`
	MemberOfStartups string `json:"member_of_startups,omitempty" es:"member_of_startups"`
	Avatar           string `json:"avatar,omitempty" custom:"avatar"`
	Public           bool   `json:"public" custom:"public"`
	IsInvestor       bool   `json:"is_investor" custom:"main_title"`
	CurrentPosition  string `json:"current_position,omitempty" custom:"current_position"`
	External         bool   `json:"external_investor" custom:"external_investor"`
	YoutubeVideo     string `json:"youtube_video,omitempty" custom:"youtube_video"`
	YoutubeVideoID   string `json:"youtube_video_id,omitempty" custom:"youtube_video_id"`
	InvestorType     string `json:"investor_type,omitempty" custom:"investor_type"`
	MainTitle        string `json:"main_title,omitempty" custom:"main_title"`
	Achievements     string `json:"achievements,omitempty" custom:"achievements"`
	HashedID         string `json:"_id,omitempty" custom:"_id"`
	Email            string `json:"email,omitempty" custom:"email"`
	BusinessEmail    string `json:"business_email,omitempty" custom:"business_email"`
	URLName          string `json:"url_name,omitempty" custom:"url_name"`
	Password         string `json:"password,omitempty" custom:"password"`
	FirstName        string `json:"first_name,omitempty" custom:"first_name"`
	LastName         string `json:"last_name,omitempty" custom:"last_name"`
	FullName         string `json:"full_name,omitempty" custom:"full_name"`
	Username         string `json:"username,omitempty" custom:"username"`
	HashedName       string `json:"hashed_name,omitempty" custom:"hashed_name"`
	Usernamehashed   string `json:"username_hashed,omitempty" custom:"username_hashed"`
	Company          string `json:"company,omitempty" custom:"company"`
	Skills           string `json:"skills,omitempty" custom:"skills"`
	Country          string `json:"country,omitempty" custom:"country"`
	City             string `json:"city,omitempty" custom:"city"`
	WhatDo           string `json:"what_do,omitempty" custom:"what_do"`
	Saved            bool   `json:"saved,omitempty"`
	References       string `json:"references,omitempty"`
	Locations        string `json:"locations,omitempty"`
	Titles           string `json:"titles,omitempty"`
	Education        string `json:"education,omitempty"`
	Experience       string `json:"experience,omitempty"`
}
