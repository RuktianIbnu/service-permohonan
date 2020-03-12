package main

import (
	"./config"
	"./handler"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db := config.DbConn()
	router := gin.Default()

	defer db.Close()
	router.POST("/service-permohonan/simpan-data", handler.InsertData)
	router.Run(":3400")
}