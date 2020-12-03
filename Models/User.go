package Models

import (
	"github.com/getsentry/sentry-go"
	"theapp/Config"
)

func GetAllUser(user *[]User) (err error)  {
	if err = Config.DB.Find(user).Error; err != nil{
		return err
	}
	return nil
}

func CreateUser(user *User) (err error)  {
	if err = Config.DB.Create(user).Error; err != nil{
		return err
	}
	sentry.CaptureMessage("User created")
	return nil
}

func GetUserByID(user *User, id string)(err error)  {
	if err = Config.DB.Where("id=?", id).First(user).Error; err != nil{
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {
	Config.DB.Save(user)
	return nil
}

func DeleteUser(user *User, id string)(err error)  {
	Config.DB.Where("id=?", id).Delete(user)
	return nil
}
