package db_mng

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func OpenDB(path string) error {
	database, err := sql.Open("sqlite3", path)
	DB = database
	return err
}

func openDBTemp(path string) (*sql.DB, error) {
	database, err := sql.Open("sqlite3", path)
	return database, err
}

func CloseDB() {
	DB.Close()
}

func CloseDBTemp(database *sql.DB) {
	database.Close()
}

func CreateTable() (*sql.Stmt, error) {
	statement, err := DB.Prepare("CREATE TABLE IF NOT EXISTS backup (id integer primary key, filename TEXT, path TEXT, hash TEXT, dateBackup TEXT, size INTEGER)")
	if err != nil {
		return statement, err
	}
	statement.Exec()
	return statement, nil
}

func IsFileAlreadyBackup(path string, hash string, size int64) (bool, error) {

	rows, err := DB.Query("SELECT id from backup where path = ? and hash = ? and size = ?", path, hash, size)

	if !rows.Next() {
		rows.Close()
		return false, err
	}
	rows.Close()
	return true, err
}

func GetFileInDB(hash string) (string, string, error) {
	row := DB.QueryRow("SELECT path, dateBackup from backup where hash = ? limit 1", hash)

	var path string
	var dateBackup string
	err := row.Scan(&path, &dateBackup)
	if err != sql.ErrNoRows {
		return path, dateBackup, nil
	}
	return "", "", errors.New("Error fetching row ")
}

// TODO Inserting operation does not work with passed database connection (crash with large amount of operations), needs to open a new connection each time :(
func AddFile(databasePath string, filename string, path string, hash string, dateBackup string, size int64) (sql.Result, error) {

	var err error

	database1, err := openDBTemp(databasePath)

	statement, err := database1.Prepare("INSERT INTO backup (filename, path, hash, dateBackup, size) values (?,?,?,?,?)")
	if err != nil {
		return nil, err
	}

	res, err := statement.Exec(filename, path, hash, dateBackup,size)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}
	CloseDBTemp(database1)
	return res, err
}

func DeleteFile(filename string, path string, hash string, dateBackup string) (sql.Result, error)  {

	statement, err := DB.Prepare("DELETE FROM backup WHERE filename = ? and path = ? and hash = ? and dateBackup = ?) values (?,?,?,?)")
	if err != nil {
		return nil, err
	}
	return statement.Exec(filename, path, hash, dateBackup)

}

func GetBackedUpDates() []string {
	var dates []string
	var tempDate string

	rows, err := DB.Query("SELECT DISTINCT dateBackup from backup")
	if err != nil{
		return dates
	}

	for rows.Next(){
		rows.Scan(&tempDate)
		dates = append(dates, tempDate)
	}

	return dates
}