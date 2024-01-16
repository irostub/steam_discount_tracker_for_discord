package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

const URL string = "https://store.steampowered.com/api/featuredcategories"

type Client struct {
	url string
	cc  string
}

type Response struct {
	Specials Specials `json:"specials"`
}

type Specials struct {
	Id    string  `json:"id"`
	Items []Items `json:"items"`
}

type Items struct {
	ID                      int    `json:"id"`
	Type                    int    `json:"type"`
	Name                    string `json:"name"`
	Discounted              bool   `json:"discounted"`
	DiscountPercent         int    `json:"discount_percent"`
	OriginalPrice           int64  `json:"original_price"`
	FinalPrice              int64  `json:"final_price"`
	Currency                string `json:"currency"`
	LargeCapsuleImage       string `json:"large_capsule_image"`
	SmallCapsuleImage       string `json:"small_capsule_image"`
	WindowsAvailable        bool   `json:"windows_available"`
	MacAvailable            bool   `json:"mac_available"`
	LinuxAvailable          bool   `json:"linux_available"`
	StreamingVideoAvailable bool   `json:"streamingvideo_available"`
	DiscountExpiration      int64  `json:"discount_expiration"`
	HeaderImage             string `json:"header_image"`
}

var lock = &sync.Mutex{}
var client *Client

func GetInstance(cc string) *Client {
	if client == nil {
		lock.Lock()
		defer lock.Unlock()
		if client == nil {
			client = &Client{url: URL, cc: cc}
		}
	}
	return client
}

func (client Client) Get() *Response {
	resp, err := http.Get(client.url + "?cc=" + client.cc)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("data %v\n", response)
	return &response
}

func (items Items) IsExpired() bool {
	now := time.Now().Unix()
	expirationTime := items.DiscountExpiration
	standardTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	if items.DiscountExpiration <= 0 || standardTime.After(time.Unix(items.DiscountExpiration, 0)) {
		return true
	}
	return now > expirationTime
}
