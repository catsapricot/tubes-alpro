package main

type mahasiswa struct {
	nama, jurusan string
	nilai int
	status string // keterima atau tidak
	username, password string // akun
}
const NMAX int = 100
type arrMahasiswa [NMAX]mahasiswa

func main() {
	var nData int
	var data arrMahasiswa

	// Data dummy
	data[0].nama, data[0].jurusan, data[0].nilai, data[0].username, data[0].password = "Devo_Hidayatullah", "Sains Data", 100, "dev", "o"
	data[1].nama, data[1].jurusan, data[1].nilai, data[1].username, data[1].password = "Ghanif_Akbar_Christy", "Informatika", 75, "ghan", "if"
	data[2].nama, data[2].jurusan, data[2].nilai, data[2].username, data[2].password = "Adit_Hidayatullah", "Sains Data", 55, "dit", "a"
	data[3].nama, data[3].jurusan, data[3].nilai, data[3].username, data[3].password = "Damar_China", "Rekayasa Perangkat Lunak", 85, "mar", "da"
	data[4].nama, data[4].jurusan, data[4].nilai, data[4].username, data[4].password = "Nina_Sari_Putri", "Sains Data", 95, "nin", "a"
	data[5].nama, data[5].jurusan, data[5].nilai, data[5].username, data[5].password = "Budi_Wijaya", "Informatika", 65, "bud", "i"
	data[6].nama, data[6].jurusan, data[6].nilai, data[6].username, data[6].password = "Citra_Dewi", "Teknologi Informasi", 80, "cit", "ra"
	data[7].nama, data[7].jurusan, data[7].nilai, data[7].username, data[7].password = "Pascabowo_Subiyanto", "Sains Data", 70, "eko", "ko"
	data[8].nama, data[8].jurusan, data[8].nilai, data[8].username, data[8].password = "Fajar_Nugroho", "Informatika", 90, "faj", "ar"
	data[9].nama, data[9].jurusan, data[9].nilai, data[9].username, data[9].password = "Gita_Purnama", "Rekayasa Perangkat Lunak", 85, "git", "a"
	data[10].nama, data[10].jurusan, data[10].nilai, data[10].username, data[10].password = "Hari_Setiawan", "Teknologi Informasi", 60, "har", "i"
	data[11].nama, data[11].jurusan, data[11].nilai, data[11].username, data[11].password = "Irma_Wati", "Informatika", 55, "irm", "a"
	data[12].nama, data[12].jurusan, data[12].nilai, data[12].username, data[12].password = "Joko_Wididi", "Rekayasa Perangkat Lunak", 75, "jok", "o"
	data[13].nama, data[13].jurusan, data[13].nilai, data[13].username, data[13].password = "Kiki_Rahayu", "Sains Data", 85, "kik", "i"
	data[14].nama, data[14].jurusan, data[14].nilai, data[14].username, data[14].password = "Lina_Sari", "Informatika", 70, "lin", "a"
	data[15].nama, data[15].jurusan, data[15].nilai, data[15].username, data[15].password = "Mila_Astuti", "Rekayasa Perangkat Lunak", 95, "mil", "a"
	data[16].nama, data[16].jurusan, data[16].nilai, data[16].username, data[16].password = "Nando_Purnomo", "Teknologi Informasi", 80, "nan", "do"
	data[17].nama, data[17].jurusan, data[17].nilai, data[17].username, data[17].password = "Oki_Setiawan", "Informatika", 65, "oki", "i"
	data[18].nama, data[18].jurusan, data[18].nilai, data[18].username, data[18].password = "Putri_Sari", "Rekayasa Perangkat Lunak", 90, "put", "ri"
	data[19].nama, data[19].jurusan, data[19].nilai, data[19].username, data[19].password = "Qori_Amalia", "Sains Data", 75, "qor", "i"
	data[20].nama, data[20].jurusan, data[20].nilai, data[20].username, data[20].password = "Zaini_Ganteng", "Teknologi Informasi", 80, "rin", "a"
	nData = 21

	login(&data, &nData)
}