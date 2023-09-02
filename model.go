package main

import (
	"errors"
	"fmt"
)

type Tool interface {
	Add(any) error
	Drop(any) error
	Get() any
	GetArray() any
}

type Mahasiswa struct {
	nama        string
	id          string
	ambilMatkul []string
}

func (mh *Mahasiswa) Add(kodeMatkul any) error {
	var errMsg error = nil

	for _, urutan := range mh.ambilMatkul {
		if urutan == kodeMatkul {
			errMsg = errors.New(fmt.Sprintf("mahasiswa %v sudah mengambil matkul %v", mh.id, kodeMatkul))
			break
		}
	}

	if errMsg == nil {
		mh.ambilMatkul = append(mh.ambilMatkul, kodeMatkul.(string))
	}

	return errMsg
}

func (mh *Mahasiswa) Drop(kodeMatkul any) error {
	var success bool = false
	var errMsg error = nil

	for index, urutan := range mh.ambilMatkul {
		if urutan == kodeMatkul {
			mh.ambilMatkul = append(mh.ambilMatkul[:index], mh.ambilMatkul[index + 1:]...)
			success = true
			break
		}
	}

	if !success {
		errMsg = errors.New(fmt.Sprintf("matkul %v tidak diambil oleh mahasiswa %v", kodeMatkul, mh.id))
	}

	return errMsg
}

func (mh *Mahasiswa) Get() any {
	return mh.id
}

func (mh *Mahasiswa) GetArray() any {
	return mh.ambilMatkul
}

type Matkul struct {
	nama             string
	kode             string
	kapasitas        int
	diambilMahasiswa []string
}

func (mt *Matkul) Add(idMahasiswa any) error {
	var errMsg error = nil

	if len(mt.diambilMahasiswa) == mt.kapasitas {
		errMsg = errors.New(fmt.Sprintf("kapasitas mahasiswa untuk mengambil matkul %v sudah penuh", mt.kode))
	} else {
		for _, urutan := range mt.diambilMahasiswa {
			if urutan == idMahasiswa {
				errMsg = errors.New(fmt.Sprintf("mahasiswa %v sudah mengambil matkul %v", idMahasiswa, mt.kode))
				break
			}
		}
	}

	if errMsg == nil {
		mt.diambilMahasiswa = append(mt.diambilMahasiswa, idMahasiswa.(string))
	}

	return errMsg
}

func (mt *Matkul) Drop(idMahasiswa any) error {
	var success bool = false
	var errMsg error = nil

	for index, urutan := range mt.diambilMahasiswa {
		if urutan == idMahasiswa {
			mt.diambilMahasiswa = append(mt.diambilMahasiswa[:index], mt.diambilMahasiswa[index + 1:]...)
			success = true
			break
		}
	}

	if !success {
		errMsg = errors.New(fmt.Sprintf("mahasiswa %v tidak mengambil matkul %v", idMahasiswa, mt.kode))
	}

	return errMsg
}

func (mt *Matkul) Get() any {
	return mt.kode
}

func (mt *Matkul) GetArray() any {
	return mt.diambilMahasiswa
}
