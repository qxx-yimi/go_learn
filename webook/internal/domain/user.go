package domain

import "time"

// User领域对象
type User struct {
	Id              int64
	Email           string
	Password        string
	Ctime           time.Time
	Nickname        string
	Birthday        string
	PersonalProfile string
}
