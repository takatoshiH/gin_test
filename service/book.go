package service

import (
	"gin_test/model"
)

type BookService struct {
}

// create
func (BookService) SetBook(book *model.Book) error {
	_, err := DbEngine.Insert(book)
	return err
}

// index
//func (BookService) GetBookList() []model.Book {
//	books := make([]model.Book, 0)
//	err := Distinct("id", "title", "content").Limit(10, 0).Find(&books)
//	if err != nil {
//		panic(err)
//	}
//	return books
//}

// update
func (BookService) updateBook(newBook *model.Book) error {
	_, err := DbEngine.Id(newBook.Id).Update(newBook)
	return err
}

// delete
func (BookService) DeleteBook(id int) error {
	book := new(model.Book)
	_, err := DbEngine.Id(id).Delete(book)

	return err
}
