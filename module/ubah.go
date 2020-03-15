package module

import (
	"net/http"
    "../structs"
    
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

func (idb *InDB) UpdateData(c *gin.Context) {
	var (
		data structs.DataPermohonan
		NewData structs.DataPermohonan
		result gin.H
	)
	id 					:= c.PostForm("id")
    nama_server         := c.PostForm("nama_server")
    detail              := c.PostForm("detail")
    jenis_server        := c.PostForm("jenis_server")
    os                  := c.PostForm("os")
    ram                 := c.PostForm("ram")
    storage             := c.PostForm("storage")
    hostname            := c.PostForm("hostname")
    internet            := c.PostForm("internet")
    internet_status     := c.PostForm("internet_status")
    open_port           := c.PostForm("open_port")
    lokasi              := c.PostForm("lokasi")
    id_kontainment      := c.PostForm("id_kontainment")
    rak                 :=  c.PostForm("rak")

	access := CekAuth(c)

	if access == false {
		result = gin.H{
			"pesan": "not authorized",
			"status": "error",
		}
	} else {
		CekData := idb.DB.Where("id = ?", id).First(&data).Error 
		if CekData != nil {
			result = gin.H{
				"pesan": "data tidak ditemukan",
				"status": "not_found",
			}
		} else if CekData == nil {
			NewData.Nama_server        = nama_server
			NewData.Detail             = detail
			NewData.Jenis_server       = jenis_server
			NewData.Os                 = os
			NewData.Ram                = ram
			NewData.Storage            = storage
			NewData.Hostname           = hostname
			NewData.Internet           = internet
			NewData.Internet_status    = internet_status
			NewData.Open_port          = open_port
			NewData.Lokasi             = lokasi
			NewData.Id_kontainment     = id_kontainment
			NewData.Rak                = rak 
			err := idb.DB.Model(&data).Where("id = ?", id).Updates(NewData).Error
			if err != nil {
				result = gin.H{
					"pesan": "update failed",
					"status": "error",
				}
			} else {
				result = gin.H{
					"pesan": "successfully updated data",
					"status": "success",
				}
			}	
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateByPetugas(c *gin.Context) {
	var (
		data structs.DataPermohonan
		NewData structs.DataPermohonan
		result gin.H
	)
	id 					:= c.PostForm("id")
    status         		:= c.PostForm("status")
    id_petugas_approval := c.PostForm("id_petugas_approval")

	access := CekAuth(c)

	if access == false {
		result = gin.H{
			"pesan": "not authorized",
			"status": "error",
		}
	} else {
		CekData := idb.DB.Where("id = ?", id).First(&data).Error 
		if CekData != nil {
			result = gin.H{
				"pesan": "data tidak ditemukan",
				"status": "not_found",
			}
		} else if CekData == nil {
			NewData.Status        				= status
			NewData.Id_petugas_approval         = id_petugas_approval
			
			err := idb.DB.Model(&data).Where("id = ?", id).Updates(NewData).Error
			if err != nil {
				result = gin.H{
					"pesan": "update failed",
					"status": "error",
				}
			} else {
				result = gin.H{
					"pesan": "successfully updated data",
					"status": "success",
				}
			}	
		}
	}
	c.JSON(http.StatusOK, result)
}
