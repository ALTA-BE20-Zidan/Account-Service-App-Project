package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"zidan/AccountServiceAppProject/controllers"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	// open connection
	db, errOpen := sql.Open("mysql", os.Getenv("AppConnection"))
	// <username>:<password>@tcp(<hostname>:<port-db>)/<db-name>
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_elibrary?parseTime=true")
	if errOpen != nil {
		log.Fatal("open connection is failed: ", errOpen)
		return nil, errOpen
	}

	// test connection
	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("no connection available: ", errPing)
		return nil, errPing
	} else {
		fmt.Println("you are succesfully connected to database")
		return db, nil
	}
}

func main() {
	// make connection
	db, errInitDB := InitDB()
	if errInitDB != nil {
		log.Fatal("init db is failed: ", errInitDB)
	}

	defer db.Close()

	fmt.Println()

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
		controllers.AddAccountController(db)
	case 2:
		fmt.Println("Welcome to Login!")
		controllers.LoginVerificationController(db)
	case 3:
		fmt.Println("Welcome to Read Account!")
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
	}
}
