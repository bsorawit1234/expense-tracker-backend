package routes

import (
	"github.com/bsorawit1234/expense-tracker-backend/controllers"
	"github.com/bsorawit1234/expense-tracker-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	public := router.Group("/api")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}

	protected := router.Group("/api")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/expenses", controllers.GetExpenses)
		protected.POST("/expenses", controllers.CreateExpense)
		protected.PUT("/expenses/:id", controllers.UpdateExpense)
		protected.DELETE("/expenses/:id", controllers.DeleteExpense)
	}

}
