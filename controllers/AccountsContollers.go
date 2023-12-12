package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"zidan/AccountServiceAppProject/entities"
)

func AddAccountController(db *sql.DB) {
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

func ReadAccountControllerr(db *sql.DB) []entities.Accounts {
	// *** read / select data all_accounts *** //
	var All_accounts []entities.Accounts

	rows, errSelect := db.Query("select * from accounts")
	if errSelect != nil {
		log.Fatal("cannot run select query: ", errSelect)
	}

	for rows.Next() {
		var Row_accounts entities.Accounts
		errScan := rows.Scan(&Row_accounts.User_id, &Row_accounts.Username, &Row_accounts.User_nama, &Row_accounts.User_phone, &Row_accounts.User_email, &Row_accounts.User_address, &Row_accounts.User_balance, &Row_accounts.User_pswd)
		if errScan != nil {
			log.Fatal("cannot run scan query: ", errScan.Error())
		}
		All_accounts = append(All_accounts, Row_accounts)
	}

	return All_accounts
}
