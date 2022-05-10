package main

import "github.com/VyacheslavBurnashev/db"

func main() {
	db.Connect("root:root@tcp(localhost:3333)/database?parseTime=true")
	db.Migrated()
}
