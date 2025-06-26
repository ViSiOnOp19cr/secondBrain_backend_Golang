package repository

import (
	"github.com/ViSiOnOp19cr/secondBrain_backend_Golang/config"
	"github.com/ViSiOnOp19cr/secondBrain_backend_Golang/models"
)


func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func GetAllUsers(users *[]models.User) error {
	return config.DB.Find(users).Error
}

func GetUserByID(id uint, user *models.User) error {
	return config.DB.First(user, id).Error
}

func GetUserByEmail(email string, user *models.User) error {
	return config.DB.Where("email = ?", email).First(user).Error
}


func CreateContent(content *models.Content) error {
	return config.DB.Create(content).Error
}

func GetAllContent(userID uint, contents *[]models.Content) error {
	return config.DB.Where("user_id = ?", userID).Find(contents).Error
}

func GetContentByID(id uint, content *models.Content) error {
	return config.DB.First(content, id).Error
}

func UpdateContent(id uint, content *models.Content) error {
	return config.DB.Model(content).Where("id = ?", id).Updates(content).Error
}

func DeleteContent(id uint) error {
	return config.DB.Delete(&models.Content{}, id).Error
}
