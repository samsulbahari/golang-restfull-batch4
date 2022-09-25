package user

import (
	"fmt"
	"restfull/internal/domain"
	"restfull/internal/libraries"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func UserUseCase(dbcon *gorm.DB) *UserService {
	return &UserService{
		db: dbcon,
	}
}

func (uu UserService) Register(ctx *gin.Context) {

	var user domain.User

	//validasi
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		validation_response := libraries.Validation(err)
		ctx.JSON(404, gin.H{
			"code":    "404",
			"message": validation_response,
		})
		return
	}

	//check username all ready
	check := uu.CheckEmail(user.Email)

	if check == true {
		ctx.JSON(200, gin.H{
			"code":    "200",
			"message": "Email already use",
		})
		return
	}
	//create password
	err = user.CreatePassword(user.Password)
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    "500",
			"message": "Server Error create password",
		})
		return
	}
	//create user
	err = uu.db.Create(&user).Error

	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    "500",
			"message": "Server create user",
		})
		return
	}

	//generate token
	stringtoken, err := user.GenerateJwt()
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    "500",
			"message": "Server generate jwt error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    "200",
		"message": "Success Register",
		"token":   stringtoken,
	})

}

func (uu UserService) CheckEmail(email string) bool {
	var user domain.User
	err := uu.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return false
	}
	return true

}

func (uu UserService) Login(ctx *gin.Context) {
	var user domain.User
	//change struct name menjadi tidak required
	var Login_struct domain.Login = domain.Login(*&user)

	err := ctx.ShouldBind(&Login_struct)
	if err != nil {
		message := libraries.Validation(err)
		ctx.JSON(404, gin.H{
			"code":    "404",
			"message": message,
		})
		return
	}

	err = uu.db.Where("email = ?", Login_struct.Email).Take(&user).Error

	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "Invalid email / password",
		})
		return
	}

	err = user.ComparePassword(Login_struct.Password)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "Invalid email / password",
		})
		return
	}

	token, err := user.GenerateJwt()
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    "500",
			"message": "Server generate jwt error",
		})
		return
	}

	fmt.Println(Login_struct)

	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "Success Login",
		"token":   token,
	})

}
