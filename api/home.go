package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"log"
	"strings"
	"strconv"

	"github.com/fzxiehui/simple-gin-restful/model"

)

func Home(c *gin.Context) {

	var cmd string
	cmd = "ps aux --sort=-%cpu | awk '{print $2,$3,$4,$11}'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	var process_list = model.ProcessList{}
	// string to line
	var top_n = 1

	var cpu float64 = 0.0
	var mem float64 = 0.0

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		// conntinue if line is title
		if strings.Contains(line, "PID") {
			continue
		}
		// split line to words
		words := strings.Fields(line)
		// log.Println(words)
		if len(words) == 4 {
			// return process_list 
			var item_cpu float64 = 0.0
			item_cpu, err = strconv.ParseFloat(words[1], 64)
			if err != nil {
				log.Println(err)
			}
			cpu += item_cpu

			var item_mem float64 = 0.0
			item_mem, err = strconv.ParseFloat(words[2], 64)

			if err != nil {
				log.Println(err)
			}
			mem += item_mem

			if top_n > 10 {
				break

			}else{
				top_n += 1
				process_list.ProcessList = append(process_list.ProcessList, 
					model.Process{
						Pid:  words[0],
						Cpu:  item_cpu,
						Mem:  item_mem,
						Name: words[3],
					})
			}
		}

	}

	var diskcmd string
	diskcmd = "df | grep ' /$'| awk '{print $1,$3,$4,$5}'"
	diskout, err := exec.Command("bash", "-c", diskcmd).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// var Device = diskout[0]
	// split line to words
	diskwords := strings.Fields(string(diskout))

	var disk = model.Disk{
		Device: diskwords[0],
		Used: diskwords[1],
		Avail: diskwords[2],
		Use: diskwords[3],
	}

	// get cpu thead number
	var cpu_thread string
	cpu_thread = "cat /proc/cpuinfo | grep 'processor' | wc -l"
	cpu_thread_out, err := exec.Command("bash", "-c", cpu_thread).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var cpu_thread_num int
	cpu_thread_num, err = strconv.Atoi(strings.TrimSpace(string(cpu_thread_out)))
	if err != nil {
		log.Println(err)
	}

	cpu = cpu / float64(cpu_thread_num)

	if cpu > 100 {
		cpu = 100
	}

	c.JSON(http.StatusOK, model.Home{
		ProcessList: process_list.ProcessList,
		Cpu: cpu,
		Mem: mem,
		Disk: disk,
	})
}

