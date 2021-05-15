package main

// 個別で作成した関数を importする場合、gomod で設定したカレントからのパス指定が必要
import (
	"readinglist-backend-api/db"
	"readinglist-backend-api/server"
)


func main() {
		db.Init()
		server.Init()
}
