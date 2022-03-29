package dice

import (
	"fmt"
	"math/rand"
	"time"
)

//opsi
var options string
var optionsPlayer int

//Player
var playerSatu string
var playerDua string
var playerTiga string
var playerEmpat string
var playerLima string

//Dadu
var NomorDadu = []int{1, 2, 3, 4, 5, 6}

//menyimpan nilai dadu
var DaduPlayerSatu = []int{}
var DaduPlayerDua = []int{}
var DaduPlayerTiga = []int{}
var DaduPlayerEmpat = []int{}
var DaduPlayerLima = []int{}

//sum skor player
var NilaiPlayerSatu int
var NilaiPlayerDua int
var NilaiPlayerTiga int
var NilaiPlayerEmpat int
var NilaiPlayerLima int

//putaran permainan
var batasPermainan = struct {
	BatasBawah int
	BatasAtas  int
}{
	BatasBawah: 1,
	BatasAtas:  5,
}

var Putaran = batasPermainan.BatasBawah + rand.Intn(batasPermainan.BatasAtas-batasPermainan.BatasBawah+1)

func Greeting() {
	fmt.Println("Selamat Datang di Permainan Dadu!")
}

func MemulaiPermainan() {
	fmt.Println("'Apakah anda ingin melanjutkan permainan? (y/n)'")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&options)
	fmt.Println("")
	if options == "y" {
		fmt.Println("Masukkan banyaknya karakter yang diinginkan (min:2 & maks: 5)")
		Karakter()
	} else if options == "n" {
		fmt.Println("Permainan tidak dilanjutkan!")
	} else {
		fmt.Println("Masukkan pilihan anda dengan benar!")
	}
}

func Karakter() {
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&optionsPlayer)
	fmt.Println("")

	if optionsPlayer == 2 {
		fmt.Print("Masukkan nama karakter pertama: ")
		fmt.Scan(&playerSatu)

		fmt.Print("Masukkan nama karakter kedua: ")
		fmt.Scan(&playerDua)

		permainanDadu()
	} else if optionsPlayer == 3 {
		fmt.Print("Masukkan nama karakter pertama: ")
		fmt.Scan(&playerSatu)

		fmt.Print("Masukkan nama karakter kedua: ")
		fmt.Scan(&playerDua)

		fmt.Print("Masukkan nama karakter ketiga: ")
		fmt.Scan(&playerTiga)

		permainanDadu()
	} else if optionsPlayer == 4 {
		fmt.Print("Masukkan nama karakter pertama: ")
		fmt.Scan(&playerSatu)

		fmt.Print("Masukkan nama karakter kedua: ")
		fmt.Scan(&playerDua)

		fmt.Print("Masukkan nama karakter ketiga: ")
		fmt.Scan(&playerTiga)

		fmt.Print("Masukkan nama karakter keempat: ")
		fmt.Scan(&playerEmpat)

		permainanDadu()
	} else if optionsPlayer == 5 {
		fmt.Print("Masukkan nama karakter pertama: ")
		fmt.Scan(&playerSatu)

		fmt.Print("Masukkan nama karakter kedua: ")
		fmt.Scan(&playerDua)

		fmt.Print("Masukkan nama karakter ketiga: ")
		fmt.Scan(&playerTiga)

		fmt.Print("Masukkan nama karakter keempat: ")
		fmt.Scan(&playerEmpat)

		fmt.Print("Masukkan nama karakter kelima: ")
		fmt.Scan(&playerLima)

		permainanDadu()
	} else {
		fmt.Println("Masukkan pilihan karakter anda dengan benar!")
	}
}

func PutaranPermainan() {
	rand.Seed(time.Now().Unix())
	fmt.Println("")
	fmt.Println("Banyaknya putaran permainan dadu adalah:", Putaran)
}

func GreetingLemparDadu() {
	fmt.Println("")
	fmt.Println("Memulai permainan dadu!")
}

