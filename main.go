package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	absen		    int
	nama              		  string
	alamat             		  string
	pekerjaan          		  string
	alasanMemilihKelas Golang string
}

func main() {

	argument := os.Args
	printData(argument)

}

func printData(arguments []string) {

	startIndex, _ := strconv.Atoi(arguments[1])

	for _, value := range dummyData() {
		if value.absen == startIndex {
			fmt.Println("---------------------------- Data Siswa -------------------------------")
			fmt.Println("Nama							:", value.nama)
			fmt.Println("Alamat							:", value.alamat)
			fmt.Println("Pekerjaan						:", value.pekerjaan)
			fmt.Println("Alasan memilih kelas golang	:", value.alasanMemilihkelasGolang)
			fmt.Println("-----------------------------------------------------------------------")
			return
		}
	}

	fmt.Println("Siswa tidak ditemukan")

}

func dummyData() []student {
	students := []student{
		{1, "Mursyied Qathafi", "Surabaya", "Freshgraduate", "Because i need as backend developer"},
		{2, "Mursyied Q", "Surabaya", "Freelance", "Because golang very populer for backend"},
		{3, "Qathafi", "Surabaya", "Fullsatck developer", "Golang very sustainable for backend"},
	}

	return students
}