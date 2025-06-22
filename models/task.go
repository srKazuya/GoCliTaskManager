package models

import (
	"time"
)

type Task struct {
	Id          int `json:"id"`
	Description string `json:"description"`
	Status      string
	CreatedAt   time.Time
	UpdatedAt time.Time
}



func NewTask(desc string) Task{
	return Task{
		Id: 0,
		Description: desc, 
		Status: "В работе",  
		CreatedAt: time.Now(), 
		UpdatedAt: time.Now(),
	}
}

