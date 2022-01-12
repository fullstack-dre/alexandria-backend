package handler

import (
	"alexandria/helper"
	"alexandria/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (handler *userHandler) RegisterUser(context *gin.Context) {
	var input user.UserInput

	err := context.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to register user due to bad inputs",
			http.StatusUnprocessableEntity,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := handler.service.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to register user due to server error",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"User successfully registered!",
		http.StatusOK,
		"success",
		newUser,
	)

	context.JSON(http.StatusOK, response)
}