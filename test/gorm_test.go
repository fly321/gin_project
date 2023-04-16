package test

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type User struct {
	gorm.Model
	Name string
	Age  int64
}

func TestGetViper(t *testing.T) {
	viper.SetConfigFile("../config/app.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		t.Error(err)
		return
	}
	//log.Fatalln(viper.Get("database.main_mysql"))
	//fmt.Println(viper.Get("database.main_mysql"))
}

func TestGorm(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:3307)/tp6?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
		return
	}
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "Jinzhu", Age: 18})
	db.Create(&User{Name: "Jinzhu 2", Age: 20})
}
