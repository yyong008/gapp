package db

import (
	"gapp1/internal/app/model"
	"log"
)

// Migrate 执行数据库迁移
func Migrate() {
	if err := Database.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration completed successfully")
}
