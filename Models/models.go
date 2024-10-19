package models

import (
    "gorm.io/gorm"
)


type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Password string
}


type Event struct {
    gorm.Model
    Name          string
    Description   string
    TotalSeats    int
    ReservedSeats int
}


type Booking struct {
    gorm.Model
    UserID  uint
    EventID uint
    Seats   int
}
