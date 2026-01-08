package main

import "fmt"

func main() {
	passwords := make(map[string]string) 
	var choice int

	for {
		fmt.Println("\n=== Password Manager Sederhana ===")
		fmt.Println("1. Tambah Password")
		fmt.Println("2. Lihat Password")
		fmt.Println("3. Hapus Password")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var akun, password string
			fmt.Print("Nama akun: ")
			fmt.Scanln(&akun)
			fmt.Print("Password: ")
			fmt.Scanln(&password)

			passwords[akun] = password
			fmt.Println("Password berhasil disimpan.")

		case 2:
			if len(passwords) == 0 {
				fmt.Println("Belum ada password tersimpan.")
			} else {
				fmt.Println("Daftar Password:")
				for akun, password := range passwords {
					fmt.Println("Akun:", akun, "| Password:", password)
				}
			}

		case 3:
			var akun string
			fmt.Print("Masukkan nama akun yang ingin dihapus: ")
			fmt.Scanln(&akun)

			if _, exists := passwords[akun]; exists {
				delete(passwords, akun)
				fmt.Println("Password berhasil dihapus.")
			} else {
				fmt.Println("Akun tidak ditemukan.")
			}

		case 4:
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
