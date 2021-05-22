package db

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
	"readinglist-backend-api/entity"
)

var (
    db  *gorm.DB
    err error
)

// Init is initialize db from main function
func Init() {
    // dsn := "host=0.0.0.0 port=3306 user=gorm dbname=keep password=gorm sslmode=disable"
    dsn := "gorm:gorm@tcp(127.0.0.1:3306)/keep?charset=utf8mb4&parseTime=True&loc=Local"
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
    return db
}

// Close is closing db
// func Close() {
//     if err := db.clo; err != nil {
//         panic(err)
//     }
// }

func autoMigration() {
    db.AutoMigrate(&entity.User{}, &entity.Leadinglist{}, &entity.Article{})

}