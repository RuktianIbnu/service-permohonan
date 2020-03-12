package database


import (
	"strconv"
	"../structs"
	"../config"
	"time"
)
type Mysql struct {
	table string
}

//set a table name to perform action
func (m *Mysql)Table(name string) {
	m.table = name
}

//perform your raw query
func (m Mysql)Query(q string) string {
	return "Query on -> " + m.table + " " + q
}

//find a record using it's primary id
func (m Mysql)Find(id int) string {
	return "SELECT * FROM " + m.table + " WHERE id=" + strconv.Itoa(id)
}

//fetch all the records from the table
func (m Mysql)Get() string {
	return "SELECT * FROM " + m.table
}

func Insert(data structs.DataPermohonan) bool{
	db := config.DbConn()
	dt := time.Now()

	_ , err := db.Exec("INSERT INTO permohonan (email_user_pemohon, status, nama_server,  detail, jenis_server, os, ram, storage, hostname,	internet, internet_status, open_port, lokasi, id_kontainment, rak, id_petugas_approval, date_insert_at, date_update_at, date_delete_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", 
		data.Email_user_pemohon,
		data.Status,
		data.Nama_server,
		data.Detail,
		data.Jenis_server,
		data.Os,
		data.Ram,
		data.Storage,
		data.Hostname,
		data.Internet,
		data.Internet_status,
		data.Open_port,
		data.Lokasi,
		data.Id_kontainment,
		data.Rak,
		data.Id_petugas_approval,
		dt,
		"",
		"",
	)

	if err !=  nil {
		return false
	}  else {
		return  true
	}
}