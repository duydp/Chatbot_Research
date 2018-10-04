package db

import (
	"database/sql"
	"fmt"

	"BVGI/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init(cfg *config.DBConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	return nil
}