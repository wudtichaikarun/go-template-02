package database

import "gorm.io/gorm"

type Database interface {
	Connect(connectionString string, isDebug bool) error
	Close() error
	GetDB() *gorm.DB
}
