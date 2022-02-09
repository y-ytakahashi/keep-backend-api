package main

// 個別で作成した関数を importする場合、gomod で設定したカレントからのパス指定が必要
import (
	"fmt"
	db "keep-back-api/infrastructure/db"
	"keep-back-api/server"
)

var DB db.DBService

func main() {
	DB.InitDB()
	if DB.Con().DB.Ping() == nil {
		fmt.Println("DB Connection Success")
	}
	server.Init()
}
