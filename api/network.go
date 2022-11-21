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

func GetNetwork(c *gin.Context) {
	
	// get network interface list 
	var cmd string
	cmd = "ls /sys/class/net"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// string to line
	lines := strings.Split(string(out), "\n")
	var network_list = model.NetworkList{}
	for _, line := range lines {
		// conntinue if line is title
		if strings.Contains(line, "lo") {
			continue
		}
		// split line to words
		words := strings.Fields(line)
		// log.Println(words)
		if len(words) == 1 {
			// return network_list 
			// get network interface operstate
			cmd = "cat /sys/class/net/" + words[0] + "/operstate"
			out, err := exec.Command("bash", "-c", cmd).Output()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}
			// string to line
			operstate := strings.Split(string(out), "\n")
			// log.Println(operstate)
			// get network interface inet
			cmd = "ip addr show " + words[0] + " | grep inet | awk '{print $2}'"
			out, err = exec.Command("bash", "-c", cmd).Output()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}
			// string to line
			inet := strings.Split(string(out), "\n")
			// log.Println(inet)
			// get network interface netmask
			// del mask cmd = "ip addr show " + words[0] + " | grep inet | awk '{print $2}' | cut -d '/' -f 2"
			// del mask out, err = exec.Command("bash", "-c", cmd).Output()
			// del mask if err != nil {
			// del mask 	c.JSON(http.StatusInternalServerError, gin.H{
			// del mask 		"message": err.Error(),
			// del mask 	})
			// del mask 	return
			// del mask }
			// del mask // string to line
			// del mask netmask := strings.Split(string(out), "\n")
			// log.Println(netmask)
			// del mask // get network interface broadcast
			// del mask cmd = "ip addr show " + words[0] + " | grep inet | awk '{print $4}' | cut -d '/' -f 1"
			// del mask out, err = exec.Command("bash", "-c", cmd).Output()
			// del mask if err != nil {
			// del mask 	c.JSON(http.StatusInternalServerError, gin.H{
			// del mask 		"message": err.Error(),
			// del mask 	})
			// del mask 	return
			// del mask }
			// del mask // string to line
			// del mask broadcast := strings.Split(string(out), "\n")
			// log.Println(broadcast)
			// get network interface gateway
			gateway := ""
			var dhcp_tag = true
			if operstate[0] == "up" {
				cmd = "ip route show | grep default | awk '{print $3}'"
				out, err = exec.Command("bash", "-c", cmd).Output()
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": err.Error(),
					})
					return
				}
				// string to line
				gateway = strings.TrimSpace(string(out))

				// get network interface dhcp 
				cmd = "nmcli connection show " + words[0] + "_config"
				out, _ = exec.Command("bash", "-c", cmd).Output()
				// string to line
				dhcp := strings.Split(string(out), "\n")
				// log.Println(dhcp)

				// Error in dhcp[0] is "Error: Connection 'eth0_config' not found."
				if strings.Contains(dhcp[0], "Error") {
					dhcp_tag = true
				} else {
					dhcp_tag = false
				}
			}

		
			var dns string = gateway
			if !dhcp_tag {
				// get network interface dns
				cmd = "nmcli connection show " + words[0] + "_config | grep dns | awk '{print $2}'"
				out, _ = exec.Command("bash", "-c", cmd).Output()
				// if err != nil {
				// 	c.JSON(http.StatusInternalServerError, gin.H{
				// 		"message": err.Error(),
				// 	})
				// 	return
				// }
				// string
				dns = string(out)
			}


			// log.Println(gateway)
			// add network interface to network_list
			network_list.NetworkList = append(network_list.NetworkList, model.Network{
				Name: words[0],
				Operstate: operstate[0],
				Inet: inet[0],
				// Netmask: netmask[0],
				// Broadcast: broadcast[0],
				Gateway: gateway,
				DNS: dns,
				DHCP: dhcp_tag,
			})


			// var item = model.Network{}
			// item.Name = words[0]
			// network_list.NetworkList = append(network_list.NetworkList, item)

		}
	}

	// return
	c.JSON(http.StatusOK, network_list.NetworkList)


}
