package orm

import (
	"log"
	"time"

	"github.com/kiwisheets/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(cfg *util.DatabaseConfig) *gorm.DB {
	connectionString := constructConnectionString(cfg)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connectionString,
	}), &gorm.Config{
		AllowGlobalUpdate: false,
		Logger:            logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println("Failed to connect to db")
		log.Println(cfg.Host)
		panic(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Println("Failed to connect to db")
		log.Println(cfg.Host)
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(cfg.MaxConnections)
	sqlDB.SetConnMaxLifetime(time.Hour * 1)

	return db
}

func constructConnectionString(dbCfg *util.DatabaseConfig) string {
	return "host=" + dbCfg.Host + " user=" + dbCfg.User + " password=" + dbCfg.Password + " dbname=" + dbCfg.Database + " port=" + dbCfg.Port
}
