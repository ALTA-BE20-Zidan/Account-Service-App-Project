package controllers

import (
	"database/sql"
	"strings"
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
