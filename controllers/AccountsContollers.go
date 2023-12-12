package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"zidan/AccountServiceAppProject/entities"
)

func AddAccountControllers(db *sql.DB) {
	var newAccount entities.Accounts
	newAccount.User_balance = 0

	fmt.Println("Masukkan Username:")
	fmt.Scanln(&newAccount.Username)
	fmt.Println("Masukkan Nama:")
	fmt.Scanln(&newAccount.User_nama)
	fmt.Println("Masukkan No. Telpon:")
	fmt.Scanln(&newAccount.User_phone)
	fmt.Println("Masukkan Email:")
	fmt.Scanln(&newAccount.User_email)
	fmt.Println("Masukkan Address:")
	fmt.Scanln(&newAccount.User_address)
	fmt.Println("Atur Password:")
	fmt.Scanln(&newAccount.User_pswd)

	fmt.Println()
	statement, errPrepare := db.Prepare("insert into accounts (username, user_nama, user_phone, user_email, user_address, user_balance, user_pswd) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if errPrepare != nil {
		log.Fatal("err prepare: ", errPrepare)
	}

	result, errExec := statement.Exec(&newAccount.Username, &newAccount.User_nama, &newAccount.User_phone, &newAccount.User_email, &newAccount.User_address, &newAccount.User_balance, &newAccount.User_pswd)
	if errExec != nil {
		log.Fatal("insert data is failed: ", errExec)
	}

	//fmt.Println(result)

	hasilID, errID := result.LastInsertId()
	hasilRows, errRows := result.RowsAffected()
	if errID != nil || errRows != nil {
		log.Println(errID)
		log.Println(errRows)
	} else {
		fmt.Println("Succesfully insert data. Last inserted ID:", hasilID)
		fmt.Println("Succesfully insert data. Row affected ID:", hasilRows)
	}
	fmt.Println()
}
