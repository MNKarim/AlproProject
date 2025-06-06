package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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
	load()
	header()
}

func header() {
	opsi := []string{"1. Kelola Workout", "2. Rekomendasi Latihan", "3. Laporan", "4. Keluar"}
	menu("Pilih keperluan anda", opsi, "MAIN")
	main_menu()
}

func main_menu() {
	var pilihan int
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
		menu("Pilih keperluan anda", opsi, "KELOLA")
		kelolaWorkout()
	case 2:
		rekomendasi()
	case 3:
		laporan()
	case 4:
		clear()
	default:
		opsi := []string{"1. Kelola Workout", "2. Rekomendasi Latihan", "3. Laporan", "4. Keluar"}
		menu("Masukkan angka sesuai opsi!", opsi, "MAIN")
		main_menu()
	}
}

func menu(teks string, opsi []string, tipeMenu string) {
	if tipeMenu == "MAIN" {
		clear()
		fmt.Printf("%25s┌─────────────────────────────────────────────────────────────────────────┐\n", "")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s│%50s%-23s│\n", "", "Aplikasi Manajemen Workout", "")
		fmt.Printf("%25s│%73s│\n", "", "")
	} else if tipeMenu == "KELOLA" {
		clear()
		printRiwayat()
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
	fmt.Printf("%25sPILIH OPSI: ", "")
}

func kelolaWorkout() {
	var pilihan int
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahWorkout()
	case 2:
		ubahWorkout()
	case 3:
		hapusWorkout()
	case 4:
		selectionSortDurasi()
		opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
		menu("Data berhasil diurutkan", opsi, "KELOLA")
		kelolaWorkout()
	case 5:
		insertionSortKalori()
		opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
		menu("Data berhasil diurutkan", opsi, "KELOLA")
		kelolaWorkout()
	case 6:
		cariLatihan()
	case 7:
		opsi := []string{"1. Kelola Workout", "2. Rekomendasi Latihan", "3. Laporan", "4. Keluar"}
		menu("Pilih opsi: ", opsi, "MAIN")
		main_menu()
	default:
		opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
		menu("Masukkan angka sesuai opsi!", opsi, "KELOLA")
		kelolaWorkout()
	}
}

func tambahWorkout() {
	var w workout
	fmt.Printf("%25sLatihan: ", "")
	fmt.Scan(&w.latihan)
	fmt.Printf("%25sDurasi (menit): ", "")
	fmt.Scan(&w.durasi)
	fmt.Printf("%25sKalori: ", "")
	fmt.Scan(&w.kalori)
	fmt.Printf("%25sJadwal (jj:mm): ", "")
	fmt.Scan(&w.jadwal)
	dW[jumlahData] = w
	jumlahData++
	opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
	menu("Pilih opsi:", opsi, "KELOLA")
	kelolaWorkout()
}

func ubahWorkout() {
	if jumlahData == 0 {
		opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
		menu("belum ada data, silahkan tambah terlebih dahulu", opsi, "KELOLA")
		kelolaWorkout()
	}
	var index int
	fmt.Printf("%25sMasukkan nomor workout yang ingin diubah (1 -  %d): ", "", jumlahData)
	fmt.Scan(&index)
	index--
	if index < 0 || index >= jumlahData {
		opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
		menu("Nomor tidak valid!", opsi, "KELOLA")
		kelolaWorkout()
	}
	fmt.Printf("%25sMasukkan data baru:\n", "")
	fmt.Printf("%25sLatihan: ", "")
	fmt.Scan(&dW[index].latihan)
	fmt.Printf("%25sDurasi (menit): ", "")
	fmt.Scan(&dW[index].durasi)
	fmt.Printf("%25sKalori: ", "")
	fmt.Scan(&dW[index].kalori)
	fmt.Printf("%25sJadwal (jj:mm): ", "")
	fmt.Scan(&dW[index].jadwal)
	opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
	menu("Workout berhasil diubah", opsi, "KELOLA")
	kelolaWorkout()
}

