package main

import "fmt"

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
	fmt.Print("Masukkan nama (Tanpa spasi): ")
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
			fmt.Print("Masukkan nama baru (Tanpa spasi): ")
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

	passingGrade(A, n)
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
			fmt.Printf("--------------------------------------------------------------------------------------------\n")
			fmt.Printf("| %-5s | %-20s | %-25s | %-10s | %-16s | \n", "No", "Nama", "Jurusan", "Nilai Test", "Status")
			fmt.Printf("--------------------------------------------------------------------------------------------\n")
			for i = 0; i < *n; i++ {
				fmt.Printf("| %-5d | %-20s | %-25s | %-10d | %-16s | \n", i+1, A[i].nama, A[i].jurusan, A[i].nilai, A[i].status)
			}
			fmt.Printf("--------------------------------------------------------------------------------------------\n")
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
		fmt.Printf("----------------------------------------------------------------------\n")
		fmt.Printf("| %-5s | %-16s | %-16s | %-20s | \n", "No", "Username", "Password", "Nama")
		fmt.Printf("----------------------------------------------------------------------\n")
		for i = 0; i < *n; i++ {
			fmt.Printf("| %-5d | %-16s | %-16s | %-20s | \n", i+1, A[i].username, A[i].password,A[i].nama)
		}
		fmt.Printf("----------------------------------------------------------------------\n")
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

func passingGrade(A *arrMahasiswa, n *int) {
	var i int
	var IF, RPL, SD, TI int
	var nIF, nRPL, nSD, nTI int
	var PGIF, PGRPL, PGSD, PGTI int

	for i = 0; i < *n; i++ {
		switch A[i].jurusan {
		case "Informatika":
			IF = IF + A[i].nilai
			nIF++
		case "Rekayasa Perangkat Lunak":
			RPL = RPL + A[i].nilai
			nRPL++
		case "Sains Data":
			SD = SD + A[i].nilai
			nSD++
		case "Teknologi Informasi":
			TI = TI + A[i].nilai
			nTI++
		}
	}

	if nIF != 0 {
		PGIF = IF / nIF
		for i = 0; i < *n; i++ {
			if A[i].jurusan == "Informatika" {
				if A[i].nilai >= PGIF {
					A[i].status = "Diterima"
				} else {
					A[i].status = "Ditolak"
				}
			}
		}
	}

	if nRPL != 0 {
		PGRPL = RPL / nRPL
		for i = 0; i < *n; i++ {
			if A[i].jurusan == "Rekayasa Perangkat Lunak" {
				if A[i].nilai >= PGRPL {
					A[i].status = "Diterima"
				} else {
					A[i].status = "Ditolak"
				}
			}
		}
	}

	if nSD != 0 {
		PGSD = SD / nSD
		for i = 0; i < *n; i++ {
			if A[i].jurusan == "Sains Data" {
				if A[i].nilai >= PGSD {
					A[i].status = "Diterima"
				} else {
					A[i].status = "Ditolak"
				}
			}
		}
	}

	if nTI != 0 {
		PGTI = TI / nTI
		for i = 0; i < *n; i++ {
			if A[i].jurusan == "Teknologi Informasi" {
				if A[i].nilai >= PGTI {
					A[i].status = "Diterima"
				} else {
					A[i].status = "Ditolak"
				}
			}
		}
	}
}