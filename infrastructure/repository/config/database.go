package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type Repository struct {
	DB *gorm.DB
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
}

func GormOpen() (gormDb *gorm.DB, err error) {
	var infoDatabase InfoDatabaseSQL
	err = infoDatabase.getDriverConn("Databases.MySQL.BoilerplateGo")
	if err != nil {
		return nil, err
	}

	gormDb, err = gorm.Open(mysql.Open(infoDatabase.Write.DriverConn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return
	}

	dialector := mysql.New(mysql.Config{
		DSN: infoDatabase.Read.DriverConn,
	})

	err = gormDb.Use(dbresolver.Register(
		dbresolver.Config{
			Replicas: []gorm.Dialector{dialector},
		},
	))
	if err != nil {
		return
	}

	var result int
	if err = gormDb.Raw("SELECT 1").Scan(&result).Error; err != nil {
		return nil, err
	}

	return
}
