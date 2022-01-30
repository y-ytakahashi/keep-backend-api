package infrastructure

// GORMは 生成されるSQLがブラックボックスになるた
// そのため、SQL制御ライブラリは SQLXに変更する

import (
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	"fmt"
	config "keep-back-api/config"

	"github.com/jmoiron/sqlx"
)

type DBService struct{}

var (
	dbConfig config.Config
	db       *sqlx.DB
)

// Init is initialize db from main function
func (DBService) Init() {
	params := dbConfig.ConfigParams()
	fmt.Println(params)
	dsn := params.DBUser + ":" + params.DBPass + "@tcp(" + params.DBHost + ":" + params.DBPort + ")/" + params.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	sqlxDB, sqlxErr := sqlx.Connect("mysql", dsn)

	if sqlxErr != nil {
		panic(sqlxErr)
	}
	db = sqlxDB
}

// GetDB is called in models
func (*DBService) DBSession() *sqlx.DB {
	return db
}
