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

	// get time zone
	var cmd string

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
