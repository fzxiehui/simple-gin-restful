package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	// "log"
	// "strings"

	"github.com/fzxiehui/simple-gin-restful/model"
	"github.com/fzxiehui/simple-gin-restful/utils"
	
)

func GetSensor(c *gin.Context) {

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

	sensor, err := utils.ReadJSONFile("/root/work/config.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// log.Println(sensor)
	// utils.PrintJSON(sensor)

	var devConfig model.DevConfig
	err = utils.MapToStruct(sensor, &devConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, devConfig.Devices)
	return
}

// UpdateSensor 
func UpdateSensor(c *gin.Context) {

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

	// var devConfig model.DevConfig
	var devlist []model.Device
	err = c.BindJSON(&devlist)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// read config.json
	sensor, err := utils.ReadJSONFile("/root/work/config.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var devConfig model.DevConfig
	err = utils.MapToStruct(sensor, &devConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// mapDevList := utils.StructToMap(devlist)

	devConfig.Devices = devlist

	mapConfig := utils.StructToMap(devConfig)

	// write config.json
	err = utils.WriteJSONFile("/root/work/config.json", mapConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// restart eunuch
	cmd = "supervisorctl restart gateway"
	_, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "restart failed",
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{"message": "ok"})
	return

}
