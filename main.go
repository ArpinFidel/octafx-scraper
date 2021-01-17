package main

import (
	"log"
	"os"

	"github.com/ArpinFidel/octafx-scraper/external"
	"github.com/ArpinFidel/octafx-scraper/model"
	"github.com/ArpinFidel/octafx-scraper/repo"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	os.Remove("db/octa.db")

	log.Println("[init] creating octa.db...")
	file, err := os.Create("db/octa.db")
	if err != nil {
		log.Fatalf("[init] create db file: %v", err.Error())
	}
	file.Close()
	log.Println("[init] octa.db created")

	db, _ := gorm.Open(sqlite.Open("./octa.db"), &gorm.Config{})
	repo.NewMasterSQL(db)

	m := external.NewMaster("https://www.octafx.com/copy-trade/master")
	m.Refresh(model.Master{Code: "11467328"})
}
