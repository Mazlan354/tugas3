package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tugas3/model"
	"tugas3/node"
)

func Insert() {
	var kota, nama, notelp, email, jalan string
	var id, nomer int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Pegawai: ")
	fmt.Scan(&id)
	reader.ReadString('\n')

	fmt.Print("Masukkan Nama Pegawai: ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan Jalan: ")
	jalan, _ = reader.ReadString('\n')
	jalan = strings.TrimSpace(jalan)

	fmt.Print("Masukkan Nomor Rumah: ")
	fmt.Scan(&nomer)
	reader.ReadString('\n')

	fmt.Print("Masukkan Kota: ")
	kota, _ = reader.ReadString('\n')
	kota = strings.TrimSpace(kota)

	fmt.Print("Masukkan Nomor Telepon : ")
	notelp, _ = reader.ReadString('\n')
	notelp = strings.TrimSpace(notelp)

	fmt.Print("Masukkan Email: ")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	// create new pegawai
	pegawai := node.Pegawai{
		ID:     id,
		Nama:   nama,
		Alamat: node.Address{Jalan: jalan, Kota: kota, Nomer: nomer},
		NoTelp: notelp,
		Email:  email,
	}

	// insert to DaftarPegawai
	cek := model.CreatePegawai(pegawai)
	if cek {
		fmt.Println("== Pegawai Berhasil ditambahkan ==")
	} else {
		fmt.Println("== Pegawai gagal ditambahkan ==")
	}
	fmt.Println()
}

func Update() {
	var id, nomer int
	var nama, kota, notelp, email, jalan string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Pegawai yang ingin diupdate: ")
	fmt.Scan(&id)
	reader.ReadString('\n')

	fmt.Print("Masukkan Nama Pegawai: ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan Kota: ")
	kota, _ = reader.ReadString('\n')
	kota = strings.TrimSpace(kota)

	fmt.Print("Masukkan Jalan: ")
	jalan, _ = reader.ReadString('\n')
	jalan = strings.TrimSpace(jalan)

	fmt.Print("Masukkan Nomer Rumah: ")
	fmt.Scan(&nomer)
	reader.ReadString('\n')

	fmt.Print("Masukkan Nomor Telepon : ")
	notelp, _ = reader.ReadString('\n')
	notelp = strings.TrimSpace(notelp)

	fmt.Print("Masukkan Email: ")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)

	// Create New Pegawai
	pegawai := node.Pegawai{
		ID:   id,
		Nama: nama,
		Alamat: node.Address{
			Jalan: jalan,
			Kota:  kota,
			Nomer: nomer},
		NoTelp: notelp,
		Email:  email,
	}

	// Update Pegawai
	cek := model.UpdatePegawai(pegawai, id)
	if cek {
		fmt.Println("== Pegawai Berhasil DiUpdate == ")
	} else {
		fmt.Println("Pegawai gagal diUpdate")
	}
	fmt.Println()
}

func Views() {
	fmt.Println("Daftar Pegawai")
	for i, emp := range model.ReadPegawai() {
		fmt.Println("--- Pegawai ke -", i+1, "---")
		fmt.Println("ID Pegawai: \t", emp.ID)
		fmt.Println("Nama Pegawai: \t", emp.Nama)
		fmt.Println("Alamat\t\t: ", emp.Alamat.Jalan, emp.Alamat.Nomer, emp.Alamat.Kota)
		fmt.Println("No. Telepon\t: ", emp.NoTelp)
		fmt.Println("Email\t\t: ", emp.Email)
		fmt.Println()
	}
}

func Delete() {
	var id int
	fmt.Print("Masukkan ID Pegawai yang akan dihapus: ")
	fmt.Scan(&id)

	cek := model.DeletePegawai(id)
	if cek {
		fmt.Println("== Pegawai Berhasil Dihapus ==")
	} else {
		fmt.Println("Pegawai gagal dihapus")
	}
	fmt.Println()
}
