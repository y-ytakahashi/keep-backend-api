package infrastructure

// GORMは 生成されるSQLがブラックボックスになるた
// そのため、SQL制御ライブラリは SQLXに変更する

import (
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	"fmt"
	config "keep-back-api/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBService struct{}

var dbConfig config.Config // db       *sqlx.DB

var db *sqlx.DB

func (DBService) InitDB() {
	params := dbConfig.ConfigParams()
	fmt.Println(params)
	dsn := params.DBUser + ":" + params.DBPass + "@tcp(" + params.DBHost + ":" + params.DBPort + ")/" + params.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	sqlxDB, sqlxErr := sqlx.Connect("mysql", dsn)

	if sqlxErr != nil {
		fmt.Println("DB Connection Error")
		panic(sqlxErr)
	}

	if sqlxDB.Ping() != nil {
		fmt.Println("DB connection Failed sqlx")
	}

	db = sqlxDB
}

func (DBService) Con() sqlx.DB {
	return *db
}
