package main

import "fmt"

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

	// buat test data
	data[0].nama, data[0].jurusan, data[0].nilai, data[0].username, data[0].password = "Devo", "Sains Data", 100, "dev", "o"
	data[1].nama, data[1].jurusan, data[1].nilai, data[1].username, data[1].password = "Ghanif", "Informatika", 75, "ghan", "if"
	data[2].nama, data[2].jurusan, data[2].nilai, data[2].username, data[2].password = "Adit", "Sains Data", 30, "dit", "a"
	data[3].nama, data[3].jurusan, data[3].nilai, data[3].username, data[3].password = "Damar", "Informatika", 85, "mar", "da"
	nData = 4

	login(&data, &nData)
}

// TAMPILAN ------------------------------------------------------------------------------------------------------------------------------------

func login(A *arrMahasiswa, n *int) {
	var akses int

	akses = 1 // 1 supaya dia ga masuk ke "Pilihan tidak valid. Silahkan pilih lagi [1/2/3]: "
	for akses != 3 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("---------------------------")
		fmt.Println("         MAIN MENU         ")
		fmt.Println("---------------------------")
		fmt.Println("1. Admin                   ")
		fmt.Println("2. Calon Mahasiswa         ")
		fmt.Println("3. Keluar dari aplikasi    ")
		fmt.Println("---------------------------")
		if (akses != 1) && (akses != 2) && (akses != 3) { // cabang ini kalau user milih selain pilihan
			fmt.Print("Pilihan tidak valid. Silahkan pilih lagi [1/2/3]: ")
		} else {
			fmt.Print("Pilih [1/2/3]: ")
		}
		fmt.Scan(&akses)
		if akses == 1 {
			menuAdmin(A, n)
		} else if akses == 2 {
			loginMahasiswa(A, n)
		} else if akses == 3 {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Terima kasih")
		}
	}
}

func menuAdmin(A *arrMahasiswa, n *int) {
	var akses, idx int
	var aksesDalam, user string
	
	akses = 1
	for akses != 6 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("---------------------------")
		fmt.Println("         MENU ADMIN        ")
		fmt.Println("---------------------------")
		fmt.Println("1. Tambah     			    ")
		fmt.Println("2. Ubah        			")
		fmt.Println("3. Hapus       			")
		fmt.Println("4. Tampil data  			")
		fmt.Println("5. Tampil akun calon  		")
		fmt.Println("6. Kembali ke main menu    ")
		fmt.Println("---------------------------")
		if (akses != 1) && (akses != 2) && (akses != 3) && (akses != 4) && (akses != 5) && (akses != 6) {
			fmt.Print("Pilihan tidak valid. Silahkan pilih lagi [1/2/3/4/5/6]: ")
		} else {
			fmt.Print("Pilih [1/2/3/4/5/6]: ")
		}
		fmt.Scan(&akses)
		if akses == 1 {
			aksesDalam = "y"
			for aksesDalam != "n" {
				fmt.Print("\033[H\033[2J")
				fmt.Print("Masukkan nama username: ")
				fmt.Scan(&user)
			    idx = seqSearch(*A, *n, user)
				if idx == -1 {
					fmt.Print("Username tidak ditemukan. Lanjutkan [y/n]: ")
					fmt.Scan(&aksesDalam)
				} else {
					tambah(A, idx)
					aksesDalam = "n"
				}
			}
		} else if akses == 2 {
			aksesDalam = "y"
			for aksesDalam != "n" {
				fmt.Print("\033[H\033[2J")
				fmt.Print("Masukkan nama username: ")
				fmt.Scan(&user)
			    idx = seqSearch(*A, *n, user)
				if idx == -1 {
					fmt.Print("Username tidak ditemukan. Lanjutkan [y/n]: ")
					fmt.Scan(&aksesDalam)
				} else {
					ubah(A, idx)
					aksesDalam = "n"
				}
			}
		} else if akses == 3 {
			aksesDalam = "y"
			for aksesDalam != "n" {
				fmt.Print("\033[H\033[2J")
				fmt.Print("Masukkan nama username: ")
				fmt.Scan(&user)
			    idx = seqSearch(*A, *n, user)
				if idx == -1 {
					fmt.Print("Username tidak ditemukan. Lanjutkan [y/n]: ")
					fmt.Scan(&aksesDalam)
				} else {
					hapus(A, n, idx)
					aksesDalam = "n"
				}
			}
		} else if akses == 4 {
			tampilData(A, n)
		} else if akses == 5 {
			tampilAkun(A, n)
		}
	}
}

