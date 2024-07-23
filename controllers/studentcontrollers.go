package controllers

import (
	"net/http"
	"os"
	"time"

	initializers "github.com/Shreedeepdatta/rankandmarks/Initializers"
	"github.com/Shreedeepdatta/rankandmarks/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	var body struct {
		Name     string
		Roll     int16
		Class    string
		Password string
	}
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "could not read body",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "could not hash password",
		})
		return
	}
	student := models.Student{Name: body.Name, Roll: body.Roll, Class: body.Class, Password: string(hash)}
	result := initializers.DB.Create(&student)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "could not sign you up",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "student enrolled",
	})
}
func Login(ctx *gin.Context){
	var body struct{
		Roll string
		Password string
	}
	if ctx.Bind(&body)!=nil{
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":"could not read body",
		})
	}
	var student models.Student
	initializers.DB.First(&student, "roll=?", body.Roll)
	if student.ID==0{
		ctx.JSON(http.StatusNotFound, gin.H{
				"message":"You are not enrolled in this academy",
			})
		return
	}
	err:=bcrypt.CompareHashAndPassword([]byte(body.Password), []byte(student.Password))
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid password",
		})
		return
	}
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":student.ID,
		"exp": time.Now().Add(time.Hour*24*30).Unix(),
	})
	tokenString, err:= token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err!=nil{
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error":"could not authorize you",
		})
		return
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Enrollmenttoken", tokenString, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK,gin.H{})
}