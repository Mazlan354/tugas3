package model

import "tugas3/node"

var DaftarPegawai [10]node.Pegawai
var Kuota int = 10
var Index int = 0

func CreatePegawai(emp node.Pegawai) bool {
	if Index < Kuota {
		DaftarPegawai[Index] = emp
		Index = Index + 1
		return true
	}
	return false
}

func ReadPegawai() []node.Pegawai {
	return DaftarPegawai[0:Index]
}

func UpdatePegawai(emp node.Pegawai, id int) bool {
	for i := 0; i < Index; i++ {
		if DaftarPegawai[i].ID == id {
			DaftarPegawai[i] = emp
			return true
		}
	}
	return false
}

func DeletePegawai(id int) bool {
	for i := 0; i < Index; i++ {
		if DaftarPegawai[i].ID == id {
			DaftarPegawai[i] = DaftarPegawai[Index-1]
			Index = Index - 1
			return true
		}
	}
	return false
}
