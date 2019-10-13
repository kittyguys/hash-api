package db

import (
	"github.com/jinzhu/gorm"
	"github.com/kittyguys/hash/api/config"
	"github.com/kittyguys/hash/api/model"
)

var db *gorm.DB

// New Connect to MySQL
func New(d *config.Config) *gorm.DB {
	conn, err := gorm.Open("mysql", d.MySQL.User+":"+d.MySQL.Password+"@tcp("+d.MySQL.Host+")/"+d.MySQL.Name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db = conn
	return db
}

// Init Migration
func Init() {
	db.AutoMigrate(&model.User{}, &model.Tag{}, &model.Subtag{}, &model.SocialLogin{})
	return
}

// GetDB for getting db
func GetDB() *gorm.DB {
	return db
}

// func (d *DB) connect() *gorm.DB {
// 	db, err := gorm.Open("mysql", d.Username+":"+d.Password+"@tcp("+d.Host+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

// Create 保存
func Create(value interface{}) *gorm.DB {
	return db.Create(value)
}

// func (db *DB) Exec(sql string, values ...interface{}) *gorm.DB {
// 	return db.Connect.Exec(sql, values...)
// }

// Find 検索
func Find(out interface{}, where ...interface{}) *gorm.DB {
	return db.Find(out, where...)
}

// func (db *DB) First(out interface{}, where ...interface{}) *gorm.DB {
// 	return db.Connect.First(out, where...)
// }

// NewRecord 新しいレコード
func NewRecord(value interface{}) bool {
	return db.NewRecord(value)
}

// func (db *DB) Raw(sql string, values ...interface{}) *gorm.DB {
// 	return db.Connect.Raw(sql, values...)
// }

// Save SAVE
func Save(value interface{}) *gorm.DB {
	return db.Save(value)
}

// func (db *DB) Where(query interface{}, args ...interface{}) *gorm.DB {
// 	return db.Connect.Where(query, args...)
// }
