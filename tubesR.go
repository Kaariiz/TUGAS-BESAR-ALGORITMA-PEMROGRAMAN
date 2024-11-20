package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// struct untuk menyimpan informasi tentang jam kerja karyawan
type Job struct {
	date                string
	tipe1, tipe2, tipe3 int
}

// struct untuk menyimpan informasi karyawan
type Karyawan struct {
	id                    int
	name, alamat, telepon string
	workHours             [NMAX]Job
	jobCount              int
}

// struct untuk menyimpan rekap data karyawan
type JobRekap struct {
	karyawan string
	durasi   int
}

// constanta yang digunakan untuk menghitung jam kerja karyawan
const (
	NMAX           = 100
	totalWorkHours = 160
	minTipe1       = 0.25
	maxTipe1       = 0.50
	minTipe23      = 0.10
)

// tipe alias karyawan
type tabKar [NMAX]Karyawan

// variabel global
var arrKar tabKar //var untuk menyimpan data karyawan
var karCount int //var untuk menghitung banyak data karyawan

// main function
func main() {
	menu()
}

func menu() {
	var pilih int
	var stop bool

	stop = false
	for !stop {
		fmt.Println("------------------------------------------------------------")
		fmt.Println("    Selamat Datang di Program Penilaian Kinerja Karyawan   ")
		fmt.Println("------------------------------------------------------------")
		fmt.Println("1. Menambahkan Data Karyawan")
		fmt.Println("2. Pengubahan Data Karyawan")
		fmt.Println("3. Menghapus Data Karyawan")
		fmt.Println("4. Lihat Data Karyawan")
		fmt.Println("5. Menambahkan Data Log Pekerjaan Karyawan")
		fmt.Println("6. Mengubah Data Log Pekerjaan Karyawan")
		fmt.Println("7. Menghapus Data Log Pekerjaan Karyawan")
		fmt.Println("8. Rekap Data Pekerjaan Tiap Tipe dalam Bulan Tertentu")
		fmt.Println("9. Rekap Data Terurut Aktivitas berdasarkan durasi")
		fmt.Println("0. Kembali")
		fmt.Println("-----------------------------------------------------------")

		fmt.Scan(&pilih)

		if pilih == 1 {
			clearScreen()
			menambah_data_karyawan(&arrKar)
		} else if pilih == 2 {
			clearScreen()
			ubah_data_karyawan(&arrKar)
		} else if pilih == 3 {
			clearScreen()
			hapus_data_karyawan(&arrKar)
		} else if pilih == 4 {
			lihat_data_karyawan(arrKar)
		} else if pilih == 5 {
			clearScreen()
			menambah_data_log(&arrKar)
		} else if pilih == 6 {
			clearScreen()
			ubah_data_log(&arrKar)
		} else if pilih == 7 {
			clearScreen()
			hapus_data_log(&arrKar)
		} else if pilih == 8 {
			clearScreen()
			rekap_tipe_bulanan(arrKar)
		} else if pilih == 9 {
			clearScreen()
			rekap_terurut_durasi(arrKar)

		} else if pilih == 0 {
			clearScreen()
			fmt.Println("Terima Kasih Sudah Menggunakan Pelayanan!")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}

		stop = pilih == 0

	}
}

