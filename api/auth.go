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

	c.JSON(http.StatusOK, auth)
}



// func Register(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "register",
// 	})
// }


