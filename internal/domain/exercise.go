package domain

import "time"

type Exercise struct {
	ID          int
	Title       string `binding:"required"`
	Description string `binding:"required"`
	Questions   []Question
}

type Question struct {
	ID            int
	ExerciseID    int    `json:"-"`
	Body          string `binding:"required"`
	OptionA       string `binding:"required"`
	OptionB       string `binding:"required"`
	OptionC       string `binding:"required"`
	OptionD       string `binding:"required"`
	CorrectAnswer string `json:"-"`
	Score         int
	CreatorID     int `json:"-"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Answer struct {
	ID         int
	ExerciseID int
	QuestionID int
	UserID     int
	Answer     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Questions  Question `gorm:"foreignKey:QuestionID;references:ID"`
}
