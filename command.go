package main

import (
	"errors"
	"fmt"
)

var (
	Nama string = "Bryan Raihan 'Ilman"
	NPM string = "2106704351"
	DaftarMahasiswa []Tool
	DaftarMatkul []Tool
)

func AddMahasiswa(nama, id string) (string, error) {
	var successMsg string = ""
	var errMsg error = nil

	for _, mahasiswa := range DaftarMahasiswa {
		if mahasiswa.Get() == id {
			errMsg = errors.New(fmt.Sprintf("mahasiswa %v sudah ada di daftar mahasiswa", id))
			break
		}
	}

	if errMsg == nil {
		successMsg = fmt.Sprintf("berhasil menambahkan mahasiswa %v ke daftar mahasiswa", id)
		DaftarMahasiswa = append(DaftarMahasiswa, Mahasiswa{nama: nama, id: id, ambilMatkul: []string{}})
	}

	return successMsg, errMsg
}

func DropMahasiswa(id string) (string, error) {
	//TODO implement here
	panic("fix me")
}

func AddMatkul(nama, kode string, kapasitas int) (string, error) {
	var successMsg string = ""
	var errMsg error = nil

	for _, matkul := range DaftarMatkul {
		if matkul.Get() == kode {
			errMsg = errors.New(fmt.Sprintf("matkul %v sudah ada di daftar mata kuliah", kode))
			break
		}
	}

	if errMsg == nil {
		successMsg = fmt.Sprintf("berhasil menambahkan matkul %v ke daftar mata kuliah", kode)
		DaftarMatkul = append(DaftarMatkul, Matkul{nama: nama, kode: kode, kapasitas: kapasitas, diambilMahasiswa: []string{}})
	}

	return successMsg, errMsg
}

func DropMatkul(kode string) (string, error) {
	//TODO implement here
	panic("fix me")
}

func AddMahasiswaMatkul(id, kode string) (string, error) {
	//TODO implement here
	panic("fix me")
}

func DropMahasiswaMatkul(id, kode string) (string, error) {
	//TODO implement here
	panic("fix me")
}

func PrintMessage(successMsg string, errMsg error) string {
	if errMsg != nil {
		if e, ok := errMsg.(error); ok {
            return fmt.Sprintf("[FAILED] %v", e.Error())
        }
	} else {
		return fmt.Sprintf("[SUCCESS] %v", successMsg)
	}

	return ""
}