func permainanDadu() {
	PutaranPermainan()
	GreetingLemparDadu()
	for i := 1; i < Putaran+1; i++ {
		fmt.Println("")
		fmt.Printf("Putaran Permainan ke-%d\n", i)
		if optionsPlayer == 2 {
			DaduPlayerSatu = append(DaduPlayerSatu, NomorDadu[rand.Intn(len(NomorDadu))])
			DaduPlayerDua = append(DaduPlayerDua, NomorDadu[rand.Intn(len(NomorDadu))])
			fmt.Printf("%s: %d\n", playerSatu, DaduPlayerSatu)
			fmt.Printf("%s: %d\n", playerDua, DaduPlayerDua)
			for _, nilai := range DaduPlayerSatu {
				if nilai == 6 {
					var skor int
					skor += 1
					NilaiPlayerSatu += skor
					fmt.Printf("Skor %s adalah %d\n", playerSatu, NilaiPlayerSatu)
				}
			}
			for _, nilai := range DaduPlayerDua {
				if nilai == 6 {
					var skor int
					skor += 1
					NilaiPlayerDua += skor
					fmt.Printf("Skor %s adalah %d\n", playerDua, NilaiPlayerDua)
				}
			}
		} else if optionsPlayer == 3 {
			DaduPlayerSatu = append(DaduPlayerSatu, NomorDadu[rand.Intn(len(NomorDadu))])
			DaduPlayerDua = append(DaduPlayerDua, NomorDadu[rand.Intn(len(NomorDadu))])
			DaduPlayerTiga = append(DaduPlayerTiga, NomorDadu[rand.Intn(len(NomorDadu))])
			fmt.Printf("%s: %d\n", playerSatu, DaduPlayerSatu)
			fmt.Printf("%s: %d\n", playerDua, DaduPlayerDua)
			fmt.Printf("%s: %d\n", playerTiga, DaduPlayerTiga)
			for _, nilai := range DaduPlayerSatu {
				if nilai == 6 {
					var skor int
					skor += 1
					NilaiPlayerSatu += skor
					fmt.Printf("Skor %s adalah %d\n", playerSatu, NilaiPlayerSatu)
				}
			}
			for _, nilai := range DaduPlayerDua {
				if nilai == 6 {
					var skor int
					skor += 1
					NilaiPlayerDua += skor
					fmt.Printf("Skor %s adalah %d\n", playerDua, NilaiPlayerDua)
				}
			}
			for _, nilai := range DaduPlayerTiga {
				if nilai == 6 {
					var skor int
					skor += 1
					NilaiPlayerTiga += skor
					fmt.Printf("Skor %s adalah %d\n", playerTiga, NilaiPlayerTiga)
				}
			}
		} else if optionsPlayer == 4 {
			DaduPlayerSatu = append(DaduPlayerSatu, NomorDadu[rand.Intn(len(NomorDadu))])
			DaduPlayerDua = append(DaduPlayerDua, NomorDadu[rand.Intn(len(NomorDadu))])
			DaduPlayerTiga = append(DaduPlayerTiga, NomorDadu[rand.Intn(len(NomorDadu))])
			DaduPlayerEmpat = append(DaduPlayerEmpat, NomorDadu[rand.Intn(len(NomorDadu))])
			fmt.Printf("%s: %d\n", playerSatu, DaduPlayerSatu)
			fmt.Printf("%s: %d\n", playerDua, DaduPlayerDua)
			fmt.Printf("%s: %d\n", playerTiga, DaduPlayerTiga)
			fmt.Printf("%s: %d\n", playerEmpat, DaduPlayerEmpat)
			for _, nilai := range DaduPlayerSatu {
				if nilai == 6 {
					var skor int
					skor += 1
					NilaiPlayerSatu += skor
					fmt.Printf("Skor %s adalah %d\n", playerSatu, NilaiPlayerSatu)
				}
			}
			for _, nilai := range DaduPlayerDua {
				if nilai == 6 {
					var skor int
					skor += 1
					NilaiPlayerDua += skor
					fmt.Printf("Skor %s adalah %d\n", playerDua, NilaiPlayerDua)
				}
			}
			for _, nilai := range DaduPlayerTiga {
				if nilai == 6 {
					var skor int
					skor += 1
					NilaiPlayerTiga += skor
					fmt.Printf("Skor %s adalah %d\n", playerTiga, NilaiPlayerTiga)
				}
			}
			for _, nilai := range DaduPlayerEmpat {
				if nilai == 6 {
					var skor int
					skor += 1
					NilaiPlayerEmpat += skor
					fmt.Printf("Skor %s adalah %d\n", playerEmpat, NilaiPlayerEmpat)
				}
			}
		} else if optionsPlayer == 5 {
			DaduPlayerSatu = append(DaduPlayerSatu, NomorDadu[rand.Intn(len(NomorDadu))])
			DaduPlayerDua = append(DaduPlayerDua, NomorDadu[rand.Intn(len(NomorDadu))])
			DaduPlayerTiga = append(DaduPlayerTiga, NomorDadu[rand.Intn(len(NomorDadu))])
			DaduPlayerEmpat = append(DaduPlayerEmpat, NomorDadu[rand.Intn(len(NomorDadu))])
			DaduPlayerLima = append(DaduPlayerLima, NomorDadu[rand.Intn(len(NomorDadu))])
			fmt.Printf("%s: %d\n", playerSatu, DaduPlayerSatu)
			fmt.Printf("%s: %d\n", playerDua, DaduPlayerDua)
			fmt.Printf("%s: %d\n", playerTiga, DaduPlayerTiga)
			fmt.Printf("%s: %d\n", playerEmpat, DaduPlayerEmpat)
			fmt.Printf("%s: %d\n", playerLima, DaduPlayerLima)
			for _, nilai := range DaduPlayerSatu {
				if nilai == 6 {
					var skor int
					skor += 1
					NilaiPlayerSatu += skor
					fmt.Printf("Skor %s adalah %d\n", playerSatu, NilaiPlayerSatu)
				}
			}
			for _, nilai := range DaduPlayerDua {
				if nilai == 6 {
					var skor int
					skor += 1
					NilaiPlayerDua += skor
					fmt.Printf("Skor %s adalah %d\n", playerDua, NilaiPlayerDua)
				}
			}
			for _, nilai := range DaduPlayerTiga {
				if nilai == 6 {
					var skor int
					skor += 1
					NilaiPlayerTiga += skor
					fmt.Printf("Skor %s adalah %d\n", playerTiga, NilaiPlayerTiga)
				}
			}
			for _, nilai := range DaduPlayerEmpat {
				if nilai == 6 {
					var skor int
					skor += 1
					NilaiPlayerEmpat += skor
					fmt.Printf("Skor %s adalah %d\n", playerEmpat, NilaiPlayerEmpat)
				}
			}
			for _, nilai := range DaduPlayerLima {
				if nilai == 6 {
					var skor int
					skor += 1
					NilaiPlayerLima += skor
					fmt.Printf("Skor %s adalah %d\n", playerLima, NilaiPlayerLima)
				}
			}
		}
	}
}
