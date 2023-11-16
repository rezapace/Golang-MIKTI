package main

import (
	"fmt"
	"os"
	"strconv"
)

type Barang struct {
	Nama   string
	Harga  int
}

var daftarBarang []Barang

func main() {
	for {
		tampilkanMenuUtama()

		var pilihan string
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			inputDataBarang()
		case "2":
			beliBarang()
		case "3":
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilkanMenuUtama() {
	fmt.Println("\nMenu Utama:")
	fmt.Println("1. Input Data Barang")
	fmt.Println("2. Beli Barang")
	fmt.Println("3. Keluar")
	fmt.Print("Masukkan pilihan: ")
}

func inputDataBarang() {
	var nama string
	var harga string

	fmt.Print("Masukkan nama barang: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan harga barang: ")
	fmt.Scanln(&harga)

	hargaInt, err := strconv.Atoi(harga)
	if err != nil {
		fmt.Println("Harga harus berupa angka.")
		return
	}

	daftarBarang = append(daftarBarang, Barang{nama, hargaInt})
	fmt.Println("Barang berhasil ditambahkan!")
}

func beliBarang() {
	tampilkanDaftarBarang()

	var kodeBarang, jumlah string
	fmt.Println("Masukkan kode barang (0-", len(daftarBarang)-1, "): ")
	fmt.Scanln(&kodeBarang)

	index, err := strconv.Atoi(kodeBarang)
	if err != nil || index < 0 || index >= len(daftarBarang) {
		fmt.Println("Kode barang tidak valid.")
		return
	}

	fmt.Println("Masukkan jumlah yang dibeli: ")
	fmt.Scanln(&jumlah)

	qty, err := strconv.Atoi(jumlah)
	if err != nil || qty <= 0 {
		fmt.Println("Jumlah tidak valid.")
		return
	}

	total := daftarBarang[index].Harga * qty
	fmt.Println("Total belanjaan Anda adalah: ", total)

	cetakInvoice(daftarBarang[index], qty, total)
}

func tampilkanDaftarBarang() {
	fmt.Println("\nDaftar Barang:")
	for i, b := range daftarBarang {
		fmt.Printf("%d. %s - Rp%d\n", i, b.Nama, b.Harga)
	}
}

func cetakInvoice(barang Barang, jumlah, total int) {
	invoice := fmt.Sprintf(
		"Nama Barang: %s\nHarga Satuan: Rp%d\nJumlah: %d\nTotal: Rp%d\n",
		barang.Nama, barang.Harga, jumlah, total,
	)

	file, err := os.Create("invoice.txt")
	if err != nil {
		fmt.Println("Gagal mencetak invoice: ", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(invoice)
	if err != nil {
		fmt.Println("Gagal menulis ke file: ", err)
		return
	}

	fmt.Println("Invoice telah dicetak ke file invoice.txt")
}
