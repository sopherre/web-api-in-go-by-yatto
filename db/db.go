package db

import (
	"fmt"
	"os"
	"web-api-in-go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	GormDb *gorm.DB
}

func Init() *Database {
	// app.envに定義したDB関係の環境変数を取得
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	// tcp（）の中にdocker-composeで定義したDB用コンテナのサービス名を入れれば、
	// 自動的にホストとポートを読み取ってくれる
	dsn := fmt.Sprintf(
		"%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		dbUser,
		dbPassword,
		dbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database.")
	}
	db.AutoMigrate(&models.Task{})
	TaskSeeder(db)
	return &Database{db}
}

func (db *Database) Close() {
	sqlDB, _ := db.GormDb.DB()
	sqlDB.Close()
}

func TaskSeeder(db *gorm.DB) {
	tasks := []models.Task{
		{Title: "セクション1学習", Completed: false},
		{Title: "セクション2学習", Completed: false},
		{Title: "セクション3学習", Completed: false},
	}
	for _, task := range tasks {
		t := task
		db.FirstOrCreate(&t, models.Task{Title: t.Title})
	}
}
