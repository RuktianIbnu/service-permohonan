package handler

import (
	"net/http"
    "../structs"
    "../database"
    
    "github.com/gin-gonic/gin"
)

func InsertData(c *gin.Context){
	var (
		data structs.DataPermohonan
	)

    email_user_pemohon := c.PostForm("email_user_pemohon")
    status := c.PostForm("status")
    nama_server := c.PostForm("nama_server")
    detail := c.PostForm("detail")
    jenis_server := c.PostForm("jenis_server")
    os := c.PostForm("os")
    ram := c.PostForm("ram")
    storage := c.PostForm("storage")
    hostname := c.PostForm("hostname")
    internet := c.PostForm("internet")
    internet_status := c.PostForm("internet_status")
    open_port := c.PostForm("open_port")
    lokasi := c.PostForm("lokasi")
    id_kontainment := c.PostForm("id_kontainment")
    rak :=  c.PostForm("rak")


    data.Email_user_pemohon = email_user_pemohon
    data.Status             = status
    data.Nama_server        = nama_server
    data.Detail             = detail
    data.Jenis_server       = jenis_server
    data.Os                 = os
    data.Ram                = ram
    data.Storage            = storage
    data.Hostname           = hostname
    data.Internet           = internet
    data.Internet_status    = internet_status
    data.Open_port          = open_port
    data.Lokasi             = lokasi
    data.Id_kontainment     = id_kontainment
    data.Rak                = rak 

    err := database.Insert(data)
    if err == false {
        c.JSON(http.StatusInternalServerError, gin.H {
            "pesan": "gagal simpan data",
            "status": "error",
        })
        c.Abort()
    } else {
        c.JSON(http.StatusOK, gin.H {
            "status": "success",
            "pesan": "Data berhasil di simpan",
        })
    }
}