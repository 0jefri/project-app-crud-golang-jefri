package main

import (
	"fmt"
	"os"
	"time"

	"lumosh_klinik/model"
	"lumosh_klinik/service"
	"lumosh_klinik/utils"
	"lumosh_klinik/view"
)

func main() {
	// Load data from files
	if err := service.LoadUsers(); err != nil {
		fmt.Printf("Gagal memuat pengguna: %s\n", err)
		os.Exit(1)
	}
	if err := service.LoadPatients(); err != nil {
		fmt.Printf("Gagal memuat pasien: %s\n", err)
		os.Exit(1)
	}
	if err := service.LoadAppointments(); err != nil {
		fmt.Printf("Gagal memuat janji temu: %s\n", err)
		os.Exit(1)
	}

	for {
		fmt.Println("=== Aplikasi Klinik ===")
		fmt.Println("1. Login")
		fmt.Println("2. Exit")
		var choice int
		fmt.Print("Pilih: ")
		fmt.Scan(&choice)

		if choice == 2 {
			break
		} else if choice == 1 {
			handleLogin()
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func handleLogin() {
	var username, password string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	user, err := service.Login(username, password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Login berhasil! Selamat datang, %s!\n", user.Username)

	if user.Role == "admin" {
		handleAdminMenu()
	} else {
		handleUserMenu()
	}
}

func handleAdminMenu() {
	for {
		fmt.Println("\n=== Menu Admin ===")
		fmt.Println("1. Tambah Pengguna")
		fmt.Println("2. Lihat Pengguna")
		fmt.Println("3. Tambah Pasien")
		fmt.Println("4. Lihat Pasien")
		fmt.Println("5. Tambah Janji Temu")
		fmt.Println("6. Lihat Janji Temu")
		fmt.Println("7. Logout")
		var choice int
		fmt.Print("Pilih: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addUser()
		case 2:
			view.ShowUsers(service.GetUsers())
		case 3:
			addPatient()
		case 4:
			view.ShowPatients(service.GetPatients())
		case 5:
			addAppointment()
		case 6:
			view.ShowAppointments(service.GetAppointments())
		case 7:
			service.Logout("admin")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func handleUserMenu() {
	for {
		fmt.Println("\n=== Menu Pengguna ===")
		fmt.Println("1. Lihat Pasien")
		fmt.Println("2. Tambah Janji Temu")
		fmt.Println("3. Logout")
		var choice int
		fmt.Print("Pilih: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			view.ShowPatients(service.GetPatients())
		case 2:
			addAppointment()
		case 3:
			service.Logout("user")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func addUser() {
	var username, password, role string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)
	fmt.Print("Role (admin/doctor/nurse): ")
	fmt.Scan(&role)

	if err := utils.ValidateUser(username, password, role); err != nil {
		fmt.Println(err)
		return
	}

	user := model.User{ID: len(service.GetUsers()) + 1, Username: username, Password: password, Role: role}
	if err := service.AddUser(user); err != nil {
		fmt.Printf("Gagal menambahkan pengguna: %s\n", err)
		return
	}
	fmt.Println("Pengguna berhasil ditambahkan!")
}

func addPatient() {
	var name, gender, address string
	var age int
	fmt.Print("Nama: ")
	fmt.Scan(&name)
	fmt.Print("Usia: ")
	fmt.Scan(&age)
	fmt.Print("Jenis Kelamin: ")
	fmt.Scan(&gender)
	fmt.Print("Alamat: ")
	fmt.Scan(&address)

	if err := utils.ValidatePatient(name, age, gender, address); err != nil {
		fmt.Println(err)
		return
	}

	patient := model.Patient{ID: len(service.GetPatients()) + 1, Name: name, Age: age, Gender: gender, Address: address}
	if err := service.AddPatient(patient); err != nil {
		fmt.Printf("Gagal menambahkan pasien: %s\n", err)
		return
	}
	fmt.Println("Pasien berhasil ditambahkan!")
}

func addAppointment() {
	var patientID, doctorID int
	var notes string
	var dateStr string
	fmt.Print("ID Pasien: ")
	fmt.Scan(&patientID)
	fmt.Print("ID Dokter: ")
	fmt.Scan(&doctorID)
	fmt.Print("Tanggal (YYYY-MM-DD): ")
	fmt.Scan(&dateStr)
	fmt.Print("Catatan: ")
	fmt.Scan(&notes)

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Println("Format tanggal tidak valid!")
		return
	}

	appointment := model.Appointment{ID: len(service.GetAppointments()) + 1, PatientID: patientID, DoctorID: doctorID, Date: date, Notes: notes}
	if err := service.AddAppointment(appointment); err != nil {
		fmt.Printf("Gagal menambahkan janji temu: %s\n", err)
		return
	}
	fmt.Println("Janji temu berhasil ditambahkan!")
}
