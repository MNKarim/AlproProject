package main

import (
	"fmt"
)

const NMAX int = 100


type workout struct {
	latihan string
	durasi int
	kalori int
	jadwal string
}

var dW [NMAX]workout
var jumlahData int

func main() {
	header()
}

func header() {
	opsi := []string{"1. Kelola Workout", "2. Rekomendasi Latihan", "3. Laporan", "4. Keluar"}
	menu("Pilih opsi:", opsi, "MAIN", false)
	main_menu()
}

func main_menu() {
	var pilihan int
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		printRiwayat()
		opsi := []string{"1. Tambah Workout", "2. Urutkan Riwayat", "3. Hapus Workout", "4. Ubah Workout", "5. Cari Latihan", "6. Kembali",}
		menu("Pilih opsi:", opsi, "KELOLA", false)
		kelolaWorkout()
	case 2:
		rekomendasi()
	case 3:
		laporan()
	case 4:
		keluar()
	default:
		opsi := []string{"1. Kelola Workout", "2. Rekomendasi Latihan", "3. Laporan", "4. Keluar"}
		menu("Masukkan angka sesuai opsi!", opsi, "MAIN", false)
		main_menu()
	}
}

func menu(teks string, opsi []string, tipeMenu string, hide bool) {
if tipeMenu == "MAIN" {
		fmt.Printf("%25s┌─────────────────────────────────────────────────────────────────────────┐\n", "")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s│%50s%-23s│\n", "", "Aplikasi Manajemen Workout", "")
		fmt.Printf("%25s│%73s│\n", "", "")
	} else if tipeMenu == "KELOLA" {
		fmt.Printf("%25s┌─────────────────────────────────────────────────────────────────────────┐\n", "")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s│%43s%-30s│\n", "", "Kelola Riwayat", "")
		fmt.Printf("%25s│%73s│\n", "", "")
	}
		fmt.Printf("%25s├─────────────────────────────────────────────────────────────────────────┤\n", "")
	fmt.Printf("%25s│%73s│\n", "", "")
	for _, option := range opsi {
		if option != "" {
			fmt.Printf("%25s│%4s%-69s│\n", "", "", option)
		}
	}
	fmt.Printf("%25s│%73s│\n", "", "")
	if teks != "" {
		fmt.Printf("%25s│  %-71s│\n", "", teks)
		fmt.Printf("%25s│%73s│\n", "", "")
	}

	fmt.Printf("%25s└─────────────────────────────────────────────────────────────────────────┘\n", "")
	if !hide {
		fmt.Printf("%25sPILIH OPSI: ", "")
	}
}

func kelolaWorkout(){
	var pilihan int
	fmt.Scan(&pilihan)
	
	switch pilihan {
	case 1:
		tambahWorkout()
	case 2:
		urutkanWorkout()
	case 3:
		hapusWorkout()
	case 4:
		ubahWorkout()
	case 5:
		cariLatihan()
	case 6:
		opsi := []string{"1. Kelola Workout", "2. Rekomendasi Latihan", "3. Laporan", "4. Keluar"}
		menu("Pilih opsi: ", opsi, "MAIN", false)
		main_menu()
	default:
		printRiwayat()
		opsi := []string{"1. Tambah Workout", "2. Urutkan Riwayat", "3. Hapus Workout", "4. Ubah Workout", "5. Cari Latihan", "6. Kembali",}
		menu("Masukkan angka sesuai opsi!", opsi, "KELOLA", false)
		kelolaWorkout()
	}
}

func tambahWorkout(){
	var w workout

	fmt.Printf("%25sLatihan: ", "")
	fmt.Scan(&w.latihan)
	fmt.Printf("%25sDurasi (menit): ", "")
	fmt.Scan(&w.durasi)
	fmt.Printf("%25sKalori: ", "")
	fmt.Scan(&w.kalori)
	fmt.Printf("%25stanggal (YYYY-MM-DD): ", "")
	fmt.Scan(&w.jadwal)
	dW[jumlahData] = w
	jumlahData++
	printRiwayat()
	opsi := []string{"1. Tambah Workout", "2. Urutkan Riwayat", "3. Hapus Workout", "4. Ubah Workout", "5. Cari Latihan", "6. Kembali",}
	menu("Pilih opsi:", opsi, "KELOLA", false)
	kelolaWorkout()
}

