package repo

import (
	. "app.nazul/errors"
	"app.nazul/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

func UsersList() (users []models.User, errs ApiError) {
	if result := CONN.Select("login_name, nick_name, sex, email, mobile, status, create_time").Order("create_time").Find(&users); result.Error != nil {
		errs = NewErrorWithMessage(DB_ERROR, result.Error.Error())
	}
	return
}

func UsersPage(pageNo, pageSize int) (users []models.User, errs ApiError) {
	return
}

func FindUserByIdOrLoginName(idOrName string) (user models.User, errs ApiError) {
	if result := CONN.Where("id = ?", idOrName).Or("loginName = ?", idOrName).First(&user); result.Error != nil {
		errs = NewErrorWithMessage(DB_ERROR, result.Error.Error())
	}
	return
}

func SaveUser(user *models.User) (err ApiError) {
	if user.LoginName == "" || user.Password == "" || user.NickName == "" {
		err = NewError(PARAMS_ERROR)
	} else {
		id, _ := uuid.NewV4()
		user.Id = id.String()
		if result := CONN.Create(user); result.Error != nil {
			err = NewErrorWithMessage(DB_ERROR, result.Error.Error())
		}
	}
	return
}

func UpdateUser(id string, user *models.User) (err ApiError) {
	if id == "" || user.Password == "" || user.NickName == "" {
		err = NewError(PARAMS_ERROR)
	} else {
		if result := CONN.Model(&models.User{Id: id}).Updates(models.User{
			NickName: user.NickName,
			Email:    user.Email,
			Mobile:   user.Mobile,
			Sex:      user.Sex,
			Status:   user.Status,
		}); result.Error != nil {
			err = NewErrorWithMessage(DB_ERROR, result.Error.Error())
		}
	}
	return
}

func DeleteUser(id string) (err ApiError) {
	if result := CONN.Delete(&models.User{Id: id}); result.Error != nil {
		err = NewErrorWithMessage(DB_ERROR, result.Error.Error())
	}
	return
}
