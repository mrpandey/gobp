package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mrpandey/gobp/src/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

func NewGormPostgresClient(
	_ context.Context,
	cfg *util.Config,
	logger *util.StandardLogger,
	gormLogLevel glogger.LogLevel,
) *gorm.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		cfg.PGHost, cfg.PGPort, cfg.PGUser, cfg.PGPass, cfg.PGDB,
	)

	gormCfg := &gorm.Config{
		TranslateError: true, // translate db error to gorm error e.g. gorm.ErrRecordNotFound
		Logger:         glogger.Default.LogMode(gormLogLevel),
	}

	db, err := gorm.Open(
		postgres.Open(connStr),
		gormCfg,
	)

	if err != nil {
		logger.Panic(err)
	}

	sqlDB := GetGormConnPool(logger, db)
	sqlDB.SetMaxIdleConns(int(cfg.SqlIdleConnectionsCount))
	sqlDB.SetMaxOpenConns(int(cfg.SqlMaxConnectionsCount))
	sqlDB.SetConnMaxLifetime(cfg.SqlConnectionLifetime)

	logger.Info("Postgres initialized.")

	return db
}

func GetPgDbUri(cfg *util.Config) string {
	pgDbUri := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v",
		cfg.PGUser,
		cfg.PGPass,
		cfg.PGHost,
		cfg.PGPort,
		cfg.PGDB,
	)
	return pgDbUri
}

func GetGormConnPool(logger *util.StandardLogger, db *gorm.DB) *sql.DB {
	sqlDB, err := db.DB()
	if err != nil {
		logger.Panic(err)
	}
	return sqlDB
}

func CreatePgDB(ctx context.Context, logger *util.StandardLogger, db *gorm.DB, dbName string) {
	sqlCmd := fmt.Sprintf("create database %s", dbName)
	err := db.WithContext(ctx).Exec(sqlCmd).Error
	if err != nil {
		logger.Panic(err)
	}
	logger.Infof("Created database %s.", dbName)
}

func DropPgDB(ctx context.Context, logger *util.StandardLogger, db *gorm.DB, dbName string) {
	sqlCmd := fmt.Sprintf("drop database if exists %s", dbName)
	err := db.WithContext(ctx).Exec(sqlCmd).Error
	if err != nil {
		logger.Panic(err)
	}
	logger.Infof("Dropped database %s.", dbName)
}

func GetDefaultGormLogLevel(dgn util.DGNType) glogger.LogLevel {
	gormLogLevel := glogger.Silent
	if dgn == util.Local {
		gormLogLevel = glogger.Info
	} else if dgn == util.Dev {
		gormLogLevel = glogger.Warn
	}
	return gormLogLevel
}
