package dto

import (
	"time"
)

//DataUpdateDTO is a model that client use when updating data from table
type DataUpdateDTO struct {
	ID        int64  `json:"id,string,omitempty"`
	Title     string `json:"title"`
	CreatedAt time.Time
}

//DataCreateDTO is is a model that client use when create a new data
type DataCreateDTO struct {
	ID        int64  `json:"id,string,omitempty"`
	Title     string `json:"title"`
	CreatedAt time.Time
}
