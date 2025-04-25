package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type account struct {
	Id       int
	Username *string
	Email    *string
	Active   *bool
}

func newDBConn() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "gitm.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initDB() error {
	db, err := newDBConn()
	if err != nil {
		return err
	}

	query := `CREATE TABLE IF NOT EXISTS account (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL,
		active INT NOT NULL DEFAULT(0)
	);`

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	defer db.Close()

	fmt.Println("Gitm initialized successfully")
	return nil
}

func viewAccounts() ([]account, error) {
	db, err := newDBConn()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var accounts []account

	query := `SELECT * FROM account;`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var acc account

		if err := rows.Scan(&acc.Id, &acc.Username, &acc.Email, &acc.Active); err != nil {
			return nil, err
		}

		accounts = append(accounts, acc)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func viewAccountByID(id int) (account, error) {
	var account account

	db, err := newDBConn()
	if err != nil {
		return account, err
	}

	defer db.Close()

	query := `SELECT * FROM account WHERE id = ?`

	err = db.QueryRow(query, id).Scan(&account.Id, &account.Username, &account.Email, &account.Active)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return account, ErrAccountNotFound
		}

		return account, err
	}

	return account, nil
}

func getActiveAccount() (account, error) {
	var account account

	db, err := newDBConn()
	if err != nil {
		return account, nil
	}

	query := `SELECT * FROM account WHERE active = 1`

	err = db.QueryRow(query).Scan(&account.Id, &account.Username, &account.Email, &account.Active)
	if err != nil {
		return account, err
	}

	return account, nil
}

func addNewAccount(account account) error {
	db, err := newDBConn()
	if err != nil {
		return err
	}

	defer db.Close()

	query := `INSERT INTO account (username, email) VALUES (?, ?);`

	_, err = db.Exec(query, account.Username, account.Email)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func updateAccount(id int, account account) error {
	db, err := newDBConn()
	if err != nil {
		return err
	}

	acc, err := viewAccountByID(id)
	if err != nil {
		if errors.Is(err, ErrAccountNotFound) {
			return ErrAccountNotFound
		}

		return err
	}

	if account.Username == nil {
		account.Username = acc.Username
	}

	if account.Email == nil {
		account.Email = acc.Email
	}

	if account.Active == nil {
		account.Active = acc.Active
	}

	query := `UPDATE account SET username = ?, email = ?, active = ? WHERE id = ?`

	_, err = db.Exec(query, account.Username, account.Email, account.Active, id)
	if err != nil {
		return err
	}

	return nil
}

func deleteAccount(id int) error {
	db, err := newDBConn()
	if err != nil {
		return err
	}

	query := `DELETE FROM account WHERE id = ?`

	_, err = db.Exec(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrAccountNotFound
		}

		return err
	}

	return nil
}
