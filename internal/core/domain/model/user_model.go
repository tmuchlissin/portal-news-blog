package model

import "time"

// User represents one user entity in Go code.
// The struct name uses singular form because one struct value represents one row,
// while the database table uses plural form (`users`) because it stores many rows.
// The fields use capital letters so they are exported and can be accessed by GORM
// and other packages; the gorm tags map them to snake_case database column names.
type User struct {
	ID        int64      `gorm:"id"`
	Name      string     `gorm:"name"`
	Email     string     `gorm:"email"`
	Password  string     `gorm:"password"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at"` // why use pointer (*) here? Because the value can be nullable.
}
