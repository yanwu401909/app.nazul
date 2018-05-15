package repo

import (
	. "app.nazul/errors"
	"app.nazul/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

var listTemp = "id, login_name, nick_name, sex, email, mobile, status, create_time"

func UsersList() (users []models.User, errs ApiError) {
	if result := CONN.Select(listTemp).Order("create_time desc").Find(&users); result.Error != nil {
		errs = NewErrorWithMessage(DB_ERROR, result.Error.Error())
	}
	return
}

func UsersPage(pageNo, pageSize int) (page models.Pageable, errs ApiError) {
	users := []models.User{}
	count := 0
	if result := CONN.Limit(pageSize).Offset((pageNo - 1) * pageSize).Select(listTemp).Order("create_time desc").Find(&users); result.Error != nil {
		errs = NewErrorWithMessage(DB_ERROR, result.Error.Error())
	} else {
		CONN.Model(&models.User{}).Select(listTemp).Count(&count)
		page.PageNo = pageNo
		page.PageSize = pageSize
		page.TotalRecord = count
		page.Data = users
	}
	return
}

func FindUserByIdOrLoginName(idOrName string) (user models.User, errs ApiError) {
	if result := CONN.Where("id = ?", idOrName).Or("login_name = ?", idOrName).First(&user); result.Error != nil {
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
