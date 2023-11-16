package mysql

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/saitamau-maximum/meline/config"
)

const (
	DRIVER = "mysql"
	PORT   = "3306"
	NET    = "tcp"
)

func ConnectDB(host string) (*sql.DB, error) {
	c := mysql.Config{
		User:                 config.MYSQL_USER,
		Passwd:               config.MYSQL_PASSWORD,
		Net:                  NET,
		Addr:                 fmt.Sprintf("%s:%s", host, PORT),
		DBName:               config.MYSQL_DATABASE,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open(DRIVER, c.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}
