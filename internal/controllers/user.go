package controllers

import (
	"errors"
	"fmt"
	"try/internal/models"
)

type UserController struct {
	model *models.UserModel
}

func NewUserController(m *models.UserModel) *UserController {
	return &UserController{
		model: m,
	}
}

func (uc *UserController) Login() (models.User, error) {
	var email, password string
	fmt.Print("Masukkan email ")
	fmt.Scanln(&email)
	fmt.Print("Masukkan password ")
	fmt.Scanln(&password)
	result, err := uc.model.Login(email, password)
	if err != nil {
		return models.User{}, errors.New("terjadi masalah ketika login")
	}
	return result, nil
}

func (uc *UserController) Register() (models.User, error) {
	var newData models.User
	fmt.Print("Masukkan Nama ")
	fmt.Scanln(&newData.Name)
	fmt.Print("Masukkan Password ")
	fmt.Scanln(&newData.Password)
	fmt.Print("Masukkan Email ")
	fmt.Scanln(&newData.Email)
	fmt.Print("Masukkan HP ")
	fmt.Scanln(&newData.Phone)
	result, err := uc.model.Register(newData)
	if err != nil && !result {
		return models.User{}, errors.New("terjadi masalah ketika registrasi")
	}
	return newData, nil
}

func (uc *UserController) updatekegiatan() (models.User, error) {
	var upKegiatan models.User
	fmt.Print("Masukkan Nama ")
	fmt.Scanln(&upKegiatan.Name)
	fmt.Print("Masukkan Password ") 
	fmt.Scanln(&upKegiatan.Password)
	fmt.Print("Masukkan Email ")
	fmt.Scanln(&upKegiatan.Email)
	fmt.Print("Masukkan HP ")
	fmt.Scanln(&upKegiatan.Phone)
	result, err := uc.model.Register(upKegiatan)
	if err != nil && !result {
		return models.User{}, errors.New("terjadi masalah ketika Update kegiatan")
	}
	return upKegiatan, nil
}

func (uc *UserController) hapusKegiatan() (models.User, error) {
	var delKegiatan models.User
	fmt.Print("Masukkan Nama ")
	fmt.Scanln(&delKegiatan.Name)
	fmt.Print("Masukkan Password ") 
	fmt.Scanln(&delKegiatan.Password)
	fmt.Print("Masukkan Email ")
	fmt.Scanln(&delKegiatan.Email)
	fmt.Print("Masukkan HP ")
	fmt.Scanln(&delKegiatan.Phone)
	result, err := uc.model.Register(delKegiatan)
	if err != nil && !result {
		return models.User{}, errors.New("terjadi masalah ketika Delete kegiatan")
	}
	return delKegiatan, nil
}
