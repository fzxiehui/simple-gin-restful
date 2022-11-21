package model

type Process struct {
	Pid  string `json:"pid"`
	Cpu  float64 `json:"cpu"`
	Mem  float64 `json:"mem"`
	Name string `json:"name"`
}

type ProcessList struct {
	ProcessList []Process `json:"process_list"`
}

type Disk struct {
	Device string `json:"device"`
	Used string `json:"used"`
	Avail string `json:"avail"`
	Use string `json:"use"`
}

type Home struct {
	ProcessList []Process `json:"process_list"`
	Cpu float64 `json:"cpu"`
	Mem float64 `json:"mem"`
	Disk Disk `json:"disk"`
}
