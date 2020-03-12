package structs

import "github.com/jinzhu/gorm"

type DataPermohonan struct{
	gorm.Model
	Email_user_pemohon 		string
	Nip 					string
	Status 					string
	Nama_server 			string
	Detail 					string
	Jenis_server 			string
	Os 						string
	Ram 					string
	Storage					string
	Hostname				string
	Internet				string
	Internet_status			string
	Open_port				string
	Lokasi					string	
	Id_kontainment			string
	Rak						string
	Id_petugas_approval		string
}

func (DataPermohonan) DataPermohonan()string{
	return "dataPermohonan"
}