func hapusWorkout() {
	if jumlahData == 0 {
		opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
		menu("Belum ada data, silahkan tambah terlebih dahulu", opsi, "KELOLA")
		kelolaWorkout()
	}
	var index int
	fmt.Printf("%25sMasukkan nomor workout yang ingin dihapus (1 -  %d): ", "", jumlahData)
	fmt.Scan(&index)
	index--
	if index < 0 || index >= jumlahData {
		opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
		menu("Nomor tidak valid!", opsi, "KELOLA")
		kelolaWorkout()
	}
	for i := index; i < jumlahData-1; i++ {
		dW[i] = dW[i+1]
	}
	jumlahData--
	opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
	menu("workout berhasil dihapus", opsi, "KELOLA")
	kelolaWorkout()
}

func selectionSortDurasi() {
	if jumlahData == 0 {
		opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
		menu("belum ada data, silahkan tambah terlebih dahulu", opsi, "KELOLA")
		kelolaWorkout()
	}
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
	if jumlahData == 0 {
		opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
		menu("belum ada data, silahkan tambah terlebih dahulu", opsi, "KELOLA")
		kelolaWorkout()
	}
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

func cariLatihan() {
	if jumlahData == 0 {
		opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
		menu("belum ada data, silahkan tambah terlebih dahulu", opsi, "KELOLA")
		kelolaWorkout()
	}
	var keyword string
	fmt.Printf("%25sMasukkan nama latihan yang ingin dicari: ", "")
	fmt.Scan(&keyword)
	hasilIndex := sequentialSearchLatihan(keyword)
	printCari()
	if hasilIndex[0] == -1 {
		opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
		menu("Latihan tidak ditemukan", opsi, "KELOLA")
		kelolaWorkout()
	} else {
		fmt.Printf("%25s│  %-21s  %-10s  %-15s  %-19s│\n", "", "Workout", "Durasi", "Kalori", "Jadwal")
		i := 0
		for hasilIndex[i] != -1 {
			idx := hasilIndex[i]
			durasiStr := fmt.Sprintf("%d menit", dW[idx].durasi)
			kaloriStr := fmt.Sprintf("%d kkal", dW[idx].kalori)
			fmt.Printf("%25s│ %2d.  %-17s  %-10s  %-15s  %-19s│\n", "", i+1, dW[idx].latihan, durasiStr, kaloriStr, dW[idx].jadwal)
			i++
		}
	}
	fmt.Printf("%25s└─────────────────────────────────────────────────────────────────────────┘\n", "")
	fmt.Printf("%25sTekan apa saja untuk kembali: ", "")
	var pilihan int
	fmt.Scan(&pilihan)
	opsi := []string{"1. Tambah Workout", "2. Ubah Workout", "3. Hapus Workout", "4. Urutkan Berdasarkan Durasi", "5. Urutkan Berdasarkan Kalori", "6. Cari Latihan", "7. Kembali",}
	menu("belum ada data, silahkan tambah terlebih dahulu", opsi, "KELOLA")
	kelolaWorkout()
}

func sequentialSearchLatihan(keyword string) [NMAX]int {
	var hasil [NMAX]int
	jumlahHasil := 0
	for i := 0; i < jumlahData; i++ {
		if strings.ToLower(dW[i].latihan) == strings.ToLower(keyword) {
			hasil[jumlahHasil] = i
			jumlahHasil++
		}
	}
	if jumlahHasil < 100 {
		hasil[jumlahHasil] = -1
	}
	return hasil
}

func printCari() {
	clear()
	if jumlahData == 0 {
		fmt.Printf("%25s┌─────────────────────────────────────────────────────────────────────────┐\n", "")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s│%29s%-44s│\n", "", "", "CARI WORKOUT: ")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s├─────────────────────────────────────────────────────────────────────────┤\n", "")
		fmt.Printf("%25s│%25s%-48s│\n", "", "", "Belum ada data workout.")
		fmt.Printf("%25s└─────────────────────────────────────────────────────────────────────────┘\n", "")
	} else {
		fmt.Printf("%25s┌─────────────────────────────────────────────────────────────────────────┐\n", "")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s│%29s%-44s│\n", "", "", "CARI WORKOUT: ")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s├─────────────────────────────────────────────────────────────────────────┤\n", "")
	}
}

