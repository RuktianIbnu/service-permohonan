package module

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
	tokenIsExist := CekAuth(c)
	if tokenIsExist == true {
		idb.DB.Find(&data)
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
	} else {
		result = gin.H {
			"status": "error",
			"pesan": "not authorized",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetDataById(c *gin.Context){
	var (
		data structs.DataPermohonan
		result gin.H
		id string
	)
	tokenIsExist := CekAuth(c)
	if tokenIsExist == true {
		id = c.Param("id")
		err := idb.DB.Find(&data, id).Error
		if err != nil {
			result = gin.H {
				"result": nil,
				"count": 0,
				"status": "Data Kosong",
			}
		} else {
			result = gin.H {
				"status": "success",
				"pesan": "get data by id = " + id,
				"data_permohonan": data,
			}
		}
	} else {
		result = gin.H {
			"status": "error",
			"pesan": "not authorized",
		}
	}
	c.JSON(http.StatusOK, result)
}