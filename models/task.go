package models

import "time"

type Task struct {
	ID          int
	Description string
	Created     time.Time
	Done        bool
}
