package main

import (
	"./config"
	"./handler"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db := config.DBInit()
	inDB := &handler.InDB{DB: db}
	router := gin.Default()

	router.POST("/service-permohonan/simpan-data", inDB.InsertData)
	router.GET("/service-permohonan/semua-data", inDB.GetAllData)
	router.Run(":3400")
}