func loginMahasiswa(A *arrMahasiswa, n *int) {
	var akses int

	akses = 1
	for akses != 3 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("---------------------------")
		fmt.Println("    MENU LOGIN MAHASISWA   ")
		fmt.Println("---------------------------")
		fmt.Println("1. Daftar akun 		    ")
		fmt.Println("2. Masuk        			")
		fmt.Println("3. Kembali ke main menu    ")
		fmt.Println("---------------------------")
		if (akses != 1) && (akses != 2) && (akses != 3) {
			fmt.Print("Pilihan tidak valid. Silahkan pilih lagi [1/2/3]: ")
		} else {
			fmt.Print("Pilih [1/2/3]: ")
		}
		fmt.Scan(&akses)
		if akses == 1 {
			daftar(A, n)
		} else if akses == 2 {
			masuk(A, n)
			akses = 3
		}
	}
}

func menuMahasiswa(A *arrMahasiswa, n *int, user string) {
	var akses, idx int

	akses = 1
	idx = seqSearch(*A, *n, user)
	for akses != 2 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("---------------------------")
		fmt.Println("       MENU MAHASISWA      ")
		fmt.Println("---------------------------")
		fmt.Println("1. Daftar jurusan          ")
		fmt.Println("2. Kembali ke main menu    ")
		fmt.Println("---------------------------")
		if (akses != 1) && (akses != 2) {
			fmt.Printf("Pilihan tidak valid. Login as %s. Silahkan pilih lagi [1/2]: ", user)
		} else if A[idx].nama != "null" {
			fmt.Printf("Akun %s sudah mendaftar. Silahkan pilih lagi [1/2]: ", user)
		} else {
			fmt.Printf("Login as %s. Pilih [1/2]: ", user)
		}
		fmt.Scan(&akses)
		if akses == 1 && A[idx].nama == "null" {
			tambah(A, idx)
		}
	}
}

// FUNGSI DAN PROSEDUR ------------------------------------------------------------------------------------------------------------------------------------

func daftar(A *arrMahasiswa, n *int) {
	var akses, user, pass string

	akses = "y"
	for akses != "n" {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Masukkan username baru: ")
		fmt.Scan(&user)
		fmt.Print("Masukkan password baru: ")
		fmt.Scan(&pass)
		if seqSearch(*A, *n, user) != -1 { // username sudah ada
			fmt.Print("Username sudah ada. Masukkan username berbeda. Lanjutkan [y/n]: ")
		} else { // username belum ada
			A[*n].username = user
			A[*n].password = pass
			A[*n].nama = "null"
			*n++ // jumlah akun bertambah
			fmt.Print("Akun telah terdaftar. Lanjutkan [y/n]: ")
		}
		fmt.Scan(&akses)
	}
}

func masuk(A *arrMahasiswa, n *int) {
	var akses, user, pass string
	var idx int

	akses = "y"
	for akses != "n" {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Masukkan username anda: ")
		fmt.Scan(&user)
		idx = seqSearch(*A, *n, user)
		if idx != -1 { // username ketemu
			fmt.Print("Masukkan password anda: ")
			fmt.Scan(&pass)
			if A[idx].password == pass { // password benar
				menuMahasiswa(A, n, user)
				akses = "n"
			} else { // password salah
				fmt.Print("Password salah. Lanjutkan [y/n]: ")
				fmt.Scan(&akses)
			}
		} else { // username tidak ketemu
			fmt.Print("Akun tidak tersedia. Lanjutkan [y/n]: ")
			fmt.Scan(&akses)
		}
	}
}

