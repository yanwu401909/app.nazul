package repo

import (
	"fmt"
	"log"

	. "app.nazul/config"
	"app.nazul/models"
	"github.com/jinzhu/gorm"
)

var connUrl = "%s:%s@tcp(%s:%d)/%s?charset=%s&autocommit=%s&parseTime=true"
var CONN *gorm.DB

func init() {
	var err error
	CONN, err = gorm.Open("mysql", fmt.Sprintf(connUrl, CONFIG.Database.User, CONFIG.Database.Pass, CONFIG.Database.Host, CONFIG.Database.Port, CONFIG.Database.Name, CONFIG.Database.Charset, CONFIG.Database.Autocommit))
	// CONN.LogMode(true)
	log.Printf("DB init....%v", err == nil)
	if err != nil {
		log.Fatal(err)
	}
	CONN.DB().SetMaxIdleConns(CONFIG.Database.MaxIdleConns)
	CONN.DB().SetMaxOpenConns(CONFIG.Database.MaxOpenConns)
	if !CONN.HasTable(&models.User{}) {
		if err := CONN.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&models.User{}).Error; err != nil {
			log.Fatal(err)
		}
	}
	if !CONN.HasTable(&models.RsaPair{}) {
		if err := CONN.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&models.RsaPair{}).Error; err != nil {
			log.Fatal(err)
		}
	}
}
