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

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var request input.EmailRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errors := helper.FormatValidationErrorResponse(err)
		errMsg := gin.H{"errors": errors}

		response := helper.ApiResponse("Check email failed", http.StatusNotAcceptable, "error", errMsg)
		c.JSON(http.StatusNotAcceptable, response)
		return
	}

	isAvailable, err := h.service.IsEmailAvailable(request)
	if err != nil {
		errMsg := gin.H{"errors": err.Error()}

		response := helper.ApiResponse("Check email failed", http.StatusNotAcceptable, "error", errMsg)
		c.JSON(http.StatusNotAcceptable, response)
		return
	}

	data := gin.H{
		"is_available": isAvailable,
	}

	metaMsg := "Email has been registered"
	if isAvailable {
		metaMsg = "Email is available"
	}

	response := helper.ApiResponse(metaMsg, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
