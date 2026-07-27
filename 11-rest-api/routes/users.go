package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func HandleSignupPOST(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to parse request data",
		})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save user",
		})
		return
	}

	// happy path
	context.JSON(http.StatusCreated, gin.H{
		"message": "Signup complete",
	})
}

func HandleLoginPOST(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to parse request data",
		})
		return
	}

	// now we need to retrieve the user based on JUST email
	userFromDB, err := models.GetUserByEmail(user.Email)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Password or Email",
		})
		return
	}

	isPasswordCorrect := utils.DoesHashMatchTerm(user.Password, userFromDB.Password)

	if !isPasswordCorrect {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Password or Email",
		})
		return
	}

	token, err := utils.GenerateToken(userFromDB.Email, userFromDB.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid Password or Email",
		})
		return
	}
	// happy path
	context.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})
}
