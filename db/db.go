package db

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
	"readinglist-backend-api/entity"
)

var (
    db  *gorm.DB
    err error
)

// Init is initialize db from main function
func Init() {
    dsn := "host=0.0.0.0 port=5432 user=gorm dbname=gorm password=gorm sslmode=disable"
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
    db.AutoMigrate(&entity.User{}, &entity.Leadinglist{})

}