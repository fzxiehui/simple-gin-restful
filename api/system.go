package api


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	// "log"
	"strings"
	// "strconv"

	"github.com/fzxiehui/simple-gin-restful/model"
	"github.com/fzxiehui/simple-gin-restful/utils"
)

func GetSystem(c *gin.Context) {

	// check if config.json exists
	var cmd string
	cmd = "ls /root/work/config.json"
	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		cmd = "cd /root/work && eunuch init"
		_, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			c.JSON(http.StatusInternalServerError, 
			gin.H{"message": err.Error()})
			return
		}
	}

	// get time zone
	var system model.System

	cmd = "timedatectl | grep 'Time zone' | awk '{print $3}'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return 
	}
	system.TimeZone = strings.TrimSpace(string(out))

	// get time zone list
	cmd = "timedatectl list-timezones"
	out, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	var timezones = strings.Split(string(out), "\n")

	// TrimSpace
	for _, v := range timezones {
		var tmp = strings.TrimSpace(v)
		if len(tmp) > 2 {
			system.TimeZoneList = append(system.TimeZoneList, tmp)
			// timezones[i] = tmp
		}

	}
	// system.TimeZoneList = timezones
	
	// get uart list
	cmd = "ls /dev/ttyS*"
	out, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	var uarts = strings.TrimSpace(string(out))
	// system.UartList = strings.Split(uarts, " ")
	var tmp_uarts = strings.Split(uarts, "\n")
	for _, v := range tmp_uarts {
		var tmp = strings.TrimSpace(v)
		if len(tmp) > 2 {
			system.UartList = append(system.UartList, tmp)
		}
	}

	var devConfig model.DevConfig

	// get config.json uart
	jsonConfig, err := utils.ReadJSONFile("/root/work/config.json")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// map to struct
	err = utils.MapToStruct(jsonConfig, &devConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	system.Uart = devConfig.Uart
	system.BaudRate = devConfig.BaudRate
	system.Interval = devConfig.Interval
	system.MaxError = devConfig.MaxError

	system.ClientId = devConfig.ClientId
	
	system.PubTopic = devConfig.PubTopic
	system.SubTopic = devConfig.SubTopic
	system.UpdateTopic = devConfig.UpdateTopic
	system.MqttHostUrl = devConfig.MqttHostUrl
	system.Username = devConfig.Username
	system.Passwd = devConfig.Passwd
	system.Port = devConfig.Port

	// var uartList []string
	c.JSON(http.StatusOK, system)
	return
}

func UpdateTimeZone(c *gin.Context) {
	var timezone model.TimeZone
	c.BindJSON(&timezone)
	var cmd string
	cmd = "timedatectl set-timezone " + timezone.TimeZone
	_, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "update time zone success",
	})
	return
}


func UpdateUart(c *gin.Context) {

	// check if config.json exists
	var cmd string
	cmd = "ls /root/work/config.json"
	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		cmd = "cd /root/work && eunuch init"
		_, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			c.JSON(http.StatusInternalServerError, 
			gin.H{"message": err.Error()})
			return
		}
	}

	var uart model.Uart
	c.BindJSON(&uart)
	// Read config.json
	var devConfig model.DevConfig
	jsonConfig, err := utils.ReadJSONFile("/root/work/config.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	// map to struct
	err = utils.MapToStruct(jsonConfig, &devConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	// update uart
	devConfig.Uart = uart.Uart
	devConfig.BaudRate = uart.BaudRate
	devConfig.Interval = uart.Interval
	devConfig.MaxError = uart.MaxError

	// map to json
	mapConfig := utils.StructToMap(devConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// write to config.json
	err = utils.WriteJSONFile("/root/work/config.json", mapConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// restart eunuch 
	cmd = "systemctl restart gateway"
	_, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "restart failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update uart success",
	})
	return
}

func UpdateMqtt(c *gin.Context) {


	var mqtt model.Mqtt
	c.BindJSON(&mqtt)

	// check if config.json exists
	var cmd string
	cmd = "ls /root/work/config.json"
	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		cmd = "cd /root/work && eunuch init"
		_, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			c.JSON(http.StatusInternalServerError, 
			gin.H{"message": err.Error()})
			return
		}
	}

	// Read config.json
	var devConfig model.DevConfig
	jsonConfig, err := utils.ReadJSONFile("/root/work/config.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// map to struct
	err = utils.MapToStruct(jsonConfig, &devConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// update mqtt
	devConfig.ClientId = mqtt.ClientId
	devConfig.PubTopic = mqtt.PubTopic
	devConfig.SubTopic = mqtt.SubTopic
	devConfig.UpdateTopic = mqtt.UpdateTopic
	devConfig.MqttHostUrl = mqtt.MqttHostUrl
	devConfig.Username = mqtt.Username
	devConfig.Passwd = mqtt.Passwd
	devConfig.Port = mqtt.Port

	// map to json
	mapConfig := utils.StructToMap(devConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// write to config.json
	err = utils.WriteJSONFile("/root/work/config.json", mapConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// restart eunuch 
	cmd = "systemctl restart gateway"
	_, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "restart failed",
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "update mqtt success",
	})
	return

}

func DownloadLog(c *gin.Context) {
	// c.File("/root/work/log.txt")
	var cmd string
	cmd = "rm -rf /root/work/log.zip"
	exec.Command("bash", "-c", cmd).Output()
	// cmd = "cd /root/work && tar -zcvf log.tar.gz info.log errors.log"
	cmd = "cd /root/work && zip log.zip *.log"
	_, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.File("/root/work/log.zip")
	return

}

func DeleteLog(c *gin.Context) {
	var cmd string
	cmd = "rm -rf /root/work/info.log & " +
					"rm -rf /root/work/errors.log & " +
					"rm -rf /root/work/info.log.* & " +
					"rm -rf /root/work/errors.log.*"
	exec.Command("bash", "-c", cmd).Output()

	c.JSON(http.StatusOK, gin.H{
		"message": "delete log success",
	})
	return
}


func RestartGateway(c *gin.Context) {
	var cmd string
	cmd = "systemctl restart gateway"
	_, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "restart gateway success",
	})
	return
}
