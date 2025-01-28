package models

import "time"

type Reminder struct {
	UserEmail    string `json:"user_email"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ReminderTime time.Time `json:"reminder_time"`
}
