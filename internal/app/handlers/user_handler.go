package handlers

import "github.com/gin-gonic/gin"

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// CreateUserHandler godoc
// @Summary user by id
// @Schemes
// @Description do ping
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user/ [post]
func (u *UserHandler) CreateUserHandler(c *gin.Context) {

}

// GetUserByIDHandler godoc
// @Summary user by id
// @Schemes
// @Description do ping
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user/ [get]
func (u *UserHandler) GetUserByIDHandler(c *gin.Context) {

}

// UpdateUserByIDHandler godoc
// @Summary user by id
// @Schemes
// @Description do ping
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user/ [put]
func (u *UserHandler) UpdateUserByIDHandler(c *gin.Context) {

}

// DeleteUserByIDsHandler godoc
// @Summary user by ids
// @Schemes
// @Description do ping
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user/ [delete]
func (u *UserHandler) DeleteUserByIDsHandler(c *gin.Context) {

}

// GetUserListHandler godoc
// @Summary list user
// @Schemes
// @Description do ping
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user/ [get]
func (u *UserHandler) GetUserListHandler(c *gin.Context) {

}
