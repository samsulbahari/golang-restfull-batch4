package question

import (
	"restfull/internal/domain"
	"restfull/internal/libraries"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuestionService struct {
	db *gorm.DB
}

func QuestionUseCase(dbcon *gorm.DB) *QuestionService {
	return &QuestionService{
		db: dbcon,
	}
}

func (qu QuestionService) CreateQuestionByIdExercises(ctx *gin.Context) {
	id := ctx.Param("id")
	exerciseID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    "400",
			"message": "Invalid input ID",
		})
		return
	}

	//var question domain.Question
	var Question_create domain.Question_Create
	err = ctx.ShouldBindJSON(&Question_create)

	if err != nil {
		validation_response := libraries.Validation(err)
		ctx.JSON(404, gin.H{
			"code":    "404",
			"message": validation_response,
		})
		return
	}

	var exercise domain.Exercise
	err = qu.db.First(&exercise, exerciseID).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    "404",
			"message": "Exercise not found",
		})
		return
	}

	userId := ctx.Request.Context().Value("user_id").(int)
	Question_create.CreatorID = userId
	Question_create.ExerciseID = exerciseID

	err = qu.db.Table("Questions").Create(&Question_create).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    "404",
			"message": "Error Create question",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    "200",
		"message": "Success create question",
	})
	return

}
func (qu QuestionService) GetAnswerQuestionByIdExercises(ctx *gin.Context) {
	exercise_id := ctx.Param("id")
	question_id := ctx.Param("idquestion")

	exerciseID, err := strconv.Atoi(exercise_id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    "400",
			"message": "Invalid input ID Exercise",
		})
		return
	}

	QuestionID, err := strconv.Atoi(question_id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    "400",
			"message": "Invalid input ID Question",
		})
		return
	}

	var ResponseAnswer domain.Question_answer

	err = qu.db.Table("Questions").Where("exercise_id = ? AND id = ?", exerciseID, QuestionID).First(&ResponseAnswer).Error

	if err != nil {
		ctx.JSON(404, gin.H{
			"code":    "404",
			"message": "question not found",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":   "200",
		"Answer": ResponseAnswer.CorrectAnswer,
	})
}
