package handler

import (
	"fmt"
	"net/http"

	"houshold-app/internal/model"
	"houshold-app/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (userHandler *UserHandler) CreateUser(context *gin.Context) {
	var user model.CreateUser

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdUser, err := userHandler.UserService.Save(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, createdUser)
}

func (userHandler *UserHandler) UpdateUser(context *gin.Context) {
	var updateUser model.UpdatedUser

	if err := context.ShouldBindJSON(&updateUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := getIdFromContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	updatedUser, err := userHandler.UserService.Update(id, &updateUser)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, updatedUser)
}

func getIdFromContext(context *gin.Context) (int, error) {
	idParam := context.Param("id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