// binary search
func mencari_index_karyawan(A tabKar, cari int) int {
	// mengembalikan index karyawan
	left := 0
	right := karCount - 1

	for left <= right {
		mid := left + (right-left)/2
		if A[mid].id == cari {
			return mid
		}
		if A[mid].id < cari {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// sequential search
func search_by_id(A tabKar, id int) int {
	// mengembalikan index dari id karyawan
	for i := 0; i < karCount; i++ {
		if A[i].id == id {
			return i
		}
	}
	return -1
}

func namaKaryawan() string {
	// mengembalikan string dari nama agar bisa menginput lebih dari satu kata
	var nama, result string

	for {
		fmt.Scan(&nama)
		if nama == "." {
			break
		}
		if result == "" {
			result = nama
		} else {
			result += " " + nama
		}
	}
	return result
}

func alamatKaryawan() string {
	// mengembalikan string dari alamat agar bisa menginput lebih dari satu kata
	var alamat, result string

	for {
		fmt.Scan(&alamat)
		if alamat == "." {
			break
		}
		if result == "" {
			result = alamat
		} else {
			result += " " + alamat
		}
	}
	return result
}

func menambah_data_karyawan(A *tabKar) {
	// IS: terdefinisi A
	// FS: menambah data karyawan ke dalam array
	var nama, alamat, telepon string
	var id int

	if karCount > 0 {
		id = (*A)[karCount-1].id + 1
	} else {
		id = 1
	}

	fmt.Print("Masukkan nama karyawan: ")
	nama = namaKaryawan()
	fmt.Print("Masukkan alamat karyawan: ")
	alamat = alamatKaryawan()
	fmt.Print("Masukkan nomor telepon karyawan: ")
	fmt.Scan(&telepon)

	(*A)[karCount] = Karyawan{
		id:        id,
		name:      nama,
		alamat:    alamat,
		telepon:   telepon,
		workHours: [NMAX]Job{},
		jobCount:  0,
	}

	karCount++
	fmt.Println("Data karyawan berhasil ditambahkan.")

}

func hapus_data_karyawan(A *tabKar) {
	// IS: terdefinisi A
	// FS: menghapus data karyawan yang diinginkan
	var idPilih int
	var tempAns string

	lihat_data_karyawan(*A)
	fmt.Print("Pilih id karyawan yang ingin dihapus: ")
	fmt.Scan(&idPilih)

	indexSearch := mencari_index_karyawan(*A, idPilih)

	if indexSearch >= 0 {
		fmt.Printf("Karyawan yang anda pilih: %s\n", (*A)[indexSearch].name)
		fmt.Print("Apakah anda yakin? (y/n) : ")
		fmt.Scan(&tempAns)

		if tempAns == "y" {
			for i := indexSearch; i < karCount-1; i++ {
				(*A)[i] = (*A)[i+1]
			}
			karCount--
			fmt.Println("Data karyawan berhasil dihapus.")
			lihat_data_karyawan(*A)

		} else {
			fmt.Println("Data karyawan batal dihapus.")
		}
	} else {
		fmt.Println("Karyawan tidak ditemukan")
	}

}

func lihat_data_karyawan(A tabKar) {
	// IS: terdefinisi A
	// FS: menampilkan data karyawan
	if karCount == 0 {
		fmt.Println("Tidak ada data karyawan.")
	}

	for i := 0; i < karCount; i++ {
		fmt.Printf("id: %d\n", A[i].id)
		fmt.Printf("Nama Karyawan: %s\n", A[i].name)
		fmt.Printf("Alamat Karyawan: %s\n", A[i].alamat)
		fmt.Printf("Nomor telepon karyawan: %s\n", A[i].telepon)

		for j := 0; j < A[i].jobCount; j++ {
			fmt.Printf("Tanggal: %s\n", A[i].workHours[j].date)
			fmt.Printf("Tipe 1: %d jam\n", A[i].workHours[j].tipe1)
			fmt.Printf("Tipe 2: %d jam\n", A[i].workHours[j].tipe2)
			fmt.Printf("Tipe 3: %d jam\n", A[i].workHours[j].tipe3)
		}
	}
}

func ubah_data_karyawan(A *tabKar) {
	// IS: terdefinisi A
	// FS: mengubah data karyawan yang diinginkan
	var idPilih int
	var tempAns string

	lihat_data_karyawan(*A)
	fmt.Print("Pilih id karyawan yang ingin diubah: ")
	fmt.Scan(&idPilih)

	indexSearch := search_by_id(*A, idPilih)

	if indexSearch == -1 {
		fmt.Println("ID tidak tersedia")
	}

	if indexSearch >= 0 {
		fmt.Printf("Karyawan yang anda pilih: %s\n", (*A)[indexSearch].name)
		fmt.Print("Apakah anda yakin? (y/n) : ")
		fmt.Scan(&tempAns)

		if tempAns == "y" {
			fmt.Println("Karyawan yang anda pilih: ")
			fmt.Printf("Nama Karyawan: %s\n", (*A)[indexSearch].name)
			fmt.Printf("Alamat Karyawan: %s\n", (*A)[indexSearch].alamat)
			fmt.Printf("Nomor telepon karyawan: %s\n", (*A)[indexSearch].telepon)

			fmt.Print("Masukkan nama karyawan: ")
			nama := namaKaryawan()
			fmt.Print("Masukkan alamat karyawan: ")
			alamat := alamatKaryawan()
			fmt.Print("Masukkan telepon karyawan: ")
			fmt.Scan(&(*A)[indexSearch].telepon)

			(*A)[indexSearch].name = nama
			(*A)[indexSearch].alamat = alamat

			fmt.Println("Data berhasil diubah")
			lihat_data_karyawan(*A)

		} else {
			fmt.Println("Data karyawan batal diubah.")
		}
	}
}

func validasi_jam_menambah_data_log(tipe1, tipe2, tipe3 int) (bool, string) {
	// mengembalikan nilai boolean dan string dari semua tipe, agar terlihat jelas di mana letak tipe yang tidak valid
	totalHours := 160
	minTipe1 := 0.25 * float64(totalHours)
	maxTipe1 := 0.50 * float64(totalHours)
	minTipe23 := 0.10 * float64(totalHours)

	if float64(tipe1) < minTipe1 || float64(tipe1) > maxTipe1 {
		return false, "Jam tipe 1 harus antara 25% sampai 50% dari total jam kerja (160 jam)"

	}

	if float64(tipe2) < minTipe23 {
		return false, "Jam tipe 2 harus minimal 10% dari total jam kerja (160 jam)"
	}

	if float64(tipe3) < minTipe23 {
		return false, "Jam tipe 3 harus minimal 10% dari total jam kerja (160 jam)"
	}

	return true, "Semua tipe valid"
}

func validasi_jam_mengubah_data_log(tipe1, tipe2, tipe3 int) bool {
	//mengembalikan true jika valid, dan false jika tidak
	totalHours := 160
	minTipe1 := 0.25 * float64(totalHours)
	maxTipe1 := 0.50 * float64(totalHours)
	minTipe2 := 0.10 * float64(totalHours)
	minTipe3 := 0.10 * float64(totalHours)

	if float64(tipe1) < minTipe1 || float64(tipe1) > maxTipe1 {
		fmt.Println("Jam kerja tipe 1 tidak valid.")
		return false
	}
	if float64(tipe2) < minTipe2 {
		fmt.Println("Jam kerja tipe 2 tidak valid.")
		return false
	}
	if float64(tipe3) < minTipe3 {
		fmt.Println("Jam kerja tipe 3 tidak valid.")
		return false
	}
	return true
}

func menambah_data_log(A *tabKar) {
	// IS: terdefinisi A
	// FS: menambah data log karyawan yang diinginkan
	var id, tipe1, tipe2, tipe3 int
	var date string

	lihat_data_karyawan(*A)
	fmt.Print("Masukkan id karyawan: ")
	fmt.Scan(&id)

	indexSearch := search_by_id((*A), id)

	if indexSearch >= 0 {
		fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
		fmt.Scan(&date)

		fmt.Print("Masukkan jumlah jam kerja untuk tipe 1: ")
		fmt.Scan(&tipe1)
		fmt.Print("Masukkan jumlah jam kerja untuk tipe 2: ")
		fmt.Scan(&tipe2)
		fmt.Print("Masukkan jumlah jam kerja untuk tipe 3: ")
		fmt.Scan(&tipe3)

		valid, tipeInvalid := validasi_jam_menambah_data_log(tipe1, tipe2, tipe3)

		if !valid {
			fmt.Printf("Jam kerja untuk %s tidak valid. Silakan masukkan data yang sesuai.\n", tipeInvalid)

		} else {
			jobBaru := &Job{
				date:  date,
				tipe1: tipe1,
				tipe2: tipe2,
				tipe3: tipe3,
			}

			A[indexSearch].workHours[A[indexSearch].jobCount] = *jobBaru
			A[indexSearch].jobCount++
			fmt.Println("Data log pekerjaan berhasil ditambahkan.")
		}

	} else {
		fmt.Println("Karyawan tidak ditemukan.")
	}
}

func ubah_data_log(A *tabKar) {
	// IS: terdefinisi A
	// FS: mengubah data log karyawan yang diinginkan
	var id, logIndex, tipe1, tipe2, tipe3 int
	var date string

	lihat_data_karyawan(*A)
	fmt.Print("Masukkan id karyawan yang ingin diubah: ")
	fmt.Scan(&id)

	indexSearch := mencari_index_karyawan(*A, id)

	if indexSearch == -1 {
		fmt.Println("Karyawan tidak ditemukan")

	} else {
		if ((*A)[indexSearch].jobCount) == 0 {
			fmt.Println("Karyawan ini tidak memiliki log pekerjaan")

		} else {
			fmt.Println("Data log pekerjaan yang ada:")
			for i := 0; i < ((*A)[indexSearch].jobCount); i++ {
				job := (*A)[indexSearch].workHours[i]
				fmt.Printf("%d. Tanggal: %s, Tipe 1: %d jam, Tipe 2: %d jam, Tipe 3: %d jam\n", i+1, job.date, job.tipe1, job.tipe2, job.tipe3)
			}

			fmt.Print("Pilih index log pekerjaan yang ingin diubah: ")
			fmt.Scan(&logIndex)

			if logIndex <= 0 || logIndex > ((*A)[indexSearch].jobCount) {
				fmt.Println("Index log pekerjaan tidak valid.")

			} else {
				fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
				fmt.Scan(&date)
				fmt.Print("Masukkan jumlah jam kerja untuk tipe 1: ")
				fmt.Scan(&tipe1)
				fmt.Print("Masukkan jumlah jam kerja untuk tipe 2: ")
				fmt.Scan(&tipe2)
				fmt.Print("Masukkan jumlah jam kerja untuk tipe 3: ")
				fmt.Scan(&tipe3)
			}
		}

		if !validasi_jam_mengubah_data_log(tipe1, tipe2, tipe3) {
			fmt.Println("Jam kerja tidak valid. Silakan masukkan data yang sesuai.")

		} else {

			(*A)[indexSearch].workHours[logIndex-1] = Job{
				date:  date,
				tipe1: tipe1,
				tipe2: tipe2,
				tipe3: tipe3,
			}
		}

		fmt.Println("Data log pekerjaan berhasil diubah.")
	}
}

func hapus_data_log(A *tabKar) {
	// IS: terdefinisi A
	// FS: menghapus data log karyawan yang diinginkan
	var id, logIndex int

	lihat_data_karyawan(*A)
	fmt.Print("Masukkan id karyawan yang ingin dihapus log pekerjaannya: ")
	fmt.Scan(&id)

	indexSearch := mencari_index_karyawan(*A, id)

	if indexSearch == -1 {
		fmt.Println("Karyawan tidak ditemukan.")

	}

	if ((*A)[indexSearch].jobCount) == 0 {
		fmt.Println("Karyawan ini tidak memiliki log pekerjaan.")
	}

	fmt.Println("Data log pekerjaan yang ada:")
	for i := 0; i < ((*A)[indexSearch].jobCount); i++ {
		job := (*A)[indexSearch].workHours[i]
		fmt.Printf("%d. Tanggal: %s, Tipe 1: %d jam, Tipe 2: %d jam, Tipe 3: %d jam\n", i+1, job.date, job.tipe1, job.tipe2, job.tipe3)
	}

	fmt.Print("Pilih index log pekerjaan yang ingin dihapus: ")
	fmt.Scan(&logIndex)

	if logIndex <= 0 || logIndex > ((*A)[indexSearch].jobCount) {
		fmt.Println("Index log pekerjaan tidak valid.")
	} else {
		for i := logIndex - 1; i < (*A)[indexSearch].jobCount-1; i++ {
			(*A)[indexSearch].workHours[i] = (*A)[indexSearch].workHours[i+1]
		}
		(*A)[indexSearch].jobCount--
		fmt.Println("Data log pekerjaan berhasil dihapus.")
	}
}

// BATAS SPESIFIKASI POINT A

func validasi_proporsi(tipe1, tipe2, total int) bool {
	//mengembalikan true jika total jam kerja sesuai tipe, false jika tidak memenuhi
	if total == 0 {
		return false

	} else {
		tipe1Percent := float64(tipe1) / float64(total)
		tipe2Percent := float64(tipe2) / float64(total)
		return (tipe1Percent >= minTipe1 && tipe1Percent <= maxTipe1 && tipe2Percent >= minTipe23)

	}
}

func rekap_log_bulanan(A tabKar) {
	// IS: terdefinisi A
	// FS: menampilkan rekap data log karyawan selama sebulan
	var id int
	var bulanTahun string

	fmt.Print("Masukkan id karyawan: ")
	fmt.Scan(&id)

	indexSearch := search_by_id(A, id)

	if indexSearch == -1 {
		fmt.Println("Karyawan tidak ditemukan.")

	} else {
		fmt.Print("Masukkan bulan dan tahun (YYYY-MM): ")
		fmt.Scan(&bulanTahun)

		// Mengurutkan pekerjaan (ascending) berdasarkan tanggal sebelum rekap (insertion sort)
		for i := 1; i < (A[indexSearch].jobCount); i++ {
			key := A[indexSearch].workHours[i]
			j := i - 1

			for j >= 0 && A[indexSearch].workHours[j].date > key.date {
				A[indexSearch].workHours[j+1] = A[indexSearch].workHours[j]
				j = j - 1
			}
			A[indexSearch].workHours[j+1] = key
		}

		totalTipe1, totalTipe2, totalTipe3 := 0, 0, 0

		for i := 0; i < ((A)[indexSearch].jobCount); i++ {
			job := A[indexSearch].workHours[i]
			if job.date[:7] == bulanTahun {
				totalTipe1 += job.tipe1
				totalTipe2 += job.tipe2
				totalTipe3 += job.tipe3
			}
		}

		totalHours := totalTipe1 + totalTipe2 + totalTipe3

		fmt.Printf("Rekap kerja bulan %s untuk karyawan %s:\n", bulanTahun, A[indexSearch].name)
		fmt.Printf("Total Tipe 1: %d jam\n", totalTipe1)
		fmt.Printf("Total Tipe 2: %d jam\n", totalTipe2)
		fmt.Printf("Total Tipe 3: %d jam\n", totalTipe3)

		if validasi_proporsi(totalTipe1, totalTipe2, totalHours) {
			fmt.Println("Proporsi kerja sesuai ketentuan.")

		} else {
			fmt.Println("Proporsi kerja tidak sesuai ketentuan.")

		}

	}
}

// BATAS SPESIFIKASI POINT B

func rekap_tipe_bulanan(A tabKar) {
	// IS: terdefinisi A
	// FS: menampilkan rekap data pekerjaan dari tiap tipe, selama kurun waktu bulan tertentu
	var bulanTahun string
	var tipe, totalKaryawan int
	var tipeTotal [NMAX]int
	var tipeKaryawan [NMAX]string

	fmt.Print("Masukkan bulan dan tahun (YYYY-MM): ")
	fmt.Scan(&bulanTahun)
	fmt.Print("Masukkan tipe pekerjaan (1/2/3): ")
	fmt.Scan(&tipe)

	for i := 0; i < karCount; i++ {
		total := 0

		for j := 0; j < (A[i].jobCount); j++ {
			if A[i].workHours[j].date[:7] == bulanTahun {
				if tipe == 1 {
					total += A[i].workHours[j].tipe1
				} else if tipe == 2 {
					total += A[i].workHours[j].tipe2
				} else {
					total += A[i].workHours[j].tipe3
				}
			}
		}

		if total > 0 {
			tipeTotal[totalKaryawan] = total
			tipeKaryawan[totalKaryawan] = A[i].name
			totalKaryawan++
		}

		fmt.Printf("Rekap pekerjaan tipe %d bulan %s:\n", tipe, bulanTahun)
		for i := 0; i < totalKaryawan; i++ {
			fmt.Printf("Karyawan: %s, Total Waktu: %d jam\n", tipeKaryawan[i], tipeTotal[i])

		}
	}
}

// BATAS SPESIFIKASI POINT C

func rekap_terurut_durasi(A tabKar) {
	// IS: terdefinisi A
	// FS: menampilkan rekap data terurut aktivitas pekerja berdasarkan durasi waktu
	var bulanTahun string
	var tipe, totalKaryawan int
	var rekaps [NMAX]JobRekap

	fmt.Print("Masukkan bulan dan tahun (YYYY-MM): ")
	fmt.Scan(&bulanTahun)
	fmt.Print("Masukkan tipe pekerjaan (1/2/3): ")
	fmt.Scan(&tipe)

	for i := 0; i < karCount; i++ {
		total := 0
		for j := 0; j < A[i].jobCount; j++ {
			if (A)[i].workHours[j].date[:7] == bulanTahun { // :7 itu fungsinya biar ngebaca sampe karakter ke 7 dari suatu elemen array
				if tipe == 1 {
					total += (A)[i].workHours[j].tipe1
				} else if tipe == 2 {
					total += (A)[i].workHours[j].tipe2
				} else {
					total += (A)[i].workHours[j].tipe3
				}
			}
		}

		rekaps[totalKaryawan] = JobRekap{
			karyawan: A[i].name,
			durasi:   total,
		}
		totalKaryawan++
	}

	// selection sort descending (mengurutkan array rekaps berdasarkan durasi)
	for i := 0; i < totalKaryawan-1; i++ {
		maxIdx := i
		for j := i + 1; j < totalKaryawan; j++ {
			if rekaps[j].durasi > rekaps[maxIdx].durasi {
				maxIdx = j
			}
		}
		rekaps[i], rekaps[maxIdx] = rekaps[maxIdx], rekaps[i]
	}

	fmt.Printf("Rekap pekerjaan tipe %d berdasarkan durasi waktu kerja:\n", tipe)
	for i := 0; i < totalKaryawan; i++ {
		fmt.Printf("Karyawan: %s, Total Waktu: %d jam\n", rekaps[i].karyawan, rekaps[i].durasi)
	}

}

//BATAS SPESIFIKASI POINT D

func clearScreen() {
	//function untuk menghapus tampilan, agar output terlihat lebih rapi
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
