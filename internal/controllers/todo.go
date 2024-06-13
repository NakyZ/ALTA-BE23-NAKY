package controllers

import (
	"fmt"
	"try/internal/models"
)

type TodoController struct {
	model *models.TodoModel
}

func NewTodoController(m *models.TodoModel) *TodoController {
	return &TodoController{
		model: m,
	}
}

func (tc *TodoController) AddTodo(id uint) (bool, error) {
	var newData models.Todo
	fmt.Print("Masukkan Aktivitas ")
	fmt.Scanln(&newData.Activity)
	newData.Owner = id
	_, err := tc.model.AddTodo(newData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (tc *TodoController) AddTodo1(id uint) (bool, error) {
	var upKegiatan models.Todo
	fmt.Print("Update Data")
	fmt.Scanln(&upKegiatan.Activity)
	upKegiatan.Owner = id
	_, err := tc.model.AddTodo(upKegiatan)
	if err != nil {
		return false, err
	}
	return true,nil
}

func (tc *TodoController) AddTodo2(id uint) (bool, error) {
	var delKegiatan models.Todo
	fmt.Print("Delete Kegiatan")
	fmt.Scanln(&delKegiatan.Activity)
	delKegiatan.Owner = id
	_, err := tc.model.AddTodo(delKegiatan)
	if err != nil {
		return false, err
	}
	return true,nil
}