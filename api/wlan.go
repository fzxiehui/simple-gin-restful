package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	// "log"
	"strings"
	// "strconv"

	"github.com/fzxiehui/simple-gin-restful/model"
)

func GetWlan(c *gin.Context) {
	
	var cmd string
	// get nmcli wlan status
	var wlan model.Wlan
	cmd = "nmcli connection show |grep ap"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		wlan.Status = false
		wlan.Ssid = ""
		wlan.Ip = ""
		c.JSON(http.StatusOK, wlan)
		return
	}
	// log.Println(string(out))
	// parse nmcli wlan status
	if strings.Contains(string(out), "ap") {
		wlan.Status = true
	} else {
		wlan.Status = false
	}

	// get nmcli wlan ssid
	cmd = "nmcli -t -f active,ssid dev wifi |grep yes"
	out, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// log.Println(string(out))
	// parse nmcli wlan ssid
	if strings.Contains(string(out), "yes") {
		wlan.Ssid = strings.TrimSpace(
										strings.Split(string(out), ":")[1])
	} else {
		wlan.Ssid = ""
	}
	
	// get ip addr  wlan ip
	cmd = "ip addr show wlan0 |grep inet |grep -v inet6 |awk '{print $2}' |awk -F '/' '{print $1}'"
	out, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// log.Println(string(out))
	// TrimSpace 
	wlan.Ip = strings.TrimSpace(string(out))



	c.JSON(http.StatusOK, wlan)
	return

}

func UpdateWlan(c *gin.Context) {

	var wlan model.WlanAp
	if err := c.ShouldBindJSON(&wlan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// delete old ap connection
	var cmd string
	cmd = "nmcli connection show |grep ap"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if strings.Contains(string(out), "unknown connection") {
		cmd = "nmcli connection delete ap"
		out, err = exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "delete old ap connection failed"})
			return
		}
	}

	// add new ap connection
	cmd = "nmcli device wifi hotspot ifname wlan0 con-name ap ssid " + 
						wlan.Ssid + 
						" password " + 
						wlan.Password

	out, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "add new ap connection failed"})
		return
	}

	// set auto connect
	cmd = "nmcli connection modify ap connection.autoconnect yes"
	out, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "set auto connect failed"})
		return
	}


	// return ok
	c.JSON(http.StatusOK, gin.H{"message": "ok"})

}

