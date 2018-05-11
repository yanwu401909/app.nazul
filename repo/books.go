package repo

import (
	"math/rand"
	"strconv"

	. "app.nazul/errors"
	"app.nazul/mock"
	"app.nazul/models"
)

func GetBooks() (books []models.Book, err ApiError) {
	books = mock.BOOKS_DATA
	return
}

func GetBook(id string) (book models.Book, err ApiError) {
	for _, item := range mock.BOOKS_DATA {
		if item.Id == id {
			book = item
			return
		}
	}
	err = NewError(NOTEXIST_ERROR)
	return
}

func SaveBook(book models.Book) (err ApiError) {
	if book.Title == "" || book.Isbn == "" {
		err = NewError(PARAMS_ERROR)
	}
	book.Id = strconv.Itoa(rand.Intn(100000) + 10)
	mock.BOOKS_DATA = append(mock.BOOKS_DATA, book)
	return
}

func UpdateBook(id string, book models.Book) (err ApiError) {
	for index, item := range mock.BOOKS_DATA {
		if item.Id == id {
			if book.Title == "" || book.Isbn == "" {
				err = NewError(PARAMS_ERROR)
			} else {
				mock.BOOKS_DATA = append(mock.BOOKS_DATA[:index], mock.BOOKS_DATA[index+1:]...)
				book.Id = id
				mock.BOOKS_DATA = append(mock.BOOKS_DATA, book)
			}
			break
		}
	}
	return
}

func DeleteBook(id string) (err ApiError) {
	for index, item := range mock.BOOKS_DATA {
		if item.Id == id {
			mock.BOOKS_DATA = append(mock.BOOKS_DATA[:index], mock.BOOKS_DATA[index+1:]...)
			break
		}
	}
	return
}
