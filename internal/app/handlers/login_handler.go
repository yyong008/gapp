package handlers

import (
	"gapp1/internal/app/services"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	loginSerivce *services.LoginService
}

func NewLoginHandler(loginSerivce *services.LoginService) *LoginHandler {
	return &LoginHandler{
		loginSerivce: loginSerivce,
	}
}

// @BasePath /api/v1/user

// Login godoc
// @Summary login
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /login [post]
func (h *LoginHandler) Login(c *gin.Context) {
	// l, _ = h.loginSerivce.Login()
	c.JSON(200, gin.H{"message": "pong"})
}

// Register godoc
// @Summary login
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /register [post]
func (h *LoginHandler) Register(c *gin.Context) {
	// l, _ = h.loginSerivce.Login()
	c.JSON(200, gin.H{"message": "pong"})
}

// Logout godoc
// @Summary login
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /logout [get]
func (h *LoginHandler) Logout(c *gin.Context) {
	// l, _ = h.loginSerivce.Login()
	c.JSON(200, gin.H{"message": "pong"})
}

// refresh_token godoc
// @Summary refresh_token
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /refresh_token [post]
func (h *LoginHandler) RefreshToken(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
