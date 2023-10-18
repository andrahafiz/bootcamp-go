package main

import "fmt"

type OrangTua struct {
	Nama string
	Umur int
}

type Siswa struct {
	OrangTua OrangTua // Embedded struct
	Nama     string
	Umur     int
	Kelas    string
}

type Siswas []Siswa

func main() {

	var daftarSiswa Siswas

	var siswa1 = Siswa{
		OrangTua: OrangTua{
			Nama: "Rudi",
			Umur: 40,
		},
		Nama:  "Eva",
		Umur:  12,
		Kelas: "6B",
	}

	var siswa2 = Siswa{
		OrangTua: OrangTua{
			Nama: "Dewi",
			Umur: 38,
		},
		Nama:  "Faisal",
		Umur:  11,
		Kelas: "5A",
	}

	var siswa3 = Siswa{
		OrangTua: OrangTua{
			Nama: "Eva",
			Umur: 35,
		},
		Nama:  "Fina",
		Umur:  16,
		Kelas: "10C",
	}
	// Informasi Siswa 1:
	// Nama: Ali, Umur: 15, Kelas: 9A
	// Orang Tua: Budi, Umur: 40
	fmt.Println("Informasi Siswa 1 :")
	fmt.Printf("Nama : %s , Umur : %d , Kelas : %s\n", siswa1.Nama, siswa1.Umur, siswa1.Kelas)
	fmt.Printf("Orang tua : %s, Umur : %d \n\n", siswa1.OrangTua.Nama, siswa1.OrangTua.Umur)

	fmt.Println("Informasi Siswa 2 :")
	fmt.Printf("Nama : %s , Umur : %d , Kelas : %s\n", siswa2.Nama, siswa2.Umur, siswa2.Kelas)
	fmt.Printf("Orang tua : %s, Umur : %d \n\n", siswa2.OrangTua.Nama, siswa2.OrangTua.Umur)

	fmt.Println("Informasi Siswa 3 :")
	fmt.Printf("Nama : %s , Umur : %d , Kelas : %s\n", siswa3.Nama, siswa3.Umur, siswa3.Kelas)
	fmt.Printf("Orang tua : %s, Umur : %d \n\n", siswa3.OrangTua.Nama, siswa3.OrangTua.Umur)

	daftarSiswa = append(daftarSiswa, siswa1)
	daftarSiswa = append(daftarSiswa, siswa2)
	daftarSiswa = append(daftarSiswa, siswa3)

	fmt.Println("Daftar Siswa:")
	for _, student := range daftarSiswa {
		fmt.Printf("Nama : %s , Umur : %d , Kelas : %s\n", student.Nama, student.Umur, student.Kelas)
		fmt.Printf("Orang tua : %s, Umur : %d \n\n", student.OrangTua.Nama, student.OrangTua.Umur)
	}
}