func tambah(A *arrMahasiswa, idx int) {
	var jurusan int

	fmt.Print("\033[H\033[2J")
	fmt.Print("Masukkan nama lengkap: ")
	fmt.Scan(&A[idx].nama)
	fmt.Print("Masukkan nilai test: ")
	fmt.Scan(&A[idx].nilai)
	fmt.Println("Pilihan Jurusan:            ")
	fmt.Println("1. Informatika        	     ")
	fmt.Println("2. Rekayasa Perangkat Lunak ")
	fmt.Println("3. Sains Data         	     ")
	fmt.Println("4. Teknologi Informasi	     ")
	fmt.Print("Masukkan jurusan [1/2/3/4]: ")
	jurusan = 0
	for (jurusan != 1) && (jurusan != 2) && (jurusan != 3) && (jurusan != 4) {
		fmt.Scan(&jurusan)
		if jurusan == 1 {
			A[idx].jurusan = "Informatika"
		} else if jurusan == 2 {
			A[idx].jurusan = "Rekayasa Perangkat Lunak"
		} else if jurusan == 3 {
			A[idx].jurusan = "Sains Data"
		} else if jurusan == 4 {
			A[idx].jurusan = "Teknologi Informasi"
		} else {
			fmt.Print("Pilihan tidak valid. Silahkan masukkan jurusan lagi [1/2/3/4]: ")
		}
	}
}

func ubah(A *arrMahasiswa, idx int) {
	var akses, jurusan int
	var aksesDalam string

	akses = 1; aksesDalam = "n"
	for aksesDalam != "y" && akses != 4 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("--------------------------")
		fmt.Println("    	    UBAH           ")
		fmt.Println("--------------------------")
		fmt.Println("1. Nama       		       ")
		fmt.Println("2. Jurusan   		       ")
		fmt.Println("3. Nilai test		       ")
		fmt.Println("4. Kembali ke menu admin  ")
		fmt.Println("--------------------------")
		if akses != 1 && akses != 2 && akses != 3 && akses != 4 {
			fmt.Printf("Pilihan tidak valid. Mengubah akun %s. Silahkan pilih lagi [1/2/3/4]: ", A[idx].username)
		} else {
			fmt.Printf("Mengubah akun %s. Silahkan pilih [1/2/3/4]: ", A[idx].username)
		}
		fmt.Scan(&akses)
		if akses == 1 {
			fmt.Print("\033[H\033[2J")
			fmt.Print("Masukkan nama baru: ")
			fmt.Scan(&A[idx].nama)
			fmt.Print("Berhasil mengubah. Kembali [y/n]: ")
			fmt.Scan(&aksesDalam)
		} else if akses == 2 {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Pilihan Jurusan:            ")
			fmt.Println("1. Informatika        	     ")
			fmt.Println("2. Rekayasa Perangkat Lunak ")
			fmt.Println("3. Sains Data         	     ")
			fmt.Println("4. Teknologi Informasi	     ")
			fmt.Print("Masukkan jurusan [1/2/3/4]: ")
			jurusan = 0
			for (jurusan != 1) && (jurusan != 2) && (jurusan != 3) && (jurusan != 4) {
				fmt.Scan(&jurusan)
				if jurusan == 1 {
					A[idx].jurusan = "Informatika"
				} else if jurusan == 2 {
					A[idx].jurusan = "Rekayasa Perangkat Lunak"
				} else if jurusan == 3 {
					A[idx].jurusan = "Sains Data"
				} else if jurusan == 4 {
					A[idx].jurusan = "Teknologi Informasi"
				} else {
					fmt.Print("Pilihan tidak valid. Silahkan masukkan jurusan lagi [1/2/3/4]: ")
				}
			}
			fmt.Print("Berhasil mengubah. Kembali [y/n]: ")
			fmt.Scan(&aksesDalam)
		} else if akses == 3 {
			fmt.Print("\033[H\033[2J")
			fmt.Print("Masukkan nilai test baru: ")
			fmt.Scan(&A[idx].nilai)
			fmt.Print("Berhasil mengubah. Kembali [y/n]: ")
			fmt.Scan(&aksesDalam)
		}
	}
}

func hapus(A *arrMahasiswa, n *int, idx int) {
	var i int
	var setuju string

	fmt.Print("Yakin untuk menghapus data [y/n]: ")
	fmt.Scan(&setuju)
	if setuju == "y" {
		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
	}
}

