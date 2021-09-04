package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func (s *student) printDataStudent() {
	fmt.Printf("Nama siswa %v \nasal %v \npekerjaan %v \nalasan ikut course %v \n", s.nama, s.alamat, s.alasan, s.alasan)
}

func main() {

	var allStudent = []student{
		{
			nama:      "Firmat Aulia",
			alamat:    "Sleman",
			pekerjaan: "Mahasiswa",
			alasan:    "Belajar API",
		},
		{
			nama:      "Andri Nur Hidayatulloh",
			alamat:    "Bantul Yogyakarta",
			pekerjaan: "Mahasiswa",
			alasan:    "Belajar Back End",
		},
		{
			nama:      "I Komang Widnyana",
			alamat:    "Bali",
			pekerjaan: "Mahasiswa",
			alasan:    "Belajar membuat API",
		},
		{
			nama:      "Erico",
			alamat:    "Medan",
			pekerjaan: "Mahasiswa",
			alasan:    "Belajar menjadi Fullstack Dev",
		},
		{
			nama:      "Jose Yolanda Purba",
			alamat:    "Bandung",
			pekerjaan: "Mahasiswa",
			alasan:    "Ingin mempelajari Back End",
		},
		{
			nama:      "Andri Kuwito",
			alamat:    "Jakarta",
			pekerjaan: "Mahasiswa",
			alasan:    "Penasaran sama Golang",
		},
		{
			nama:      "Sandy Budiman",
			alamat:    "Bandung",
			pekerjaan: "Mahasiswa",
			alasan:    "Migrasi dari Dev Ops ke Golang",
		},
		{
			nama:      "Rafly Andreansyah",
			alamat:    "Malang",
			pekerjaan: "Mahasiswa",
			alasan:    "Mencoba Golang",
		},
		{
			nama:      "Taufiq Hidayah",
			alamat:    "Kalimantan",
			pekerjaan: "Mahasiswa",
			alasan:    "Penasaran sama Micrservices Golang",
		},
		{
			nama:      "Melvita Sari",
			alamat:    "Aceh",
			pekerjaan: "Mahasiswa",
			alasan:    "Ingin menjadi Back End Dev",
		},
	}

	valueOfARGS, _ := strconv.Atoi(os.Args[1])
	valueOfARGS -= 1

	if valueOfARGS >= 0 && valueOfARGS < len(allStudent) {
		allStudent[valueOfARGS].printDataStudent()
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}
