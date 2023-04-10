package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService Service
}

func NewBookHandler(bookService Service) *userHandler {
	return &userHandler{bookService}
}
func (h *userHandler) GetUsers(c *gin.Context){
	users,err := h.userService.FindAll()
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"data": users,
	})
}