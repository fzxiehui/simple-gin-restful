package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	
	"log"
	"github.com/fzxiehui/simple-gin-restful/model"
)


func Login(c *gin.Context) {

	var auth model.Auth
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(auth.Username)
	log.Println(auth.Password)
	// get header Referer from request
	referer := c.Request.Header.Get("Referer")
	if auth.Username == "admin" && auth.Password == "admin" {

		// set cookie
		c.SetCookie("token", "admin", 3600, "/", referer, false, false)

		c.JSON(http.StatusOK, gin.H{
			"message": "login success",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "login failed",
		})
	}
}



// func Register(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "register",
// 	})
// }


