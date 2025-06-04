package routes

import(

	"todoAPI/handlers"
	"github.com/gin-gonic/gin"

)

func RegisterRoutes(router *gin.Engine){
	router.GET("/todos", handlers.GetTodos)
	router.GET("/todos/:id",handlers.GetTodoById)
	router.POST("/todos", handlers.CreateTodo)
	router.PUT("/todos/:id",handlers.UpdateTodos)
	router.DELETE("/todos/:id",handlers.DeleteTodo)
	
}
	
