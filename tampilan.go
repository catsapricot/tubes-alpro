package main

import "fmt"

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