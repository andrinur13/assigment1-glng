package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var students = []map[string]string{
		map[string]string{
			"nama":      "Firman Aulia",
			"alamat":    "Sleman",
			"pekerjaan": "Mahasiswa",
			"alasan":    "Belajar API",
		},

		map[string]string{
			"nama":      "Andri Nur Hidayatulloh",
			"alamat":    "Bantul, Yogyakarta",
			"pekerjaan": "Mahasiswa",
			"alasan":    "Belajar Back End",
		},

		map[string]string{
			"nama":      "I Komang Widnyana",
			"alamat":    "Bali",
			"pekerjaan": "Mahasiswa",
			"alasan":    "Belajar membuat API",
		},

		map[string]string{
			"nama":      "Erico",
			"alamat":    "Medan",
			"pekerjaan": "Mahasiswa",
			"alasan":    "Belajar menjadi Fullstack Dev.",
		},

		map[string]string{
			"nama":      "Jose Yolanda Purba",
			"alamat":    "Bandung",
			"pekerjaan": "Mahasiswa",
			"alasan":    "Ingin mempelajari Back End",
		},

		map[string]string{
			"nama":      "Andri Kuwito",
			"alamat":    "Jakarta",
			"pekerjaan": "Mahasiswa",
			"alasan":    "Penasaran sama Golang",
		},

		map[string]string{
			"nama":      "Sandy Budiman",
			"alamat":    "Bandung",
			"pekerjaan": "Mahasiswa",
			"alasan":    "Migrasi dari Dev Ops ke Golang",
		},

		map[string]string{
			"nama":      "Rafly Andreansyah",
			"alamat":    "Malang",
			"pekerjaan": "Mahasiswa",
			"alasan":    "Mencoba Golang",
		},

		map[string]string{
			"nama":      "Taufiq Hidayah",
			"alamat":    "Kalimantan",
			"pekerjaan": "Mahasiswa",
			"alasan":    "Penasaran sama Microservices Golang",
		},

		map[string]string{
			"nama":      "Melvita Sari",
			"alamat":    "Aceh",
			"pekerjaan": "Mahasiswa",
			"alasan":    "Ingin menjadi Back End Dev.",
		},
	}

	valueOfARGS, _ := strconv.Atoi(os.Args[1])
	valueOfARGS -= 1

	if valueOfARGS >= 0 && valueOfARGS < len(students) {
		fmt.Printf("Nama siswa %v \nasal %v \npekerjaan %v \nalasan ikut course %v \n", students[valueOfARGS]["nama"], students[valueOfARGS]["alamat"], students[valueOfARGS]["pekerjaan"], students[valueOfARGS]["alasan"])
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}