func urutkanWorkout(){
	if jumlahData == 0 {
		printRiwayat()
		opsi := []string{"1. Tambah Workout", "2. Urutkan Riwayat", "3. Hapus Workout", "4. Ubah Workout", "5. Cari Latihan", "6. Kembali",}
		menu("belum ada data, silahkan tambah terlebih dahulu", opsi, "KELOLA", false)
		kelolaWorkout()
	}

	printRiwayat()
	opsi := []string{"1. Urutkan berdasarkan durasi", "2. Urutkan berdasarkan kalori"}
	menu("Pilih metode pengurutan: ", opsi, "KELOLA", false)

	var pilihan int
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		selectionSortDurasi()
	case 2:
		insertionSortKalori()
	default:
		printRiwayat()
		opsi := []string{"1. Tambah Workout", "2. Urutkan Riwayat", "3. Hapus Workout", "4. Ubah Workout", "5. Cari Latihan", "6. Kembali",}
		menu("Masukkan angka sesuai opsi!", opsi, "KELOLA", false)
		kelolaWorkout()
	}

	printRiwayat()
	opsi = []string{"1. Tambah Workout", "2. Urutkan Riwayat", "3. Hapus Workout", "4. Ubah Workout", "5. Cari Latihan", "6. Kembali",}
	menu("Data berhasil diurutkan", opsi, "KELOLA", false)
	kelolaWorkout()
}

func hapusWorkout(){
	if jumlahData == 0 {
		printRiwayat()
		opsi := []string{"1. Tambah Workout", "2. Urutkan Riwayat", "3. Hapus Workout", "4. Ubah Workout", "5. Cari Latihan", "6. Kembali",}
		menu("Belum ada data, silahkan tambah terlebih dahulu", opsi, "KELOLA", false)
		kelolaWorkout()
	}

	var index int
	fmt.Print("Masukkan nomor workout yang ingin dihapus (1 - ", jumlahData, "): ")
	fmt.Scan(&index)

	index--

	if index < 0 || index >= jumlahData {
		printRiwayat()
		opsi := []string{"1. Tambah Workout", "2. Urutkan Riwayat", "3. Hapus Workout", "4. Ubah Workout", "5. Cari Latihan", "6. Kembali",}
		menu("Nomor tidak valid!", opsi, "KELOLA", false)
		kelolaWorkout()
	}

	for i := index; i < jumlahData-1; i++ {
		dW[i] = dW[i+1]
	}

	jumlahData--
	printRiwayat()
	opsi := []string{"1. Tambah Workout", "2. Urutkan Riwayat", "3. Hapus Workout", "4. Ubah Workout", "5. Cari Latihan", "6. Kembali",}
	menu("workout berhasil dihapus", opsi, "KELOLA", false)
	kelolaWorkout()
}

func ubahWorkout(){
	if jumlahData == 0 {
		printRiwayat()
		opsi := []string{"1. Tambah Workout", "2. Urutkan Riwayat", "3. Hapus Workout", "4. Ubah Workout", "5. Cari Latihan", "6. Kembali",}
		menu("belum ada data, silahkan tambah terlebih dahulu", opsi, "KELOLA", false)
		kelolaWorkout()
	}

	var index int
	printRiwayat()
	fmt.Print("Masukkan nomor workout yang ingin diubah (1 - ", jumlahData, "): ")
	fmt.Scan(&index)

	index--

	if index < 0 || index >= jumlahData {
		printRiwayat()
		opsi := []string{"1. Tambah Workout", "2. Urutkan Riwayat", "3. Hapus Workout", "4. Ubah Workout", "5. Cari Latihan", "6. Kembali",}
		menu("Nomor tidak valid!", opsi, "KELOLA", false)
		kelolaWorkout()
	}

	fmt.Println("Masukkan data baru:")
	fmt.Print("Latihan: ")
	fmt.Scan(&dW[index].latihan)
	fmt.Print("Durasi (menit): ")
	fmt.Scan(&dW[index].durasi)
	fmt.Print("Kalori: ")
	fmt.Scan(&dW[index].kalori)
	fmt.Print("Tanggal (YYYY-MM-DD): ")
	fmt.Scan(&dW[index].jadwal)

	printRiwayat()
	opsi := []string{"1. Tambah Workout", "2. Urutkan Riwayat", "3. Hapus Workout", "4. Ubah Workout", "5. Cari Latihan", "6. Kembali"}
	menu("Workout berhasil diubah", opsi, "KELOLA", false)
	kelolaWorkout()
}

