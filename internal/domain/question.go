package domain

import "time"

type Question_Create struct {
	ID            int
	ExerciseID    int    `json:"-"`
	Body          string `binding:"required"`
	OptionA       string `binding:"required"`
	OptionB       string `binding:"required"`
	OptionC       string `binding:"required"`
	OptionD       string `binding:"required"`
	CorrectAnswer string
	Score         int
	CreatorID     int `json:"-"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Question_answer struct {
	CorrectAnswer string
}
