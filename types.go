package main

import "time"

type EerrorResponse struct {
	Error string `json:"Error"`
}

type Task struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ProjectID    int64     `json:"projectID"`
	AssignedToID int64     `json:"assignedTo"`
	CreatedAt    time.Time `json:"createdAt"`
}
