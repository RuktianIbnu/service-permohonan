package main

import (
	"./config"
	"./module"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db := config.DBInit()
	inDB := &module.InDB{DB: db}
	router := gin.Default()

	router.POST("/service-permohonan/simpan-data", inDB.InsertData)
	router.GET("/service-permohonan/semua-data", inDB.GetAllData)
	router.POST("/service-permohonan/ubah-data", inDB.UpdateData)
	router.DELETE("/service-permohonan/hapus-data/:id", inDB.DeleteData)
	router.GET("/service-permohonan/test-MappingToken", inDB.TestMapping)
	router.Run(":3400")
}