func tampilData(A *arrMahasiswa, n *int) {
	var i, terurut int
	var akses, berdasarkan string

	akses = "n"; terurut = 1
	for akses != "y" {
		fmt.Print("\033[H\033[2J")
		fmt.Println("-----------------------------")
		fmt.Println("           PILIHAN           ")
		fmt.Println("-----------------------------")
		fmt.Println("1. Nilai test secara menaik  ")
		fmt.Println("2. Nilai test secara menurun ")
		fmt.Println("3. Jurusan 				  ")
		fmt.Println("4. Nama mahasiswa A-Z 		  ")
		fmt.Println("5. Nama mahasiswa Z-A 		  ")
		fmt.Println("-----------------------------")
		if terurut != 1 && terurut != 2 && terurut != 3 && terurut != 4 && terurut != 5 {
			fmt.Print("Pilihan tidak valid. Silahkan pilih lagi [1/2/3/4/5]: ")
		} else {
			fmt.Print("Tampilkan terurut berdasarkan [1/2/3/4/5]: ")
		}
		fmt.Scan(&terurut)
		if terurut == 1 {
			nilaiAscend(A, n)
			berdasarkan = "nilai test secara menaik"
		} else if terurut == 2 {
			nilaiDescend(A, n)
			berdasarkan = "nilai test secara menurun"
		} else if terurut == 3 {
			jurusan(A, n)
			berdasarkan = "jurusan"
		} else if terurut == 4 {
			namaAscend(A, n)
			berdasarkan = "nama mahasiswa A-Z"
		} else if terurut == 5 {
			namaDescend(A, n)
			berdasarkan = "nama mahasiswa Z-A"
		}
		if terurut >= 1 && terurut <= 5 { // <- maka akan keprint
			fmt.Print("\033[H\033[2J")
			fmt.Printf("| %-5s | %-16s | %-25s | %-10s | %-16s | \n", "No", "Nama", "Jurusan", "Nilai Test", "Status")
			fmt.Printf("---------------------------------------------------------------------------------------- \n")
			for i = 0; i < *n; i++ {
				fmt.Printf("| %-5d | %-16s | %-25s | %-10d | %-16s | \n", i+1, A[i].nama, A[i].jurusan, A[i].nilai, A[i].status)
			}
			fmt.Printf("Data berdasarkan %s. Kembali [y/n]: ", berdasarkan)
			fmt.Scan(&akses)
		}
	}
}

func tampilAkun(A *arrMahasiswa, n *int) {
	var i int
	var akses string

	akses = "n"
	for akses != "y" {
		fmt.Print("\033[H\033[2J")
		fmt.Printf("| %-5s | %-16s | %-16s | \n", "No", "Username", "Password")
		fmt.Printf("----------------------------------------------- \n")
		for i = 0; i < *n; i++ {
			fmt.Printf("| %-5d | %-16s | %-16s | \n", i+1, A[i].username, A[i].password)
		}
		fmt.Print("Kembali [y/n]: ")
		fmt.Scan(&akses)
	}
}

func seqSearch(A arrMahasiswa, n int, x string) int {
	var i int
	var ketemu int

	i = 0; ketemu = -1
	for i < n && ketemu == -1 {
		if A[i].username == x {
			ketemu = i
		}
		i++
	}
	return ketemu
}

// SORTING ------------------------------------------------------------------------------------------------------------------------------------

func nilaiAscend(A *arrMahasiswa, n *int) { // selection sort
	var pass, idx, i int
	var temp mahasiswa

	pass = 1
	for pass <= *n - 1 {
		idx = pass - 1
		i = pass
		for i < *n {
			if A[i].nilai < A[idx].nilai {
				idx = i
			}
			i++
		}
		temp = A[pass - 1]
		A[pass - 1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func nilaiDescend(A *arrMahasiswa, n *int) { // selection sort
	var pass, idx, i int
	var temp mahasiswa

	pass = 1
	for pass <= *n - 1 {
		idx = pass - 1
		i = pass
		for i < *n {
			if A[i].nilai > A[idx].nilai {
				idx = i
			}
			i++
		}
		temp = A[pass - 1]
		A[pass - 1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func jurusan(A *arrMahasiswa, n *int) { // selection sort
	var pass, idx, i int
	var temp mahasiswa

	pass = 1
	for pass <= *n - 1 {
		idx = pass - 1
		i = pass
		for i < *n {
			if A[i].jurusan < A[idx].jurusan {
				idx = i
			}
			i++
		}
		temp = A[pass - 1]
		A[pass - 1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func namaAscend(A *arrMahasiswa, n *int) { // insertion sort
	var pass, i int
	var temp mahasiswa

	pass = 1
	for pass <= *n - 1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.nama < A[i-1].nama {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}

func namaDescend(A *arrMahasiswa, n *int) { // insertion sort
	var pass, i int
	var temp mahasiswa

	pass = 1
	for pass <= *n - 1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.nama > A[i-1].nama {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}