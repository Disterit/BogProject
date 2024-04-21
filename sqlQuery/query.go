package sqlQuery

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func connection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/practica")
	if err != nil {
		panic(err.Error())
	}

	// Проверяем соединение
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}

func createUser(userID int64, userName string) {
	db := connection()
	defer db.Close()

	_, err := db.Exec("INSERT INTO user (user_id, user_name) VALUES (?, ?)", userID, userName)
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("INSERT INTO permissions (user_id) VALUES (?)", userID)
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("INSERT INTO personality (user_id) VALUES (?)", userID)
	if err != nil {
		panic(err.Error())
	}
}

func CheckUser(userID int64, userName string) bool {
	db := connection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user where user_id = ?", userID)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		return true
	} else {
		createUser(userID, userName)
		return false
	}
}

func CheckAgreement(userID int64) bool {
	db := connection()
	defer db.Close()

	// Выполняем SQL-запрос для обновления значения
	rows, err := db.Query("SELECT * FROM permissions where user_id = ?", userID)
	if err != nil {
		if err != nil {
			panic(err.Error())
		}
	}

	for rows.Next() {
		var blocked, agreement bool
		var user_id int64
		if err := rows.Scan(&blocked, &agreement, &user_id); err != nil {
			panic(err.Error())
		}
		return agreement
	}
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
	return false
}

func UserCurrency(userID int64) string {
	db := connection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM personality where user_id = ?", userID)
	if err != nil {
		if err != nil {
			panic(err.Error())
		}
	}

	for rows.Next() {
		var user_id int64
		var email, passport, currency, FIO string
		if err := rows.Scan(&user_id, &email, &passport, &currency, &FIO); err != nil {
			panic(err.Error())
		}
		return currency
	}
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
	return ""
}

func UpdateUserCurrency(userID int64, newCurrency string) {
	db := connection()
	defer db.Close()

	_, err := db.Exec("UPDATE personality SET currency = ? WHERE user_id = ?", newCurrency, userID)
	if err != nil {
		panic(err.Error())
	}
}

func UpdateUserPassport(userID int64, newPassport string) {
	db := connection()
	defer db.Close()

	_, err := db.Exec("UPDATE personality SET passport = ? WHERE user_id = ?", newPassport, userID)
	if err != nil {
		panic(err.Error())
	}
}

func UpdateUserEmail(userID int64, newEmail string) {
	db := connection()
	defer db.Close()

	_, err := db.Exec("UPDATE personality SET email = ? WHERE user_id = ?", newEmail, userID)
	if err != nil {
		panic(err.Error())
	}
}

func UpdateUserFIO(userID int64, newFIO string) {
	db := connection()
	defer db.Close()

	_, err := db.Exec("UPDATE personality SET FIO = ? WHERE user_id = ?", newFIO, userID)
	if err != nil {
		panic(err.Error())
	}
}

func GetAccountInfo(userID int64) (string, string, string, string) {
	db := connection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM personality where user_id = ?", userID)
	if err != nil {
		if err != nil {
			panic(err.Error())
		}
	}

	for rows.Next() {
		var user_id int64
		var email, passport, currency, FIO string
		if err := rows.Scan(&user_id, &email, &passport, &currency, &FIO); err != nil {
			panic(err.Error())
		}
		return email, passport, currency, FIO
	}
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
	return "", "", "", ""
}

func ChangeAgreemant(userID int64) {
	db := connection()
	defer db.Close()

	OK := true

	// Выполняем SQL-запрос для обновления значения
	_, err := db.Exec("UPDATE permissions SET agreement = ? WHERE user_id = ?", OK, userID)
	if err != nil {
		panic(err.Error())
	}
}

func CheckBlocked(userId int64) bool {
	return true
}
