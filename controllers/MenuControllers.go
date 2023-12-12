package controllers

import (
	"database/sql"
	"fmt"
	"os"
)

var status_login bool

func Menu(db *sql.DB) {

	fmt.Println("These are our menu:")
	fmt.Println("1. Add Account (Register)")
	fmt.Println("2. Login")
	fmt.Println("3. Read Account")
	fmt.Println("4. Update Account")
	fmt.Println("5. Delete Account")
	fmt.Println("6. Top-Up")
	fmt.Println("7. Transfer")
	fmt.Println("8. History Top-Up")
	fmt.Println("9. History Transfer")
	fmt.Println("10. Read User's Profile")
	fmt.Println("0. Quit")

	fmt.Println("\nChoose Your Menu:")

	var menu int
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		fmt.Println("Welcome to Add Account (Register)!")
		AddAccountController(db)
		status_login = true
	case 2:
		fmt.Println("Welcome to Login!")
		LoginVerificationController(db)
	case 3:
		if status_login == false {
			fmt.Println("Silakan login terlebih dahulu")
			fmt.Println()
			Menu(db)
		} else if status_login == true {
			fmt.Println("Welcome to Read Account!")
			ReadMyAccountController(db)
		}
	case 4:
		fmt.Println("Welcome to Update Account!")
	case 5:
		fmt.Println("Welcome to Delete Account!")
	case 6:
		fmt.Println("Welcome to Top-Up!")
	case 7:
		fmt.Println("Welcome to Transfer!")
	case 8:
		fmt.Println("Welcome to History Top-Up!")
	case 9:
		fmt.Println("Welcome to History Transfer!")
	case 10:
		fmt.Println("Welcome to Read User's Profile")
	case 0:
		fmt.Println("Terima kasih telah bertransaksi")
		os.Exit(0)
		status_login = false
	default:
		fmt.Println("Menu tidak tersedia")
	}
}