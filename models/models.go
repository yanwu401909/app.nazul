package models

import (
	"time"
)

type Pageable struct {
	PageNo      int         `json:"pageNo,omitempty"`
	PageSize    int         `json:"pageSize,omitempty"`
	TotalRecord int         `json:"totalRecord,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

func (p *Pageable) TotalPage() int {
	if p.PageSize == 0 {
		return 0
	}
	if p.TotalRecord%p.PageSize > 0 {
		return p.TotalRecord/p.PageSize + 1
	} else {
		return p.TotalRecord % p.PageSize
	}
}

func (p *Pageable) HasPrevious() bool {
	return p.PageNo > 0
}

func (p *Pageable) HasNext() bool {
	return p.PageNo < p.TotalPage()
}

type ApiRequest struct {
	AuthToken string `json:"authToken,omitempty"`
	Platform  string `json:"platform,omitempty"`
	Agent     string `json:"agent,omitempty"`
	Version   string `json:"version,omitempty"`
}

type ApiResponse struct {
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}

type EchoResponse struct {
	ApiResponse
	// ResultCode int32 `json:"resultCode,omitempty"`
	// ResultMessage string `json:"resultMessage,omitempty"`
	Data string `json:"data,omitempty"`
}

type BooksListResponse struct {
	ApiResponse
	Data []Book `json:"data,omitempty"`
}
type UsersListResponse struct {
	ApiResponse
	Data []User `json:"data,omitempty"`
}

type UsersPageResponse struct {
	ApiResponse
	Data Pageable `json:"data,omitempty"`
}

type UserResponse struct {
	ApiResponse
	Data User `json:"data,omitempty"`
}

type BookResponse struct {
	ApiResponse
	Data Book `json:"data,omitempty"`
}

type Book struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Isbn        string `json:"isbn,omitempty"`
	Author      string `json:"author,omitempty"`
	Image       string `json:"image,omitempty"`
	Description string `json:"description,omitempty"`
}

type User struct {
	Id         string    `gorm:"primary_key" json:"id,omitempty"`
	LoginName  string    `gorm:"type:varchar(50);not null;" json:"loginName,omitempty"`
	Password   string    `gorm:"type:varchar(50);not null;" json:"password,omitempty"`
	NickName   string    `gorm:"type:varchar(50);not null;" json:"nickName,omitempty"`
	Sex        uint8     `gorm:"type:tinyint;default 1" json:"sex,omitempty"`
	Email      string    `gorm:"type:varchar(50);" json:"email,omitempty"`
	Mobile     string    `gorm:"type:varchar(20);" json:"mobile,omitempty"`
	Status     uint8     `gorm:"type:tinyint;default 0" json:"status,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
}
