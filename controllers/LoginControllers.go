package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"zidan/AccountServiceAppProject/entities"
)

func AddLoginController(db *sql.DB, user_id uint) {
	var newLogin entities.Login

	statement, errPrepare := db.Prepare("insert into login (login_id, user_id, login_time) VALUES (?, ?, ?)")
	if errPrepare != nil {
		log.Fatal("err prepare: ", errPrepare)
	}

	result, errExec := statement.Exec(&newLogin.Login_id, user_id, &newLogin.Login_time)
	if errExec != nil {
		log.Fatal("insert data is failed: ", errExec)
	}

	_, errID := result.LastInsertId()
	_, errRows := result.RowsAffected()
	if errID != nil || errRows != nil {
		log.Println("errID:", errID)
		log.Println("errRows:", errRows)
	}
	// else {
	// 	fmt.Println("Succesfully insert data. Last inserted ID:", hasilID)
	// 	fmt.Println("Succesfully insert data. Row affected ID:", hasilRows)
	// }
}

func LoginVerificationController(db *sql.DB) {
	var phone_number string
	var password string

	fmt.Println("Masukkan No. Telpon:")
	fmt.Scanln(&phone_number)
	fmt.Println("Masukkan Password:")
	fmt.Scanln(&password)

	result := ReadAccountControllerr(db)

	for _, value := range result {
		if phone_number == value.User_phone && password == value.User_pswd {
			fmt.Println("Silakan Menikmati Layanan Kami")
			AddLoginController(db, value.User_id)
		} else {
			fmt.Println("Anda belum terdaftar")
		}
	}
}
