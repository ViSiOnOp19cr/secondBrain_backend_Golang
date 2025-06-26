package repository

import (
	"github.com/ViSiOnOp19cr/secondBrain_backend_Golang/config"
	"github.com/ViSiOnOp19cr/secondBrain_backend_Golang/models"
)


func CreateUser(user *models.User) error{
	return config.DB.Create(user).Error
}

func GetAllUsers(users *[]models.User) error{
	return config.DB.Find(users).Error
}

func GetUserByID(id uint , user *models.User) error{
	return config.DB.First(user,id).Error
}

func CreateContent(content *models.Content) error{
	return config.DB.Create(content).Error
}

func GetAllContent(id uint, content *[]models.Content) error{
	return config.DB.Find(id,content).Error
}

func UpdateContent(id uint , content *[]models.Content) error{
	return config.DB.Save(content).Error
}
