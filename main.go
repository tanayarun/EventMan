package main

import (
    "fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "event-booking/models"
)

var db *gorm.DB
var err error

func init() {
    dsn := "host=localhost user=event_user password=your_password dbname=event_booking port=5432 sslmode=disable"
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Auto-migrate the models
    db.AutoMigrate(&models.User{}, &models.Event{}, &models.Booking{})
}

func main() {
    fmt.Println("Database connected and migrations done!")
}
