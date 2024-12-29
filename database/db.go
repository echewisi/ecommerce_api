package database

import (
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/echewisi/ecommerce_api/models"
)

func ConnectDB(dsn string) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Auto-migrate the models
    err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
    if err != nil {
        log.Fatalf("Error migrating database: %v", err)
    }

    log.Println("Database connection and migration established")
    return db, nil
}
