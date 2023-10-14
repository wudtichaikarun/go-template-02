package database_mysql

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logGorm "gorm.io/gorm/logger"
)

type MySQLDatabase struct {
	db *gorm.DB
}

func setupLoggerMode(isDebug bool) logGorm.Interface {
	if isDebug {
		return logGorm.Default.LogMode(logGorm.Info)
	}
	return logGorm.Default.LogMode(logGorm.Silent)
}

func (m *MySQLDatabase) Connect(connectionString string, isDebug bool) error {
	configDb := &gorm.Config{
		Logger:  setupLoggerMode(isDebug),
		NowFunc: func() time.Time { return time.Now().UTC() },
	}

	db, err := gorm.Open(mysql.Open(connectionString), configDb)
	if err != nil {
		return err
	}

	sqlDB, _ := db.DB()
	if err != nil {
		return err
	}

	// Connection Pool Configs
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Assign the connected database instance to the 'db' field
	m.db = db

	return nil
}

func (m *MySQLDatabase) Close() error {
	if m.db != nil {
		sqlDB, _ := m.db.DB()
		sqlDB.Close()
	}

	return nil
}

func (m *MySQLDatabase) GetDB() *gorm.DB {
	return m.db
}
