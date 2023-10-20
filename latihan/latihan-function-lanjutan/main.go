package main

import "fmt"

func main() {
	nilaiSiswa := []float32{85.5, 92.0, 78.5, 90.0, 88.5}

	sumSiswa, averageSiswa := calculate(nilaiSiswa)

	fmt.Println("Jumlah siswa dalam kelas:", len(nilaiSiswa))
	fmt.Println("Jumlah nilai siswa dalam kelas:", sumSiswa)
	fmt.Println("Nilai rata-rata siswa dalam kelas:", averageSiswa)
}

func calculate(nilaiSiswa []float32) (sumCalculate float32, avgCalucalte float32) {

	for _, siswa := range nilaiSiswa {
		sumCalculate += siswa
	}
	avgCalucalte = sumCalculate / float32(len(nilaiSiswa))

	return
}
