package config

import (
	"fmt"
	"time"

	"github.com/auliardana/fleetify-test/internal/entity"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
)

func NewDatabase(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	uri := viper.GetString("database.uri")
	idleConnection := viper.GetInt("database.pool.idle")
	maxConnection := viper.GetInt("database.pool.maxconnection")
	maxLifeTimeConnection := viper.GetInt("database.pool.lifetime")

	fmt.Println("uri : ", uri)
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed connect to database: %v", err)
		return nil
	}

	conn, err := db.DB()
	if err != nil {
		log.Fatalf("failed connect to database: %v", err)
	}

	conn.SetMaxIdleConns(idleConnection)
	conn.SetMaxOpenConns(maxConnection)
	conn.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	err = db.AutoMigrate(
		&entity.Employee{},
		&entity.Departement{},
		&entity.AttendanceHistory{},
		&entity.Attendance{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
		return nil
	}

	log.Info("Database migration completed successfully")

	return db

}
