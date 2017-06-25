package database

import (
	"fmt"

	"github.com/go-shosa/shosa/log"
	"github.com/jinzhu/gorm"
	// gormの関連パッケージ
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBConf is database config
type DBConf struct {
	DB   string `env:"DB_TYPE"`
	User string `env:"DB_USER"`
	Pass string `env:"DB_PASS"`
	IP   string `env:"DB_IP"`
	Port string `env:"DB_PORT"`
}

// Connect returns database connection
func Connect(isTests ...bool) (*gorm.DB, error) {
	// set config
	conf := DBConf{
		DB:   "mysql",
		User: "v6jym_user",
		Pass: "!QAZse4rfvgy7y",
		IP:   "public.v6jym.tyo1.database-hosting.conoha.io",
		Port: "3306",
	}

	dbName := "v6jym_suggesta"

	// connect database
	db, err := gorm.Open(conf.DB, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.User, conf.Pass, conf.IP, conf.Port, dbName))
	log.Debugf("try connect database ( %s:%s@tcp(%s:%s)/%s?parseTime=true )", conf.User, conf.Pass, conf.IP, conf.Port, dbName)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Open returns database connection
// If Error occurs, not return error, output error message to std.Output
func Open(isTests ...bool) *gorm.DB {
	db, err := Connect(isTests...)
	if err == nil {
		return db
	}
	log.Printf("Can't connect database. %v", err)
	return nil
}

// Close closes current database connection
func Close(d *gorm.DB) error {
	return d.Close()
}
