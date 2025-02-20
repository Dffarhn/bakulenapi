package v1

import (
	"fmt"
	"io"
	"net/http"

	service "github.com/Dffarhn/bakulenapi/internal/services"
	"github.com/Dffarhn/bakulenapi/pkg/utils"
	_ "image/jpeg"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		UserService: service.NewUserService(),
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	// Retrieve the userId from the context
	userID, exists := c.Get("userId")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User ID not found in context")
		return
	}

	// Convert userID to string (it may have been stored as an interface{})
	userIDStr, ok := userID.(string)
	if !ok {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid User ID format")
		return
	}

	// Now you can use userIDStr to fetch the user data
	user, err := h.UserService.GetUser(userIDStr)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User retrieved successfully", user)
}

//update user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	// Retrieve the userId from the context
	userID, exists := c.Get("userId")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User ID not found in context")
		return
	}

	// Convert userID to string (it may have been stored as an interface{})
	userIDStr, ok := userID.(string)
	if !ok {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid User ID format")
		return
	}

	// Prepare a map to hold the fields to update
	data := make(map[string]interface{})

	// Check if username is provided
	if name := c.PostForm("name"); name != "" {
		data["name"] = name
	}

	// Check if profile picture is provided
	file, _, err := c.Request.FormFile("profile_picture")
	if err == nil {
		// Convert and upload the image if a profile picture is provided
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error reading file: %v", err))
			return
		}

		// Upload the webp image
		imageURL, err := utils.UploadImage(fmt.Sprintf("%s.webp", utils.GenerateUniqueFilename("user")), fileBytes)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error uploading image: %v", err))
			return
		}
		data["profile_picture"] = imageURL
	}

	// Call the service to update user fields
	err = h.UserService.UpdateUser(userIDStr, data)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error updating user: %v", err))
		return
	}

	// Return success response
	utils.SuccessResponse(c, http.StatusOK, "Profile updated successfully", nil)
}
