package model

//request
//response
type AccountResponse struct {
	ErrorCode int           `json:"ErrorCode"`
	ErrorMsg  string        `json:"ErrorMsg"`
	Data      []AccountData `json:"Data"`
	Extra     AccountExtra  `json:"Extra"`
}

type AccountExtra struct {
	Version string `json:"version"`
}

type AccountData struct {
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Cardid       string `json:"cardid"`
	Packages     int    `json:"packages"`
	Id           int    `json:"id"`
	Enabled      string `json:"enabled"`
	BindIfname   string `json:"bind_ifname"`
	Comment      string `json:"comment"`
	LastConntime int    `json:"last_conntime"`
	Username     string `json:"username"`
	Passwd       string `json:"passwd"`
	Packname     string `json:"packname"`
	Expires      int    `json:"expires"`
	StartTime    int    `json:"start_time"`
	CreateTime   int    `json:"create_time"`
	Ppptype      string `json:"ppptype"`
	BindVlanid   string `json:"bind_vlanid"`
	Pppname      string `json:"pppname"`
	Share        int    `json:"share"`
	AutoMac      int    `json:"auto_mac"`
	Upload       int    `json:"upload"`
	Download     int    `json:"download"`
	IpType       int    `json:"ip_type"`
	Mac          string `json:"mac"`
	AutoVlanid   int    `json:"auto_vlanid"`
	Address      string `json:"address"`
}
