package model

type Wlan struct {
	Status bool   `json:"status"`
	Ssid   string `json:"ssid"`
	Ip     string `json:"ip"`
	Password string `json:"password"`
}

type WlanAp struct {
	Ssid string `json:"ssid"`
	Password string `json:"password"`
}
