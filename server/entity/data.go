package entity

import (
	"time"
)

//Data struct represents data table in database
type Data struct {
	ID        int64  `json:"id,string,omitempty"`
	Title     string ` json:"title"`
	CreatedAt time.Time
}
