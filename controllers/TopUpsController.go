package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"zidan/AccountServiceAppProject/entities"
)

func FormatCheckController(db *sql.DB, amount string) bool {
	var valid string = "0123456789"

	for _, karakter := range amount {
		if !strings.Contains(valid, string(karakter)) {
			return false
		}
	}

	return true
}

func AddHistoryController(db *sql.DB, user_id uint, amount uint) {
	var newTopUp entities.TopUps

	statement, errPrepare := db.Prepare("insert into topups (topup_id, user_id, topup_amount) VALUES (?, ?, ?)")
	if errPrepare != nil {
		log.Fatal("err prepare: ", errPrepare)
	}

	result, errExec := statement.Exec(&newTopUp.Topup_id, user_id, amount)
	if errExec != nil {
		log.Fatal("insert history topup is failed: ", errExec)
	}

	_, errID := result.LastInsertId()
	_, errRows := result.RowsAffected()
	if errID != nil || errRows != nil {
		log.Println("errID:", errID)
		log.Println("errRows:", errRows)
	}
}

func ReadHistoryTopUpController(db *sql.DB) {
	var allHistory []entities.TopUps
	var rowTopUp entities.TopUps

	// row := db.QueryRow("select accounts.user_id from accounts inner join topups on accounts.user_id = topups.user_id where accounts.user_id = (select user_id from topups where topup_id = (select max(topup_id) from topups)) limit 1")
	// if err := row.Scan(&rowTopUp.User_id); err != nil {
	// 	log.Fatal("cannot read user id. Please be patient: ", err)
	// }

	// fmt.Println("user id: ", rowTopUp.User_id)

	rowFull, errQuery := db.Query("select topups.topup_id , topups.topup_amount, topups.topup_time from accounts inner join topups on accounts.user_id = topups.user_id where accounts.user_id = (select user_id from topups where topup_id = (select max(topup_id) from topups));")
	if errQuery != nil {
		log.Fatal("cannot do select: ", errQuery)
	}

	for rowFull.Next() {
		errScan := rowFull.Scan(&rowTopUp.Topup_id, &rowTopUp.Topup_amount, &rowTopUp.Topup_time)
		if errScan != nil {
			log.Fatal("cannot do scan from row: ", errScan.Error())
		}
		allHistory = append(allHistory, rowTopUp)
	}

	fmt.Println("This is your Top-up History")
	for _, value := range allHistory {
		fmt.Printf("\nTop-Up ID: %v \nTop-Up Amount: %v \nTop-Up Time: %v\n", value.Topup_id, value.Topup_amount, value.Topup_time)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Want to do another transaction?")
	Menu(db)
}
