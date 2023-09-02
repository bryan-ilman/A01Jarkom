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
		DaftarMahasiswa = append(DaftarMahasiswa, &Mahasiswa{nama: nama, id: id, ambilMatkul: []string{}})
	}

	return successMsg, errMsg
}

func DropMahasiswa(id string) (string, error) {
	var successMsg string = ""
	var errMsg error = nil

	for index, mahasiswa := range DaftarMahasiswa {
		if mahasiswa.Get() == id {
			if len(mahasiswa.GetArray().([]string)) == 0 {
				successMsg = fmt.Sprintf("berhasil menghapus mahasiswa %v dari daftar mahasiswa", id)
				DaftarMahasiswa = append(DaftarMahasiswa[:index], DaftarMahasiswa[index + 1:]...)
			} else {
				errMsg = errors.New(fmt.Sprintf("terdapat matkul yang diambil oleh mahasiswa %v", id))
			}
			break
		}
	}

	if successMsg == "" && errMsg == nil {
		errMsg = errors.New(fmt.Sprintf("mahasiswa %v tidak ada di daftar mahasiswa", id))
	}

	return successMsg, errMsg
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
		DaftarMatkul = append(DaftarMatkul, &Matkul{nama: nama, kode: kode, kapasitas: kapasitas, diambilMahasiswa: []string{}})
	}

	return successMsg, errMsg
}

func DropMatkul(kode string) (string, error) {
	var successMsg string = ""
	var errMsg error = nil

	for index, matkul := range DaftarMatkul {
		if matkul.Get() == kode {
			if len(matkul.GetArray().([]string)) == 0 {
				successMsg = fmt.Sprintf("berhasil menghapus matkul %v dari daftar mata kuliah", kode)
				DaftarMatkul = append(DaftarMatkul[:index], DaftarMatkul[index + 1:]...)
			} else {
				errMsg = errors.New(fmt.Sprintf("terdapat mahasiswa yang mengambil matkul %v", kode))
			}
			break
		}
	}

	if successMsg == "" && errMsg == nil {
		errMsg = errors.New(fmt.Sprintf("matkul %v tidak ada di daftar mata kuliah", kode))
	}

	return successMsg, errMsg
}

func AddMahasiswaMatkul(id, kode string) (string, error) {
	var successMsg string = ""
	var errMsg error = nil

	for _, mahasiswa := range DaftarMahasiswa {
		if mahasiswa.Get() == id {
			errMsg = errors.New(fmt.Sprintf("mahasiswa %v sudah ada di daftar mahasiswa", id))
			break
		}
	}

	for _, matkul := range DaftarMatkul {
		if matkul.Get() == kode {
			errMsg = errors.New(fmt.Sprintf("matkul %v sudah ada di daftar mata kuliah", kode))
			break
		}
	}

	return successMsg, errMsg
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