func rekomendasi() {
	var chest, legs, core, fullbody, back int
    for i := 0; i < jumlahData; i++ {
        bagian := bagianTubuh(dW[i].latihan)
        if bagian == "Chest" {
            chest++
        } else if bagian == "Legs" {
            legs++
        } else if bagian == "Core" {
            core++
        } else if bagian == "Full Body" {
            fullbody++
        } else if bagian == "Back" {
			back++
		}
    }
    max := chest
    bagianDominan := "Chest"
    if legs > max {
        max = legs
        bagianDominan = "Legs"
    }
    if core > max {
        max = core
        bagianDominan = "Core"
    }
    if fullbody > max {
        max = fullbody
        bagianDominan = "Full Body"
    }
	if back> max {
        max = back
        bagianDominan = "Back"
    }
    printRekomendasi(bagianDominan)
}

func bagianTubuh(nama string) string {
    if nama == "Push_Up" || nama == "Bench_Press" {
        return "Chest"
    } else if nama == "Box_Jump" || nama == "Deadlift" || nama == "High_Knees" || nama == "Jogging" || nama == "Bersepeda" || nama == "Jump Rope" || nama == "Squat" || nama == "Lunges"{
        return "Legs"
    } else if nama == "Plank" || nama == "Sit_Up" || nama == "Burpees" || nama == "Side_Plank" || nama == "Crunch"{
        return "Core"
    } else if nama == "Pull_Up" || nama == "Mountain_Climber" {
        return "Back"
    } else if nama == "Yoga" || nama == "Jumping_Jack" {
        return "Full Body"
    } else {
        return "Unknown"
    }
}

func printRekomendasi(bagian string) {
    clear()
	if jumlahData == 0 {
		fmt.Printf("%25s┌─────────────────────────────────────────────────────────────────────────┐\n", "")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s│%26s%-47s│\n", "", "", "REKOMENDASI WORKOUT: ")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s├─────────────────────────────────────────────────────────────────────────┤\n", "")
		fmt.Printf("%25s│%25s%-48s│\n", "", "", "Belum ada data workout.")
		fmt.Printf("%25s└─────────────────────────────────────────────────────────────────────────┘\n", "")
		fmt.Printf("%25sTekan apa saja untuk kembali: ", "")
		var pilihan int
		fmt.Scan(&pilihan)
		opsi := []string{"1. Kelola Workout", "2. Rekomendasi Latihan", "3. Laporan", "4. Keluar"}
		menu("Pilih keperluan anda", opsi, "MAIN")
		main_menu()
	} else {
		fmt.Printf("%25s┌─────────────────────────────────────────────────────────────────────────┐\n", "")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s│%26s%-47s│\n", "", "", "REKOMENDASI WORKOUT: ")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s├─────────────────────────────────────────────────────────────────────────┤\n", "")
	}
	if bagian == "Legs" || bagian == "Core" || bagian == "Back" {
		fmt.Printf("%25s│ Bagian tubuh yang paling sering dilatih adalah: %s %-18s │\n", "", bagian, "")
	} else if bagian == "Chest" {
		fmt.Printf("%25s│ Bagian tubuh yang paling sering dilatih adalah: %s %-17s │\n", "", bagian, "")
	} else {
		fmt.Printf("%25s│ Bagian tubuh yang paling sering dilatih adalah: %s %-13s │\n", "", bagian, "")
	}
	fmt.Printf("%25s│ %s %41s │\n", "","Rekomendasi latihan tambahan:", "")
    if bagian == "Chest" {
		fmt.Printf("%25s│ %s %-53s │\n", "","- Incline Push Up", "")
		fmt.Printf("%25s│ %s %-53s │\n", "","- Diamond Push Up", "")
		fmt.Printf("%25s│ %s %-58s │\n", "","- Chest Dips", "")
    } else if bagian == "Legs" {
		fmt.Printf("%25s│ %s %-59s │\n", "","- Steps_Ups", "")
		fmt.Printf("%25s│ %s %-57s │\n", "","- Calf_Raises", "")
		fmt.Printf("%25s│ %s %-60s │\n", "","- Wall_Sit", "")
    } else if bagian == "Core" {
		fmt.Printf("%25s│ %s %-58s │\n", "","- Side_Plank", "")
		fmt.Printf("%25s│ %s %-55s │\n", "","- Russian_Twist", "")
		fmt.Printf("%25s│ %s %-58s │\n", "","- Leg_Raises", "")
    } else if bagian == "Full Body" {
		fmt.Printf("%25s│ %s %-55s │\n", "","- Jumping Jacks", "")
		fmt.Printf("%25s│ %s %-51s │\n", "","- Mountain Climbers", "")
		fmt.Printf("%25s│ %s %-58s │\n", "","- Bear Crawl", "")
	}else if bagian == "Back" {
		fmt.Printf("%25s│ %s %-56s │\n", "","- Lat_Pulldown", "")
		fmt.Printf("%25s│ %s %-56s │\n", "","- Inverted_Row", "")
		fmt.Printf("%25s│ %s %-60s │\n", "","- Bird_Dog", "")
    } else {
        fmt.Println("- Tidak ada rekomendasi yang tersedia.")
    }
	fmt.Printf("%25s└─────────────────────────────────────────────────────────────────────────┘\n", "")
	fmt.Printf("%25sTekan apa saja untuk kembali: ", "")
	var pilihan int
	fmt.Scan(&pilihan)
	opsi := []string{"1. Kelola Workout", "2. Rekomendasi Latihan", "3. Laporan", "4. Keluar"}
	menu("Pilih keperluan anda", opsi, "MAIN")
	main_menu()
}

