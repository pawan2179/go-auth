package config

import (
	env "auth-go/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
	config := mysql.NewConfig()
	config.User = env.GetString("DB_USER", "root")
	config.Passwd = env.GetString("DB_PASSWORD", "root")
	config.Net = env.GetString("DB_NET", "tcp")
	config.Addr = env.GetString("DB_ADDR", "127.0.0.1:3306")
	config.DBName = env.GetString("DB_NAME", "auth_dev")

	fmt.Println("Connecting to database: ", config.DBName)
	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		fmt.Println("Error in connecting to database: ", err)
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Error in pinging database: ", pingErr)
		return nil, pingErr
	}

	fmt.Println("Connected to DB successfully")
	return db, nil
}
