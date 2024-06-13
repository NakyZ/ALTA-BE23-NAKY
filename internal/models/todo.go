package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Activity string
	Mark     bool
	Owner    uint
}

type TodoModel struct {
	db *gorm.DB
}

func NewTodoModel(connection *gorm.DB) *TodoModel {
	return &TodoModel{
		db: connection,
	}
}

func (tm *TodoModel) AddTodo(newData Todo) (Todo, error) {
	newData.Mark = false
	err := tm.db.Create(&newData).Error
	if err != nil {
		return Todo{}, err
	}

	return newData, nil
}

func (tm *TodoModel) AddTodo1(upKegiatan Todo) (Todo, error) {
	upKegiatan.Mark = false
	err := tm.db.Create(&upKegiatan).Error
	if err != nil {
		return Todo{}, err
	}

	return upKegiatan, nil
}

func (tm *TodoModel) AddTodo2(delKegiatan Todo) (Todo, error) {
	delKegiatan.Mark = false
	err := tm.db.Create(&delKegiatan).Error
	if err != nil {
		return Todo{}, err
	}

	return delKegiatan, nil
}

func (tm *TodoModel) AddTodo3(dftKegiatan Todo) (Todo, error) {
	dftKegiatan.Mark = false
	err := tm.db.Create(&dftKegiatan).Error
	if err != nil {
		return Todo{}, err
	}

	return dftKegiatan, nil
}