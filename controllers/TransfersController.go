package controllers

import (
	"database/sql"
	"fmt"
	"zidan/AccountServiceAppProject/entities"
)

func DoTransferController(db *sql.DB) {
	var akun1 entities.Accounts
	var akun2 entities.Accounts
	//var newTransfer entities.Transfers
	var phone_target string
	var my_phone string
	var amount uint

	fmt.Println("Masukkan nomor anda:")
	fmt.Scanln(&my_phone)
	fmt.Println("Masukkan nomor tujuan transfer:")
	fmt.Scanln(&phone_target)
	fmt.Println("Masukkan jumlah transfer:")
	fmt.Scanln(&amount)

	//cek jumlah amount pada balance. Apa cukup untuk transfer?
	row := db.QueryRow("select user_phone, user_balance from accounts where user_phone = ?", my_phone)
	if err := row.Scan(&akun1.User_phone, &akun1.User_balance); err != nil {
		fmt.Println("cannot read balance", err)
	}

	if amount <= akun1.User_balance {
		row := db.QueryRow("select user_phone, user_balance from accounts where user_phone = ?", phone_target)
		if err := row.Scan(&akun2.User_phone, &akun2.User_balance); err != nil {
			fmt.Println("cannot read balance", err)
		}

		balance2 := akun2.User_balance + amount
		fmt.Println(balance2)
		balance1 := akun1.User_balance - amount
		fmt.Println(balance1)

		result, errExec := db.Exec("UPDATE accounts SET user_balance = ? WHERE user_balance = ? and user_phone = ?", balance2, akun2.User_balance, akun2.User_phone)
		if errExec != nil {
			fmt.Println("saldo target tidak bisa update", errExec)
		}

		resultRow, _ := result.RowsAffected()
		fmt.Println("resultRow:", resultRow)

		result2, errExec2 := db.Exec("UPDATE accounts SET user_balance = ? WHERE user_balance = ? and user_phone = ?", balance1, akun1.User_balance, akun1.User_phone)
		if errExec2 != nil {
			fmt.Println("saldo anda tidak bisa diupdate", errExec2)
		}

		resultRow2, _ := result2.RowsAffected()
		fmt.Println("resultRow2:", resultRow2)

		fmt.Println()
		fmt.Println("Want to do another transaction?")
		Menu(db)
	} else if amount > akun1.User_balance {
		fmt.Println("Your balance is not enough. Please top up.")
		fmt.Println()
		fmt.Println("Want to do another transaction?")
		Menu(db)
	}

}
