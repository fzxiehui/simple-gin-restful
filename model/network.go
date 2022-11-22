package model

type Network struct {
	Name string `json:"name" binding:"required"`
	Operstate string `json:"operstate"`
	Inet string `json:"inet"`
	// Netmask string `json:"netmask"`
	// Broadcast string `json:"broadcast"`
	Gateway string `json:"gateway"`
	DNS string `json:"dns"`
	DHCP bool `json:"dhcp"`
}

type NetworkList struct {
	NetworkList []Network `json:"network_list"`
}



