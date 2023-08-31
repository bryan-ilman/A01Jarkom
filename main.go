package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	DaftarMahasiswa = []Tool{}
	DaftarMatkul = []Tool{}
	run()
}

func run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[CRASH] ", r)
		}
	}()

	fmt.Printf("Nama: %s, NPM: %s\n", Nama, NPM)
	fmt.Println("========================================")
	fmt.Println("Selamat datang di SIKA-NG")
	fmt.Println("Silakan masukkan perintah anda di bawah")
	fmt.Println("========================================")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		line := scanner.Text()
		err := scanner.Err()
		if err != nil {
			fmt.Println("[CRASH] ", err.Error())
		}

		spl := strings.Split(line, " ")
		switch spl[0] {
		case "ADD_MAHASISWA":
			fmt.Println(PrintMessage(AddMahasiswa(spl[1], spl[2])))
		case "DROP_MAHASISWA":
			fmt.Println(PrintMessage(DropMahasiswa(spl[1])))
		case "ADD_MATKUL":
			c, err := strconv.Atoi(spl[3])
			if err != nil {
				fmt.Println("[CRASH] ", err.Error())
			} else {
				fmt.Println(PrintMessage(AddMatkul(spl[1], spl[2], c)))
			}
		case "DROP_MATKUL":
			fmt.Println(PrintMessage(DropMatkul(spl[1])))
		case "ADD_MAHASISWA_MATKUL":
			fmt.Println(PrintMessage(AddMahasiswaMatkul(spl[1], spl[2])))
		case "DROP_MAHASISWA_MATKUL":
			fmt.Println(PrintMessage(DropMahasiswaMatkul(spl[1], spl[2])))
		case "EXIT":
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}
