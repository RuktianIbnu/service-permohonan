package module

import (
	"net/http"
	"../structs"
	
	"github.com/gin-gonic/gin"
)

func (idb *InDB) DeleteData(c *gin.Context){
	var (
		data structs.DataPermohonan
		result gin.H
	)
		id := c.Param("id")
		err := idb.DB.Where("id = ?", id).First(&data).Error
		
		access := CekAuth(c)
		if access == false {
			result = gin.H{
				"pesan": "not authorized",
				"status": "error",
			}
		} else {
			if err != nil {
				result = gin.H{
					"pesan": "data tidak ditemukan",
					"status": "not found",
				}
			} else {
				err = idb.DB.Delete(&data).Error
				if err != nil {
					result = gin.H{
						"pesan": "delete failed",
						"status": "error",
					}
				} else {
					result = gin.H{
						"pesan": "Data deleted successfully",
						"status": "success",
					}
				}
			}
		}	
	c.JSON(http.StatusOK, result)
}