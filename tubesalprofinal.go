package main

import (
	"fmt"
)

type mahasiswa struct {
	Name       string
	Department string
	Score      int
}

const maxStudents = 100

var students [maxStudents]mahasiswa
var studentCount int

func main() {
	userRole := ""
	fmt.Println("masukan peran pengguna (admin/mahasiswa):")
	fmt.Scanln(&userRole)

	if userRole == "admin" {
		adminMenu()
	} else if userRole == "mahasiswa" {
		studentMenu()
	} else {
		fmt.Println("peran pengguna tidak valid.")
	}
}

func adminMenu() {
	for {
		fmt.Println("admin Menu:")
		fmt.Println("1. input mahasiswa")
		fmt.Println("2. edit mahasiswa")
		fmt.Println("3. hapus mahasiswa")
		fmt.Println("4. tampilan mahasiswa")
		fmt.Println("5. input/edit/hapus nilai mahasiswa")
		fmt.Println("6. output data siswa yang diurutkan")
		fmt.Println("7. cari mahasiswa")
		fmt.Println("8. exit")
		choice := 0
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addStudent()
		case 2:
			editStudent()
		case 3:
			deleteStudent()
		case 4:
			displayStudents()
		case 5:
			manageStudentScores()
		case 6:
			outputSortedData()
		case 7:
			searchStudent()
		case 8:
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func studentMenu() {
	for {
		fmt.Println("menu mahasiswa:")
		fmt.Println("1. input mahasiswa")
		fmt.Println("2. hapus mahasiswa")
		fmt.Println("3. tampilan mahasiswa")
		fmt.Println("4. output data siswa yang diurutkan")
		fmt.Println("5. cari mahasiswa")
		fmt.Println("6. exit")
		choice := 0
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addStudent()
		case 2:
			deleteStudent()
		case 3:
			displayStudents()
		case 4:
			outputSortedData()
		case 5:
			searchStudent()
		case 6:
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func addStudent() {
	if studentCount >= maxStudents {
		fmt.Println("Kapasitas mahasiswa penuh.")
		return
	}

	var name, department string
	var score int
	fmt.Println("Masukan nama mahasiswa:")
	fmt.Scanln(&name)
	fmt.Println("Masukan fakultas mahasiswa:")
	fmt.Scanln(&department)
	fmt.Println("Masukan nilai mahasiswa (0-100):")
	fmt.Scanln(&score)

	if score >= 0 && score <= 100 {
		students[studentCount] = mahasiswa{name, department, score}
		studentCount++
		fmt.Println("Mahasiswa berhasil terdaftar.")
	} else {
		fmt.Println("Masukkan nilai yang valid (0-100).")
	}
}

func editStudent() {
	var name string
	fmt.Println("masukan nama mahasiswa yang akan diedit:")
	fmt.Scanln(&name)
	for i := 0; i < studentCount; i++ {
		if students[i].Name == name {
			fmt.Println("masukan fakultas baru:")
			fmt.Scanln(&students[i].Department)
			fmt.Println("masukan nilai baru:")
			fmt.Scanln(&students[i].Score)
			fmt.Println("mahasiswa berhasil diedit.")
			return
		}
	}
	fmt.Println("mahasiswa tidak ditemukan.")
}

func deleteStudent() {
	var name string
	fmt.Println("masukan nama mahasiswa yang akan dihapus:")
	fmt.Scanln(&name)
	for i := 0; i < studentCount; i++ {
		if students[i].Name == name {
			for j := i; j < studentCount-1; j++ {
				students[j] = students[j+1]
			}
			students[studentCount-1] = mahasiswa{}
			studentCount--
			fmt.Println("mahasiswa berhasil dihapus.")
			return
		}
	}
	fmt.Println("mahasiswa tidak terdaftar.")
}

func displayStudents() {
	if studentCount == 0 {
		fmt.Println("tidak ada mahasiswa yang tersedia.")
		return
	}
	for i := 0; i < studentCount; i++ {
		s := students[i]
		if s.Score > 80 {
			fmt.Printf("nama: %s, fakultas: %s, nilai: %d status: diterima\n", s.Name, s.Department, s.Score)
		} else {
			fmt.Printf("nama: %s, fakultas: %s, nilai: %d status: ditolak\n", s.Name, s.Department, s.Score)
		}
	}
}

func manageStudentScores() {
	for {
		fmt.Println("1. Tambah nilai mahasiswa")
		fmt.Println("2. Edit nilai mahasiswa")
		fmt.Println("3. Hapus nilai mahasiswa")
		fmt.Println("4. Kembali")
		choice := 0
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addStudentScore()
		case 2:
			editStudentScore()
		case 3:
			deleteStudentScore()
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func addStudentScore() {
	var name string
	fmt.Println("masukan nama mahasiswa yang akan ditambah nilainya:")
	fmt.Scanln(&name)
	for i := 0; i < studentCount; i++ {
		if students[i].Name == name {
			var score int
			fmt.Println("masukan nilai baru:")
			fmt.Scanln(&score)
			students[i].Score = score
			fmt.Println("nilai mahasiswa berhasil ditambahkan.")
			return
		}
	}
	fmt.Println("mahasiswa tidak ditemukan.")
}

func editStudentScore() {
	var name string
	fmt.Println("masukan nama mahasiswa yang akan diedit nilainya:")
	fmt.Scanln(&name)
	for i := 0; i < studentCount; i++ {
		if students[i].Name == name {
			fmt.Println("masukan nilai baru:")
			fmt.Scanln(&students[i].Score)
			fmt.Println("nilai terbaru mahasiswa berhasil diedit.")
			return
		}
	}
	fmt.Println("mahasiswa tidak ditemukan.")
}

func deleteStudentScore() {
	var name string
	fmt.Println("masukan nama mahasiswa yang akan dihapus nilainya:")
	fmt.Scanln(&name)
	for i := 0; i < studentCount; i++ {
		if students[i].Name == name {
			students[i].Score = 0
			fmt.Println("nilai mahasiswa berhasil dihapus.")
			return
		}
	}
	fmt.Println("mahasiswa tidak ditemukan.")
}

func outputSortedData() {
	if studentCount == 0 {
		fmt.Println("tidak ada mahasiswa yang tersedia.")
		return
	}
	fmt.Printf("terurut oleh: \n1. nilai \n2. jurusan \n3. nama\n")
	choice := 0
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		sortByScore()
	case 2:
		sortByMajor()
	case 3:
		insertionSortByName()
	default:
		fmt.Println("pilihan yang tidak valid.")
	}
}

func sortByScore() {
	fmt.Println("Urutkan berdasarkan nilai:")
	fmt.Println("1. Menaik")
	fmt.Println("2. Menurun")
	order := 0
	fmt.Scanln(&order)

	for i := 0; i < studentCount-1; i++ {
		selectedIdx := i
		for j := i + 1; j < studentCount; j++ {
			if (order == 1 && students[j].Score < students[selectedIdx].Score) ||
				(order == 2 && students[j].Score > students[selectedIdx].Score) {
				selectedIdx = j
			}
		}
		students[i], students[selectedIdx] = students[selectedIdx], students[i]
	}
	displayStudents()
}

func sortByMajor() {
	for i := 0; i < studentCount-1; i++ {
		minIdx := i
		for j := i + 1; j < studentCount; j++ {
			if students[j].Department < students[minIdx].Department {
				minIdx = j
			}
		}
		students[i], students[minIdx] = students[minIdx], students[i]
	}
	displayStudents()
}

func insertionSortByName() {
	for i := 1; i < studentCount; i++ {
		key := students[i]
		j := i - 1
		for j >= 0 && students[j].Name > key.Name {
			students[j+1] = students[j]
			j = j - 1
		}
		students[j+1] = key
	}
	displayStudents()
}

func searchStudent() {
	fmt.Println("Cari mahasiswa berdasarkan:")
	fmt.Println("1. Nama")
	fmt.Println("2. Fakultas")
	fmt.Println("3. Nilai")
	choice := 0
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var name string
		fmt.Println("masukkan nama:")
		fmt.Scanln(&name)
		insertionSortByName() // Ensure the list is sorted by name
		index := binarySearchByName(name)
		if index != -1 {
			fmt.Printf("Ditemukan mahasiswa: nama: %s, fakultas: %s, nilai: %d\n", students[index].Name, students[index].Department, students[index].Score)
		} else {
			fmt.Println("Mahasiswa tidak ditemukan.")
		}
	case 2:
		var department string
		fmt.Println("masukkan fakultas:")
		fmt.Scanln(&department)
		sequentialSearchByDepartment(department)
	case 3:
		var score int
		fmt.Println("masukkan nilai:")
		fmt.Scanln(&score)
		sequentialSearchByScore(score)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func binarySearchByName(name string) int {
	low, high := 0, studentCount-1
	for low <= high {
		mid := (low + high) / 2
		if students[mid].Name == name {
			return mid
		} else if students[mid].Name < name {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func sequentialSearchByDepartment(department string) {
	for i := 0; i < studentCount; i++ {
		if students[i].Department == department {
			fmt.Printf("Ditemukan mahasiswa: nama: %s, fakultas: %s, nilai: %d\n", students[i].Name, students[i].Department, students[i].Score)
			return
		}
	}
	fmt.Println("Mahasiswa tidak ditemukan.")
}

func sequentialSearchByScore(score int) {
	for i := 0; i < studentCount; i++ {
		if students[i].Score == score {
			fmt.Printf("Ditemukan mahasiswa: nama: %s, fakultas: %s, nilai: %d\n", students[i].Name, students[i].Department, students[i].Score)
			return
		}
	}
	fmt.Println("Mahasiswa tidak ditemukan.")
}
