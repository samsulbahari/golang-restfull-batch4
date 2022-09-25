package exercise

import (
	"net/http"
	"strconv"
	"strings"

	"restfull/internal/domain"
	"restfull/internal/libraries"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExerciseSevice struct {
	db *gorm.DB
}

func ExerciseUseCase(dbcon *gorm.DB) *ExerciseSevice {
	return &ExerciseSevice{
		db: dbcon,
	}
}

func (eu ExerciseSevice) CreateExercise(ctx *gin.Context) {
	var exercise domain.Exercise
	err := ctx.ShouldBindJSON(&exercise)
	if err != nil {
		validation_response := libraries.Validation(err)
		ctx.JSON(404, gin.H{
			"code":    "404",
			"message": validation_response,
		})
	}
	err = eu.db.Create(&exercise).Error
	if err != nil {
		ctx.JSON(404, gin.H{
			"code":    "500",
			"message": "error create exercises",
		})
	}

	ctx.JSON(200, gin.H{
		"code":    "200",
		"message": "Succes create Exercises",
		"data":    exercise,
	})
}

func (eu ExerciseSevice) GetExerciseByid(ctx *gin.Context) {

	id := ctx.Param("id")
	exerciseID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"code":    "404",
			"message": "Invalid input ID",
		})
		return
	}

	var exercise domain.Exercise
	err = eu.db.Preload("Questions").First(&exercise, exerciseID).Error
	if err != nil {
		ctx.JSON(404, gin.H{
			"code":    "404",
			"message": "Exercise not found",
		})
		return

	}
	//hide correct ansewr

	ctx.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "Succes get data",
		"data":    exercise,
	})
}

func (eu ExerciseSevice) GetAnswerByUser(ctx *gin.Context) {
	id := ctx.Param("id")
	exerciseID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"code":    "404",
			"message": "Invalid input ID",
		})
		return
	}

	var exercise domain.Exercise

	err = eu.db.Preload("Questions").First(&exercise, exerciseID).Error

	if err != nil {
		ctx.JSON(404, gin.H{
			"code":    "404",
			"message": "Exercise not found",
		})
		return
	}
	userId := ctx.Request.Context().Value("user_id").(int)

	var answer []domain.Answer
	err = eu.db.Where("exercise_id = ? AND user_id = ?", exerciseID, userId).Preload("Questions").Find(&answer).Error

	if err != nil || len(answer) == 0 {
		ctx.JSON(200, gin.H{
			"code":  "200",
			"score": 0,
		})
		return
	}

	var score int
	for _, answers := range answer {
		if strings.EqualFold(answers.Answer, answers.Questions.CorrectAnswer) {
			score += answers.Questions.Score
		}
	}

	ctx.JSON(200, gin.H{
		"code":  "200",
		"score": score,
	})
}
