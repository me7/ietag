package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Equipment struct {
	gorm.Model
	Desc     string
	Brand    string
	Serial   string
	Type     string
	Cus      string
	Ietag    string `gorm:"unique"`
	Loc      string
	Cost     float32
	Currency string
	Invioce  string
	Invref   string
	Receive  time.Time
	Remark   string
}

func main() {
	var err error
	var equipments []Equipment
	// var equipment Equipment

	db, err := gorm.Open(sqlite.Open("ietag.db"))

	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&Equipment{})

	result := db.Find(&equipments)
	// fmt.Printf("%v", result)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	for _, v := range equipments {
		fmt.Println(v.Desc, v.Brand)
	}

	// db.Create(&Equipment{Brand: "Samsung", Ietag: "IE101", Cus: "SST", Cost: 119, Receive: time.Now()})
	// db.First(&eq1, "desc=?", "CPU")
	// db.Raw("select * from equipment where desc=?", "DB Test Fixture").Scan(&eq1)
	// fmt.Println(eq1)
	// db.Delete(&Equipment{}, "Ietag=?", "IE101")
}
