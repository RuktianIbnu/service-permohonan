package structs

type DataPermohonan struct{
	Id						int
	Email_user_pemohon 		string
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
	Date_insert_at			string
	Date_update_at			string
	Date_delete_at			string
}

func (DataPermohonan) DataPermohonan()string{
	return "dataPermohonan"
}