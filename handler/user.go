package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verryp/go-startup/app/formatter"
	"github.com/verryp/go-startup/app/input"
	"github.com/verryp/go-startup/helper"
	"github.com/verryp/go-startup/service"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *userHandler {
	return &userHandler{s}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input input.UserRegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationErrorResponse(err)
		errMsg := gin.H{"errors": errors}

		response := helper.ApiResponse("Register account failed", http.StatusNotAcceptable, "error", errMsg)
		c.JSON(http.StatusNotAcceptable, response)
		return
	}

	newUser, err := h.service.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// TODO: generate token using jwt
	
	formatUser := formatter.FormatUser(newUser, "lalala")
	response := helper.ApiResponse("Account has been registered", http.StatusOK, "success", formatUser)
	c.JSON(http.StatusOK, response)
	
}