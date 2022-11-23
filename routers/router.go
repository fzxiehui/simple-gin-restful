package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/fzxiehui/simple-gin-restful/api"
)


func InitRouter(r *gin.Engine) {
	V1API(r)
}

func V1API(r *gin.Engine) {
	v := r.Group("/v1")
	{
		// auth
		v.GET("/auth/login", api.Login)
		v.POST("/auth/login", api.Login)

		// home 
		v.GET("/home", api.Home)

		// network
		v.GET("/network", api.GetNetwork)
		v.PUT("/network", api.UpdateNetwork)

		// wlan
		v.GET("/wlan", api.GetWlan)
		v.PUT("/wlan", api.UpdateWlan)

		// sensor
		v.GET("/sensor", api.GetSensor)
		v.PUT("/sensor", api.UpdateSensor)

		// system
		v.GET("/system", api.GetSystem)

		// system -> time zone
		// v.GET("/system/timezone", api.GetTimeZone)
		v.PUT("/system/timezone", api.UpdateTimeZone)



		// system -> uart
		// v.GET("/system/uart", api.GetUart)
		v.PUT("/system/uart", api.UpdateUart)

		// system -> mqtt
		// v.GET("/system/mqtt", api.GetMqtt)
		v.PUT("/system/mqtt", api.UpdateMqtt)

	}

}

