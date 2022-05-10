package main

import "github.com/VyacheslavBurnashev/db"

func main() {
	db.Connect("root:root@tcp(127.0.0.1:3306)/database?parseTime=true")
	db.Migrated()

}
