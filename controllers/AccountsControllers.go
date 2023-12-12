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
	fmt.Println("Want to do another transaction?")
	Menu(db)
}

func ReadAllAccountsController(db *sql.DB) []entities.Accounts {
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

func ReadMyAccountController(db *sql.DB) {
	// *** read / select data my account *** //
	var my_account entities.Accounts

	rowID := db.QueryRow("select accounts.user_id from accounts inner join login on accounts.user_id = login.user_id where accounts.user_id = (select user_id from login where login_id = (select max(login_id) from login)) limit 1;")

	if err := rowID.Scan(&my_account.User_id); err != nil {
		//
	}

	fmt.Println(my_account.User_id)

	// // query select
	rowFull := db.QueryRow("select * from accounts where user_id = ?", my_account.User_id)

	if err := rowFull.Scan(&my_account.User_id, &my_account.Username, &my_account.User_nama, &my_account.User_phone, &my_account.User_email, &my_account.User_address, &my_account.User_balance, &my_account.User_pswd); err != nil {
		if err == sql.ErrNoRows {
			log.Fatal(err)
		}
	}

	fmt.Println("This is your profile")
	fmt.Printf("\nUser ID: %v \nUsername: %v \nDisplay Name: %v \nPhone: %v \nEmail: %v \nAddress: %v \nBalance: %v \nPassword: %v\n", my_account.User_id, my_account.Username, my_account.User_nama, my_account.User_phone, my_account.User_email, my_account.User_address, my_account.User_balance, my_account.User_pswd)
	fmt.Println()
	fmt.Println("Want to do another transaction?")
	Menu(db)
}
