package model

type Device struct {
	Name string `json:"name"`
	SlaveId int `json:"slave_id"`
	Function int `json:"function"`
	Addr int `json:"addr"`
	Quantity int `json:"quantity"`
	Rule float64 `json:"rule"`
	Tick bool `json:"tick"`
}

type DeviceList struct {
	Devices []Device `json:"devices"`
}

type DevConfig struct {
	Uart string `json:"uart"`
	BaudRate int `json:"baudrate"`
	ClientId string `json:"clientId"`
	PubTopic string `json:"pubTopic"`
	SubTopic string `json:"subTopic"`
	UpdateTopic string `json:"updateTopic"`
	MqttHostUrl string `json:"mqttHostUrl"`
	Username string `json:"username"`
	Passwd string `json:"passwd"`
	Port int `json:"port"`
	Interval int `json:"interval"`
	MaxError int `json:"maxError"`
	Devices []Device `json:"devices"`
}

// type DevConfig struct {
// 	Uart string `json:"uart"`
// 	BaudRate int `json:"baud_rate"`
// 	ClientId string `json:"client_id"`
// 	PubTopic string `json:"pub_topic"`
// 	SubTopic string `json:"sub_topic"`
// 	UpdateTopic string `json:"update_topic"`
// 	MqttHostUrl string `json:"mqtt_host_url"`
// 	Username string `json:"username"`
// 	Passwd string `json:"password"`
// 	Port int `json:"port"`
// 	Inetrval int `json:"interval"`
// 	MaxError int `json:"max_error"`
// 	Devices []Device `json:"devices"`
// }