func cariLatihan(){
	if jumlahData == 0 {
		printRiwayat()
		opsi := []string{"1. Tambah Workout", "2. Urutkan Riwayat", "3. Hapus Workout", "4. Ubah Workout", "5. Cari Latihan", "6. Kembali",}
		menu("belum ada data, silahkan tambah terlebih dahulu", opsi, "KELOLA", false)
		kelolaWorkout()
	}
}

func rekomendasi(){
	if jumlahData == 0 {
		opsi := []string{"1. Kelola Workout", "2. Rekomendasi Latihan", "3. Laporan", "4. Keluar"}
		menu("belum ada data, silahkan tambah terlebih dahulu", opsi, "MAIN", false)
		main_menu()
	}
}

func laporan(){
	if jumlahData == 0 {
		opsi := []string{"1. Kelola Workout", "2. Rekomendasi Latihan", "3. Laporan", "4. Keluar"}
		menu("belum ada data, silahkan tambah terlebih dahulu", opsi, "MAIN", false)
		main_menu()
	}
}

func printRiwayat(){
	if jumlahData == 0 {
		fmt.Printf("%25s┌─────────────────────────────────────────────────────────────────────────┐\n", "")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s│%29s%-44s│\n", "", "", "Riwayat Workout: ")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s├─────────────────────────────────────────────────────────────────────────┤\n", "")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s│%25s%-48s│\n", "", "", "Belum ada data workout.")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s└─────────────────────────────────────────────────────────────────────────┘\n", "")
	} else {
		fmt.Printf("%25s┌─────────────────────────────────────────────────────────────────────────┐\n", "")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s│%29s%-44s│\n", "", "", "RIWAYAT WORKOUT: ")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s├─────────────────────────────────────────────────────────────────────────┤\n", "")
    	for i := 0; i < jumlahData; i++ {
        	fmt.Printf("%25s│%15s%-58d│\n", "", "Workout #", i+1)
			fmt.Printf("%25s│%35s%-38s│\n", "", "Jenis   : ", dW[i].latihan )
			fmt.Printf("%25s│%35s%-38d│\n", "", "Durasi   : ", dW[i].durasi)
        	fmt.Printf("%25s│%35s%-38d│\n", "", "Kalori   : ", dW[i].kalori)
        	fmt.Printf("%25s│%35s%-38s│\n", "", "Tanggal   : ", dW[i].jadwal )
		}
		fmt.Printf("%25s└─────────────────────────────────────────────────────────────────────────┘\n", "")
    }
}

func selectionSortDurasi() {
	for i := 0; i < jumlahData-1; i++ {
		maxIdx := i
		for j := i + 1; j < jumlahData; j++ {
			if dW[j].durasi > dW[maxIdx].durasi {
				maxIdx = j
			}
		}
		if maxIdx != i {
			dW[i], dW[maxIdx] = dW[maxIdx], dW[i]
		}
	}
}

func insertionSortKalori() {
	for i := 1; i < jumlahData; i++ {
		key := dW[i]
		j := i - 1
		for j >= 0 && dW[j].kalori < key.kalori {
			dW[j+1] = dW[j]
			j--
		}
		dW[j+1] = key
	}
}

func keluar(){

}