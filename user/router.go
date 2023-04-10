package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(route *gin.Engine, db *gorm.DB) {
	userRepository := NewRepository(db)
	serviceRepository := NewService(userRepository)
	userHandler := NewBookHandler(serviceRepository)
	user := route.Group("/users")
	user.GET("", userHandler.GetUsers)
	user.POST("")

}