func laporan() {
	printLaporan()
	var totalDurasi, totalKalori int
	maxDurasi := dW[0].durasi
	maxKalori := dW[0].kalori
	idxMaxDurasi := 0
	idxMaxKalori := 0
	for i := 0; i < jumlahData; i++ {
		totalDurasi += dW[i].durasi
		totalKalori += dW[i].kalori

		if dW[i].durasi > maxDurasi {
			maxDurasi = dW[i].durasi
			idxMaxDurasi = i
		}

		if dW[i].kalori > maxKalori {
			maxKalori = dW[i].kalori
			idxMaxKalori = i
		}
	}
	rataDurasi := float64(totalDurasi) / float64(jumlahData)
	rataKalori := float64(totalKalori) / float64(jumlahData)
	printLaporan()
	fmt.Printf("%25s│  %-21s  %-10d  %-15s  %-19s│\n", "", "Jumlah data workout :", jumlahData, "", "")
	fmt.Printf("%25s│  %-21s  %-5d menit %-15s  %-19s│\n", "", "Total durasi :", totalDurasi, "", "")
	fmt.Printf("%25s│  %-21s  %-10d  %-15s  %-19s│\n", "", "Total kalori :", totalKalori, "", "")
	fmt.Printf("%25s│  %-21s  %.2f menit %-15s  %-19s│\n", "", "Rata-rata durasi :", rataDurasi, "", "")
	fmt.Printf("%25s│  %-21s  %.2f kkal %-15s  %-19s│\n", "", "Rata-rata kalori :", rataKalori, "", "")
	fmt.Printf("%25s│%73s│\n", "", "")
	fmt.Printf("%25s│  %-21s  %-10s  %-15s  %-9s│\n", "", "Workout dengan durasi terlama :", "", "", "")
	fmt.Printf("%25s│  - %s (%d menit, %d kalori, %s) %-27s│\n", "", dW[idxMaxDurasi].latihan, dW[idxMaxDurasi].durasi, dW[idxMaxDurasi].kalori, dW[idxMaxDurasi].jadwal, "")
	fmt.Printf("%25s│  %-21s  %-10s  %-15s  %-7s│\n", "", "Workout dengan kalori terbanyak :", "", "", "")
	fmt.Printf("%25s│  - %s (%d menit, %d kalori, %s) %-22s│\n", "", dW[idxMaxKalori].latihan, dW[idxMaxKalori].durasi, dW[idxMaxKalori].kalori, dW[idxMaxKalori].jadwal, "")
	fmt.Printf("%25s└─────────────────────────────────────────────────────────────────────────┘\n", "")
	fmt.Printf("%25sTekan apa saja untuk kembali: ", "")
	var pilihan int
	fmt.Scan(&pilihan)
	opsi := []string{"1. Kelola Workout", "2. Rekomendasi Latihan", "3. Laporan", "4. Keluar"}
	menu("Pilih keperluan anda", opsi, "MAIN")
	main_menu()
}

