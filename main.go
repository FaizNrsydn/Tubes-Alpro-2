package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tim struct {
	Nama  string
	Peran string
}

type Startup struct {
	Nama, BidangUsaha string
	TahunBerdiri      int
	TotalPendanaan    float64
	Tim               []Tim
}

var data []Startup

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("----------------------------------------------------")
		fmt.Println("\t\tMENU UTAMA")
		fmt.Println("----------------------------------------------------")
		fmt.Println("1. Tambah data Startup")
		fmt.Println("2. Ubah data Startup")
		fmt.Println("3. Hapus data Startup")
		fmt.Println("4. Tambah daftar anggota tim dan peran dalam Startup")
		fmt.Println("5. Cari data Startup")
		fmt.Println("6. Urutkan data Startup")
		fmt.Println("7. Tampilkan data Startup")
		fmt.Println("8. Laporan jumlah startup per kategori bidang usaha")
		fmt.Println("9. Keluar")
		fmt.Println("----------------------------------------------------")
		fmt.Print("Pilih Menu(1-9): ")
		pilihan := readLineInt(reader)
		fmt.Println("----------------------------------------------------")
		switch pilihan {
		case 1:
			tambahStartup(reader)
		case 2:
			ubahStartup(reader)
		case 3:
			hapusStartup(reader)
		case 4:
			tambahAnggotaTim(reader)
		case 5:
			cari(reader)
		case 6:
			urut(reader)
		case 7:
			tampilkanStartup()
		case 8:
			laporanBidang()
		case 9:
			fmt.Println()
			fmt.Println("-----x-----x-----x-----x-----x-----x-----x-----x----")
			fmt.Println("\t\t    KELUAR")
			fmt.Println("\t\tPROGRAM SELESAI")
			fmt.Println("-----x-----x-----x-----x-----x-----x-----x-----x----")
			return
		}
	}
}

