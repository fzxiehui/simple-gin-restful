package model

type System struct {
	TimeZone string `json:"time_zone"`
	TimeZoneList []string `json:"time_zone_list"`
	Uart string `json:"uart"`
	UartList []string `json:"uart_list"`
	BaudRate int `json:"baud_rate"`
	Interval int `json:"interval"`
	MaxError int `json:"max_error"`

	ClientId string `json:"client_id"`
	PubTopic string `json:"pub_topic"`
	SubTopic string `json:"sub_topic"`
	UpdateTopic string `json:"update_topic"`
	MqttHostUrl string `json:"mqtt_host_url"`
	Username string `json:"username"`
	Passwd string `json:"password"`
	Port int `json:"port"`
}

type TimeZone struct {
	TimeZone string `json:"time_zone"`
}


type Uart struct {
	Uart string `json:"uart"`
	BaudRate int `json:"baud_rate"`
	Interval int `json:"interval"`
	MaxError int `json:"max_error"`
}


type Mqtt struct {
	ClientId string `json:"client_id"`
	PubTopic string `json:"pub_topic"`
	SubTopic string `json:"sub_topic"`
	UpdateTopic string `json:"update_topic"`
	MqttHostUrl string `json:"mqtt_host_url"`
	Username string `json:"username"`
	Passwd string `json:"password"`
	Port int `json:"port"`
}

