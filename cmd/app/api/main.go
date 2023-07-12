package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"path/filepath"
	"sharing-vision-2021/initiator"
	"strings"
)

func main() {
	LoadEnvVars()
	i := initiator.NewInit()

	r := i.GetGin()
	db := i.GetDB()

	r.GET("/ping", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
		if err = sqlDB.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run("localhost:7000")
}

func LoadEnvVars() {
	cwd, _ := os.Getwd()
	dirString := strings.Split(cwd, "sharing-vision-2021")
	dir := strings.Join([]string{dirString[0], "sharing-vision-2021"}, "")
	AppPath := dir

	godotenv.Load(filepath.Join(AppPath, "/.env"))
}