package v1

import (
	"log"
	"net/http"

	service "github.com/Dffarhn/bakulenapi/internal/services"
	"github.com/Dffarhn/bakulenapi/pkg/utils"
	"github.com/gin-gonic/gin"
)

type GoogleAuthRequest struct {
	IDToken string `json:"idToken"`
}

// AuthHandler struct
type AuthHandler struct {
	AuthService *service.AuthService
}

// NewAuthHandler initializes AuthHandler
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		AuthService: service.NewAuthService(),
	}
}

// RegisterUser handles user registration
func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var req struct {
		Username        string `json:"username"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		RetypedPassword string `json:"retyped_password"`
		FCMToken        string `json:"fcm_token"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format")
		return
	}

	user, token, err := h.AuthService.Register(req.Email, req.Username, req.Password, req.FCMToken)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "User registered successfully", gin.H{
		"token": token,
		"user":  user,
	})
}

// LoginUser handles user login
func (h *AuthHandler) LoginUser(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("[ERROR] Invalid request format:", err)
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format")
		return
	}

	log.Printf("[INFO] User login attempt: Email - %s", req.Email)

	token, err := h.AuthService.Login(req.Email, req.Password)
	if err != nil {
		log.Println("[ERROR] Login failed:", err)
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	log.Println("[INFO] Login successful, token generated")

	utils.SuccessResponse(c, http.StatusOK, "Login successful", gin.H{
		"token": token,
	})
}

// made it when use google idtoken
// GoogleLogin handles Google login verification
func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	var req struct {
		IDToken string `json:"idToken"`
	}

	// Bind JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format")
		return
	}

	// Verify Google ID Token
	token, err := h.AuthService.VerifyGoogleIDToken(req.IDToken)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Respond with user details
	utils.SuccessResponse(c, http.StatusOK, "Login or Register successful", gin.H{
		"token": token,
	})
}

// store fcm token
func (h *AuthHandler) StoreFCMToken(c *gin.Context) {
	var req struct {
		FCMToken string `json:"fcm_token"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format")
		return
	}

	userID := c.GetString("userID")
	if err := h.AuthService.StoreFCMToken(userID, req.FCMToken); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "FCM token stored successfully", nil)
}
