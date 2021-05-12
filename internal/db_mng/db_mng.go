package db_mng

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(path string) (*sql.DB,error) {
	database,err := sql.Open("sqlite3",path)
	return database,err
}

func CloseDB(database sql.DB)  {
	database.Close()
}

func CreateTable(database sql.DB) (*sql.Stmt, error) {
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS backup (id integer primary key, filename TEXT, path TEXT, hash TEXT, dateBackup TEXT)")
	if err!= nil {
		return statement, err
	}
	statement.Exec()
	return statement, nil
}

func IsFileAlreadyBackup(database sql.DB, path string, hash string) (bool, error) {

	rows, err := database.Query("SELECT id from backup where path = ? and hash = ?",path, hash)

	if !rows.Next() {
		rows.Close()
		return false, err
	}
	rows.Close()
	return true, err
}

func GetFileInDB(database sql.DB, hash string) (string, string, error) {
	row := database.QueryRow("SELECT path, dateBackup from backup where hash = ? limit 1", hash)

	var path string
	var dateBackup string
	err := row.Scan(&path, &dateBackup)
	if err != sql.ErrNoRows {
		return path, dateBackup,nil
	}
	return "", "", errors.New("Error fetching row ")
}

func AddFile(database sql.DB, filename string, path string, hash string, dateBackup string) (sql.Result, error) {

	statement, err := database.Prepare("INSERT INTO backup (filename, path, hash, dateBackup) values (?,?,?,?)")
	if err != nil{
		return nil, err
	}
	return statement.Exec(filename, path, hash, dateBackup)
}