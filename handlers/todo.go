//Хендлеры 
package handlers

import (
	"net/http"
	"todoAPI/db"
	"todoAPI/models"

	"github.com/gin-gonic/gin"
)



func GetTodos(c *gin.Context) {
	rows,err := db.DB.Query("SELECT id, title, completed FROM todos")
	if err !=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Ошибка при получении задачи",
		})
		return
	}

	defer rows.Close()

	var todos []models.Todo

	for rows.Next(){
		var t models.Todo
		 err := rows.Scan(&t.ID, &t.Title,&t.Completed)
		 if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":"Ошибка при чтении строк",
			})
			return
		}
		todos = append(todos, t)
	}
	c.JSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context){
	id := c.Param("id")
	var t models.Todo
	query := "SELECT id, title, completed FROM 	todos WHERE id = $1"
	err := db.DB.QueryRow(query, id).Scan(&t.ID, &t.Title, &t.Completed)


	if err!= nil{
		c.JSON(http.StatusNotFound, gin.H{
			"error":"Задача не найдена",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":t,
	})
}

func CreateTodo(c *gin.Context) {
	var input models.CreateTodoInput

	if err:= c.ShouldBindJSON(&input)
	err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Неверные данные. Убедитесь, что поля заполнены правильно",
		})		
		return
	}

	query := "INSERT INTO todos (title, completed) VALUES ($1,$2)"
	_,err := db.DB.Exec(query, input.Title, input.Completed)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при добовлении",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":"Задача добавлена ✅",
	})

}

func UpdateTodos(c *gin.Context) {
	id := c.Param("id")
	var updatedData models.CreateTodoInput

	if err:= c.ShouldBindJSON(&updatedData); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Некорректный формат JSON",
		})
		return
	}

	query := "UPDATE todos SET title=$1, completed=$2 WHERE id=$3"
	res, err := db.DB.Exec(query, updatedData.Title, updatedData.Completed, id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Ошибка при обновлении",
		})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0{
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Задача не найдена",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"Задача обновлена",
	})


}
func DeleteTodo(c *gin.Context){
	id:= c.Param("id")
	query := "DELETE FROM todos WHERE id = $1"

	result,err := db.DB.Exec(query,id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Ошибка при удалении",
		})
		return
	}

	rowsAffected, _ :=result.RowsAffected()
	if rowsAffected == 0{
		c.JSON(http.StatusNotFound, gin.H{
			"error":"Задана не найдена",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"Задача удалена",
	})


	
}
