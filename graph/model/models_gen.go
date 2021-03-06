// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type CreateActivityInput struct {
	ProjectID int    `json:"projectId"`
	Name      string `json:"name"`
}

type CreateProjectInput struct {
	Name string `json:"name"`
}

type CreateTimesheetInput struct {
	ActivityID int        `json:"activityId"`
	StartedAt  *time.Time `json:"startedAt"`
	EndedAt    *time.Time `json:"endedAt"`
	Duration   int        `json:"duration"`
}

type CreateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserInput struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}
