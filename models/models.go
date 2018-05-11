package models

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

type BooksResponse struct {
	ApiResponse
	Data []Book `json:"data,omitempty"`
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
	Id         string `json:"id,omitempty"`
	LoginName  string `json:"loginName,omitempty"`
	Password   string `json:"password,omitempty"`
	NickName   string `json:"nickName,omitempty"`
	Sex        uint8  `json:"sex,omitempty"`
	Email      string `json:"email,omitempty"`
	Mobile     string `json:"mobile,omitempty"`
	Status     uint8  `json:"status,omitempty"`
	CreateTime int32  `json:"createTime,omitempty"`
}
