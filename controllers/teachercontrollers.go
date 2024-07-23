package controllers

import (
	"net/http"

	initializers "github.com/Shreedeepdatta/rankandmarks/Initializers"
	"github.com/Shreedeepdatta/rankandmarks/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUpTeacher(ctx *gin.Context){
	var body struct{
		Name string
		Subject string
		Qualifications string
		Experience string
		Password string
	}
	if ctx.Bind(&body)!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"failed to read body",
		})
		return
	}
	hash, err :=bcrypt.GenerateFromPassword([]byte(body.Password),10)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"Failed to register password",
		})
		return
	}
	teacher:=models.Teacher{Name: body.Name, Subject: body.Subject, Qualifications: body.Qualifications, Experience: body.Experience, Password: string(hash)}
	result:=initializers.DB_TEACHER.Create(&teacher)
	if result.Error!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"Could not Sign you up",
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message":"Teacher added",
	})
}