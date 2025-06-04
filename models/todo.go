// Структура данных запроса
package models

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
type CreateTodoInput struct{
		Title string `json:"title" binding:"required"`
		Completed *bool `json:"completed" binding:"required"`
}

