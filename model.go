package main

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
	//TODO implement here
	panic("fix me")
}

func (mh *Mahasiswa) Drop(kodeMatkul any) error {
	//TODO implement here
	panic("fix me")
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
	//TODO implement here
	panic("fix me")
}

func (mt *Matkul) Drop(idMahasiswa any) error {
	//TODO implement here
	panic("fix me")
}

func (mt *Matkul) Get() any {
	return mt.kode
}

func (mt *Matkul) GetArray() any {
	return mt.diambilMahasiswa
}