func printLaporan() {
	clear()
	if jumlahData == 0 {
		fmt.Printf("%25s┌─────────────────────────────────────────────────────────────────────────┐\n", "")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s│%28s%-45s│\n", "", "", "LAPORAN WORKOUT: ")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s├─────────────────────────────────────────────────────────────────────────┤\n", "")
		fmt.Printf("%25s│%25s%-48s│\n", "", "", "Belum ada data workout.")
		fmt.Printf("%25s└─────────────────────────────────────────────────────────────────────────┘\n", "")
		fmt.Printf("%25sTekan apa saja untuk kembali: ", "")
		var pilihan int
		fmt.Scan(&pilihan)
		opsi := []string{"1. Kelola Workout", "2. Rekomendasi Latihan", "3. Laporan", "4. Keluar"}
		menu("Pilih keperluan anda", opsi, "MAIN")
		main_menu()
	} else {
		fmt.Printf("%25s┌─────────────────────────────────────────────────────────────────────────┐\n", "")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s│%28s%-45s│\n", "", "", "LAPORAN WORKOUT: ")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s├─────────────────────────────────────────────────────────────────────────┤\n", "")
	}
}

func printRiwayat() {
	if jumlahData == 0 {
		fmt.Printf("%25s┌─────────────────────────────────────────────────────────────────────────┐\n", "")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s│%29s%-44s│\n", "", "", "RIWAYAT WORKOUT: ")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s├─────────────────────────────────────────────────────────────────────────┤\n", "")
		fmt.Printf("%25s│%25s%-48s│\n", "", "", "Belum ada data workout.")
	} else {
		fmt.Printf("%25s┌─────────────────────────────────────────────────────────────────────────┐\n", "")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s│%29s%-44s│\n", "", "", "RIWAYAT WORKOUT: ")
		fmt.Printf("%25s│%73s│\n", "", "")
		fmt.Printf("%25s├─────────────────────────────────────────────────────────────────────────┤\n", "")
		fmt.Printf("%25s│  %-21s  %-10s  %-15s  %-19s│\n", "", "Workout", "Durasi", "Kalori", "Jadwal")
		for i := 0; i < jumlahData; i++ {
			durasiStr := fmt.Sprintf("%d menit", dW[i].durasi)
			kaloriStr := fmt.Sprintf("%d kkal", dW[i].kalori)
			fmt.Printf("%25s│ %2d.  %-17s  %-10s  %-15s  %-19s│\n", "", i+1,dW[i].latihan, durasiStr, kaloriStr, dW[i].jadwal)
		}
	}
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func load() {
	dW[0] = workout{"Push_Up", 30, 150, "Senin, 08:00"}
	dW[1] = workout{"Jogging", 45, 300, "Senin, 08:35"}
	dW[2] = workout{"Plank", 10, 50, "Senin, 09:25"}
	dW[3] = workout{"Squat", 20, 120, "Senin, 09:40"}
	dW[4] = workout{"Yoga", 60, 200, "Senin, 10:05"}
	dW[5] = workout{"Bersepeda", 50, 400, "Senin, 15:20"}
	dW[6] = workout{"Lunges", 25, 140, "Senin, 16:15"}
	dW[7] = workout{"Burpees", 15, 180, "Senin, 16:55"}
	dW[8] = workout{"Sit_Up", 35, 160, "Senin, 17:15"}
	dW[9] = workout{"Jump Rope", 20, 220, "Senin, 17:55"}
	dW[10] = workout{"Pull_Up", 10, 100, "Sabtu, 06:30"}
	dW[11] = workout{"Mountain_Climber", 15, 130, "Sabtu, 06:50"}
	dW[12] = workout{"High_Knees", 20, 150, "Sabtu, 07:10"}
	dW[13] = workout{"Side Plank", 12, 60, "Sabtu, 07:30"}
	dW[14] = workout{"Jumping_Jack", 25, 180, "Sabtu, 07:50"}
	dW[15] = workout{"Deadlift", 30, 250, "Sabtu, 08:10"}
	dW[16] = workout{"Bench_Press", 25, 220, "Sabtu, 08:30"}
	dW[17] = workout{"Crunch", 20, 90, "Sabtu, 08:50"}
	dW[18] = workout{"Stretching", 15, 70, "Sabtu, 09:10"}
	dW[19] = workout{"Box_Jump", 20, 190, "Sabtu, 09:30"}
	jumlahData = 20
}
