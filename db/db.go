package db

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n===================================================\n", sql)
}

func New() *gorm.DB {

	dsn := "root:boatboat224@tcp(127.0.0.1:3306)/test?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})
	if err != nil {
		fmt.Println("db error", err)
		panic(err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	// db.AutoMigrate(repository.Customer{})
}
