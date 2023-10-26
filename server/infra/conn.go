package infra

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

const (
	driver = "mysql"
	port = "3306"
	net = "tcp"
)

func ConnectDB() (*sql.DB, error) {
	c := mysql.Config{
		User: os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_ROOT_PASSWORD"),
		Net: net,
		Addr: fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), port),
		DBName: os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords: true,
		ParseTime: true,
	}
	
	db, err := sql.Open(driver, c.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}