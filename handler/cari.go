package handler

import (
	"net/http"
	"../structs"
	"github.com/gin-gonic/gin"

)

func (idb *InDB) GetAllData(c *gin.Context){
	var (
		data []structs.DataPermohonan
		result gin.H
	)

	idb.DB.Limit(20).Find(&data)
	if len(data) <= 0 {
		result = gin.H {
			"result": nil,
			"count": 0,
			"status": "Data Kosong",
		}
	} else {
		result = gin.H {
			"status": "success",
			"pesan": "get all data",
			"data_permohonan": data,
			"total": len(data),
		}
	}
	c.JSON(http.StatusOK, result)
}