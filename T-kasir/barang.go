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

var daftarBarang = []Barang{
	{"Buku", 10000},
	{"Pulpen", 2000},
	{"Penggaris", 5000},
}

func main() {
	tampilkanDaftarBarang()

	var kodeBarang, jumlah string
	fmt.Println("Masukkan kode barang (0-2): ")
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
	fmt.Println("Daftar Barang:")
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

