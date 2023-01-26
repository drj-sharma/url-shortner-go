package database

import (
	"fmt"

	"url-shortner/config"

	log "github.com/drj-sharma/glogger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	logging = log.Logger{}
	logger  = logging.GetLogger()
)

type Database struct {
}

// type Config struct {
// 	Host     string
// 	Port     string
// 	Password string
// 	User     string
// 	DBName   string
// }

func (d *Database) getDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Username, config.Password, config.Hostname, config.DBname)
}

func (d *Database) Connect() (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := d.getDsn()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Debug("trip client: err: ", err)
		return db, err
	}
	return db, nil
}
