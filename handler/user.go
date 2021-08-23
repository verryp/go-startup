package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verryp/go-startup/app/formatter"
	"github.com/verryp/go-startup/app/input"
	"github.com/verryp/go-startup/helper"
	"github.com/verryp/go-startup/service"
)

type userHandler struct {
	service     service.UserService
	authService service.AuthService
}

func NewUserHandler(s service.UserService, authService service.AuthService) *userHandler {
	return &userHandler{s, authService}
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

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatUser := formatter.FormatUser(newUser, token)
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

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Upload avatar failed", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 1 //TODO: get from jwt
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Upload avatar failed", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Upload avatar failed", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Avatar successfully uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input input.LoginRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationErrorResponse(err)
		errMsg := gin.H{"errors": errors}

		response := helper.ApiResponse("Login failed", http.StatusUnauthorized, "error", errMsg)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	userLogged, err := h.service.Login(input)
	if err != nil {
		errMsg := gin.H{"errors": err.Error()}

		response := helper.ApiResponse("Login failed", http.StatusUnauthorized, "error", errMsg)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	token, err := h.authService.GenerateToken(userLogged.ID)
	if err != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatUser := formatter.FormatUser(userLogged, token)
	response := helper.ApiResponse("Successfully loggedin", http.StatusOK, "success", formatUser)
	c.JSON(http.StatusOK, response)
}
