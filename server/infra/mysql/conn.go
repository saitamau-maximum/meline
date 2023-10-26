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

func ConnectDB(host string) (*sql.DB, error) {
	user := os.Getenv("MYSQL_USER")
	if user == "" {
		os.Exit(1)
	} 
	pass := os.Getenv("MYSQL_PASSWORD")
	if pass== "" {
		os.Exit(1)
	}
	dbname := os.Getenv("MYSQL_DATABASE")
	if dbname == "" {
		os.Exit(1)
	}
	

	c := mysql.Config{
		User: user,
		Passwd: pass,
		Net: net,
		Addr: fmt.Sprintf("%s:%s", host, port),
		DBName: dbname,
		AllowNativePasswords: true,
		ParseTime: true,
	}

	db, err := sql.Open(driver, c.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}
