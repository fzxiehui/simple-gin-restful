package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "log"

	"github.com/fzxiehui/simple-gin-restful/model"
	"github.com/fzxiehui/simple-gin-restful/utils"
	
)

func GetSensor(c *gin.Context) {

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

	// var devConfig model.DevConfig
	var devlist []model.Device
	err := c.BindJSON(&devlist)
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

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
	return

}