func readLineStr(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
func readLineInt(reader *bufio.Reader) int {
	inputInt, _ := reader.ReadString('\n')
	input, _ := strconv.Atoi(strings.TrimSpace(inputInt))
	return input
}
func readLineFloat(reader *bufio.Reader) float64 {
	inputF, _ := reader.ReadString('\n')
	input, _ := strconv.ParseFloat(strings.TrimSpace(inputF), 64)
	return input
}

func tambahStartup(reader *bufio.Reader) {
	fmt.Println()
	fmt.Println("----------------------------------")
	fmt.Println("||\tTambah data Startup\t||")
	fmt.Println("----------------------------------")
	for {
		fmt.Print("Nama Startup\t: ")
		Nama := readLineStr(reader)
		fmt.Print("Bidang Usaha\t: ")
		BidangUsaha := readLineStr(reader)
		fmt.Print("Pendanaan\t: ")
		TotalPendanaan := readLineFloat(reader)
		fmt.Print("Tahun Berdiri\t: ")
		TahunBerdiri := readLineInt(reader)
		startups := Startup{
			Nama:           Nama,
			BidangUsaha:    BidangUsaha,
			TahunBerdiri:   TahunBerdiri,
			TotalPendanaan: TotalPendanaan,
		}
		data = append(data, startups)
		fmt.Println(">> Startup berhasil ditambahkan!")
		if konvirmasi() {
			break
		}
	}
	fmt.Println()
}

func tampilkanStartup() {
	for {
		if len(data) == 0 {
			fmt.Println(">> Belum ada startup yang tersedia.")
			break
		}
		for i, s := range data {
			fmt.Println()
			fmt.Println("-----+-----+-----+-----+-----+-----")
			fmt.Printf("[%d] %s\n", i+1, s.Nama)
			fmt.Printf("BIdang USaha\t: %s\n", s.BidangUsaha)
			fmt.Printf("Tahun Berdiri\t: %d\n", s.TahunBerdiri)
			fmt.Printf("Pendanaan\t: Rp.%.2f\n", s.TotalPendanaan)
			fmt.Println("Tim\t\t:")
			for _, t := range s.Tim {
				fmt.Printf("  -> %s sebagai %s\n", t.Nama, t.Peran)
			}
			fmt.Println("-----+-----+-----+-----+-----+-----")
		}
		if konvirmasi() {
			break
		}
	}
	fmt.Println()
}

func Tampilkan() {
	for i, s := range data {
		fmt.Println()
		fmt.Println("-----+-----+-----+-----+-----+-----")
		fmt.Printf("[%d] %s\n", i+1, s.Nama)
		fmt.Printf("BIdang USaha\t: %s\n", s.BidangUsaha)
		fmt.Printf("Tahun Berdiri\t: %d\n", s.TahunBerdiri)
		fmt.Printf("Pendanaan\t: Rp.%.2f\n", s.TotalPendanaan)
		fmt.Println("Tim\t\t:")
		for _, t := range s.Tim {
			fmt.Printf("  -> %s sebagai %s\n", t.Nama, t.Peran)
		}
		fmt.Println("-----+-----+-----+-----+-----+-----")
	}
}

func ubahStartup(reader *bufio.Reader) {
	fmt.Println()
	fmt.Println("----------------------------------")
	fmt.Println("||\tUbah data Startup\t||")
	fmt.Println("----------------------------------")
	for {
		if len(data) == 0 {
			fmt.Println("Belum ada startup yang tersedia.")
			break
		}
		Tampilkan()
		fmt.Println()
		fmt.Println("----------------------------------------------------")
		fmt.Print("Masukkan nomor startup yang ingin diubah: ")
		var idx int
		idx = readLineInt(reader)
		idx--
		if idx < 0 || idx >= len(data) {
			fmt.Println(">> Indeks tidak valid.")
			return
		}
		fmt.Printf("Nama [%s]: ", data[idx].Nama)
		input := readLineStr(reader)
		if input != "" {
			data[idx].Nama = input
		}
		fmt.Printf("Bidang [%s]: ", data[idx].BidangUsaha)
		input = readLineStr(reader)
		if input != "" {
			data[idx].BidangUsaha = input
		}
		fmt.Printf("Pendanaan [%.2f]: ", data[idx].TotalPendanaan)
		n := readLineFloat(reader)
		if n > 0 {
			data[idx].TotalPendanaan = n
		}
		fmt.Printf("Tahun Berdiri [%d]: ", data[idx].TahunBerdiri)
		y := readLineInt(reader)
		if y > 0 {
			data[idx].TahunBerdiri = y
		}
		fmt.Println(">> Data berhasil diubah.")
		if konvirmasi() {
			break
		}
	}
	fmt.Println()
}

func hapusStartup(reader *bufio.Reader) {
	fmt.Println()
	fmt.Println("----------------------------------")
	fmt.Println("||\tHapus data Startup\t||")
	fmt.Println("----------------------------------")
	for {
		if len(data) == 0 {
			fmt.Println("Belum ada startup yang tersedia.")
			break
		}
		Tampilkan()
		fmt.Print("Masukkan nomor startup yang ingin dihapus: ")
		var idx int
		idx = readLineInt(reader)
		idx--
		if idx < 0 || idx >= len(data) {
			fmt.Println(">> Indeks tidak valid.")
			return
		}
		fmt.Println(">> Startup", data[idx].Nama, "berhasil dihapus.")
		data = append(data[:idx], data[idx+1:]...)
		if konvirmasi() {
			break
		}
	}
	fmt.Println()
}

func cari(reader *bufio.Reader) {
	fmt.Println()
	fmt.Println("----------------------------------")
	fmt.Println("||\tCari data Startup\t||")
	fmt.Println("----------------------------------")
	for {
		if len(data) == 0 {
			fmt.Println("Belum ada startup yang tersedia.")
			break
		}
		fmt.Println("1.	Cari Berdasarkan Bidang(Squential Search)")
		fmt.Println("2.	Cari Berdasarkan Nama(Binary Search)")
		fmt.Print("Piliha(1-2) :")
		pilih := readLineInt(reader)
		if pilih == 1 {
			cariBidang(reader)
		} else if pilih == 2 {
			cariNama(reader)
		} else {
			fmt.Println("Pilihan Salah!!!")
		}
		if konvirmasi() {
			break
		}
	}
	fmt.Println()
}

func cariBidang(reader *bufio.Reader) {
	fmt.Print("Masukkan kata kunci bidang usaha: ")
	keyword := strings.ToLower(readLineStr(reader))
	found := false
	for i := 0; i < len(data); i++ {
		bidang := strings.ToLower(data[i].BidangUsaha)
		// Proses pencarian sekuensial manual
		match := true
		if !strings.Contains(bidang, keyword) {
			match = false
		}
		if match {
			fmt.Println()
			fmt.Println("==================== Ditemukan =====================")
			fmt.Printf("[%d] %s\n", i+1, data[i].Nama)
			fmt.Printf("BIdang USaha\t: %s\n", data[i].BidangUsaha)
			fmt.Printf("Tahun Berdiri\t: %d\n", data[i].TahunBerdiri)
			fmt.Printf("Pendanaan\t: Rp.%.2f\n", data[i].TotalPendanaan)
			fmt.Println("Tim\t\t:")
			for _, t := range data[i].Tim {
				fmt.Printf("  -> %s sebagai %s\n", t.Nama, t.Peran)
			}
			fmt.Println("====================================================")
			found = true
		}
	}
	if !found {
		fmt.Println(">> Data tidak ditemukan.")
	}
}

func insertionSortByNama() {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && strings.ToLower(data[j].Nama) > strings.ToLower(key.Nama) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func cariNama(reader *bufio.Reader) {
	if len(data) == 0 {
		fmt.Println("Belum ada startup yang tersedia.")
		return
	}
	insertionSortByNama() // pastikan diurutkan dulu berdasarkan nama
	fmt.Print("\nMasukkan nama startup yang ingin dicari: ")
	keyword := readLineStr(reader)
	low := 0
	high := len(data) - 1
	found := false
	for low <= high {
		mid := (low + high) / 2
		namaMid := strings.ToLower(data[mid].Nama)
		if namaMid == keyword {
			fmt.Println("\n==================== Ditemukan =====================")
			fmt.Printf("[%d] %s\n", mid+1, data[mid].Nama)
			fmt.Printf("BIdang USaha\t: %s\n", data[mid].BidangUsaha)
			fmt.Printf("Tahun Berdiri\t: %d\n", data[mid].TahunBerdiri)
			fmt.Printf("Pendanaan\t: Rp.%.2f\n", data[mid].TotalPendanaan)
			fmt.Println("Tim\t\t:")
			for _, t := range data[mid].Tim {
				fmt.Printf("  -> %s sebagai %s\n", t.Nama, t.Peran)
			}
			fmt.Println("====================================================")
			found = true
			break
		} else if namaMid < keyword {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if !found {
		fmt.Println("Startup tidak ditemukan.")
	}
}

func urut(reader *bufio.Reader) {
	fmt.Println()
	fmt.Println("----------------------------------")
	fmt.Println("||\tUrutkan data Startup\t||")
	fmt.Println("----------------------------------")
	for {
		if len(data) == 0 {
			fmt.Println("Belum ada startup yang tersedia.")
			break
		}
		fmt.Println("1.	Urutkan Berdasarkan Total Pendanaan (Selection Sort)")
		fmt.Println("2.	Urutkan Berdasarkan Tahun Berdiri (Insertion Sort)")
		fmt.Print("Piliha(1-2) :")
		pilih := readLineInt(reader)
		if pilih == 1 {
			urutDana()
		} else if pilih == 2 {
			urutTahunBerdiri()
		} else {
			fmt.Println("Pilihan Salah!!!")
		}
		if konvirmasi() {
			break
		}
	}
	fmt.Println()
}

func urutDana() {
	for i := 0; i < len(data)-1; i++ {
		maxIdx := i
		for j := i + 1; j < len(data); j++ {
			if data[j].TotalPendanaan > data[maxIdx].TotalPendanaan {
				maxIdx = j
			}
		}
		data[i], data[maxIdx] = data[maxIdx], data[i]
	}
	fmt.Println("Daftar Startup (Urut Berdasarkan Total Pendanaan - Terbesar ke Terkecil):")
	Tampilkan()
}

func urutTahunBerdiri() {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].TahunBerdiri > key.TahunBerdiri {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
	fmt.Println("Daftar Startup (Urut Berdasarkan Total Pendanaan - Terlama ke Terbaru):")
	Tampilkan()
}

func tambahAnggotaTim(reader *bufio.Reader) {
	fmt.Println()
	fmt.Println("-------------------------------------------------------")
	fmt.Println("|| Tambah daftar anggota tim dan peran dalam Startup ||")
	fmt.Println("-------------------------------------------------------")
	for {
		if len(data) == 0 {
			fmt.Println("Belum ada startup yang tersedia.")
			break
		}
		Tampilkan()
		fmt.Print("\nMasukkan nomor startup yang ingin ditambahkan anggota tim: ")
		idx := readLineInt(reader)
		idx-- // konversi ke indeks array (mulai dari 0)
		if idx < 0 || idx >= len(data) {
			fmt.Println("nomor tidak valid.")
			break
		}
		fmt.Print("Nama Anggota Tim: ")
		nama := readLineStr(reader)
		fmt.Print("Peran Anggota Tim: ")
		peran := readLineStr(reader)
		anggota := Tim{
			Nama:  nama,
			Peran: peran,
		}
		data[idx].Tim = append(data[idx].Tim, anggota)
		fmt.Println(">> Anggota tim berhasil ditambahkan.")
		if konvirmasi() {
			break
		}
	}
	fmt.Println()
}

func laporanBidang() {
	fmt.Println()
	fmt.Println("------------------------------------------------------")
	fmt.Println("|| Laporan jumlah startup per kategori bidang usaha ||")
	fmt.Println("------------------------------------------------------")
	for {
		laporan := make(map[string]int)
		for _, s := range data {
			laporan[s.BidangUsaha]++
		}
		fmt.Println("\nLaporan Jumlah Startup per Bidang Usaha:")
		for bidang, jumlah := range laporan {
			fmt.Printf("<-> %s : %d startup\n", bidang, jumlah)
		}
		if konvirmasi() {
			break
		}
	}
}

func konvirmasi() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Println("----------------------------------------------------")
	fmt.Print("Kembali ke Menu Utama? (ya/tidak): ")
	konvirmasi, _ := reader.ReadString('\n')
	konvirmasi = strings.TrimSpace(strings.ToLower(konvirmasi))
	fmt.Println("----------------------------------------------------")
	if konvirmasi == "ya" {
		return true
	}
	return false
}
