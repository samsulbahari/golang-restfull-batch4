package main

import (
	"restfull/internal/database"
	"restfull/internal/exercise"
	"restfull/internal/middleware"
	"restfull/internal/question"
	"restfull/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	authorized := r.Group("/")
	db := database.Connetion()

	authorized.Use(middleware.WithAuth())
	{
		//modul exercise
		exercise := exercise.ExerciseUseCase(db)
		authorized.POST("/exercises", exercise.CreateExercise)
		authorized.GET("/exercises/:id", exercise.GetExerciseByid)
		authorized.GET("/exercises/:id/score", exercise.GetAnswerByUser)
		authorized.GET("/exercises/:id/question", exercise.GetAnswerByUser)

		question := question.QuestionUseCase(db)
		authorized.POST("/exercises/:id/questions", question.CreateQuestionByIdExercises)
		authorized.POST("/exercises/:id/questions/:idquestion/answer", question.GetAnswerQuestionByIdExercises)
	}

	//modul user
	user := user.UserUseCase(db)
	r.POST("/register", user.Register)
	r.POST("/login", user.Login)

	r.Run(":1234")
}
