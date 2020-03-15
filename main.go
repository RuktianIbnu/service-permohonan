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
	router.POST("/service-permohonan/ubah-data", inDB.UpdateData)
	router.POST("/service-permohonan/ubahdaripetugas-data", inDB.UpdateByPetugas)
	router.POST("/service-permohonan/bystatus-data", inDB.GetDataStatus)
	router.GET("/service-permohonan/semua-data", inDB.GetAllData)
	router.GET("/service-permohonan/byid-data/:id", inDB.GetDataById)
	router.DELETE("/service-permohonan/hapus-data/:id", inDB.DeleteData)
	router.GET("/service-permohonan/test-MappingToken", inDB.TestMapping)
	router.Run(":3400")
}