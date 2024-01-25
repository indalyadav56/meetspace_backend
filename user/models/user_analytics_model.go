package models

import "gorm.io/gorm"

type UserAnalytics struct {
    gorm.Model
    UserID          uint   `gorm:"not null;uniqueIndex"`
    PageViews       int    `gorm:"default:0"`
    TimeOnSite      int    `gorm:"default:0"`
    ActionsTaken    int    `gorm:"default:0"`
}
